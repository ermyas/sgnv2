package keeper

import (
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRewardContract(ctx sdk.Context, contract commontypes.ContractInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRewardContractKey(contract.ChainId), k.cdc.MustMarshal(&contract))
}

func (k Keeper) GetRewardContract(ctx sdk.Context, chainId uint64) (contract commontypes.ContractInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRewardContractKey(chainId))
	if bz == nil {
		return contract, false
	}
	k.cdc.MustUnmarshal(bz, &contract)
	return contract, true
}

func (k Keeper) IterateAllRewardContracts(
	ctx sdk.Context, handler func(info commontypes.ContractInfo) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.RewardContractPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info commontypes.ContractInfo
		k.cdc.MustUnmarshal(iter.Value(), &info)
		if handler(info) {
			break
		}
	}
}
