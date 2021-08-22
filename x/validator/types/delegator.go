package types

import (
	"github.com/celer-network/sgn-v2/contracts"
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewDelegator(valAddr, delAddr string) *Delegator {
	return &Delegator{
		ValAddress: contracts.FormatAddrHex(valAddr),
		DelAddress: contracts.FormatAddrHex(delAddr),
	}
}

func MustMarshalDelegator(cdc codec.BinaryCodec, delegator *Delegator) []byte {
	return cdc.MustMarshal(delegator)
}

func MustUnmarshalDelegator(cdc codec.BinaryCodec, value []byte) Delegator {
	delegator, err := UnmarshalDelegator(cdc, value)
	if err != nil {
		panic(err)
	}

	return delegator
}

func UnmarshalDelegator(cdc codec.BinaryCodec, value []byte) (d Delegator, err error) {
	err = cdc.Unmarshal(value, &d)
	return d, err
}
