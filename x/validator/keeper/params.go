package keeper

import (
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
func (k Keeper) SyncerDuration(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeySyncerDuration, &res)
	return
}

// EpochLength - epoch length
func (k Keeper) EpochLength(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyEpochLength, &res)
	return
}

// MaxValidatorDiff - max validator add
func (k Keeper) MaxValidatorDiff(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxValidatorDiff, &res)
	return
}

func (k Keeper) PowerReduction(ctx sdk.Context) sdk.Int {
	return sdk.DefaultPowerReduction
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.SyncerDuration(ctx),
		k.EpochLength(ctx),
		k.MaxValidatorDiff(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
