package relayer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const PendingNonceLimit = 20

var evNames = []string{
	cbrtypes.CbrEventSend,
	cbrtypes.CbrEventRelay,
	cbrtypes.CbrEventLiqAdd,
	cbrtypes.CbrEventWithdraw,
	cbrtypes.CbrEventSignersUpdated,
	cbrtypes.CbrEventWithdrawalRequest,
}

// funcs for monitor cbridge events
func (c *CbrOneChain) startMon() {
	smallDelay := func() {
		time.Sleep(20 * time.Millisecond)
	}
	if c.FlowClient != nil {
		c.monitorFlow()
		return
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
	c.monWithdrawalRequest(blkNum)
	smallDelay()
	c.monSignersUpdated(blkNum)

	smallDelay()
	c.monPegbrDeposited(blkNum)
	smallDelay()
	c.monPegbrMint(blkNum)
	smallDelay()
	c.monPegbrBurn(blkNum)
	smallDelay()
	c.monPegbrWithdrawn(blkNum)
	smallDelay()
	c.monMessageWithTransfer(blkNum)
	smallDelay()
	c.monMessageBusEventExecuted(blkNum)
	smallDelay()
	c.monMessage(blkNum)
}

func (c *CbrOneChain) monSend(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     cbrtypes.CbrEventSend,
		Contract:      c.cbrContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(cbrtypes.CbrEventSend),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.cbrContract.ParseSend(eLog)
		if err != nil {
			log.Errorf("monSend: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventSend, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monRelay(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     cbrtypes.CbrEventRelay,
		Contract:      c.cbrContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(cbrtypes.CbrEventRelay),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.cbrContract.ParseRelay(eLog)
		if err != nil {
			log.Errorf("monRelay: chain %d cannot parse event: %s", c.chainid, err)
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
		return false
	})
}

func (c *CbrOneChain) monLiqAdd(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     cbrtypes.CbrEventLiqAdd,
		Contract:      c.cbrContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(cbrtypes.CbrEventLiqAdd),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.cbrContract.ParseLiquidityAdded(eLog)
		if err != nil {
			log.Errorf("monLiqAdd: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventLiqAdd, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monWithdraw(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     cbrtypes.CbrEventWithdraw,
		Contract:      c.cbrContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(cbrtypes.CbrEventWithdraw),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.cbrContract.ParseWithdrawDone(eLog)
		if err != nil {
			log.Errorf("monWithdraw: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventWithdraw, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monWithdrawalRequest(blk *big.Int) {
	if c.wdiContract.GetAddr() == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     cbrtypes.CbrEventWithdrawalRequest,
		Contract:      c.wdiContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(cbrtypes.CbrEventWithdrawalRequest),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.wdiContract.ParseWithdrawalRequest(eLog)
		if err != nil {
			log.Errorf("monWithdrawalRequest: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(cbrtypes.CbrEventWithdrawalRequest, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monSignersUpdated(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:       c.chainid,
		EventName:     cbrtypes.CbrEventSignersUpdated,
		Contract:      c.cbrContract,
		StartBlock:    blk,
		ForwardDelay:  c.forwardBlkDelay,
		CheckInterval: c.getEventCheckInterval(cbrtypes.CbrEventSignersUpdated),
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.cbrContract.ParseSignersUpdated(eLog)
		if err != nil {
			log.Errorf("monSignersUpdated: chain %d cannot parse event: %s", c.chainid, err)
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
	err := c.checkPendingNonce()
	if err != nil {
		return "", fmt.Errorf("Pending nonce check failed. %w", err)
	}
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
			return c.cbrContract.Relay(opts, relayBytes, sigs, curss.addrs, curss.powers)
		},
	)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (c *CbrOneChain) existTransferId(transferId eth.Hash) (bool, error) {
	return c.cbrContract.BridgeCaller.Transfers(&bind.CallOpts{}, transferId)
}

func (c *CbrOneChain) checkPendingNonce() error {
	nonce, err := c.Client.NonceAt(context.Background(), c.Transactor.Address(), nil)
	if err != nil {
		return fmt.Errorf("NonceAt %w", err)
	}
	pendingNonce, err := c.Client.PendingNonceAt(context.Background(), c.Transactor.Address())
	if err != nil {
		return fmt.Errorf("PendingNonceAt %w", err)
	}
	if pendingNonce-nonce > PendingNonceLimit {
		return fmt.Errorf("pendingNonce %d nonce %d", pendingNonce, nonce)
	}
	return nil
}
