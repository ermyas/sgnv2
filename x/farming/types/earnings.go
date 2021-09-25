package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewEarnings creates a new instance of Earnings
func NewEarnings(targetBlockHeight int64, stakedAmount sdk.DecCoin, rewardAmount sdk.DecCoins) Earnings {
	return Earnings{
		TargetBlockHeight: targetBlockHeight,
		StakedAmount:      stakedAmount,
		RewardAmount:      rewardAmount,
	}
}
