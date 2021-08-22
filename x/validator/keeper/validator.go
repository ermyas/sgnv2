package keeper

import (
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Get the entire Validator metadata
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

// Get the list of all validators
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []*types.Validator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		validators = append(validators, &validator)
	}

	return validators
}

// Sets the Validator metadata
func (k Keeper) SetValidator(ctx sdk.Context, validator *types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validatorKey := types.GetValidatorKey(validator.EthAddress)
	store.Set(validatorKey, types.MustMarshalValidator(k.cdc, validator))
}

// Get sdk validators metadata
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
