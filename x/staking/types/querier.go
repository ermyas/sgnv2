package types

import "github.com/celer-network/sgn-v2/eth"

const (
	QueryValidator           = "validator"
	QueryValidatorBySgnAddr  = "validator-sgnaddr"
	QueryValidatorByConsAddr = "validator-consaddr"
	QueryValidators          = "validators"
	QueryDelegator           = "delegator"
	QueryDelegators          = "delegators"
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

type QueryDelegatorParams struct {
	ValAddress string
	DelAddress string
}

func NewQueryDelegatorParams(valAddress, delAddress string) QueryDelegatorParams {
	return QueryDelegatorParams{
		ValAddress: eth.FormatAddrHex(valAddress),
		DelAddress: eth.FormatAddrHex(delAddress),
	}
}

type QueryDelegatorsParams struct {
	ValAddress string
}

func NewQueryDelegatorsParams(ethAddress string) QueryDelegatorsParams {
	return QueryDelegatorsParams{
		ValAddress: eth.FormatAddrHex(ethAddress),
	}
}

type QuerySgnAccountParams struct {
	SgnAddress string
}

func NewQuerySgnAccountParams(sgnAddress string) QuerySgnAccountParams {
	return QuerySgnAccountParams{SgnAddress: sgnAddress}
}
