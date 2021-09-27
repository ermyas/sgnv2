package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
		err = dal.UpdateTransferStatus(common.Hash(ev.TransferId).String(), uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE))
		if err != nil {
			log.Errorln("UpdateTransfer err:", err)
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
		err = dal.TransferCompleted(common.Hash(ev.TransferId).String(), eLog.TxHash.String())
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
		newContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		chainId, err := c.ChainID(newContext)
		if err != nil {
			log.Errorln("get chain id err:", err)
			return false
		}
		token, found, err := dal.GetTokenByAddr(ev.Token.String(), chainId.Uint64())
		if err != nil || !found {
			return false
		}
		err = dal.UpsertLP(ev.Provider.String(), token.Token.Symbol, token.Token.Address, ev.Amount.String(), eLog.TxHash.String(), chainId.Uint64(), uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN), uint64(webapi.LPType_LP_TYPE_ADD), ev.Seqnum)
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
		transferId, found, err := dal.GetTransferBySeqNum(ev.Seqnum)
		if err != nil {
			return false
		}
		if found {
			dbErr := dal.UpdateTransferStatus(transferId, uint64(types.TransferHistoryStatus_TRANSFER_REFUNDED))
			if dbErr != nil {
				log.Errorln("db when UpdateTransferStatus to TRANSFER_REFUNDED err:", err)
			}
		} else {
			dbErr := dal.UpdateLPStatus(ev.Seqnum, uint64(types.LPHistoryStatus_LP_COMPLETED))
			if dbErr != nil {
				log.Errorln("db when UpdateLPStatus to LP_COMPLETED err:", err)
			}
		}
		return false
	})
}

func (c *CbrOneChain) monNewSigners(blk *big.Int) {
	cfg := &monitor.Config{
		EventName:  CbrEventSignersUpdated,
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
		err = c.saveEvent(CbrEventSignersUpdated, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		c.curss.setSigners(ev.CurSigners)
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

// query chain to verify event is the same, return err if mismatch
// TODO: impl logic
func (c *CbrOneChain) CheckEvent(evtype string, tocheck *ethtypes.Log) (retry bool, err error) {
	switch evtype {
	case CbrEventLiqAdd:
		return false, nil
	case CbrEventSend:
		return false, nil
	case CbrEventRelay:
		return false, nil
	case CbrEventSignersUpdated:
		ev, err := c.contract.ParseSignersUpdated(*tocheck)
		if err != nil {
			return false, err
		}
		ssHash, err := c.contract.SsHash(&bind.CallOpts{})
		if err != nil {
			return true, err
		}
		if eth.Bytes2Hash(crypto.Keccak256(ev.CurSigners)) != ssHash {
			return false, fmt.Errorf("ssHash not match onchain value")
		}
		return false, nil
	}
	return false, fmt.Errorf("invalid event type %s", evtype)
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
