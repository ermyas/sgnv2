package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

const (
	QuerySyncer     = "syncer"
	QueryValidator  = "validator"
	QueryValidators = "validators"
	QueryDelegator  = "delegator"
	QueryDelegators = "delegators"
	QueryParameters = "parameters"
)

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

type QueryValidatorParams struct {
	ValidatorAddress string
}

func NewQueryValidatorParams(validatorAddress string) QueryValidatorParams {
	return QueryValidatorParams{
		ValidatorAddress: contracts.FormatAddrHex(validatorAddress),
	}
}
