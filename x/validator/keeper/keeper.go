package keeper

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_auth_keeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk_staking_keeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc        codec.Codec  // The wire codec for binary encoding/decoding.
	storeKey   sdk.StoreKey // Unexposed key to access store from sdk.Context
	sdkacct    sdk_auth_keeper.AccountKeeperI
	sdkval     sdk_staking_keeper.Keeper
	paramstore sdk_params.Subspace
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey,
	sdkacct sdk_auth_keeper.AccountKeeperI, sdkval sdk_staking_keeper.Keeper, paramstore sdk_params.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		sdkacct:    sdkacct,
		sdkval:     sdkval,
		paramstore: paramstore.WithKeyTable(ParamKeyTable()),
	}
}

// Get validators metadata
func (k Keeper) GetBondedSgnValidators(ctx sdk.Context) []sdk_staking.Validator {
	return k.sdkval.GetBondedValidatorsByPower(ctx)
}

// Get a validator by consencus address
func (k Keeper) GetSgnValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) (sdk_staking.Validator, bool) {
	return k.sdkval.GetValidatorByConsAddr(ctx, addr)
}

// Get a validator by validator account address
func (k Keeper) GetSgnValidator(ctx sdk.Context, addr sdk.ValAddress) (sdk_staking.Validator, bool) {
	return k.sdkval.GetValidator(ctx, addr)
}

// func (k Keeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(index int64, validator exported.ValidatorI) (stop bool)) {
// 	k.sdkval.IterateBondedValidatorsByPower(ctx, fn)
// }

// Get the entire Syncer metadata
func (k Keeper) GetSyncer(ctx sdk.Context) *types.Syncer {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(types.SyncerKey) {
		return &types.Syncer{}
	}

	//value := store.Get(SyncerKey)
	var syncer types.Syncer
	//k.cdc.MustUnmarshalBinaryBare(value, &syncer)
	return &syncer
}

// Sets the entire Syncer metadata
func (k Keeper) SetSyncer(ctx sdk.Context, syncer *types.Syncer) {
	//store := ctx.KVStore(k.storeKey)
	//store.Set(SyncerKey, k.cdc.MustMarshalBinaryBare(syncer))
}

// Get the entire Delegator metadata for a validatorAddr and delegatorAddr
func (k Keeper) GetDelegator(ctx sdk.Context, validatorAddr, delegatorAddr string) (delegator *types.Delegator, found bool) {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(types.GetDelegatorKey(validatorAddr, delegatorAddr)) {
		return delegator, false
	}

	//value := store.Get(GetDelegatorKey(validatorAddr, delegatorAddr))
	//k.cdc.MustUnmarshalBinaryBare(value, &delegator)
	return delegator, true
}

// Get the list of all delegators
func (k Keeper) GetAllDelegators(ctx sdk.Context, validatorAddr string) (delegators []*types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetDelegatorsKey(validatorAddr))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var delegator types.Delegator
		//k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &delegator)
		delegators = append(delegators, &delegator)
	}
	return delegators
}

// Sets the entire Delegator metadata for a validatorAddr and delegatorAddr
func (k Keeper) SetDelegator(ctx sdk.Context, delegator *types.Delegator) {
	//store := ctx.KVStore(k.storeKey)
	//store.Set(GetDelegatorKey(delegator.ValidatorAddr, delegator.DelegatorAddr), k.cdc.MustMarshalBinaryBare(delegator))
}

func (k Keeper) RemoveDelegator(ctx sdk.Context, delegator *types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDelegatorKey(delegator.ValAddress, delegator.EthAddress))
}

// Get the entire Validator metadata
func (k Keeper) GetValidator(ctx sdk.Context, validatorAddr string) (validator *types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	validatorKey := types.GetValidatorKey(validatorAddr)

	if !store.Has(validatorKey) {
		return validator, false
	}

	//value := store.Get(validatorKey)
	//k.cdc.MustUnmarshalBinaryBare(value, &validator)
	return validator, true
}

// Get the list of all validators
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []*types.Validator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		//k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &validator)
		validators = append(validators, &validator)
	}
	return validators
}

// Sets the Validator metadata
func (k Keeper) SetValidator(ctx sdk.Context, validator *types.Validator) {
	//store := ctx.KVStore(k.storeKey)
	//validatorKey := GetValidatorKey(validator.EthAddress)
	//store.Set(validatorKey, k.cdc.MustMarshalBinaryBare(validator))
}

func (k Keeper) InitAccount(ctx sdk.Context, accAddress sdk.AccAddress) {
	err := sdk.VerifyAddressFormat(accAddress)
	if err != nil {
		log.Errorf("InitAccount %s err: %s", accAddress, err)
		return
	}
	account := k.sdkacct.GetAccount(ctx, accAddress)
	if account == nil {
		log.Infof("Set new account %s", accAddress)
		//account = k.sdkacct.NewAccountWithAddress(ctx, accAddress)
		k.sdkacct.SetAccount(ctx, account)
	}
}

func (k Keeper) RemoveAccount(ctx sdk.Context, accAddress sdk.AccAddress) {
	account := k.sdkacct.GetAccount(ctx, accAddress)
	if account != nil {
		log.Infof("Remove account %s", accAddress)
		k.sdkacct.RemoveAccount(ctx, account)
	}
}
