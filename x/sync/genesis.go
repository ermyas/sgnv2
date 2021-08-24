package sync

import (
	"github.com/celer-network/sgn-v2/x/sync/keeper"
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) []abci.ValidatorUpdate {
	keeper.SetParams(ctx, data.Params)
	keeper.SetNextUpdateId(ctx, data.StartUpdateId)

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	startUpdateId := keeper.GetNextUpdateId(ctx)

	return &types.GenesisState{
		Params:        params,
		StartUpdateId: startUpdateId,
	}
}

func ValidateGenesis(data *types.GenesisState) error {
	return data.Params.Validate()
}
