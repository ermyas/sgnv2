package keeper

import (
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Get the entire Syncer metadata
func (k Keeper) GetSyncer(ctx sdk.Context) *types.Syncer {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.SyncerKey)
	if value == nil {
		return &types.Syncer{}
	}
	syncer := types.MustUnmarshalSyncer(k.cdc, value)
	return &syncer
}

// Sets the entire Syncer metadata
func (k Keeper) SetSyncer(ctx sdk.Context, syncer *types.Syncer) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SyncerKey, types.MustMarshalSyncer(k.cdc, syncer))
}
