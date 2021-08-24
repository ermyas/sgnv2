package types

import "github.com/celer-network/sgn-v2/eth"

const (
	QueryValidator  = "validator"
	QueryValidators = "validators"
	QueryDelegator  = "delegator"
	QueryDelegators = "delegators"
	QuerySyncer     = "syncer"
	QueryParams     = "params"
)

type QueryValidatorParams struct {
	EthAddress string
}

func NewQueryValidatorParams(ethAddress string) QueryValidatorParams {
	return QueryValidatorParams{
		EthAddress: eth.FormatAddrHex(ethAddress),
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
