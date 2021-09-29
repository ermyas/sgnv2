package keeper

import (
	"fmt"
	"math/big"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingkeeper "github.com/celer-network/sgn-v2/x/farming/keeper"
	stakingkeeper "github.com/celer-network/sgn-v2/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	paramstore    params.Subspace
	stakingKeeper stakingkeeper.Keeper
	farmingKeeper farmingkeeper.Keeper
	// this line is used by starport scaffolding # ibc/keeper/attribute
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	params params.Subspace,
	stakingKeeper stakingkeeper.Keeper,
	farmingKeeper farmingkeeper.Keeper) Keeper {
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
	return k.farmingKeeper.Stake(ctx, fmt.Sprintf("%s-%d", sym, chid), lpAddr, sdk.NewCoin(sym, sdk.NewIntFromBigInt(delta)), true)
}

func (k Keeper) FarmUnStake(ctx sdk.Context, sym string, chid uint64, lpAddr eth.Addr, delta *big.Int) error {
	return k.farmingKeeper.Unstake(ctx, fmt.Sprintf("%s-%d", sym, chid), lpAddr, sdk.NewCoin(sym, sdk.NewIntFromBigInt(delta)), true)
}
