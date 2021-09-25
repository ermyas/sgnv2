package types

import (
	"fmt"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	pools FarmingPools,
	stakeInfos []StakeInfo,
	historicalRewards []PoolHistoricalRewardsRecord,
	currentRewards []PoolCurrentRewardsRecord,
) *GenesisState {
	return &GenesisState{
		Pools:                 pools,
		StakeInfos:            stakeInfos,
		PoolHistoricalRewards: historicalRewards,
		PoolCurrentRewards:    currentRewards,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Pools:                 FarmingPools{},
		StakeInfos:            []StakeInfo{},
		PoolHistoricalRewards: []PoolHistoricalRewardsRecord{},
		PoolCurrentRewards:    []PoolCurrentRewardsRecord{},
	}
}

// ValidateGenesis validates the farm genesis parameters
func ValidateGenesis(data GenesisState) error {
	if len(data.Pools) != len(data.PoolCurrentRewards) {
		return fmt.Errorf("count of pools(%d) is not equal to that of current rewards(%d)",
			len(data.Pools), len(data.PoolCurrentRewards))
	}

	var expectedReferenceCount uint16
	for _, h := range data.PoolHistoricalRewards {
		expectedReferenceCount += uint16(h.Rewards.ReferenceCount)
	}

	actualReferenceCount := len(data.StakeInfos) + len(data.PoolCurrentRewards)
	if actualReferenceCount != int(expectedReferenceCount) {
		return fmt.Errorf("actual reference count(%d) is not equal to expected reference count(%d)",
			actualReferenceCount, expectedReferenceCount)
	}
	return nil
}
