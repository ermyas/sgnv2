package keeper

import (
	"fmt"
	"math/big"

	clog "github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	paramstore    params.Subspace
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	farmingKeeper types.FarmingKeeper
	distrKeeper   types.DistributionKeeper

	feeCollectorName string // name of the FeeCollector ModuleAccount
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	params params.Subspace,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	farmingKeeper types.FarmingKeeper,
	distrKeeper types.DistributionKeeper,
	feeCollectorName string) Keeper {
	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		paramstore:       params,
		bankKeeper:       bankKeeper,
		stakingKeeper:    stakingKeeper,
		farmingKeeper:    farmingKeeper,
		distrKeeper:      distrKeeper,
		feeCollectorName: feeCollectorName,
	}
}

// SyncFarming attempts to sync liquidity with the stake in the farming module
func (k Keeper) SyncFarming(ctx sdk.Context, sym string, chid uint64, lpAddr eth.Addr, liquidityBigInt *big.Int) error {
	var farmingErr error
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
			amount := sdk.NewCoin(denom, liquidity.Sub(stake))
			// Mint stakes and send to lp address in farming module
			farmingErr = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
			if farmingErr != nil {
				clog.Errorf("Failed to mint stake, error %s, poolName %s, lpAddr %s, liquidity %s, stake %s",
					farmingErr, poolName, lpAddr, liquidity, stake)
			}
			derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(farmingtypes.ModuleName, lpAddr)
			farmingErr = k.bankKeeper.SendCoinsFromModuleToAccount(
				ctx, types.ModuleName, derivedAccAddress, sdk.NewCoins(amount),
			)
			if farmingErr != nil {
				clog.Errorf("Failed to send stake, error %s, poolName %s, lpAddr %s, liquidity %s, stake %s",
					farmingErr, poolName, lpAddr, liquidity, stake)
			}
			// Stake
			farmingErr = k.farmingKeeper.Stake(ctx, poolName, lpAddr, amount)
			if farmingErr != nil {
				clog.Errorf("Failed to stake, error %s, poolName %s, lpAddr %s, liquidity %s, stake %s",
					farmingErr, poolName, lpAddr, liquidity, stake)
			}
		} else if liquidity.LT(stake) {
			amount := sdk.NewCoin(denom, stake.Sub(liquidity))
			// Unstake
			farmingErr = k.farmingKeeper.Unstake(ctx, poolName, lpAddr, amount)
			if farmingErr != nil {
				clog.Errorf("Failed to unstake, error %s, poolName %s, lpAddr %s, liquidity %s, stake %s",
					farmingErr, poolName, lpAddr, liquidity, stake)
			}
			// Burn stakes
			derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(farmingtypes.ModuleName, lpAddr)
			farmingErr = k.bankKeeper.SendCoinsFromAccountToModule(
				ctx, derivedAccAddress, types.ModuleName, sdk.NewCoins(amount),
			)
			if farmingErr != nil {
				clog.Errorf("Failed to send stake back, error %s, poolName %s, lpAddr %s, liquidity %s, stake %s",
					farmingErr, poolName, lpAddr, liquidity, stake)
			}
			farmingErr = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
			if farmingErr != nil {
				clog.Errorf("Failed to burn stake, error %s, poolName %s, lpAddr %s, liquidity %s, stake %s",
					farmingErr, poolName, lpAddr, liquidity, stake)
			}
		}
	}
	return farmingErr
}

func derivePoolNameAndDenom(symbol string, chainId uint64) (poolName string, denom string) {
	denom = fmt.Sprintf("CB-%s/%d", symbol, chainId)
	poolName = fmt.Sprintf("%s-%s/%d", types.ModuleName, symbol, chainId)
	return poolName, denom
}
