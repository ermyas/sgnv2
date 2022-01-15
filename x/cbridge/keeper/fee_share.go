package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintAndSendFeeShare(ctx sdk.Context, addr eth.Addr, amount sdk.Coin) error {
	if err := k.MintFeeShare(ctx, amount); err != nil {
		return err
	}
	// Send coins from module to address directly, bypassing distribution mechanism
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(distrtypes.ModuleName, addr)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, k.feeCollectorName, derivedAccAddress, sdk.NewCoins(amount)); err != nil {
		return err
	}
	return nil
}

func (k Keeper) MintFeeShare(ctx sdk.Context, amount sdk.Coin) error {
	// Mint coins to distribution module's fee collector account
	if err := k.bankKeeper.MintCoins(ctx, k.feeCollectorName, sdk.NewCoins(amount)); err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnFeeShare(ctx sdk.Context, addr eth.Addr, amount sdk.Coin) error {
	// Send coins from address to module
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(distrtypes.ModuleName, addr)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, derivedAccAddress, k.feeCollectorName, sdk.NewCoins(amount)); err != nil {
		return err
	}
	// Burn coins
	if err := k.bankKeeper.BurnCoins(ctx, k.feeCollectorName, sdk.NewCoins(amount)); err != nil {
		return err
	}
	return nil
}
