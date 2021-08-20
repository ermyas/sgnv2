package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

const (
	// module name
	ModuleName = "validator"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	AttributeKeyEthAddress = "eth_address"

	ActionInitiateWithdraw = "initate_withdraw"
)

var (
	SyncerKey              = []byte{0x01} // key for syncer
	DelegatorKeyPrefix     = []byte{0x03} // Key prefix for delegator
	ValidatorKeyPrefix     = []byte{0x04} // Key prefix for validator
	RewardKeyPrefix        = []byte{0x05} // Key prefix for reward
	RewardEpochKey         = []byte{0x06} // Key for reward epoch
	PendingRewardKeyPrefix = []byte{0x07} // Key for pending reward
)

// get delegators key from validator address
func GetDelegatorsKey(validatorAddr string) []byte {
	return append(DelegatorKeyPrefix, contracts.Hex2Addr(validatorAddr).Bytes()...)
}

// get delegator key from validator address and delegator address
func GetDelegatorKey(validatorAddr, delegatorAddr string) []byte {
	return append(GetDelegatorsKey(validatorAddr), contracts.Hex2Addr(delegatorAddr).Bytes()...)
}

// get validator key from validatorAddr
func GetValidatorKey(validatorAddr string) []byte {
	return append(ValidatorKeyPrefix, contracts.Hex2Addr(validatorAddr).Bytes()...)
}

// get reward key from ethAddr
func GetRewardKey(ethAddr string) []byte {
	return append(RewardKeyPrefix, contracts.Hex2Addr(ethAddr).Bytes()...)
}

func GetPendingRewardKey(validatorAddr string) []byte {
	return append(PendingRewardKeyPrefix, contracts.Hex2Addr(validatorAddr).Bytes()...)
}
