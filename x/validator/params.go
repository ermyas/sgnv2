package validator

import (
	"time"

	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

// ParamTable for validator module
func ParamKeyTable() sdk_params.KeyTable {
	return sdk_params.NewKeyTable().RegisterParamSet(&types.Params{})
}

// SyncerDuration - syncer duration
func (k Keeper) SyncerDuration(ctx sdk.Context) (res uint) {
	k.paramstore.Get(ctx, types.KeySyncerDuration, &res)
	return
}

// EpochLength - epoch length
func (k Keeper) EpochLength(ctx sdk.Context) (res uint) {
	k.paramstore.Get(ctx, types.KeyEpochLength, &res)
	return
}

// MaxValidatorDiff - max validator add
func (k Keeper) MaxValidatorDiff(ctx sdk.Context) (res uint) {
	k.paramstore.Get(ctx, types.KeyMaxValidatorDiff, &res)
	return
}

// ClaimWindow - withdraw window
func (k Keeper) ClaimWindow(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyClaimWindow, &res)
	return
}

// MiningReward - mining reward
func (k Keeper) MiningReward(ctx sdk.Context) (res sdk.Int) {
	k.paramstore.Get(ctx, types.KeyMiningReward, &res)
	return
}

// PullerReward - puller reward
func (k Keeper) PullerReward(ctx sdk.Context) (res sdk.Int) {
	k.paramstore.Get(ctx, types.KeyPullerReward, &res)
	return
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.SyncerDuration(ctx),
		k.EpochLength(ctx),
		k.MaxValidatorDiff(ctx),
		k.ClaimWindow(ctx),
		k.MiningReward(ctx),
		k.PullerReward(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
