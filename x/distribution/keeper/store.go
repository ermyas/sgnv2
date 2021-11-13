package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// get the delegator withdraw address, defaulting to the delegator address
func (k Keeper) GetDelegatorWithdrawAddr(ctx sdk.Context, delAddr eth.Addr) eth.Addr {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetDelegatorWithdrawAddrKey(delAddr))
	if b == nil {
		return delAddr
	}
	return eth.Bytes2Addr(b)
}

// set the delegator withdraw address
func (k Keeper) SetDelegatorWithdrawAddr(ctx sdk.Context, delAddr eth.Addr, withdrawAddr eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDelegatorWithdrawAddrKey(delAddr), withdrawAddr.Bytes())
}

// delete a delegator withdraw addr
func (k Keeper) DeleteDelegatorWithdrawAddr(ctx sdk.Context, delAddr eth.Addr, withdrawAddr eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDelegatorWithdrawAddrKey(delAddr))
}

// iterate over delegator withdraw addrs
func (k Keeper) IterateDelegatorWithdrawAddrs(ctx sdk.Context, handler func(del eth.Addr, addr eth.Addr) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.DelegatorWithdrawAddrPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr := eth.Bytes2Addr(iter.Value())
		del := types.GetDelegatorWithdrawInfoAddress(iter.Key())
		if handler(del, addr) {
			break
		}
	}
}

// get the global fee pool distribution info
func (k Keeper) GetFeePool(ctx sdk.Context) (feePool types.FeePool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.FeePoolKey)
	if b == nil {
		panic("Stored fee pool should not have been nil")
	}
	k.cdc.MustUnmarshal(b, &feePool)
	return
}

// set the global fee pool distribution info
func (k Keeper) SetFeePool(ctx sdk.Context, feePool types.FeePool) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&feePool)
	store.Set(types.FeePoolKey, b)
}

// GetPreviousProposerConsAddr returns the proposer consensus address for the
// current block.
func (k Keeper) GetPreviousProposerConsAddr(ctx sdk.Context) sdk.ConsAddress {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ProposerKey)
	if bz == nil {
		panic("previous proposer not set")
	}

	addrValue := gogotypes.BytesValue{}
	k.cdc.MustUnmarshal(bz, &addrValue)
	return addrValue.GetValue()
}

// set the proposer public key for this block
func (k Keeper) SetPreviousProposerConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.BytesValue{Value: consAddr})
	store.Set(types.ProposerKey, bz)
}

// get the starting info associated with a delegator
func (k Keeper) GetDelegatorStartingInfo(ctx sdk.Context, val eth.Addr, del eth.Addr) (period types.DelegatorStartingInfo) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetDelegatorStartingInfoKey(val, del))
	k.cdc.MustUnmarshal(b, &period)
	return
}

// set the starting info associated with a delegator
func (k Keeper) SetDelegatorStartingInfo(ctx sdk.Context, val eth.Addr, del eth.Addr, period types.DelegatorStartingInfo) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&period)
	store.Set(types.GetDelegatorStartingInfoKey(val, del), b)
}

// check existence of the starting info associated with a delegator
func (k Keeper) HasDelegatorStartingInfo(ctx sdk.Context, val eth.Addr, del eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetDelegatorStartingInfoKey(val, del))
}

// delete the starting info associated with a delegator
func (k Keeper) DeleteDelegatorStartingInfo(ctx sdk.Context, val eth.Addr, del eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetDelegatorStartingInfoKey(val, del))
}

// iterate over delegator starting infos
func (k Keeper) IterateDelegatorStartingInfos(ctx sdk.Context, handler func(val eth.Addr, del eth.Addr, info types.DelegatorStartingInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.DelegatorStartingInfoPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.DelegatorStartingInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		val, del := types.GetDelegatorStartingInfoAddresses(iter.Key())
		if handler(val, del, info) {
			break
		}
	}
}

// get historical rewards for a particular period
func (k Keeper) GetValidatorHistoricalRewards(ctx sdk.Context, val eth.Addr, period uint64) (rewards types.ValidatorHistoricalRewards) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetValidatorHistoricalRewardsKey(val, period))
	k.cdc.MustUnmarshal(b, &rewards)
	return
}

// set historical rewards for a particular period
func (k Keeper) SetValidatorHistoricalRewards(ctx sdk.Context, val eth.Addr, period uint64, rewards types.ValidatorHistoricalRewards) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&rewards)
	store.Set(types.GetValidatorHistoricalRewardsKey(val, period), b)
}

// iterate over historical rewards
func (k Keeper) IterateValidatorHistoricalRewards(ctx sdk.Context, handler func(val eth.Addr, period uint64, rewards types.ValidatorHistoricalRewards) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorHistoricalRewardsPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var rewards types.ValidatorHistoricalRewards
		k.cdc.MustUnmarshal(iter.Value(), &rewards)
		addr, period := types.GetValidatorHistoricalRewardsAddressPeriod(iter.Key())
		if handler(addr, period, rewards) {
			break
		}
	}
}

// delete a historical reward
func (k Keeper) DeleteValidatorHistoricalReward(ctx sdk.Context, val eth.Addr, period uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorHistoricalRewardsKey(val, period))
}

// delete historical rewards for a validator
func (k Keeper) DeleteValidatorHistoricalRewards(ctx sdk.Context, val eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetValidatorHistoricalRewardsPrefix(val))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// delete all historical rewards
func (k Keeper) DeleteAllValidatorHistoricalRewards(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorHistoricalRewardsPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// historical reference count (used for testcases)
func (k Keeper) GetValidatorHistoricalReferenceCount(ctx sdk.Context) (count uint64) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorHistoricalRewardsPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var rewards types.ValidatorHistoricalRewards
		k.cdc.MustUnmarshal(iter.Value(), &rewards)
		count += uint64(rewards.ReferenceCount)
	}
	return
}

// get current rewards for a validator
func (k Keeper) GetValidatorCurrentRewards(ctx sdk.Context, val eth.Addr) (rewards types.ValidatorCurrentRewards) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetValidatorCurrentRewardsKey(val))
	k.cdc.MustUnmarshal(b, &rewards)
	return
}

// set current rewards for a validator
func (k Keeper) SetValidatorCurrentRewards(ctx sdk.Context, val eth.Addr, rewards types.ValidatorCurrentRewards) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&rewards)
	store.Set(types.GetValidatorCurrentRewardsKey(val), b)
}

// delete current rewards for a validator
func (k Keeper) DeleteValidatorCurrentRewards(ctx sdk.Context, val eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorCurrentRewardsKey(val))
}

// iterate over current rewards
func (k Keeper) IterateValidatorCurrentRewards(ctx sdk.Context, handler func(val eth.Addr, rewards types.ValidatorCurrentRewards) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorCurrentRewardsPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var rewards types.ValidatorCurrentRewards
		k.cdc.MustUnmarshal(iter.Value(), &rewards)
		addr := types.GetValidatorCurrentRewardsAddress(iter.Key())
		if handler(addr, rewards) {
			break
		}
	}
}

// get accumulated commission for a validator
func (k Keeper) GetValidatorAccumulatedCommission(ctx sdk.Context, val eth.Addr) (commission types.ValidatorAccumulatedCommission) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetValidatorAccumulatedCommissionKey(val))
	if b == nil {
		return types.ValidatorAccumulatedCommission{}
	}
	k.cdc.MustUnmarshal(b, &commission)
	return
}

// set accumulated commission for a validator
func (k Keeper) SetValidatorAccumulatedCommission(ctx sdk.Context, val eth.Addr, commission types.ValidatorAccumulatedCommission) {
	var bz []byte

	store := ctx.KVStore(k.storeKey)
	if commission.Commission.IsZero() {
		bz = k.cdc.MustMarshal(&types.ValidatorAccumulatedCommission{})
	} else {
		bz = k.cdc.MustMarshal(&commission)
	}

	store.Set(types.GetValidatorAccumulatedCommissionKey(val), bz)
}

// delete accumulated commission for a validator
func (k Keeper) DeleteValidatorAccumulatedCommission(ctx sdk.Context, val eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorAccumulatedCommissionKey(val))
}

// iterate over accumulated commissions
func (k Keeper) IterateValidatorAccumulatedCommissions(ctx sdk.Context, handler func(val eth.Addr, commission types.ValidatorAccumulatedCommission) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorAccumulatedCommissionPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var commission types.ValidatorAccumulatedCommission
		k.cdc.MustUnmarshal(iter.Value(), &commission)
		addr := types.GetValidatorAccumulatedCommissionAddress(iter.Key())
		if handler(addr, commission) {
			break
		}
	}
}

// get validator outstanding rewards
func (k Keeper) GetValidatorOutstandingRewards(ctx sdk.Context, val eth.Addr) (rewards types.ValidatorOutstandingRewards) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorOutstandingRewardsKey(val))
	k.cdc.MustUnmarshal(bz, &rewards)
	return
}

// set validator outstanding rewards
func (k Keeper) SetValidatorOutstandingRewards(ctx sdk.Context, val eth.Addr, rewards types.ValidatorOutstandingRewards) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&rewards)
	store.Set(types.GetValidatorOutstandingRewardsKey(val), b)
}

// delete validator outstanding rewards
func (k Keeper) DeleteValidatorOutstandingRewards(ctx sdk.Context, val eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorOutstandingRewardsKey(val))
}

// iterate validator outstanding rewards
func (k Keeper) IterateValidatorOutstandingRewards(ctx sdk.Context, handler func(val eth.Addr, rewards types.ValidatorOutstandingRewards) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorOutstandingRewardsPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		rewards := types.ValidatorOutstandingRewards{}
		k.cdc.MustUnmarshal(iter.Value(), &rewards)
		addr := types.GetValidatorOutstandingRewardsAddress(iter.Key())
		if handler(addr, rewards) {
			break
		}
	}
}
