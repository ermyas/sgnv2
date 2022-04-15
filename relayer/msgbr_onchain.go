package relayer

import (
	"time"

	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) MonMessage() {
	if c.msgContract.GetAddr() != eth.ZeroAddr {
		go c.mon.MonAddr(mon2.PerAddrCfg{
			Addr:    c.msgContract.GetAddr(),
			ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
			AbiStr:  c.msgContract.GetABI(), // to parse event name by topics[0]
		}, c.msgEvCallback)
	}
}

func (c *CbrOneChain) msgEvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "Message":
		c.handleMessage(elog)
	case "MessageWithTransfer":
		c.handleMessageWithTransfer(elog)
	case "Executed":
		c.handleMessageBusEventExecuted(elog)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) handleMessage(eLog ethtypes.Log) {
	ev, err := c.msgContract.ParseMessage(eLog)
	if err != nil {
		log.Errorln("monMessage: cannot parse event:", err)
		return
	}
	msgId, _ := msgtypes.NewMessage(ev, c.chainid)
	log.Infof("MonEv: %s. msgId: %x", ev.PrettyLog(c.chainid), msgId)
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnln("eth addrs blocked", ev.String())
		return
	}

	err = c.saveEvent(msgtypes.MsgEventMessage, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) handleMessageWithTransfer(eLog ethtypes.Log) {
	ev, err := c.msgContract.ParseMessageWithTransfer(eLog)
	if err != nil {
		log.Errorln("monMessageWithTransfer: cannot parse event:", err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid))
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnln("eth addrs blocked", ev.String())
		return
	}

	err = c.saveEvent(msgtypes.MsgEventMessageWithTransfer, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) handleMessageBusEventExecuted(eLog ethtypes.Log) {
	ev, err := c.msgContract.ParseExecuted(eLog)
	if err != nil {
		log.Errorln("monMessageBusEventExecuted: cannot parse event:", err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid))

	err = c.saveEvent(msgtypes.MsgEventExecuted, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}
