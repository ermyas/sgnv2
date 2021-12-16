package keeper

import (
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetOrigPeggedPair(ctx sdk.Context, pair types.OrigPeggedPair) {
	if pair.Obsolete {
		k.DeleteOrigPeggedPair(ctx, pair.Orig.ChainId, eth.Hex2Addr(pair.Orig.Address), pair.Pegged.ChainId, eth.Hex2Addr(pair.Pegged.Address))
		return
	}

	storedPair, found := k.GetOrigPeggedPair(ctx, pair.Orig.ChainId, eth.Hex2Addr(pair.Orig.Address), pair.Pegged.ChainId)
	if found && eth.Hex2Addr(storedPair.Pegged.Address) != eth.Hex2Addr(pair.Pegged.Address) {
		k.DeletePeggedOrigIndex(ctx, storedPair.Pegged.ChainId, eth.Hex2Addr(storedPair.Pegged.Address))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetOrigPeggedPairKey(pair.Orig.ChainId, eth.Hex2Addr(pair.Orig.Address), pair.Pegged.ChainId),
		k.cdc.MustMarshal(&pair))

	index := types.PeggedOrigIndex{
		Pegged: commontypes.ContractInfo{
			ChainId: pair.Pegged.ChainId,
			Address: pair.Pegged.Address,
		},
		Orig: commontypes.ContractInfo{
			ChainId: pair.Orig.ChainId,
			Address: pair.Orig.Address,
		},
	}
	k.SetPeggedOrigIndex(ctx, index)
}

func (k Keeper) GetOrigPeggedPair(
	ctx sdk.Context, origChainId uint64, origAddress eth.Addr, peggedChainId uint64) (pair types.OrigPeggedPair, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetOrigPeggedPairKey(origChainId, origAddress, peggedChainId))
	if bz == nil {
		return pair, false
	}
	k.cdc.MustUnmarshal(bz, &pair)
	return pair, true
}

func (k Keeper) HasOrigPeggedPair(
	ctx sdk.Context, origChainId uint64, origAddress eth.Addr, peggedChainId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetOrigPeggedPairKey(origChainId, origAddress, peggedChainId))
}

func (k Keeper) DeleteOrigPeggedPair(
	ctx sdk.Context, origChainId uint64, origAddress eth.Addr, peggedChainId uint64, peggedAddress eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetOrigPeggedPairKey(origChainId, origAddress, peggedChainId))
	k.DeletePeggedOrigIndex(ctx, peggedChainId, peggedAddress)
}

func (k Keeper) IterateAllOrigPeggedPairs(
	ctx sdk.Context, handler func(pair types.OrigPeggedPair) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.OrigPeggedPairPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var pair types.OrigPeggedPair
		k.cdc.MustUnmarshal(iter.Value(), &pair)
		if handler(pair) {
			break
		}
	}
}

func (k Keeper) IterateOrigPeggedPairsByOrig(
	ctx sdk.Context, origChainId uint64, origAddress eth.Addr, handler func(pair types.OrigPeggedPair) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	prefix := types.GetOrigPeggedByOrigPrefix(origChainId, origAddress)
	iter := sdk.KVStorePrefixIterator(store, prefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var pair types.OrigPeggedPair
		k.cdc.MustUnmarshal(iter.Value(), &pair)
		if handler(pair) {
			break
		}
	}
}

func (k Keeper) SetPeggedOrigIndex(ctx sdk.Context, index types.PeggedOrigIndex) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetPeggedOrigIndexKey(index.Pegged.ChainId, eth.Hex2Addr(index.Pegged.Address)),
		k.cdc.MustMarshal(&index),
	)
}

func (k Keeper) GetPeggedOrigIndex(
	ctx sdk.Context, peggedChainId uint64, peggedAddress eth.Addr) (index types.PeggedOrigIndex, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPeggedOrigIndexKey(peggedChainId, peggedAddress))
	if bz == nil {
		return index, false
	}
	k.cdc.MustUnmarshal(bz, &index)
	return index, true
}

func (k Keeper) HasPeggedOrigIndex(ctx sdk.Context,
	peggedChainId uint64, peggedAddress eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetPeggedOrigIndexKey(peggedChainId, peggedAddress))
}

func (k Keeper) DeletePeggedOrigIndex(ctx sdk.Context,
	peggedChainId uint64, peggedAddress eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPeggedOrigIndexKey(peggedChainId, peggedAddress))
}

func (k Keeper) IterateAllPeggedOrigIndices(
	ctx sdk.Context, handler func(index types.PeggedOrigIndex) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PeggedOrigIndexPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var index types.PeggedOrigIndex
		k.cdc.MustUnmarshal(iter.Value(), &index)
		if handler(index) {
			break
		}
	}
}

func (k Keeper) GetOrigPeggedPairByPegged(
	ctx sdk.Context, peggedChainId uint64, peggedAddress eth.Addr) (pair types.OrigPeggedPair, found bool) {
	index, found := k.GetPeggedOrigIndex(ctx, peggedChainId, peggedAddress)
	if !found {
		return pair, false
	}
	origInfo := index.Orig
	return k.GetOrigPeggedPair(ctx, origInfo.ChainId, eth.Hex2Addr(origInfo.Address), peggedChainId)
}
