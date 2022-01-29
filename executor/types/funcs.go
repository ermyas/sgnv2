package types

import (
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ExecuteRefund func(wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error

type RefundTxFunc func(
	opts *bind.TransactOpts,
	wdOnchain []byte,
	sortedSigs [][]byte,
	signers []eth.Addr,
	powers []*big.Int) (*ethtypes.Transaction, error)
