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
	smallDelay()
	c.monDelayXferAdd(blkNum)
	smallDelay()
	c.monDelayXferExec(blkNum)
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

		err = GatewayOnSend(common.Hash(ev.TransferId).String(), ev.Sender.String(), ev.Token.String(), ev.Amount.String(), eLog.TxHash.String(), c.chainid, ev.DstChainId)
		if err != nil {
			log.Warnf("GatewayOnSend err: %s, txId %x, txHash %x, chainId %d", err, ev.TransferId, eLog.TxHash, c.chainid)
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
			CurRelayerInstance.dbDelete(GetCbrXferKey(ev.SrcTransferId[:], c.chainid))
		}

		err = c.saveEvent(cbrtypes.CbrEventRelay, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		err = GatewayOnRelay(common.Hash(ev.SrcTransferId).String(), eLog.TxHash.String(), common.Hash(ev.TransferId).String(), ev.Amount.String())
		if err != nil {
			log.Warnf("UpdateTransfer err: %s, srcId %x, dstId %x, txHash %x, chainId %d", err, ev.SrcTransferId, ev.TransferId, eLog.TxHash, c.chainid)
		}
		return false
	})
}

func (c *CbrOneChain) monDelayXferAdd(blk *big.Int) {
	blkDelay := c.blkDelay / 2
	if blkDelay < 1 {
		blkDelay = 1
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventDelayXferAdd,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
		// much lower than the blk delay in cfg because the gateway service needs "Relay"
		// to be precedented by "DelayedTransferAdded".
		// this is ok because the the result action of monitoring "DelayedTransferAdded"
		// only changes the status for display use and is not related to fund safety.
		BlockDelay: blkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseDelayedTransferAdded(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		idstr := common.Hash(ev.Id).String()
		log.Infof("MonEv: DelayedTransferAdded chainId: %d, tx: %s, id %s", c.chainid, eLog.TxHash.String(), idstr)
		err = GatewayOnDelayXferAdd(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("GatewayOnDelayXferAdd err:", err)
		}
		return false
	})
}

func (c *CbrOneChain) monDelayXferExec(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    cbrtypes.CbrEventDelayXferExec,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseDelayedTransferExecuted(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		idstr := common.Hash(ev.Id).String()
		log.Infof("MonEv: DelayedTransferExecuted chainId: %d, tx: %s, id %s", c.chainid, eLog.TxHash.String(), idstr)
		err = GatewayOnDelayXferExec(idstr)
		if err != nil {
			log.Errorln("DelayedTransferExecuted err:", err)
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
		token, found := c.getTokenFromDB(ev.Token.String())
		if !found {
			return false
		}
		err = GatewayOnLiqAdd(ev.Provider.String(), token.Token.Symbol, token.Token.Address, ev.Amount.String(), eLog.TxHash.String(), c.chainid, ev.Seqnum)
		if err != nil {
			log.Warnf("UpsertLP err: %s, seqNum %d, amt %s, txHash %x, chainId %d", err, ev.Seqnum, ev.Amount.String(), eLog.TxHash, c.chainid)
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
		idstr := common.Hash(ev.WithdrawId).String()
		GatewayOnLiqWithdraw(idstr, c.chainid, ev.Seqnum, ev.Receiver.String())
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
					log.Warnf("Relay transaction failed, tx %x. %s", receipt.TxHash, logmsg)
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
