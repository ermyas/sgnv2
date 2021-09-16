package slash

import (
	"github.com/celer-network/sgn-v2/x/slash/keeper"
	"github.com/celer-network/sgn-v2/x/slash/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	keeper.SetParams(ctx, data.Params)
	keeper.SetSlashNonce(ctx, data.SlashNonce)

	for _, slash := range data.Slashes {
		keeper.SetSlash(ctx, slash)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) types.GenesisState {
	params := keeper.GetParams(ctx)
	slashNonce := keeper.GetSlashNonce(ctx)
	slashes := keeper.GetSlashes(ctx)

	return types.GenesisState{
		Params:     params,
		SlashNonce: slashNonce,
		Slashes:    slashes,
	}
}
