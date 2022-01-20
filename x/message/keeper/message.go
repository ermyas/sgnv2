package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetMessage(ctx sdk.Context, messageId eth.Hash, message types.Message) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetMessageKey(messageId), k.cdc.MustMarshal(&message))
}

func (k Keeper) GetMessage(ctx sdk.Context, messageId eth.Hash) (message types.Message, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMessageKey(messageId))
	if bz == nil {
		return message, false
	}
	k.cdc.MustUnmarshal(bz, &message)
	return message, true
}

func (k Keeper) HasMessage(ctx sdk.Context, messageId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetMessageKey(messageId))
}

func (k Keeper) DeleteMessage(ctx sdk.Context, messageId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMessageKey(messageId))
}

func (k Keeper) IterateAllMessages(
	ctx sdk.Context, handler func(message types.Message) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MessagePrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var message types.Message
		k.cdc.MustUnmarshal(iter.Value(), &message)
		if handler(message) {
			break
		}
	}
}
