package dal

import (
	"fmt"

	"github.com/celer-network/sgn-v2/gateway/webapi"
)

// wrapper for other package usage, out of gateway
// for query api, if DB is nil, will a common err
// for update api, will return error==nil if DB is nil

// UpdateTransferStatus update api
func UpdateTransferStatus(transferId string, status uint64, txHash string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateTransferRefundStatus(transferId, status, txHash)
	}
}

// UpsertTransferOnSend update api
func UpsertTransferOnSend(transferId, usrAddr, tokenSymbol, amt, sendTxHash string, srcChainId, dsChainId uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpsertTransferOnSend(transferId, usrAddr, tokenSymbol, amt, sendTxHash, srcChainId, dsChainId)
	}
}

// TransferCompleted update api
func TransferCompleted(transferId, txHash, dstTransferId, amt string) error {
	if DB == nil {
		return nil
	} else {
		return DB.TransferCompleted(transferId, txHash, dstTransferId, amt)
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

// UpsertLPForLiqAdd update api
func UpsertLPForLiqAdd(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpsertLPWithTx(usrAddr, tokenSymbol, tokenAddr, amt, txHash, chainId, status, lpType, seqNum)
	}
}

// GetTransferByRefundSeqNum query api
func GetTransferByRefundSeqNum(chainId, seqNum uint64, usrAddr string) (string, bool, error) {
	if DB == nil {
		return "", false, noDBErrorForQuery()
	} else {
		return DB.GetTransferByRefundSeqNum(chainId, seqNum, usrAddr)
	}
}

// UpdateLPStatusForWithdraw update api
func UpdateLPStatusForWithdraw(chainId, seqNum, status uint64, addr, txHash string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateLPStatusForWithdrawWithTx(chainId, seqNum, status, addr, txHash)
	}
}

// common error
func noDBErrorForQuery() error {
	return fmt.Errorf("no gateway DB support")
}
