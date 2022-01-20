package executor

import (
	"context"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/viper"
)

type SgnClient struct {
	txrs *transactor.TransactorPool
}

func NewSgnClient() *SgnClient {
	txrs := newSgnTransactors()
	return &SgnClient{txrs}
}

func newSgnTransactors() *transactor.TransactorPool {
	encoding := app.MakeEncodingConfig()
	txrAddrs := viper.GetStringSlice(common.FlagSgnTransactors)
	chainId := viper.GetString(common.FlagSgnChainId)
	nodeUri := viper.GetString(common.FlagSgnNodeURI)
	home := viper.GetString(flags.FlagHome)
	log.Infof("Initializing sgn transactors with args: home %s, chainId %s, nodeuri %s, addrs %v", home, chainId, nodeUri, txrAddrs)
	txrs := transactor.NewTransactorPool(home, chainId, encoding.Amino, encoding.Codec, encoding.InterfaceRegistry)
	err := txrs.AddTransactors(
		nodeUri,
		viper.GetString(common.FlagSgnPassphrase),
		txrAddrs)
	if err != nil {
		log.Fatalf("failed to add transactors: %s", err.Error())
	}
	log.Infof("Initialized %d sgn transactors", len(txrAddrs))
	return txrs
}

func (c *SgnClient) GetExecutionContexts(chainIds ...uint64) ([]msgtypes.ExecutionContext, error) {
	qc := msgtypes.NewQueryClient(c.txrs.GetTransactor().CliCtx)
	var contractInfos []*commontypes.ContractInfo
	for _, chainId := range chainIds {
		contractInfos = append(contractInfos, &commontypes.ContractInfo{ChainId: chainId, Address: ""})
	}
	req := &msgtypes.QueryExecutionContextsRequest{
		ContractInfos: contractInfos,
	}
	res, err := qc.ExecutionContexts(context.Background(), req)
	if err != nil {
		log.Errorln("failed to query messages from sgn", err)
		return nil, err
	}
	return res.GetExecutionContexts(), nil
}

func (c *SgnClient) InitWithdraw(srcXferId []byte, nonce uint64) error {
	txr := c.txrs.GetTransactor()
	wdReq := &cbrtypes.WithdrawReq{
		XferId:       eth.Bytes2Hex(srcXferId),
		ReqId:        nonce,
		WithdrawType: cbrtypes.RefundTransfer,
	}
	wdReqBytes, err := wdReq.Marshal()
	if err != nil {
		return err
	}
	msg := &cbrtypes.MsgInitWithdraw{
		WithdrawReq: wdReqBytes,
		Creator:     txr.Key.GetAddress().String(),
	}
	_, err = cbrcli.InitWithdraw(txr, msg)
	return err
}

func (c *SgnClient) PollAndExecuteWithdraw(addr string, nonce uint64, chainId uint64, callback ExecuteWithdraw) {
	maxTries := 10
	for try := 1; try <= maxTries; try++ {
		log.Debugf("polling withdraw status (try %d/%d): addr %s, nonce %d, chainId %d", try, maxTries, addr, nonce, chainId)
		time.Sleep(6 * time.Second)
		// poll withdraw status until its status reaches WD_WAITING_FOR_LP
		detail, status, err := c.GetWithdrawStatus(addr, nonce, chainId)
		if err != nil {
			log.Errorf("failed to get withdraw status (addr %x, nonce %d, chainId %d): %s", addr, nonce, chainId, err.Error())
			return
		}
		if status != cbrtypes.WithdrawStatus_WD_WAITING_FOR_LP {
			continue
		}
		if status == cbrtypes.WithdrawStatus_WD_COMPLETED ||
			status == cbrtypes.WithdrawStatus_WD_DELAYED ||
			status == cbrtypes.WithdrawStatus_WD_FAILED {
			log.Warnf("withdraw status %v not executable, skipping", status)
			return
		}
		// prepare withdraw req info
		signers, powers, err := c.QueryChainSigners(chainId)
		if err != nil {
			log.Errorf("failed to query chain signers with chainId %d", chainId)
			return
		}
		wdOnchain, sortedSigs := detail.GetWdOnchain(), detail.GetSortedSigsBytes()

		// execute the withdraw req onchain
		err = callback(wdOnchain, sortedSigs, signers, powers)
		if err != nil {
			log.Errorf("failed to execute withdraw: %s", err.Error())
		}
		return
	}
}

func (c *SgnClient) GetWithdrawStatus(
	addr string, nonce uint64, chainId uint64) (*cbrtypes.WithdrawDetail, cbrtypes.WithdrawStatus, error) {

	txr := c.txrs.GetTransactor()
	req := &cbrtypes.QueryWithdrawLiquidityStatusRequest{
		SeqNum:  nonce,
		UsrAddr: addr,
	}
	res, err := cbrcli.QueryWithdrawLiquidityStatus(txr.CliCtx, req)
	if err != nil {
		return nil, cbrtypes.WithdrawStatus_WD_UNKNOWN, err
	}
	log.Debugf("withdraw status %v", res.GetStatus())

	return res.GetDetail(), res.GetStatus(), nil
}

func (c *SgnClient) QueryChainSigners(chainId uint64) (addrs []eth.Addr, powers []*big.Int, err error) {
	res, err := cbrcli.QueryChainSigners(c.txrs.GetTransactor().CliCtx, chainId)
	if err != nil {
		return
	}
	addrs, powers = res.GetAddrsPowers()
	return
}
