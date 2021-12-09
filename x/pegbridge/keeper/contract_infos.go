package keeper

import (
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetOriginalTokenVault(ctx sdk.Context, vault commontypes.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetOriginalTokenVaultKey(vault.ChainId), k.cdc.MustMarshal(&vault))
}

func (k Keeper) GetOriginalTokenVault(ctx sdk.Context, chainId uint64) (vault commontypes.ContractInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetOriginalTokenVaultKey(chainId))
	if bz == nil {
		return vault, false
	}
	k.cdc.MustUnmarshal(bz, &vault)
	return vault, true
}

func (k Keeper) HasOriginalTokenVault(ctx sdk.Context, chainId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetOriginalTokenVaultKey(chainId))
}

func (k Keeper) DeleteOriginalTokenVault(ctx sdk.Context, chainId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetOriginalTokenVaultKey(chainId))
}

func (k Keeper) IterateAllOriginalTokenVaults(
	ctx sdk.Context, handler func(vault commontypes.ContractInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.OriginalTokenVaultPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var vault commontypes.ContractInfo
		k.cdc.MustUnmarshal(iter.Value(), &vault)
		if handler(vault) {
			break
		}
	}
}

func (k Keeper) SetPeggedTokenBridge(ctx sdk.Context, bridge commontypes.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPeggedTokenBridgeKey(bridge.ChainId), k.cdc.MustMarshal(&bridge))
}

func (k Keeper) GetPeggedTokenBridge(ctx sdk.Context, chainId uint64) (bridge commontypes.ContractInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPeggedTokenBridgeKey(chainId))
	if bz == nil {
		return bridge, false
	}
	k.cdc.MustUnmarshal(bz, &bridge)
	return bridge, true
}

func (k Keeper) HasPeggedTokenBridge(ctx sdk.Context, chainId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetPeggedTokenBridgeKey(chainId))
}

func (k Keeper) DeletePeggedTokenBridge(ctx sdk.Context, chainId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPeggedTokenBridgeKey(chainId))
}

func (k Keeper) IterateAllPeggedTokenBridges(
	ctx sdk.Context, handler func(bridge commontypes.ContractInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PeggedTokenBridgePrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var bridge commontypes.ContractInfo
		k.cdc.MustUnmarshal(iter.Value(), &bridge)
		if handler(bridge) {
			break
		}
	}
}
