package keeper

import (
	"context"
	"fmt"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
	stakingkeeper "github.com/celer-network/sgn-v2/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	memKey        sdk.StoreKey
	paramstore    params.Subspace
	stakingKeeper stakingkeeper.Keeper
	// this line is used by starport scaffolding # ibc/keeper/attribute
}

func (k Keeper) ChainTokensConfig(ctx context.Context, request *types.ChainTokensConfigRequest) (*types.ChainTokensConfigResponse, error) {
	panic("implement me")
}

func (k Keeper) GetFee(ctx context.Context, request *types.GetFeeRequest) (*types.GetFeeResponse, error) {
	panic("implement me")
}

func (k Keeper) QueryTransferStatus(ctx context.Context, request *types.QueryTransferStatusRequest) (*types.QueryTransferStatusResponse, error) {
	panic("implement me")
}

func (k Keeper) LiquidityDetailList(ctx context.Context, request *types.LiquidityDetailListRequest) (*types.LiquidityDetailListResponse, error) {
	panic("implement me")
}

func (k Keeper) QueryAddLiquidityStatus(ctx context.Context, request *types.QueryAddLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	panic("implement me")
}

func (k Keeper) QueryWithdrawLiquidityStatus(ctx context.Context, request *types.QueryWithdrawLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	panic("implement me")
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	params params.Subspace,
	stakingKeeper stakingkeeper.Keeper,
	// this line is used by starport scaffolding # ibc/keeper/parameter

) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    params,
		stakingKeeper: stakingKeeper,
		// this line is used by starport scaffolding # ibc/keeper/return

	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
