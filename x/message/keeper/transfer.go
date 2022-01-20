package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetTransfer(ctx sdk.Context, messageId eth.Hash, transfer types.Transfer) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetTransferKey(messageId), k.cdc.MustMarshal(&transfer))
}

func (k Keeper) GetTransfer(ctx sdk.Context, messageId eth.Hash) (transfer types.Transfer, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTransferKey(messageId))
	if bz == nil {
		return transfer, false
	}
	k.cdc.MustUnmarshal(bz, &transfer)
	return transfer, true
}

func (k Keeper) HasTransfer(ctx sdk.Context, messageId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetTransferKey(messageId))
}

func (k Keeper) DeleteTransfer(ctx sdk.Context, messageId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetTransferKey(messageId))
}

func (k Keeper) IterateAllTransfers(
	ctx sdk.Context, handler func(transfer types.Transfer) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.TransferPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var transfer types.Transfer
		k.cdc.MustUnmarshal(iter.Value(), &transfer)
		if handler(transfer) {
			break
		}
	}
}
