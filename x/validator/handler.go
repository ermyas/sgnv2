package validator

import (
	"fmt"

	"github.com/celer-network/sgn-v2/seal"
	"github.com/celer-network/sgn-v2/x/validator/keeper"
	"github.com/celer-network/sgn-v2/x/validator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		logEntry := seal.NewMsgLog()
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var res *sdk.Result
		var err error

		switch msg := msg.(type) {
		case *types.MsgSetTransactors:
			res, err = handleMsgSetTransactors(ctx, keeper, msg, logEntry)
		case *types.MsgEditDescription:
			res, err = handleMsgEditDescription(ctx, keeper, msg, logEntry)
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

// Handle a message to set transactors
func handleMsgSetTransactors(
	ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSetTransactors, logEntry *seal.MsgLog) (*sdk.Result, error) {

	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender

	sdkVal, found := keeper.GetSdkValidator(ctx, sdk.ValAddress(msg.Sender))
	if !found {
		return nil, fmt.Errorf("Sender is not a validator")
	}

	validator, found := keeper.GetValidator(ctx, sdkVal.Description.Identity)
	if !found {
		return nil, fmt.Errorf("Validator does not exist")
	}

	dedup := make(map[string]bool)
	oldTransactors := validator.Transactors
	validator.Transactors = []string{}
	for _, transactor := range msg.Transactors {
		if transactor != (validator.SgnAddress) {
			if _, exist := dedup[transactor]; !exist {
				logEntry.Transactor = append(logEntry.Transactor, transactor)
				validator.Transactors = append(validator.Transactors, transactor)
				dedup[transactor] = true
				acctAddr, err := types.SdkAccAddrFromSgnBech32(transactor)
				if err != nil {
					return nil, fmt.Errorf("Invalid bech32 addr %s, %s", transactor, err)
				}
				keeper.InitAccount(ctx, acctAddr)
			}
		}
	}

	for _, transactor := range oldTransactors {
		if _, exist := dedup[transactor]; !exist {
			acctAddr, err := types.SdkAccAddrFromSgnBech32(transactor)
			if err != nil {
				return nil, fmt.Errorf("Invalid bech32 addr %s, %s", transactor, err)
			}
			keeper.RemoveAccount(ctx, acctAddr)
		}
	}

	keeper.SetValidator(ctx, validator)
	return &sdk.Result{}, nil
}

// Handle a message to edit validator description
func handleMsgEditDescription(
	ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgEditDescription, logEntry *seal.MsgLog) (*sdk.Result, error) {

	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Sender
	logEntry.EthAddress = msg.EthAddress

	validator, found := keeper.GetValidator(ctx, msg.EthAddress)
	if !found {
		return nil, fmt.Errorf("Validator does not exist")
	}
	// TODO: copy update validator description from sdk_staking

	validator.Description = msg.Description
	keeper.SetValidator(ctx, validator)
	return &sdk.Result{}, nil
}
