package relayer

import (
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	CbrEventSignersUpdated = "SignersUpdated"
)

var evNames = []string{CbrEventSend, CbrEventRelay, CbrEventLiqAdd, CbrEventWithdraw, CbrEventSignersUpdated}

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
	c.monSignersUpdated(blkNum)
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
		err = GatewayOnSend(common.Hash(ev.TransferId).String())
		if err != nil {
			log.Errorln("GatewayOnSend err:", err)
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
		err = GatewayOnRelay(common.Hash(ev.TransferId).String(), eLog.TxHash.String())
		if err != nil {
			log.Errorln("UpdateTransfer err:", err)
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
		token, chainId, found := c.getTokenFromDB(ev.Token.String())
		if !found {
			return false
		}
		err = GatewayOnLiqAdd(ev.Provider.String(), token.Token.Symbol, token.Token.Address, ev.Amount.String(), eLog.TxHash.String(), chainId, types.LPHistoryStatus_LP_WAITING_FOR_SGN, webapi.LPType_LP_TYPE_ADD, ev.Seqnum)
		if err != nil {
			log.Errorln("UpsertLP db err:", err)
			return false
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
		GatewayOnLiqWithdraw(ev.Seqnum)
		return false
	})
}

func (c *CbrOneChain) monSignersUpdated(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventSignersUpdated,
		Contract:   c.contract,
		StartBlock: blk,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseSignersUpdated(eLog)
		if err != nil {
			log.Errorln("monSignersUpdated: cannot parse event:", err)
			return false
		}
		logmsg := fmt.Sprintf("Catch event SignersUpdated: tx hash: %x, blknum: %d, signers %x",
			eLog.TxHash, eLog.BlockNumber, ev.CurSigners)
		err = c.saveEvent(CbrEventSignersUpdated, eLog)
		if err != nil {
			log.Errorf("%s, saveEvent err: %s", logmsg, err)
			return true // ask to recreate to process event again
		}
		c.setCurss(ev.CurSigners)
		log.Infoln(logmsg, c.getCurss().signers.String())
		return false
	})
}

// send relay tx onchain to cbridge contract, no wait mine
func (c *CbrOneChain) SendRelay(relay, curss []byte, sigs [][]byte) error {
	tx, err := c.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("Relay transaction %x succeeded", receipt.TxHash)
				} else {
					log.Errorf("Relay transaction %x failed", receipt.TxHash)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Errorf("Relay transaction %x err: %s", tx.Hash(), err)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.contract.Relay(opts, relay, curss, sigs)
		},
	)

	if err != nil {
		return err
	}

	log.Infoln("Relay tx submitted", tx.Hash().Hex())
	return nil
}

// send updateSigners tx onchain to cbridge contract, no wait mine
func (c *CbrOneChain) UpdateSigners(newss, curss []byte, sigs [][]byte) error {
	tx, err := c.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("UpdateSigners transaction %x succeeded", receipt.TxHash)
				} else {
					log.Errorf("UpdateSigners transaction %x failed", receipt.TxHash)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Errorf("UpdateSigners transaction %x err: %s", tx.Hash(), err)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.contract.UpdateSigners(opts, newss, curss, sigs)
		},
	)

	if err != nil {
		return err
	}

	log.Infoln("UpdateSigners tx submitted", tx.Hash().Hex())
	return nil
}
