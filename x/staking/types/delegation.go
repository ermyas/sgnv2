package types

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

// Implements Delegation interface
var _ DelegationI = Delegation{}

func NewDelegation(
	delegatorAddress eth.Addr,
	validatorAddress eth.Addr,
	shares sdk.Int,
) Delegation {
	return Delegation{
		DelegatorAddress: eth.Addr2Hex(delegatorAddress),
		ValidatorAddress: eth.Addr2Hex(validatorAddress),
		Shares:           shares,
	}
}

func MustMarshalDelegation(cdc codec.BinaryCodec, delegation Delegation) []byte {
	return cdc.MustMarshal(&delegation)
}

func MustUnmarshalDelegation(cdc codec.BinaryCodec, value []byte) Delegation {
	delegator, err := UnmarshalDelegation(cdc, value)
	if err != nil {
		panic(err)
	}

	return delegator
}

func UnmarshalDelegation(cdc codec.BinaryCodec, value []byte) (d Delegation, err error) {
	err = cdc.Unmarshal(value, &d)
	return d, err
}

func (d Delegation) GetDelegatorAddr() eth.Addr {
	delAddr := eth.Hex2Addr(d.DelegatorAddress)
	return delAddr
}

func (d Delegation) GetValidatorAddr() eth.Addr {
	valAddr := eth.Hex2Addr(d.ValidatorAddress)
	return valAddr
}

func (d Delegation) GetShares() sdk.Int { return d.Shares }

// String returns a human readable string representation of a Delegation.
func (d Delegation) String() string {
	return proto.CompactTextString(&d)
}