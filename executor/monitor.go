package executor

import (
	"time"

	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *Chain) startMonitoring() {
	c.monitorBusExecuted()
	smallDelay()
}

func (c *Chain) monitorBusExecuted() {
	addrConfig := mon2.PerAddrCfg{
		Addr:    c.MsgBus.Address,
		ChkIntv: time.Minute,
		AbiStr:  eth.MessageBusABI,
	}
	if c.filterAddr != "" {
		// first topic would be a signature of the event
		// while we switch eventName later, so we don't filter event signature here.
		addrConfig.Topics = [][]common.Hash{[]common.Hash{}, []common.Hash{common.HexToHash(c.filterAddr)}}
	}
	go c.monitor2.MonAddr(addrConfig, func(evName string, eLog ethtypes.Log) {
		switch evName {
		case "Executed":
			e, err := c.MsgBus.ParseExecuted(eLog)
			if err != nil {
				log.Errorln("cannot parse event Executed", err)
				return
			}
			log.Infof("monitorBusExecuted: got event Executed %v", e)
			status, err := types.NewExecutionStatus(e.Status)
			if err != nil {
				log.Errorln("monitorBusExecuted: ", err)
				return
			}
			err = Dal.UpdateStatus(e.MsgId[:], status)
			if err != nil {
				log.Errorf("failed to update execution_context %x: %v", e.MsgId[:], err)
			}
		}
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
