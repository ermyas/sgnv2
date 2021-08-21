package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

func NewDelegator(validatorAddr, delegatorAddr string) Delegator {
	return Delegator{
		EthAddress: contracts.FormatAddrHex(delegatorAddr),
		ValAddress: contracts.FormatAddrHex(delegatorAddr),
	}
}
