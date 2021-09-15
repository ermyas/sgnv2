package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	// event names
	CbrEventSend  = "Send"
	CbrEventRelay = "Relay"
	// from pool.sol
	CbrEventLiqAdd   = "LiquidityAdded"
	CbrEventWithdraw = "WithdrawDone" // could be LP or user
	// from signers.sol
	CbrEventNewSigners = "SignersUpdated"
)

var evNames = []string{CbrEventSend, CbrEventRelay, CbrEventLiqAdd, CbrEventWithdraw, CbrEventNewSigners}

// funcs for monitor cbridge events
func (c *CbrOneChain) startMon() {
	smallDelay := func() {
		time.Sleep(100 * time.Millisecond)
	}
	// avoid repeated get block number calls
	blkNum := c.mon.GetCurrentBlockNumber()
	c.monSend(blkNum)
	smallDelay()
	c.monRelay(blkNum)
	smallDelay()
	c.monLiqAdd(blkNum)
	smallDelay()
	c.monWithdraw(blkNum)
	smallDelay()
	c.monNewSigners(blkNum)
}

func (c *CbrOneChain) monSend(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventSend,
		Contract:   c.contract,
		StartBlock: blk,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseSend(eLog)
		if err != nil {
			log.Errorln("monSend: cannot parse event:", err)
			return false
		}
		log.Infof("Send event: %+v", ev)
		err = c.saveEvent(CbrEventSend, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monRelay(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventRelay,
		Contract:   c.contract,
		StartBlock: blk,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseRelay(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		log.Infof("Relay event: %+v", ev)
		err = c.saveEvent(CbrEventRelay, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monLiqAdd(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventLiqAdd,
		Contract:   c.contract,
		StartBlock: blk,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseLiquidityAdded(eLog)
		if err != nil {
			log.Errorln("monLiqAdd: cannot parse event:", err)
			return false
		}
		log.Infof("LiqAdd event: %+v", ev)

		err = c.saveEvent(CbrEventLiqAdd, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monWithdraw(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventWithdraw,
		Contract:   c.contract,
		StartBlock: blk,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseWithdrawDone(eLog)
		if err != nil {
			log.Errorln("monWithdraw: cannot parse event:", err)
			return false
		}
		log.Infof("Withdraw event: %+v", ev)
		err = c.saveEvent(CbrEventWithdraw, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monNewSigners(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventNewSigners,
		Contract:   c.contract,
		StartBlock: blk,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseSignersUpdated(eLog)
		if err != nil {
			log.Errorln("monNewSigners: cannot parse event:", err)
			return false
		}
		log.Infof("NewSigners event: %+v", ev)
		err = c.saveEvent(CbrEventNewSigners, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

// each event's key is name-blkNum-index, value is json marshaled elog
func (c *CbrOneChain) saveEvent(name string, elog ethtypes.Log) error {
	key := fmt.Sprintf("%s-%d-%d", name, elog.BlockNumber, elog.Index)
	val, _ := json.Marshal(elog)
	return c.db.Set([]byte(key), val)
}

func (c *CbrOneChain) delEvent(name string, blknum, idx uint64) error {
	return c.db.Delete([]byte(fmt.Sprintf("%s-%d-%d", name, blknum, idx)))
}

// query chain to verify event is the same, return true if match
// TODO: impl logic
func (oc *CbrOneChain) CheckEvent(tocheck *ethtypes.Log) (bool, error) {
	return true, nil
}
