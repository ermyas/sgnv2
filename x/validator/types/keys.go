package types

import (
	"github.com/celer-network/sgn-v2/contracts"
)

const (
	// module name
	ModuleName = "validator"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the validator module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName
)

var (
	ValidatorKey = []byte{0x01} // key prefix for validator
	DelegatorKey = []byte{0x11} // key prefix for delegator
	SyncerKey    = []byte{0x21} // key for syncer
)

// get delegators key from validator address
func GetDelegatorsKey(valAddr string) []byte {
	return append(DelegatorKey, contracts.Hex2Addr(valAddr).Bytes()...)
}

// get delegator key from validator address and delegator address
func GetDelegatorKey(valAddr, delAddr string) []byte {
	return append(GetDelegatorsKey(valAddr), contracts.Hex2Addr(delAddr).Bytes()...)
}

// get validator key from valAddr
func GetValidatorKey(ethAddr string) []byte {
	return append(ValidatorKey, contracts.Hex2Addr(ethAddr).Bytes()...)
}
