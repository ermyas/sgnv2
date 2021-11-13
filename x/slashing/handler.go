package slashing

import (
	"fmt"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/seal"
	"github.com/celer-network/sgn-v2/x/slashing/keeper"
	"github.com/celer-network/sgn-v2/x/slashing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "slash" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		logEntry := seal.NewMsgLog(types.ModuleName)
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		var res *sdk.Result
		var err error
		switch msg := msg.(type) {
		case *types.MsgSignSlash:
			res, err = handleMsgSignSlash(ctx, keeper, msg, logEntry)
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

// Handle a message to sign slash
func handleMsgSignSlash(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSignSlash, logEntry *seal.MsgLog) (*sdk.Result, error) {
	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender
	logEntry.Slash.Nonce = msg.Nonce

	res := &sdk.Result{}
	senderAcct, _ := sdk.AccAddressFromBech32(msg.Sender)
	validator, found := keeper.StakingKeeper.GetValidatorBySgnAddr(ctx, senderAcct)
	if !found {
		return res, fmt.Errorf("sender is not a validator")
	}
	if !validator.IsBonded() {
		return res, fmt.Errorf("validator is not bonded")
	}

	slash, found := keeper.GetSlash(ctx, msg.Nonce)
	if !found {
		return res, fmt.Errorf("slash does not exist")
	}
	logEntry.Slash.ValAddr = slash.Validator
	logEntry.Slash.Reason = slash.Reason

	err := slash.AddSig(msg.Sig, eth.Addr2Hex(validator.GetSignerAddr()))
	if err != nil {
		return res, fmt.Errorf(fmt.Sprintf("failed to add sig: %s", err))
	}

	keeper.SetSlash(ctx, slash)
	return res, nil
}
