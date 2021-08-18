package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

const (
	QuerySyncer              = "syncer"
	QueryDelegator           = "delegator"
	QueryCandidate           = "candidate"
	QueryCandidates          = "candidates"
	QueryCandidateDelegators = "candidate-delegators"
	QueryReward              = "reward"
	QueryRewardEpoch         = "reward-epoch"
	QueryRewardStats         = "reward-stats"
	QueryParameters          = "parameters"
)

type QueryDelegatorParams struct {
	CandidateAddress string
	DelegatorAddress string
}

func NewQueryDelegatorParams(candidateAddress, delegatorAddress string) QueryDelegatorParams {
	return QueryDelegatorParams{
		CandidateAddress: contracts.FormatAddrHex(candidateAddress),
		DelegatorAddress: contracts.FormatAddrHex(delegatorAddress),
	}
}

type QueryCandidateParams struct {
	CandidateAddress string
}

func NewQueryCandidateParams(candidateAddress string) QueryCandidateParams {
	return QueryCandidateParams{
		CandidateAddress: contracts.FormatAddrHex(candidateAddress),
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
