package relayer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// to be called by r.verifyUpdate
// decode event and check if it matches onchain
func (r *Relayer) verifyMsgbrEventUpdate(update *synctypes.PendingUpdate) (done, approve bool) {
	onchev := new(cbrtypes.OnChainEvent)
	err := onchev.Unmarshal(update.Data)
	if err != nil {
		log.Errorf("failed to unmarshal %x to onchain event msg", update.Data)
		return true, false
	}
	elog := new(ethtypes.Log)
	err = json.Unmarshal(onchev.Elog, elog)
	if err != nil {
		log.Errorf("failed to unmarshal %x to eth Log", onchev.Elog)
		return true, false
	}
	if elog == nil {
		log.Errorf("unmarshal %x to to nil", onchev.Elog)
		return true, false
	}

	cbrOneChain := r.cbrMgr[onchev.Chainid]
	if cbrOneChain == nil {
		log.Errorf("cbrMgr not finish initialization yet, updates from chain: %d", onchev.Chainid)
		return false, false
	}

	skip, reason := cbrOneChain.skipMsgbrEvent(onchev.Evtype, elog, r.Transactor.CliCtx)
	if skip {
		log.Debugf("skip msg event: chain %d addr %x tx %x, reason: %s", onchev.Chainid, elog.Address, elog.TxHash, reason)
		return true, false
	}

	logmsg := fmt.Sprintf("verify update %d", update.Id)

	switch onchev.Evtype {
	case msgtypes.MsgEventMessage:
		return cbrOneChain.verifyMessage(r.Transactor.CliCtx, elog, logmsg)
	case msgtypes.MsgEventMessageWithTransfer:
		return cbrOneChain.verifyMessageEventTransfer(r.Transactor.CliCtx, elog, logmsg)
	case msgtypes.MsgEventExecuted:
		return cbrOneChain.verifyMessageEventExecuted(r.Transactor.CliCtx, elog, logmsg)
	default:
		log.Errorf("%s. invalid type", logmsg)
		return true, false
	}
}

func (c *CbrOneChain) verifyMessage(cliCtx client.Context, eLog *ethtypes.Log, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.msgContract.ParseMessage(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	msgId, _ := msgtypes.NewMessage(ev, c.chainid)
	logmsg = fmt.Sprintf("%s. %s, msgId %x", logmsg, ev.PrettyLog(c.chainid), msgId)
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnf("%s, eth addrs blocked", logmsg)
		return true, false
	}

	// check on chain
	done, approve, msgLog := c.verifyEventLog(eLog, eth.ContractTypeMsgBus, msgtypes.MsgEventMessage, c.msgContract.GetAddr(), logmsg)
	if msgLog == nil {
		return done, approve
	}

	msgEv, err := c.msgContract.ParseMessage(*msgLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	// cmp ev and msgEv
	if !ev.Equal(msgEv) {
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, msgEv, msgEv.Raw, ev, ev.Raw)
		return true, false
	}
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyMessageEventTransfer(cliCtx client.Context, eLog *ethtypes.Log, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.msgContract.ParseMessageWithTransfer(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.PrettyLog(c.chainid))
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnf("%s, eth addrs blocked", logmsg)
		return true, false
	}

	// check on chain
	done, approve, msgLog := c.verifyEventLog(eLog, eth.ContractTypeMsgBus, msgtypes.MsgEventMessageWithTransfer, c.msgContract.GetAddr(), logmsg)
	if msgLog == nil {
		return done, approve
	}

	msgEv, err := c.msgContract.ParseMessageWithTransfer(*msgLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	// now cmp ev and msgEv
	if !ev.Equal(msgEv) {
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, msgEv, msgEv.Raw, ev, ev.Raw)
		return true, false
	}

	// check transfer from sgn
	srcBridgeType := c.getSrcBridgeType(ev.Bridge)
	// check: bridge is either liquidity bridge, peg src vault, or peg dst bridge
	if srcBridgeType == msgtypes.BRIDGE_TYPE_NULL {
		log.Errorln(logmsg, "unknown bridge type")
		return true, false
	}

	xferSender, xferDstChainId, refunded, err := getTransferInfo(cliCtx, ev.SrcTransferId, srcBridgeType)
	if err != nil {
		if !strings.Contains(err.Error(), "no info found") ||
			(c.mon.GetCurrentBlockNumber().Int64()-int64(ev.Raw.BlockNumber))*int64(c.blkInterval) > 120 {
			log.Debugf("%s. getTransferInfo err: %s", logmsg, err)
		}
		return false, false
	}

	if refunded {
		log.Debugf("%s. approve message because found refund", logmsg)
		return true, true
	}

	if ev.Sender != xferSender {
		log.Errorf("%s. transfer sender not match: %x", logmsg, xferSender)
		return true, false
	}
	if ev.DstChainId.Uint64() != xferDstChainId {
		log.Errorf("%s. transfer dst chain not match: %d", logmsg, xferDstChainId)
		return true, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyMessageEventExecuted(cliCtx client.Context, eLog *ethtypes.Log, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.msgContract.ParseExecuted(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.PrettyLog(c.chainid))

	// check on chain
	done, approve, msgLog := c.verifyEventLog(eLog, eth.ContractTypeMsgBus, msgtypes.MsgEventExecuted, c.msgContract.GetAddr(), logmsg)
	if msgLog == nil {
		return done, approve
	}

	msgEv, err := c.msgContract.ParseExecuted(*msgLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	// now cmp ev and msgEv
	if !ev.Equal(msgEv) {
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, msgEv, msgEv.Raw, ev, ev.Raw)
		return true, false
	}

	// check status
	executedStatus, err := c.msgContract.ExecutedMessages(&bind.CallOpts{}, ev.MsgId)
	if err != nil {
		log.Errorln("could not verify MessageEventExecuted", err)
		return true, false
	}
	if executedStatus == 0 {
		log.Errorln(logmsg, "executedStatus is invalid. got:", executedStatus, "expect > 0")
		return true, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func getTransferInfo(cliCtx client.Context, srcTransferId eth.Hash, srcBridgeType msgtypes.BridgeType) (eth.Addr, uint64, bool, error) {
	xferId := srcTransferId.String()
	switch srcBridgeType {
	case msgtypes.BRIDGE_TYPE_LIQUIDITY:
		req := &cbrtypes.QueryTransferStatusRequest{TransferId: []string{xferId}}
		res, err := cbrcli.QueryTransferStatus(cliCtx, req)
		if err != nil {
			return makeErr(err)
		}
		status := res.GetStatus()[xferId]
		if status.GetSgnStatus() == cbrtypes.XferStatus_UNKNOWN {
			return makeErr(fmt.Errorf("xfer no info found")) // same err msg with peg not found
		}
		if status.GetSgnStatus() != cbrtypes.XferStatus_SUCCESS && status.GetSgnStatus() != cbrtypes.XferStatus_OK_TO_RELAY {
			return foundRefund()
		}
		relay, err := cbrcli.QueryRelay(cliCtx, srcTransferId.Bytes())
		if err != nil {
			return makeErr(err)
		}
		relayOnChain := new(cbrtypes.RelayOnChain)
		err = relayOnChain.Unmarshal(relay.Relay)
		if err != nil {
			return makeErr(err)
		}
		return eth.Bytes2Addr(relayOnChain.Sender), relayOnChain.DstChainId, false, nil
	case msgtypes.BRIDGE_TYPE_PEG_VAULT:
		deposit, err := pegcli.QueryDepositInfo(cliCtx, srcTransferId.String())
		if err != nil {
			return makeErr(err)
		}
		if len(deposit.GetMintId()) == 0 {
			return foundRefund()
		}
		mint, err := pegcli.QueryMintInfo(cliCtx, eth.Bytes2Hash(deposit.GetMintId()).String())
		if err != nil {
			return makeErr(err)
		}
		mintOnChain := new(pegbrtypes.MintOnChain)
		err = mintOnChain.Unmarshal(mint.GetMintProtoBytes())
		if err != nil {
			log.Errorf("Unmarshal mintInfo.MintProtoBytes err %s", err)
			return makeErr(err)
		}
		return eth.Bytes2Addr(mintOnChain.GetDepositor()), mint.GetChainId(), false, nil
	case msgtypes.BRIDGE_TYPE_PEG_BRIDGE:
		burn, err := pegcli.QueryBurnInfo(cliCtx, srcTransferId.String())
		if err != nil {
			return makeErr(err)
		}
		if len(burn.GetWithdrawId()) == 0 {
			return foundRefund()
		}
		withdraw, err := pegcli.QueryWithdrawInfo(cliCtx, eth.Bytes2Hash(burn.GetWithdrawId()).String())
		if err != nil {
			return makeErr(err)
		}
		withdrawOnChain := new(pegbrtypes.WithdrawOnChain)
		err = withdrawOnChain.Unmarshal(withdraw.GetWithdrawProtoBytes())
		if err != nil {
			return makeErr(err)
		}
		return eth.Bytes2Addr(withdrawOnChain.GetBurnAccount()), withdraw.GetChainId(), false, nil
	}
	return makeErr(fmt.Errorf("unsupported transfer type (%v)", srcBridgeType))
}

func makeErr(err error) (eth.Addr, uint64, bool, error) {
	return eth.ZeroAddr, 0, false, err
}

func foundRefund() (eth.Addr, uint64, bool, error) {
	return eth.ZeroAddr, 0, true, nil
}
