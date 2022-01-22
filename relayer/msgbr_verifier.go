package relayer

import (
	"encoding/json"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgbrtypes "github.com/celer-network/sgn-v2/x/message/types"
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
		log.Debugf("skip msgbr event: %s, reason: %s", string(onchev.Elog), reason)
		return true, false
	}

	logmsg := fmt.Sprintf("verify update %d cbr chain %d", update.Id, onchev.Chainid)

	switch onchev.Evtype {
	case msgbrtypes.MsgEventMessage:
		return cbrOneChain.verifyMessage(r.Transactor.CliCtx, elog, logmsg)
	case msgbrtypes.MsgEventMessageWithTransfer:
		return cbrOneChain.verifyMessageEventTransfer(r.Transactor.CliCtx, elog, logmsg)
	case msgbrtypes.MsgEventExecuted:
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
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	done, approve, msgLog := c.verifyEventLog(eLog, eth.MsgBridge, msgbrtypes.MsgEventMessage, c.msgContract.Address, logmsg)
	if msgLog == nil {
		return done, approve
	}

	msgEv, err := c.msgContract.ParseMessage(*msgLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	// now cmp ev and msgEv
	if !ev.Equal(msgEv) {
		log.Errorln(logmsg, "ev not equal. got:", msgEv.String(), "expect:", ev.String())
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
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	done, approve, msgLog := c.verifyEventLog(eLog, eth.MsgBridge, msgbrtypes.MsgEventMessageWithTransfer, c.msgContract.Address, logmsg)
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
		log.Errorln(logmsg, "ev not equal. got:", msgEv.String(), "expect:", ev.String())
		return true, false
	}

	// check info in cbr
	return c.verifyTransferForMessageBus(cliCtx, ev)
}

func (c *CbrOneChain) verifyMessageEventExecuted(cliCtx client.Context, eLog *ethtypes.Log, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.msgContract.ParseExecuted(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	done, approve, msgLog := c.verifyEventLog(eLog, eth.MsgBridge, msgbrtypes.MsgEventExecuted, c.msgContract.Address, logmsg)
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
		log.Errorln(logmsg, "ev not equal. got:", msgEv.String(), "expect:", ev.String())
		return true, false
	}

	// check status
	executedStatus, err := c.msgContract.ExecutedMessages(&bind.CallOpts{}, ev.Id)
	if err != nil {
		log.Errorln("could not verify MessageEventExecuted", err)
		return true, false
	}
	if executedStatus == 0 {
		log.Errorln(logmsg, "executedStatus is invalid. got:", executedStatus, "expect > 0")
		return true, false
	}
	log.Infof("verify MessageEventExecuted success, %s", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyTransferForMessageBus(cliCtx client.Context, ev *eth.MessageBusMessageWithTransfer) (done, approve bool) {
	// check transfer from sgn
	transferType := c.getTransferType(ev.Bridge)
	// check: bridge is either liquidity bridge, peg src vault, or peg dst bridge
	if transferType != msgbrtypes.TRANSFER_TYPE_LIQUIDITY_SEND &&
		transferType != msgbrtypes.TRANSFER_TYPE_PEG_MINT &&
		transferType != msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW {
		log.Warnf("verifyTransferForMessageBus failed, unknown bridge type, msg chainId:%d, transfer bridge:%s", c.chainid, ev.Bridge)
		return true, false
	}

	senderInTx, dstChainIdInTx, foundRefund, err := getTransferInfo(cliCtx, ev.SrcTransferId, transferType)
	if err != nil {
		log.Warnf("cannot verify xfer msg (srcXferId %x) %s", ev.SrcTransferId, err.Error())
		return false, false
	}

	if foundRefund {
		log.Debugf("approve message (srcXferId %x) because found refund", ev.SrcTransferId)
		return true, true
	}

	log.Infof("verifying message with transfer (srcTransferId %x)", ev.SrcTransferId)
	return verifySenderAndDstChainId(ev.Sender, senderInTx, ev.DstChainId.Uint64(), dstChainIdInTx)
}

func getTransferInfo(cliCtx client.Context, srcTransferId eth.Hash, transferType msgbrtypes.TransferType) (eth.Addr, uint64, bool, error) {
	xferId := srcTransferId.String()
	switch transferType {
	case msgbrtypes.TRANSFER_TYPE_LIQUIDITY_SEND:
		req := &cbrtypes.QueryTransferStatusRequest{TransferId: []string{xferId}}
		res, err := cbrcli.QueryTransferStatus(cliCtx, req)
		if err != nil {
			return makeErr(err)
		}
		status := res.GetStatus()[xferId]
		if status.GetSgnStatus() == cbrtypes.XferStatus_UNKNOWN {
			return makeErr(fmt.Errorf("xfer status unknown"))
		}
		if status.GetSgnStatus() != cbrtypes.XferStatus_SUCCESS || status.GetSgnStatus() != cbrtypes.XferStatus_OK_TO_RELAY {
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
	case msgbrtypes.TRANSFER_TYPE_PEG_MINT:
		deposit, err := pegcli.QueryDepositInfo(cliCtx, srcTransferId.String())
		if err != nil {
			return makeErr(err)
		}
		if deposit.GetMintId() == nil {
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
	case msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		burn, err := pegcli.QueryBurnInfo(cliCtx, srcTransferId.String())
		if err != nil {
			return makeErr(err)
		}
		if burn.GetWithdrawId() == nil {
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
	return makeErr(fmt.Errorf("unsupported transfer type (%v)", transferType))
}

func verifySenderAndDstChainId(sender, senderInTx eth.Addr, dstChainId, dstChainIdInTx uint64) (done, approve bool) {
	if sender != senderInTx {
		log.Warnf("verifyTransferForMessageBus failed, msg.sender doesn't match sender of the src transfer, msg sender:%s, transfer sender:%s", sender, senderInTx)
		return true, false
	}
	// check: dstChainId matches dstChainId of the src transfer
	if dstChainId != dstChainIdInTx {
		log.Warnf("verifyTransferForMessageBus failed, dstChainId doesn't match dstChainId of the src transfer, msg dstChainId:%d, transfer dstChainId:%d", dstChainId, dstChainIdInTx)
		return true, false
	}
	return true, true
}

func makeErr(err error) (eth.Addr, uint64, bool, error) {
	return eth.ZeroAddr, 0, false, err
}

func foundRefund() (eth.Addr, uint64, bool, error) {
	return eth.ZeroAddr, 0, true, nil
}
