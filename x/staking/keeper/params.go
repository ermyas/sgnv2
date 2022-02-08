package keeper

import (
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

// SyncerDuration - syncer duration
func (k Keeper) SyncerDuration(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeySyncerDuration, &res)
	return
}

func (k Keeper) SyncerCandidates(ctx sdk.Context) (res []string) {
	k.paramstore.GetIfExists(ctx, types.KeySyncerCandidates, &res)
	return
}

func (k Keeper) PowerReduction(ctx sdk.Context) sdk.Int {
	return types.DefaultPowerReduction
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.SyncerDuration(ctx),
		k.SyncerCandidates(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
