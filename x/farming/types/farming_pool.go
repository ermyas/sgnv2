package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewFarmingPool creates a new instance of FarmingPool
func NewFarmingPool(
	name string,
	stakeToken ERC20Token,
	rewardTokens []ERC20Token,
	totalStakedAmount sdk.DecCoin,
	rewardTokenInfos RewardTokenInfos,
	totalAccumulatedRewards sdk.DecCoins,
) FarmingPool {
	return FarmingPool{
		Name:                    name,
		StakeToken:              stakeToken,
		RewardTokens:            rewardTokens,
		TotalStakedAmount:       totalStakedAmount,
		RewardTokenInfos:        rewardTokenInfos,
		TotalAccumulatedRewards: totalAccumulatedRewards,
	}
}

func (fp FarmingPool) Finished() bool {
	for _, rewardTokenInfo := range fp.RewardTokenInfos {
		if rewardTokenInfo.RemainingAmount.IsPositive() {
			return false
		}
	}
	return fp.TotalStakedAmount.IsZero()
}

// FarmingPools is a collection of FarmingPool
type FarmingPools []FarmingPool

// String returns a human readable string representation of FarmingPools
func (fps FarmingPools) String() (out string) {
	for _, fp := range fps {
		out += fp.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// NewNumPools creates a new instance of NumPools
func NewNumPools(num uint64) NumPools {
	return NumPools{
		NumPools: num,
	}
}
