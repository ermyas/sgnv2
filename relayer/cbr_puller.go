package relayer

import (
	"time"

	"github.com/celer-network/goutils/log"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// sleep, check if syncer, if yes, go over cbr dbs to send tx
func (r *Relayer) doCbridge(cbrMgr CbrMgr) {
	for {
		time.Sleep(15 * time.Second)
		r.verifyCbrUpdates()
		if !r.isSyncer() {
			continue
		}
		// find all events need to be sent out, batch into one msg
		msg := &cbrtypes.MsgOnchainManyEvents{
			Creator: r.Transactor.Key.GetAddress().String(),
		}
		for chid, onech := range cbrMgr {
			// go over each chain db events, send msg
			msg.Events = append(msg.Events, onech.pullEvents(chid)...)
		}
		if len(msg.Events) > 0 {
			// or we should call cbridge grpc here?
			r.Transactor.AddTxMsg(msg)
		}
	}
}

func (c *CbrOneChain) pullEvents(chid uint64) []*cbrtypes.MsgOnchainEvent {
	var ret []*cbrtypes.MsgOnchainEvent
	// 1st loop over event names, then go over iter
	for _, evn := range evNames {
		iter, err := c.db.ReverseIterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn)))
		if err != nil {
			log.Errorln("chainID:", chid, evn, "iter err:", err)
			continue
		}
		for ; iter.Valid(); iter.Next() {
			// creator is not needed as MsgOnchainManyEvents has it
			ret = append(ret, &cbrtypes.MsgOnchainEvent{
				Chainid: chid,
				Evtype:  evn,
				Elog:    iter.Value(),
			})
		}
		iter.Close()
	}
	return ret
}

func (r *Relayer) verifyCbrUpdates() {}
