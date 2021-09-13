package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramstore params.Subspace
	// this line is used by starport scaffolding # ibc/keeper/attribute
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	params params.Subspace,
	// this line is used by starport scaffolding # ibc/keeper/parameter

) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: params,
		// this line is used by starport scaffolding # ibc/keeper/return

	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
