package keeper

import (
	"time"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of farming parameters.
func (k Keeper) GetParams(clientCtx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(clientCtx, &params)
	return params
}

// SetParams sets the farming parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetClaimCooldown returns the current farming claim cooldown period.
func (k Keeper) GetClaimCooldown(ctx sdk.Context) (cooldown time.Duration) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyClaimCooldown, &cooldown)
	return cooldown
}
