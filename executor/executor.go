package executor

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	ethtypes "github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/sgn"
	"github.com/celer-network/sgn-v2/executor/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

type Executor struct {
	dal         *DAL
	chains      *ChainMgr
	sgn         *sgn.SgnClient
	gateway     *GatewayClient
	wg          sync.WaitGroup
	contracts   []*types.ContractConfig
	parallelism int
	testMode    bool
	autoRefund  bool
}

func NewExecutor(dal *DAL, testMode bool) *Executor {
	var gateway *GatewayClient
	if !testMode {
		gateway = NewGatewayClient(viper.GetString(types.FlagGatewayGrpcUrl))
	}
	sgn := sgn.NewSgnClient(viper.GetString(types.FlagSgnGrpcUrl), testMode)
	chains := NewChainMgr(dal)

	contracts := []*types.ContractConfig{}
	err := viper.UnmarshalKey(types.FlagExecutorContracts, &contracts)
	if err != nil {
		log.Fatalln("failed to initialize contract filters", err)
	}
	if len(contracts) == 0 {
		log.Fatalln("empty executor contract filter")
	}
	log.Infoln("executor will submit execution for contracts:")
	for _, contract := range contracts {
		log.Infof("(chainId %d, addr %s, value %s)",
			contract.ChainId, contract.Address, contract.PayableValue)
	}
	autoRefundEnabled := viper.GetBool(types.FlagEnableAutoRefund)
	if autoRefundEnabled {
		log.Infoln("auto refund enabled")
	} else {
		log.Infoln("auto refund disabled")
	}
	return &Executor{
		dal:         dal,
		chains:      chains,
		sgn:         sgn,
		gateway:     gateway,
		contracts:   contracts,
		parallelism: 10,
		testMode:    testMode,
		autoRefund:  autoRefundEnabled,
	}
}

func (e *Executor) Start() {
	done := make(chan bool)
	go e.startFetchingExecCtxsFromSgn()
	go e.startProcessingExecCtxsFromDb()
	go e.chains.StartMonitoring()
	log.Info("executor started")
	<-done
}

func (e *Executor) startFetchingExecCtxsFromSgn() {
	log.Infoln("Start fetching execution contexts from SGN")
	for {
		time.Sleep(8 * time.Second)
		execCtxs, err := e.sgn.GetExecutionContexts(e.contracts)
		if err != nil {
			log.Errorln("failed to get messages", err)
			continue
		}
		if len(execCtxs) == 0 {
			continue
		}
		log.Tracef("Got %d execution contexts", len(execCtxs))
		execCtxsToSave := []*msgtypes.ExecutionContext{}
		for i := range execCtxs {
			execCtxsToSave = append(execCtxsToSave, &execCtxs[i])
		}
		e.dal.SaveExecutionContexts(execCtxsToSave)
	}
}

func (e *Executor) startProcessingExecCtxsFromDb() {
	log.Infoln("Start processing execution contexts from DB")
	for {
		time.Sleep(3 * time.Second)
		execCtxs, statuses := e.dal.GetExecutionContextsToExecute()
		if len(execCtxs) == 0 {
			continue
		}
		e.executeInParallel(execCtxs, statuses)
	}
}

func (e *Executor) executeInParallel(execCtxs []*msgtypes.ExecutionContext, statuses []types.ExecutionStatus) {
	// X workers processing messages at once
	// each worker is responsible for a chunk of the msgs
	chunkSize := len(execCtxs) / e.parallelism
	if chunkSize < 1 {
		chunkSize = 1
	}
	log.Debugf("Executing %d messages with parallelism %d, chunk size %d", len(execCtxs), e.parallelism, chunkSize)
	workerNum := 0
	for i := 0; i < len(execCtxs); i += chunkSize {
		end := i + chunkSize
		if end > len(execCtxs) {
			end = len(execCtxs)
		}
		chunk := execCtxs[i:end]
		statusChunk := statuses[i:end]
		e.wg.Add(1)
		log.Debugf("Worker #%d executing messages [%d:%d]", workerNum, i, end)
		go e.execute(chunk, statusChunk)
		workerNum++
	}
	// block until the current round of msgs are all done executing
	e.wg.Wait()
}

func (e *Executor) execute(execCtxs []*msgtypes.ExecutionContext, statuses []types.ExecutionStatus) {
	defer e.wg.Done()
	for i, execCtx := range execCtxs {
		e.routeExecution(execCtx, statuses[i])
	}
}

func (e *Executor) routeExecution(execCtx *msgtypes.ExecutionContext, status types.ExecutionStatus) {
	// same chain ids mean it's a refund
	if execCtx.Message.SrcChainId == execCtx.Message.DstChainId {
		if !e.autoRefund {
			log.Debugf("skip executing refund for message (id %x) because enable_auto_refund is off", execCtx.MessageId)
			return
		}
		if status == types.ExecutionStatus_Init_Refund_Executed {
			e.executeMsgWithTransferRefund(execCtx)
		} else if status == types.ExecutionStatus_Unexecuted {
			err := e.routeInitRefund(execCtx)
			if err != nil {
				log.Errorln("init refund failed", err)
				Dal.UpdateStatus(execCtx.MessageId, types.ExecutionStatus_Init_Refund_Failed)
			}
		}
		return
	}
	// handle normal execution
	switch execCtx.Message.GetTransferType() {
	case msgtypes.TRANSFER_TYPE_NULL:
		e.executeMsgNoTransfer(execCtx)
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_SEND:
		e.executeMsgWithTransfer(execCtx)
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		e.executeMsgWithTransfer(execCtx)
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		e.executeMsgWithTransfer(execCtx)
	default:
		log.Errorf("normal execution not possible for message (id %x) transfer type %v, status %d",
			execCtx.MessageId, execCtx.Message.GetTransferType(), status)
	}
}

func (e *Executor) routeInitRefund(execCtx *msgtypes.ExecutionContext) error {
	switch execCtx.Message.GetTransferType() {
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		return e.executeRefundWithdraw(execCtx)
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		return e.initAndExecutePegRefundMint(execCtx)
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		return e.initAndExecutePegRefundWithdraw(execCtx)
	default:
		return fmt.Errorf("init refund not possible for message (id %x) transfer type %v",
			execCtx.MessageId, execCtx.Message.GetTransferType())
	}
}

func (e *Executor) executeMsgNoTransfer(execCtx *msgtypes.ExecutionContext) {
	message := &execCtx.Message
	chain, err := e.chains.GetChain(message.DstChainId)
	if err != nil {
		log.Errorln("cannot executeMessage", err)
		return
	}
	id := execCtx.Message.ComputeMessageIdNoTransfer()
	route := ethtypes.MessageBusReceiverRouteInfo{
		Sender:     ethtypes.Hex2Addr(message.Sender),
		Receiver:   ethtypes.Hex2Addr(message.Receiver),
		SrcChainId: message.SrcChainId,
	}

	err = Dal.UpdateStatus(id, types.ExecutionStatus_Executing)
	if err != nil {
		log.Errorln("cannot executeMessage", err)
		return
	}
	msg, sigs, signers, powers, err := e.getMsgSignInfo(execCtx)
	if err != nil {
		log.Errorf("failed to query chain signers with chainId %d", execCtx.Message.DstChainId)
		return
	}
	contractConf, found := types.GetContractConfig(e.contracts, execCtx.Message.Receiver)
	if !found {
		log.Errorf("message receiver (address %s) not found in contract configs", execCtx.Message.Receiver)
		return
	}
	log.Infof("executing msg (id %x)...", id)
	tx, err := chain.Transactor.Transact(
		getTransactionHandler(id, execCtx, "executeMsg"),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			// Executor can be optionally configured to include a payable value for message execution.
			// This value acts as message fee and is needed when calling executeMessage results in
			// sending another message.
			setValue(opts, contractConf.PayableValue)
			return chain.MsgBus.ExecuteMessage(opts, msg, route, sigs, signers, powers)
		})
	if err != nil {
		log.Errorf("cannot execute msg (id %x): %s", id, err.Error())
		err = Dal.UpdateStatus(id, types.ExecutionStatus_Failed)
		if err != nil {
			log.Errorf("cannot update execution_context (id %x): %s", id, err.Error())
		}
		return
	}
	log.Infof("executed msg (id %x): txhash %x", id, tx.Hash())
}

func (e *Executor) initAndExecutePegRefundMint(execCtx *msgtypes.ExecutionContext) error {
	message := &execCtx.Message
	burnId := message.GetTransferRefId()

	var err error
	if e.testMode {
		err = e.sgn.InitPegRefund(burnId)
	} else {
		err = e.gateway.InitPegRefund(burnId)
	}
	if err != nil {
		return fmt.Errorf("failed to init claim refund: %s", err.Error())
	}
	chain, err := e.chains.GetChain(message.DstChainId)
	if err != nil {
		return fmt.Errorf("failed to initAndExecutePegRefundMint: %s", err.Error())
	}
	mintExecutor := chain.NewExecuteRefundHandler(execCtx.MessageId, chain.ExecutePegMint)
	return e.sgn.PollAndExecutePegRefundMint(burnId, message.DstChainId, mintExecutor)
}

func (e *Executor) initAndExecutePegRefundWithdraw(execCtx *msgtypes.ExecutionContext) error {
	message := &execCtx.Message
	depositId := message.GetTransferRefId()

	var err error
	if e.testMode {
		err = e.sgn.InitPegRefund(depositId)
	} else {
		err = e.gateway.InitPegRefund(depositId)
	}
	if err != nil {
		return fmt.Errorf("failed to init claim refund: %s", err.Error())
	}
	chain, err := e.chains.GetChain(message.DstChainId)
	if err != nil {
		return fmt.Errorf("failed to initAndExecutePegRefundWithdraw: %s", err.Error())
	}
	withdrawExecutor := chain.NewExecuteRefundHandler(execCtx.MessageId, chain.ExecutePegWithdraw)
	return e.sgn.PollAndExecutePegRefundWithdraw(depositId, message.DstChainId, withdrawExecutor)
}

func (e *Executor) executeRefundWithdraw(execCtx *msgtypes.ExecutionContext) error {
	receiver := execCtx.Message.Receiver
	nonce := execCtx.GetTransfer().GetSeqNum()
	chainId := execCtx.Message.DstChainId
	srcXferId := execCtx.Message.GetTransferRefId()

	log.Infof("execute refund withdrawal: srcTransferId %x, nonce %d", srcXferId, nonce)
	chain, err := e.chains.GetChain(chainId)
	if err != nil {
		return fmt.Errorf("failed to executeRefundWithdraw: %s", err.Error())
	}
	execHandler := chain.NewExecuteRefundHandler(execCtx.MessageId, chain.ExecuteLiqWithdraw)
	return e.sgn.PollAndExecuteWithdraw(receiver, nonce, chainId, execHandler)
}

func (e *Executor) executeMsgWithTransferRefund(execCtx *msgtypes.ExecutionContext) {
	chain, err := e.chains.GetChain(execCtx.Message.DstChainId)
	if err != nil {
		log.Errorln("cannot executeMsgWithTransferRefund", err)
		return
	}
	message := execCtx.Message
	id := execCtx.MessageId

	transfer := execCtx.GetTransfer()
	amount, _ := new(big.Int).SetString(transfer.Amount, 10)
	xfer := ethtypes.MessageBusReceiverTransferInfo{
		T:          uint8(message.GetTransferType()),
		Sender:     ethtypes.Hex2Addr(message.Sender),
		Receiver:   ethtypes.Hex2Addr(message.Receiver),
		Token:      ethtypes.Bytes2Addr(transfer.Token),
		Amount:     amount,
		Seqnum:     transfer.SeqNum,
		SrcChainId: message.SrcChainId,
		RefId:      ethtypes.Bytes2Hash(execCtx.Message.GetTransferRefId()),
	}

	err = Dal.UpdateStatus(id, types.ExecutionStatus_Executing)
	if err != nil {
		log.Errorf("cannot execute refund %s", err.Error())
		return
	}
	msg, sigs, signers, powers, err := e.getMsgSignInfo(execCtx)
	if err != nil {
		log.Errorf("failed to query chain signers with chainId %d", execCtx.Message.DstChainId)
		return
	}
	log.Infof("executing refund (id %x)...", id)
	execCtx.Message.PrettyLog()
	tx, err := chain.Transactor.Transact(
		getTransactionHandler(id, execCtx, "executeMsgWithTransferRefund"),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			return chain.MsgBus.ExecuteMessageWithTransferRefund(opts, msg, xfer, sigs, signers, powers)
		})
	if err != nil {
		log.Errorf("cannot execute refund (id %x): %s", id, err.Error())
		err = Dal.UpdateStatus(id, types.ExecutionStatus_Failed)
		if err != nil {
			log.Errorf("cannot update execution_context (id %x): %s", id, err.Error())
		}
		return
	}
	log.Infof("executed refund (id %x): txhash %x", id, tx.Hash())
}

func (e *Executor) isTransferReady(chain *Chain, execCtx *msgtypes.ExecutionContext) (ready bool) {
	dstTransferId := ethtypes.Bytes2Hash(execCtx.ComputeDstTransferId())
	var err error
	switch execCtx.Message.TransferType {
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_SEND:
		ready, err = chain.LiqBridge.Transfers(&bind.CallOpts{}, dstTransferId)
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		ready, err = chain.PegBridge.Records(&bind.CallOpts{}, dstTransferId)
	case msgtypes.TRANSFER_TYPE_PEG_MINT_V2:
		ready, err = chain.PegBridgeV2.Records(&bind.CallOpts{}, dstTransferId)
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		ready, err = chain.PegVault.Records(&bind.CallOpts{}, dstTransferId)
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW_V2:
		ready, err = chain.PegVaultV2.Records(&bind.CallOpts{}, dstTransferId)
	default:
		log.Panicf("unsupported transfer type %s", execCtx.Message.TransferType)
	}
	if err != nil {
		log.Errorf("[skip execution] failed to query on-chain transfer for message (id %x, transferType %s, dstTransferId %x)",
			execCtx.MessageId, execCtx.Message.TransferType, dstTransferId)
	}
	return
}

func (e *Executor) executeMsgWithTransfer(execCtx *msgtypes.ExecutionContext) {
	chain, err := e.chains.GetChain(execCtx.Message.DstChainId)
	if err != nil {
		log.Errorln("failed to get chain", err)
		return
	}
	message := execCtx.Message
	id := execCtx.MessageId
	log.Infof("executeMsgWithTransfer %x", message.TransferRefId)

	if !e.isTransferReady(chain, execCtx) {
		log.Infof("[skip execution] message (id %x) because trasnfer is not seen on dst chain yet", execCtx.MessageId)
		return
	}

	transfer := execCtx.GetTransfer()
	amount, _ := new(big.Int).SetString(transfer.Amount, 10)
	xfer := ethtypes.MessageBusReceiverTransferInfo{
		T:          uint8(message.GetTransferType()),
		Sender:     ethtypes.Hex2Addr(message.Sender),
		Receiver:   ethtypes.Hex2Addr(message.Receiver),
		Token:      ethtypes.Bytes2Addr(transfer.Token),
		Amount:     amount,
		Seqnum:     transfer.SeqNum,
		SrcChainId: message.SrcChainId,
		RefId:      ethtypes.Bytes2Hash(execCtx.Message.GetTransferRefId()),
	}
	contractConf, found := types.GetContractConfig(e.contracts, execCtx.Message.Receiver)
	if !found {
		log.Errorf("message receiver (address %s) not found in contract configs", execCtx.Message.Receiver)
		return
	}

	err = Dal.UpdateStatus(id, types.ExecutionStatus_Executing)
	if err != nil {
		log.Errorf("cannot execute xferMsg %s", err.Error())
		return
	}
	msg, sigs, signers, powers, err := e.getMsgSignInfo(execCtx)
	if err != nil {
		log.Errorf("failed to query chain signers with chainId %d", execCtx.Message.DstChainId)
		return
	}
	log.Infof("executing xferMsg (id %x)...", id)
	tx, err := chain.Transactor.Transact(
		getTransactionHandler(id, execCtx, "executeMsgWithXfer"),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			// Executor can be optionally configured to include a payable value for message execution.
			// This value acts as message fee and is needed when calling executeMessageWithTransfer results
			// in sending another message.
			setValue(opts, contractConf.PayableValue)
			return chain.MsgBus.ExecuteMessageWithTransfer(opts, msg, xfer, sigs, signers, powers)
		})
	if err != nil {
		log.Errorf("cannot execute xferMsg (id %x): %s", id, err.Error())
		err = Dal.UpdateStatus(id, types.ExecutionStatus_Failed)
		if err != nil {
			log.Errorf("cannot update execution_context (id %x): %s", id, err.Error())
		}
		return
	}
	log.Infof("executed xferMsg (id %x): txhash %x", id, tx.Hash())
}

func (e *Executor) getMsgSignInfo(execCtx *msgtypes.ExecutionContext) (msg []byte, sigs [][]byte, signers []ethtypes.Addr, powers []*big.Int, err error) {
	msg = execCtx.Message.Data
	sigs = execCtx.Message.GetSigBytes()
	chainSigners, err := e.sgn.GetChainSigners(execCtx.Message.DstChainId)
	if err != nil {
		return
	}
	signers, powers = chainSigners.GetAddrsPowers()
	return
}

func getTransactionHandler(id []byte, execCtx *msgtypes.ExecutionContext, logmsg string) *eth.TransactionStateHandler {
	return &eth.TransactionStateHandler{
		OnMined: func(receipt *gethtypes.Receipt) {
			log.Infof("%s: tx %x mined, status %v message id", logmsg, receipt.TxHash, receipt.Status)
			status := types.ExecutionStatus_Executed
			if receipt.Status == gethtypes.ReceiptStatusFailed {
				status = types.ExecutionStatus_Failed
			}
			Dal.UpdateStatus(id, status)
		},
		OnError: func(tx *gethtypes.Transaction, err error) {
			log.Errorf("%s error: txhash %s, err %v", logmsg, tx.Hash(), err)
			Dal.UpdateStatus(id, types.ExecutionStatus_Failed)
		},
	}
}

func setValue(opts *bind.TransactOpts, value string) {
	if len(value) > 0 {
		val, ok := new(big.Int).SetString(value, 10)
		if ok {
			opts.Value = val
		}
	}
}
