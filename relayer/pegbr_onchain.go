package relayer

import (
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	pegtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) MonVault() {
	if c.pegContracts.vault.GetAddr() != eth.ZeroAddr {
		go c.mon.MonAddr(mon2.PerAddrCfg{
			Addr:    c.pegContracts.vault.GetAddr(),
			ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
			AbiStr:  c.pegContracts.vault.GetABI(), // to parse event name by topics[0]
		}, c.vaultEvCallback)
	}

	if c.pegContracts.vault2.GetAddr() != eth.ZeroAddr {
		go c.mon.MonAddr(mon2.PerAddrCfg{
			Addr:    c.pegContracts.vault2.GetAddr(),
			ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
			AbiStr:  c.pegContracts.vault2.GetABI(), // to parse event name by topics[0]
		}, c.vault2EvCallback)
	}
}

func (c *CbrOneChain) vaultEvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "Deposited":
		c.pbrDepositMonCb(elog, 0)
	case "Withdrawn":
		c.pbrWithdrawnMonCb(elog, 0)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) vault2EvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "Deposited":
		c.pbrDepositMonCb(elog, 2)
	case "Withdrawn":
		c.pbrWithdrawnMonCb(elog, 2)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) pbrDepositMonCb(eLog ethtypes.Log, version uint32) {
	if version == 2 {
		ev, err := c.pegContracts.vault2.ParseDeposited(eLog)
		if err != nil {
			log.Errorf("monPegbrV2Deposited: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		if relayerInstance.isEthAddrBlocked(ev.Depositor, ev.MintAccount) {
			log.Warnln("eth addrs blocked", ev.String())
			return
		}
	} else {
		ev, err := c.pegContracts.vault.ParseDeposited(eLog)
		if err != nil {
			log.Errorf("monPegbrDeposited: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		if relayerInstance.isEthAddrBlocked(ev.Depositor, ev.MintAccount) {
			log.Warnln("eth addrs blocked", ev.String())
			return
		}
	}
	err := c.saveEvent(pegtypes.PegbrEventDeposited, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) pbrWithdrawnMonCb(eLog ethtypes.Log, version uint32) {
	var refChainId uint64
	var refId []byte
	if version == 2 {
		ev, err := c.pegContracts.vault2.ParseWithdrawn(eLog)
		if err != nil {
			log.Errorf("monPegbrV2Withdrawn: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		refChainId = ev.RefChainId
		refId = ev.RefId[:]
	} else {
		ev, err := c.pegContracts.vault.ParseWithdrawn(eLog)
		if err != nil {
			log.Errorf("monPegbrWithdrawn: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		refChainId = ev.RefChainId
		refId = ev.RefId[:]
	}
	// delete to-submit withdraw at local if have, as it's been submitted (by other nodes or me)
	relayerInstance.dbDelete(GetPegbrWdKey(c.chainid, refChainId, refId))

	err := c.saveEvent(pegtypes.PegbrEventWithdrawn, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) MonPegBridge() {
	if c.pegContracts.bridge.GetAddr() != eth.ZeroAddr {
		go c.mon.MonAddr(mon2.PerAddrCfg{
			Addr:    c.pegContracts.bridge.GetAddr(),
			ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
			AbiStr:  c.pegContracts.bridge.GetABI(), // to parse event name by topics[0]
		}, c.pegbrEvCallback)
	}

	if c.pegContracts.bridge2.GetAddr() != eth.ZeroAddr {
		go c.mon.MonAddr(mon2.PerAddrCfg{
			Addr:    c.pegContracts.bridge2.GetAddr(),
			ChkIntv: 4 * time.Duration(c.blkInterval) * time.Second,
			AbiStr:  c.pegContracts.bridge2.GetABI(), // to parse event name by topics[0]
		}, c.pegbr2EvCallback)
	}
}

func (c *CbrOneChain) pegbrEvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "Mint":
		c.pbrMintMonCb(elog, 0)
	case "Burn":
		c.pbrBurnMonCb(elog, 0)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) pegbr2EvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "Mint":
		c.pbrMintMonCb(elog, 2)
	case "Burn":
		c.pbrBurnMonCb(elog, 2)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (c *CbrOneChain) pbrMintMonCb(eLog ethtypes.Log, version uint32) {
	var refChainId uint64
	var refId []byte
	if version == 2 {
		ev, err := c.pegContracts.bridge2.ParseMint(eLog)
		if err != nil {
			log.Errorf("monPegbrV2Mint: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		refChainId = ev.RefChainId
		refId = ev.RefId[:]
	} else {
		ev, err := c.pegContracts.bridge.ParseMint(eLog)
		if err != nil {
			log.Errorf("monPegbrMint: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		refChainId = ev.RefChainId
		refId = ev.RefId[:]
	}

	// delete to-submit mint at local if have, as it's been submitted (by other nodes or me)
	relayerInstance.dbDelete(GetPegbrMintKey(c.chainid, refChainId, refId[:]))

	err := c.saveEvent(pegtypes.PegbrEventMint, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

func (c *CbrOneChain) pbrBurnMonCb(eLog ethtypes.Log, version uint32) {
	if version == 2 {
		ev, err := c.pegContracts.bridge2.ParseBurn(eLog)
		if err != nil {
			log.Errorf("monPegbrV2Burn: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		if relayerInstance.isEthAddrBlocked(ev.Account, ev.ToAccount) {
			log.Warnln("eth addrs blocked", ev.String())
			return
		}
	} else {
		ev, err := c.pegContracts.bridge.ParseBurn(eLog)
		if err != nil {
			log.Errorf("monPegbrBurn: chain %d cannot parse event: %s", c.chainid, err)
			return
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		if relayerInstance.isEthAddrBlocked(ev.Account, ev.WithdrawAccount) {
			log.Warnln("eth addrs blocked", ev.String())
			return
		}
	}
	err := c.saveEvent(pegtypes.PegbrEventBurn, eLog)
	if err != nil {
		log.Errorln("saveEvent err:", err)
	}
}

// SendMint sends mint tx onchain to PeggedTokenBridge contract, no wait mine
func (c *CbrOneChain) SendMint(
	mintBytes []byte, sigs [][]byte, curss currentSigners, mint *pegbrtypes.MintOnChain, bridgeVersion uint32) (string, error) {
	logmsg := fmt.Sprintf(
		"mint %s of token %x for user %x, refChainId %d, refId %x, mintChainId %d, depositor %x",
		new(big.Int).SetBytes(mint.GetAmount()).String(), mint.GetToken(), mint.GetAccount(), mint.GetRefChainId(), mint.GetRefId(), c.chainid, mint.GetDepositor())
	if types.IsFlowChain(c.chainid) {
		go c.FlowClient.sendMint(logmsg, mintBytes, string(mint.Token), sigs)
		return "", nil
	}
	// EVM chains
	err := c.checkPendingNonce()
	if err != nil {
		return "", fmt.Errorf("Pending nonce check failed. %w", err)
	}
	tx, err := c.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("Mint transaction succeeded, tx %x. %s", receipt.TxHash, logmsg)
				} else {
					log.Warnf("Mint transaction failed, tx %x. %s", receipt.TxHash, logmsg)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("Mint transaction err: %s. %s", err, logmsg)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			if bridgeVersion == 0 {
				return c.pegContracts.bridge.Mint(opts, mintBytes, sigs, curss.addrs, curss.powers)
			}
			return c.pegContracts.bridge2.Mint(opts, mintBytes, sigs, curss.addrs, curss.powers)
		},
	)
	if err != nil {
		return "", fmt.Errorf("%s. Failed to send mint: %w", logmsg, err)
	}

	return tx.Hash().Hex(), nil
}

// SendWithdraw sends withdraw tx onchain to OriginalTokenVault contract, no wait mine
func (c *CbrOneChain) SendWithdraw(
	wdBytes []byte, sigs [][]byte, curss currentSigners, withdraw *pegbrtypes.WithdrawOnChain, vaultVersion uint32) (string, error) {
	logmsg := fmt.Sprintf(
		"withdraw %s of token %x for user %x, refChainId %d, refId %x, withdrawChainId %d, burnAccount %x",
		new(big.Int).SetBytes(withdraw.GetAmount()).String(), withdraw.GetToken(), withdraw.GetReceiver(), withdraw.GetRefChainId(), withdraw.GetRefId(), c.chainid, withdraw.BurnAccount)
	if types.IsFlowChain(c.chainid) {
		go c.FlowClient.sendWithdraw(logmsg, wdBytes, string(withdraw.Token), sigs)
		return "", nil
	}
	// EVM chains
	err := c.checkPendingNonce()
	if err != nil {
		return "", fmt.Errorf("Pending nonce check failed. %w", err)
	}
	tx, err := c.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("Withdraw transaction succeeded, tx %x. %s", receipt.TxHash, logmsg)
				} else {
					log.Warnf("Withdraw transaction failed, tx %x. %s", receipt.TxHash, logmsg)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("Withdraw transaction err: %s. %s", err, logmsg)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			if vaultVersion == 0 {
				return c.pegContracts.vault.Withdraw(opts, wdBytes, sigs, curss.addrs, curss.powers)
			}
			return c.pegContracts.vault2.Withdraw(opts, wdBytes, sigs, curss.addrs, curss.powers)
		},
	)
	if err != nil {
		return "", fmt.Errorf("%s. Failed to send withraw: %w", logmsg, err)
	}

	return tx.Hash().Hex(), nil
}
