package monitor

import (
	"fmt"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (m *Monitor) monitorEthValidatorNotice() {
	_, err := m.ethMonitor.Monitor(
		&monitor.Config{
			EventName:     eth.EventValidatorNotice,
			Contract:      m.EthClient.Contracts.Staking,
			StartBlock:    m.startEthBlock,
			Reset:         true,
			CheckInterval: getEventCheckInterval(eth.EventValidatorNotice),
		},
		func(cb monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
			e, err := m.EthClient.Contracts.Staking.ParseValidatorNotice(eLog)
			if err != nil {
				log.Errorln("parse event err", err)
			}
			if e.From != eth.ZeroAddr && e.From != m.EthClient.Contracts.Sgn.Address {
				return
			}
			log.Infof("Catch event ValidatorNotice %s, tx hash: %x, blknum: %d", e.Key, eLog.TxHash, eLog.BlockNumber)
			event := eth.NewEvent(eth.EventValidatorNotice, eLog)
			err = m.dbSet(GetPullerKey(eLog), event.MustMarshal())
			if err != nil {
				log.Errorln("db Set err", err)
			}
			return false
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Monitor) monitorEthValidatorStatusUpdate() {
	_, err := m.ethMonitor.Monitor(
		&monitor.Config{
			EventName:     eth.EventValidatorStatusUpdate,
			Contract:      m.EthClient.Contracts.Staking,
			StartBlock:    m.startEthBlock,
			Reset:         true,
			CheckInterval: getEventCheckInterval(eth.EventValidatorStatusUpdate),
		},
		func(cb monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
			logmsg := fmt.Sprintf("Catch event ValidatorStatusUpdate, tx hash: %x, blknum: %d", eLog.TxHash, eLog.BlockNumber)
			e, err := m.EthClient.Contracts.Staking.ParseValidatorStatusUpdate(eLog)
			if err != nil {
				log.Errorf("%s. parse event err: %s", logmsg, err)
			}
			if e.Status == eth.Bonded {
				m.setBootstrapped()
				if e.ValAddr == m.valAddr {
					log.Infof("%s. Init my own validator.", logmsg)
					m.setBonded()
					go m.selfSyncValidator()
				} else {
					log.Infof("%s. Validator %x bonded.", logmsg, e.ValAddr)
				}
			} else {
				// only put unbonding or unbonded event to puller queue
				log.Infof("%s. Validator %x %s.", logmsg, e.ValAddr, eth.ParseValStatus(e.Status))
				if e.ValAddr == m.valAddr {
					m.clearBonded()
				}
				event := eth.NewEvent(eth.EventValidatorStatusUpdate, eLog)
				err = m.dbSet(GetPullerKey(eLog), event.MustMarshal())
				if err != nil {
					log.Errorln("db Set err", err)
				}
			}
			return false
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Monitor) monitorEthDelegationUpdate() {
	_, err := m.ethMonitor.Monitor(
		&monitor.Config{
			EventName:     eth.EventDelegationUpdate,
			Contract:      m.EthClient.Contracts.Staking,
			StartBlock:    m.startEthBlock,
			Reset:         true,
			CheckInterval: getEventCheckInterval(eth.EventDelegationUpdate),
		},
		func(cb monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
			log.Infof("Catch event DelegationUpdate, tx hash: %x, blknum: %d", eLog.TxHash, eLog.BlockNumber)
			event := eth.NewEvent(eth.EventDelegationUpdate, eLog)
			err := m.dbSet(GetPullerKey(eLog), event.MustMarshal())
			if err != nil {
				log.Errorln("db Set err", err)
			}
			if !m.isBonded() {
				e, err2 := m.EthClient.Contracts.Staking.ParseDelegationUpdate(eLog)
				if err2 != nil {
					log.Errorln("parse event err", err2)
					return
				}
				if e.ValAddr == m.valAddr && m.shouldBondValidator() {
					m.bondValidator()
				}
			}
			return false
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Monitor) shouldBondValidator() bool {
	// TODO: use eth dry run?
	validator, err := m.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, m.valAddr)
	if err != nil {
		log.Errorln("get validator err", err)
		return false
	}

	if validator.Status == 0 {
		log.Debug("Validator not initialized")
		return false
	}

	if validator.Status == eth.Bonded {
		log.Debug("Validator already  bonded")
		return false
	}

	hasMinRequiredTokens, err :=
		m.EthClient.Contracts.Staking.HasMinRequiredTokens(&bind.CallOpts{}, m.valAddr, true)
	if err != nil {
		log.Errorln("Get min required tokens err", err)
		return false
	}
	if !hasMinRequiredTokens {
		log.Debug("Not have min required tokens")
		return false
	}

	minValTokens, err := m.EthClient.Contracts.Staking.GetMinValidatorTokens(&bind.CallOpts{})
	if err != nil {
		log.Errorln("Get min validator tokens err", err)
		return false
	}
	if validator.Tokens.Cmp(minValTokens) <= 0 {
		log.Debugf("Token less than current min validator tokens: %s < %s", validator.Tokens, minValTokens)
		return false
	}

	currBlkNum := m.ethMonitor.GetCurrentBlockNumber()
	if currBlkNum.Cmp(validator.BondBlock) < 0 {
		log.Debugf("Not validator bond block %d yet", validator.BondBlock)
		return false
	}

	nextBondBlock, err := m.EthClient.Contracts.Staking.NextBondBlock(&bind.CallOpts{})
	if err != nil {
		log.Errorln("Get next bond block err", err)
		return false
	}
	if currBlkNum.Cmp(nextBondBlock) < 0 {
		log.Debugf("Not next bond block %d yet", validator.BondBlock)
		return false
	}

	sgnAddr, err := m.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, m.valAddr)
	if err != nil {
		log.Errorln("Get sgn addr err", err)
		return false
	}
	if !sdk.AccAddress(sgnAddr).Equals(m.sgnAcct) {
		log.Debugf("sidechain address not match, %s %s", sdk.AccAddress(sgnAddr), m.sgnAcct)
		return false
	}

	return true
}

func (m *Monitor) bondValidator() {
	_, err := m.EthClient.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("BondValidator transaction %x succeeded", receipt.TxHash)
				} else {
					log.Errorf("BondValidator transaction %x failed", receipt.TxHash)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Errorf("BondValidator transaction %x err: %s", tx.Hash(), err)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return m.EthClient.Contracts.Staking.BondValidator(opts)
		},
	)
	if err != nil {
		log.Errorln("BondValidator tx err", err)
		return
	}
	log.Infof("Bond validator %x on mainchain", m.valAddr)
}

func (m *Monitor) selfSyncValidator() {
	var i int
	for i = 1; i < 5; i++ {
		updated := m.SyncValidator(m.EthClient.Address, m.ethMonitor.GetCurrentBlockNumber())
		if updated {
			return
		}
		time.Sleep(60 * time.Second)
	}
	log.Warn("self validator not synced yet")
}
