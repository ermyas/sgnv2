package types

import (
	"encoding/binary"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "distribution"

	// StoreKey is the store key string for distribution
	StoreKey = ModuleName

	// RouterKey is the message route for distribution
	RouterKey = ModuleName

	// QuerierRoute is the querier route for distribution
	QuerierRoute = ModuleName

	StakingRewardDenom = common.CelrDenom + "/reward"
)

// Keys for distribution store
// Items are stored with the following key: values
//
// - 0x00<proposalID_Bytes>: FeePool
//
// - 0x01: sdk.ConsAddress
//
// - 0x02<valAddrLen (1 Byte)><valAddr_Bytes>: ValidatorOutstandingRewards
//
// - 0x03<delAddrLen (1 Byte)><delAddr_Bytes>: eth.Addr
//
// - 0x04<valAddrLen (1 Byte)><valAddr_Bytes><delAddrLen (1 Byte)><delAddr_Bytes>: DelegatorStartingInfo
//
// - 0x05<valAddrLen (1 Byte)><valAddr_Bytes><period_Bytes>: ValidatorHistoricalRewards
//
// - 0x06<valAddrLen (1 Byte)><valAddr_Bytes>: ValidatorCurrentRewards
//
// - 0x07<valAddrLen (1 Byte)><valAddr_Bytes>: ValidatorAccumulatedCommission
//
// - 0x08<delAddr_Bytes>: StakingRewardClaimInfo
var (
	FeePoolKey                        = []byte{0x00} // key for global distribution state
	ProposerKey                       = []byte{0x01} // key for the proposer operator address
	ValidatorOutstandingRewardsPrefix = []byte{0x02} // key for outstanding rewards

	DelegatorWithdrawAddrPrefix          = []byte{0x03} // key for delegator withdraw address
	DelegatorStartingInfoPrefix          = []byte{0x04} // key for delegator starting info
	ValidatorHistoricalRewardsPrefix     = []byte{0x05} // key for historical validators rewards / stake
	ValidatorCurrentRewardsPrefix        = []byte{0x06} // key for current validator rewards
	ValidatorAccumulatedCommissionPrefix = []byte{0x07} // key for accumulated validator commission

	StakingRewardClaimInfoPrefix = []byte{0x08} // key for delegator staking reward claim info
)

// GetValidatorOutstandingRewardsAddress creates an address from a validator's outstanding rewards key.
func GetValidatorOutstandingRewardsAddress(key []byte) (valAddr eth.Addr) {
	// key is in the format:
	// 0x02<valAddrLen (1 Byte)><valAddr_Bytes>

	// Remove prefix and address length.
	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return eth.Bytes2Addr(addr)
}

// GetDelegatorWithdrawInfoAddress creates an address from a delegator's withdraw info key.
func GetDelegatorWithdrawInfoAddress(key []byte) (delAddr eth.Addr) {
	// key is in the format:
	// 0x03<delAddrLen (1 Byte)><delAddr_Bytes>

	// Remove prefix and address length.
	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return eth.Bytes2Addr(addr)
}

// GetDelegatorStartingInfoAddresses creates the addresses from a delegator starting info key.
func GetDelegatorStartingInfoAddresses(key []byte) (valAddr eth.Addr, delAddr eth.Addr) {
	// key is in the format:
	// 0x04<valAddrLen (1 Byte)><valAddr_Bytes><accAddrLen (1 Byte)><accAddr_Bytes>
	valAddrLen := int(key[1])
	valAddr = eth.Bytes2Addr(key[2 : 2+valAddrLen])
	delAddrLen := int(key[2+valAddrLen])
	delAddr = eth.Bytes2Addr(key[3+valAddrLen:])
	if len(delAddr.Bytes()) != delAddrLen {
		panic("unexpected key length")
	}

	return
}

// GetValidatorHistoricalRewardsAddressPeriod creates the address & period from a validator's historical rewards key.
func GetValidatorHistoricalRewardsAddressPeriod(key []byte) (valAddr eth.Addr, period uint64) {
	// key is in the format:
	// 0x05<valAddrLen (1 Byte)><valAddr_Bytes><period_Bytes>
	valAddrLen := int(key[1])
	valAddr = eth.Bytes2Addr(key[2 : 2+valAddrLen])
	b := key[2+valAddrLen:]
	if len(b) != 8 {
		panic("unexpected key length")
	}
	period = binary.LittleEndian.Uint64(b)
	return
}

// GetValidatorCurrentRewardsAddress creates the address from a validator's current rewards key.
func GetValidatorCurrentRewardsAddress(key []byte) (valAddr eth.Addr) {
	// key is in the format:
	// 0x06<valAddrLen (1 Byte)><valAddr_Bytes>: ValidatorCurrentRewards

	// Remove prefix and address length.
	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return eth.Bytes2Addr(addr)
}

// GetValidatorAccumulatedCommissionAddress creates the address from a validator's accumulated commission key.
func GetValidatorAccumulatedCommissionAddress(key []byte) (valAddr eth.Addr) {
	// key is in the format:
	// 0x07<valAddrLen (1 Byte)><valAddr_Bytes>: ValidatorCurrentRewards

	// Remove prefix and address length.
	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return eth.Bytes2Addr(addr)
}

// GetValidatorSlashEventAddressHeight creates the height from a validator's slash event key.
func GetValidatorSlashEventAddressHeight(key []byte) (valAddr eth.Addr, height uint64) {
	// key is in the format:
	// 0x08<valAddrLen (1 Byte)><valAddr_Bytes><height>: ValidatorSlashEvent
	valAddrLen := int(key[1])
	valAddr = eth.Bytes2Addr(key[2 : 2+valAddrLen])
	startB := 2 + valAddrLen
	b := key[startB : startB+8] // the next 8 bytes represent the height
	height = binary.BigEndian.Uint64(b)
	return
}

// GetValidatorOutstandingRewardsKey creates the outstanding rewards key for a validator.
func GetValidatorOutstandingRewardsKey(valAddr eth.Addr) []byte {
	return append(ValidatorOutstandingRewardsPrefix, address.MustLengthPrefix(valAddr.Bytes())...)
}

// GetDelegatorWithdrawAddrKey creates the key for a delegator's withdraw addr.
func GetDelegatorWithdrawAddrKey(delAddr eth.Addr) []byte {
	return append(DelegatorWithdrawAddrPrefix, address.MustLengthPrefix(delAddr.Bytes())...)
}

// GetDelegatorStartingInfoKey creates the key for a delegator's starting info.
func GetDelegatorStartingInfoKey(v eth.Addr, d eth.Addr) []byte {
	return append(append(DelegatorStartingInfoPrefix, address.MustLengthPrefix(v.Bytes())...), address.MustLengthPrefix(d.Bytes())...)
}

// GetValidatorHistoricalRewardsPrefix creates the prefix key for a validator's historical rewards.
func GetValidatorHistoricalRewardsPrefix(v eth.Addr) []byte {
	return append(ValidatorHistoricalRewardsPrefix, address.MustLengthPrefix(v.Bytes())...)
}

// GetValidatorHistoricalRewardsKey creates the key for a validator's historical rewards.
func GetValidatorHistoricalRewardsKey(v eth.Addr, k uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, k)
	return append(append(ValidatorHistoricalRewardsPrefix, address.MustLengthPrefix(v.Bytes())...), b...)
}

// GetValidatorCurrentRewardsKey creates the key for a validator's current rewards.
func GetValidatorCurrentRewardsKey(v eth.Addr) []byte {
	return append(ValidatorCurrentRewardsPrefix, address.MustLengthPrefix(v.Bytes())...)
}

// GetValidatorAccumulatedCommissionKey creates the key for a validator's current commission.
func GetValidatorAccumulatedCommissionKey(v eth.Addr) []byte {
	return append(ValidatorAccumulatedCommissionPrefix, address.MustLengthPrefix(v.Bytes())...)
}

// GetStakingRewardClaimInfoKey gets staking reward claim info key from an Ethereum address
func GetStakingRewardClaimInfoKey(addr eth.Addr) []byte {
	return append(StakingRewardClaimInfoPrefix, addr.Bytes()...)
}
