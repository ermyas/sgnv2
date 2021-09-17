package validator

import (
	"fmt"

	"github.com/celer-network/sgn-v2/seal"
	"github.com/celer-network/sgn-v2/x/staking/keeper"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		logEntry := seal.NewMsgLog(types.ModuleName)
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var res *sdk.Result
		var err error

		switch msg := msg.(type) {
		case *types.MsgSetTransactors:
			res, err = handleMsgSetTransactors(ctx, keeper, msg, logEntry)
		case *types.MsgEditDescription:
			res, err = handleMsgEditDescription(ctx, keeper, msg, logEntry)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}

		if err != nil {
			logEntry.Error = append(logEntry.Error, err.Error())
		}

		seal.CommitMsgLog(logEntry)
		return res, err
	}
}

// Handle a message to set transactors
func handleMsgSetTransactors(
	ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSetTransactors, logEntry *seal.MsgLog) (*sdk.Result, error) {

	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender
	sgnAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, fmt.Errorf("invalid sender %w", err)
	}
	err = keeper.SetTransactors(ctx, msg.GetOperation(), sgnAddr, msg.GetTransactors())
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

// Handle a message to edit validator description
func handleMsgEditDescription(
	ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgEditDescription, logEntry *seal.MsgLog) (*sdk.Result, error) {

	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender

	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	v, found := keeper.GetValidatorBySgnAddr(ctx, addr)
	if !found {
		return nil, fmt.Errorf("validator not found")
	}
	validator := v.(types.Validator)
	logEntry.Staking.ValAddr = validator.GetEthAddr().Hex()

	err = validator.Description.UpdateDescription(msg.Description)
	if err != nil {
		return nil, err
	}
	keeper.SetValidator(ctx, &validator)
	return &sdk.Result{}, nil
}
