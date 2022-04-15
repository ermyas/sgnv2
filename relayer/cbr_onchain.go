package relayer

import (
	"context"
	"fmt"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
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
	if relayerInstance == nil {
		log.Fatal("relayer instance not initiated")
	}
	smallDelay := func() {
		time.Sleep(20 * time.Millisecond)
	}
	if c.FlowClient != nil {
		c.monitorFlow()
		return
	}

	c.MonCbridge()
	smallDelay()
	c.MonWdInbox()
	smallDelay()
	c.MonVault()
	smallDelay()
	c.MonPegBridge()
	smallDelay()
	c.MonMessage()
}

func (c *CbrOneChain) MonCbridge() {
	go c.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    c.cbrContract.Address,
		ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
		AbiStr:  c.cbrContract.GetABI(), // to parse event name by topics[0]
	}, c.cbrEvCallback)
}

func (c *CbrOneChain) cbrEvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "Send":
		c.handleSend(elog)
	case "Relay":
		c.handleRelay(elog)
	case "LiquidityAdded":
		c.handleLiqAdd(elog)
	case "WithdrawDone":
		c.handleWithdraw(elog)
	case "SignersUpdated":
		c.handleSignersUpdated(elog)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) handleSend(eLog ethtypes.Log) {
	ev, err := c.cbrContract.ParseSend(eLog)
	if err != nil {
		log.Errorf("monSend: chain %d cannot parse event: %s", c.chainid, err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnln("eth addrs blocked", ev.String())
		return
	}

	err = c.saveEvent(cbrtypes.CbrEventSend, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) handleRelay(eLog ethtypes.Log) {
	ev, err := c.cbrContract.ParseRelay(eLog)
	if err != nil {
		log.Errorf("monRelay: chain %d cannot parse event: %s", c.chainid, err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

	// delete to-submit relay at local if have, as it's been submitted (by other nodes or me)
	relayerInstance.dbDelete(GetCbrXferKey(ev.SrcTransferId[:], c.chainid))

	err = c.saveEvent(cbrtypes.CbrEventRelay, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) handleLiqAdd(eLog ethtypes.Log) {
	ev, err := c.cbrContract.ParseLiquidityAdded(eLog)
	if err != nil {
		log.Errorf("monLiqAdd: chain %d cannot parse event: %s", c.chainid, err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
	if relayerInstance.isEthAddrBlocked(ev.Provider) {
		log.Warnln("eth addrs blocked", ev.String())
		return
	}

	err = c.saveEvent(cbrtypes.CbrEventLiqAdd, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) handleWithdraw(eLog ethtypes.Log) {
	ev, err := c.cbrContract.ParseWithdrawDone(eLog)
	if err != nil {
		log.Errorf("monWithdraw: chain %d cannot parse event: %s", c.chainid, err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

	err = c.saveEvent(cbrtypes.CbrEventWithdraw, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) handleSignersUpdated(eLog ethtypes.Log) {
	ev, err := c.cbrContract.ParseSignersUpdated(eLog)
	if err != nil {
		log.Errorf("monSignersUpdated: chain %d cannot parse event: %s", c.chainid, err)
		return
	}
	c.setCurssByEvent(ev)
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

	err = c.saveEvent(cbrtypes.CbrEventSignersUpdated, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) MonWdInbox() {
	if c.wdiContract.GetAddr() != eth.ZeroAddr {
		go c.mon.MonAddr(mon2.PerAddrCfg{
			Addr:    c.wdiContract.Address,
			ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
			AbiStr:  c.wdiContract.GetABI(), // to parse event name by topics[0]
		}, c.wdiEvCallback)
	}
}

func (c *CbrOneChain) wdiEvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "WithdrawalRequest":
		c.handleWithdrawalRequest(elog)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) handleWithdrawalRequest(eLog ethtypes.Log) {
	ev, err := c.wdiContract.ParseWithdrawalRequest(eLog)
	if err != nil {
		log.Errorf("monWithdrawalRequest: chain %d cannot parse event: %s", c.chainid, err)
		return
	}
	log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnln("eth addrs blocked", ev.String())
		return
	}

	err = c.saveEvent(cbrtypes.CbrEventWithdrawalRequest, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
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
