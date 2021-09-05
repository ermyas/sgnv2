package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/sync/types"
	vtypes "github.com/celer-network/sgn-v2/x/validator/types"
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
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator sgn addr %s", v.String())
	// TODO: remove previous sgnaddr account
	acct, err := vtypes.SdkAccAddrFromSgnBech32(v.SgnAddress)
	if err != nil {
		return false, err
	}
	err = k.valKeeper.InitAccount(ctx, acct)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (k Keeper) applyValidatorParams(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
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
	val, found := k.valKeeper.GetValidator(ctx, v.EthAddress)
	if !found {
		val = vtypes.NewValidator(v.EthAddress, v.EthSigner, v.SgnAddress)
		val.Description = v.Description
	} else {
		val.EthSigner = v.EthSigner
		val.SgnAddress = v.SgnAddress
	}
	val.ConsensusPubkey = v.ConsensusPubkey
	val.CommissionRate = v.CommissionRate
	k.valKeeper.SetValidator(ctx, val)
	//TODO: gas coins
	return true, nil
}

func (k Keeper) applyValidatorStates(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator states %s", v.String())
	_, ok := sdk.NewIntFromString(v.Tokens)
	if !ok {
		return false, fmt.Errorf("invalid tokens %s", v.Tokens)
	}
	_, ok = sdk.NewIntFromString(v.Shares)
	if !ok {
		return false, fmt.Errorf("invalid shares %s", v.Shares)
	}
	val, found := k.valKeeper.GetValidator(ctx, v.EthAddress)
	if !found {
		return false, fmt.Errorf("validator %s not found", val.EthAddress)
	}
	val.Status = v.Status
	val.Tokens = v.Tokens
	val.Shares = v.Shares

	err = k.valKeeper.SetValidatorStates(ctx, val)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (k Keeper) applyDelegatorShares(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	d, err := vtypes.UnmarshalDelegator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply delegator shares valAddr %s delAddr %s shares %s", d.ValAddress, d.DelAddress, d.Shares)
	err = k.valKeeper.SetDelegatorShares(ctx, d.ValAddress, d.DelAddress, d.Shares)
	if err != nil {
		return false, err
	}
	return true, nil
}
