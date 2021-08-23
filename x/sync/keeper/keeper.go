package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_auth_keeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk_staking_keeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc        codec.BinaryCodec // The wire codec for binary encoding/decoding.
	storeKey   sdk.StoreKey      // Unexposed key to access store from sdk.Context
	sdkacct    sdk_auth_keeper.AccountKeeperI
	sdkval     sdk_staking_keeper.Keeper
	paramstore sdk_params.Subspace
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	sdkacct sdk_auth_keeper.AccountKeeperI,
	sdkval sdk_staking_keeper.Keeper,
	paramstore sdk_params.Subspace,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		sdkacct:    sdkacct,
		sdkval:     sdkval,
		paramstore: paramstore.WithKeyTable(ParamKeyTable()),
	}
}
