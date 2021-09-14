package types

import "github.com/celer-network/sgn-v2/eth"

const (
	QueryValidator           = "validator"
	QueryValidatorBySgnAddr  = "validator-sgnaddr"
	QueryValidatorByConsAddr = "validator-consaddr"
	QueryValidators          = "validators"
	QueryDelegation          = "delegation"
	QueryDelegations         = "delegations"
	QuerySgnAccount          = "account"
	QuerySyncer              = "syncer"
	QueryParams              = "params"
)

type QueryValidatorParams struct {
	EthAddress string
}

func NewQueryValidatorParams(ethAddress string) QueryValidatorParams {
	return QueryValidatorParams{
		EthAddress: eth.FormatAddrHex(ethAddress),
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
