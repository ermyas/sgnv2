package keeper

import (
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

func (k Keeper) GetAllDelegations(ctx sdk.Context, valAddr eth.Addr) (delegations types.Delegations) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetDelegationsKey(valAddr))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		delegations = append(delegations, delegation)
	}
	return delegations
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
	}

	// Call the after-modification hook
	k.AfterDelegationModified(ctx, delAddr, valAddr)
}
