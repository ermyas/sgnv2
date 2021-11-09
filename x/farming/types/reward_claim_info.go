package types

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
)

// AddSig adds a signature to a reward claim details
func (d *RewardClaimDetails) AddSig(sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(d.Signatures, d.RewardProtoBytes, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	d.Signatures = sigs
	return nil
}

func (r RewardClaimInfo) LogStr() string {
	res := fmt.Sprintf("recipient:%s last_claim_time:%s reward_claim_details_list:", r.GetRecipient(), r.GetLastClaimTime().UTC())
	for _, detail := range r.GetRewardClaimDetailsList() {
		res += fmt.Sprintf(" <chain_id:%d cumulative_amount:%s> ", detail.GetChainId(), detail.GetCumulativeRewardAmounts())
	}
	return res
}
