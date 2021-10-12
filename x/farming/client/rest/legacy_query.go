package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	// get the current state of the all farming pools
	r.HandleFunc(
		"/farming/pools",
		queryPoolsHandlerFn(cliCtx),
	).Methods("GET")

	// get a single pool info by the farming pool's name
	r.HandleFunc(
		"/farming/pool/{poolName}",
		queryPoolHandlerFn(cliCtx),
	).Methods("GET")

	// get the current earnings of an account in a farming pool
	r.HandleFunc(
		"/farming/earnings/{poolName}/{addr}",
		queryEarningsHandlerFn(cliCtx),
	).Methods("GET")

	// get all the farming pools that the account has staked tokens in
	r.HandleFunc(
		"/farming/staked-pools/{addr}",
		QueryAccountHandlerFn(cliCtx),
	).Methods("GET")
}

func QueryAccountHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		addr := eth.Hex2Addr(mux.Vars(r)["addr"])
		jsonBytes, err := cliCtx.LegacyAmino.MarshalJSON(types.NewQueryAccountParams(addr))
		if rest.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAccountInfo)
		res, height, err := cliCtx.QueryWithData(route, jsonBytes)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryEarningsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		varsMap := mux.Vars(r)
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		addr := eth.Hex2Addr(mux.Vars(r)["addr"])
		jsonBytes, err := cliCtx.LegacyAmino.MarshalJSON(types.NewQueryPoolAccountParams(varsMap["poolName"], addr))
		if rest.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryEarnings)
		res, height, err := cliCtx.QueryWithData(route, jsonBytes)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryPoolHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		poolName := mux.Vars(r)["poolName"]
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryPoolParams(poolName)

		jsonBytes, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryPool)
		res, height, err := cliCtx.QueryWithData(route, jsonBytes)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryPoolsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryPoolsParams(page, limit)
		jsonBytes, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryPools)
		res, height, err := cliCtx.QueryWithData(route, jsonBytes)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
