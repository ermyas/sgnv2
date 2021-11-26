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
