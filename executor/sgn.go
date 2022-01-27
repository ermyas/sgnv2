package executor

import (
	"context"
	"fmt"
	"math/big"
	"strings"
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
	pegbrcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
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

func (c *SgnClient) GetExecutionContexts(filters []*commontypes.ContractInfo) ([]msgtypes.ExecutionContext, error) {
	qc := msgtypes.NewQueryClient(c.txrs.GetTransactor().CliCtx)
	req := &msgtypes.QueryExecutionContextsRequest{
		ContractInfos: filters,
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

func (c *SgnClient) InitPegRefund(refId []byte) error {
	txr := c.txrs.GetTransactor()
	msg := &pegbrtypes.MsgClaimRefund{
		RefId:  eth.Bytes2Hex(refId),
		Sender: txr.Key.GetAddress().String(),
	}
	log.Infof("init peg refund (refId %x)", refId)
	_, err := pegbrcli.InitClaimRefund(txr, msg)
	return err
}

func (c *SgnClient) PollAndExecuteWithdraw(addr string, nonce uint64, chainId uint64, execute ExecuteRefund) error {
	for try := 1; try <= MaxPollingRetries; try++ {
		log.Debugf("polling withdraw status (try %d/%d): addr %s, nonce %d, chainId %d", try, MaxPollingRetries, addr, nonce, chainId)
		time.Sleep(PollingInterval)
		// poll withdraw status until its status reaches WD_WAITING_FOR_LP
		detail, status, err := c.GetWithdrawStatus(addr, nonce, chainId)
		if err != nil {
			return fmt.Errorf("failed to get withdraw status (addr %x, nonce %d, chainId %d): %s", addr, nonce, chainId, err.Error())
		}
		if status != cbrtypes.WithdrawStatus_WD_WAITING_FOR_LP {
			log.Debugf("withdraw status %v is not WAITING_FOR_LP yet", status)
			continue
		}
		if status == cbrtypes.WithdrawStatus_WD_COMPLETED ||
			status == cbrtypes.WithdrawStatus_WD_DELAYED ||
			status == cbrtypes.WithdrawStatus_WD_FAILED {
			log.Warnf("withdraw status %v not executable, skipping", status)
			return nil
		}
		// prepare withdraw req info
		signers, powers, err := c.QueryChainSigners(chainId)
		if err != nil {
			return fmt.Errorf("failed to query chain signers with chainId %d", chainId)
		}
		wdOnchain, sortedSigs := detail.GetWdOnchain(), detail.GetSortedSigsBytes()

		// execute the withdraw req onchain
		err = execute(wdOnchain, sortedSigs, signers, powers)
		if err != nil {
			return fmt.Errorf("failed to execute liq withdraw (usr %s, nonce %d, chainId %d)", addr, nonce, chainId)
		} else {
			return nil
		}
	}
	return fmt.Errorf("PollAndExecuteWithdraw max retry reached for user %s, nonce %d, chainId %d", addr, nonce, chainId)
}

func (c *SgnClient) PollAndExecutePegRefundMint(burnId []byte, chainId uint64, execute ExecuteRefund) error {
	cliCtx := c.txrs.GetTransactor().CliCtx
	for try := 1; try <= MaxPollingRetries; try++ {
		log.Debugf("polling ClaimRefund status (try %d/%d): burnId %x, chainId %d", try, MaxPollingRetries, burnId, chainId)
		time.Sleep(PollingInterval)
		res, err := cbrcli.QueryChainSigners(cliCtx, chainId)
		if err != nil {
			return fmt.Errorf("failed to query chain signers: %s", err.Error())
		}
		mintId, err := pegbrcli.QueryRefundClaimInfo(cliCtx, eth.Bytes2Hex(burnId))
		if err != nil {
			return fmt.Errorf("failed to query refund claim info for deposit (id %x): %s", burnId, err.Error())
		}
		mintInfo, err := pegbrcli.QueryMintInfo(cliCtx, mintId)
		if err != nil {
			if strings.Contains(err.Error(), pegbrtypes.ErrNoInfoFound.Error()) {
				log.Infof("peg mint info for burnId %x not found yet", burnId)
				continue
			}
			return fmt.Errorf("failed to query peg mint info: %s", err.Error())
		}
		pass, sigBytes := cbrtypes.ValidateSignatureQuorum(mintInfo.GetSignatures(), res.GetSortedSigners())
		if !pass {
			log.Infof("skip peg burn refund (burnId %x): not enough sigs yet", burnId)
			continue
		}
		signers, powers := res.GetAddrsPowers()
		err = execute(mintInfo.MintProtoBytes, sigBytes, signers, powers)
		if err != nil {
			return fmt.Errorf("failed to execute peg burn (burnId %x) refund mint: %s", burnId, err.Error())
		} else {
			return nil
		}
	}
	return fmt.Errorf("PollAndExecutePegRefundMint max retry reached for burnId %x chainId %d", burnId, chainId)
}

func (c *SgnClient) PollAndExecutePegRefundWithdraw(depositId []byte, chainId uint64, execute ExecuteRefund) error {
	cliCtx := c.txrs.GetTransactor().CliCtx
	for try := 1; try <= MaxPollingRetries; try++ {
		log.Debugf("polling ClaimRefund status (try %d/%d): depositId %x, chainId %d", try, MaxPollingRetries, depositId, chainId)
		time.Sleep(PollingInterval)
		res, err := cbrcli.QueryChainSigners(cliCtx, chainId)
		if err != nil {
			return fmt.Errorf("failed to query chain signers %s", err.Error())
		}
		withdrawId, err := pegbrcli.QueryRefundClaimInfo(cliCtx, eth.Bytes2Hex(depositId))
		if err != nil {
			return fmt.Errorf("failed to query refund claim info for deposit (id %x): %s", depositId, err.Error())
		}
		withdrawInfo, err := pegbrcli.QueryWithdrawInfo(cliCtx, withdrawId)
		if err != nil {
			if strings.Contains(err.Error(), pegbrtypes.ErrNoInfoFound.Error()) {
				log.Infof("peg withdraw info for depositId %x not found yet", depositId)
				continue
			}
			return fmt.Errorf("failed to query withdraw info %s", err.Error())
		}
		pass, sigBytes := cbrtypes.ValidateSignatureQuorum(withdrawInfo.GetSignatures(), res.GetSortedSigners())
		if !pass {
			log.Infof("skip peg deposit refund (depositId %X): not enough sigs yet", depositId)
			continue
		}
		signers, powers := res.GetAddrsPowers()
		err = execute(withdrawInfo.WithdrawProtoBytes, sigBytes, signers, powers)
		if err != nil {
			return fmt.Errorf("failed to execute peg deposit (depositId %x) refund withdraw: %s", depositId, err.Error())
		} else {
			return nil
		}
	}
	return fmt.Errorf("PollAndExecutePegRefundWithdraw max retry reached for depositId %x chainId %d", depositId, chainId)
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
