package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	paramSpace    paramstypes.Subspace
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	cbridgeKeeper types.CbridgeKeeper
	distrKeeper   types.DistributionKeeper

	feeCollectorName string // name of the FeeCollector ModuleAccount
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	cbridgeKeeper types.CbridgeKeeper,
	distrKeeper types.DistributionKeeper,
	feeCollectorName string,
) Keeper {

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		paramSpace:       paramSpace,
		bankKeeper:       bankKeeper,
		stakingKeeper:    stakingKeeper,
		cbridgeKeeper:    cbridgeKeeper,
		distrKeeper:      distrKeeper,
		feeCollectorName: feeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
