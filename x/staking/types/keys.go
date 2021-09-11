package types

import (
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// module name
	ModuleName = "staking"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the validator module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName
)

var (
	ValidatorKey           = []byte{0x11} // prefix for each key to a validator
	ValidatorBySgnAddrKey  = []byte{0x12} // prefix for each key to a validator index, by accAddress
	ValidatorByConsAddrKey = []byte{0x13} // prefix for each key to a validator index, by pubkey

	ValidatorPowerKey       = []byte{0x21}
	ValidatorPowerUpdateKey = []byte{0x22}

	ValidatorTransactorsKey = []byte{0x31}

	DelegatorKey = []byte{0x41} // key prefix for delegator

	SyncerKey = []byte{0x51} // key for syncer
)

func GetValidatorPowerKey(ethAddr string) []byte {
	return append(ValidatorPowerKey, eth.Hex2Addr(ethAddr).Bytes()...)
}

func GetValidatorPowerUpdateKey(ethAddr string) []byte {
	return append(ValidatorPowerUpdateKey, eth.Hex2Addr(ethAddr).Bytes()...)
}

func GetValidatorKey(ethAddr string) []byte {
	return append(ValidatorKey, eth.Hex2Addr(ethAddr).Bytes()...)
}

func GetValidatorBySgnAddrKey(addr sdk.AccAddress) []byte {
	return append(ValidatorBySgnAddrKey, address.MustLengthPrefix(addr)...)
}

func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorByConsAddrKey, address.MustLengthPrefix(addr)...)
}

func GetValidatorTransactorsKey(ethAddr string) []byte {
	return append(ValidatorTransactorsKey, eth.Hex2Addr(ethAddr).Bytes()...)
}

// get delegators key from validator address
func GetDelegatorsKey(valAddr string) []byte {
	return append(DelegatorKey, eth.Hex2Addr(valAddr).Bytes()...)
}

// get delegator key from validator address and delegator address
func GetDelegatorKey(valAddr, delAddr string) []byte {
	return append(GetDelegatorsKey(valAddr), eth.Hex2Addr(delAddr).Bytes()...)
}

func AddrFromValidatorKey(key []byte) []byte {
	return key[1:] // remove prefix
}
