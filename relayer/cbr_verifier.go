package relayer

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// to be called by r.verifyUpdate
// decode event and check if I also have event in db
// TODO: query x/cbridge to make sure event not processed
func (r *Relayer) verifyCbrEventUpdate(update *synctypes.PendingUpdate) (done, approve bool) {
	onchev := new(cbrtypes.OnChainEvent)
	err := onchev.Unmarshal(update.Data)
	if err != nil {
		log.Errorf("failed to unmarshal %x to onchain event msg", update.Data)
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

	logmsg := fmt.Sprintf("verify update %d cbr chain %d type %s", update.Id, onchev.Chainid, onchev.Evtype)
	switch onchev.Evtype {
	case cbrtypes.CbrEventLiqAdd:
		return r.cbrMgr[onchev.Chainid].verifyLiqAdd(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventSend:
		return r.cbrMgr[onchev.Chainid].verifySend(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventRelay:
		return r.cbrMgr[onchev.Chainid].verifyRelay(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventWithdraw:
		return r.cbrMgr[onchev.Chainid].verifyWithdraw(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventSignersUpdated:
		return r.cbrMgr[onchev.Chainid].verifySigners(elog, r.Transactor.CliCtx, logmsg)

	default:
		log.Errorf("%s. invalid type", logmsg)
		return true, false
	}
}

func (c *CbrOneChain) verifyLiqAdd(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.contract.ParseLiquidityAdded(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain

	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifySend(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.contract.ParseSend(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain

	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyRelay(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.contract.ParseRelay(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain

	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyWithdraw(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.contract.ParseWithdrawDone(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain

	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifySigners(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.contract.ParseSignersUpdated(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store
	storedChainSigners, err := cbrcli.QueryChainSigners(cliCtx, c.chainid)
	if err == nil {
		if equalSigners(storedChainSigners.GetSortedSigners(), ev) {
			log.Infof("%s. already updated", logmsg)
			return true, false
		}
	}

	// check on chain
	ssHash, err := c.contract.SsHash(&bind.CallOpts{})
	if err != nil {
		log.Errorf("%s. query ssHash err: %s", logmsg, err)
		return false, false
	}
	curssHash := eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(ev.Signers, ev.Powers)))
	if curssHash != ssHash {
		log.Errorf("%s. curss hash %x not match onchain values: %x", logmsg, curssHash, ssHash)
		return true, false
	}
	c.setCurssByEvent(ev)

	log.Infof("%s, success", logmsg)
	return true, true
}

func equalSigners(ss []*cbrtypes.Signer, ev *eth.BridgeSignersUpdated) bool {
	if len(ss) != len(ev.Signers) {
		return false
	}
	for i, s := range ss {
		if !bytes.Equal(s.Addr, ev.Signers[i].Bytes()) {
			return false
		}
		if !bytes.Equal(s.Power, ev.Powers[i].Bytes()) {
			return false
		}
	}
	return true
}
