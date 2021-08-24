package types

import (
	"encoding/binary"
)

const (
	// module name
	ModuleName = "sync"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the sync module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName
)

var (
	PendingUpdateKey   = []byte{0x01}
	PendingUpdateIdKey = []byte{0x11}
)

// ChangeKey gets a specific change from the store
func GetPendingUpdateKey(updateId uint64) []byte {
	return append(PendingUpdateKey, GetPendingUpdateIdBytes(updateId)...)
}

// GetChangeIDFromBytes returns changeID in uint64 format from a byte array
func GetPendingUpdateIdFromBytes(bz []byte) (updateId uint64) {
	return binary.BigEndian.Uint64(bz)
}

func GetPendingUpdateIdBytes(changeID uint64) (idBytes []byte) {
	idBytes = make([]byte, 8)
	binary.BigEndian.PutUint64(idBytes, changeID)
	return
}
