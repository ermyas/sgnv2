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
	cdc        codec.BinaryCodec // The wire codec for binary encoding/decoding.
	storeKey   sdk.StoreKey      // Unexposed key to access store from sdk.Context
	sdkacct    sdk_auth.AccountKeeperI
	sdkval     sdk_staking.Keeper
	paramstore sdk_params.Subspace
}

// NewKeeper creates new instances of the validator Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	sdkacct sdk_auth.AccountKeeperI,
	sdkval sdk_staking.Keeper,
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

func (k Keeper) InitAccount(ctx sdk.Context, accAddress sdk.AccAddress) error {
	err := sdk.VerifyAddressFormat(accAddress)
	if err != nil {
		return fmt.Errorf("InitAccount %s err: %s", accAddress, err)
	}
	account := k.sdkacct.GetAccount(ctx, accAddress)
	if account == nil {
		log.Infof("Set sdk account %s", accAddress)
		account = k.sdkacct.NewAccountWithAddress(ctx, accAddress)
		k.sdkacct.SetAccount(ctx, account)
	}
	return err
}

func (k Keeper) RemoveAccount(ctx sdk.Context, accAddress sdk.AccAddress) {
	account := k.sdkacct.GetAccount(ctx, accAddress)
	if account != nil {
		log.Infof("Remove sdk account %s", accAddress)
		k.sdkacct.RemoveAccount(ctx, account)
	}
}
