package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k Keeper) GetValidator(ctx sdk.Context, ethAddr string) (validator *types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorKey(ethAddr))
	if value == nil {
		return validator, false
	}

	v := types.MustUnmarshalValidator(k.cdc, value)
	validator = &v
	return validator, true
}

func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []*types.Validator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		validators = append(validators, &validator)
	}

	return validators
}

func (k Keeper) SetValidator(ctx sdk.Context, validator *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validatorKey := types.GetValidatorKey(validator.EthAddress)
	store.Set(validatorKey, types.MustMarshalValidator(k.cdc, validator))
}

func (k Keeper) SetValidatorStates(
	ctx sdk.Context, ethAddr string, status types.ValidatorStatus, tokens, shares string) error {

	val, found := k.GetValidator(ctx, ethAddr)
	if !found {
		return fmt.Errorf("validator %s not found", ethAddr)
	}
	tkInt, ok := sdk.NewIntFromString(tokens)
	if !ok {
		return fmt.Errorf("Invalid tokens %s", tokens)
	}
	shInt, ok := sdk.NewIntFromString(shares)
	if !ok {
		return fmt.Errorf("Invalid shares %s", shares)
	}
	val.Status = status
	val.Tokens = tokens
	val.Shares = shares
	k.SetValidator(ctx, val)

	sdkValAddr, err := types.SdkValAddrFromSgnBech32(val.SgnAddress)
	if err != nil {
		return err
	}
	sdkVal, found := k.sdkval.GetValidator(ctx, sdkValAddr)
	if !found {
		if val.Status == types.ValidatorStatus_Bonded {
			// TODO: create sdk validator
			return nil
		} else if val.Status == types.ValidatorStatus_Unbonded {
			log.Debugf("Validator %s %s not bonded", ethAddr, val.SgnAddress)
			return nil
		} else {
			log.Debugf("Validator %s %s %s not found", ethAddr, val.SgnAddress, val.Status)
			return nil
		}
	}

	k.sdkval.DeleteValidatorByPowerIndex(ctx, sdkVal)
	sdkVal.Status = sdk_staking.BondStatus(val.Status)
	sdkVal.Tokens = tkInt
	sdkVal.DelegatorShares = shInt.ToDec()
	k.sdkval.SetValidator(ctx, sdkVal)

	if val.Status == types.ValidatorStatus_Bonded {
		k.sdkval.SetNewValidatorByPowerIndex(ctx, sdkVal)
	} else if val.Status == types.ValidatorStatus_Unbonded {
		k.sdkval.RemoveValidator(ctx, sdkValAddr)
	} else if val.Status == types.ValidatorStatus_Unbonding {
	}

	return nil
}

// Get sdk validators
func (k Keeper) GetBondedSdkValidators(ctx sdk.Context) []sdk_staking.Validator {
	return k.sdkval.GetBondedValidatorsByPower(ctx)
}

// Get a sdk validator by consencus address
func (k Keeper) GetSdkValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) (sdk_staking.Validator, bool) {
	return k.sdkval.GetValidatorByConsAddr(ctx, addr)
}

// Get a sdk validator by validator account address
func (k Keeper) GetSdkValidator(ctx sdk.Context, addr sdk.ValAddress) (sdk_staking.Validator, bool) {
	return k.sdkval.GetValidator(ctx, addr)
}
