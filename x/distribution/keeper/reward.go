package keeper

import (
	"github.com/armon/go-metrics"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) withdrawDelegatorRewardForOneValidator(ctx sdk.Context, delAddrStr string, valAddrStr string) error {
	valAddr := eth.Hex2Addr(valAddrStr)
	delAddr := eth.Hex2Addr(delAddrStr)
	amount, err := k.WithdrawDelegationRewards(ctx, delAddr, valAddr)
	if err != nil {
		return err
	}

	defer func() {
		for _, a := range amount {
			if a.Amount.IsInt64() {
				telemetry.SetGaugeWithLabels(
					[]string{"tx", "msg", "withdraw_reward"},
					float32(a.Amount.Int64()),
					[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
				)
			}
		}
	}()

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, delAddrStr),
		),
	)

	return nil
}
