package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// Get all delegations from a delegator
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/delegations",
		delegatorDelegationsHandlerFn(clientCtx),
	).Methods("GET")

	// Query all validators that a delegator is bonded to
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/validators",
		delegatorValidatorsHandlerFn(clientCtx),
	).Methods("GET")

	// Query a validator that a delegator is bonded to
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/validators/{validatorAddr}",
		delegatorValidatorHandlerFn(clientCtx),
	).Methods("GET")

	// Query a delegation between a delegator and a validator
	r.HandleFunc(
		"/staking/delegators/{delegatorAddr}/delegations/{validatorAddr}",
		delegationHandlerFn(clientCtx),
	).Methods("GET")

	// Get all validators
	r.HandleFunc(
		"/staking/validators",
		validatorsHandlerFn(clientCtx),
	).Methods("GET")

	// Get a single validator info
	r.HandleFunc(
		"/staking/validators/{validatorAddr}",
		validatorHandlerFn(clientCtx),
	).Methods("GET")

	// Get all delegations to a validator
	r.HandleFunc(
		"/staking/validators/{validatorAddr}/delegations",
		validatorDelegationsHandlerFn(clientCtx),
	).Methods("GET")

	// Get the current staking parameter values
	r.HandleFunc(
		"/staking/parameters",
		paramsHandlerFn(clientCtx),
	).Methods("GET")
}

// HTTP request handler to query a delegator delegations
func delegatorDelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return queryDelegator(clientCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegatorDelegations))
}

// HTTP request handler to query a delegation
func delegationHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return queryBonds(clientCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegation))
}

// HTTP request handler to query all delegator bonded validators
func delegatorValidatorsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return queryDelegator(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegatorValidators))
}

// HTTP request handler to get information from a currently bonded validator
func delegatorValidatorHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return queryBonds(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegatorValidator))
}

// HTTP request handler to query list of validators
func validatorsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		status := r.FormValue("status")
		// These are query params that were available in =<0.39. We show a nice
		// error message for this breaking change.
		if status == "bonded" || status == "unbonding" || status == "unbonded" {
			err := fmt.Errorf("cosmos sdk v0.40 introduces a breaking change on this endpoint:"+
				" instead of querying using `?status=%s`, please use `status=BOND_STATUS_%s`. For more"+
				" info, please see our REST endpoint migration guide at %s", status, strings.ToUpper(status), clientrest.DeprecationURL)

			if rest.CheckBadRequestError(w, err) {
				return
			}

		}

		if status == "" {
			status = types.Bonded.String()
		}

		params := types.NewQueryValidatorsParams(page, limit, status)

		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidators)

		res, height, err := clientCtx.QueryWithData(route, bz)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

// HTTP request handler to query the validator information from a given validator address
func validatorHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return queryValidator(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidator))
}

// HTTP request handler to query all unbonding delegations from a validator
func validatorDelegationsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return queryValidator(clientCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidatorDelegations))
}

// HTTP request handler to query the staking params values
func paramsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams), nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
