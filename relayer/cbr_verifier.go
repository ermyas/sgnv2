package relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// to be called by r.verifyUpdate
// decode event and check if it matches onchain
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
		log.Debugf("skip cbr event: chain %d addr %x tx %x, reason: %s", onchev.Chainid, elog.Address, elog.TxHash, reason)
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
	if relayerInstance.isEthAddrBlocked(ev.Provider) {
		log.Warnf("%s, eth addrs blocked", logmsg)
		return true, false
	}

	// check on chain
	done, approve, addLiqLog := c.verifyEventLog(eLog, eth.ContractTypeLiquidityBridge, cbrtypes.CbrEventLiqAdd, c.cbrContract.GetAddr(), logmsg)
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
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, addLiqEv, addLiqEv.Raw, ev, ev.Raw)
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
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnf("%s, eth addrs blocked", logmsg)
		return true, false
	}

	// check on chain
	done, approve, sendLog := c.verifyEventLog(eLog, eth.ContractTypeLiquidityBridge, cbrtypes.CbrEventSend, c.cbrContract.GetAddr(), logmsg)
	if sendLog == nil {
		return done, approve
	}
	sendEv, err := c.cbrContract.ParseSend(*sendLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	if !ev.Equal(sendEv) {
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, sendEv, sendEv.Raw, ev, ev.Raw)
		return true, false
	}

	// event log and block delay already checked, so everything should be valid,
	// continue to check the onchain state again for extra safety
	// the following checks should never fail in normal cases
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
		log.Warnln(logmsg, "send not found on chain")
		return false, false
	}

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

	// check on chain
	done, approve, sendLog := c.verifyEventLog(eLog, eth.ContractTypeLiquidityBridge, cbrtypes.CbrEventRelay, c.cbrContract.GetAddr(), logmsg)
	if sendLog == nil {
		return done, approve
	}
	relayEv, err := c.cbrContract.ParseRelay(*sendLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	if !ev.Equal(relayEv) {
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, relayEv, relayEv.Raw, ev, ev.Raw)
		return true, false
	}

	// event log and block delay already checked, so everything should be valid,
	// continue to check the onchain state again for extra safety
	// the following checks should never fail in normal cases
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
		log.Warnln(logmsg, "relay not found on chain")
		return false, false
	}

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

	// check on chain
	done, approve, sendLog := c.verifyEventLog(eLog, eth.ContractTypeLiquidityBridge, cbrtypes.CbrEventWithdraw, c.cbrContract.GetAddr(), logmsg)
	if sendLog == nil {
		return done, approve
	}
	withdrawEv, err := c.cbrContract.ParseWithdrawDone(*sendLog)
	if err != nil {
		log.Errorln(logmsg, "parse log err:", err)
		return true, false
	}
	if !ev.Equal(withdrawEv) {
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, withdrawEv, withdrawEv.Raw, ev, ev.Raw)
		return true, false
	}

	// event log and block delay already checked, so everything should be valid,
	// continue to check the onchain state again for extra safety
	// the following checks should never fail in normal cases
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
		// wdid doesn't exist, vote no
		log.Warnln(logmsg, "withdrawal not found on chain")
		return false, false
	}
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyEventLog(
	eLog *ethtypes.Log, ctype eth.ContractType, evName string, expAddr eth.Addr, logmsg string) (
	done, approve bool, resLog *ethtypes.Log) {

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
	blk := c.mon.GetBlkNum()
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
	if relayerInstance.isEthAddrBlocked(ev.Sender, ev.Receiver) {
		log.Warnf("%s, eth addrs blocked", logmsg)
		return true, false
	}

	// check on chain
	done, approve, wdReqLog := c.verifyEventLog(eLog, eth.ContractTypeWdInbox, cbrtypes.CbrEventWithdrawalRequest, c.wdiContract.GetAddr(), logmsg)
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
		log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, wdReqEv, wdReqEv.Raw, ev, ev.Raw)
		return true, false
	}
	log.Infof("%s, success", logmsg)
	return true, true
}
