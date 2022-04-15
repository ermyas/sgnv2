package relayer

import (
	"math/big"

	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) monMessage(blk *big.Int) {
	if c.msgContract.GetAddr() == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     msgtypes.MsgEventMessage,
		Contract:      c.msgContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(msgtypes.MsgEventMessage),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContract.ParseMessage(eLog)
		if err != nil {
			log.Errorln("monMessage: cannot parse event:", err)
			return false
		}
		msgId, _ := msgtypes.NewMessage(ev, c.chainid)
		log.Infof("MonEv: %s. msgId: %x", ev.PrettyLog(c.chainid), msgId)
		if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
			log.Warnln("eth addrs blocked", ev.String())
			return false
		}

		err = c.saveEvent(msgtypes.MsgEventMessage, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monMessageWithTransfer(blk *big.Int) {
	if c.msgContract.GetAddr() == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     msgtypes.MsgEventMessageWithTransfer,
		Contract:      c.msgContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(msgtypes.MsgEventMessageWithTransfer),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContract.ParseMessageWithTransfer(eLog)
		if err != nil {
			log.Errorln("monMessageWithTransfer: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid))
		if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
			log.Warnln("eth addrs blocked", ev.String())
			return false
		}

		err = c.saveEvent(msgtypes.MsgEventMessageWithTransfer, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monMessageBusEventExecuted(blk *big.Int) {
	if c.msgContract.GetAddr() == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     msgtypes.MsgEventExecuted,
		Contract:      c.msgContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(msgtypes.MsgEventExecuted),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.msgContract.ParseExecuted(eLog)
		if err != nil {
			log.Errorln("monMessageBusEventExecuted: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid))

		err = c.saveEvent(msgtypes.MsgEventExecuted, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}
