package ropstentest

import (
	"context"
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/celer-network/goutils/eth"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	sgneth "github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/tools/cbridge/cmd"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TestChainConfig struct {
	chainId      uint64
	endpoint     string
	contractAddr string
}

type TestOnChain struct {
	transactor   *eth.Transactor
	bridge       *sgneth.Bridge
	authAccount  *bind.TransactOpts
	ec           *ethclient.Client
	contractAddr common.Addr
}

const (
	kspwd           = "123456"
	gatewayUrl      = "34.222.237.63:8082"
	testTokenSymbol = "USDT"
)

var (
	onChainConfigs = []*TestChainConfig{
		{
			chainId:      3,
			endpoint:     "wss://ropsten.infura.io/ws/v3/bf6437ebf88d487abbff85ba975def88",
			contractAddr: "0x2c60A9874493d8fE9b314c73F9d2cBe73ae18FB1",
		},
		{
			chainId:      5,
			endpoint:     "wss://goerli.infura.io/ws/v3/bf6437ebf88d487abbff85ba975def88",
			contractAddr: "0xaB2692D14898cAc8F1D5E3FB5BC95cF84092A2A6",
		},
	}
	onChainMap    = map[uint64]*TestOnChain{}
	gatewayClient *GatewayClient

	// bot behavior config
	// user addr
	userAddr              = common.Hex2Addr("0x29B563951Ed0eB9Ae5C49692266E1fbc81445cfE")
	usdtTokenAddrOnChain3 = common.Hex2Addr("0x7Df0C9680B7493Cd41332176559B8E2bA7c2A355")
	usdtTokenAddrOnChain5 = common.Hex2Addr("0xf4b2cbc3ba04c478f0dc824f4806ac39982dce73")

	signer *eth.CelerSigner // sign req msg

	defaultConfig = ProcessConfig{
		senderAddr:   userAddr,
		receiverAddr: userAddr,
		srcChainId:   5,
		dstChainId:   3,
		token:        usdtTokenAddrOnChain5,
	}
)

type ProcessConfig struct {
	senderAddr   common.Addr
	receiverAddr common.Addr
	srcChainId   uint64
	dstChainId   uint64
	token        common.Addr
}

// TestMain is used to setup/teardown a temporary CockroachDB instance
// and run all the unit tests in between.
func TestMain(m *testing.M) {
	flag.Parse()
	log.Infof("start test frame")
	err := setUp()
	if err != nil {
		log.Fatalf("fail to setup env, err:%v", err)
		return
	}
	exitCode := m.Run() // run all unittests
	os.Exit(exitCode)
}

func setUp() error {
	ksPath := "./ks/account1.json"
	ksBytes, err := ioutil.ReadFile(ksPath)
	if err != nil {
		return err
	}
	signer, err = eth.NewSignerFromKeystore(string(ksBytes), "123456", big.NewInt(0))
	if err != nil {
		log.Errorf("fail to create relay node singer from key store, err:%v", err)
		return err
	}
	gatewayClient, err = NewGatewayAPI(gatewayUrl)
	if err != nil {
		return err
	}
	for _, cfg := range onChainConfigs {
		log.Infof("start init chain:%d", cfg.chainId)
		var ec *ethclient.Client
		ec, err = ethclient.Dial(cfg.endpoint)
		if err != nil {
			return err
		}
		var cbr *sgneth.Bridge
		cbr, err = sgneth.NewBridge(sgneth.Hex2Addr(cfg.contractAddr), ec)
		if err != nil {
			return err
		}
		var trans *ethutils.Transactor
		trans, err = ethutils.NewTransactor(
			string(ksBytes),
			kspwd,
			ec,
			big.NewInt(int64(cfg.chainId)),
			ethutils.WithBlockDelay(10),
			ethutils.WithPollingInterval(15*time.Second),
		)
		if err != nil {
			log.Fatalln("ReadFile err:", err)
			return err
		}
		var ksfAc *os.File
		var authAccount *bind.TransactOpts
		ksfAc, err = os.Open(ksPath)
		if err != nil {
			return err
		}
		authAccount, err = bind.NewTransactorWithChainID(ksfAc, kspwd, big.NewInt(int64(cfg.chainId)))
		if err != nil {
			return err
		}
		onChainMap[cfg.chainId] = &TestOnChain{
			transactor:   trans,
			bridge:       cbr,
			authAccount:  authAccount,
			ec:           ec,
			contractAddr: common.Hex2Addr(cfg.contractAddr),
		}
	}
	return nil
}

func TestGenTransferId(t *testing.T) {
	tid := getTransferId(defaultConfig.senderAddr,
		defaultConfig.receiverAddr, usdtTokenAddrOnChain5, big.NewInt(10),
		defaultConfig.srcChainId,
		defaultConfig.dstChainId,
		1634032921829)
	if tid != sgneth.Hex2Hash("CBB877C0BC985F817D713ECA5039BF4E6FFD867A60E307F6EB25E103E6E30E1E") {
		log.Fatalf("unexpected transfer id")
	}
}

func TestTransfer(t *testing.T) {
	transferAmount := big.NewInt(100000)
	estimateAmountReq := &webapi.EstimateAmtRequest{
		SrcChainId:  uint32(defaultConfig.srcChainId),
		DstChainId:  uint32(defaultConfig.dstChainId),
		TokenSymbol: testTokenSymbol,
		Amt:         transferAmount.String(),
		UsrAddr:     defaultConfig.senderAddr.String(),
	}
	estimateAmountResp, err := gatewayClient.EstimateAmt(context.Background(), estimateAmountReq)
	if err != nil {
		t.Fatalf("fail to estimateAmount, err:%v", err)
		return
	}

	onChain, foundOnChain := onChainMap[defaultConfig.srcChainId]
	if !foundOnChain {
		t.Fatalf("fail to found this chain transactor, chain id:%d", 5)
		return
	}

	nowTs := time.Now()
	nonce := common.TsMilli(nowTs)
	tid := getTransferId(defaultConfig.senderAddr,
		defaultConfig.receiverAddr, usdtTokenAddrOnChain5, transferAmount,
		defaultConfig.srcChainId,
		defaultConfig.dstChainId,
		nonce)

	log.Infof("transferId:%s", tid.String())
	tx, err := send(onChain, defaultConfig.senderAddr, defaultConfig.receiverAddr, usdtTokenAddrOnChain5, transferAmount, defaultConfig.dstChainId, nonce, estimateAmountResp.GetMaxSlippage())
	if err != nil {
		log.Fatalf("fail to send on contract, err:%v", err)
		return
	}

	markTransferRequest := &webapi.MarkTransferRequest{
		TransferId: tid.String(),
		SrcSendInfo: &webapi.TransferInfo{
			Chain: &webapi.Chain{
				Id: uint32(defaultConfig.srcChainId),
			},
			Token: &types.Token{
				Symbol: testTokenSymbol,
			},
			Amount: transferAmount.String(),
		},
		DstMinReceivedInfo: &webapi.TransferInfo{
			Chain: &webapi.Chain{
				Id: uint32(defaultConfig.dstChainId),
			},
			Token: &types.Token{
				Symbol: testTokenSymbol,
			},
			Amount: transferAmount.String(),
		},
		Addr:      defaultConfig.senderAddr.String(),
		SrcTxHash: tx.Hash().String(),
		Type:      webapi.TransferType_TRANSFER_TYPE_SEND,
	}

	_, err = gatewayClient.MarkTransfer(context.Background(), markTransferRequest)
	if err != nil {
		log.Errorf("fail to mark transfer")
		return
	}

	for n := 0; n <= 60; n++ {
		getStatusResp, getStatusRespErr := gatewayClient.GetTransferStatus(context.Background(), &webapi.GetTransferStatusRequest{
			TransferId: tid.String(),
		})
		if getStatusRespErr != nil {
			log.Errorf("fail to getStatusRespErr:%v", getStatusRespErr)
		} else if getStatusResp.GetStatus() == types.TransferHistoryStatus_TRANSFER_COMPLETED {
			log.Infof("transfer success:%s", getStatusResp.GetStatus())
			break
		} else if getStatusResp.GetStatus() == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED {
			log.Infof("fail to transfer, status:%s", getStatusResp.GetStatus().String())
			processRefundErr := processRefund(onChain, tid, defaultConfig.senderAddr, usdtTokenAddrOnChain5, transferAmount, defaultConfig.srcChainId)
			if processRefundErr != nil {
				log.Errorf("fail to refund, err:%v", processRefundErr)
			} else {
				log.Infof("success to refund")
			}
			break
		} else {
			log.Infof("this transfer not completed now, status:%s", getStatusResp.GetStatus().String())
		}
		time.Sleep(4 * time.Second)
		continue
	}
}

func processRefund(onChain *TestOnChain, tid common.Hash, receiver, token common.Addr, amount *big.Int, chainId uint64) error {
	// TODO need sign again
	nowTs := time.Now()
	sigMsg, err := signer.SignEthMessage(sgneth.ToPadBytes(common.TsMilli(nowTs)))
	if err != nil {
		log.Errorf("fail to sig for ping, err:%v", err)
		return err
	}
	withdrawLiquidityResp, withdrawLiquidityRespErr := gatewayClient.WithdrawLiquidity(context.Background(), &webapi.WithdrawLiquidityRequest{
		TransferId:   tid.String(),
		ReceiverAddr: receiver.String(),
		Amount:       amount.String(),
		TokenAddr:    token.String(),
		ChainId:      uint32(chainId),
		Reqid:        common.TsMilli(nowTs),
		Sig:          sigMsg,
	})
	if withdrawLiquidityRespErr != nil {
		return withdrawLiquidityRespErr
	}
	seqNum := withdrawLiquidityResp.SeqNum
	log.Infof("withdraw seq num:%d", seqNum)

	var wmsg []byte
	var sigs [][]byte
	var signers []common.Addr
	var powers []*big.Int
	for n := 0; n <= 10; n++ {
		time.Sleep(3 * time.Second)
		transferStatusResp, getTransferStatusRespErr := gatewayClient.GetTransferStatus(
			context.Background(),
			&webapi.GetTransferStatusRequest{TransferId: tid.String()})
		if getTransferStatusRespErr != nil {
			log.Errorf("fail to get transfer status, err:%v", getTransferStatusRespErr)
			continue
		}
		if transferStatusResp.Status != types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
			log.Infof("transfer status expect:%s, actual:%s, try again", types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED.String(), transferStatusResp.Status.String())
			continue
		}
		wmsg = transferStatusResp.WdOnchain
		sigs = transferStatusResp.SortedSigs
		for _, bytes := range transferStatusResp.Signers {
			signers = append(signers, sgneth.Bytes2Addr(bytes))
		}
		for _, bytes := range transferStatusResp.Powers {
			powers = append(powers, new(big.Int).SetBytes(bytes))
		}
		break
	}
	if len(wmsg) > 0 && len(sigs) > 0 && len(signers) > 0 && len(powers) > 0 {
		// try refund
		refundOp, refundErr := withdraw(onChain, wmsg, sigs, signers, powers)
		if refundErr != nil {
			return refundErr
		}
		log.Infof("success send refund, hash:%s", refundOp.Hash())
		for n := 0; n <= 60; n++ {
			time.Sleep(4 * time.Second)
			getStatusResp, getStatusRespErr := gatewayClient.GetTransferStatus(context.Background(), &webapi.GetTransferStatusRequest{
				TransferId: tid.String(),
			})
			if getStatusRespErr != nil {
				log.Errorf("fail to getStatusRespErr:%v", getStatusRespErr)
			} else if getStatusResp.GetStatus() == types.TransferHistoryStatus_TRANSFER_COMPLETED {
				log.Infof("transfer success:%s", getStatusResp.GetStatus())
				break
			} else if getStatusResp.GetStatus() == types.TransferHistoryStatus_TRANSFER_REFUNDED {
				log.Infof("success to refund transfer, status:%s", getStatusResp.GetStatus().String())
				break
			} else {
				log.Infof("this transfer not completed now, status:%s", getStatusResp.GetStatus().String())
			}
		}
	} else {
		return fmt.Errorf("fail to do refund, because we can not get the transfer status with initWithdraw")
	}

	return nil
}

func withdraw(onChain *TestOnChain, wmsg []byte, sigs [][]byte, signers []common.Addr, powers []*big.Int) (*ethtypes.Transaction, error) {
	return onChain.transactor.Transact(
		logTransactionStateHandler(fmt.Sprintf("receipt send withdraw")),
		func(ctr bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return onChain.bridge.Withdraw(opts, wmsg, sigs, signers, powers)
		},
		eth.WithTimeout(2*time.Minute), // wait at most 1 minute
		eth.WithAddGasGwei(2),
		eth.WithBlockDelay(1),
		eth.WithPollingInterval(5*time.Second),
	)
}

func send(onChain *TestOnChain, sender, receiver, token common.Addr, amount *big.Int, dstChainId, nonce uint64, maxSlippage uint32) (*ethtypes.Transaction, error) {
	err := approveAllowance(onChain, sender, token, onChain.contractAddr)
	if err != nil {
		return nil, err
	}
	return onChain.transactor.Transact(
		logTransactionStateHandler(fmt.Sprintf("receipt send")),
		func(ctr bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return onChain.bridge.Send(opts, receiver, token, amount, dstChainId, nonce, maxSlippage)
		},
		eth.WithTimeout(2*time.Minute), // wait at most 1 minute
		eth.WithAddGasGwei(2),
		eth.WithBlockDelay(1),
	)
}

func sendWithMined(onChain *TestOnChain, sender, receiver, token common.Addr, amount *big.Int, dstChainId, nonce uint64, maxSlippage uint32) (common.Hash, error) {
	log.Infof("do send")
	err := approveAllowance(onChain, sender, token, onChain.contractAddr)
	if err != nil {
		return common.Hash{}, err
	}
	receipt, err := onChain.transactor.TransactWaitMined(
		fmt.Sprintf("send:%d", time.Now().Unix()),
		func(ctr bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return onChain.bridge.Send(opts, receiver, token, amount, dstChainId, nonce, maxSlippage)
		},
		eth.WithTimeout(2*time.Minute), // wait at most 1 minute
		eth.WithAddGasGwei(2),
		eth.WithBlockDelay(1),
	)
	if err != nil {
		log.Infof("error for send: %+v", err)
		return common.Hash{}, err
	}
	log.Infof("receipt for send: %+v", receipt)
	return receipt.TxHash, nil
}

func TestAddLp(t *testing.T) {
	addLpAmount := big.NewInt(1000)
	removeLpAmount := big.NewInt(100)
	nowTs := time.Now()
	onChain, foundOnChain := onChainMap[defaultConfig.srcChainId]
	if !foundOnChain {
		t.Fatalf("fail to found this chain transactor, chain id:%d", 5)
		return
	}

	txOp, err := addLp(onChain, defaultConfig.senderAddr, usdtTokenAddrOnChain5, addLpAmount)
	if err != nil {
		log.Errorf("fail to add lp, err:%v", err)
		return
	}
	log.Infof("add lp, tx:%s", txOp.Hash().String())
	addLpReq := &webapi.MarkLiquidityRequest{
		LpAddr:    defaultConfig.senderAddr.String(),
		Amt:       addLpAmount.String(),
		TokenAddr: usdtTokenAddrOnChain5.String(),
		ChainId:   uint32(defaultConfig.srcChainId),
		SeqNum:    common.TsMilli(nowTs),
		TxHash:    txOp.Hash().String(),
		Type:      webapi.LPType_LP_TYPE_ADD,
	}

	_, err = gatewayClient.MarkLiquidity(context.Background(), addLpReq)
	if err != nil {
		log.Errorf("fail to mark lp, err:%v", err)
		return
	}
	log.Infof("success to mark lp")

	sigMsg, err := signer.SignEthMessage(sgneth.ToPadBytes(common.TsMilli(nowTs)))
	if err != nil {
		log.Errorf("fail to sig for ping, err:%v", err)
		return
	}
	withdrawLiquidityReq := &webapi.WithdrawLiquidityRequest{
		ReceiverAddr: defaultConfig.senderAddr.String(),
		Amount:       removeLpAmount.String(),
		TokenAddr:    usdtTokenAddrOnChain5.String(),
		ChainId:      uint32(defaultConfig.srcChainId),
		Reqid:        common.TsMilli(nowTs),
		Sig:          sigMsg,
	}

	withdrawLiquidityResp, err := gatewayClient.WithdrawLiquidity(context.Background(), withdrawLiquidityReq)
	if err != nil {
		log.Errorf("fail to wd lp, err:%v", err)
		return
	}

	wdSeqNum := withdrawLiquidityResp.GetSeqNum()
	log.Infof("wd seq num:%d", wdSeqNum)

	var qResp *webapi.QueryLiquidityStatusResponse
	var qRespErr error
	for n := 0; n <= 30; n++ {
		time.Sleep(2 * time.Second)
		queryLiquidityStatusReq := &webapi.QueryLiquidityStatusRequest{
			SeqNum:  wdSeqNum,
			LpAddr:  defaultConfig.senderAddr.String(),
			ChainId: uint32(defaultConfig.srcChainId),
			Type:    webapi.LPType_LP_TYPE_REMOVE,
		}
		qResp, qRespErr = gatewayClient.QueryLiquidityStatus(context.Background(), queryLiquidityStatusReq)
		if qRespErr != nil {
			log.Errorf("fail to wd lp, err:%v", qRespErr)
			continue
		}
		log.Infof("query this withdrawal status:%s", qResp.Status.String())
		if qResp.GetStatus() == types.LPHistoryStatus_LP_WAITING_FOR_LP {
			break
		}
	}

	if qResp.GetStatus() != types.LPHistoryStatus_LP_WAITING_FOR_LP {
		log.Warnf("status expect:%s, actual:%s", types.LPHistoryStatus_LP_WAITING_FOR_LP.String(), qResp.GetStatus().String())
		return
	} else {
		log.Infof("sgn sign is ready, try to call contract for withdraw")
	}

	var wmsg []byte
	var sigs [][]byte
	var signers []common.Addr
	var powers []*big.Int
	wmsg = qResp.WdOnchain
	sigs = qResp.SortedSigs
	for _, bytes := range qResp.Signers {
		signers = append(signers, sgneth.Bytes2Addr(bytes))
	}
	for _, bytes := range qResp.Powers {
		powers = append(powers, new(big.Int).SetBytes(bytes))
	}

	txOp, err = withdraw(onChain, wmsg, sigs, signers, powers)
	if err != nil {
		log.Errorf("fail to send wd lp on chain, err:%v", err)
		return
	}

	for n := 0; n <= 60; n++ {
		time.Sleep(2 * time.Second)
		queryLiquidityStatusReq := &webapi.QueryLiquidityStatusRequest{
			SeqNum:  wdSeqNum,
			LpAddr:  defaultConfig.senderAddr.String(),
			ChainId: uint32(defaultConfig.srcChainId),
			Type:    webapi.LPType_LP_TYPE_REMOVE,
		}
		qResp, qRespErr = gatewayClient.QueryLiquidityStatus(context.Background(), queryLiquidityStatusReq)
		if qRespErr != nil {
			log.Errorf("fail to wd lp, err:%v", qRespErr)
			continue
		}
		log.Infof("query this withdrawal status:%s", qResp.Status.String())
		if qResp.GetStatus() == types.LPHistoryStatus_LP_COMPLETED {
			break
		}
	}

	if qResp.GetStatus() != types.LPHistoryStatus_LP_COMPLETED {
		log.Warnf("status expect:%s, actual:%s", types.LPHistoryStatus_LP_COMPLETED.String(), qResp.GetStatus().String())
		return
	} else {
		log.Infof("withdraw complete")
	}
}

func addLp(onChain *TestOnChain, sender, token common.Addr, amount *big.Int) (*ethtypes.Transaction, error) {
	err := approveAllowance(onChain, sender, token, onChain.contractAddr)
	if err != nil {
		return nil, err
	}
	return onChain.transactor.Transact(
		logTransactionStateHandler(fmt.Sprintf("receipt add lp")),
		func(ctr bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return onChain.bridge.AddLiquidity(opts, token, amount)
		},
		eth.WithTimeout(2*time.Minute), // wait at most 1 minute
		eth.WithAddGasGwei(2),
		eth.WithBlockDelay(1),
		eth.WithPollingInterval(5*time.Second),
	)
}

func approveAllowance(onChain *TestOnChain, userAccount, tokenAddr, contractAddr common.Addr) error {
	erc20, err := sgneth.NewErc20(tokenAddr, onChain.ec)
	if err != nil {
		return err
	}
	curAllowance, err := erc20.Allowance(&bind.CallOpts{}, userAccount, contractAddr)
	if err != nil {
		return err
	}
	if curAllowance.Cmp(new(big.Int).Div(cmd.MaxUint256, big.NewInt(2))) < 0 {
		approveTx, approveErr := erc20.Approve(onChain.authAccount, contractAddr, cmd.MaxUint256)
		if approveErr != nil {
			return approveErr
		}
		_, approveReceiptErr := eth.WaitMinedWithTxHash(context.Background(),
			onChain.ec,
			approveTx.Hash().String(),
			eth.WithTimeout(1*time.Minute),
			eth.WithBlockDelay(2))

		if approveReceiptErr != nil {
			return approveReceiptErr
		}
	}
	return nil
}

func getTransferId(sender, receiver, token common.Addr, amount *big.Int, srcChainId, dstChainId, nonce uint64) common.Hash {
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"address", "address", "address", "uint256", "uint64", "uint64", "uint64"},
		// values
		[]interface{}{
			sender,
			receiver,
			token,
			amount.String(),
			fmt.Sprintf("%d", dstChainId),
			fmt.Sprintf("%d", nonce),
			fmt.Sprintf("%d", srcChainId),
		},
	)
	return common.Bytes2Hash(hash)
}

func logTransactionStateHandler(desc string) *eth.TransactionStateHandler {
	return &eth.TransactionStateHandler{
		OnMined: func(receipt *ethtypes.Receipt) {
			if receipt.Status == ethtypes.ReceiptStatusSuccessful {
				log.Infof("%s transaction %x succeeded", desc, receipt.TxHash)
			} else {
				log.Errorf("%s transaction %x failed", desc, receipt.TxHash)
			}
		},
		OnError: func(tx *ethtypes.Transaction, err error) {
			log.Errorf("%s transaction %x err: %s", desc, tx.Hash(), err)
		},
	}
}
