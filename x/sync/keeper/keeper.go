package keeper

import (
	valkeeper "github.com/celer-network/sgn-v2/x/validator/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc        codec.BinaryCodec // The wire codec for binary encoding/decoding.
	storeKey   sdk.StoreKey      // Unexposed key to access store from sdk.Context
	paramstore sdk_params.Subspace
	bankKeeper sdk_bank.Keeper
	valKeeper  valkeeper.Keeper
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	valKeeper valkeeper.Keeper,
	paramstore sdk_params.Subspace,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		valKeeper:  valKeeper,
		paramstore: paramstore.WithKeyTable(ParamKeyTable()),
	}
}
