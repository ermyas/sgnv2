package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/client/common"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type (
	withdrawRewardsReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	}

	setWithdrawalAddrReq struct {
		BaseReq         rest.BaseReq `json:"base_req" yaml:"base_req"`
		WithdrawAddress eth.Addr     `json:"withdraw_address" yaml:"withdraw_address"`
	}

	fundCommunityPoolReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
		Amount  sdk.Coins    `json:"amount" yaml:"amount"`
	}
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// Withdraw all delegator rewards
	r.HandleFunc(
		"/distribution/delegators/{delegatorAddr}/rewards",
		newWithdrawDelegatorRewardsHandlerFn(clientCtx),
	).Methods("POST")

	// Withdraw delegation rewards
	r.HandleFunc(
		"/distribution/delegators/{delegatorAddr}/rewards/{validatorAddr}",
		newWithdrawDelegationRewardsHandlerFn(clientCtx),
	).Methods("POST")

	// Replace the rewards withdrawal address
	r.HandleFunc(
		"/distribution/delegators/{delegatorAddr}/withdraw_address",
		newSetDelegatorWithdrawalAddrHandlerFn(clientCtx),
	).Methods("POST")

	// Withdraw validator rewards and commission
	r.HandleFunc(
		"/distribution/validators/{validatorAddr}/rewards",
		newWithdrawValidatorRewardsHandlerFn(clientCtx),
	).Methods("POST")

	// Fund the community pool
	r.HandleFunc(
		"/distribution/community_pool",
		newFundCommunityPoolHandlerFn(clientCtx),
	).Methods("POST")
}

func newWithdrawDelegatorRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawRewardsReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// read and validate URL's variables
		delAddr, ok := checkDelegatorAddressVar(w, r)
		if !ok {
			return
		}

		msgs, err := common.WithdrawAllDelegatorRewards(clientCtx, delAddr)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msgs...)
	}
}

func newWithdrawDelegationRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawRewardsReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// read and validate URL's variables
		delAddr, ok := checkDelegatorAddressVar(w, r)
		if !ok {
			return
		}

		valAddr, ok := checkValidatorAddressVar(w, r)
		if !ok {
			return
		}

		msg := types.NewMsgWithdrawDelegatorReward(delAddr, valAddr)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newSetDelegatorWithdrawalAddrHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setWithdrawalAddrReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// read and validate URL's variables
		delAddr, ok := checkDelegatorAddressVar(w, r)
		if !ok {
			return
		}

		msg := types.NewMsgSetWithdrawAddress(delAddr, req.WithdrawAddress)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newWithdrawValidatorRewardsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req withdrawRewardsReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// read and validate URL's variable
		valAddr, ok := checkValidatorAddressVar(w, r)
		if !ok {
			return
		}

		valEthAddr, ok := checkValidatorEthAddressVar(w, r)
		if !ok {
			return
		}

		// prepare multi-message transaction
		msgs, err := common.WithdrawValidatorRewardsAndCommission(valAddr, valEthAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msgs...)
	}
}

func newFundCommunityPoolHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req fundCommunityPoolReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		msg := types.NewMsgFundCommunityPool(req.Amount, fromAddr)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

// Auxiliary

func checkDelegatorAddressVar(w http.ResponseWriter, r *http.Request) (eth.Addr, bool) {
	addr := eth.Hex2Addr(mux.Vars(r)["delegatorAddr"])
	return addr, true
}

func checkValidatorAddressVar(w http.ResponseWriter, r *http.Request) (eth.Addr, bool) {
	addr := eth.Hex2Addr(mux.Vars(r)["validatorAddr"])
	return addr, true
}

func checkValidatorEthAddressVar(w http.ResponseWriter, r *http.Request) (eth.Addr, bool) {
	addr := eth.Hex2Addr(mux.Vars(r)["validatorEthAddr"])
	return addr, true
}
