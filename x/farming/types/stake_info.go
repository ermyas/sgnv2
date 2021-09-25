package types

import (
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStakeInfo creates a new instance of StakeInfo
func NewStakeInfo(
	address eth.Addr,
	poolName string,
	amount sdk.DecCoin,
	startBlockHeight int64,
	referencePeriod uint64,
) StakeInfo {
	return StakeInfo{
		StakerAddress:    eth.Addr2Hex(address),
		PoolName:         poolName,
		Amount:           amount,
		StartBlockHeight: startBlockHeight,
		ReferencePeriod:  referencePeriod,
	}
}
