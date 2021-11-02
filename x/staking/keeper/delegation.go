package keeper

import (
	"bytes"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetDelegation(
	ctx sdk.Context,
	delAddr eth.Addr,
	valAddr eth.Addr,
) (delegation types.Delegation, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetDelegationKey(delAddr, valAddr))
	if value == nil {
		return delegation, false
	}
	delegation = types.MustUnmarshalDelegation(k.cdc, value)
	return delegation, true
}

// IterateAllDelegations iterate through all of the delegations
func (k Keeper) IterateAllDelegations(ctx sdk.Context, cb func(delegation types.Delegation) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		if cb(delegation) {
			break
		}
	}
}

// GetAllDelegations returns all delegations. NOTE: This is only used during genesis dump
func (k Keeper) GetAllDelegations(ctx sdk.Context, valAddr eth.Addr) (delegations types.Delegations) {
	k.IterateAllDelegations(ctx, func(delegation types.Delegation) bool {
		delegations = append(delegations, delegation)
		return false
	})

	return delegations
}

// return all delegations to a specific validator. Useful for querier.
func (k Keeper) GetValidatorDelegations(ctx sdk.Context, valAddr eth.Addr) (delegations []types.Delegation) { //nolint:interfacer
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.DelegationKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		if bytes.Equal(delegation.GetValidatorAddr().Bytes(), valAddr.Bytes()) {
			delegations = append(delegations, delegation)
		}
	}

	return delegations
}

// return a given amount of all the delegations from a delegator
func (k Keeper) GetDelegatorDelegations(ctx sdk.Context, delegator eth.Addr,
	maxRetrieve uint16) (delegations []types.Delegation) {
	delegations = make([]types.Delegation, maxRetrieve)
	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetDelegationsKey(delegator)

	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		delegations[i] = delegation
		i++
	}

	return delegations[:i] // trim if the array length < maxRetrieve
}

func (k Keeper) SetDelegation(ctx sdk.Context, delegation types.Delegation) {
	store := ctx.KVStore(k.storeKey)
	delegatorKey := types.GetDelegationKey(
		eth.Hex2Addr(delegation.DelegatorAddress),
		eth.Hex2Addr(delegation.ValidatorAddress))
	store.Set(delegatorKey, types.MustMarshalDelegation(k.cdc, delegation))
}

func (k Keeper) RemoveDelegation(ctx sdk.Context, delegation types.Delegation) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDelegationKey(
		eth.Hex2Addr(delegation.DelegatorAddress),
		eth.Hex2Addr(delegation.ValidatorAddress)))
}

func (k Keeper) SetDelegationShares(
	ctx sdk.Context,
	delAddr eth.Addr,
	valAddr eth.Addr,
	shares sdk.Int,
) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetDelegationKey(delAddr, valAddr))
	if k.Validator(ctx, valAddr) == nil {
		// The delegation happened before we have seen the validator. Add a placeholder for distribution.
		k.SetValidator(ctx, &types.Validator{
			EthAddress:      eth.Addr2Hex(valAddr),
			DelegatorShares: shares,
		})
	}
	// call the appropriate hook if present
	if value == nil {
		// New delegation
		k.BeforeDelegationCreated(ctx, delAddr, valAddr)
	} else {
		k.BeforeDelegationModified(ctx, delAddr, valAddr)
	}

	delegation := types.NewDelegation(delAddr, valAddr, shares)
	if shares.IsZero() {
		k.RemoveDelegation(ctx, delegation)
	} else {
		k.SetDelegation(ctx, delegation)
		// Call the after-modification hook
		k.AfterDelegationModified(ctx, delAddr, valAddr)
	}
}
