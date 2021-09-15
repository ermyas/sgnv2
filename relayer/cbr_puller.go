package relayer

import (
	"encoding/json"
	"time"

	"github.com/celer-network/goutils/log"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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
	// now we directly verify this event onchain
	// why we have to do onchain check instead of local db only: our local db could
	// be behind so we can't differentiate not yet see event vs. faked event
	// only onchain query can give us 100% certainty
	// we have 2 ways: getTransactionReceipt or getLogs
	equal, err := r.cbrMgr[onchev.Chainid].CheckEvent(elog)
	if err != nil {
		return false, false // onchain error, allows retry
	}
	// delete my local db event so this event won't be picked again when I become syncer
	defer r.cbrMgr[onchev.Chainid].delEvent(onchev.Evtype, elog.BlockNumber, uint64(elog.Index))
	if !equal {
		return true, false
	}
	// event is the same as onchain, now move on to per event logic
	// now query x/cbridge if this event has already been handled
	// but if we only query applied state, we may still vote again?
	// unless we can query x/cbridge w/ state plus pending?
	// TODO: query x/cbridge

	// now per event logic
	switch onchev.Evtype {
	case CbrEventLiqAdd:
		return true, true
	}
	return true, false
}
