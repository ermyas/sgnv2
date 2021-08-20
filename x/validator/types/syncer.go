package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Syncer struct {
	ValidatorIdx uint           `json:"validator_idx"`
	SgnAddress   sdk.AccAddress `json:"validator_acct"`
}

func NewSyncer(validatorIdx uint, sgnAddress sdk.AccAddress) Syncer {
	return Syncer{
		ValidatorIdx: validatorIdx,
		SgnAddress:   sgnAddress,
	}
}

func (m Syncer) Reset()         { m = Syncer{} }
func (m Syncer) String() string { return proto.CompactTextString(m) }
func (Syncer) ProtoMessage()    {}
