package message

import (
	"fmt"

	"github.com/celer-network/sgn-v2/common"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	msgkeeper "github.com/celer-network/sgn-v2/x/message/keeper"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgProposalHandler(k msgkeeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MsgProposal:
			return handleMsgProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unsupported peg proposal content type: %T", c)
		}
	}
}

func handleMsgProposal(ctx sdk.Context, k msgkeeper.Keeper, p *types.MsgProposal) error {
	for _, bus := range p.MessageBuses {
		if !common.IsHexAddress(bus.ContractInfo.Address) {
			return fmt.Errorf("invalid message bus address %s", bus.String())
		}
	}
	for _, bus := range p.MessageBuses {
		k.SetMessageBus(ctx, *bus)
	}
	return nil
}
