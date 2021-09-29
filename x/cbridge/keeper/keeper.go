package keeper

import (
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	paramstore    params.Subspace
	stakingKeeper types.StakingKeeper
	farmingKeeper types.FarmingKeeper
	// this line is used by starport scaffolding # ibc/keeper/attribute
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	params params.Subspace,
	stakingKeeper types.StakingKeeper,
	farmingKeeper types.FarmingKeeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramstore:    params,
		stakingKeeper: stakingKeeper,
		farmingKeeper: farmingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) FarmStake(ctx sdk.Context, sym string, chid uint64, lpAddr eth.Addr, delta *big.Int) error {
	poolName, denom := derivePoolNameAndDenom(sym, chid)
	if k.farmingKeeper.HasFarmingPool(ctx, poolName) {
		return k.farmingKeeper.Stake(
			ctx, poolName, lpAddr, sdk.NewCoin(denom, sdk.NewIntFromBigInt(delta)), true)
	}
	return nil
}

func (k Keeper) FarmUnstake(ctx sdk.Context, sym string, chid uint64, lpAddr eth.Addr, delta *big.Int) error {
	poolName, denom := derivePoolNameAndDenom(sym, chid)
	if k.farmingKeeper.HasFarmingPool(ctx, poolName) {
		return k.farmingKeeper.Unstake(
			ctx, poolName, lpAddr, sdk.NewCoin(denom, sdk.NewIntFromBigInt(delta)), true)
	}
	return nil
}

func derivePoolNameAndDenom(symbol string, chainId uint64) (poolName string, denom string) {
	denom = fmt.Sprintf("CB-%s/%d", symbol, chainId)
	poolName = fmt.Sprintf("%s-CB-%s", types.ModuleName, denom)
	return poolName, denom
}
