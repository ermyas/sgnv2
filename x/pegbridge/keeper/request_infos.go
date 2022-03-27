package keeper

import (
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetDepositInfo(ctx sdk.Context, info types.DepositInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDepositInfoKey(eth.Bytes2Hash(info.DepositId)), k.cdc.MustMarshal(&info))
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

func (k Keeper) SetBurnInfo(ctx sdk.Context, info types.BurnInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBurnInfoKey(eth.Bytes2Hash(info.BurnId)), k.cdc.MustMarshal(&info))
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

func (k Keeper) SetTotalSupply(ctx sdk.Context, peggedChainId uint64, peggedAddress eth.Addr, amount *big.Int) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetTotalSupplyKey(peggedChainId, peggedAddress), amount.Bytes())
}

func (k Keeper) GetTotalSupply(ctx sdk.Context, peggedChainId uint64, peggedAddress eth.Addr) (amount *big.Int, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTotalSupplyKey(peggedChainId, peggedAddress))
	if bz == nil {
		return amount, false
	}
	amount = new(big.Int)
	return amount.SetBytes(bz), true
}

func (k Keeper) HasTotalSupply(ctx sdk.Context, peggedChainId uint64, peggedAddress eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetTotalSupplyKey(peggedChainId, peggedAddress))
}

func (k Keeper) DeleteTotalSupply(ctx sdk.Context, peggedChainId uint64, peggedAddress eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetTotalSupplyKey(peggedChainId, peggedAddress))
}

func (k Keeper) IterateAllTotalSupplies(
	ctx sdk.Context, handler func(amount *big.Int) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.TotalSupplyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		if handler(new(big.Int).SetBytes(iter.Value())) {
			break
		}
	}
}

func (k Keeper) SetDepositRefund(ctx sdk.Context, depositId eth.Hash, wdOnChain types.WithdrawOnChain) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRefundKey(depositId), k.cdc.MustMarshal(&wdOnChain))
}

func (k Keeper) GetDepositRefund(ctx sdk.Context, depositId eth.Hash) (wdOnChain types.WithdrawOnChain, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRefundKey(depositId))
	if bz == nil {
		return wdOnChain, false
	}
	k.cdc.MustUnmarshal(bz, &wdOnChain)
	return wdOnChain, true
}

func (k Keeper) SetBurnRefund(ctx sdk.Context, burnId eth.Hash, mintOnChain types.MintOnChain) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRefundKey(burnId), k.cdc.MustMarshal(&mintOnChain))
}

func (k Keeper) GetBurnRefund(ctx sdk.Context, burnId eth.Hash) (mintOnChain types.MintOnChain, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRefundKey(burnId))
	if bz == nil {
		return mintOnChain, false
	}
	k.cdc.MustUnmarshal(bz, &mintOnChain)
	return mintOnChain, true
}

func (k Keeper) HasRefund(ctx sdk.Context, refId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetRefundKey(refId))
}

func (k Keeper) DeleteRefund(ctx sdk.Context, refId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRefundKey(refId))
}

func (k Keeper) IterateAllRefunds(
	ctx sdk.Context, handler func(sthOnChain []byte) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.RefundPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		if handler(iter.Value()) {
			break
		}
	}
}

func (k Keeper) SetRefundClaimInfo(
	ctx sdk.Context, srcId eth.Hash /* deposit or burn */, refundId eth.Hash /* withdraw or burn  */) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRefundClaimInfoKey(srcId), refundId.Bytes())
}

func (k Keeper) GetRefundClaimInfo(
	ctx sdk.Context, srcId eth.Hash /* deposit or burn */) (refundId eth.Hash /* withdraw or burn */, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRefundClaimInfoKey(srcId))
	if bz == nil {
		return refundId, false
	}
	refundId = eth.Bytes2Hash(bz)
	return refundId, true
}

func (k Keeper) HasRefundClaimInfo(ctx sdk.Context, srcId eth.Hash /* deposit or burn */) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetRefundClaimInfoKey(srcId))
}

func (k Keeper) DeleteRefundClaimInfo(ctx sdk.Context, srcId eth.Hash /* deposit or burn */) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRefundClaimInfoKey(srcId))
}

func (k Keeper) IterateAllRefundClaimInfos(
	ctx sdk.Context, handler func(refundId eth.Hash /* withdraw or burn */) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.RefundClaimInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		refundId := eth.Bytes2Hash(iter.Value())
		if handler(refundId) {
			break
		}
	}
}
