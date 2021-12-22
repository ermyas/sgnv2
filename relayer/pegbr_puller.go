package relayer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegbrcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

var pegEvNames = []string{
	pegbrtypes.PegbrEventDeposited,
	pegbrtypes.PegbrEventBurn,
	pegbrtypes.PegbrEventMint,
	pegbrtypes.PegbrEventWithdrawn,
}

// sleep, check if syncer, if yes, go over cbr dbs to send pegbr tx
func (r *Relayer) doPegbrSync(cbrMgr CbrMgr) {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalCbridge)) * time.Second
	log.Infoln("start process pegbr sync, interval:", interval)
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
			ret, isUpdateMsgFull := onech.pullPegbrEvents(chid, r.Transactor.CliCtx, &updatesBytesLen)
			msg.Updates = append(msg.Updates, ret...)
			if isUpdateMsgFull {
				break
			}
		}
		if len(msg.Updates) > 0 {
			r.Transactor.AddTxMsg(msg)
			log.Debugln("PegbrEvent updates count in one msg:", len(msg.Updates))
		}
	}
}

// sleep, check if syncer, if yes, go over cbr dbs to send pegbr tx
func (r *Relayer) doPegbrOnchain(cbrMgr CbrMgr) {
	for chid := range cbrMgr {
		go r.doPegbrOnchainByChain(chid)
	}
}

func (r *Relayer) doPegbrOnchainByChain(chid uint64) {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalCbridge)) * time.Second
	log.Infof("start process pegbr onchain, interval:%s, chainId: %d", interval, chid)
	for {
		time.Sleep(interval)

		r.processPegbrMintQueue(chid)
		r.processPegbrWithdrawQueue(chid)
	}
}

func (r *Relayer) processPegbrMintQueue(chid uint64) {
	syncer, syncerUpdateTime := r.getSyncer()
	if !syncer {
		return
	}

	var keys, vals [][]byte
	r.pegbrLock.RLock()
	prefix := GetPegbrMintPrefix(chid)
	iterator, err := r.db.Iterator(prefix, storetypes.PrefixEndBytes(prefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
		r.pegbrLock.RUnlock()
		return
	}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
		vals = append(vals, iterator.Value())
	}
	iterator.Close()
	r.pegbrLock.RUnlock()

	if len(keys) > 0 {
		log.Debugf("start process mint queue，current timestamp: %d, queue size: %d, chainid: %d", time.Now().Unix(), len(keys), chid)
	}

	newSyncer := false
	newSyncerWaitTime := time.Duration(r.cbrMgr[chid].blkInterval) * time.Second * newSyncerWaitBlk
	if syncerUpdateTime.Add(newSyncerWaitTime).After(time.Now()) {
		newSyncer = true
	}
	sigWaitTime := viper.GetDuration(common.FlagConsensusTimeoutCommit) * sigWaitSgnBlk
	for i, key := range keys {
		event := NewMintRequestFromBytes(vals[i])
		if event.CreateTime.Add(sigWaitTime).After(time.Now()) {
			// wait a while to collect validator signatures
			continue
		}
		if newSyncer && event.CreateTime.Before(syncerUpdateTime) {
			// wait for mint to be submitted by the previous syncer
			continue
		}

		err = r.dbDelete(key)
		if err != nil {
			log.Errorln("db Delete err", err)
			continue
		}

		r.submitMint(event)
	}
}

func (r *Relayer) submitMint(mintRequest MintRequest) {
	logmsg := fmt.Sprintf("Process peg mint, mintChain %d mintId %x depositChainId %d depositId %x",
		mintRequest.MintChainId, mintRequest.MintId, mintRequest.DepositChainId, mintRequest.DepositId)

	mintInfo, err := pegbrcli.QueryMintInfo(r.Transactor.CliCtx, eth.Bytes2Hex(mintRequest.MintId))
	if err != nil {
		r.requeueMint(mintRequest, fmt.Sprintf("%s. QueryMintInfo err: %s", logmsg, err), true)
		return
	}
	if mintInfo.Success {
		log.Infof("%s. mint already completed, skip it", logmsg)
		return
	}

	mintOnChain := new(pegbrtypes.MintOnChain)
	err = mintOnChain.Unmarshal(mintInfo.MintProtoBytes)
	if err != nil {
		log.Errorf("%s. Unmarshal mintInfo.MintProtoBytes err %s", logmsg, err)
		return
	}

	curss := r.cbrMgr[mintRequest.MintChainId].getCurss()
	curssList := make([]*cbrtypes.Signer, 0)
	for i, addr := range curss.addrs {
		power := curss.powers[i]
		curssList = append(curssList, &cbrtypes.Signer{
			Addr:  addr.Bytes(),
			Power: power.Bytes(),
		})
	}
	pass, sigsBytes := cbrtypes.ValidateSigQuorum(mintInfo.GetAddrSigs(), curssList)
	if !pass {
		r.requeueMint(mintRequest,
			fmt.Sprintf("%s. Not have enough sigs %s, curss %s", logmsg, mintInfo.SignersStr(), curss.String()), false)
		return
	}

	txHash, err := r.cbrMgr[mintRequest.MintChainId].SendMint(mintInfo.MintProtoBytes, sigsBytes, curss, mintOnChain)
	if err != nil {
		if strings.Contains(err.Error(), "record exists") {
			log.Infof("%s. err %s, skip it", logmsg, err)
			return
		}

		if strings.Contains(err.Error(), "Pausable: paused") ||
			strings.Contains(err.Error(), "volume exceeds cap") ||
			strings.Contains(err.Error(), "Mismatch current signers") {
			if mintRequest.RetryCount > 0 {
				mintRequest.RetryCount -= 1
			}
		}
		r.requeueMint(mintRequest, fmt.Sprintf("%s. err %s", logmsg, err), true)
		return
	}
	log.Infof("%s. tx hash %s", logmsg, txHash)
}

func (r *Relayer) requeueMint(mintRequest MintRequest, logmsg string, warn bool) {
	if mintRequest.RetryCount >= maxRelayRetry {
		log.Errorf("%s. hits retry limit", logmsg)
		return
	}
	mintRequest.RetryCount += 1
	err := r.dbSet(GetPegbrMintKey(mintRequest.MintChainId, mintRequest.DepositChainId, mintRequest.DepositId), mintRequest.MustMarshal())
	if err != nil {
		log.Errorf("%s. db Set err: %s", logmsg, err)
	}
	if warn {
		log.Warn(logmsg)
	} else {
		log.Debug(logmsg)
	}
}

func (r *Relayer) processPegbrWithdrawQueue(chid uint64) {
	syncer, syncerUpdateTime := r.getSyncer()
	if !syncer {
		return
	}

	var keys, vals [][]byte
	r.pegbrLock.RLock()
	prefix := GetPegbrWdPrefix(chid)
	iterator, err := r.db.Iterator(prefix, storetypes.PrefixEndBytes(prefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
		r.pegbrLock.RUnlock()
		return
	}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
		vals = append(vals, iterator.Value())
	}
	iterator.Close()
	r.pegbrLock.RUnlock()

	if len(keys) > 0 {
		log.Debugf("start process withdraw queue，current timestamp: %d, queue size: %d, chainid: %d", time.Now().Unix(), len(keys), chid)
	}

	newSyncer := false
	newSyncerWaitTime := time.Duration(r.cbrMgr[chid].blkInterval) * time.Second * newSyncerWaitBlk
	if syncerUpdateTime.Add(newSyncerWaitTime).After(time.Now()) {
		newSyncer = true
	}
	sigWaitTime := viper.GetDuration(common.FlagConsensusTimeoutCommit) * sigWaitSgnBlk
	for i, key := range keys {
		event := NewWithdrawRequestFromBytes(vals[i])
		if event.CreateTime.Add(sigWaitTime).After(time.Now()) {
			// wait a while to collect validator signatures
			continue
		}
		if newSyncer && event.CreateTime.Before(syncerUpdateTime) {
			// wait for withdraw to be submitted by the previous syncer
			continue
		}

		err = r.dbDelete(key)
		if err != nil {
			log.Errorln("db Delete err", err)
			continue
		}

		r.submitWithdraw(event)
	}
}

func (r *Relayer) submitWithdraw(wdRequest WithdrawRequest) {
	logmsg := fmt.Sprintf("Process peg withdraw, withdrawChain %d withdrawId %x", wdRequest.WithdrawChainId, wdRequest.WithdrawId)

	wdInfo, err := pegbrcli.QueryWithdrawInfo(r.Transactor.CliCtx, eth.Bytes2Hex(wdRequest.WithdrawId))
	if err != nil {
		r.requeueWithdraw(wdRequest, fmt.Sprintf("%s. QueryMintInfo err: %s", logmsg, err), true)
		return
	}
	if wdInfo.Success {
		log.Infof("%s. withdraw already completed, skip it", logmsg)
		return
	}

	wdOnChain := new(pegbrtypes.WithdrawOnChain)
	err = wdOnChain.Unmarshal(wdInfo.WithdrawProtoBytes)
	if err != nil {
		log.Errorf("%s. Unmarshal wdInfo.WithdrawProtoBytes err %s", logmsg, err)
		return
	}

	curss := r.cbrMgr[wdRequest.WithdrawChainId].getCurss()
	curssList := make([]*cbrtypes.Signer, 0)
	for i, addr := range curss.addrs {
		power := curss.powers[i]
		curssList = append(curssList, &cbrtypes.Signer{
			Addr:  addr.Bytes(),
			Power: power.Bytes(),
		})
	}
	pass, sigsBytes := cbrtypes.ValidateSigQuorum(wdInfo.GetAddrSigs(), curssList)
	if !pass {
		r.requeueWithdraw(wdRequest,
			fmt.Sprintf("%s. Not have enough sigs %s, curss %s", logmsg, wdInfo.SignersStr(), curss.String()), false)
		return
	}

	txHash, err := r.cbrMgr[wdRequest.WithdrawChainId].SendWithdraw(wdInfo.WithdrawProtoBytes, sigsBytes, curss, wdOnChain)
	if err != nil {
		if strings.Contains(err.Error(), "record exists") {
			log.Infof("%s. err %s, skip it", logmsg, err)
			return
		}

		if strings.Contains(err.Error(), "Pausable: paused") ||
			strings.Contains(err.Error(), "volume exceeds cap") ||
			strings.Contains(err.Error(), "Mismatch current signers") {
			if wdRequest.RetryCount > 0 {
				wdRequest.RetryCount -= 1
			}
		}
		r.requeueWithdraw(wdRequest, fmt.Sprintf("%s. err %s", logmsg, err), true)
		return
	}
	log.Infof("%s. tx hash %s", logmsg, txHash)
}

func (r *Relayer) requeueWithdraw(wdRequest WithdrawRequest, logmsg string, warn bool) {
	if wdRequest.RetryCount >= maxRelayRetry {
		log.Errorf("%s. hits retry limit", logmsg)
		return
	}
	wdRequest.RetryCount += 1
	err := r.dbSet(GetPegbrWdKey(wdRequest.WithdrawChainId, wdRequest.BurnChainId, wdRequest.BurnId), wdRequest.MustMarshal())
	if err != nil {
		log.Errorf("%s. db Set err: %s", logmsg, err)
	}
	if warn {
		log.Warn(logmsg)
	} else {
		log.Debug(logmsg)
	}
}

// TODO: query x/pegbridge to skip already processed events to avoid duplicated propose
// Note if syncer changes before EndBlock, new syncer may still propose again
// the 2nd propose shouldn't get votes because when verify, sgn nodes will find it's already processed
// even it is voted, apply will still fail because x/cbr will err
func (c *CbrOneChain) pullPegbrEvents(chid uint64, cliCtx client.Context, updatesBytesLen *int) (ret []*synctypes.ProposeUpdate, isUpdateMsgFull bool) {
	// 1st loop over event names, then go over iter
	isUpdateMsgFull = false
	for _, evn := range pegEvNames {
		var keys, vals [][]byte
		c.pegbrLock.RLock()
		iterator, err := c.db.Iterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn)))
		if err != nil {
			log.Errorln("Create db iterator err", err)
			c.pegbrLock.RUnlock()
			continue
		}
		for ; iterator.Valid(); iterator.Next() {
			keys = append(keys, iterator.Key())
			vals = append(vals, iterator.Value())
		}
		iterator.Close()
		c.pegbrLock.RUnlock()

		pegbrUserActionValidCache := make(map[string]bool)
		for i, key := range keys {
			err = c.db.Delete(key)
			if err != nil {
				log.Errorln("db Delete err", err)
				continue
			}

			evlog := new(ethtypes.Log)
			err := json.Unmarshal(vals[i], evlog)
			if err != nil {
				log.Errorf("failed to unmarshal onchev elog, key:%s, err:%s", string(key), err.Error())
				continue
			}

			skip, reason := c.skipPegbrEvent(evn, evlog, cliCtx, pegbrUserActionValidCache)
			if skip {
				log.Debugf("skip pbr event: %s, chid %d, reason: %s", string(key), c.chainid, reason)
				continue
			}

			onchev := &cbrtypes.OnChainEvent{
				Chainid: chid,
				Evtype:  evn,
				Elog:    vals[i],
			}
			data, _ := onchev.Marshal()
			update := &synctypes.ProposeUpdate{
				Type:    synctypes.DataType_PegbrOnChainEvent,
				ChainId: chid,
				Data:    data,
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

func (c *CbrOneChain) skipPegbrEvent(evn string, evlog *ethtypes.Log, cliCtx client.Context, checkedCache map[string]bool) (skip bool, reason string) {
	switch evn {
	case pegbrtypes.PegbrEventDeposited:
		skip, reason = c.skipSyncPegbrDeposit(evlog, cliCtx, checkedCache)
	case pegbrtypes.PegbrEventBurn:
		skip, reason = c.skipSyncPegbrBurn(evlog, cliCtx, checkedCache)
	case pegbrtypes.PegbrEventMint:
		skip, reason = c.skipSyncPegbrMint(evlog, cliCtx)
	case pegbrtypes.PegbrEventWithdrawn:
		skip, reason = c.skipSyncPegbrWithdrawn(evlog, cliCtx)
	}
	return
}

func (c *CbrOneChain) skipSyncPegbrDeposit(
	evlog *ethtypes.Log, cliCtx client.Context, validCache map[string]bool) (skip bool, reason string) {

	ev, err := c.pegContracts.vault.ParseDeposited(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	// we should check cache first
	cacheKey := fmt.Sprintf("%d-%d-%x", c.chainid, ev.MintChainId, ev.Token)

	if validCache != nil {
		cacheValid, found := validCache[cacheKey]
		if found && !cacheValid {
			return true, "invalid pegbr deposit"
		}
	}

	req := &pegbrtypes.QueryOrigPeggedPairsRequest{
		Orig: &commontypes.ContractInfo{
			ChainId: c.chainid,
			Address: ev.Token.Hex(),
		},
		Pegged: &commontypes.ContractInfo{
			ChainId: ev.MintChainId,
		},
	}
	pairs, err := pegbrcli.QueryOrigPeggedPairs(cliCtx, req)
	if len(pairs) == 0 {
		// If request failed, we will not break this flow.
		// As if invalid token send event go to the apply flow, sgn will also check it and set it to refund flow.
		log.Errorf("fail to lookup pegged pair, ev:%s, err:%s", ev.PrettyLog(c.chainid), err)
		// may be call sgn fail, we still send this ev to sgn and sgn to do the check again.
		return
	}
	// Only single pair
	pair := pairs[0]
	// cached and can reduce some cli call
	if validCache != nil {
		validCache[cacheKey] = pair.Pegged.Address != ""
	}

	resp, err := pegbrcli.QueryDepositInfo(cliCtx, eth.Bytes2Hex(ev.DepositId[:]))
	if err != nil && !strings.Contains(err.Error(), "no info found") {
		// log only, will not skip if request failed
		log.Errorf("QueryDepositInfo err: %s", err)
		return
	}
	if resp.DepositId != nil {
		return true, fmt.Sprintf("deposit %x already synced", ev.DepositId)
	}

	return
}

func (c *CbrOneChain) skipSyncPegbrBurn(
	evlog *ethtypes.Log, cliCtx client.Context, validCache map[string]bool) (skip bool, reason string) {

	ev, err := c.pegContracts.bridge.ParseBurn(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	// we should check cache first
	cacheKey := fmt.Sprintf("%d-%x", c.chainid, ev.Token)

	if validCache != nil {
		cacheValid, found := validCache[cacheKey]
		if found && !cacheValid {
			return true, "invalid pegbr burn"
		}
	}

	req := &pegbrtypes.QueryOrigPeggedPairsRequest{
		Pegged: &commontypes.ContractInfo{
			ChainId: c.chainid,
			Address: ev.Token.Hex(),
		},
	}

	pairs, err := pegbrcli.QueryOrigPeggedPairs(cliCtx, req)
	if len(pairs) == 0 {
		// If request failed, we will not break this flow.
		// As if invalid token send event go to the apply flow, sgn will also check it and set it to refund flow.
		log.Errorf("fail to lookup pegged pair, ev:%s, err:%s", ev.PrettyLog(c.chainid), err)
		// may be call sgn fail, we still send this ev to sgn and sgn to do the check again.
		return
	}
	// Only single pair
	pair := pairs[0]
	// cached and can reduce some cli call
	if validCache != nil {
		validCache[cacheKey] = pair.Orig.Address != ""
	}

	resp, err := pegbrcli.QueryBurnInfo(cliCtx, eth.Bytes2Hex(ev.BurnId[:]))
	if err != nil && !strings.Contains(err.Error(), "no info found") {
		// log only, will not skip if request failed
		log.Errorf("QueryBurnInfo err: %s", err)
		return
	}
	if resp.BurnId != nil {
		return true, fmt.Sprintf("burn %x already synced", ev.BurnId)
	}

	return
}

func (c *CbrOneChain) skipSyncPegbrMint(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.pegContracts.bridge.ParseMint(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	resp, err := pegbrcli.QueryMintInfo(cliCtx, eth.Hash(ev.MintId).Hex())
	if err != nil {
		// log only, will not skip if request failed
		log.Errorf("QueryMintInfo err: %s", err)
		return
	}
	if resp.Success {
		return true, fmt.Sprintf("mint %x already synced", ev.MintId)
	}
	return
}

func (c *CbrOneChain) skipSyncPegbrWithdrawn(evlog *ethtypes.Log, cliCtx client.Context) (skip bool, reason string) {
	ev, err := c.pegContracts.vault.ParseWithdrawn(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	resp, err := pegbrcli.QueryWithdrawInfo(cliCtx, eth.Hash(ev.WithdrawId).Hex())
	if err != nil {
		// log only, will not skip if request failed
		log.Errorf("QueryWithdrawInfo err: %s", err)
		return
	}
	if resp.Success {
		return true, fmt.Sprintf("withdraw %x already synced", ev.WithdrawId)
	}
	return
}
