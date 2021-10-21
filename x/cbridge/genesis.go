package cbridge

import (
	"github.com/celer-network/sgn-v2/x/cbridge/keeper"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	if err := genState.Validate(); err != nil {
		panic(err)
	}
	k.SetCbrConfig(ctx, genState.Config)
	// set initial price so base fee will work before new prices are set in x/cbridge kv
	k.SetCbrPrice(ctx, genState.Price)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
