package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) setCurss(curssBytes []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.setSigners(curssBytes)
}

func (c *CbrOneChain) getCurss() *sortedSigners {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.curss
}

// each event's key is name-blkNum-index, value is json marshaled elog
func (c *CbrOneChain) saveEvent(name string, elog ethtypes.Log) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key := fmt.Sprintf("%s-%d-%d", name, elog.BlockNumber, elog.Index)
	val, _ := json.Marshal(elog)
	return c.db.Set([]byte(key), val)
}

func (c *CbrOneChain) delEvent(name string, blknum, idx uint64) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.db.Delete([]byte(fmt.Sprintf("%s-%d-%d", name, blknum, idx)))
}

func (c *CbrOneChain) getTokenFromDB(tokenAddr string) (*webapi.TokenInfo, uint64, bool) {
	newContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainId, err := c.ChainID(newContext)
	if err != nil {
		log.Errorln("get chain id err:", err)
		return nil, 0, false
	}
	token, found, err := dal.GetTokenByAddr(tokenAddr, chainId.Uint64())
	if err != nil || !found {
		return nil, 0, false
	}
	return token, chainId.Uint64(), true
}

func GatewayOnSend(transferId string) error {
	if dal.DB == nil {
		return nil
	}
	return dal.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE))
}

func GatewayOnRelay(transferId, txHash, dstTransferId, amt string) error {
	if dal.DB == nil {
		return nil
	}
	return dal.TransferCompleted(transferId, txHash, dstTransferId, amt)
}

func GatewayOnLiqAdd(lpAddr, token, tokenAddr, amt, txHash string, chainId uint64, seqNum uint64) error {
	if dal.DB == nil {
		return nil
	}
	status := cbrtypes.LPHistoryStatus_LP_WAITING_FOR_SGN
	lpType := webapi.LPType_LP_TYPE_ADD
	return dal.UpsertLP(lpAddr, token, tokenAddr, amt, txHash, chainId, uint64(status), uint64(lpType), seqNum)
}

func GatewayOnLiqWithdraw(seqNum uint64) {
	if dal.DB == nil {
		return
	}
	transferId, found, err := dal.GetTransferByRefundSeqNum(seqNum)
	if err != nil {
		log.Errorln("error when get transfer by seq num:", err)
	}
	if found {
		dbErr := dal.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_REFUNDED))
		if dbErr != nil {
			log.Errorln("db when UpdateTransferStatus to TRANSFER_REFUNDED err:", dbErr)
		}
	} else {
		dbErr := dal.UpdateLPStatusForWithdraw(seqNum, uint64(cbrtypes.LPHistoryStatus_LP_COMPLETED))
		if dbErr != nil {
			log.Errorln("db when UpdateLPStatus to LP_COMPLETED err:", dbErr)
		}
	}
}
