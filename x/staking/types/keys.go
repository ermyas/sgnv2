package types

import (
	"github.com/celer-network/sgn-v2/common"
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

	StakeDenom = common.CelrDenom + "/stake"

	DefaultQueryLimit = 100
)

var (
	ValidatorKey           = []byte{0x11} // prefix for each key to a validator
	ValidatorBySgnAddrKey  = []byte{0x12} // prefix for each key to a validator index, by accAddress
	ValidatorByConsAddrKey = []byte{0x13} // prefix for each key to a validator index, by pubkey

	ValidatorPowerKey       = []byte{0x21}
	ValidatorPowerUpdateKey = []byte{0x22}

	ValidatorTransactorsKey = []byte{0x31}

	DelegationKey = []byte{0x41} // key prefix for delegation

	SyncerKey = []byte{0x51} // key for syncer
)

func GetValidatorPowerKey(ethAddr eth.Addr) []byte {
	return append(ValidatorPowerKey, ethAddr.Bytes()...)
}

func GetValidatorPowerUpdateKey(ethAddr eth.Addr) []byte {
	return append(ValidatorPowerUpdateKey, ethAddr.Bytes()...)
}

func GetValidatorKey(ethAddr eth.Addr) []byte {
	return append(ValidatorKey, ethAddr.Bytes()...)
}

func GetValidatorBySgnAddrKey(addr sdk.AccAddress) []byte {
	return append(ValidatorBySgnAddrKey, address.MustLengthPrefix(addr)...)
}

func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorByConsAddrKey, address.MustLengthPrefix(addr)...)
}

func GetValidatorTransactorsKey(ethAddr eth.Addr) []byte {
	return append(ValidatorTransactorsKey, ethAddr.Bytes()...)
}

// get delegations key from delegator address
func GetDelegationsKey(delAddr eth.Addr) []byte {
	return append(DelegationKey, delAddr.Bytes()...)
}

// get delegation key from delegator address and validator address
func GetDelegationKey(delAddr eth.Addr, valAddr eth.Addr) []byte {
	return append(GetDelegationsKey(delAddr), valAddr.Bytes()...)
}

func AddrFromValidatorKey(key []byte) []byte {
	return key[1:] // remove prefix
}
