package main

import (
	"testing"

	nftbr "github.com/celer-network/sgn-v2/tools/nft-bridge"
)

func TestAddrFmt(t *testing.T) {
	var adr nftbr.Addr
	t.Errorf("%s\n%x", adr, adr)
}
