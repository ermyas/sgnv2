package keeper

import (
	"fmt"

	"github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the farm store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	authKeeper types.AccountKeeper
	bankKeeper types.BankKeeper
}

// NewKeeper creates a farm keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	authKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		authKeeper: authKeeper,
		bankKeeper: bankKeeper,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
