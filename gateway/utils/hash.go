package utils

import (
	"fmt"

	"github.com/celer-network/sgn-v2/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func GenWithdrawId(chainId, seqnum uint64, addr string, tokenAddr string, amt string) common.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"uint64", "uint64", "address", "address", "uint256"},
		[]interface{}{
			fmt.Sprintf("%d", chainId),
			fmt.Sprintf("%d", seqnum),
			addr,
			tokenAddr,
			amt,
		},
	)
	return common.Bytes2Hash(hash)
}