package keeper

import (
	"fmt"
	"math/big"

	clog "github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	paramstore    params.Subspace
	stakingKeeper types.StakingKeeper
	farmingKeeper types.FarmingKeeper
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

// SyncFarming attempts to sync liquidity with the stake in the farming module
func (k Keeper) SyncFarming(ctx sdk.Context, sym string, chid uint64, lpAddr eth.Addr, liquidityBigInt *big.Int) error {
	var err error
	poolName, denom := derivePoolNameAndDenom(sym, chid)
	// Only sync if the pool exists
	if k.farmingKeeper.HasFarmingPool(ctx, poolName) {
		liquidity := sdk.NewIntFromBigInt(liquidityBigInt)
		stake := sdk.ZeroInt()
		stakeInfo, found := k.farmingKeeper.GetStakeInfo(ctx, lpAddr, poolName)
		if found {
			stake = stakeInfo.Amount.Amount.RoundInt()
		}
		if liquidity.GT(stake) {
			err = k.farmingKeeper.Stake(
				ctx, poolName, lpAddr, sdk.NewCoin(denom, liquidity.Sub(stake)), true)
			if err != nil {
				clog.Errorf("Failed to stake, poolName %s, lpAddr %s, liquidity %s, stake %s", poolName, lpAddr, liquidity, stake)
			}
		} else if liquidity.LT(stake) {
			err = k.farmingKeeper.Unstake(
				ctx, poolName, lpAddr, sdk.NewCoin(denom, stake.Sub(liquidity)), true)
			if err != nil {
				clog.Errorf("Failed to unstake, poolName %s, lpAddr %s, liquidity %s, stake %s", poolName, lpAddr, liquidity, stake)
			}
		}
	}
	return err
}

func derivePoolNameAndDenom(symbol string, chainId uint64) (poolName string, denom string) {
	denom = fmt.Sprintf("CB-%s/%d", symbol, chainId)
	poolName = fmt.Sprintf("%s-%s/%d", types.ModuleName, symbol, chainId)
	return poolName, denom
}
