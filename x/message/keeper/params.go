package keeper

import (
	"time"

	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetTriggerSignCooldown returns the current message trigger sign cooldown period.
func (k Keeper) GetTriggerSignCooldown(ctx sdk.Context) (cooldown time.Duration) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyTriggerSignCooldown, &cooldown)
	return cooldown
}
