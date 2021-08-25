package relayer

import (
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	synccli "github.com/celer-network/sgn-v2/x/sync/client/cli"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	validatorcli "github.com/celer-network/sgn-v2/x/validator/client/cli"
	validatortypes "github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
)

func (r *Relayer) verifyPendingUpdates() {
	v, _ := validatorcli.QueryValidator(r.Transactor.CliCtx, r.Transactor.Key.GetAddress().String())
	if v.GetStatus() != validatortypes.ValidatorStatus_Bonded {
		log.Traceln("skip verifying pending updates as I am not a bonded validator")
		return
	}
	pendingUpdates, err := synccli.QueryPendingUpdates(r.Transactor.CliCtx)
	if err != nil {
		log.Errorln("Query pending updates error:", err)
		return
	}

	msgs := synctypes.MsgProposeUpdates{
		Updates:  make([]*synctypes.ProposeUpdate, 0),
		EthBlock: r.getCurrentBlockNumber().Uint64(),
		Sender:   string(r.Transactor.Key.GetAddress()),
	}
	for _, update := range pendingUpdates {
		_, err = r.verifiedUpdates.Get(strconv.Itoa(int(update.Id)))
		if err == nil {
			continue
		}

		done, approve := r.verifyUpdate(update)
		if done {
			err = r.verifiedUpdates.Set(strconv.Itoa(int(update.Id)), []byte{})
			if err != nil {
				log.Errorln("verifiedChanges Set err", err)
				continue
			}
			if approve {
				msgs.Updates = append(msgs.Updates, &synctypes.ProposeUpdate{
					Type: update.Type,
					Data: update.Data,
				})
			}
		}
	}

	if len(msgs.Updates) > 0 {
		r.Transactor.AddTxMsg(&msgs)
	}
}

func (r *Relayer) verifyUpdate(update *synctypes.PendingUpdate) (done, approve bool) {
	switch update.Type {
	case synctypes.DataType_EthBlkNum:
		return r.verifyEthBlkNum(update)
	case synctypes.DataType_StakingContractParam:
		return r.verifyStakingContractParam(update)
	case synctypes.DataType_ValidatorAddrs:
		return r.verifyValidatorAddrs(update)
	case synctypes.DataType_ValidatorStates:
		return r.verifyValidatorStates(update)
	case synctypes.DataType_ValidatorCommissionRate:
		return r.verifyValidatorCommissionRate(update)
	case synctypes.DataType_DelegatorShares:
		return r.verifyDelegatorShares(update)
	default:
		return false, false
	}
}

func (r *Relayer) verifyEthBlkNum(update *synctypes.PendingUpdate) (done, approve bool) {
	log.Infof("Verify sync mainchain block: %d", update.EthBlock)
	accceptedBlkRange := viper.GetUint64(common.FlagEthAcceptedBlkRange)
	currentBlkNum := r.getCurrentBlockNumber().Uint64()

	if update.EthBlock-currentBlkNum < accceptedBlkRange || currentBlkNum-update.EthBlock < accceptedBlkRange {
		return true, true
	}

	return true, false
}

func (r *Relayer) verifyStakingContractParam(update *synctypes.PendingUpdate) (done, approve bool) {
	// TODO
	return true, true
}

func (r *Relayer) verifyValidatorAddrs(update *synctypes.PendingUpdate) (done, approve bool) {
	v, err := validatortypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}

	logmsg := fmt.Sprintf("verify update id %d, sidechain addr for validator: %s", update.Id, v.String())
	c, err := validatorcli.QueryValidator(r.Transactor.CliCtx, v.EthAddress)
	if err == nil {
		if v.SgnAddress == c.SgnAddress {
			log.Infof("%s. sgn addr already updated", logmsg)
			return true, false
		}
	}

	sgnAddr, err := r.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, eth.Hex2Addr(v.EthAddress))
	if err != nil {
		log.Errorf("%s. query sgn address err: %s", logmsg, err)
		return false, false
	}

	// TODO check the format...
	if v.SgnAddress != string(sgnAddr) {
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Infof("%s. validator sgn address not match mainchain value: %s", logmsg, sgnAddr)
			return true, false
		}
		log.Infof("%s. mainchain block not passed, validator addr: %s", logmsg, sgnAddr)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyValidatorStates(update *synctypes.PendingUpdate) (done, approve bool) {
	v, err := validatortypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}

	logmsg := fmt.Sprintf("verify update id %d, states for validator: %s", update.Id, v.String())
	c, err := validatorcli.QueryValidator(r.Transactor.CliCtx, v.EthAddress)
	if err == nil {
		if v.Status == c.Status && v.Tokens == c.Tokens && v.Shares == c.Shares {
			log.Infof("%s. states already updated", logmsg)
			return true, false
		}
	}

	vFromEth, err := r.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, eth.Hex2Addr(v.EthAddress))
	if err != nil {
		log.Errorf("%s. query validator info err: %s", logmsg, err)
		return false, false
	}

	if v.Status != validatortypes.ValidatorStatus(vFromEth.Status) || v.Tokens != vFromEth.Tokens.String() || v.Shares != vFromEth.Shares.String() {
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Infof("%s. validator states not match mainchain value: %s", logmsg, vFromEth)
			return true, false
		}
		log.Infof("%s. mainchain block not passed, mainchain value: %s", logmsg, vFromEth)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyValidatorCommissionRate(update *synctypes.PendingUpdate) (done, approve bool) {
	v, err := validatortypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}

	logmsg := fmt.Sprintf("verify update id %d, commission rate for validator: %s", update.Id, v.String())
	c, err := validatorcli.QueryValidator(r.Transactor.CliCtx, v.EthAddress)
	if err == nil {
		if v.CommissionRate == c.CommissionRate {
			log.Infof("%s. commission rate already updated", logmsg)
			return true, false
		}
	}

	vFromEth, err := r.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, eth.Hex2Addr(v.EthAddress))
	if err != nil {
		log.Errorf("%s. query validator info err: %s", logmsg, err)
		return false, false
	}

	if v.CommissionRate != vFromEth.CommissionRate.Uint64() {
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Infof("%s. validator commission rate not match mainchain value: %s", logmsg, vFromEth.CommissionRate.Uint64())
			return true, false
		}
		log.Infof("%s. mainchain block not passed, mainchain value: %s", logmsg, vFromEth.CommissionRate.Uint64())
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyDelegatorShares(update *synctypes.PendingUpdate) (done, approve bool) {
	d, err := validatortypes.UnmarshalDelegator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}

	logmsg := fmt.Sprintf("verify update id %d, shares for delegator: %s", update.Id, d)
	c, err := validatorcli.QueryDelegator(r.Transactor.CliCtx, d.ValAddress, d.DelAddress)
	if err == nil {
		if d.Shares == c.Shares {
			log.Infof("%s. shares already updated", logmsg)
			return true, false
		}
	}

	dFromEth, err := r.EthClient.Contracts.Staking.GetDelegatorInfo(&bind.CallOpts{}, eth.Hex2Addr(d.ValAddress), eth.Hex2Addr(d.DelAddress))
	if err != nil {
		log.Errorf("%s. query delegator info err: %s", logmsg, err)
		return false, false
	}

	if d.Shares != dFromEth.Shares.String() {
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Infof("%s. delegator shares not match mainchain value: %s", logmsg, dFromEth.Shares.String())
			return true, false
		}
		log.Infof("%s. mainchain block not passed, mainchain value: %s", logmsg, dFromEth.Shares.String())
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) cmpBlkNum(blkNum uint64) int8 {
	currentBlkNum := r.getCurrentBlockNumber().Uint64()
	if currentBlkNum > blkNum {
		return 1
	} else if currentBlkNum < blkNum {
		return -1
	}
	return 0
}
