package types

import (
	"encoding/binary"

	"github.com/celer-network/sgn-v2/eth"
	ec "github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the name of the module
	ModuleName = "farming"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// TStoreKey is the string transient store representation
	TStoreKey = "transient_" + ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// RewardModuleAccountName is the name of the module account to store all minted reward tokens to be distributed
	RewardModuleAccountName = "reward_module_account"

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName

	// Byte length of period occupied
	PeriodByteArrayLength = 8
)

var (
	FarmingPoolPrefix           = []byte{0x01}
	PoolToAddressPrefix         = []byte{0x02}
	AddressToPoolPrefix         = []byte{0x03}
	PoolHistoricalRewardsPrefix = []byte{0x04}
	PoolCurrentRewardsPrefix    = []byte{0x05}

	ERC20TokenPrefix = []byte{0x06}
)

const (
	poolNameFromStakeInfoKeyIndex = ec.AddressLength + 1
)

func GetFarmingPoolKey(poolName string) []byte {
	return append(FarmingPoolPrefix, []byte(poolName)...)
}

func GetAddressInFarmingPoolKey(poolName string, addr eth.Addr) []byte {
	return append(PoolToAddressPrefix, append([]byte(poolName), addr.Bytes()...)...)
}

func GetStakeInfoKey(addr eth.Addr, poolName string) []byte {
	return append(AddressToPoolPrefix, append(addr.Bytes(), []byte(poolName)...)...)
}

// SplitPoolNameFromStakeInfoKey splits the pool name out from a StakeInfoKey
func SplitPoolNameFromStakeInfoKey(stakeInfoKey []byte) string {
	return string(stakeInfoKey[poolNameFromStakeInfoKeyIndex:])
}

// GetPoolHistoricalRewardsKey gets the key for a pool's historical reward
func GetPoolHistoricalRewardsKey(poolName string, period uint64) []byte {
	b := make([]byte, PeriodByteArrayLength)
	binary.LittleEndian.PutUint64(b, period)
	return append(PoolHistoricalRewardsPrefix, append([]byte(poolName), b...)...)
}

// GetPoolHistoricalRewardsPrefix gets the prefix key with pool name for a pool's historical rewards
func GetPoolHistoricalRewardsPrefix(poolName string) []byte {
	return append(PoolHistoricalRewardsPrefix, []byte(poolName)...)
}

// GetPoolCurrentRewardsKey gets the key for a pool's current period reward
func GetPoolCurrentRewardsKey(poolName string) []byte {
	return append(PoolCurrentRewardsPrefix, []byte(poolName)...)
}

// GetERC20TokenKey gets the key for an ERC-20 token
func GetERC20TokenKey(chainId uint64, symbol string) []byte {
	chainIdBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(ERC20TokenPrefix, append(chainIdBytes, []byte(symbol)...)...)
}
