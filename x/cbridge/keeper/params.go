package keeper

import (
	"time"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	return types.NewParams(k.GetSignerUpdateDuration(ctx), k.GetSignAgainCoolDownDuration(ctx))
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetSignerUpdateDuration(ctx sdk.Context) (duration time.Duration) {
	k.paramstore.GetIfExists(ctx, types.KeySignerUpdateDuration, &duration)
	if duration == 0 {
		duration = types.DefaultSignerUpdateDuration
	}
	return
}

func (k Keeper) GetSignAgainCoolDownDuration(ctx sdk.Context) (duration time.Duration) {
	k.paramstore.GetIfExists(ctx, types.KeySignAgainCoolDownDuration, &duration)
	if duration == 0 {
		duration = types.DefaultSignAgainCoolDownDuration
	}
	return
}
