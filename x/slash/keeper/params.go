package keeper

import (
	"github.com/celer-network/sgn-v2/x/slash/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

// EnableSlash - if enable slash event
func (k Keeper) EnableSlash(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyEnableSlash, &res)
	return
}

// SignedBlocksWindow - sliding window for downtime slashing
func (k Keeper) SignedBlocksWindow(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeySignedBlocksWindow, &res)
	return
}

// SlashTimeout - how long does slash stay valid
func (k Keeper) SlashTimeout(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeySlashTimeout, &res)
	return
}

// SlashFactorDoubleSign - slash factor in case of double sign
func (k Keeper) SlashFactorDoubleSign(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeySlashFactorDoubleSign, &res)
	return
}

// SlashFactorDowntime - slash factor for downtime
func (k Keeper) SlashFactorDowntime(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeySlashFactorDowntime, &res)
	return
}

// SlashFactorDowntime - slash factor for downtime
func (k Keeper) JailPeriod(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyJailPeriod, &res)
	return
}

// MinSignedPerWindow - minimum blocks signed per window
func (k Keeper) MinSignedPerWindow(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyMinSignedPerWindow, &res)
	return
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.EnableSlash(ctx),
		k.SignedBlocksWindow(ctx),
		k.SlashTimeout(ctx),
		k.SlashFactorDoubleSign(ctx),
		k.SlashFactorDowntime(ctx),
		k.JailPeriod(ctx),
		k.MinSignedPerWindow(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
