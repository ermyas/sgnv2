package keeper

import (
	"context"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) OnchainEvent(goCtx context.Context, msg *types.MsgOnchainEvent) (*types.MsgOnchainEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgOnchainEventResponse{}, nil
}
