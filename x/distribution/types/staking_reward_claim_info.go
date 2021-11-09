package types

import (
	fmt "fmt"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
)

// AddSig adds a signature to a staking reward claim info
func (i *StakingRewardClaimInfo) AddSig(sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(i.Signatures, i.RewardProtoBytes, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	i.Signatures = sigs
	return nil
}

func (r StakingRewardClaimInfo) LogStr() string {
	res := fmt.Sprintf("recipient:%s last_claim_time:%s cumulative_amount:%s",
		r.GetRecipient(), r.GetLastClaimTime().UTC(), r.GetCumulativeRewardAmount())
	return res
}
