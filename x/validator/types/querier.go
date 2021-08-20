package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

const (
	QuerySyncer              = "syncer"
	QueryDelegator           = "delegator"
	QueryValidator           = "validator"
	QueryValidators          = "validators"
	QueryValidatorDelegators = "validator-delegators"
	QueryReward              = "reward"
	QueryRewardEpoch         = "reward-epoch"
	QueryRewardStats         = "reward-stats"
	QueryParameters          = "parameters"
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

type QueryRewardParams struct {
	EthAddress string
}

func NewQueryRewardParams(ethAddress string) QueryRewardParams {
	return QueryRewardParams{
		EthAddress: contracts.FormatAddrHex(ethAddress),
	}
}
