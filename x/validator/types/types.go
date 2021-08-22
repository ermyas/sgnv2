package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

func NewValidator(ethAddress, ethSigner, sgnAddress string) *Validator {
	return &Validator{
		EthAddress: contracts.FormatAddrHex(ethAddress),
		EthSigner:  contracts.FormatAddrHex(ethSigner),
		SgnAddress: sgnAddress,
	}
}

func NewDelegator(validatorAddr, delegatorAddr string) *Delegator {
	return &Delegator{
		EthAddress: contracts.FormatAddrHex(delegatorAddr),
		ValAddress: contracts.FormatAddrHex(delegatorAddr),
	}
}

func NewSyncer(validatorIdx uint64, sgnAddress string) *Syncer {
	return &Syncer{
		ValIndex:   validatorIdx,
		SgnAddress: sgnAddress,
	}
}
