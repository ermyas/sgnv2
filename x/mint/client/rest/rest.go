package rest

import (
	"net/http"

	govrest "github.com/celer-network/sgn-v2/x/gov/client/rest"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/mint/types"
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

// TODO add proto compatible Handler after x/gov migration

// AdjustProvisionsProposalRESTHandler returns an AdjustProvisionsProposalRESTHandler that exposes the adjust provisions REST handler with a given sub-route.
func AdjustProvisionsProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "mint_adjust_provisions",
		Handler:  postAdjustProvisionsProposalHandlerFn(clientCtx),
	}
}

func postAdjustProvisionsProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AdjustProvisionsProposalReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := types.NewAdjustProvisionsProposal(req.Title, req.Description, req.NewAnnualProvisions)

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
