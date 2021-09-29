package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
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

func validateCbrSigs(sigs []*cbrtypes.AddrSig, curss *cbrtypes.SortedSigners) bool {
	if len(curss.GetSigners()) == 0 {
		return false
	}
	totalPower := big.NewInt(0)
	cursMap := make(map[eth.Addr]*cbrtypes.AddrAmt)
	for _, s := range curss.GetSigners() {
		power := big.NewInt(0).SetBytes(s.Amt)
		totalPower.Add(totalPower, power)
		cursMap[eth.Bytes2Addr(s.Addr)] = s
	}

	signedPower := big.NewInt(0)
	i := 0
	for _, s := range sigs {
		if addrAmt, ok := cursMap[eth.Bytes2Addr(s.Addr)]; ok {
			power := big.NewInt(0).SetBytes(addrAmt.Amt)
			signedPower.Add(signedPower, power)
			sigs[i] = s
			i++
		}
	}
	// truncate sigs not in the current signers set
	for j := i; j < len(sigs); j++ {
		sigs[j] = nil
	}
	sigs = sigs[:i]

	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))

	if signedPower.Cmp(quorumStake) > 0 {
		return true
	}

	return false
}

func GatewayOnSend(transferId string) error {
	return dal.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE))
}

func GatewayOnRelay(transferId, txHash string) error {
	return dal.TransferCompleted(transferId, txHash)
}

func GatewayOnLiqAdd(lpAddr, token, tokenAddr, amt, txHash string, chainId uint64, seqNum uint64) error {
	status := cbrtypes.LPHistoryStatus_LP_WAITING_FOR_SGN
	lpType := webapi.LPType_LP_TYPE_ADD
	return dal.UpsertLP(lpAddr, token, tokenAddr, amt, txHash, chainId, uint64(status), uint64(lpType), seqNum)
}

func GatewayOnLiqWithdraw(seqNum uint64) {
	transferId, found, err := dal.GetTransferBySeqNum(seqNum)
	if err != nil {
		log.Errorln("error when get transfer by seq num:", err)
	}
	if found {
		dbErr := dal.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_REFUNDED))
		if dbErr != nil {
			log.Errorln("db when UpdateTransferStatus to TRANSFER_REFUNDED err:", err)
		}
	} else {
		dbErr := dal.UpdateLPStatus(seqNum, uint64(cbrtypes.LPHistoryStatus_LP_COMPLETED))
		if dbErr != nil {
			log.Errorln("db when UpdateLPStatus to LP_COMPLETED err:", err)
		}
	}
}
