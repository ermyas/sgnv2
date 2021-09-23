package relayer

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/staking/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	mapset "github.com/deckarep/golang-set"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
)

const (
	maxRelayRetry = 5
)

// sleep, check if syncer, if yes, go over cbr dbs to send tx
func (r *Relayer) doCbridge(cbrMgr CbrMgr) {
	for {
		time.Sleep(15 * time.Second)
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
	}
}

func (r *Relayer) processCbridgeQueue() {
	var keys, vals [][]byte
	r.lock.RLock()
	iterator, err := r.db.Iterator(CbrXferKeyPrefix, storetypes.PrefixEndBytes(CbrXferKeyPrefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
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
	log.Infoln("Process relay ", string(relayEvent.XferId))

	relay, err := cbrcli.QueryRelay(r.Transactor.CliCtx, relayEvent.XferId)
	if err != nil {
		log.Errorln("QuerySlash err", err)
		return
	}

	relayOnChain := new(cbrtypes.RelayOnChain)
	err = relayOnChain.Unmarshal(relay.Relay)
	if err != nil {
		log.Errorln("Unmarshal relay.Relay err", err)
		return
	}

	signedValidators := mapset.NewSet()
	for _, sig := range relay.SortedSigs {
		signedValidators.Add(string(sig.Addr))
	}
	pass, allValidators := r.validateSigs(signedValidators)
	if !pass {
		log.Debugf("relay %s does not have enough sigs", relayEvent.XferId)
		r.requeueRelay(relayEvent)
		return
	}

	currss, _ := proto.Marshal(GetSortedSigners(allValidators))

	err = r.cbrMgr[relayOnChain.DstChainId].SendRelay(relay.Relay, currss, relay.GetSortedSigsBytes())
	if err != nil {
		r.requeueRelay(relayEvent)
		log.Errorln("relay err", err)
		return
	}
}

func GetSortedSigners(validators types.Validators) *cbrtypes.SortedSigners {
	signers := make([]*cbrtypes.AddrAmt, 0)
	for _, v := range validators {
		signers = append(signers, &cbrtypes.AddrAmt{
			Addr: []byte(eth.Addr2Hex(v.GetSignerAddr())),
			Amt:  v.BondedTokens().BigInt().Bytes(),
		})
	}
	sort.Slice(signers, func(i, j int) bool {
		return string(signers[i].Addr) < string(signers[j].Addr)
	})
	return &cbrtypes.SortedSigners{
		Signers: signers,
	}
}

func (r *Relayer) requeueRelay(relayEvent RelayEvent) {
	if relayEvent.RetryCount >= maxRelayRetry {
		log.Infof("relay %s hits retry limit", relayEvent.XferId)
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
		iter, err := c.db.Iterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn)))
		if err != nil {
			log.Errorln("chainID:", chid, evn, "iter err:", err)
			continue
		}
		for ; iter.Valid(); iter.Next() {
			onchev := &cbrtypes.OnChainEvent{
				Chainid: chid,
				Evtype:  evn,
				Elog:    iter.Value(),
			}
			data, _ := onchev.Marshal()
			ret = append(ret, &synctypes.ProposeUpdate{
				Type:       synctypes.DataType_CbrOnchainEvent,
				ChainId:    chid,
				ChainBlock: 0, // why do we need this in ProposeUpdate?
				Data:       data,
			})
		}
		iter.Close()
	}
	return ret
}

// to be called by r.verifyUpdate
// decode event and check if I also have event in db
// TODO: query x/cbridge to make sure event not processed
// data is marshaled OnChainEvent, see above line 53
func (r *Relayer) verifyCbrEventUpdate(data []byte) (done, approve bool) {
	onchev := new(cbrtypes.OnChainEvent)
	err := onchev.Unmarshal(data)
	if err != nil {
		log.Errorf("failed to unmarshal %x to onchain event msg", data)
		return true, false
	}
	elog := new(ethtypes.Log)
	err = json.Unmarshal(onchev.Elog, elog)
	if err != nil {
		log.Errorf("failed to unmarshal %x to eth Log", onchev.Elog)
		return true, false
	}

	// delete my local db event so this event won't be picked again when I become syncer
	defer r.cbrMgr[onchev.Chainid].delEvent(onchev.Evtype, elog.BlockNumber, uint64(elog.Index))

	// now we directly verify this event onchain
	// why we have to do onchain check instead of local db only: our local db could
	// be behind so we can't differentiate not yet see event vs. faked event
	// only onchain query can give us 100% certainty
	// we have 2 ways: getTransactionReceipt or getLogs
	equal, err := r.cbrMgr[onchev.Chainid].CheckEvent(onchev.Evtype, elog)
	if err != nil {
		return false, false // onchain error, so don't vote
	}
	if !equal {
		return true, false
	}
	// event is the same as onchain, now move on to per event logic
	// eg. query x/cbridge if this event has already been handled
	// but if we only query applied state, we may still vote again?
	// unless we can query x/cbridge w/ state plus pending?
	// TODO: query x/cbridge

	// now per event logic
	switch onchev.Evtype {
	case CbrEventLiqAdd:
		// if chid-seq already processed, return true, false
		return true, true
	case CbrEventSend:
		// if transferid is waiting for vote status, return true, true
		// otherwise, true, false
		return true, true
	case CbrEventRelay:
		// this event means syncer already submitted relay tx onchain
		return true, true
	}
	return true, false
}
