package executor

import (
	"time"

	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *Chain) startMonitoring() {
	c.startBlk = c.monitor.GetCurrentBlockNumber()
	c.monitorBusExecuted()
	smallDelay()
	c.monitorBridgeRelay()
	smallDelay()
	c.monitorPegMint()
	smallDelay()
	c.monitorPegWithdrawn()
}

func (c *Chain) monitorBridgeRelay() {
	if c.LiqBridge.GetAddr() == eth.ZeroAddr {
		return
	}
	cfg := c.makeDefaultMonConf(types.LiqBridgeEventRelay, c.LiqBridge)
	c.monitor.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) bool {
		e, err := c.LiqBridge.ParseRelay(eLog)
		if err != nil {
			log.Errorln("cannot ParseRelay", err)
			return false
		}
		log.Infof("monitorBridgeRelay: got event Relay %v", e)
		messageId := msgtypes.ComputeMessageIdFromDstTransfer(e.TransferId[:], c.LiqBridge.Address)
		log.Infoln("message id", eth.Bytes2Hex(messageId))
		err = Dal.SaveTransfer(messageId)
		if err != nil {
			log.Errorf("failed to update execution_context %x: %v", messageId, err)
		}
		return false
	})
}

func (c *Chain) monitorPegMint() {
	if c.PegBridge.GetAddr() == eth.ZeroAddr {
		return
	}
	cfg := c.makeDefaultMonConf(types.PegBridgeEventMint, c.PegBridge)
	c.monitor.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) bool {
		e, err := c.PegBridge.ParseMint(eLog)
		if err != nil {
			log.Errorln("cannot ParseMint", err)
			return false
		}
		log.Infof("monitorPegMint: got event Mint %v", e)
		msg := msgtypes.Message{
			SrcChainId:    e.RefChainId,
			Sender:        eth.Addr2Hex(e.Depositor),
			DstChainId:    c.ChainID,
			Receiver:      eth.Addr2Hex(e.Account),
			TransferType:  msgtypes.TRANSFER_TYPE_PEG_MINT,
			TransferRefId: e.RefId[:],
		}
		transfer := &msgtypes.Transfer{
			Token:  e.Token.Bytes(),
			Amount: e.Amount.String(),
		}
		execCtx := &msgtypes.ExecutionContext{
			Message:  msg,
			Transfer: transfer,
		}
		messageId := getMessageIdWithTransfer(c, execCtx)
		err = Dal.SaveTransfer(messageId)
		if err != nil {
			log.Errorf("failed to update execution_context %x: %v", messageId, err)
		}
		return false
	})
}

func (c *Chain) monitorPegWithdrawn() {
	if c.PegVault.GetAddr() == eth.ZeroAddr {
		return
	}
	cfg := c.makeDefaultMonConf(types.PegVaultEventWithdrawn, c.PegVault)
	c.monitor.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) bool {
		e, err := c.PegVault.ParseWithdrawn(eLog)
		if err != nil {
			log.Errorln("cannot ParseWithdrawn", err)
			return false
		}
		log.Infof("monitorPegWithdrawn: got event Withdrawn %v", e)
		msg := msgtypes.Message{
			SrcChainId:    e.RefChainId,
			Sender:        eth.Addr2Hex(e.BurnAccount),
			DstChainId:    c.ChainID,
			Receiver:      eth.Addr2Hex(e.Receiver),
			TransferType:  msgtypes.TRANSFER_TYPE_PEG_WITHDRAW,
			TransferRefId: e.RefId[:],
		}
		transfer := &msgtypes.Transfer{
			Token:  e.Token.Bytes(),
			Amount: e.Amount.String(),
		}
		execCtx := &msgtypes.ExecutionContext{
			Message:  msg,
			Transfer: transfer,
		}
		messageId := getMessageIdWithTransfer(c, execCtx)
		err = Dal.SaveTransfer(messageId)
		if err != nil {
			log.Errorf("failed to update execution_context %x: %v", messageId, err)
		}
		return false
	})
}

func (c *Chain) monitorBusExecuted() {
	cfg := c.makeDefaultMonConf(types.MessageBusEventExecuted, c.MsgBus)
	c.monitor.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) bool {
		e, err := c.MsgBus.ParseExecuted(eLog)
		if err != nil {
			log.Errorln("cannot parse event Executed", err)
			return false
		}
		log.Infof("monitorBusExecuted: got event Executed %v", e)
		status, err := types.NewExecutionStatus(e.Status)
		if err != nil {
			log.Errorln("monitorBusExecuted: ", err)
			return false
		}
		err = Dal.UpdateStatus(e.Id[:], status)
		if err != nil {
			log.Errorf("failed to update execution_context %x: %v", e.Id[:], err)
		}
		return false
	})
}

func (c *Chain) makeDefaultMonConf(name string, contract monitor.Contract) *monitor.Config {
	return &monitor.Config{
		ChainId:      c.ChainID,
		EventName:    name,
		Contract:     contract,
		StartBlock:   c.startBlk,
		ForwardDelay: c.fwdBlkDelay,
	}
}

func smallDelay() {
	time.Sleep(100 * time.Millisecond)
}
