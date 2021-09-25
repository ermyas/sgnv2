package keeper

import (
	"time"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

func ParamKeyTable() sdk_params.KeyTable {
	return sdk_params.NewKeyTable().RegisterParamSet(&types.Params{})
}

func (k Keeper) GetSignerUpdateDuraion(ctx sdk.Context) time.Duration {
	var duration time.Duration
	k.paramstore.Get(ctx, types.KeySignerUpdateDuration, &duration)
	return duration
}

// set the SignerUpdateDuraion
func (k Keeper) SetSignerUpdateDuraion(ctx sdk.Context, duration time.Duration) {
	k.paramstore.Set(ctx, types.KeySignerUpdateDuration, &duration)
}
