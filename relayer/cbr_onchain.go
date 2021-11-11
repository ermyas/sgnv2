package relayer

import (
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var evNames = []string{
	cbrtypes.CbrEventSend,
	cbrtypes.CbrEventRelay,
	cbrtypes.CbrEventLiqAdd,
	cbrtypes.CbrEventWithdraw,
	cbrtypes.CbrEventSignersUpdated,
}

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
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventSend,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseSend(eLog)
		if err != nil {
			log.Errorln("monSend: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventSend, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		log.Infof("transferId in Send event: %s", common.Hash(ev.TransferId).String())

		err = GatewayOnSend(common.Hash(ev.TransferId).String(), ev.Sender.String(), ev.Token.String(), ev.Amount.String(), eLog.TxHash.String(), c.chainid, ev.DstChainId)
		if err != nil {
			log.Errorln("GatewayOnSend err:", err)
		}
		return false
	})
}

func (c *CbrOneChain) monRelay(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventRelay,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseRelay(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		// delete to-submit relay at local if have, as it's been submitted (by other nodes or me)
		if CurRelayerInstance == nil {
			log.Errorln("CurRelayerInstance not initialized", err)
		} else {
			CurRelayerInstance.dbDelete(GetCbrXferKey(ev.TransferId[:]))
		}

		err = c.saveEvent(cbrtypes.CbrEventRelay, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		err = GatewayOnRelay(common.Hash(ev.SrcTransferId).String(), eLog.TxHash.String(), common.Hash(ev.TransferId).String(), ev.Amount.String())
		if err != nil {
			log.Errorln("UpdateTransfer err:", err)
		}
		return false
	})
}

func (c *CbrOneChain) monLiqAdd(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventLiqAdd,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseLiquidityAdded(eLog)
		if err != nil {
			log.Errorln("monLiqAdd: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventLiqAdd, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		// todo: use cbr query to get symbol to avoid query db, as we already have all other info
		token, chainId, found := c.getTokenFromDB(ev.Token.String())
		if !found {
			return false
		}
		err = GatewayOnLiqAdd(ev.Provider.String(), token.Token.Symbol, token.Token.Address, ev.Amount.String(), eLog.TxHash.String(), chainId, ev.Seqnum)
		if err != nil {
			log.Errorln("UpsertLP db err:", err)
			return false
		}
		return false
	})
}

func (c *CbrOneChain) monWithdraw(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventWithdraw,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseWithdrawDone(eLog)
		if err != nil {
			log.Errorln("monWithdraw: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventWithdraw, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		GatewayOnLiqWithdraw(c.chainid, ev.Seqnum, ev.Receiver.String())
		return false
	})
}

func (c *CbrOneChain) monSignersUpdated(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventSignersUpdated,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseSignersUpdated(eLog)
		if err != nil {
			log.Errorln("monSignersUpdated: cannot parse event:", err)
			return false
		}
		c.setCurssByEvent(ev)
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventSignersUpdated, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

// send relay tx onchain to cbridge contract, no wait mine
func (c *CbrOneChain) SendRelay(relayBytes []byte, sigs [][]byte, curss currentSigners, relayMsg *cbrtypes.RelayOnChain) (string, error) {
	logmsg := fmt.Sprintf("srcXferId %x chain %d->%d", relayMsg.GetSrcTransferId(), relayMsg.GetSrcChainId(), relayMsg.GetDstChainId())
	tx, err := c.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("Relay transaction succeeded, tx %x. %s", receipt.TxHash, logmsg)
				} else {
					log.Errorf("Relay transaction failed, tx %x. %s", receipt.TxHash, logmsg)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("Relay transaction err: %s. %s", err, logmsg)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.contract.Relay(opts, relayBytes, sigs, curss.addrs, curss.powers)
		},
	)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (c *CbrOneChain) existTransferId(transferId common.Hash) (bool, error) {
	return c.contract.BridgeCaller.Transfers(&bind.CallOpts{}, transferId)
}
