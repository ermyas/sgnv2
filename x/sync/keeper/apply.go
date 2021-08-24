package keeper

import (
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
	case types.DataType_ValidatorAddrs:
		applied, err = k.applyValidatorAddrs(ctx, update)
	case types.DataType_ValidatorStates:
		applied, err = k.applyValidatorStates(ctx, update)
	case types.DataType_ValidatorCommissionRate:
		applied, err = k.applyValidatorCommissionRate(ctx, update)
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

func (k Keeper) applyValidatorAddrs(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator %s signer %s sgnaddr %s", v.EthAddress, v.EthSigner, v.SgnAddress)
	val, found := k.valKeeper.GetValidator(ctx, v.EthAddress)
	if found {
		val.EthSigner = v.EthSigner
		val.SgnAddress = v.SgnAddress
	} else {
		val = vtypes.NewValidator(v.EthAddress, v.EthSigner, v.SgnAddress)
	}
	acct, err := vtypes.SdkAccAddrFromSgnBech32(v.SgnAddress)
	if err != nil {
		return false, err
	}
	err = k.valKeeper.InitAccount(ctx, acct)
	if err != nil {
		return false, err
	}
	k.valKeeper.SetValidator(ctx, &v)

	//TODO: gas coins
	return true, nil
}

func (k Keeper) applyValidatorStates(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := vtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator %s status %s tokens %s shares %s", v.EthAddress, v.Status, v.Tokens, v.Shares)
	err = k.valKeeper.SetValidatorStates(ctx, v.EthAddress, v.Status, v.Tokens, v.Shares)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (k Keeper) applyValidatorCommissionRate(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
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
