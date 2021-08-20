package types

import (
	"github.com/celer-network/sgn-v2/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Delegator struct {
	ValidatorAddr string  `json:"validator_addr"`
	DelegatorAddr string  `json:"delegator_addr"`
	Shares        sdk.Int `json:"shares"`
}

func NewDelegator(validatorAddr, delegatorAddr string) Delegator {
	return Delegator{
		ValidatorAddr: contracts.FormatAddrHex(validatorAddr),
		DelegatorAddr: contracts.FormatAddrHex(delegatorAddr),
	}
}

func (m Delegator) Reset()         { m = Delegator{} }
func (m Delegator) String() string { return proto.CompactTextString(m) }
func (Delegator) ProtoMessage()    {}
