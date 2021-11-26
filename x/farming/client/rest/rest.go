package rest

import (
	"net/http"

	"github.com/celer-network/sgn-v2/x/farming/types"
	govrest "github.com/celer-network/sgn-v2/x/gov/client/rest"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
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
		SubRoute: "farming_add_pool",
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

// AddTokensProposalRESTHandler returns an AddTokensProposalRESTHandler that exposes the add tokens REST handler with a given sub-route.
func AddTokensProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "farming_add_tokens",
		Handler:  postAddTokensProposalHandlerFn(clientCtx),
	}
}

func postAddTokensProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddTokensProposalReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := types.NewAddTokensProposal(req.Title, req.Description, req.Tokens)

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

// AdjustRewardProposalRESTHandler returns an AdjustRewardProposalRESTHandler that exposes the adjust reward REST handler with a given sub-route.
func AdjustRewardProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "farming_adjust_reward",
		Handler:  postAdjustRewardProposalHandlerFn(clientCtx),
	}
}

func postAdjustRewardProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AdjustRewardProposalReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := types.NewAdjustRewardProposal(req.Title, req.Description, req.PoolName, req.RewardAdjustmentInputs)

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

// SetRewardContractsProposalRESTHandler returns an SetRewardContractsProposalRESTHandler that exposes the add tokens REST handler with a given sub-route.
func SetRewardContractsProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "farming_set_reward_contracts",
		Handler:  postSetRewardContractsProposalHandlerFn(clientCtx),
	}
}

func postSetRewardContractsProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SetRewardContractsProposalReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		content := types.NewSetRewardContractsProposal(req.Title, req.Description, req.RewardContracts)

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
