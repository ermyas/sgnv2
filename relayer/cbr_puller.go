package relayer

import (
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

const (
	maxRelayRetry = 5
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
			msg.Updates = append(msg.Updates, onech.pullEvents(chid)...)
		}
		if len(msg.Updates) > 0 {
			// or we should call cbridge grpc here?
			r.Transactor.AddTxMsg(msg)
		}

		r.processCbridgeQueue()

		if r.isCbrSsUpdating() {
			latestSs, err := cbrcli.QueryLatestSigners(r.Transactor.CliCtx)
			if err != nil {
				log.Errorln("failed to get latest signers", err)
				continue
			}
			r.updateSigners(latestSs)
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
	pass := validateCbrSigs(relay.SortedSigs, curss.signers)
	if !pass {
		log.Debugf("%s. Not have enough sigs %s, curss %s", logmsg, relay.SignersStr(), curss.signers.String())
		r.requeueRelay(relayEvent)
		return
	}
	// TODO: check if relay already sent on chain
	log.Infof("%s with signers %s", logmsg, relay.SignersStr())
	err = r.cbrMgr[relayOnChain.DstChainId].SendRelay(relay.Relay, curss.bytes, relay.GetSortedSigsBytes())
	if err != nil {
		r.requeueRelay(relayEvent)
		log.Errorln("relay err", err)
		return
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
// the 2nd propose shouldn't get votes? why? MUST confirm this
func (c *CbrOneChain) pullEvents(chid uint64) []*synctypes.ProposeUpdate {
	var ret []*synctypes.ProposeUpdate
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

func (r *Relayer) updateSigners(latestSs *cbrtypes.LatestSigners) {
	updated := true
	for chainId, c := range r.cbrMgr {
		ssHash, err := c.contract.SsHash(&bind.CallOpts{})
		if err != nil {
			log.Errorln("failed to get sshash", chainId, err)
			updated = false
			continue
		}
		if eth.Bytes2Hash(crypto.Keccak256(latestSs.GetSignersBytes())) == ssHash {
			log.Debugf("signers for chain %d already updated", chainId)
			continue
		}
		// TODO: fine-grainded per chain updated flag
		updated = false

		if !validateCbrSigs(latestSs.SortedSigs, c.curss.signers) {
			log.Infof("chain %d signers not enough yet", chainId)
			continue
		}
		err = c.UpdateSigners(latestSs.SignersBytes, c.curss.bytes, latestSs.GetSortedSigsBytes())
		if err != nil {
			log.Error(err)
		}
	}
	if updated {
		r.setCbrSsUpdated()
	}
}
