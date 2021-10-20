package types

import (
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
