package keeper

import (
	"encoding/binary"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
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

// ----------- versioned store ---------------

func (k Keeper) SetVersionedVault(ctx sdk.Context, version uint32, vault commontypes.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetVersionedVaultKey(vault.ChainId, version), eth.Hex2Addr(vault.Address).Bytes())
	bz := k.cdc.MustMarshal(&gogotypes.UInt32Value{Value: version})
	store.Set(types.GetVaultVersionKey(vault.ChainId, eth.Hex2Addr(vault.Address)), bz)
}

func (k Keeper) GetVersionedVault(ctx sdk.Context, chainId uint64, version uint32) (addr eth.Addr, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVersionedVaultKey(chainId, version))
	if bz == nil {
		return eth.ZeroAddr, false
	}
	return eth.Bytes2Addr(bz), true
}

func (k Keeper) GetOriginalVault(ctx sdk.Context, chainId uint64, version uint32) (addr eth.Addr, found bool) {
	addr, found = k.GetVersionedVault(ctx, chainId, version)
	if found {
		return
	}
	vault, found := k.GetOriginalTokenVault(ctx, chainId)
	return eth.Hex2Addr(vault.Address), found
}

func (k Keeper) GetVaultVersion(ctx sdk.Context, chainId uint64, addr eth.Addr) (version uint32, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVaultVersionKey(chainId, addr))
	if bz == nil {
		return 0, false
	}
	v := gogotypes.UInt32Value{}
	k.cdc.MustUnmarshal(bz, &v)
	return v.Value, true
}

func (k Keeper) DeleteVersionedVault(ctx sdk.Context, chainId uint64, version uint32) {
	addr, found := k.GetVersionedVault(ctx, chainId, version)
	if !found {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetVersionedVaultKey(chainId, version))
	store.Delete(types.GetVaultVersionKey(chainId, addr))
}

func (k Keeper) GetAllVersionedVaults(ctx sdk.Context) (vaults []types.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.VersionedVaultPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		chainId := binary.LittleEndian.Uint64(iterator.Key()[1:9])
		version := binary.LittleEndian.Uint32(iterator.Key()[9:])
		addrress := eth.Bytes2AddrHex(iterator.Value())
		vaults = append(vaults,
			types.ContractInfo{
				Contract: commontypes.NewContractInfo(chainId, addrress),
				Version:  version,
			},
		)
	}
	return
}

func (k Keeper) SetVersionedBridge(ctx sdk.Context, version uint32, vault commontypes.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetVersionedBridgeKey(vault.ChainId, version), eth.Hex2Addr(vault.Address).Bytes())
	bz := k.cdc.MustMarshal(&gogotypes.UInt32Value{Value: version})
	store.Set(types.GetBridgeVersionKey(vault.ChainId, eth.Hex2Addr(vault.Address)), bz)
}

func (k Keeper) GetVersionedBridge(ctx sdk.Context, chainId uint64, version uint32) (addr eth.Addr, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVersionedBridgeKey(chainId, version))
	if bz == nil {
		return eth.ZeroAddr, false
	}
	return eth.Bytes2Addr(bz), true
}

func (k Keeper) GetPeggedBridge(ctx sdk.Context, chainId uint64, version uint32) (addr eth.Addr, found bool) {
	addr, found = k.GetVersionedBridge(ctx, chainId, version)
	if found {
		return
	}
	bridge, found := k.GetPeggedTokenBridge(ctx, chainId)
	return eth.Hex2Addr(bridge.Address), found
}

func (k Keeper) GetBridgeVersion(ctx sdk.Context, chainId uint64, addr eth.Addr) (version uint32, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetBridgeVersionKey(chainId, addr))
	if bz == nil {
		return 0, false
	}
	v := gogotypes.UInt32Value{}
	k.cdc.MustUnmarshal(bz, &v)
	return v.Value, true
}

func (k Keeper) DeleteVersionedBridge(ctx sdk.Context, chainId uint64, version uint32) {
	addr, found := k.GetVersionedBridge(ctx, chainId, version)
	if !found {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetVersionedBridgeKey(chainId, version))
	store.Delete(types.GetBridgeVersionKey(chainId, addr))
}

func (k Keeper) GetAllVersionedBridges(ctx sdk.Context) (bridges []types.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.VersionedBridgePrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		chainId := binary.LittleEndian.Uint64(iterator.Key()[1:9])
		version := binary.LittleEndian.Uint32(iterator.Key()[9:])
		addrress := eth.Bytes2AddrHex(iterator.Value())
		bridges = append(bridges,
			types.ContractInfo{
				Contract: commontypes.NewContractInfo(chainId, addrress),
				Version:  version,
			},
		)
	}
	return
}
