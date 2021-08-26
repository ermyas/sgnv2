package sync

import (
	"github.com/celer-network/sgn-v2/seal"
	"github.com/celer-network/sgn-v2/x/sync/keeper"
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		logEntry := seal.NewMsgLog(types.ModuleName)
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var res *sdk.Result
		var err error

		switch msg := msg.(type) {
		case *types.MsgProposeUpdates:
			res, err = handleMsgProposeUpdates(ctx, keeper, msg, logEntry)
		case *types.MsgVoteUpdates:
			res, err = handleMsgVoteUpdates(ctx, keeper, msg, logEntry)
		default:
			return nil, sdk_errors.Wrapf(sdk_errors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}

		if err != nil {
			logEntry.Error = append(logEntry.Error, err.Error())
		}

		seal.CommitMsgLog(logEntry)
		return res, err
	}
}

func handleMsgProposeUpdates(
	ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgProposeUpdates, logEntry *seal.MsgLog) (*sdk.Result, error) {

	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender

	err := keeper.ProposeUpdates(ctx, msg.Updates, msg.EthBlock, msg.Sender, logEntry)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

func handleMsgVoteUpdates(
	ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgVoteUpdates, logEntry *seal.MsgLog) (*sdk.Result, error) {

	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender

	keeper.VoteUpdates(ctx, msg.Votes, msg.Sender, logEntry)
	return &sdk.Result{}, nil
}
