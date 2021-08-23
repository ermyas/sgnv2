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
	UpdateKey   = []byte{0x01}
	UpdateIdKey = []byte{0x11}
)

// ChangeKey gets a specific change from the store
func GetUpdateKey(updateId uint64) []byte {
	idbytes := make([]byte, 8)
	binary.BigEndian.PutUint64(idbytes, updateId)
	return append(UpdateKey, idbytes...)
}
