package onchain

import (
	"context"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func GetCbrLog(chainid uint64, txhash, evname string) (*ethtypes.Log, error) {
	logmsg := "could not cbr elog:"
	chain, ok := Chains[chainid]
	if !ok {
		return nil, fmt.Errorf(logmsg+"chain not found for chain id %s", chainid)
	}
	txReceipt, err := chain.Client.TransactionReceipt(context.Background(), eth.Hex2Hash(txhash))
	if err != nil {
		return nil, fmt.Errorf(logmsg+"failed to get transaction receipt, err %v", err)
	}
	elog := eth.FindMatchCbrEvent(evname, chain.contract.Address, txReceipt.Logs)
	if elog == nil {
		log.Warnf("no match event found in tx:%s", txhash)
		return nil, fmt.Errorf("no match event found in tx: %s", txhash)
	}
	return elog, nil
}
