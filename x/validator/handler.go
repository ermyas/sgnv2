package validator

import (
	"fmt"

	"github.com/celer-network/sgn-v2/seal"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "validator" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		logEntry := seal.NewMsgLog()
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var res *sdk.Result
		var err error

		switch msg := msg.(type) {
		// case MsgSetTransactors:
		// 	res, err = handleMsgSetTransactors(ctx, keeper, msg, logEntry)
		// case MsgEditValidatorDescription:
		// 	res, err = handleMsgEditValidatorDescription(ctx, keeper, msg, logEntry)
		default:
			return nil, sdk_errors.Wrapf(sdk_errors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}

		if err != nil {
			logEntry.Error = append(logEntry.Error, err.Error())
		}

		seal.CommitMsgLog(logEntry)
		return res, err
	}
}

// Handle a message to set transactors
func handleMsgSetTransactors(ctx sdk.Context, keeper Keeper, msg MsgSetTransactors, logEntry *seal.MsgLog) (*sdk.Result, error) {
	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender.String()

	sgnVal, found := keeper.GetSgnValidator(ctx, sdk.ValAddress(msg.Sender))
	if !found {
		return nil, fmt.Errorf("Sender is not a validator")
	}

	validator, found := keeper.GetValidator(ctx, sgnVal.Description.Identity)
	if !found {
		return nil, fmt.Errorf("Validator does not exist")
	}

	dedup := make(map[string]bool)
	oldTransactors := validator.Transactors
	validator.Transactors = []sdk.AccAddress{}
	for _, transactor := range msg.Transactors {
		if !transactor.Equals(validator.SgnAddress) {
			if _, exist := dedup[transactor.String()]; !exist {
				logEntry.Transactor = append(logEntry.Transactor, transactor.String())
				validator.Transactors = append(validator.Transactors, transactor)
				dedup[transactor.String()] = true
				keeper.InitAccount(ctx, transactor)
			}
		}
	}

	for _, transactor := range oldTransactors {
		if _, exist := dedup[transactor.String()]; !exist {
			keeper.RemoveAccount(ctx, transactor)
		}
	}

	keeper.SetValidator(ctx, validator)
	return &sdk.Result{}, nil
}

// Handle a message to edit validator description
func handleMsgEditValidatorDescription(ctx sdk.Context, keeper Keeper, msg MsgEditValidatorDescription, logEntry *seal.MsgLog) (*sdk.Result, error) {
	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender.String()
	logEntry.EthAddress = msg.EthAddress

	validator, found := keeper.GetValidator(ctx, msg.EthAddress)
	if !found {
		return nil, fmt.Errorf("Validator does not exist")
	}

	description, err := validator.Description.UpdateDescription(msg.Description)
	if err != nil {
		return nil, err
	}

	validator.Description = description
	keeper.SetValidator(ctx, validator)
	return &sdk.Result{}, nil
}
