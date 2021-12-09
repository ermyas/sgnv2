package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// to be called by r.verifyUpdate
// decode event and check if it matches onchain
func (r *Relayer) verifyPegbrEventUpdate(update *synctypes.PendingUpdate) (done, approve bool) {
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
	if elog == nil {
		log.Errorf("unmarshal %x to to nil", onchev.Elog)
		return true, false
	}

	cbrOneChain := r.cbrMgr[onchev.Chainid]
	if cbrOneChain == nil {
		log.Errorf("cbrMgr not finish initialization yet, updates from chain: %d", onchev.Chainid)
		return false, false
	}

	skip, reason := cbrOneChain.skipPegbrEvent(onchev.Evtype, elog, r.Transactor.CliCtx, nil)
	if skip {
		log.Debugf("skip pbr event: %s, reason: %s", string(onchev.Elog), reason)
		return true, false
	}

	logmsg := fmt.Sprintf("verify update %d cbr chain %d type %s", update.Id, onchev.Chainid, onchev.Evtype)
	switch onchev.Evtype {
	case pegtypes.PegbrEventDeposited:
		return cbrOneChain.verifyPegbrDeposit(elog, r.Transactor.CliCtx, logmsg)

	case pegtypes.PegbrEventBurn:
		return cbrOneChain.verifyPegbrBurn(elog, r.Transactor.CliCtx, logmsg)

	case pegtypes.PegbrEventMint:
		return cbrOneChain.verifyPegbrMint(elog, r.Transactor.CliCtx, logmsg)

	case pegtypes.PegbrEventWithdrawn:
		return cbrOneChain.verifyPegbrWithdrawn(elog, r.Transactor.CliCtx, logmsg)

	default:
		log.Errorf("%s. invalid type", logmsg)
		return true, false
	}
}

func (c *CbrOneChain) verifyPegbrDeposit(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.pegContracts.vault.ParseDeposited(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	return c.verifyOriginalTokenRecord(ev.DepositId, eLog.BlockNumber, cliCtx, logmsg)

}

func (c *CbrOneChain) verifyPegbrWithdrawn(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.pegContracts.vault.ParseWithdrawn(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	return c.verifyOriginalTokenRecord(ev.WithdrawId, eLog.BlockNumber, cliCtx, logmsg)
}

func (c *CbrOneChain) verifyPegbrBurn(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.pegContracts.bridge.ParseBurn(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	return c.verifyPeggedTokenRecord(ev.BurnId, eLog.BlockNumber, cliCtx, logmsg)
}

func (c *CbrOneChain) verifyPegbrMint(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.pegContracts.bridge.ParseMint(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	return c.verifyPeggedTokenRecord(ev.MintId, eLog.BlockNumber, cliCtx, logmsg)
}

func (c *CbrOneChain) verifyPeggedTokenRecord(recordId eth.Hash, blockNumber uint64, cliCtx client.Context, logmsg string) (done, approve bool) {
	// check on chain
	exist, err := c.pegContracts.bridge.Records(nil, recordId)
	if err != nil {
		log.Warnf("%s. query burn records err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		if c.mon.GetCurrentBlockNumber().Uint64() < blockNumber {
			log.Warnln(logmsg, "block number not passed", c.mon.GetCurrentBlockNumber(), blockNumber)
			return false, false
		}
		// burn doesn't exist, vote no
		log.Warnln(logmsg, "record id not found")
		return true, false
	}
	// latest has the state, now check if it has been long enough
	safeBlkNum := c.mon.GetCurrentBlockNumber().Uint64() - c.blkDelay
	exist, err = c.pegContracts.bridge.Records(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(safeBlkNum),
	}, recordId)
	if err != nil {
		log.Warnf("%s. query safe record err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		log.Infoln(logmsg, "record id not found in safeblk")
		return false, false
	}
	// now both latest and safeblk has the state, ok to vote yes
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyOriginalTokenRecord(recordId eth.Hash, blockNumber uint64, cliCtx client.Context, logmsg string) (done, approve bool) {
	// check on chain
	exist, err := c.pegContracts.vault.Records(nil, recordId)
	if err != nil {
		log.Warnf("%s. query deposit records err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		if c.mon.GetCurrentBlockNumber().Uint64() < blockNumber {
			log.Warnln(logmsg, "block number not passed", c.mon.GetCurrentBlockNumber(), blockNumber)
			return false, false
		}
		// deposit doesn't exist, vote no
		log.Warnln(logmsg, "record id not found")
		return true, false
	}
	// latest has the state, now check if it has been long enough
	safeBlkNum := c.mon.GetCurrentBlockNumber().Uint64() - c.blkDelay
	exist, err = c.pegContracts.vault.Records(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(safeBlkNum),
	}, recordId)
	if err != nil {
		log.Warnf("%s. query safe record err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		log.Infoln(logmsg, "record id not found in safeblk")
		return false, false
	}
	// now both latest and safeblk has the state, ok to vote yes
	log.Infof("%s, success", logmsg)
	return true, true
}
