package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgbrtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

var msgEvNames = []string{
	msgbrtypes.MsgEventMessage,
	msgbrtypes.MsgEventMessageWithTransfer,
	msgbrtypes.MsgEventExecuted,
}

// sleep, check if syncer, if yes, go over cbr dbs to send msgbr tx
func (r *Relayer) doMsgbrSync(cbrMgr CbrMgr) {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalCbridge)) * time.Second
	log.Infoln("start process msgbr sync, interval:", interval)
	for {
		time.Sleep(interval)
		if !r.isSyncer() {
			continue
		}
		// find all events need to be sent out, batch into one msg
		msg := &synctypes.MsgProposeUpdates{
			Sender: r.Transactor.Key.GetAddress().String(),
		}

		var updatesBytesLen int
		for chid, onech := range cbrMgr {
			// go over each chain db events, send msg
			ret, isUpdateMsgFull := onech.pullMsgbrEvents(chid, r.Transactor.CliCtx, &updatesBytesLen)
			msg.Updates = append(msg.Updates, ret...)
			if isUpdateMsgFull {
				break
			}
		}
		if len(msg.Updates) > 0 {
			r.Transactor.AddTxMsg(msg)
			log.Debugln("MsgbrEvent updates count in one msg:", len(msg.Updates))
		}
	}
}

// Note if syncer changes before EndBlock, new syncer may still propose again
// the 2nd propose shouldn't get votes because when verify, sgn nodes will find it's already processed
// even it is voted, apply will still fail because x/massage will err
func (c *CbrOneChain) pullMsgbrEvents(chid uint64, cliCtx client.Context, updatesBytesLen *int) (ret []*synctypes.ProposeUpdate, isUpdateMsgFull bool) {
	// 1st loop over event names, then go over iter
	isUpdateMsgFull = false
	for _, evn := range msgEvNames {
		var keys, vals [][]byte
		c.msgbrLock.RLock()
		iterator, err := c.db.Iterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn+"-")))
		if err != nil {
			log.Errorln("Create db iterator err", err)
			c.msgbrLock.RUnlock()
			continue
		}
		for ; iterator.Valid(); iterator.Next() {
			keys = append(keys, iterator.Key())
			vals = append(vals, iterator.Value())
		}
		iterator.Close()
		c.msgbrLock.RUnlock()

		for i, key := range keys {
			err = c.db.Delete(key)
			if err != nil {
				log.Errorln("db Delete err", err)
				continue
			}

			evlog := new(ethtypes.Log)
			err := json.Unmarshal(vals[i], evlog)
			if err != nil {
				log.Errorf("failed to unmarshal onchev elog, key:%s, err:%s", string(key), err.Error())
				continue
			}

			skip, reason := c.skipMsgbrEvent(evn, evlog, cliCtx)
			if skip {
				log.Debugf("skip msgbr event: %s, chid %d, reason: %s", string(key), c.chainid, reason)
				continue
			}

			onchev := &cbrtypes.OnChainEvent{
				Chainid: chid,
				Evtype:  evn,
				Elog:    vals[i],
			}
			data, _ := onchev.Marshal()
			update := &synctypes.ProposeUpdate{
				Type:    synctypes.DataType_MsgbrOnChainEvent,
				ChainId: chid,
				Data:    data,
			}

			updateBytes, _ := proto.Marshal(update)
			*updatesBytesLen += len(updateBytes)
			if *updatesBytesLen > maxBytesPerUpdate {
				isUpdateMsgFull = true
				c.db.Set(key, vals[i]) // adds back to db
				break
			}

			ret = append(ret, update)
		}
		if isUpdateMsgFull {
			break
		}
	}
	return
}

func (c *CbrOneChain) getTransferType(srcChainBridgeAddr eth.Addr) msgbrtypes.TransferType {
	if c.cbrContract.GetAddr() == srcChainBridgeAddr {
		return msgbrtypes.TRANSFER_TYPE_LIQUIDITY_SEND
	} else if c.pegContracts.GetPegVaultContract().GetAddr() == srcChainBridgeAddr {
		// srcChain: deposit
		return msgbrtypes.TRANSFER_TYPE_PEG_MINT
	} else if c.pegContracts.GetPegBridgeContract().GetAddr() == srcChainBridgeAddr {
		// srcChain: burn
		return msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW
	} else {
		return msgbrtypes.TRANSFER_TYPE_NULL
	}
}

func getTransferInfoBySrcTransferId(
	cliCtx client.Context, srcTransferId eth.Hash, transferType msgbrtypes.TransferType) (sender eth.Addr, dstChainId uint64, token eth.Addr, amt *big.Int, isRefund bool, err error) {
	switch transferType {
	case msgbrtypes.TRANSFER_TYPE_LIQUIDITY_SEND:
		xferIdStr := srcTransferId.String()
		var resp *cbrtypes.QueryTransferStatusResponse
		resp, err = cbrcli.QueryTransferStatus(cliCtx, &cbrtypes.QueryTransferStatusRequest{
			TransferId: []string{xferIdStr},
		})
		if err == nil {
			status := resp.GetStatus()[xferIdStr]
			if status.GetSgnStatus() == cbrtypes.XferStatus_UNKNOWN {
				log.Debugf("unknown xfer status: status %s", resp.GetStatus()[xferIdStr].GetSgnStatus())
				err = fmt.Errorf("unknown xfer status: status %s", resp.GetStatus()[xferIdStr].GetSgnStatus())
			} else if status.GetSgnStatus() != cbrtypes.XferStatus_SUCCESS ||
				status.GetSgnStatus() != cbrtypes.XferStatus_OK_TO_RELAY {
				isRefund = true
			}
		}
		if isRefund {
			return
		}
		var relay *cbrtypes.XferRelay
		relay, err = cbrcli.QueryRelay(cliCtx, srcTransferId.Bytes())
		relayOnChain := new(cbrtypes.RelayOnChain)
		if err == nil {
			err = relayOnChain.Unmarshal(relay.Relay)
		}
		if err == nil {
			sender = eth.Bytes2Addr(relayOnChain.Sender)
			dstChainId = relayOnChain.DstChainId
			token = eth.Bytes2Addr(relayOnChain.Token)
			amt = new(big.Int).SetBytes(relayOnChain.Amount)
		}
	case msgbrtypes.TRANSFER_TYPE_PEG_MINT:
		var deposit pegbrtypes.DepositInfo
		deposit, err = pegcli.QueryDepositInfo(cliCtx, srcTransferId.String())
		if err == nil {
			if deposit.GetMintId() == nil {
				isRefund = true
				return
			}
			var mint pegbrtypes.MintInfo
			mint, err = pegcli.QueryMintInfo(cliCtx, eth.Bytes2Hash(deposit.GetMintId()).String())
			if err == nil {
				dstChainId = mint.GetChainId()
				mintOnChain := new(pegbrtypes.MintOnChain)
				err = mintOnChain.Unmarshal(mint.GetMintProtoBytes())
				if err != nil {
					log.Errorf("Unmarshal mintInfo.MintProtoBytes err %s", err)
					return sender, dstChainId, eth.ZeroAddr, new(big.Int).SetUint64(0), isRefund, err
				}
				sender = eth.Bytes2Addr(mintOnChain.GetDepositor())
				token = eth.Bytes2Addr(mintOnChain.GetToken())
				amt = new(big.Int).SetBytes(mintOnChain.GetAmount())
			}
		}
	case msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		var burn pegbrtypes.BurnInfo
		burn, err = pegcli.QueryBurnInfo(cliCtx, srcTransferId.String())
		if err == nil {
			if burn.GetWithdrawId() == nil {
				isRefund = true
				return
			}
			var withdraw pegbrtypes.WithdrawInfo
			withdraw, err = pegcli.QueryWithdrawInfo(cliCtx, eth.Bytes2Hash(burn.GetWithdrawId()).String())
			if err == nil {
				dstChainId = withdraw.GetChainId()
				withdrawOnChain := new(pegbrtypes.WithdrawOnChain)
				err = withdrawOnChain.Unmarshal(withdraw.GetWithdrawProtoBytes())
				if err != nil {
					log.Errorf("Unmarshal withdrawInfo.WithdrawProtoBytes err %s", err)
					return sender, dstChainId, eth.ZeroAddr, new(big.Int).SetUint64(0), isRefund, err
				}
				sender = eth.Bytes2Addr(withdrawOnChain.GetBurnAccount())
				token = eth.Bytes2Addr(withdrawOnChain.GetToken())
				amt = new(big.Int).SetBytes(withdrawOnChain.GetAmount())
			}
		}
	default:
		err = fmt.Errorf("unknown transfer type")
	}
	return sender, dstChainId, token, amt, isRefund, err
}

func (c *CbrOneChain) skipMsgbrEvent(evn string, evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	switch evn {
	case msgbrtypes.MsgEventMessage:
		skip, reason = c.skipMessageNoTransfer(evlog, cliCtx)
	case msgbrtypes.MsgEventMessageWithTransfer:
		skip, reason = c.skipMessageWithTransfer(evlog, cliCtx)
		if skip {
			return skip, reason
		}
		skip, reason = c.skipMessageWithTransferRefund(evlog, cliCtx)
	case msgbrtypes.MsgEventExecuted:
		skip, reason = c.skipMessageExecuted(evlog, cliCtx)
	}
	return
}

func (c *CbrOneChain) skipMessageExecuted(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	messageId, err := c.getMessageIdFromExecutedEvent(evlog)
	if err != nil {
		log.Errorf("getMessageId from evlog err: %s", err)
		return
	}
	exist, reason := c.checkMessageActive(messageId, cliCtx)
	skip = !exist
	log.Debugf("check skipMessageExecuted:%s, skip:%t", messageId, skip)
	return
}

func (c *CbrOneChain) skipMessageNoTransfer(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	messageId, err := c.getMessageIdFromMessageNoTransferEvent(cliCtx, evlog)
	if err != nil {
		log.Warnf("getMessageId from message evlog err: %s", err)
		return
	}
	return c.checkMessageActive(messageId, cliCtx)
}

func (c *CbrOneChain) skipMessageWithTransfer(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	messageId, err := c.getMessageIdFromMessageWithTransferEvent(cliCtx, evlog)
	if err != nil {
		reason = fmt.Sprintf("getMessageId from messageWithTransfer evlog err: %s", err)
		log.Warnln(reason)
		return
	}
	return c.checkMessageActive(messageId, cliCtx)
}

func (c *CbrOneChain) skipMessageWithTransferRefund(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.msgContracts.ParseMessageWithTransfer(*evlog)
	if err != nil {
		log.Warnf("getMessageId from messageWithTransfer evlog err: %s", err)
		return
	}

	return c.checkRefundExists(eth.Bytes2Hex(ev.SrcTransferId[:]), cliCtx)
}

// if message exist, skip
func (c *CbrOneChain) checkMessageActive(messageId string, cliCtx client.Context) (skip bool, reason string) {
	queryClient := msgbrtypes.NewQueryClient(cliCtx)
	resp, err := queryClient.IsMessageActive(context.Background(), &msgbrtypes.IsMessageActiveRequest{MessageId: messageId})
	if err != nil {
		log.Errorf("Query MessageExists err: %s", err)
		return
	}
	if resp.Exists {
		return true, fmt.Sprintf("msg with msgId %s already synced", messageId)
	}
	return
}

func (c *CbrOneChain) checkRefundExists(srcXferId string, cliCtx client.Context) (skip bool, reason string) {
	queryClient := msgbrtypes.NewQueryClient(cliCtx)
	resp, err := queryClient.RefundExists(context.Background(), &msgbrtypes.QueryRefundExistsRequest{SrcTransferId: srcXferId})
	if err != nil {
		log.Errorf("Query MessageExists err: %s", err)
		return
	}
	if resp.Exists {
		return true, fmt.Sprintf("msg refund (srcXferId %s) already applied", srcXferId)
	}
	return
}

func (c *CbrOneChain) getMessageIdFromExecutedEvent(evlog *ethtypes.Log) (string, error) {
	ev, err := c.msgContracts.ParseExecuted(*evlog)
	if err != nil {
		log.Errorln("getMessageId: cannot parse event:", err)
		return "", err
	}
	return eth.Hash(ev.Id).String(), nil
}

func (c *CbrOneChain) getMessageIdFromMessageNoTransferEvent(cliCtx client.Context, evlog *ethtypes.Log) (string, error) {
	ev, err := c.msgContracts.ParseMessage(*evlog)
	if err != nil {
		log.Errorln("getMessageIdFromMessageEvent: cannot parse event:", err)
		return "", err
	}
	message := msgbrtypes.Message{
		SrcChainId: c.chainid,
		Sender:     ev.Sender.String(),
		DstChainId: ev.DstChainId.Uint64(),
		Receiver:   ev.Receiver.String(),
		Data:       ev.Message,
	}
	messageId := message.ComputeMessageIdNoTransfer()
	return eth.Bytes2Hash(messageId).String(), nil
}

func (c *CbrOneChain) getMessageIdFromMessageWithTransferEvent(cliCtx client.Context, evlog *ethtypes.Log) (string, error) {
	ev, err := c.msgContracts.ParseMessageWithTransfer(*evlog)
	if err != nil {
		log.Errorln("getMessageIdFromMessageWithTransferEvent: cannot parse event:", err)
		return "", err
	}
	transferType := c.getTransferType(ev.Bridge)
	// check: bridge is either liquidity bridge, peg src vault, or peg dst bridge
	if transferType == msgbrtypes.TRANSFER_TYPE_NULL {
		return "", fmt.Errorf("getMessageIdFromMessageBusEvent failed, unknown bridge type, msg chainId:%d, transfer bridge:%s", c.chainid, ev.Bridge)
	}

	log.Debugf("getTransferInfoBySrcTransferId, srcTransferId:%s, type:%s", eth.Hash(ev.SrcTransferId), transferType)
	_, _, token, amt, _, err := getTransferInfoBySrcTransferId(cliCtx, ev.SrcTransferId, transferType)
	if err != nil {
		log.Warnf("getTransferInfoBySrcTransferId err:%+v", err)
		return "", err
	}
	message := msgbrtypes.Message{
		SrcChainId:   c.chainid,
		Sender:       ev.Sender.String(),
		DstChainId:   ev.DstChainId.Uint64(),
		Receiver:     ev.Receiver.String(),
		Data:         ev.Message,
		TransferType: transferType,
	}
	transfer := &msgbrtypes.Transfer{
		Token:  token.Bytes(),
		Amount: amt.String(),
		RefId:  ev.SrcTransferId[:],
	}
	execCtx := &msgbrtypes.ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	dstChain := CbrMgrInstance[ev.DstChainId.Uint64()]
	if dstChain == nil {
		log.Errorf("dstChain cannot be found in CbrMgrInstance, dstChainId:%d", ev.DstChainId.Uint64())
		return "", nil
	}
	messageId := execCtx.ComputeMessageId(dstChain.getBridgeAddrOnDstChain(transferType))
	return eth.Bytes2Hash(messageId).String(), nil
}

func (c *CbrOneChain) getBridgeAddrOnDstChain(transferType msgbrtypes.TransferType) eth.Addr {
	switch transferType {
	case msgbrtypes.TRANSFER_TYPE_LIQUIDITY_SEND:
		return c.cbrContract.Address
	case msgbrtypes.TRANSFER_TYPE_PEG_MINT:
		return c.pegContracts.GetPegBridgeContract().GetAddr()
	case msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		return c.pegContracts.GetPegVaultContract().GetAddr()
	}
	return eth.ZeroAddr
}
