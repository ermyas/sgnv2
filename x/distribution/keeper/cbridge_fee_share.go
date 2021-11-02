package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetWithdrawableCBridgeFeeShare(ctx sdk.Context, delAddr eth.Addr, coin sdk.Coin) sdk.Coin {
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, delAddr)
	return k.bankKeeper.GetBalance(ctx, derivedAccAddress, coin.Denom)
}

func (k Keeper) ClaimCBridgeFeeShare(ctx sdk.Context, delAddr eth.Addr) error {
	// 1. Withdraw reward for all validators
	allValidators := k.stakingKeeper.GetAllValidators(ctx)
	for _, validator := range allValidators {
		valAddr := eth.Hex2Addr(validator.EthAddress)
		// TODO: Check residual
		delegation := k.stakingKeeper.Delegation(ctx, delAddr, valAddr)
		if delegation != nil {
			withdrawErr := k.withdrawDelegatorRewardForOneValidator(ctx, delAddr.String(), validator.EthAddress)
			if withdrawErr != nil {
				return withdrawErr
			}
		}
	}

	// 2. Emit claim_cbridge_fee_share event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaimCBridgeFeeShare,
		sdk.NewAttribute(types.AttributeKeyDelegatorAddress, delAddr.String()),
	))

	return nil
}
