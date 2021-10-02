package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetFarmingPool(ctx sdk.Context, pool types.FarmingPool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetFarmingPoolKey(pool.Name), k.cdc.MustMarshal(&pool))
}

func (k Keeper) GetFarmingPool(ctx sdk.Context, poolName string) (pool types.FarmingPool, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFarmingPoolKey(poolName))
	if bz == nil {
		return pool, false
	}
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}

func (k Keeper) HasFarmingPool(ctx sdk.Context, poolName string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFarmingPoolKey(poolName))
}

func (k Keeper) DeleteFarmingPool(ctx sdk.Context, poolName string) {
	store := ctx.KVStore(k.storeKey)
	// delete pool key
	store.Delete(types.GetFarmingPoolKey(poolName))
}

// GetFarmingPoolNamesForAccount gets all pool names that an account has staked coins in from the store
func (k Keeper) GetFarmingPoolNamesForAccount(ctx sdk.Context, addr eth.Addr) (poolNames types.PoolNameList) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.AddressToPoolPrefix, addr.Bytes()...))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		poolNames = append(poolNames, types.SplitPoolNameFromStakeInfoKey(iterator.Key()))
	}

	return
}

// GetAccountsStakedIn gets all addresses of accounts that have staked coins in a pool
func (k Keeper) GetAccountsStakedIn(ctx sdk.Context, poolName string) (stakerAddrList types.AddrList) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, append(types.PoolToAddressPrefix, []byte(poolName)...))
	defer iterator.Close()

	splitIndex := 1 + len(poolName)
	for ; iterator.Valid(); iterator.Next() {
		stakerAddrList = append(stakerAddrList, eth.Bytes2Addr(iterator.Key()[splitIndex:]))
	}

	return
}

// getNumPools gets the number of pools that already exist
func (k Keeper) GetNumPools(ctx sdk.Context) types.NumPools {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.FarmingPoolPrefix)
	defer iterator.Close()
	var num uint64
	for ; iterator.Valid(); iterator.Next() {
		num++
	}

	return types.NewNumPools(num)
}

// GetFarmingPools gets all pools that exist currently in the store
func (k Keeper) GetFarmingPools(ctx sdk.Context) (pools types.FarmingPools) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.FarmingPoolPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var pool types.FarmingPool
		k.cdc.MustUnmarshal(iterator.Value(), &pool)
		pools = append(pools, pool)
	}

	return
}

func (k Keeper) SetAddressInFarmingPool(ctx sdk.Context, poolName string, addr eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetAddressInFarmingPoolKey(poolName, addr), []byte(""))
}

// HasAddressInFarmingPool check existence of the pool associated with a address
func (k Keeper) HasAddressInFarmingPool(ctx sdk.Context, poolName string, addr eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetAddressInFarmingPoolKey(poolName, addr))
}

func (k Keeper) DeleteAddressInFarmingPool(ctx sdk.Context, poolName string, addr eth.Addr) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAddressInFarmingPoolKey(poolName, addr))
}

func (k Keeper) SetStakeInfo(ctx sdk.Context, stakeInfo types.StakeInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetStakeInfoKey(eth.Hex2Addr(stakeInfo.StakerAddress), stakeInfo.PoolName), k.cdc.MustMarshal(&stakeInfo))
}

func (k Keeper) GetStakeInfo(ctx sdk.Context, addr eth.Addr, poolName string) (info types.StakeInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetStakeInfoKey(addr, poolName))
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

// HasStakeInfo check existence of the address associated with a pool
func (k Keeper) HasStakeInfo(ctx sdk.Context, addr eth.Addr, poolName string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetStakeInfoKey(addr, poolName))
}

func (k Keeper) DeleteStakeInfo(ctx sdk.Context, addr eth.Addr, poolName string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetStakeInfoKey(addr, poolName))
}

// Iterate over all stake infos
func (k Keeper) IterateAllStakeInfos(
	ctx sdk.Context, handler func(stakeInfo types.StakeInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.AddressToPoolPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var stakeInfo types.StakeInfo
		k.cdc.MustUnmarshal(iter.Value(), &stakeInfo)
		if handler(stakeInfo) {
			break
		}
	}
}
