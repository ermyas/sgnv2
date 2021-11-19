package cbridge

import (
	"github.com/celer-network/sgn-v2/x/cbridge/keeper"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}
	k.SetParams(ctx, genState.Params)
	if len(genState.Kvs) == 0 {
		// no kv, assume first start from manual genesis.json
		if err := genState.Config.Validate(); err != nil {
			panic(err)
		}
		k.SetCbrConfig(ctx, genState.Config)
		// set initial price so base fee will work before new prices are set in x/cbridge kv
		if genState.Price != nil {
			k.SetCbrPrice(ctx, genState.Price)
		}
	} else {
		// go over Kvs and set all into keeper
		k.ImportAllKV(ctx.KVStore(k.StoreKey()), genState.Kvs)
	}
}

// ExportGenesis returns the capability module's exported genesis.
// NOTE we DONOT export CbrConfig/CbrPrice, instead we export everything as KV
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	ret := new(types.GenesisState)
	ret.Params = k.GetParams(ctx)
	ret.Kvs = k.ExportAllKV(ctx.KVStore(k.StoreKey()))
	return ret
}
