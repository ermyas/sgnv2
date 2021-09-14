package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyUpdate(ctx sdk.Context, update *types.PendingUpdate) bool {
	var applied bool
	var err error
	switch update.Type {
	case types.DataType_EthBlkNum:
		applied, err = k.applyEthBlkNum(ctx, update)
	case types.DataType_StakingContractParam:
		applied, err = k.applyStakingContractParam(ctx, update)
	case types.DataType_ValidatorSgnAddr:
		applied, err = k.applyValidatorSgnAddr(ctx, update)
	case types.DataType_ValidatorParams:
		applied, err = k.applyValidatorParams(ctx, update)
	case types.DataType_ValidatorStates:
		applied, err = k.applyValidatorStates(ctx, update)
	case types.DataType_DelegatorShares:
		applied, err = k.applyDelegatorShares(ctx, update)
	}

	if err != nil {
		log.Errorln("Apply update err:", err)
	}
	return applied
}

func (k Keeper) applyEthBlkNum(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	return true, nil
}

func (k Keeper) applyStakingContractParam(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	return true, nil
}

func (k Keeper) applyValidatorSgnAddr(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator sgn addr %s", v.String())
	// TODO: remove previous sgnaddr account
	acct, err := sdk.AccAddressFromBech32(v.SgnAddress)
	if err != nil {
		return false, err
	}
	err = k.stakingKeeper.InitAccount(ctx, acct)
	if err != nil {
		return false, err
	}
	return true, nil
}

// TODO: handle/restrict sgnaddr/consaddr update
func (k Keeper) applyValidatorParams(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	if update.Proposer != v.SgnAddress {
		return false, fmt.Errorf("Validator %s not msg sender: %s", v.EthAddress, update.Proposer)
	}
	if v.ConsensusPubkey == nil {
		return false, fmt.Errorf("empty consensus pub key")
	}
	log.Infof("Apply validator params %s", v.String())
	val, found := k.stakingKeeper.GetValidator(ctx, eth.Hex2Addr(v.EthAddress))
	if !found {
		val = *stakingtypes.NewValidator(v.EthAddress, v.EthSigner, v.SgnAddress)
		val.Description = v.Description
	} else {
		val.EthSigner = eth.FormatAddrHex(v.EthSigner)
		val.SgnAddress = v.SgnAddress
	}
	val.ConsensusPubkey = v.ConsensusPubkey
	val.CommissionRate = v.CommissionRate
	k.stakingKeeper.SetValidatorParams(ctx, &val, !found)
	//TODO: gas coins
	return true, nil
}

func (k Keeper) applyValidatorStates(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator states %s", v.String())
	val, found := k.stakingKeeper.GetValidator(ctx, eth.Hex2Addr(v.EthAddress))
	if !found {
		return false, fmt.Errorf("validator %s not found", val.EthAddress)
	}
	val.Status = v.Status
	val.Tokens = v.Tokens
	val.DelegatorShares = v.DelegatorShares

	k.stakingKeeper.SetValidatorStates(ctx, &val)
	return true, nil
}

func (k Keeper) applyDelegatorShares(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	d, err := stakingtypes.UnmarshalDelegation(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply delegator shares valAddr %s delAddr %s shares %s", d.ValidatorAddress, d.DelegatorAddress, d.Shares)
	k.stakingKeeper.SetDelegationShares(
		ctx,
		eth.Hex2Addr(d.DelegatorAddress),
		eth.Hex2Addr(d.ValidatorAddress),
		d.Shares)
	return true, nil
}
