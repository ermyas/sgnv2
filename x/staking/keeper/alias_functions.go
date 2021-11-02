package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validator Set

// iterate through the validator set and perform the provided function
func (k Keeper) IterateValidators(ctx sdk.Context, fn func(index int64, validator types.ValidatorI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorKey)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		stop := fn(i, validator) // XXX is this safe will the validator unexposed fields be able to get written to?

		if stop {
			break
		}
		i++
	}
}

// Validator gets the Validator interface for a particular address
func (k Keeper) Validator(ctx sdk.Context, address eth.Addr) types.ValidatorI {
	val, found := k.GetValidator(ctx, address)
	if !found {
		return nil
	}

	return val
}

// ValidatorByConsAddr gets the validator interface for a particular pubkey
func (k Keeper) ValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) types.ValidatorI {
	val, found := k.GetValidatorByConsAddr(ctx, addr)
	if !found {
		return nil
	}

	return val
}

// Delegation Set

// Returns self as it is both a validatorset and delegationset
func (k Keeper) GetValidatorSet() types.ValidatorSet {
	return k
}

// Delegation get the delegation interface for a particular set of delegator and validator addresses
func (k Keeper) Delegation(ctx sdk.Context, addrDel eth.Addr, addrVal eth.Addr) types.DelegationI {
	bond, ok := k.GetDelegation(ctx, addrDel, addrVal)
	if !ok {
		return nil
	}

	return bond
}

// iterate through all of the delegations from a delegator
func (k Keeper) IterateDelegations(ctx sdk.Context, delAddr eth.Addr,
	fn func(index int64, del types.DelegationI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetDelegationsKey(delAddr)

	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey) // smallest to largest
	defer iterator.Close()

	for i := int64(0); iterator.Valid(); iterator.Next() {
		del := types.MustUnmarshalDelegation(k.cdc, iterator.Value())

		stop := fn(i, del)
		if stop {
			break
		}
		i++
	}
}

// return all delegations used during genesis dump
// TODO: remove this func, change all usage for iterate functionality (follow upstream)
func (k Keeper) GetAllSDKDelegations(ctx sdk.Context) (delegations []types.Delegation) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		delegations = append(delegations, delegation)
	}

	return
}
