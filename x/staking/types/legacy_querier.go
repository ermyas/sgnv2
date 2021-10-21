package types

import "github.com/celer-network/sgn-v2/eth"

// query endpoints supported by the staking legacy querier
const (
	QueryValidators           = "validators"
	QueryValidator            = "validator"
	QueryValidatorBySgnAddr   = "validator-sgnaddr"
	QueryValidatorByConsAddr  = "validator-consaddr"
	QueryValidatorDelegations = "validator-delegations"
	QueryDelegation           = "delegation"
	QueryDelegatorDelegations = "delegator-delegations"
	QueryDelegatorValidators  = "delegator-validators"
	QueryDelegatorValidator   = "delegator-validator"
	QueryTransactors          = "transactors"
	QuerySgnAccount           = "account"
	QuerySyncer               = "syncer"
	QueryParams               = "params"
)

// defines the params for the following queries:
// - 'custom/staking/delegatorDelegations'
// - 'custom/staking/delegatorValidators'
type QueryDelegatorParams struct {
	DelegatorAddr string
}

func NewQueryDelegatorParams(delegatorAddr string) QueryDelegatorParams {
	return QueryDelegatorParams{
		DelegatorAddr: eth.FormatAddrHex(delegatorAddr),
	}
}

// defines the params for the following queries:
// - 'custom/staking/validator'
// - 'custom/staking/validatorDelegations'
type QueryValidatorParams struct {
	EthAddress  string
	Page, Limit int
}

func NewQueryValidatorParams(ethAddress string, page, limit int) QueryValidatorParams {
	return QueryValidatorParams{
		EthAddress: eth.FormatAddrHex(ethAddress),
		Page:       page,
		Limit:      limit,
	}
}

type QueryValidatorBySgnAddrParams struct {
	SgnAddress string
}

func NewQueryValidatorBySgnAddrParams(sgnAddress string) QueryValidatorBySgnAddrParams {
	return QueryValidatorBySgnAddrParams{
		SgnAddress: sgnAddress,
	}
}

type QueryValidatorByConsAddrParams struct {
	ConsAddress string
}

func NewQueryValidatorByConsAddrParams(consAddress string) QueryValidatorByConsAddrParams {
	return QueryValidatorByConsAddrParams{
		ConsAddress: consAddress,
	}
}

type QueryTransactorsParams struct {
	ValAddress string
}

func NewQueryTransactorsParams(valAddress string) QueryTransactorsParams {
	return QueryTransactorsParams{
		ValAddress: eth.FormatAddrHex(valAddress),
	}
}

type QueryDelegationParams struct {
	ValAddress string
	DelAddress string
}

func NewQueryDelegationParams(valAddress, delAddress string) QueryDelegationParams {
	return QueryDelegationParams{
		ValAddress: eth.FormatAddrHex(valAddress),
		DelAddress: eth.FormatAddrHex(delAddress),
	}
}

type QueryDelegationsParams struct {
	ValAddress string
}

func NewQueryDelegationsParams(ethAddress string) QueryDelegationsParams {
	return QueryDelegationsParams{
		ValAddress: eth.FormatAddrHex(ethAddress),
	}
}

type QuerySgnAccountParams struct {
	SgnAddress string
}

func NewQuerySgnAccountParams(sgnAddress string) QuerySgnAccountParams {
	return QuerySgnAccountParams{SgnAddress: sgnAddress}
}

// QueryValidatorsParams defines the params for the following queries:
// - 'custom/staking/validators'
type QueryValidatorsParams struct {
	Page, Limit int
	Status      string
}

func NewQueryValidatorsParams(page, limit int, status string) QueryValidatorsParams {
	return QueryValidatorsParams{page, limit, status}
}
