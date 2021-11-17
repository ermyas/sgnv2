package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) setCurss(ss []*cbrtypes.Signer) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs, c.curss.powers = cbrtypes.SignersToEthArrays(ss)
}

func (c *CbrOneChain) setCurssByEvent(e *eth.BridgeSignersUpdated) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs = make([]eth.Addr, len(e.Signers))
	c.curss.powers = make([]*big.Int, len(e.Powers))
	for i, addr := range e.Signers {
		c.curss.addrs[i] = addr
		c.curss.powers[i] = e.Powers[i]
	}
}

func (c *CbrOneChain) getCurss() currentSigners {
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

func (c *CbrOneChain) getTokenFromDB(tokenAddr string) (*webapi.TokenInfo, bool) {
	token, found, err := dal.GetTokenByAddr(tokenAddr, c.chainid)
	if err != nil || !found {
		return nil, false
	}
	return token, true
}

func GatewayOnSend(transferId, usrAddr, tokenSymbol, amt, sendTxHash string, srcChainId, dsChainId uint64) error {
	if dal.DB == nil {
		return nil
	}
	return dal.UpsertTransferOnSend(transferId, usrAddr, tokenSymbol, amt, sendTxHash, srcChainId, dsChainId)
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
	status := cbrtypes.WithdrawStatus_WD_WAITING_FOR_SGN
	lpType := webapi.LPType_LP_TYPE_ADD
	return dal.UpsertLPForLiqAdd(lpAddr, token, tokenAddr, amt, txHash, chainId, uint64(status), uint64(lpType), seqNum)
}

func GatewayOnLiqWithdraw(chainId, seqNum uint64, usrAddr string) {
	if dal.DB == nil {
		return
	}
	transferId, found, err := dal.GetTransferByRefundSeqNum(chainId, seqNum, usrAddr)
	if err != nil {
		log.Warnf("error when get transfer, usr:%s chainId:%d, seqNum:%d, err:%+v", usrAddr, chainId, seqNum, err)
	}
	if found {
		dbErr := dal.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_REFUNDED))
		if dbErr != nil {
			log.Warnf("db when UpdateTransferStatus to TRANSFER_REFUNDED, transferId:%s, err:%+v", transferId, dbErr)
		}
	} else {
		dbErr := dal.UpdateLPStatusForWithdraw(chainId, seqNum, uint64(cbrtypes.WithdrawStatus_WD_COMPLETED), usrAddr)
		if dbErr != nil {
			log.Warnf("db when UpdateLPStatus to WD_COMPLETED, transferId:%s, err:%+v", transferId, dbErr)
		}
	}
}
