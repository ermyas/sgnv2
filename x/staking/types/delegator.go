package types

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDelegator(valAddr, delAddr string, shares sdk.Int) *Delegator {
	return &Delegator{
		ValAddress: eth.FormatAddrHex(valAddr),
		DelAddress: eth.FormatAddrHex(delAddr),
		Shares:     shares,
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
