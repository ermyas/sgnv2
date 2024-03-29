package cbridge

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/cbridge/keeper"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgInitWithdraw:
			res, err := msgServer.InitWithdraw(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				log.Warn(err)
			}
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgSendMySig:
			res, err := msgServer.SendMySig(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				log.Error(err)
			}
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgSignAgain:
			res, err := msgServer.SignAgain(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				log.Warn(err)
			}
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUpdateLatestSigners:
			res, err := msgServer.UpdateLatestSigners(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				log.Warn(err)
			}
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgSyncFarming:
			res, err := msgServer.SyncFarming(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				log.Warn(err)
			}
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgTriggerSetRefund:
			res, err := msgServer.TriggerSetRefund(sdk.WrapSDKContext(ctx), msg)
			if err != nil {
				log.Warn(err)
			}
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			log.Error(errMsg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
