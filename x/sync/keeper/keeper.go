package keeper

import (
	cbrkeeper "github.com/celer-network/sgn-v2/x/cbridge/keeper"
	msgkeeper "github.com/celer-network/sgn-v2/x/message/keeper"
	pegkeeper "github.com/celer-network/sgn-v2/x/pegbridge/keeper"
	"github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc           codec.BinaryCodec // The wire codec for binary encoding/decoding.
	storeKey      sdk.StoreKey      // Unexposed key to access store from sdk.Context
	paramstore    paramstypes.Subspace
	stakingKeeper types.StakingKeeper
	cbrKeeper     cbrkeeper.Keeper
	pegbrKeeper   pegkeeper.Keeper
	msgbrKeeper   msgkeeper.Keeper
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	stakingKeeper types.StakingKeeper,
	paramstore paramstypes.Subspace,
	cbr cbrkeeper.Keeper,
	pegbr pegkeeper.Keeper,
	msg msgkeeper.Keeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		stakingKeeper: stakingKeeper,
		paramstore:    paramstore,
		cbrKeeper:     cbr,
		pegbrKeeper:   pegbr,
		msgbrKeeper:   msg,
	}
}
