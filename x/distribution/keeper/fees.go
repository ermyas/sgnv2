package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetWithdrawableBalance(ctx sdk.Context, delAddr eth.Addr, coin sdk.Coin) sdk.Coin {
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, delAddr)
	return k.bankKeeper.GetBalance(ctx, derivedAccAddress, coin.Denom)
}

func (k Keeper) ClaimCBridgeFeeShare(ctx sdk.Context, delAddr eth.Addr) error {
	// 1. Withdraw reward for all validators. If delAddr is a validator address, withdraw its commission.
	err := k.withdrawRewardsAndCommission(ctx, delAddr)
	if err != nil {
		return err
	}

	// 2. Emit claim_cbridge_fee_share event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaimCBridgeFeeShare,
		sdk.NewAttribute(types.AttributeKeyDelegatorAddress, delAddr.String()),
	))

	return nil
}

func (k Keeper) ClaimPegBridgeFees(ctx sdk.Context, delAddr eth.Addr) error {
	// 1. Withdraw reward for all validators. If delAddr is a validator address, withdraw its commission.
	err := k.withdrawRewardsAndCommission(ctx, delAddr)
	if err != nil {
		return err
	}

	// 2. Emit claim_pegbridge_fees event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaimPegBridgeFees,
		sdk.NewAttribute(types.AttributeKeyDelegatorAddress, delAddr.String()),
	))

	return nil
}
