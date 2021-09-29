package relayer

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

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
	retry, err := r.cbrMgr[onchev.Chainid].checkEventOnChain(onchev.Evtype, elog)
	if err != nil {
		if retry {
			return false, false // onchain error, so don't vote
		} else {
			return true, false
		}

	}
	// event is the same as onchain, now move on to per event logic
	// eg. query x/cbridge if this event has already been handled
	// but if we only query applied state, we may still vote again?
	// unless we can query x/cbridge w/ state plus pending?
	// TODO: query x/cbridge
	return r.cbrMgr[onchev.Chainid].checkEventInStore(onchev.Evtype, onchev.Chainid, elog, r.Transactor.CliCtx)
}

// query chain to verify event is the same, return err if mismatch
// TODO: impl logic
func (c *CbrOneChain) checkEventOnChain(evtype string, eLog *ethtypes.Log) (retry bool, err error) {
	switch evtype {
	case CbrEventLiqAdd:
		return false, nil
	case CbrEventSend:
		return false, nil
	case CbrEventRelay:
		return false, nil
	case CbrEventSignersUpdated:
		ev, err := c.contract.ParseSignersUpdated(*eLog)
		if err != nil {
			return false, err
		}
		ssHash, err := c.contract.SsHash(&bind.CallOpts{})
		if err != nil {
			return true, err
		}
		if eth.Bytes2Hash(crypto.Keccak256(ev.CurSigners)) != ssHash {
			return false, fmt.Errorf("ssHash not match onchain value")
		}
		return false, nil
	}
	return false, fmt.Errorf("invalid event type %s", evtype)
}

func (c *CbrOneChain) checkEventInStore(
	evtype string, chainId uint64, eLog *ethtypes.Log, cliCtx client.Context) (done, approve bool) {

	switch evtype {
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
	case CbrEventSignersUpdated:
		ev, err := c.contract.ParseSignersUpdated(*eLog)
		if err != nil {
			return true, false
		}
		storedChainSigners, err := cbrcli.QueryChainSigners(cliCtx, chainId)
		if err == nil {
			if bytes.Compare(storedChainSigners.GetSignersBytes(), ev.CurSigners) == 0 {
				return true, false
			}
		}
		return true, true
	}
	return true, false
}
