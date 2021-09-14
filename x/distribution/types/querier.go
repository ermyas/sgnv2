package types

import "github.com/celer-network/sgn-v2/eth"

// querier keys
const (
	QueryParams                      = "params"
	QueryValidatorOutstandingRewards = "validator_outstanding_rewards"
	QueryValidatorCommission         = "validator_commission"
	QueryValidatorSlashes            = "validator_slashes"
	QueryDelegationRewards           = "delegation_rewards"
	QueryDelegatorTotalRewards       = "delegator_total_rewards"
	QueryDelegatorValidators         = "delegator_validators"
	QueryWithdrawAddr                = "withdraw_addr"
	QueryCommunityPool               = "community_pool"
)

// params for query 'custom/distr/validator_outstanding_rewards'
type QueryValidatorOutstandingRewardsParams struct {
	ValidatorAddress eth.Addr `json:"validator_address" yaml:"validator_address"`
}

// creates a new instance of QueryValidatorOutstandingRewardsParams
func NewQueryValidatorOutstandingRewardsParams(validatorAddr eth.Addr) QueryValidatorOutstandingRewardsParams {
	return QueryValidatorOutstandingRewardsParams{
		ValidatorAddress: validatorAddr,
	}
}

// params for query 'custom/distr/validator_commission'
type QueryValidatorCommissionParams struct {
	ValidatorAddress eth.Addr `json:"validator_address" yaml:"validator_address"`
}

// creates a new instance of QueryValidatorCommissionParams
func NewQueryValidatorCommissionParams(validatorAddr eth.Addr) QueryValidatorCommissionParams {
	return QueryValidatorCommissionParams{
		ValidatorAddress: validatorAddr,
	}
}

// params for query 'custom/distr/validator_slashes'
type QueryValidatorSlashesParams struct {
	ValidatorAddress eth.Addr `json:"validator_address" yaml:"validator_address"`
	StartingHeight   uint64   `json:"starting_height" yaml:"starting_height"`
	EndingHeight     uint64   `json:"ending_height" yaml:"ending_height"`
}

// creates a new instance of QueryValidatorSlashesParams
func NewQueryValidatorSlashesParams(validatorAddr eth.Addr, startingHeight uint64, endingHeight uint64) QueryValidatorSlashesParams {
	return QueryValidatorSlashesParams{
		ValidatorAddress: validatorAddr,
		StartingHeight:   startingHeight,
		EndingHeight:     endingHeight,
	}
}

// params for query 'custom/distr/delegation_rewards'
type QueryDelegationRewardsParams struct {
	DelegatorAddress eth.Addr `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress eth.Addr `json:"validator_address" yaml:"validator_address"`
}

// creates a new instance of QueryDelegationRewardsParams
func NewQueryDelegationRewardsParams(delegatorAddr eth.Addr, validatorAddr eth.Addr) QueryDelegationRewardsParams {
	return QueryDelegationRewardsParams{
		DelegatorAddress: delegatorAddr,
		ValidatorAddress: validatorAddr,
	}
}

// params for query 'custom/distr/delegator_total_rewards' and 'custom/distr/delegator_validators'
type QueryDelegatorParams struct {
	DelegatorAddress eth.Addr `json:"delegator_address" yaml:"delegator_address"`
}

// creates a new instance of QueryDelegationRewardsParams
func NewQueryDelegatorParams(delegatorAddr eth.Addr) QueryDelegatorParams {
	return QueryDelegatorParams{
		DelegatorAddress: delegatorAddr,
	}
}

// params for query 'custom/distr/withdraw_addr'
type QueryDelegatorWithdrawAddrParams struct {
	DelegatorAddress eth.Addr `json:"delegator_address" yaml:"delegator_address"`
}

// NewQueryDelegatorWithdrawAddrParams creates a new instance of QueryDelegatorWithdrawAddrParams.
func NewQueryDelegatorWithdrawAddrParams(delegatorAddr eth.Addr) QueryDelegatorWithdrawAddrParams {
	return QueryDelegatorWithdrawAddrParams{DelegatorAddress: delegatorAddr}
}
