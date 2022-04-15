package relayer

import (
	"encoding/json"
	"fmt"

	flowtypes "github.com/celer-network/cbridge-flow/types"
	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
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
	if commontypes.IsFlowChain(onchev.Chainid) {
		ev := new(flowtypes.FlowMonitorLog)
		err = json.Unmarshal(onchev.Elog, ev)
		if err != nil {
			log.Errorf("failed to unmarshal %x to FlowMonitorLog", onchev.Elog)
			return true, false
		}
		ok, err := r.cbrMgr[onchev.Chainid].fcc.VerifyEvent(ev)
		if err != nil {
			log.Error("flow verify event err: ", err)
			return false, false
		}
		return true, ok
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
		log.Debugf("skip pbr event: chain %x addr %x tx %x, reason: %s", onchev.Chainid, elog.Address, elog.TxHash, reason)
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
	if eLog.Address == c.pegContracts.vault.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.vault.ParseDeposited(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())
		if relayerInstance.isEthAddrBlocked(ev.Depositor, ev.MintAccount) {
			log.Warnf("%s, eth addrs blocked", logmsg)
			return true, false
		}

		done, approve, depositLog := c.verifyEventLog(
			eLog, eth.ContractTypePegVault, pegtypes.PegbrEventDeposited, c.pegContracts.vault.GetAddr(), logmsg)
		if depositLog == nil {
			return done, approve
		}
		depositEv, err := c.pegContracts.vault.ParseDeposited(*depositLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(depositEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, depositEv, depositEv.Raw, ev, ev.Raw)
			return true, false
		}
		return c.verifyOriginalTokenRecord(ev.DepositId, eLog.BlockNumber, cliCtx, logmsg, false)
	} else if eLog.Address == c.pegContracts.vault2.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.vault2.ParseDeposited(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())
		if relayerInstance.isEthAddrBlocked(ev.Depositor, ev.MintAccount) {
			log.Warnf("%s, eth addrs blocked", logmsg)
			return true, false
		}

		done, approve, depositLog := c.verifyEventLog(
			eLog, eth.ContractTypePegVaultV2, pegtypes.PegbrEventDeposited, c.pegContracts.vault2.GetAddr(), logmsg)
		if depositLog == nil {
			return done, approve
		}
		depositEv, err := c.pegContracts.vault2.ParseDeposited(*depositLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(depositEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, depositEv, depositEv.Raw, ev, ev.Raw)
			return true, false
		}

		depositId := pegtypes.CalcDepositIdV2(
			ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount, ev.Nonce, c.chainid, eLog.Address)
		if depositId != ev.DepositId {
			log.Errorf("%s. mismatch depositId ev has %x, calc: %x", logmsg, ev.DepositId, depositId)
			return true, false
		}
		return c.verifyOriginalTokenRecord(ev.DepositId, eLog.BlockNumber, cliCtx, logmsg, true)
	}

	log.Errorf("chain %d invalid deposit event address %x", c.chainid, eLog.Address)
	return true, false
}

func (c *CbrOneChain) verifyPegbrWithdrawn(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	if eLog.Address == c.pegContracts.vault.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.vault.ParseWithdrawn(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

		done, approve, withdrawLog := c.verifyEventLog(
			eLog, eth.ContractTypePegVault, pegtypes.PegbrEventWithdrawn, c.pegContracts.vault.GetAddr(), logmsg)
		if withdrawLog == nil {
			return done, approve
		}

		withdrawEv, err := c.pegContracts.vault.ParseWithdrawn(*withdrawLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(withdrawEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, withdrawEv, withdrawEv.Raw, ev, ev.Raw)
			return true, false
		}

		withdrawId := pegtypes.CalcWithdrawId(ev.Receiver, ev.Token, ev.Amount, ev.BurnAccount, ev.RefChainId, ev.RefId)
		if withdrawId != ev.WithdrawId {
			log.Errorf("%s. mismatch withdrawId ev has %x, calc: %x", logmsg, ev.WithdrawId, withdrawId)
			return true, false
		}
		return c.verifyOriginalTokenRecord(ev.WithdrawId, eLog.BlockNumber, cliCtx, logmsg, false)
	} else if eLog.Address == c.pegContracts.vault2.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.vault2.ParseWithdrawn(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

		done, approve, withdrawLog := c.verifyEventLog(
			eLog, eth.ContractTypePegVaultV2, pegtypes.PegbrEventWithdrawn, c.pegContracts.vault2.GetAddr(), logmsg)
		if withdrawLog == nil {
			return done, approve
		}
		withdrawEv, err := c.pegContracts.vault2.ParseWithdrawn(*withdrawLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(withdrawEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, withdrawEv, withdrawEv.Raw, ev, ev.Raw)
			return true, false
		}

		withdrawId := pegtypes.CalcWithdrawIdV2(
			ev.Receiver, ev.Token, ev.Amount, ev.BurnAccount, ev.RefChainId, ev.RefId, eLog.Address)
		if withdrawId != ev.WithdrawId {
			log.Errorf("%s. mismatch withdrawId ev has %x, calc: %x", logmsg, ev.WithdrawId, withdrawId)
			return true, false
		}
		return c.verifyOriginalTokenRecord(ev.WithdrawId, eLog.BlockNumber, cliCtx, logmsg, true)
	}

	log.Errorf("chain %d invalid withdraw event address %x", c.chainid, eLog.Address)
	return true, false
}

func (c *CbrOneChain) verifyPegbrBurn(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	if eLog.Address == c.pegContracts.bridge.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.bridge.ParseBurn(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())
		if relayerInstance.isEthAddrBlocked(ev.Account, ev.WithdrawAccount) {
			log.Warnf("%s, eth addrs blocked", logmsg)
			return true, false
		}

		done, approve, burnLog := c.verifyEventLog(
			eLog, eth.ContractTypePegBridge, pegtypes.PegbrEventBurn, c.pegContracts.bridge.GetAddr(), logmsg)
		if burnLog == nil {
			return done, approve
		}
		burnEv, err := c.pegContracts.bridge.ParseBurn(*burnLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(burnEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, burnEv, burnEv.Raw, ev, ev.Raw)
			return true, false
		}
		return c.verifyPeggedTokenRecord(ev.BurnId, eLog.BlockNumber, cliCtx, logmsg, false)

	} else if eLog.Address == c.pegContracts.bridge2.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.bridge2.ParseBurn(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())
		if relayerInstance.isEthAddrBlocked(ev.Account, ev.ToAccount) {
			log.Warnf("%s, eth addrs blocked", logmsg)
			return true, false
		}

		done, approve, burnLog := c.verifyEventLog(
			eLog, eth.ContractTypePegBridgeV2, pegtypes.PegbrEventBurn, c.pegContracts.bridge2.GetAddr(), logmsg)
		if burnLog == nil {
			return done, approve
		}
		burnEv, err := c.pegContracts.bridge2.ParseBurn(*burnLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(burnEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, burnEv, burnEv.Raw, ev, ev.Raw)
			return true, false
		}

		burnId := pegtypes.CalcBurnIdV2(ev.Account, ev.Token, ev.Amount, ev.ToChainId, ev.ToAccount, ev.Nonce, c.chainid, eLog.Address)
		if burnId != ev.BurnId {
			log.Errorf("%s. mismatch burnId ev has %x, calc: %x", logmsg, ev.BurnId, burnId)
			return true, false
		}
		return c.verifyPeggedTokenRecord(ev.BurnId, eLog.BlockNumber, cliCtx, logmsg, true)
	}

	log.Errorf("chain %d invalid burn event address %x", c.chainid, eLog.Address)
	return true, false
}

func (c *CbrOneChain) verifyPegbrMint(eLog *ethtypes.Log, cliCtx client.Context, logmsg string) (done, approve bool) {
	if eLog.Address == c.pegContracts.bridge.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.bridge.ParseMint(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

		done, approve, mintLog := c.verifyEventLog(
			eLog, eth.ContractTypePegBridge, pegtypes.PegbrEventMint, c.pegContracts.bridge.GetAddr(), logmsg)
		if mintLog == nil {
			return done, approve
		}
		mintEv, err := c.pegContracts.bridge.ParseMint(*mintLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(mintEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, mintEv, mintEv.Raw, ev, ev.Raw)
			return true, false
		}

		mintId := pegtypes.CalcMintId(ev.Account, ev.Token, ev.Amount, ev.Depositor, ev.RefChainId, ev.RefId)
		if mintId != ev.MintId {
			log.Errorf("%s. mismatch mintId ev has %x, calc: %x", logmsg, ev.MintId, mintId)
			return true, false
		}
		return c.verifyPeggedTokenRecord(ev.MintId, eLog.BlockNumber, cliCtx, logmsg, false)
	} else if eLog.Address == c.pegContracts.bridge2.GetAddr() && eLog.Address != eth.ZeroAddr {
		ev, err := c.pegContracts.bridge2.ParseMint(*eLog)
		if err != nil {
			log.Errorf("%s. parse eLog error %s", logmsg, err)
			return true, false
		}
		logmsg = fmt.Sprintf("%s. %s", logmsg, ev.String())

		done, approve, mintLog := c.verifyEventLog(
			eLog, eth.ContractTypePegBridgeV2, pegtypes.PegbrEventMint, c.pegContracts.bridge2.GetAddr(), logmsg)
		if mintLog == nil {
			return done, approve
		}
		mintEv, err := c.pegContracts.bridge2.ParseMint(*mintLog)
		if err != nil {
			log.Errorln(logmsg, "parse log err:", err)
			return true, false
		}
		if !ev.Equal(mintEv) {
			log.Errorf("%s. ev not equal. got: %s %v. expect: %s %v", logmsg, mintEv, mintEv.Raw, ev, ev.Raw)
			return true, false
		}

		mintId := pegtypes.CalcMintIdV2(ev.Account, ev.Token, ev.Amount, ev.Depositor, ev.RefChainId, ev.RefId, eLog.Address)
		if mintId != ev.MintId {
			log.Errorf("%s. mismatch mintId ev has %x, calc: %x", logmsg, ev.MintId, mintId)
			return true, false
		}
		return c.verifyPeggedTokenRecord(ev.MintId, eLog.BlockNumber, cliCtx, logmsg, true)
	}

	log.Errorf("chain %d invalid mint event address %x", c.chainid, eLog.Address)
	return true, false
}

func (c *CbrOneChain) verifyPeggedTokenRecord(recordId eth.Hash, blockNumber uint64, cliCtx client.Context, logmsg string, v2 bool) (done, approve bool) {
	// event log and block delay already checked, so everything should be valid,
	// continue to check the onchain state again for extra safety
	// the following checks should never fail in normal cases
	var exist bool
	var err error
	if v2 {
		exist, err = c.pegContracts.bridge2.Records(nil, recordId)
	} else {
		exist, err = c.pegContracts.bridge.Records(nil, recordId)
	}

	if err != nil {
		log.Warnf("%s. query burn records err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		// doesn't exist, vote no
		log.Errorln(logmsg, "record id not found")
		return true, false
	}
	log.Infof("%s, success", logmsg)
	return true, true
}

func (c *CbrOneChain) verifyOriginalTokenRecord(recordId eth.Hash, blockNumber uint64, cliCtx client.Context, logmsg string, v2 bool) (done, approve bool) {
	// event log and block delay already checked, so everything should be valid,
	// continue to check the onchain state again for extra safety
	// the following checks should never fail in normal cases
	var exist bool
	var err error
	if v2 {
		exist, err = c.pegContracts.vault2.Records(nil, recordId)
	} else {
		exist, err = c.pegContracts.vault.Records(nil, recordId)
	}
	if err != nil {
		log.Warnf("%s. query deposit records err: %s", logmsg, err)
		return false, false
	}
	if !exist {
		// deposit doesn't exist, vote no
		log.Errorln(logmsg, "record id not found")
		return true, false
	}
	// now both latest and safeblk has the state, ok to vote yes
	log.Infof("%s, success", logmsg)
	return true, true
}
