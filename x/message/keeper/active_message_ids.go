package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetActiveMessageId(ctx sdk.Context, dstChainId uint64, target eth.Addr, messageId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetActiveMessageIdsKey(dstChainId, target, messageId), []byte{})
}

func (k Keeper) HasActiveMessageId(ctx sdk.Context, dstChainId uint64, target eth.Addr, messageId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetActiveMessageIdsKey(dstChainId, target, messageId))
}

func (k Keeper) DeleteActiveMessageId(ctx sdk.Context, dstChainId uint64, target eth.Addr, messageId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetActiveMessageIdsKey(dstChainId, target, messageId))
}

func (k Keeper) GetActiveMessageIdsByDstChainId(ctx sdk.Context, dstChainId uint64) (messageIds []eth.Hash, found bool) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetActiveMessageIdsPrefixByDstChainId(dstChainId))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		messageIds = append(messageIds, types.GetMessageIdFromActiveMessageIdsKey(iter.Key()))
	}
	return messageIds, len(messageIds) > 0
}

func (k Keeper) GetActiveMessageIdsByChainIdTarget(ctx sdk.Context, dstChainId uint64, target eth.Addr) (messageIds []eth.Hash, found bool) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetActiveMessageIdsPrefixByDstChainIdTarget(dstChainId, target))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		messageIds = append(messageIds, types.GetMessageIdFromActiveMessageIdsKey(iter.Key()))
	}
	return messageIds, len(messageIds) > 0
}

func (k Keeper) IterateAllActiveMessageIds(
	ctx sdk.Context, handler func(messageId eth.Hash) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ActiveMessageIdsPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		messageId := types.GetMessageIdFromActiveMessageIdsKey(iter.Key())
		if handler(messageId) {
			break
		}
	}
}
