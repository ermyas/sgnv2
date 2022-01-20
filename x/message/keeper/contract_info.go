package keeper

import (
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetMessageBus(ctx sdk.Context, bus types.MessageBusInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetMessageBusKey(bus.ContractInfo.ChainId), k.cdc.MustMarshal(&bus))
}

func (k Keeper) GetMessageBus(ctx sdk.Context, chainId uint64) (bus types.MessageBusInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMessageBusKey(chainId))
	if bz == nil {
		return bus, false
	}
	k.cdc.MustUnmarshal(bz, &bus)
	return bus, true
}

func (k Keeper) HasMessageBus(ctx sdk.Context, chainId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetMessageBusKey(chainId))
}

func (k Keeper) DeleteMessageBus(ctx sdk.Context, chainId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMessageBusKey(chainId))
}

func (k Keeper) IterateAllMessageBuses(
	ctx sdk.Context, handler func(info types.MessageBusInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MessageBusPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var bus types.MessageBusInfo
		k.cdc.MustUnmarshal(iter.Value(), &bus)
		if handler(bus) {
			break
		}
	}
}
