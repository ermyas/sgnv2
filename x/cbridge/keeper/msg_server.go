package keeper

import (
	"context"

	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) InitWithdraw(context.Context, *types.MsgInitWithdraw) (*types.MsgInitWithdrawResp, error) {
	return nil, nil
}

// user can request to sign a previous withdraw again
// to mitigate dos attack, we could be smart and re-use sigs if
// they are still valid. we should also deny if withdraw already
// completed
func (k msgServer) SignAgain(context.Context, *types.MsgSignAgain) (*types.MsgSignAgainResp, error) {
	return nil, nil
}

// send my sig for data, so it can be later submitted onchain
func (k msgServer) SendMySig(context.Context, *types.MsgSendMySig) (*types.MsgSendMySigResp, error) {
	// zhihua
	// if type is relay, find xfer id, update XferRelayKey
	// note sigs need to be sorted by signer address
	return nil, nil
}
