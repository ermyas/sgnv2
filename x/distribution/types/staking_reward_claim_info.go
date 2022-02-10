package types

import (
	fmt "fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// AddSig adds a signature to a staking reward claim info
func (i *StakingRewardClaimInfo) AddSig(msgToSign []byte, sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(i.Signatures, msgToSign, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	i.Signatures = sigs
	return nil
}

func (i *StakingRewardClaimInfo) EncodeDataToSign(chainId uint64, contractAddr eth.Addr) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(chainId), contractAddr, "StakingReward"},
	)
	return append(domain, i.RewardProtoBytes...)
}

func (i StakingRewardClaimInfo) LogStr() string {
	res := fmt.Sprintf("recipient:%s last_claim_time:%s cumulative_amount:%s",
		i.GetRecipient(), i.GetLastClaimTime().UTC(), i.GetCumulativeRewardAmount())
	return res
}

func (i *StakingRewardClaimInfo) GetAddrSigs() []*cbrtypes.AddrSig {
	addrSigs := make([]*cbrtypes.AddrSig, 0)
	for _, sig := range i.Signatures {
		addrSigs = append(addrSigs, &cbrtypes.AddrSig{
			Addr: eth.Hex2Bytes(sig.Signer),
			Sig:  sig.SigBytes,
		})
	}

	return addrSigs
}

func (i *StakingRewardClaimInfo) SignersStr() string {
	var signers string
	for _, s := range i.Signatures {
		signers += fmt.Sprintf("%s ", s.Signer)
	}
	return fmt.Sprintf("signers:< %s>", signers)
}

func (i *StakingRewardClaimInfo) GetSortedSigsBytes() [][]byte {
	if i != nil {
		sigs := make([][]byte, len(i.Signatures))
		for index := range i.Signatures {
			sigs[index] = i.Signatures[index].SigBytes
		}
		return sigs
	}
	return nil
}
