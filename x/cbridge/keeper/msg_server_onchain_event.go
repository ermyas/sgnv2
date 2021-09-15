package keeper

import (
	"context"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// must match cbr_monitor
const (
	// event names
	CbrEventSend  = "Send"
	CbrEventRelay = "Relay"
	// from pool.sol
	CbrEventLiqAdd   = "LiquidityAdded"
	CbrEventWithdraw = "WithdrawDone" // could be LP or user
	// from signers.sol
	CbrEventNewSigners = "SignersUpdated"
)

func (k msgServer) OnchainEvent(goCtx context.Context, msg *types.MsgOnchainEvent) (*types.MsgOnchainEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.handleOneEvent(ctx, msg)

	return &types.MsgOnchainEventResponse{}, nil
}

func (k msgServer) OnchainManyEvents(goCtx context.Context, msg *types.MsgOnchainManyEvents) (*types.MsgOnchainManyEventsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, ev := range msg.Events {
		k.handleOneEvent(ctx, ev)
	}

	return &types.MsgOnchainManyEventsResponse{}, nil
}

func (k msgServer) handleOneEvent(ctx sdk.Context, ev *types.MsgOnchainEvent) {
	switch ev.Evtype {
	case CbrEventLiqAdd:
		k.handleAddLiq(ctx, ev)
	}
}

func (k msgServer) handleAddLiq(ctx sdk.Context, ev *types.MsgOnchainEvent) {
	ctx.EventManager().EmitEvent(sdk.NewEvent("liqAdded", sdk.NewAttribute("tbd-key", "tbd-value")))
}
