package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

const (
	maxRelayRetry     = 15
	maxSigRetry       = 10
	maxBytesPerUpdate = 400000

	sigWaitSgnBlk    = 4
	newSyncerWaitBlk = 8
)

// sleep, check if syncer, if yes, go over cbr dbs to send tx
func (r *Relayer) doCbridgeSync(cbrMgr CbrMgr) {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalCbridge)) * time.Second
	log.Infoln("start process cbridge sync, interval:", interval)
	for {
		time.Sleep(interval)
		if !r.isSyncer() {
			continue
		}
		// find all events need to be sent out, batch into one msg
		msg := &synctypes.MsgProposeUpdates{
			Sender: r.Transactor.Key.GetAddress().String(),
		}

		var updatesBytesLen int
		for chid, onech := range cbrMgr {
			// go over each chain db events, send msg
			ret, isUpdateMsgFull := onech.pullEvents(chid, r.Transactor.CliCtx, &updatesBytesLen)
			msg.Updates = append(msg.Updates, ret...)
			if isUpdateMsgFull {
				break
			}
		}
		if len(msg.Updates) > 0 {
			// or we should call cbridge grpc here?
			r.Transactor.AddTxMsg(msg)
			log.Debugln("CbridgeEvent updates count in one msg:", len(msg.Updates))
		}

		if r.isCbrSsUpdating() {
			r.updateSigners()
		}
	}
}

// sleep, check if syncer, if yes, go over cbr dbs to send tx
func (r *Relayer) doCbridgeOnchain(cbrMgr CbrMgr) {
	for chid := range cbrMgr {
		go r.doCbridgeOnchainByChain(chid)
	}
}

func (r *Relayer) doCbridgeOnchainByChain(chid uint64) {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalCbridge)) * time.Second
	log.Infof("start process cbridge onchain, interval:%s, chainId: %d", interval, chid)
	for {
		time.Sleep(interval)
		r.processCbridgeQueue(chid)
	}
}

func (r *Relayer) processCbridgeQueue(chid uint64) {
	syncer, syncerUpdateTime := r.getSyncer()
	if !syncer {
		return
	}

	var keys, vals [][]byte
	r.cbrMgr[chid].lock.RLock()
	prefix := GetCbrChainXferPrefix(chid)
	iterator, err := r.db.Iterator(prefix, storetypes.PrefixEndBytes(prefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
		r.cbrMgr[chid].lock.RUnlock()
		return
	}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
		vals = append(vals, iterator.Value())
	}
	iterator.Close()
	r.cbrMgr[chid].lock.RUnlock()

	if len(keys) > 0 {
		log.Debugf("start process relay queue for dst chain %d queue size: %d", chid, len(keys))
	} else {
		return
	}

	newSyncer := false
	newSyncerWaitTime := time.Duration(r.cbrMgr[chid].blkInterval) * time.Second * newSyncerWaitBlk
	if syncerUpdateTime.Add(newSyncerWaitTime).After(time.Now()) {
		newSyncer = true
	}
	sigWaitTime := viper.GetDuration(common.FlagConsensusTimeoutCommit) * sigWaitSgnBlk
	for i, key := range keys {
		relay := NewRelayRequestFromBytes(vals[i])
		if relay.CreateTime.Add(sigWaitTime).After(time.Now()) {
			// wait a while to collect validator signatures
			continue
		}
		if newSyncer && relay.CreateTime.Before(syncerUpdateTime) {
			// wait for relay to be submitted by the previous syncer
			continue
		}

		err = r.dbDelete(key)
		if err != nil {
			log.Errorln("db Delete err", err)
			continue
		}

		r.submitRelay(relay)
	}
}

func (r *Relayer) submitRelay(relayRequest RelayRequest) {
	logmsg := fmt.Sprintf("Process relay srcId %x dstChain %d", relayRequest.XferId, relayRequest.DstChainId)

	relay, err := cbrcli.QueryRelay(r.Transactor.CliCtx, relayRequest.XferId)
	if err != nil {
		r.requeueRelay(relayRequest, fmt.Sprintf("%s. QueryRelay err: %s", logmsg, err), true)
		return
	}

	relayOnChain := new(cbrtypes.RelayOnChain)
	err = relayOnChain.Unmarshal(relay.Relay)
	if err != nil {
		log.Errorf("%s. Unmarshal relay.Relay err %s", logmsg, err)
		return
	}

	xferId := eth.Bytes2Hex(relayOnChain.SrcTransferId)
	resp, err := cbrcli.QueryTransferStatus(r.Transactor.CliCtx, &cbrtypes.QueryTransferStatusRequest{
		TransferId: []string{xferId},
	})
	if err != nil {
		r.requeueRelay(relayRequest, fmt.Sprintf("%s. QueryTransferStatus err: %s", logmsg, err), true)
		return
	}
	if resp.Status[xferId].SgnStatus == cbrtypes.XferStatus_SUCCESS {
		log.Infof("%s. transfer already completed, skip it", logmsg)
		return
	}

	curss := r.cbrMgr[relayOnChain.DstChainId].getCurss()
	curssList := make([]*cbrtypes.Signer, 0)
	for i, addr := range curss.addrs {
		power := curss.powers[i]
		curssList = append(curssList, &cbrtypes.Signer{
			Addr:  addr.Bytes(),
			Power: power.Bytes(),
		})
	}
	pass, sigsBytes := cbrtypes.ValidateSigQuorum(relay.SortedSigs, curssList)
	if !pass {
		r.requeueRelay(relayRequest,
			fmt.Sprintf("%s. Not have enough sigs %s, curss %s", logmsg, relay.SignersStr(), curss.String()), false)
		return
	}
	relayTransferId := relayOnChain.GetRelayOnChainTransferId()
	logmsg = fmt.Sprintf("%s dstId %x", logmsg, relayTransferId)
	existRelay, existRelayErr := r.cbrMgr[relayOnChain.DstChainId].existTransferId(relayTransferId)
	if existRelayErr != nil {
		// if fail to query, continue to send this relay, because we can not make sure whether the relay already exist.
		log.Warnln("fail to query transfer err:", existRelayErr)
	} else if existRelay {
		log.Infof("%s. dest transfer already exist on chain, skip it", logmsg)
		return
	}
	txHash, err := r.cbrMgr[relayOnChain.DstChainId].SendRelay(relay.Relay, sigsBytes, curss, relayOnChain)
	if err != nil {
		if strings.Contains(err.Error(), "transfer exists") {
			log.Infof("%s. err %s, skip it", logmsg, err)
			return
		}

		if strings.Contains(err.Error(), "Pausable: paused") ||
			strings.Contains(err.Error(), "volume exceeds cap") ||
			strings.Contains(err.Error(), "Mismatch current signers") {
			if relayRequest.RetryCount > 0 {
				relayRequest.RetryCount -= 1
			}
		}
		r.requeueRelay(relayRequest, fmt.Sprintf("%s. err %s", logmsg, err), true)
		return
	}
	log.Infof("%s. tx hash %s", logmsg, txHash)
}

func (r *Relayer) requeueRelay(relayRequest RelayRequest, logmsg string, warn bool) {
	if relayRequest.RetryCount >= maxRelayRetry {
		log.Errorf("%s. hits retry limit", logmsg)
		return
	}
	relayRequest.RetryCount += 1
	err := r.dbSet(GetCbrXferKey(relayRequest.XferId, relayRequest.DstChainId), relayRequest.MustMarshal())
	if err != nil {
		log.Errorf("%s. db Set err: %s", logmsg, err)
	}
	if warn {
		log.Warn(logmsg)
	} else {
		log.Debug(logmsg)
	}
}

// TODO: query x/cbridge to skip already processed events to avoid duplicated propose
// Note if syncer changes before EndBlock, new syncer may still propose again
// the 2nd propose shouldn't get votes because when verify, sgn nodes will find it's already processed
// even it is voted, apply will still fail because x/cbr will err
func (c *CbrOneChain) pullEvents(chid uint64, cliCtx client.Context, updatesBytesLen *int) (ret []*synctypes.ProposeUpdate, isUpdateMsgFull bool) {
	// to make it simple we use "srcChainId-destChainId-srcTokenAddr" as key, and valid as val.
	// this cache can only be used in only one pullEvents, if pull again, we should create and use a new cache.
	cbrSendValidCache := make(map[string]bool)
	// 1st loop over event names, then go over iter
	isUpdateMsgFull = false
	for _, evn := range evNames {
		var keys, vals [][]byte
		c.lock.RLock()
		iterator, err := c.db.Iterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn)))
		if err != nil {
			log.Errorln("Create db iterator err", err)
			c.lock.RUnlock()
			continue
		}
		for ; iterator.Valid(); iterator.Next() {
			keys = append(keys, iterator.Key())
			vals = append(vals, iterator.Value())
		}
		iterator.Close()
		c.lock.RUnlock()

		for i, key := range keys {
			err = c.db.Delete(key) // TODO: lock protection?
			if err != nil {
				log.Errorln("db Delete err", err)
				continue
			}

			evlog := new(ethtypes.Log)
			err := json.Unmarshal(vals[i], evlog)
			if err != nil {
				log.Errorf("failed to unmarshal onchev elog, key:%s, err:%s", string(key), err)
				continue
			}

			skip, reason := c.skipEvent(evn, evlog, cliCtx, cbrSendValidCache)
			if skip {
				log.Debugf("skip cbr event: %s, chid %d, reason: %s", string(key), c.chainid, reason)
				continue
			}

			onchev := &cbrtypes.OnChainEvent{
				Chainid: chid,
				Evtype:  evn,
				Elog:    vals[i],
			}
			data, _ := onchev.Marshal()
			update := &synctypes.ProposeUpdate{
				Type:       synctypes.DataType_CbrOnchainEvent,
				ChainId:    chid,
				ChainBlock: 0, // why do we need this in ProposeUpdate?
				Data:       data,
			}

			updateBytes, _ := proto.Marshal(update)
			*updatesBytesLen += len(updateBytes)
			if *updatesBytesLen > maxBytesPerUpdate {
				isUpdateMsgFull = true
				c.db.Set(key, vals[i]) // adds back to db
				break
			}

			ret = append(ret, update)
		}
		if isUpdateMsgFull {
			break
		}
	}
	return
}

func (c *CbrOneChain) skipEvent(evn string, evlog *ethtypes.Log, cliCtx client.Context, checkedCache map[string]bool) (skip bool, reason string) {
	switch evn {
	case cbrtypes.CbrEventSend:
		skip, reason = c.skipSyncCbrSend(evlog, cliCtx, checkedCache)
	case cbrtypes.CbrEventSignersUpdated:
		skip, reason = c.skipSyncCbrSignerUpdate(evlog, cliCtx)
	case cbrtypes.CbrEventLiqAdd:
		skip, reason = c.skipSyncCbrLiqAdd(evlog, cliCtx)
	case cbrtypes.CbrEventRelay:
		skip, reason = c.skipSyncCbrRelay(evlog, cliCtx)
	case cbrtypes.CbrEventWithdraw:
		skip, reason = c.skipSyncCbrWithdraw(evlog, cliCtx)
	case cbrtypes.CbrEventWithdrawalRequest:
		skip, reason = c.skipSyncCbrWithdrawalRequest(evlog, cliCtx)
	}

	return
}

func (c *CbrOneChain) skipSyncCbrSend(
	evlog *ethtypes.Log, cliCtx client.Context, validCache map[string]bool) (skip bool, reason string) {

	sendEv, err := c.cbrContract.ParseSend(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	// we should check cache first
	cacheKey := fmt.Sprintf("%d-%d-%x", c.chainid, sendEv.DstChainId, sendEv.Token)

	if validCache != nil {
		cacheValid, found := validCache[cacheKey]
		if found && !cacheValid {
			return true, "invalid cbr send"
		}
	}

	checkReq := &cbrtypes.CheckChainTokenValidRequest{
		SrcChainId:   c.chainid,
		DestChainId:  sendEv.DstChainId,
		SrcTokenAddr: eth.Addr2Hex(sendEv.Token),
	}
	checkResp, checkRespErr := cbrcli.QueryCheckChainTokenValid(cliCtx, checkReq)
	if checkRespErr != nil {
		// If request failed, we will not break this flow.
		// As if invalid token send event go to the apply flow, sgn will also check it and set it to refund flow.
		log.Errorf("fail to check chain token valid, sendEv:%s, err:%s", sendEv.PrettyLog(c.chainid), checkRespErr)
		// may be call sgn fail, we still send this ev to sgn and sgn to do the check again.
		return
	} else {
		// cached and can reduce some cli call
		if validCache != nil {
			validCache[cacheKey] = checkResp.GetValid()
		}
		if !checkResp.GetValid() {
			return true, "invalid cbr send"
		}
	}

	xferId := eth.Hash(sendEv.TransferId).String()
	resp, err := cbrcli.QueryTransferStatus(cliCtx, &cbrtypes.QueryTransferStatusRequest{
		TransferId: []string{xferId},
	})
	if err != nil {
		// log only, will not skip if request failed
		log.Errorf("QueryTransferStatus err: %s", err)
		return
	}
	if resp.Status[xferId].SgnStatus != cbrtypes.XferStatus_UNKNOWN {
		return true, fmt.Sprintf("xfer with xferId %s already synced", xferId)
	}

	return
}

func (c *CbrOneChain) skipSyncCbrSignerUpdate(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.cbrContract.ParseSignersUpdated(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	ssHash, err := c.cbrContract.SsHash(&bind.CallOpts{})
	if err != nil {
		log.Errorf("chain %d failed to get onchain sshash err %s", c.chainid, err)
		return
	}
	if eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(ev.Signers, ev.Powers))) != ssHash {
		return true, "not match onchain sshash, maybe outdated"
	}

	chainSigners, err := cbrcli.QueryChainSigners(cliCtx, c.chainid)
	if err == nil {
		addrs, powers := cbrtypes.SignersToEthArrays(chainSigners.SortedSigners)
		if eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(addrs, powers))) == ssHash {
			return true, "chain signers already updated"
		}
	}
	return
}

func (c *CbrOneChain) skipSyncCbrLiqAdd(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.cbrContract.ParseLiquidityAdded(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}

	resp, err := cbrcli.QueryAddLiquidityStatus(cliCtx, &cbrtypes.QueryAddLiquidityStatusRequest{
		ChainId: c.chainid,
		SeqNum:  ev.Seqnum,
	})
	if err != nil {
		// log only, will not skip if request failed
		log.Errorf("QueryAddLiquidityStatus err: %s", err)
		return
	}
	if resp.Status == cbrtypes.WithdrawStatus_WD_COMPLETED {
		return true, fmt.Sprintf("LiquidityAdded with seqNum %d on chain %d already synced", ev.Seqnum, c.chainid)
	}

	return
}

func (c *CbrOneChain) skipSyncCbrRelay(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.cbrContract.ParseRelay(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}

	xferId := eth.Hash(ev.SrcTransferId).String()
	resp, err := cbrcli.QueryTransferStatus(cliCtx, &cbrtypes.QueryTransferStatusRequest{
		TransferId: []string{xferId},
	})
	if err != nil {
		// log only, will not skip if request failed
		log.Errorf("QueryTransferStatus err: %s", err)
		return
	}
	if resp.Status[xferId].SgnStatus == cbrtypes.XferStatus_SUCCESS {
		return true, fmt.Sprintf("relay with xferId %s already synced", eth.Hash(ev.TransferId).String())
	}

	return
}

func (c *CbrOneChain) skipSyncCbrWithdraw(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.cbrContract.ParseWithdrawDone(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}

	resp, err := cbrcli.QueryWithdrawLiquidityStatus(cliCtx, &cbrtypes.QueryWithdrawLiquidityStatusRequest{
		SeqNum:  ev.Seqnum,
		UsrAddr: ev.Receiver.String(),
	})
	if err != nil {
		// log only, will not skip if request failed
		log.Errorf("QueryWithdrawLiquidityStatus err: %s", err)
		return
	}
	if resp.Status == cbrtypes.WithdrawStatus_WD_COMPLETED {
		return true, fmt.Sprintf("withdrawal with seqNum %d on chain %d already synced", ev.Seqnum, c.chainid)
	}

	return
}

func (c *CbrOneChain) skipSyncCbrWithdrawalRequest(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.wdiContract.ParseWithdrawalRequest(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}

	// check if this withdrawal request is valid
	deadline := time.Unix(ev.Deadline.Int64(), 0)
	if time.Now().After(deadline) {
		return true, fmt.Sprintf("this withdrawal request has passed the deadline %s", deadline.Format(time.Stamp))
	}

	// check if all tokens have the same symbol.
	var chainTokens []*cbrtypes.ChainTokenAddrPair
	for i := range ev.FromChains {
		chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{ChainId: ev.FromChains[i], TokenAddr: ev.Tokens[i].String()})
	}
	symbols, err := cbrcli.QueryAssetsSymbols(cliCtx, chainTokens)
	if err != nil {
		log.Errorf("QueryAssetsSymbols err: %s", err)
		return
	}
	for i := 0; i < len(symbols)-1; i++ {
		if symbols[i] != symbols[i+1] {
			return true, fmt.Sprintf("different kinds of token found in one withdrawal request, %s and %s", symbols[i], symbols[i+1])
		}
	}

	// check if token with the same symbol exists on toChain.
	assets, err := cbrcli.QueryAssetsInfos(cliCtx, []string{symbols[0]}, []uint64{ev.ToChain})
	if err != nil {
		log.Errorf("QueryAssetsInfos err: %s", err)
	}
	if assets[0] == nil {
		return true, fmt.Sprintf("token with symbol %s does not exist on chain %d", symbols[0], ev.ToChain)
	}

	// check for lp origin, in order to minimize invalid withdrawal requests synced.
	lpOriginResp, err := cbrcli.QueryLPOrigin(cliCtx, &cbrtypes.QueryLPOriginRequest{UsrAddr: eth.Addr2Hex(ev.Sender)})
	if err != nil {
		log.Errorf("QueryLPOrigin err: %s, lp address: %s", err, ev.Sender.Hex())
		return
	}
	if lpOriginResp.ChainId != c.chainid {
		return true, fmt.Sprintf("withdrawal with seqNum %d on chain %d is invalid, lp(%x) original chainId is %d",
			ev.SeqNum, c.chainid, ev.Sender, lpOriginResp.ChainId)
	}

	resp, err := cbrcli.QueryWithdrawLiquidityStatus(cliCtx, &cbrtypes.QueryWithdrawLiquidityStatusRequest{
		SeqNum:  ev.SeqNum,
		UsrAddr: ev.Receiver.Hex(),
	})
	if err != nil {
		// withdrawal request has not yet been synced or is syncing still.
		return
	}
	if resp.Status == cbrtypes.WithdrawStatus_WD_WAITING_FOR_SGN ||
		resp.Status == cbrtypes.WithdrawStatus_WD_WAITING_FOR_LP ||
		resp.Status == cbrtypes.WithdrawStatus_WD_COMPLETED {
		return true, fmt.Sprintf("withdrawal with seqNum %d of %s on chain %d already synced", ev.SeqNum, ev.Receiver.Hex(), c.chainid)
	}

	return
}

func (r *Relayer) updateSigners() {
	latestSigners, err := cbrcli.QueryLatestSigners(r.Transactor.CliCtx)
	if err != nil {
		log.Errorln("failed to get latest signers", err)
		return
	}
	sgnBlkTime := viper.GetDuration(common.FlagConsensusTimeoutCommit)

	log.Infoln("update latest signers to", latestSigners.String())
	for i, c := range r.cbrMgr {
		chainId := i
		// TODO, this will be added later, we'll need address to pubkey mapping in another smart contract
		if types.IsFlowChain(chainId) {
			continue
		}
		ssHash, err := c.cbrContract.SsHash(&bind.CallOpts{})
		if err != nil {
			log.Errorln("failed to get sshash", chainId, err)
			continue
		}
		signerBytes := latestSigners.GetSignersBytes()
		// len = 32 bytes Trigger time + N * 32 bytes addrs + N * 32 bytes powers
		if len(signerBytes) < 96 {
			log.Error("Invalid signers bytes")
			return
		}
		// skip the first 32 bytes triggerTime
		signerBytes = signerBytes[32:]
		if eth.Bytes2Hash(crypto.Keccak256(signerBytes)) == ssHash {
			log.Debugf("chain %d signers already updated", chainId)
			continue
		}
		curss := c.getCurss()
		if eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(curss.addrs, curss.powers))) != ssHash {
			log.Warnf("chain %d local curss not match onchain value", chainId)
			continue
		}
		var pass bool
		var sigsBytes [][]byte
		retry := 0
		for !pass && retry < maxSigRetry {
			curssList := make([]*cbrtypes.Signer, 0)
			for i, addr := range curss.addrs {
				power := curss.powers[i]
				curssList = append(curssList, &cbrtypes.Signer{
					Addr:  addr.Bytes(),
					Power: power.Bytes(),
				})
			}
			chainSigners, err := cbrcli.QueryChainSigners(r.Transactor.CliCtx, chainId)
			if err != nil {
				log.Errorf("failed to get chain %d signers, err: %s", chainId, err)
			} else {
				pass, sigsBytes = cbrtypes.ValidateSigQuorum(chainSigners.GetSortedSigs(), curssList)
				if pass {
					break
				}
			}
			time.Sleep(sgnBlkTime)
			retry++
		}
		if !pass {
			log.Errorf("chain %d signers not enough yet", chainId)
			continue
		}

		tx, err := c.Transactor.Transact(
			&ethutils.TransactionStateHandler{
				OnMined: func(receipt *ethtypes.Receipt) {
					if receipt.Status == ethtypes.ReceiptStatusSuccessful {
						log.Infof("chain %d UpdateSigners transaction %x succeeded", chainId, receipt.TxHash)
					} else {
						log.Errorf("chain %d UpdateSigners transaction %x failed", chainId, receipt.TxHash)
					}
				},
				OnError: func(tx *ethtypes.Transaction, err error) {
					log.Errorf("chain %d UpdateSigners transaction %x err: %s", chainId, tx.Hash(), err)
				},
			},
			func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
				newSignerAddrs, newSignerPowers := cbrtypes.SignersToEthArrays(latestSigners.GetSortedSigners())
				return c.cbrContract.UpdateSigners(
					opts, new(big.Int).SetUint64(latestSigners.TriggerTime),
					newSignerAddrs, newSignerPowers, sigsBytes, curss.addrs, curss.powers)
			},
		)
		if err != nil {
			log.Errorf("chain %d update signer err %s", chainId, err)
			continue
		}
		log.Infof("chain %d UpdateSigners tx %x submitted", chainId, tx.Hash())
	}
	r.setCbrSsUpdated()
}
