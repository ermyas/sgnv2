package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// NewPoolHistoricalRewards creates a new instance of PoolHistoricalRewards
func NewPoolHistoricalRewards(cumulativeRewardRatio sdk.DecCoins, referenceCount uint32) PoolHistoricalRewards {
	return PoolHistoricalRewards{
		CumulativeRewardRatio: cumulativeRewardRatio,
		ReferenceCount:        referenceCount,
	}
}

// NewPoolCurrentRewards creates a new instance of PoolCurrentRewards
func NewPoolCurrentRewards(startBlockHeight int64, period uint64, token sdk.DecCoins) PoolCurrentRewards {
	return PoolCurrentRewards{
		StartBlockHeight: startBlockHeight,
		Period:           period,
		Rewards:          token,
	}
}
