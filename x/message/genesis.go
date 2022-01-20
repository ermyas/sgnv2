package message

import (
	"github.com/celer-network/sgn-v2/x/message/keeper"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the message module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	for _, bus := range genState.MessageBuses {
		k.SetMessageBus(ctx, bus)
	}
}

// ExportGenesis returns the message module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	buses := make([]types.MessageBusInfo, 0)
	k.IterateAllMessageBuses(ctx,
		func(bus types.MessageBusInfo) (stop bool) {
			buses = append(buses, bus)
			return false
		},
	)
	genesis.MessageBuses = buses
	return genesis
}
