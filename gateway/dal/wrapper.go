package dal

import (
	"fmt"
	"github.com/celer-network/sgn-v2/gateway/webapi"
)

// wrapper for other package usage, out of gateway
// for query api, if DB is nil, will a common err
// for update api, will return error==nil if DB is nil

// UpdateTransferStatus update api
func UpdateTransferStatus(transferId string, status uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateTransferStatus(transferId, status)
	}
}

// TransferCompleted update api
func TransferCompleted(transferId, txHash string) error {
	if DB == nil {
		return nil
	} else {
		return DB.TransferCompleted(transferId, txHash)
	}
}

// GetTokenByAddr query api
func GetTokenByAddr(addr string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	if DB == nil {
		return nil, false, noDBErrorForQuery()
	} else {
		return DB.GetTokenByAddr(addr, chainId)
	}
}

// UpsertLP update api
func UpsertLP(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpsertLP(usrAddr, tokenSymbol, tokenAddr, amt, txHash, chainId, status, lpType, seqNum)
	}
}

// GetTransferBySeqNum query api
func GetTransferBySeqNum(seqNum uint64) (string, bool, error) {
	if DB == nil {
		return "", false, noDBErrorForQuery()
	} else {
		return DB.GetTransferBySeqNum(seqNum)
	}
}

// UpdateLPStatus update api
func UpdateLPStatus(seqNum, status uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateLPStatus(seqNum, status)
	}
}

// common error
func noDBErrorForQuery() error {
	return fmt.Errorf("no gateway DB support")
}
