package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetDepositInfo(ctx sdk.Context, depositId eth.Hash, info types.DepositInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDepositInfoKey(depositId), k.cdc.MustMarshal(&info))
}

func (k Keeper) GetDepositInfo(ctx sdk.Context, depositId eth.Hash) (info types.DepositInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetDepositInfoKey(depositId))
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

func (k Keeper) HasDepositInfo(ctx sdk.Context, depositId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetDepositInfoKey(depositId))
}

func (k Keeper) DeleteDepositInfo(ctx sdk.Context, depositId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDepositInfoKey(depositId))
}

func (k Keeper) IterateAllDepositInfos(
	ctx sdk.Context, handler func(info types.DepositInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.DepositInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.DepositInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		if handler(info) {
			break
		}
	}
}

func (k Keeper) SetWithdrawInfo(ctx sdk.Context, withdrawId eth.Hash, info types.WithdrawInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetWithdrawInfoKey(withdrawId), k.cdc.MustMarshal(&info))
}

func (k Keeper) GetWithdrawInfo(ctx sdk.Context, withdrawId eth.Hash) (info types.WithdrawInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetWithdrawInfoKey(withdrawId))
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

func (k Keeper) HasWithdrawInfo(ctx sdk.Context, withdrawId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetWithdrawInfoKey(withdrawId))
}

func (k Keeper) DeleteWithdrawInfo(ctx sdk.Context, withdrawId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetWithdrawInfoKey(withdrawId))
}

func (k Keeper) IterateAllWithdrawInfos(
	ctx sdk.Context, handler func(info types.WithdrawInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.WithdrawInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.WithdrawInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		if handler(info) {
			break
		}
	}
}

func (k Keeper) SetMintInfo(ctx sdk.Context, mintId eth.Hash, info types.MintInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetMintInfoKey(mintId), k.cdc.MustMarshal(&info))
}

func (k Keeper) GetMintInfo(ctx sdk.Context, mintId eth.Hash) (info types.MintInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMintInfoKey(mintId))
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

func (k Keeper) HasMintInfo(ctx sdk.Context, mintId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetMintInfoKey(mintId))
}

func (k Keeper) DeleteMintInfo(ctx sdk.Context, mintId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMintInfoKey(mintId))
}

func (k Keeper) IterateAllMintInfos(
	ctx sdk.Context, handler func(info types.MintInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MintInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.MintInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		if handler(info) {
			break
		}
	}
}

func (k Keeper) SetBurnInfo(ctx sdk.Context, burnId eth.Hash, info types.BurnInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBurnInfoKey(burnId), k.cdc.MustMarshal(&info))
}

func (k Keeper) GetBurnInfo(ctx sdk.Context, burnId eth.Hash) (info types.BurnInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetBurnInfoKey(burnId))
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

func (k Keeper) HasBurnInfo(ctx sdk.Context, burnId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetBurnInfoKey(burnId))
}

func (k Keeper) DeleteBurnInfo(ctx sdk.Context, burnId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetBurnInfoKey(burnId))
}

func (k Keeper) IterateAllBurnInfos(
	ctx sdk.Context, handler func(info types.BurnInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.BurnInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.BurnInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		if handler(info) {
			break
		}
	}
}

func (k Keeper) SetFeeClaimInfo(ctx sdk.Context, address eth.Addr, nonce uint64, info types.FeeClaimInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetFeeClaimInfoKey(address, nonce), k.cdc.MustMarshal(&info))
}

func (k Keeper) GetFeeClaimInfo(ctx sdk.Context, address eth.Addr, nonce uint64) (info types.FeeClaimInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFeeClaimInfoKey(address, nonce))
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

func (k Keeper) HasFeeClaimInfo(ctx sdk.Context, address eth.Addr, nonce uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeClaimInfoKey(address, nonce))
}

func (k Keeper) DeleteFeeClaimInfo(ctx sdk.Context, address eth.Addr, nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetFeeClaimInfoKey(address, nonce))
}

func (k Keeper) IterateAllFeeClaimInfos(
	ctx sdk.Context, handler func(info types.FeeClaimInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.FeeClaimInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.FeeClaimInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		if handler(info) {
			break
		}
	}
}
