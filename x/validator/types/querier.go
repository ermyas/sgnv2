package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

const (
	QueryValidator  = "validator"
	QueryValidators = "validators"
	QueryDelegator  = "delegator"
	QueryDelegators = "delegators"
	QuerySyncer     = "syncer"
	QueryParameters = "parameters"
)

type QueryValidatorParams struct {
	ValidatorAddress string
}

func NewQueryValidatorParams(validatorAddress string) QueryValidatorParams {
	return QueryValidatorParams{
		ValidatorAddress: contracts.FormatAddrHex(validatorAddress),
	}
}

type QueryDelegatorParams struct {
	ValidatorAddress string
	DelegatorAddress string
}

func NewQueryDelegatorParams(validatorAddress, delegatorAddress string) QueryDelegatorParams {
	return QueryDelegatorParams{
		ValidatorAddress: contracts.FormatAddrHex(validatorAddress),
		DelegatorAddress: contracts.FormatAddrHex(delegatorAddress),
	}
}
