package relayer

import (
	"fmt"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	validatorcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

func (r *Relayer) MonStaking() {
	go r.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    r.EthClient.Contracts.Staking.GetAddr(),
		ChkIntv: 4 * time.Duration(viper.GetUint64(common.FlagEthPollInterval)) * time.Second,
		AbiStr:  r.EthClient.Contracts.Staking.GetABI(), // to parse event name by topics[0]
	}, r.stakingEvCallback)
}

func (r *Relayer) stakingEvCallback(evname string, elog ethtypes.Log) {
	switch evname {
	case "ValidatorNotice":
		r.handleValidatorNotice(elog)
	case "ValidatorStatusUpdate":
		r.handleEthValidatorStatusUpdate(elog)
	case "DelegationUpdate":
		r.handleEthDelegationUpdate(elog)
	default:
		log.Infoln("unsupported evname: ", evname)
		return
	}
}

func (r *Relayer) handleValidatorNotice(eLog ethtypes.Log) {
	e, err := r.EthClient.Contracts.Staking.ParseValidatorNotice(eLog)
	if err != nil {
		log.Errorln("parse event err", err)
		return
	}
	if !(e.From == eth.ZeroAddr || e.From == r.EthClient.Contracts.Sgn.GetAddr()) {
		return
	}
	log.Infof("Catch event ValidatorNotice %s, val addr: %x tx hash: %x, blknum: %d",
		e.Key, e.ValAddr, eLog.TxHash, eLog.BlockNumber)
	if e.Key == "sgn-addr" || e.Key == "signer" || e.Key == "commission" {
		if e.Key == "sgn-addr" {
			// TODO: handle non-first-time sgn-addr update
			if r.chainMonitorStatus != ChainMonitorStatusNo {
				event := eth.NewEvent(eth.EventValidatorNotice, eLog)
				err = r.dbSet(GetPullerKey(eLog), event.MustMarshal())
				if err != nil {
					log.Errorln("db Set err", err)
				}
			}
			if e.ValAddr == r.Operator.ValAddr {
				if !r.isBonded() && r.shouldBondValidator() {
					go r.bondValidator()
				}
			}
		}
		if e.ValAddr == r.Operator.ValAddr {
			log.Debug("Self sync validator params")
			go r.selfSyncValidatorParams()
			if e.Key == "sgn-addr" && r.isBootstrapped() {
				go r.selfSyncValidatorStates()
			}
		}
	}
}

func (r *Relayer) handleEthValidatorStatusUpdate(eLog ethtypes.Log) {
	e, err := r.EthClient.Contracts.Staking.ParseValidatorStatusUpdate(eLog)
	if err != nil {
		log.Errorf("parse event err: %s", err)
		return
	}
	logmsg := fmt.Sprintf("Catch event ValidatorStatusUpdate, val addr: %x, status: %s, tx hash: %x, blknum: %d",
		e.ValAddr, eth.ParseValStatus(e.Status), eLog.TxHash, eLog.BlockNumber)

	if e.Status == eth.Bonded {
		r.setBootstrapped()
		if e.ValAddr == r.Operator.ValAddr {
			log.Infof("%s. Self sync bonded validator.", logmsg)
			r.setBonded()
			go r.selfSyncValidatorStates()
		} else {
			log.Infof("%s. Skip", logmsg)
		}
	} else {
		// only put unbonding or unbonded event to puller queue
		log.Infof("%s. Put in queue", logmsg)
		if e.ValAddr == r.Operator.ValAddr {
			r.clearBonded()
		}
		if r.chainMonitorStatus != ChainMonitorStatusNo {
			event := eth.NewEvent(eth.EventValidatorStatusUpdate, eLog)
			err = r.dbSet(GetPullerKey(eLog), event.MustMarshal())
			if err != nil {
				log.Errorln("db Set err", err)
			}
		}
	}
}

func (r *Relayer) handleEthDelegationUpdate(eLog ethtypes.Log) {
	log.Infof("Catch event DelegationUpdate, tx hash: %x, blknum: %d", eLog.TxHash, eLog.BlockNumber)
	if r.chainMonitorStatus != ChainMonitorStatusNo {
		event := eth.NewEvent(eth.EventDelegationUpdate, eLog)
		err := r.dbSet(GetPullerKey(eLog), event.MustMarshal())
		if err != nil {
			log.Errorln("db Set err", err)
		}
	}
	if !r.isBonded() {
		e, err2 := r.EthClient.Contracts.Staking.ParseDelegationUpdate(eLog)
		if err2 != nil {
			log.Errorln("parse event err", err2)
			return
		}
		if e.ValAddr == r.Operator.ValAddr && r.shouldBondValidator() {
			r.bondValidator()
		}
	}
}

func (r *Relayer) shouldBondValidator() bool {
	shouldBond, err := r.EthClient.Contracts.Viewer.ShouldBondValidator(&bind.CallOpts{}, r.Operator.ValAddr)
	if err != nil {
		log.Errorln("Check if should bond validator err:", err)
		return false
	}

	if !shouldBond {
		log.Debug("Validator not ready to be bonded")
		return false
	}

	sgnAddr, err := r.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, r.Operator.ValAddr)
	if err != nil {
		log.Errorln("Get sgn addr err", err)
		return false
	}
	if !sdk.AccAddress(sgnAddr).Equals(r.sgnAcct) {
		log.Debugf("sgn addr not match, %s, %s", sdk.AccAddress(sgnAddr), r.sgnAcct)
		return false
	}

	return true
}

func (r *Relayer) bondValidator() {
	_, err := r.EthClient.Transactor.Transact(
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
			return r.EthClient.Contracts.Staking.BondValidator(opts)
		},
	)
	if err != nil {
		log.Errorln("BondValidator tx err", err)
		return
	}
	log.Infof("Bond validator %x on mainchain", r.Operator.ValAddr)
}

func (r *Relayer) selfSyncValidatorStates() {
	r.selfSyncValidator(ValSyncOptions{states: true})
}

func (r *Relayer) selfSyncValidatorParams() {
	r.selfSyncValidator(ValSyncOptions{params: true})
}

func (r *Relayer) selfSyncValidator(options ValSyncOptions) {
	acctFound := r.waitForSgnAccountFound()
	if !acctFound {
		log.Errorf("Sgn account %s not found", r.sgnAcct.String())
		return
	}
	if options.states {
		valFound := r.waitForValidatorFound()
		if !valFound {
			log.Errorf("Validator %x not found", r.Operator.ValAddr)
			return
		}
	}
	for i := 0; i < 20; i++ {
		updated := r.SyncValidator(r.Operator.ValAddr, r.getCurrentBlockNumber().Uint64(), options)
		if updated {
			log.Debugln("Self validator synced", options)
			return
		}
		time.Sleep(30 * time.Second)
	}
	log.Warn("Self validator not synced")
}

func (r *Relayer) waitForSgnAccountFound() bool {
	var acctFound bool
	for i := 0; i < 50; i++ {
		exist, _ := validatorcli.QuerySgnAccount(r.Transactor.CliCtx, r.sgnAcct.String())
		if exist {
			log.Debugf("Sgn account %s found", r.sgnAcct.String())
			acctFound = true
			break
		}
		time.Sleep(5 * time.Second)
	}
	return acctFound
}

func (r *Relayer) waitForValidatorFound() bool {
	var valFound bool
	for i := 0; i < 50; i++ {
		storeVal, _ := validatorcli.QueryValidator(r.Transactor.CliCtx, r.Operator.ValAddr.Hex())
		if storeVal != nil {
			log.Debugf("Validator %x found", r.Operator.ValAddr)
			valFound = true
			break
		}
		time.Sleep(5 * time.Second)
	}
	return valFound
}
