package types

import (
	"time"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Reward struct {
	Recipient        string       `json:"recipient"`
	Reward           sdk.Int      `json:"reward"`
	RewardProtoBytes []byte       `json:"reward_proto_bytes"` // proto msg for reward snapshot from latest intendWithdraw
	LastClaimTime    time.Time    `json:"last_claim_time"`    // last time the user triggers claim
	Sigs             []common.Sig `json:"sigs"`
}

func NewReward(recipient string) Reward {
	return Reward{
		Recipient: contracts.FormatAddrHex(recipient),
		Reward:    sdk.ZeroInt(),
	}
}

func (r Reward) Reset()         { r = Reward{} }
func (r Reward) String() string { return proto.CompactTextString(r) }
func (Reward) ProtoMessage()    {}
