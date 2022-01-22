package relayer

import (
	"math/big"

	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	msgbrtypes "github.com/celer-network/sgn-v2/x/message/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) monMessage(blk *big.Int) {
	if c.msgContract.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     msgbrtypes.MsgEventMessage,
		Contract:      c.msgContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(msgbrtypes.MsgEventMessage),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContract.ParseMessage(eLog)
		if err != nil {
			log.Errorln("monMessage: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(msgbrtypes.MsgEventMessage, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monMessageWithTransfer(blk *big.Int) {
	if c.msgContract.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     msgbrtypes.MsgEventMessageWithTransfer,
		Contract:      c.msgContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(msgbrtypes.MsgEventMessageWithTransfer),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContract.ParseMessageWithTransfer(eLog)
		if err != nil {
			log.Errorln("monMessageWithTransfer: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(msgbrtypes.MsgEventMessageWithTransfer, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monMessageBusEventExecuted(blk *big.Int) {
	if c.msgContract.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     msgbrtypes.MsgEventExecuted,
		Contract:      c.msgContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(msgbrtypes.MsgEventExecuted),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContract.ParseExecuted(eLog)
		if err != nil {
			log.Errorln("monMessageBusEventExecuted: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(msgbrtypes.MsgEventExecuted, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}
