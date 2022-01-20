package relayer

import (
	"math/big"

	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	msgbrtypes "github.com/celer-network/sgn-v2/x/message/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) monMessage(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    msgbrtypes.MsgEventMessage,
		Contract:     c.msgContracts,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContracts.ParseMessage(eLog)
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
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    msgbrtypes.MsgEventMessageWithTransfer,
		Contract:     c.msgContracts,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContracts.ParseMessageWithTransfer(eLog)
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
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    msgbrtypes.MsgEventExecuted,
		Contract:     c.msgContracts,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContracts.ParseExecuted(eLog)
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
