package types

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// AddSig adds a signature to a reward claim details
func (d *RewardClaimDetails) AddSig(msgToSign []byte, sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(d.Signatures, msgToSign, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	d.Signatures = sigs
	return nil
}

func (d *RewardClaimDetails) EncodeDataToSign(contractAddr eth.Addr) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(d.ChainId), contractAddr, "FarmingRewards"},
	)
	return append(domain, d.RewardProtoBytes...)
}

func (r RewardClaimInfo) LogStr() string {
	res := fmt.Sprintf("recipient:%s last_claim_time:%s reward_claim_details_list:", r.GetRecipient(), r.GetLastClaimTime().UTC())
	for _, detail := range r.GetRewardClaimDetailsList() {
		res += fmt.Sprintf(" <chain_id:%d cumulative_amount:%s> ", detail.GetChainId(), detail.GetCumulativeRewardAmounts())
	}
	return res
}
