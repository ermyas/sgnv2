package keeper

import (
	"github.com/celer-network/sgn-v2/x/cbridge/types"
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

// MaxValidatorDiff - max validator add
func (k Keeper) MultiChainAsset(ctx sdk.Context) (res types.MultiChainAsset) {
	k.paramstore.Get(ctx, types.KeyMultiChainAsset, &res)
	return
}

func (k Keeper) PowerReduction(ctx sdk.Context) sdk.Int {
	return sdk.DefaultPowerReduction
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.MultiChainAsset(ctx))
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
