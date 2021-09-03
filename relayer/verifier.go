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
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
)

func (r *Relayer) verifyPendingUpdates() {
	v, _ := validatorcli.QuerySdkValidator(r.Transactor.CliCtx, r.Transactor.Key.GetAddress().String())
	if v == nil || v.Status != stakingtypes.Bonded {
		log.Traceln("skip verifying pending updates as I am not a bonded validator")
		return
	}
	pendingUpdates, err := synccli.QueryPendingUpdates(r.Transactor.CliCtx)
	if err != nil {
		log.Errorln("Query pending updates error:", err)
		return
	}

	msgs := synctypes.MsgVoteUpdates{
		Votes:  make([]*synctypes.VoteUpdate, 0),
		Sender: r.Transactor.Key.GetAddress().String(),
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
				log.Errorln("verifiedUpdates Set err", err)
				continue
			}
			if approve {
				msgs.Votes = append(msgs.Votes, &synctypes.VoteUpdate{
					Id:     update.Id,
					Option: synctypes.VoteOption_Yes,
				})
			}
		}
	}

	if len(msgs.Votes) > 0 {
		r.Transactor.AddTxMsg(&msgs)
	}
}

func (r *Relayer) verifyUpdate(update *synctypes.PendingUpdate) (done, approve bool) {
	switch update.Type {
	case synctypes.DataType_EthBlkNum:
		return r.verifyEthBlkNum(update)
	case synctypes.DataType_StakingContractParam:
		return r.verifyStakingContractParam(update)
	case synctypes.DataType_ValidatorParams:
		return r.verifyValidatorParams(update)
	case synctypes.DataType_ValidatorStates:
		return r.verifyValidatorStates(update)
	case synctypes.DataType_DelegatorShares:
		return r.verifyDelegatorShares(update)
	default:
		return false, false
	}
}

func (r *Relayer) verifyEthBlkNum(update *synctypes.PendingUpdate) (done, approve bool) {
	log.Infof("Verify sync eth block: %d", update.EthBlock)
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

func (r *Relayer) verifyValidatorParams(update *synctypes.PendingUpdate) (done, approve bool) {
	updateVal, err := validatortypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, params for validator: %s, signer %s, sgnaddr %s, commission %d",
		update.Id, updateVal.EthAddress, updateVal.EthSigner, updateVal.SgnAddress, updateVal.CommissionRate)

	storeVal, err := validatorcli.QueryValidator(r.Transactor.CliCtx, updateVal.EthAddress)
	if err == nil {
		if updateVal.EthSigner == storeVal.EthSigner && updateVal.SgnAddress == storeVal.SgnAddress &&
			updateVal.CommissionRate == storeVal.CommissionRate {
			log.Infof("%s. validator params already updated", logmsg)
			return true, false
		}
	}

	ethVal, err := r.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query validator info err: %s", logmsg, err)
		return false, false
	}

	sgnAddr, err := r.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query sgn address err: %s", logmsg, err)
		return false, false
	}

	if updateVal.EthSigner != eth.Addr2Hex(ethVal.Signer) ||
		updateVal.SgnAddress != sdk.AccAddress(sgnAddr).String() ||
		updateVal.CommissionRate != ethVal.CommissionRate {
		values := fmt.Sprintf("signer %x sgnaddr %s commission %d",
			ethVal.Signer, sdk.AccAddress(sgnAddr).String(), ethVal.CommissionRate)
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Errorf("%s. validator params not match eth values: %s", logmsg, values)
			return true, false
		}
		log.Infof("%s. eth block not passed, values: %s", logmsg, values)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyValidatorStates(update *synctypes.PendingUpdate) (done, approve bool) {
	updateVal, err := validatortypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, states for validator: %s, status %s, tokens %s, shares %s",
		update.Id, updateVal.EthAddress, updateVal.Status, updateVal.Tokens, updateVal.Shares)

	storeVal, err := validatorcli.QueryValidator(r.Transactor.CliCtx, updateVal.EthAddress)
	if err == nil {
		if updateVal.Status == storeVal.Status && updateVal.Tokens == storeVal.Tokens && updateVal.Shares == storeVal.Shares {
			log.Infof("%s. states already updated", logmsg)
			return true, false
		}
	}

	ethVal, err := r.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query validator info err: %s", logmsg, err)
		return false, false
	}

	if updateVal.Status != validatortypes.ValidatorStatus(ethVal.Status) ||
		updateVal.Tokens != ethVal.Tokens.String() || updateVal.Shares != ethVal.Shares.String() {
		values := fmt.Sprintf("status %s tokens %s shares %s", eth.ParseValStatus(ethVal.Status), ethVal.Tokens, ethVal.Shares)
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Infof("%s. validator states not match eth values: %s", logmsg, values)
			return true, false
		}
		log.Infof("%s. eth block not passed, eth values: %s", logmsg, values)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyDelegatorShares(update *synctypes.PendingUpdate) (done, approve bool) {
	updateDel, err := validatortypes.UnmarshalDelegator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, shares for delegator: %s", update.Id, updateDel)

	storeDel, err := validatorcli.QueryDelegator(r.Transactor.CliCtx, updateDel.ValAddress, updateDel.DelAddress)
	if err == nil {
		if updateDel.Shares == storeDel.Shares {
			log.Infof("%s. shares already updated", logmsg)
			return true, false
		}
	}

	ethDel, err := r.EthClient.Contracts.Staking.GetDelegatorInfo(
		&bind.CallOpts{}, eth.Hex2Addr(updateDel.ValAddress), eth.Hex2Addr(updateDel.DelAddress))
	if err != nil {
		log.Errorf("%s. query delegator info err: %s", logmsg, err)
		return false, false
	}

	if updateDel.Shares != ethDel.Shares.String() {
		if r.cmpBlkNum(update.EthBlock) == 1 {
			log.Infof("%s. delegator shares not match eth value: %s", logmsg, ethDel.Shares.String())
			return true, false
		}
		log.Infof("%s. eth block not passed, eth value: %s", logmsg, ethDel.Shares.String())
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
