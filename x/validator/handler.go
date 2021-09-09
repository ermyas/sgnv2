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

	sgnAddr, err := types.SdkAccAddrFromSgnBech32(msg.Sender)
	if err != nil {
		return nil, fmt.Errorf("invalid sender addr %s, %s", msg.Sender, err)
	}
	validator, found := keeper.GetValidatorBySgnAddr(ctx, sgnAddr)
	if !found {
		return nil, fmt.Errorf("validator does not exist")
	}
	logEntry.ValAddr = validator.EthAddress
	if validator.Status != types.ValidatorStatus_Bonded {
		return nil, fmt.Errorf("validator is not bonded")
	}

	dedup := make(map[string]bool)
	oldTransactors := validator.Transactors
	validator.Transactors = []string{}
	for _, transactor := range msg.Transactors {
		if transactor != (validator.SgnAddress) {
			if _, exist := dedup[transactor]; !exist {
				logEntry.Transactors = append(logEntry.Transactors, transactor)
				validator.Transactors = append(validator.Transactors, transactor)
				dedup[transactor] = true
				acctAddr, err := types.SdkAccAddrFromSgnBech32(transactor)
				if err != nil {
					return nil, fmt.Errorf("invalid bech32 addr %s, %s", transactor, err)
				}
				keeper.InitAccount(ctx, acctAddr)
			}
		}
	}

	for _, transactor := range oldTransactors {
		if _, exist := dedup[transactor]; !exist {
			acctAddr, err := types.SdkAccAddrFromSgnBech32(transactor)
			if err != nil {
				return nil, fmt.Errorf("invalid bech32 addr %s, %s", transactor, err)
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

	validator, found := keeper.GetValidator(ctx, msg.EthAddress)
	if !found {
		return nil, fmt.Errorf("validator does not exist")
	}
	logEntry.ValAddr = validator.EthAddress
	// TODO: copy update validator description from sdk_staking

	validator.Description = msg.Description
	keeper.SetValidator(ctx, validator)
	return &sdk.Result{}, nil
}
