package keeper

import (
	"github.com/celer-network/sgn-v2/x/sync/types"
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

func (k Keeper) VotingPeriod(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyVotingPeriod, &res)
	return
}

func (k Keeper) TallyThreshold(ctx sdk.Context) (res float32) {
	k.paramstore.Get(ctx, types.KeyTallyThreshold, &res)
	return
}

func (k Keeper) GetParams(ctx sdk.Context) *types.Params {
	return types.NewParams(
		k.VotingPeriod(ctx),
		k.TallyThreshold(ctx),
	)
}

func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	k.paramstore.SetParamSet(ctx, params)
}
