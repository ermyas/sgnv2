package relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"

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
// decode event and check if it matches onchain
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
	if elog == nil {
		log.Errorf("unmarshal %x to to nil", onchev.Elog)
		return true, false
	}

	cbrOneChain := r.cbrMgr[onchev.Chainid]
	if cbrOneChain == nil {
		log.Errorf("cbrMgr not finish initialization yet, updates from chain: %d", onchev.Chainid)
		return false, false
	}

	skip, reason := cbrOneChain.skipEvent(onchev.Evtype, elog, r.Transactor.CliCtx, nil)
	if skip {
		log.Debugf("skip cbr event: %s, reason: %s", string(onchev.Elog), reason)
		return true, false
	}

	logmsg := fmt.Sprintf("verify update %d cbr chain %d type %s", update.Id, onchev.Chainid, onchev.Evtype)
	switch onchev.Evtype {
	case cbrtypes.CbrEventLiqAdd:
		return cbrOneChain.verifyLiqAdd(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventSend:
		return cbrOneChain.verifySend(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventRelay:
		return cbrOneChain.verifyRelay(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventWithdraw:
		return cbrOneChain.verifyWithdraw(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventSignersUpdated:
		return cbrOneChain.verifySigners(elog, r.Transactor.CliCtx, logmsg)

	case cbrtypes.CbrEventWithdrawalRequest:
		return cbrOneChain.verifyWithdrawalRequest(elog, r.Transactor.CliCtx, logmsg)

	default:
		log.Errorf("%s. invalid type", logmsg)
		return true, false
	}
}

func (r *Relayer) verifyUpdateCbrPrice(update *synctypes.PendingUpdate) (done, approve bool) {
	priceFromSyncer := new(cbrtypes.CbrPrice)
	err := priceFromSyncer.Unmarshal(update.Data)
	if err != nil {
		log.Errorln("failed to unmarshal ", update.Data, " to CbrPrice msg")
		return true, false
	}
	priceIGot, success := getCbrPriceFromUrl()
	if !success {
		log.Warnln("failed to get CbrPrice from s3. priceFromSyncer:", priceFromSyncer)
		return false, false
	}
	if priceIGot.GetUpdateEpoch() < priceFromSyncer.GetUpdateEpoch() {
		log.Warnln("price I got is older than price from syncer, price I got:", priceIGot.GetUpdateEpoch(),
			" price from syncer:", priceFromSyncer.GetUpdateEpoch())
		return false, false
	} else if priceIGot.GetUpdateEpoch() > priceFromSyncer.GetUpdateEpoch() {
		log.Warnln("price I got is newer than price from syncer, price I got:", priceIGot.GetUpdateEpoch(),
			" price from syncer:", priceFromSyncer.GetUpdateEpoch())
		return true, false
	}
	p1, _ := priceIGot.Marshal()
	if !bytes.Equal(p1, update.Data) {
		log.Errorln("price I got is different from price from syncer but has same update_epoch, price I got:", priceIGot,
			" price from syncer:", priceFromSyncer)
		return true, false
	}

	log.Infof("verifyUpdateCbrPrice success, %+v", priceFromSyncer)
	return true, true
}

func (c *CbrOneChain) verifyLiqAdd(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.cbrContract.ParseLiquidityAdded(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	done, approve, addLiqLog := c.verifyEventLog(eLog, eth.LiquidityBridge, cbrtypes.CbrEventLiqAdd, c.cbrContract.GetAddr(), logmsg)
	if addLiqLog == nil {
		return done, approve
	}
	addLiqEv, err := c.cbrContract.ParseLiquidityAdded(*addLiqLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	// now cmp ev and addLiqEv
	if !ev.Equal(addLiqEv) {
		log.Errorln(logmsg, "ev not equal. got:", addLiqEv.String(), "expect:", ev.String())
		return true, false
	}
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifySend(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.cbrContract.ParseSend(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	xferId := ev.CalcXferId(c.chainid)
	if xferId != ev.TransferId {
		log.Errorf("%s. mismatch xferid ev has %x, calc: %x", logmsg, ev.TransferId, xferId)
		return true, false
	}
	// we must check both latest and latest-blkdelay have the state
	// if only latest has, means too soon, if only latest-blkdelay has, means it has been reorg
	exist, err := c.cbrContract.Transfers(nil, xferId)
	if err != nil {
		log.Warnf("%s. query transfers err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		if c.mon.GetCurrentBlockNumber().Uint64() < eLog.BlockNumber {
			log.Warnln(logmsg, "xferId:", xferId.String(), "block number not passed", c.mon.GetCurrentBlockNumber(), eLog.BlockNumber)
			return false, false
		}
		// xfer doesn't exist, vote no
		log.Warnln(logmsg, "xferId:", xferId.String(), "not found")
		return true, false
	}
	// latest has the state, now check if it has been long enough
	safeBlkNum := c.mon.GetCurrentBlockNumber().Uint64() - c.blkDelay
	exist, err = c.cbrContract.Transfers(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(safeBlkNum),
	}, xferId)
	if err != nil {
		log.Warnf("%s. query safe transfers err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		// xfer doesn't exist in history, means too soon, allow retry later
		log.Infoln(logmsg, "xferId:", xferId.String(), "not found in safeblk")
		return false, false
	}
	// now both latest and safeblk has the state, ok to vote yes
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyRelay(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.cbrContract.ParseRelay(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	xferId := ev.CalcXferId(c.chainid)
	if xferId != ev.TransferId {
		log.Errorf("%s. mismatch xferid ev has %x, calc: %x", logmsg, ev.TransferId, xferId)
		return true, false
	}
	exist, err := c.cbrContract.Transfers(nil, xferId)
	if err != nil {
		log.Warnf("%s. query transfers err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		if c.mon.GetCurrentBlockNumber().Uint64() < eLog.BlockNumber {
			log.Warnln(logmsg, "xferId:", xferId.String(), "block number not passed", c.mon.GetCurrentBlockNumber(), eLog.BlockNumber)
			return false, false
		}
		// xfer doesn't exist, vote no
		log.Warnln(logmsg, "xferId:", xferId.String(), "not found")
		return true, false
	}
	// we don't do safeblk checking as this is event when money leaving the system, so it's safe
	// to be more acceptable of event
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyWithdraw(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.cbrContract.ParseWithdrawDone(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	wdId := ev.CalcWdID(c.chainid)
	if wdId != ev.WithdrawId {
		log.Errorf("%s mismatch wdid ev has %x, calc %x", logmsg, ev.WithdrawId, wdId)
		return true, false
	}
	exist, err := c.cbrContract.Withdraws(nil, wdId)
	if err != nil {
		log.Warnf("%s. query withdraws err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		if c.mon.GetCurrentBlockNumber().Uint64() < eLog.BlockNumber {
			log.Warnln(logmsg, "wdId:", wdId.String(), "block number not passed", c.mon.GetCurrentBlockNumber(), eLog.BlockNumber)
			return false, false
		}
		// wdid doesn't exist, vote no
		log.Warnln(logmsg, "wdId:", wdId.String(), "not found")
		return true, false
	}
	// we don't do safeblk checking as this is event when money leaving the system, so it's safe
	// to be more acceptable of event
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyEventLog(
	eLog *ethtypes.Log, ctype eth.ContractType, evName string, expAddr eth.Addr, logmsg string) (done, approve bool, resLog *ethtypes.Log) {

	receipt, err := c.TransactionReceipt(context.Background(), eLog.TxHash)
	if err != nil {
		log.Warnln(logmsg, "TransactionReceipt err:", err)
		return false, false, nil
	}

	resLog = eth.FindMatchContractEvent(ctype, evName, expAddr, receipt.Logs)

	if resLog == nil { // no match event in the tx, could be forged tx
		log.Errorln(logmsg, "no match event found in tx:", eLog.TxHash)
		return true, false, nil
	}
	if resLog.Removed {
		log.Errorln(logmsg, "log removed")
		return true, false, nil
	}
	// not possible as we check addr in FindMatchContractEvent, keep here for extra safety
	if resLog.Address != expAddr {
		log.Errorln(logmsg, "mismatch contract addr. log has:", resLog.Address, "expect:", expAddr)
		return true, false, nil
	}
	// check blocknumber and index because they are used in key
	if resLog.BlockNumber != eLog.BlockNumber {
		log.Errorln(logmsg, "mismatch blknum. proposal has:", eLog.BlockNumber, "log from receipt has:", resLog.BlockNumber)
		return true, false, nil
	}
	if resLog.Index != eLog.Index {
		log.Errorln(logmsg, "mismatch event index. proposal has:", eLog.Index, "log from receipt has:", resLog.Index)
		return true, false, nil
	}
	// make sure addLiqLog.BlockNumber isn't too recent
	blk := c.mon.GetCurrentBlockNumber().Uint64()
	if resLog.BlockNumber > blk-c.blkDelay {
		log.Warnf("%s evblk %d too soon, should only up to blk %d", logmsg, resLog.BlockNumber, blk-c.blkDelay)
		return false, false, nil
	}

	return true, true, resLog
}

func (c *CbrOneChain) verifySigners(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.cbrContract.ParseSignersUpdated(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store
	storedChainSigners, err := cbrcli.QueryChainSigners(cliCtx, c.chainid)
	if err == nil {
		if EqualSigners(storedChainSigners.GetSortedSigners(), ev) {
			log.Infof("%s. already updated", logmsg)
			return true, false
		}
	}

	// check on chain
	ssHash, err := c.cbrContract.SsHash(&bind.CallOpts{})
	if err != nil {
		log.Warnf("%s. query ssHash err: %s", logmsg, err)
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

func EqualSigners(ss []*cbrtypes.Signer, ev *eth.BridgeSignersUpdated) bool {
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

func (c *CbrOneChain) verifyWithdrawalRequest(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	// parse event
	ev, err := c.wdiContract.ParseWithdrawalRequest(*eLog)
	if err != nil {
		log.Errorf("%s. parse eLog error %s", logmsg, err)
		return true, false
	}
	logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

	// check in store

	// check on chain
	done, approve, wdReqLog := c.verifyEventLog(eLog, eth.WdInbox, cbrtypes.CbrEventWithdrawalRequest, c.wdiContract.GetAddr(), logmsg)
	if wdReqLog == nil {
		return done, approve
	}
	wdReqEv, err := c.wdiContract.ParseWithdrawalRequest(*wdReqLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	// now cmp ev and wdReqEv
	if !ev.Equal(wdReqEv) {
		log.Errorln(logmsg, "ev not equal. got:", wdReqEv.String(), "expect:", ev.String())
		return true, false
	}
	log.Infof("%s, success", logmsg)
	return true, true
}
