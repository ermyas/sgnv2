package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

func NewValidator(ethAddress, ethSigner, sgnAddress string) Validator {
	return Validator{
		EthAddress: contracts.FormatAddrHex(ethAddress),
		EthSigner:  contracts.FormatAddrHex(ethSigner),
		SgnAddress: sgnAddress,
	}
}
