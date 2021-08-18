package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Syncer struct {
	ValidatorIdx  uint           `json:"validator_idx"`
	ValidatorAddr sdk.AccAddress `json:"validator_addr"`
}

func NewSyncer(validatorIdx uint, validatorAddr sdk.AccAddress) Syncer {
	return Syncer{
		ValidatorIdx:  validatorIdx,
		ValidatorAddr: validatorAddr,
	}
}

// implement fmt.Stringer
/* func (r Syncer) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ValidatorIdx: %d, ValidatorAddr: %x`, r.ValidatorIdx, r.ValidatorAddr))
} */

func (m Syncer) Reset()         { m = Syncer{} }
func (m Syncer) String() string { return proto.CompactTextString(m) }
func (Syncer) ProtoMessage()    {}
