package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgbrtypes "github.com/celer-network/sgn-v2/x/message/types"
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
		c.lock.RLock()
		iterator, err := c.db.Iterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn+"-")))
		if err != nil {
			log.Errorln("Create db iterator err", err)
			c.lock.RUnlock()
			continue
		}
		for ; iterator.Valid(); iterator.Next() {
			keys = append(keys, iterator.Key())
			vals = append(vals, iterator.Value())
		}
		iterator.Close()
		c.lock.RUnlock()

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

func (c *CbrOneChain) getSrcBridgeType(srcChainBridgeAddr eth.Addr) msgbrtypes.BridgeType {
	if c.cbrContract.GetAddr() == srcChainBridgeAddr {
		return msgbrtypes.BRIDGE_TYPE_LIQUIDITY
	} else if c.pegContracts.GetPegVaultContract().GetAddr() == srcChainBridgeAddr ||
		c.pegContracts.GetPegVaultV2Contract().GetAddr() == srcChainBridgeAddr {
		// srcChain: deposit
		return msgbrtypes.BRIDGE_TYPE_PEG_VAULT
	} else if c.pegContracts.GetPegBridgeContract().GetAddr() == srcChainBridgeAddr ||
		c.pegContracts.GetPegBridgeV2Contract().GetAddr() == srcChainBridgeAddr {
		// srcChain: burn
		return msgbrtypes.BRIDGE_TYPE_PEG_BRIDGE
	} else {
		return msgbrtypes.BRIDGE_TYPE_NULL
	}
}

func (c *CbrOneChain) skipMsgbrEvent(evn string, evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	switch evn {
	case msgbrtypes.MsgEventMessage:
		skip, reason = c.skipMessageNoTransfer(evlog, cliCtx)
	case msgbrtypes.MsgEventMessageWithTransfer:
		skip, reason = c.skipMessageWithTransfer(evlog, cliCtx)
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
	queryClient := msgbrtypes.NewQueryClient(cliCtx)
	resp, err := queryClient.IsMessageActive(context.Background(), &msgbrtypes.IsMessageActiveRequest{MessageId: messageId})
	if err != nil {
		return
	}
	if !resp.Exists {
		return true, fmt.Sprintf("msgId %s not active", messageId)
	}
	return
}

func (c *CbrOneChain) skipMessageNoTransfer(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.msgContract.ParseMessage(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	messageId, _ := msgbrtypes.NewMessage(ev, c.chainid)
	queryClient := msgbrtypes.NewQueryClient(cliCtx)
	resp, err := queryClient.MessageExists(
		context.Background(), &msgbrtypes.QueryMessageExistsRequest{MessageId: messageId.Hex()})
	if err != nil {
		return
	}
	if resp.Exists {
		return true, fmt.Sprintf("msgId %x already synced", messageId)
	}
	return
}

func (c *CbrOneChain) skipMessageWithTransfer(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.msgContract.ParseMessageWithTransfer(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	srcBridgeType := c.getSrcBridgeType(ev.Bridge)
	queryClient := msgbrtypes.NewQueryClient(cliCtx)
	resp, err := queryClient.ExecutionContextBySrcTransfer(context.Background(),
		&msgbrtypes.QueryExecutionContextBySrcTransferRequest{
			SrcTransferId: eth.Hash(ev.SrcTransferId).Hex(),
			SrcBridgeType: srcBridgeType,
			MessageIdOnly: true,
		})
	if err != nil {
		return
	}
	messageId := eth.Bytes2Hash(resp.ExecutionContext.GetMessageId())
	if messageId != eth.ZeroHash {
		return true, fmt.Sprintf("srcTransfer %x with msgId %x already synced", ev.SrcTransferId, messageId)
	}
	return
}

func (c *CbrOneChain) getMessageIdFromExecutedEvent(evlog *ethtypes.Log) (string, error) {
	ev, err := c.msgContract.ParseExecuted(*evlog)
	if err != nil {
		log.Errorln("getMessageId: cannot parse event:", err)
		return "", err
	}
	return eth.Hash(ev.MsgId).String(), nil
}

func (c *CbrOneChain) getMessageIdFromMessageNoTransferEvent(cliCtx client.Context, evlog *ethtypes.Log) (string, error) {
	ev, err := c.msgContract.ParseMessage(*evlog)
	if err != nil {
		log.Errorln("getMessageIdFromMessageEvent: cannot parse event:", err)
		return "", err
	}
	messageId, _ := msgbrtypes.NewMessage(ev, c.chainid)
	return messageId.String(), nil
}

func (c *CbrOneChain) getBridgeAddrOnDstChain(transferType msgbrtypes.TransferType) eth.Addr {
	switch transferType {
	case msgbrtypes.TRANSFER_TYPE_LIQUIDITY_SEND:
		return c.cbrContract.GetAddr()
	case msgbrtypes.TRANSFER_TYPE_PEG_MINT:
		return c.pegContracts.GetPegBridgeContract().GetAddr()
	case msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		return c.pegContracts.GetPegVaultContract().GetAddr()
	case msgbrtypes.TRANSFER_TYPE_PEG_MINT_V2:
		return c.pegContracts.GetPegBridgeV2Contract().GetAddr()
	case msgbrtypes.TRANSFER_TYPE_PEG_WITHDRAW_V2:
		return c.pegContracts.GetPegVaultV2Contract().GetAddr()
	}
	return eth.ZeroAddr
}
