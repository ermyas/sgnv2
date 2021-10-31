package relayer

import (
	"encoding/json"
	"fmt"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

const (
	maxRelayRetry = 5
	maxSigRetry   = 10
)

// sleep, check if syncer, if yes, go over cbr dbs to send tx
func (r *Relayer) doCbridge(cbrMgr CbrMgr) {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalCbridge)) * time.Second
	log.Infoln("start process cbridge queue, interval:", interval)
	for {
		time.Sleep(interval)
		if !r.isSyncer() {
			continue
		}
		// find all events need to be sent out, batch into one msg
		msg := &synctypes.MsgProposeUpdates{
			Sender: r.Transactor.Key.GetAddress().String(),
		}

		for chid, onech := range cbrMgr {
			// go over each chain db events, send msg
			msg.Updates = append(msg.Updates, onech.pullEvents(chid, r.Transactor.CliCtx)...)
		}
		if len(msg.Updates) > 0 {
			// or we should call cbridge grpc here?
			r.Transactor.AddTxMsg(msg)
		}

		r.processCbridgeQueue()

		if r.isCbrSsUpdating() {
			r.updateSigners()
		}
	}
}

func (r *Relayer) processCbridgeQueue() {
	var keys, vals [][]byte
	r.lock.RLock()
	iterator, err := r.db.Iterator(CbrXferKeyPrefix, storetypes.PrefixEndBytes(CbrXferKeyPrefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
		r.lock.RUnlock()
		return
	}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
		vals = append(vals, iterator.Value())
	}
	iterator.Close()
	r.lock.RUnlock()

	for i, key := range keys {
		event := NewRelayEventFromBytes(vals[i])
		err = r.dbDelete(key)
		if err != nil {
			log.Errorln("db Delete err", err)
			continue
		}
		r.submitRelay(event)
	}
}

func (r *Relayer) submitRelay(relayEvent RelayEvent) {
	logmsg := fmt.Sprintf("Process relay %x", relayEvent.XferId)

	relay, err := cbrcli.QueryRelay(r.Transactor.CliCtx, relayEvent.XferId)
	if err != nil {
		log.Errorf("%s. QueryRelay err: %s", logmsg, err)
		return
	}

	relayOnChain := new(cbrtypes.RelayOnChain)
	err = relayOnChain.Unmarshal(relay.Relay)
	if err != nil {
		log.Errorf("%s. Unmarshal relay.Relay err %s", logmsg, err)
		return
	}

	curss := r.cbrMgr[relayOnChain.DstChainId].getCurss()
	pass, sigsBytes := validateSigQuorum(relay.SortedSigs, curss)
	if !pass {
		log.Debugf("%s. Not have enough sigs %s, curss %s", logmsg, relay.SignersStr(), curss.String())
		r.requeueRelay(relayEvent)
		return
	}
	relayTransferId := relayOnChain.GetRelayOnChainTransferId()
	existRelay, existRelayErr := r.cbrMgr[relayOnChain.DstChainId].existTransferId(relayTransferId)
	if existRelayErr != nil {
		// if fail to query, continue to send this relay, because we can not make sure whether the relay already exist.
		log.Warnf("fail to query this relay by transfer id, relay src transfer id:%x, dest transfer id:%x, err:%s",
			relayOnChain.SrcChainId, relayTransferId, existRelayErr.Error())
	} else if existRelay {
		log.Infof("relay already exist on chain, skip it, relay src transfer id:%x, dest transfer id:%x",
			relayOnChain.SrcChainId, relayTransferId)
		return
	}
	log.Infof("%s with signers %s", logmsg, relay.SignersStr())
	txHash, err := r.cbrMgr[relayOnChain.DstChainId].SendRelay(relay.Relay, sigsBytes, curss)
	if err != nil {
		r.requeueRelay(relayEvent)
		log.Errorln("relay err", err)
		return
	}
	err = dal.UpdateTransferRelayedStatus(common.Bytes2Hash(relayEvent.XferId).String(), txHash)
	if err != nil {
		log.Errorln("failed in UpdateTransferRelayedStatus:", err)
	}
}

func (r *Relayer) requeueRelay(relayEvent RelayEvent) {
	if relayEvent.RetryCount >= maxRelayRetry {
		log.Infof("relay %x hits retry limit", relayEvent.XferId)
		return
	}

	relayEvent.RetryCount = relayEvent.RetryCount + 1
	err := r.dbSet(GetCbrXferKey(relayEvent.XferId), relayEvent.MustMarshal())
	if err != nil {
		log.Errorln("db Set err", err)
	}
}

// TODO: query x/cbridge to skip already processed events to avoid duplicated propose
// Note if syncer changes before EndBlock, new syncer may still propose again
// the 2nd propose shouldn't get votes because when verify, sgn nodes will find it's already processed
// even it is voted, apply will still fail because x/cbr will err
func (c *CbrOneChain) pullEvents(chid uint64, CliCtx client.Context) []*synctypes.ProposeUpdate {
	var ret []*synctypes.ProposeUpdate
	// to make it simple we use "srcChainId-destChainId-srcTokenAddr" as key, and valid as val.
	// this cache can only be used in only one pullEvents, if pull again, we should create and use a new cache.
	cbrSendValidCache := make(map[string]bool)
	// 1st loop over event names, then go over iter
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
				log.Errorf("failed to unmarshal onchev elog, key:%s, err:%s", string(key), err.Error())
				continue
			}
			var skip bool
			var reason string
			// TODO: other events
			switch evn {
			case cbrtypes.CbrEventSend:
				skip, reason = c.skipSyncCbrSend(evlog, CliCtx, cbrSendValidCache)
			case cbrtypes.CbrEventSignersUpdated:
				skip, reason = c.skipSyncCbrSignerUdpate(evlog, CliCtx)
			}
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
			ret = append(ret,
				&synctypes.ProposeUpdate{
					Type:       synctypes.DataType_CbrOnchainEvent,
					ChainId:    chid,
					ChainBlock: 0, // why do we need this in ProposeUpdate?
					Data:       data,
				},
			)
		}
	}
	return ret
}

func (c *CbrOneChain) skipSyncCbrSend(
	evlog *ethtypes.Log, CliCtx client.Context, validCache map[string]bool) (skip bool, reason string) {

	sendEv, err := c.contract.ParseSend(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	// we should check cache first
	cacheKey := fmt.Sprintf("%d-%d-%x", c.chainid, sendEv.DstChainId, sendEv.Token)
	cacheValid, found := validCache[cacheKey]
	if found && !cacheValid {
		return true, "invalid cbr send"
	}
	checkReq := &cbrtypes.CheckChainTokenValidRequest{
		SrcChainId:   c.chainid,
		DestChainId:  sendEv.DstChainId,
		SrcTokenAddr: eth.Addr2Hex(sendEv.Token),
	}
	checkResp, checkRespErr := cbrcli.QueryCheckChainTokenValid(CliCtx, checkReq)
	if checkRespErr != nil {
		// If request failed, we will not break this flow.
		// As if invalid token send event go to the apply flow, sgn will also check it and set it to refund flow.
		log.Errorf("fail to check chain token valid, sendEv:%s, err:%s", sendEv.PrettyLog(c.chainid), checkRespErr.Error())
		// may be call sgn fail, we still send this ev to sgn and sgn to do the check again.
		return
	} else {
		// cached and can reduce some cli call
		validCache[cacheKey] = checkResp.GetValid()
		if !checkResp.GetValid() {
			return true, "invalid cbr send"
		}
	}
	// TODO: check is sendEv already snynced
	return
}

func (c *CbrOneChain) skipSyncCbrSignerUdpate(evlog *ethtypes.Log, CliCtx client.Context) (skip bool, reason string) {
	ev, err := c.contract.ParseSignersUpdated(*evlog)
	if err != nil {
		return true, fmt.Sprintf("fail to parse event, txHash:%x, err:%s", evlog.TxHash, err)
	}
	ssHash, err := c.contract.SsHash(&bind.CallOpts{})
	if err != nil {
		log.Errorf("chain %d failed to get onchain sshash err %s", c.chainid, err)
		return
	}
	if eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(ev.Signers, ev.Powers))) != ssHash {
		return true, "not match onchain sshash, maybe outdated"
	}

	chainSigners, err := cbrcli.QueryChainSigners(CliCtx, c.chainid)
	if err == nil {
		addrs, powers := cbrtypes.SignersToEthArrays(chainSigners.SortedSigners)
		if eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(addrs, powers))) == ssHash {
			return true, "chain signers already updated"
		}
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
	for chainId, c := range r.cbrMgr {
		ssHash, err := c.contract.SsHash(&bind.CallOpts{})
		if err != nil {
			log.Errorln("failed to get sshash", chainId, err)
			continue
		}
		if eth.Bytes2Hash(crypto.Keccak256(latestSigners.GetSignersBytes())) == ssHash {
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
			pass, sigsBytes = validateSigQuorum(latestSigners.GetSortedSigs(), curss)
			if pass {
				break
			}
			time.Sleep(sgnBlkTime)
			latestSigners, err = cbrcli.QueryLatestSigners(r.Transactor.CliCtx)
			if err != nil {
				log.Errorln("failed to get latest signers", err)
			}
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
				return c.contract.UpdateSigners(
					opts, newSignerAddrs, newSignerPowers, sigsBytes, curss.addrs, curss.powers)
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
