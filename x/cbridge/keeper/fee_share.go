package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddFeeShare(ctx sdk.Context, amount sdk.Coin) error {
	// Mint coins to distribution module's fee collector account
	if err := k.bankKeeper.MintCoins(ctx, k.feeCollectorName, sdk.NewCoins(amount)); err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnFeeShare(ctx sdk.Context, delAddr eth.Addr, amount sdk.Coin) error {
	// Send coins from delegator address to module
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(distrtypes.ModuleName, delAddr)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, derivedAccAddress, k.feeCollectorName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}
	// Burn coins
	err = k.bankKeeper.BurnCoins(ctx, k.feeCollectorName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}
	return nil
}
