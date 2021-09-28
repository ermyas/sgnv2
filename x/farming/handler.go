package farming

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/celer-network/sgn-v2/x/farming/keeper"
	"github.com/celer-network/sgn-v2/x/farming/types"
)

// NewHandler creates an sdk.Handler for all the farming type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgClaimAllRewards:
			res, err := msgServer.ClaimAllRewards(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		// TODO: MsgClaimRewards
		case *types.MsgSignRewards:
			res, err := msgServer.SignRewards(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, types.WrapErrUnknownFarmingMsgType(errMsg)
		}
	}
}
