package gov

import (
	"fmt"

	"github.com/celer-network/sgn-v2/seal"
	"github.com/celer-network/sgn-v2/x/gov/keeper"
	"github.com/celer-network/sgn-v2/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the gov type messages
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		logEntry := seal.NewMsgLog(types.ModuleName)
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		var res *sdk.Result
		var err error

		switch msg := msg.(type) {
		case *types.MsgDeposit:
			res, err = handleMsgDeposit(ctx, keeper, msg, logEntry)

		case *types.MsgSubmitProposal:
			res, err = handleMsgSubmitProposal(ctx, keeper, msg, logEntry)

		case *types.MsgVote:
			res, err = handleMsgVote(ctx, keeper, msg, logEntry)

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

func handleMsgSubmitProposal(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgSubmitProposal, logEntry *seal.MsgLog) (*sdk.Result, error) {
	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Proposer

	proposal, err := keeper.SubmitProposal(ctx, msg.GetContent())
	if err != nil {
		return nil, err
	}

	logEntry.Govern.ProposalId = proposal.ProposalId
	acc, _ := sdk.AccAddressFromBech32(msg.Proposer)
	votingStarted, err := keeper.AddDeposit(ctx, proposal.ProposalId, acc, msg.InitialDeposit)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Proposer),
		),
	)

	submitEvent := sdk.NewEvent(types.EventTypeSubmitProposal, sdk.NewAttribute(types.AttributeKeyProposalType, msg.GetContent().ProposalType()))
	if votingStarted {
		submitEvent = submitEvent.AppendAttributes(
			sdk.NewAttribute(types.AttributeKeyVotingPeriodStart, fmt.Sprintf("%d", proposal.ProposalId)),
		)
	}
	ctx.EventManager().EmitEvent(submitEvent)

	return &sdk.Result{
		Data:   types.GetProposalIDBytes(proposal.ProposalId),
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

func handleMsgDeposit(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgDeposit, logEntry *seal.MsgLog) (*sdk.Result, error) {
	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Depositor
	logEntry.Govern.ProposalId = msg.ProposalId
	logEntry.Govern.Amount = msg.Amount.Uint64()

	acc, _ := sdk.AccAddressFromBech32(msg.Depositor)
	votingStarted, err := keeper.AddDeposit(ctx, msg.ProposalId, acc, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor),
		),
	)

	if votingStarted {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProposalDeposit,
				sdk.NewAttribute(types.AttributeKeyVotingPeriodStart, fmt.Sprintf("%d", msg.ProposalId)),
			),
		)
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMsgVote(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgVote, logEntry *seal.MsgLog) (*sdk.Result, error) {
	logEntry.Type = msg.Type()
	logEntry.Sender = msg.Voter
	logEntry.Govern.ProposalId = msg.ProposalId
	logEntry.Govern.Option = uint32(msg.Option)

	acc, _ := sdk.AccAddressFromBech32(msg.Voter)
	err := keeper.AddVote(ctx, msg.ProposalId, acc, msg.Option)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Voter),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
