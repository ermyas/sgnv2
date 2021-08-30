package keeper

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	cdc              codec.BinaryCodec // The wire codec for binary encoding/decoding.
	storeKey         sdk.StoreKey      // Unexposed key to access store from sdk.Context
	sdkAccountKeeper sdk_auth.AccountKeeperI
	sdkStakingKeeper sdk_staking.Keeper
	paramstore       sdk_params.Subspace
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	sdkAccountKeeper sdk_auth.AccountKeeperI,
	sdkStakingKeeper sdk_staking.Keeper,
	paramstore sdk_params.Subspace,
) Keeper {
	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		sdkAccountKeeper: sdkAccountKeeper,
		sdkStakingKeeper: sdkStakingKeeper,
		paramstore:       paramstore.WithKeyTable(ParamKeyTable()),
	}
}

func (k Keeper) InitAccount(ctx sdk.Context, accAddress sdk.AccAddress) error {
	err := sdk.VerifyAddressFormat(accAddress)
	if err != nil {
		return fmt.Errorf("InitAccount %s err: %s", accAddress, err)
	}
	account := k.sdkAccountKeeper.GetAccount(ctx, accAddress)
	if account == nil {
		log.Infof("Set sdk account %s", accAddress)
		account = k.sdkAccountKeeper.NewAccountWithAddress(ctx, accAddress)
		k.sdkAccountKeeper.SetAccount(ctx, account)
	}
	return err
}

func (k Keeper) RemoveAccount(ctx sdk.Context, accAddress sdk.AccAddress) {
	account := k.sdkAccountKeeper.GetAccount(ctx, accAddress)
	if account != nil {
		log.Infof("Remove sdk account %s", accAddress)
		k.sdkAccountKeeper.RemoveAccount(ctx, account)
	}
}
