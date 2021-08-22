package keeper

import (
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Get the entire Delegator metadata for a validator and delegator addresses
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

// Get the list of all delegators
func (k Keeper) GetAllDelegators(ctx sdk.Context, valAddr string) (delegators []*types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetDelegatorsKey(valAddr))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		delegator := types.MustUnmarshalDelegator(k.cdc, iterator.Value())
		delegators = append(delegators, &delegator)
	}
	return delegators
}

// Sets the entire Delegator metadata for a validatorAddr and delegatorAddr
func (k Keeper) SetDelegator(ctx sdk.Context, delegator *types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	delegatorKey := types.GetDelegatorKey(delegator.ValAddress, delegator.EthAddress)
	store.Set(delegatorKey, types.MustMarshalDelegator(k.cdc, delegator))
}

func (k Keeper) RemoveDelegator(ctx sdk.Context, delegator *types.Delegator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDelegatorKey(delegator.ValAddress, delegator.EthAddress))
}
