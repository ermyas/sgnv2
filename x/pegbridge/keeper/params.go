package keeper

import (
	"time"

	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of pegbridge parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the pegbridge parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetTriggerSignCooldown returns the current pegbridge trigger sign cooldown period.
func (k Keeper) GetTriggerSignCooldown(ctx sdk.Context) (cooldown time.Duration) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyTriggerSignCooldown, &cooldown)
	return cooldown
}
