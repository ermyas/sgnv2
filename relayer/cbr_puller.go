package relayer

import (
	"time"

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
		// find all events need to be sent out, batch into one propose msg
		for chid, onech := range cbrMgr {
			// go over each chain db events, send msg
			onech.pullEvents(chid)
		}
	}
}

func (r *Relayer) sendCbrMsg() {}

func (c *CbrOneChain) pullEvents(chid uint64) {
	for _, evn := range evNames {
		c.db.ReverseIterator([]byte(evn), storetypes.PrefixEndBytes([]byte(evn)))

	}
}

func (r *Relayer) verifyCbrUpdates() {}
