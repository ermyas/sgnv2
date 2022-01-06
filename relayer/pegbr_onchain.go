package relayer

import (
	"fmt"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	pegtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) monPegbrDeposited(blk *big.Int) {
	if c.pegContracts.vault.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventDeposited,
		Contract:     c.pegContracts.vault,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.pegContracts.vault.ParseDeposited(eLog)
		if err != nil {
			log.Errorf("monPegbrDeposited: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(pegtypes.PegbrEventDeposited, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monPegbrMint(blk *big.Int) {
	if c.pegContracts.bridge.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventMint,
		Contract:     c.pegContracts.bridge,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.pegContracts.bridge.ParseMint(eLog)
		if err != nil {
			log.Errorf("monPegbrMint: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		// delete to-submit mint at local if have, as it's been submitted (by other nodes or me)
		if CurRelayerInstance == nil {
			log.Errorln("CurRelayerInstance not initialized", err)
		} else {
			CurRelayerInstance.dbDelete(GetPegbrMintKey(c.chainid, ev.RefChainId, ev.RefId[:]))
		}

		err = c.saveEvent(pegtypes.PegbrEventMint, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monPegbrBurn(blk *big.Int) {
	if c.pegContracts.bridge.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventBurn,
		Contract:     c.pegContracts.bridge,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.pegContracts.bridge.ParseBurn(eLog)
		if err != nil {
			log.Errorf("monPegbrBurn: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}

		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = c.saveEvent(pegtypes.PegbrEventBurn, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

func (c *CbrOneChain) monPegbrWithdrawn(blk *big.Int) {
	if c.pegContracts.vault.Address == eth.ZeroAddr {
		return
	}

	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventWithdrawn,
		Contract:     c.pegContracts.vault,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.pegContracts.vault.ParseWithdrawn(eLog)
		if err != nil {
			log.Errorf("monPegbrWithdrawn: chain %d cannot parse event: %s", c.chainid, err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		// delete to-submit withdraw at local if have, as it's been submitted (by other nodes or me)
		if CurRelayerInstance == nil {
			log.Errorln("CurRelayerInstance not initialized", err)
		} else {
			CurRelayerInstance.dbDelete(GetPegbrWdKey(c.chainid, ev.RefChainId, ev.RefId[:]))
		}

		err = c.saveEvent(pegtypes.PegbrEventWithdrawn, eLog)
		if err != nil {
			log.Errorln("saveEvent err:", err)
			return true // ask to recreate to process event again
		}
		return false
	})
}

// SendMint sends mint tx onchain to PeggedTokenBridge contract, no wait mine
func (c *CbrOneChain) SendMint(mintBytes []byte, sigs [][]byte, curss currentSigners, mint *pegbrtypes.MintOnChain) (string, error) {
	logmsg := fmt.Sprintf(
		"mint %s of token %x for user %x, refChainId %d, refId %x, mintChainId %d, depositor %x",
		new(big.Int).SetBytes(mint.GetAmount()).String(), mint.GetToken(), mint.GetAccount(), mint.GetRefChainId(), mint.GetRefId(), c.chainid, mint.GetDepositor())
	err := c.checkPendingNonce()
	if err != nil {
		log.Warnf("Pending nonce check failed: %s. %s", err, logmsg)
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
			return c.pegContracts.bridge.Mint(opts, mintBytes, sigs, curss.addrs, curss.powers)
		},
	)
	if err != nil {
		return "", fmt.Errorf("%s. Failed to send mint: %w", logmsg, err)
	}

	return tx.Hash().Hex(), nil
}

// SendWithdraw sends withdraw tx onchain to OriginalTokenVault contract, no wait mine
func (c *CbrOneChain) SendWithdraw(wdBytes []byte, sigs [][]byte, curss currentSigners, withdraw *pegbrtypes.WithdrawOnChain) (string, error) {
	logmsg := fmt.Sprintf(
		"withdraw %s of token %x for user %x, refChainId %d, refId %x, withdrawChainId %d, burnAccount %x",
		new(big.Int).SetBytes(withdraw.GetAmount()).String(), withdraw.GetToken(), withdraw.GetReceiver(), withdraw.GetRefChainId(), withdraw.GetRefId(), c.chainid, withdraw.BurnAccount)
	err := c.checkPendingNonce()
	if err != nil {
		log.Warnf("Pending nonce check failed: %s. %s", err, logmsg)
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
			return c.pegContracts.vault.Withdraw(opts, wdBytes, sigs, curss.addrs, curss.powers)
		},
	)
	if err != nil {
		return "", fmt.Errorf("%s. Failed to send withraw: %w", logmsg, err)
	}

	return tx.Hash().Hex(), nil
}
