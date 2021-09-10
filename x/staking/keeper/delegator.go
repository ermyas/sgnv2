package keeper

import (
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetDelegator(ctx sdk.Context, valAddr, delAddr string) (delegator *types.Delegator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetDelegatorKey(valAddr, delAddr))
	if value == nil {
		return delegator, false
	}
	d := types.MustUnmarshalDelegator(k.cdc, value)
	delegator = &d
	return delegator, true
}

func (k Keeper) GetAllDelegators(ctx sdk.Context, valAddr string) (delegators []types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetDelegatorsKey(valAddr))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegator := types.MustUnmarshalDelegator(k.cdc, iterator.Value())
		delegators = append(delegators, delegator)
	}
	return delegators
}

func (k Keeper) SetDelegator(ctx sdk.Context, delegator *types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	delegatorKey := types.GetDelegatorKey(delegator.ValAddress, delegator.DelAddress)
	store.Set(delegatorKey, types.MustMarshalDelegator(k.cdc, delegator))
}

func (k Keeper) RemoveDelegator(ctx sdk.Context, delegator *types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDelegatorKey(delegator.ValAddress, delegator.DelAddress))
}

func (k Keeper) SetDelegatorShares(ctx sdk.Context, valAddr, delAddr string, shares sdk.Int) {
	delegator := types.NewDelegator(valAddr, delAddr, shares)
	if shares.IsZero() {
		k.RemoveDelegator(ctx, delegator)
	} else {
		k.SetDelegator(ctx, delegator)
	}
}
