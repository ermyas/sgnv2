package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewRewardTokenInfo creates a new instance of RewardTokenInfo
func NewRewardTokenInfo(
	remainingAmount sdk.DecCoin, rewardStartBlockHeight int64, rewardAmountPerBlock sdk.Dec,
) RewardTokenInfo {
	return RewardTokenInfo{
		RemainingAmount:        remainingAmount,
		RewardStartBlockHeight: rewardStartBlockHeight,
		RewardAmountPerBlock:   rewardAmountPerBlock,
	}
}

// RewardTokenInfos is a collection of RewardTokenInfo
type RewardTokenInfos []RewardTokenInfo

// NewRewardTokenInfo creates a new instance of RewardTokenInfo
func NewRewardTokenInfos(rewardTokenInfos ...RewardTokenInfo) RewardTokenInfos {
	return rewardTokenInfos
}

// String returns a human readable string representation of RewardTokenInfos
func (infos RewardTokenInfos) String() (out string) {
	for _, info := range infos {
		out += info.String() + "\n"
	}
	return strings.TrimSpace(out)
}
