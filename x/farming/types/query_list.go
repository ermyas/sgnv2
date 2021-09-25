package types

import (
	"fmt"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
)

// PoolNameList is the type alias for []string
type PoolNameList []string

// String returns a human readable string representation of PoolNameList
func (pnl PoolNameList) String() string {
	out := "Pool Name List:\n"
	for _, poolName := range pnl {
		out = fmt.Sprintf("%s  %s\n", out, poolName)
	}
	return strings.TrimSpace(out)
}

// AddrList is the type alias for []eth.Addr
type AddrList []eth.Addr

// String returns a human readable string representation of AddrList
func (al AddrList) String() string {
	out := "Address List:\n"
	for _, addr := range al {
		out = fmt.Sprintf("%s  %s\n", out, addr.String())
	}
	return strings.TrimSpace(out)
}
