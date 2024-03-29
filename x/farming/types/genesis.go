package types

import (
	"fmt"

	commontypes "github.com/celer-network/sgn-v2/common/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params,
	pools FarmingPools,
	stakeInfos []StakeInfo,
	historicalRewards []PoolHistoricalRewardsRecord,
	currentRewards []PoolCurrentRewardsRecord,
	rewardContracts []commontypes.ContractInfo,
) *GenesisState {
	return &GenesisState{
		Params:                params,
		Pools:                 pools,
		StakeInfos:            stakeInfos,
		PoolHistoricalRewards: historicalRewards,
		PoolCurrentRewards:    currentRewards,
		RewardContracts:       rewardContracts,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:                DefaultParams(),
		Pools:                 FarmingPools{},
		StakeInfos:            []StakeInfo{},
		PoolHistoricalRewards: []PoolHistoricalRewardsRecord{},
		PoolCurrentRewards:    []PoolCurrentRewardsRecord{},
		RewardContracts:       []commontypes.ContractInfo{},
	}
}

// ValidateGenesis validates the farming genesis parameters
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

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
