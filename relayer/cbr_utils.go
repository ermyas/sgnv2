package relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"time"

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

func validateSigQuorum(sortedSigs []*cbrtypes.AddrSig, curss currentSigners) (pass bool, sigsBytes [][]byte) {
	if len(curss.addrs) == 0 {
		return false, nil
	}
	totalPower := big.NewInt(0)
	signerPowers := make(map[eth.Addr]*big.Int)
	for i, power := range curss.powers {
		totalPower.Add(totalPower, power)
		signerPowers[curss.addrs[i]] = power
	}
	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))

	var filteredSigs []*cbrtypes.AddrSig
	for _, sig := range sortedSigs {
		if _, ok := signerPowers[eth.Bytes2Addr(sig.GetAddr())]; ok {
			filteredSigs = append(filteredSigs, sig)
		}
	}
	// sort signer by power
	var quorumSigners []*cbrtypes.AddrSig
	sort.Slice(filteredSigs, func(i, j int) bool {
		return signerPowers[eth.Bytes2Addr(filteredSigs[i].GetAddr())].Cmp(signerPowers[eth.Bytes2Addr(filteredSigs[j].GetAddr())]) > 0
	})

	signedPower := big.NewInt(0)
	for _, s := range filteredSigs {
		if power, ok := signerPowers[eth.Bytes2Addr(s.Addr)]; ok {
			signedPower.Add(signedPower, power)
			quorumSigners = append(quorumSigners, s)
			if signedPower.Cmp(quorumStake) > 0 {
				sort.Slice(quorumSigners, func(i, j int) bool {
					return bytes.Compare(eth.Pad20Bytes(quorumSigners[i].Addr), eth.Pad20Bytes(quorumSigners[j].Addr)) == -1
				})
				for _, signer := range quorumSigners {
					sigsBytes = append(sigsBytes, signer.Sig)
				}
				return true, sigsBytes
			}
			delete(signerPowers, eth.Bytes2Addr(s.Addr))
		}
	}

	return false, nil
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
		log.Errorln("error when get transfer by seq num:", err)
	}
	if found {
		dbErr := dal.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_REFUNDED))
		if dbErr != nil {
			log.Errorln("db when UpdateTransferStatus to TRANSFER_REFUNDED err:", dbErr)
		}
	} else {
		dbErr := dal.UpdateLPStatusForWithdraw(chainId, seqNum, uint64(cbrtypes.WithdrawStatus_WD_COMPLETED), usrAddr)
		if dbErr != nil {
			log.Errorln("db when UpdateLPStatus to WD_COMPLETED err:", dbErr)
		}
	}
}
