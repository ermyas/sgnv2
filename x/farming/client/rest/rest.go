package rest

import (
	"net/http"

	"github.com/celer-network/sgn-v2/x/farming/types"
	govrest "github.com/celer-network/sgn-v2/x/gov/client/rest"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// RegisterRoutes registers farm-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

// TODO add proto compatible Handler after x/gov migration
// AddPoolProposalRESTHandler returns an AddPoolProposalRESTHandler that exposes the add pool REST handler with a given sub-route.
func AddPoolProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "add_pool",
		Handler:  postAddPoolProposalHandlerFn(clientCtx),
	}
}

func postAddPoolProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddPoolProposalReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := types.NewAddPoolProposal(
			req.Title, req.Description,
			req.PoolName, req.StakeToken, req.RewardTokens, req.InitialRewardInputs)

		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit[0].Amount, req.Proposer)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
