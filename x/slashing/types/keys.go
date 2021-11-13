package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// module name
	ModuleName = "slashing"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// QuerierRoute is the querier route for gov
	QuerierRoute = ModuleName
)

var (
	SlashKeyPrefix = []byte{0x11} // Key prefix for slash
	SlashNonceKey  = []byte{0x12} // Key for slash nonce
)

// GetSlashKey gets slash key from nonce
func GetSlashKey(nonce uint64) []byte {
	return append(SlashKeyPrefix, sdk.Uint64ToBigEndian(nonce)...)
}
