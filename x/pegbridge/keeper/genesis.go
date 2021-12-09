package keeper

import (
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the pegbridge module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	k.SetParams(ctx, data.Params)
	k.SetPegConfig(ctx, data.Config)
}

// ExportGenesis returns the pegbridge module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(k.GetParams(ctx), k.GetConfig(ctx))
}

func (k Keeper) GetConfig(ctx sdk.Context) types.PegConfig {
	vaults := make([]commontypes.ContractInfo, 0)
	k.IterateAllOriginalTokenVaults(ctx,
		func(vault commontypes.ContractInfo) (stop bool) {
			vaults = append(vaults, vault)
			return false
		},
	)
	bridges := make([]commontypes.ContractInfo, 0)
	k.IterateAllPeggedTokenBridges(ctx,
		func(bridge commontypes.ContractInfo) (stop bool) {
			bridges = append(bridges, bridge)
			return false
		},
	)
	pairs := make([]types.OrigPeggedPair, 0)
	k.IterateAllOrigPeggedPairs(ctx,
		func(pair types.OrigPeggedPair) (stop bool) {
			pairs = append(pairs, pair)
			return false
		},
	)
	return types.PegConfig{
		PeggedTokenBridges:  bridges,
		OriginalTokenVaults: vaults,
		OrigPeggedPairs:     pairs,
	}
}
