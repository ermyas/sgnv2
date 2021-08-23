package types

import (
	"github.com/celer-network/sgn-v2/contracts"
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewValidator(ethAddress, ethSigner, sgnAddress string) *Validator {
	return &Validator{
		EthAddress: contracts.FormatAddrHex(ethAddress),
		EthSigner:  contracts.FormatAddrHex(ethSigner),
		SgnAddress: sgnAddress,
		Status:     ValidatorStatus_Unbonded,
	}
}

func MustMarshalValidator(cdc codec.BinaryCodec, validator *Validator) []byte {
	return cdc.MustMarshal(validator)
}

func MustUnmarshalValidator(cdc codec.BinaryCodec, value []byte) Validator {
	validator, err := UnmarshalValidator(cdc, value)
	if err != nil {
		panic(err)
	}

	return validator
}

func UnmarshalValidator(cdc codec.BinaryCodec, value []byte) (v Validator, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}
