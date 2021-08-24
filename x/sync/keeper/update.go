package keeper

import (
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k Keeper) ProposeUpdates(ctx sdk.Context, updates []*types.ProposeUpdate, ethBlock uint64, sender string) error {
	updateId := k.GetNextUpdateId(ctx)
	proposeTs := uint64(ctx.BlockHeader().Time.Unix())
	closingTs := proposeTs + k.VotingPeriod(ctx)

	for _, u := range updates {
		update := types.NewPendingUpdate(updateId, u.Type, u.Data, ethBlock, sender, proposeTs, closingTs)
		k.SetPendingUpdate(ctx, update)
		updateId++
	}

	k.SetNextUpdateId(ctx, updateId)
	return nil
}

func (k Keeper) VoteUpdates(ctx sdk.Context, votes []*types.VoteUpdate, sender string) {
	for _, v := range votes {
		update, ok := k.GetPendingUpdate(ctx, v.Id)
		if !ok {
			continue
		}
		update.Votes = append(update.Votes, &types.Vote{Voter: sender, Option: v.Option})
		k.SetPendingUpdate(ctx, update)
	}
}

func (k Keeper) GetPendingUpdate(ctx sdk.Context, updateId uint64) (update *types.PendingUpdate, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetPendingUpdateKey(update.Id))
	if value == nil {
		return update, false
	}
	u := types.MustUnmarshalPendingUpdate(k.cdc, value)
	update = &u
	return update, true
}

func (k Keeper) GetAllPendingUpdates(ctx sdk.Context) (udpates []*types.PendingUpdate) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PendingUpdateKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		update := types.MustUnmarshalPendingUpdate(k.cdc, iterator.Value())
		udpates = append(udpates, &update)
	}

	return udpates
}

func (k Keeper) SetPendingUpdate(ctx sdk.Context, update *types.PendingUpdate) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPendingUpdateKey(update.Id), types.MustMarshalPendingUpdate(k.cdc, update))
}

func (keeper Keeper) RemovePendingUpdate(ctx sdk.Context, updateId uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(types.GetPendingUpdateKey(updateId))
}

// GetNextUpdateId gets the highest update id
func (k Keeper) GetNextUpdateId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextUpdateIdKey)
	if bz == nil {
		return 0
	}
	return types.GetUpdateIdFromBytes(bz)
}

// SetNextUpdateId sets the new update ID to the store
func (k Keeper) SetNextUpdateId(ctx sdk.Context, updateId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextUpdateIdKey, types.GetUpdateIdToBytes(updateId))
}

func (k Keeper) GetBondedValidators(ctx sdk.Context) []sdk_staking.Validator {
	return k.valKeeper.GetBondedSdkValidators(ctx)
}
