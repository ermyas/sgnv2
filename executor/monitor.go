package executor

import (
	"time"

	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/executor/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *Chain) startMonitoring() {
	c.startBlk = c.monitor.GetCurrentBlockNumber()
	c.monitorBusExecuted()
	smallDelay()
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
		err = Dal.UpdateStatus(e.MsgId[:], status)
		if err != nil {
			log.Errorf("failed to update execution_context %x: %v", e.MsgId[:], err)
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
