package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

	cbrflowtypes "github.com/celer-network/cbridge-flow/types"
	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// data is serialized OnChainEvent
func (k Keeper) ApplyEvent(ctx sdk.Context, data []byte) (bool, error) {
	onchev := new(cbrtypes.OnChainEvent)
	err := onchev.Unmarshal(data)
	if err != nil {
		return false, err
	}
	elog := new(ethtypes.Log) // only parse elog for evm
	if !commontypes.IsNonEvm(onchev.Chainid) {
		err = json.Unmarshal(onchev.Elog, elog)
		if err != nil {
			return false, err
		}
	}
	switch onchev.Evtype {
	case types.PegbrEventDeposited:
		depositChainId := onchev.Chainid
		var depositToken string // must be string as flow token is like A.xxxx.SomeToken. for evm, it's hex string
		var ev *eth.OriginalTokenVaultV2Deposited
		var version uint32 // only needed for evm chains, will be set by GetVaultVersion
		if commontypes.IsFlowChain(depositChainId) {
			// onchev.Elog is json serialized cbridge-flow/types/FlowMonitorLog so we deserialize and assign to ev fields
			flev := new(cbrflowtypes.FlowMonitorLog)
			err = json.Unmarshal(onchev.Elog, flev)
			if err != nil {
				return false, err
			}
			flowEv := new(cbrflowtypes.FlowSafeBoxDeposited)
			err = json.Unmarshal(flev.Event, flowEv)
			if err != nil {
				return false, err
			}
			depositToken = flowEv.Token // human string
			// set ev fields to flowEv values
			ev = &eth.OriginalTokenVaultV2Deposited{}
			ev.SetByFlow(flowEv)
		} else {
			version, _ = k.GetVaultVersion(ctx, onchev.Chainid, elog.Address)
			if version == 0 {
				otvContract, _ := eth.NewOriginalTokenVaultFilterer(eth.ZeroAddr, nil)
				ev0, err := otvContract.ParseDeposited(*elog)
				if err != nil {
					return false, err
				}
				ev = &eth.OriginalTokenVaultV2Deposited{
					DepositId:   ev0.DepositId,
					Depositor:   ev0.Depositor,
					Token:       ev0.Token,
					Amount:      ev0.Amount,
					MintChainId: ev0.MintChainId,
					MintAccount: ev0.MintAccount,
				}
			} else {
				ptbContract, _ := eth.NewOriginalTokenVaultV2Filterer(eth.ZeroAddr, nil)
				ev, err = ptbContract.ParseDeposited(*elog)
				if err != nil {
					return false, err
				}
			}
			depositToken = eth.Addr2Hex(ev.Token)
		}

		if k.HasDepositInfo(ctx, ev.DepositId) {
			log.Infof("skip already applied pegbr deposit event. chainid %d depositId %x version %d", depositChainId, ev.DepositId, version)
			return false, nil
		}

		mintId, mintAmount, baseFee, percFee, refundMsg, err := k.pegMint(
			ctx, depositChainId, depositChainId, ev.MintChainId, ev.Depositor, ev.MintAccount, depositToken, ev.Amount, ev.DepositId)
		if err != nil {
			return false, err
		}
		if refundMsg != "" {
			k.manageDataForDepositRefund(ctx, depositChainId, ev, version, depositToken)
			log.Warnf("deposit to be refunded, depositId:%x. %s", ev.DepositId, refundMsg)
			return true, nil
		}

		// Record DepositInfo
		depositInfo := types.DepositInfo{
			ChainId:      depositChainId,
			DepositId:    ev.DepositId[:],
			MintId:       mintId.Bytes(),
			VaultVersion: version,
		}
		k.SetDepositInfo(ctx, depositInfo)

		log.Infof("x/pegbr applied: %s. mintId: %x mintAmt: %s baseFee: %s percFee: %s",
			ev.PrettyLog(depositChainId), mintId, mintAmount, baseFee, percFee)
		return true, nil

	case types.PegbrEventBurn:
		burnChainId := onchev.Chainid
		var burnToken string // must be string as flow token is like A.xxxx.SomeToken. for evm, it's hex
		var ev *eth.PeggedTokenBridgeV2Burn
		var version uint32 // only needed for evm chains, will be set by GetBridgeVersion
		if commontypes.IsFlowChain(burnChainId) {
			// onchev.Elog is json serialized cbridge-flow/types/FlowMonitorLog so we deserialize and assign to ev fields
			flev := new(cbrflowtypes.FlowMonitorLog)
			err = json.Unmarshal(onchev.Elog, flev)
			if err != nil {
				return false, err
			}
			flowEv := new(cbrflowtypes.FlowPegBridgeBurn)
			err = json.Unmarshal(flev.Event, flowEv)
			if err != nil {
				return false, err
			}
			burnToken = flowEv.Token // human string
			// set ev fields to flowEv values
			ev = &eth.PeggedTokenBridgeV2Burn{}
			ev.SetByFlow(flowEv)
		} else {
			version, _ = k.GetBridgeVersion(ctx, onchev.Chainid, elog.Address)
			if version == 0 {
				ptbContract, _ := eth.NewPeggedTokenBridgeFilterer(eth.ZeroAddr, nil)
				ev0, err := ptbContract.ParseBurn(*elog)
				if err != nil {
					return false, err
				}
				ev = &eth.PeggedTokenBridgeV2Burn{
					BurnId:    ev0.BurnId,
					Token:     ev0.Token,
					Account:   ev0.Account,
					Amount:    ev0.Amount,
					ToAccount: ev0.WithdrawAccount,
				}
			} else {
				ptbContract, _ := eth.NewPeggedTokenBridgeV2Filterer(eth.ZeroAddr, nil)
				ev, err = ptbContract.ParseBurn(*elog)
				if err != nil {
					return false, err
				}
			}
			burnToken = eth.Addr2Hex(ev.Token)
		}

		if k.HasBurnInfo(ctx, ev.BurnId) {
			log.Infof("skip already applied pegbr burn event. burnChainId %d burnId %x version %d", burnChainId, ev.BurnId, version)
			return false, nil
		}

		pair, pairFound := k.GetOrigPeggedPairByPeggedByStrAddr(ctx, burnChainId, burnToken)
		if !pairFound {
			// pegged pair should be found. if not, an ERROR log would be printed.
			// this burn couldn't be refunded, because totalSupply in sgn was not updated.
			return false, fmt.Errorf("burn rejected, burnId:%x, pegged pair not exists, dstChainId %d, token %x", ev.BurnId, burnChainId, ev.Token)
		}

		var burnInfo types.BurnInfo
		var refundMsg string
		if ev.ToChainId == 0 || ev.ToChainId == pair.Orig.ChainId {
			// withdraw at the original token vault
			burnInfo, refundMsg, err = k.vaultWithdraw(ctx, pair, burnChainId, ev, version)
		} else {
			// mint at another chain
			burnInfo, refundMsg, err = k.burnMint(ctx, pair, burnChainId, ev, version)
		}
		if err != nil {
			return false, err
		}
		if refundMsg != "" {
			k.manageDataForBurnRefund(ctx, burnChainId, ev, version, burnToken)
			log.Warnf("burn to be refunded, burn:%x. %s", ev.BurnId, refundMsg)
			return true, nil
		}

		// Record BurnInfo
		k.SetBurnInfo(ctx, burnInfo)
		k.burnSupply(ctx, pair, burnChainId, ev.Token, ev.Amount, version)
		return true, nil

	case types.PegbrEventMint:
		mintChainId := onchev.Chainid
		ev := new(eth.PeggedTokenBridgeMint)
		if commontypes.IsFlowChain(mintChainId) {
			flev := new(cbrflowtypes.FlowMonitorLog)
			err = json.Unmarshal(onchev.Elog, flev)
			if err != nil {
				return false, err
			}
			flowEv := new(cbrflowtypes.FlowPegBridgeMint)
			err = json.Unmarshal(flev.Event, flowEv)
			if err != nil {
				return false, err
			}
			// only need to set ev.MintId as other fields are not needed in later code
			ev.MintId = flowEv.MintId
		} else {
			ptbContract, _ := eth.NewPeggedTokenBridgeFilterer(eth.ZeroAddr, nil)
			ev, err = ptbContract.ParseMint(*elog)
			if err != nil {
				return false, err
			}
		}
		mintInfo, found := k.GetMintInfo(ctx, ev.MintId)
		if !found {
			log.Errorln("x/pegbr mint info not found", ev.PrettyLog(onchev.Chainid))
			return false, nil
		}
		mintInfo.Success = true
		k.SetMintInfo(ctx, ev.MintId, mintInfo)
		log.Infoln("x/pegbr applied:", ev.PrettyLog(onchev.Chainid))
		return true, nil

	case types.PegbrEventWithdrawn:
		wdChainId := onchev.Chainid
		ev := new(eth.OriginalTokenVaultWithdrawn)
		if commontypes.IsFlowChain(wdChainId) {
			flev := new(cbrflowtypes.FlowMonitorLog)
			err = json.Unmarshal(onchev.Elog, flev)
			if err != nil {
				return false, err
			}
			flowEv := new(cbrflowtypes.FlowSafeBoxWithdrawn)
			err = json.Unmarshal(flev.Event, flowEv)
			if err != nil {
				return false, err
			}
			// only need to set ev.WithdrawId as other fields are not needed in later code
			ev.WithdrawId = flowEv.WithdrawId
		} else {
			otvContract, _ := eth.NewOriginalTokenVaultFilterer(eth.ZeroAddr, nil)
			ev, err = otvContract.ParseWithdrawn(*elog)
			if err != nil {
				return false, err
			}
		}
		wdInfo, found := k.GetWithdrawInfo(ctx, ev.WithdrawId)
		if !found {
			log.Errorln("x/pegbr withdraw info not found", ev.PrettyLog(wdChainId))
			return false, nil
		}
		wdInfo.Success = true
		k.SetWithdrawInfo(ctx, ev.WithdrawId, wdInfo)
		log.Infoln("x/pegbr applied:", ev.PrettyLog(wdChainId))
		return true, nil
	}

	return true, nil
}

func (k Keeper) pegMint(
	ctx sdk.Context, depositChainId, refChainId, mintChainId uint64,
	depositor, mintAccount eth.Addr, token string, amount *big.Int, refId eth.Hash) (
	mintId eth.Hash, mintAmount, baseFee, percFee *big.Int, refundMsg string, err error) {

	pair, found := k.GetOrigPeggedPair(ctx, depositChainId, token, mintChainId)
	if !found {
		err = fmt.Errorf("pegged pair not exists, srcChainId %d, dstChainId %d, token %x", depositChainId, mintChainId, token)
		return
	}
	mintAmount, baseFee, percFee = k.CalcAmountAndFees(ctx, pair, amount, true /* isPeggedDest */)
	if mintAmount.Sign() <= 0 {
		refundMsg = fmt.Sprintf("amount too small to cover fees, mintAmount %s baseFee %s percFee %s", mintAmount, baseFee, percFee)
		return
	}
	mintToken := eth.Hex2Addr(pair.Pegged.Address)
	// refChainId is deposit chain ID, refId is deposit ID
	if pair.BridgeVersion == 0 {
		mintId = types.CalcMintId(mintAccount, mintToken, mintAmount, depositor, refChainId, refId)
	} else {
		bridgeAddr, found := k.GetVersionedBridge(ctx, mintChainId, pair.BridgeVersion)
		if !found {
			err = fmt.Errorf("versioned bridge not found %d %d", mintChainId, pair.BridgeVersion)
			return
		}
		mintId = types.CalcMintIdV2(mintAccount, mintToken, mintAmount, depositor, refChainId, refId, bridgeAddr)
	}

	// get supplyCap
	supplyCap := new(big.Int).SetInt64(0)
	if pair.SupplyCap != "" {
		// supply cap string was checked during config set
		supplyCap.SetString(pair.SupplyCap, 10)
	}
	// a zero supplyCap indicates infinite supply.
	// a negative supplyCap indicates a special mod of burn ONLY, NO mint.
	if supplyCap.Sign() == 0 {
		// do nothing
	} else if supplyCap.Sign() == -1 {
		refundMsg = fmt.Sprintf("burn ONLY, negative supply cap %s", supplyCap)
		return
	} else {
		// get totalSupply
		beforeMintTotalSupply, found := k.GetTotalSupply(ctx, mintChainId, mintToken)
		if !found {
			beforeMintTotalSupply = new(big.Int).SetInt64(0)
		}
		// check if mint would exceed supply cap
		afterMintTotalSupply := new(big.Int).Add(beforeMintTotalSupply, mintAmount)
		if supplyCap.Cmp(afterMintTotalSupply) == -1 {
			// in reason of big mint amount that would exceed the supply cap, this deposit would be refunded
			refundMsg = fmt.Sprintf("hits supply cap, mintAmount %s current totalSupply %s supplyCap %s",
				mintAmount, beforeMintTotalSupply, supplyCap)
			return
		}
		// reset totalSupply
		k.SetTotalSupply(ctx, mintChainId, mintToken, afterMintTotalSupply)
	}

	err = k.MintFeeAndSendToSyncer(ctx, pair.Orig, baseFee, depositChainId, mintChainId)
	if err != nil {
		return
	}
	_, err = k.MintFee(ctx, pair.Orig, percFee)
	if err != nil {
		return
	}

	// Record MintInfo
	mint := types.MintOnChain{
		Token:      mintToken.Bytes(),
		Account:    mintAccount.Bytes(),
		Amount:     mintAmount.Bytes(),
		Depositor:  depositor.Bytes(),
		RefChainId: refChainId,
		RefId:      refId.Bytes(),
	}

	var mintProtoBytes []byte // to be serlized after set Token field
	// refChainId is burn chain ID, refId is burn ID
	if commontypes.IsFlowChain(mintChainId) {
		mint.Token = []byte(pair.Pegged.Address)
		mint.Account = mint.Account[12:] // only last 8 bytes for Flow account
		mintProtoBytes, _ = mint.Marshal()
		mintId = types.CalcFlowMintId(mintProtoBytes)
	} else {
		mint.Token = mintToken.Bytes()
		mintProtoBytes, _ = mint.Marshal()
	}

	mintInfo := types.MintInfo{
		ChainId:        mintChainId,
		MintProtoBytes: mintProtoBytes,
		Signatures:     make([]commontypes.Signature, 0),
		BaseFee:        baseFee.String(),
		PercentageFee:  percFee.String(),
		LastReqTime:    ctx.BlockTime().Unix(),
		BridgeVersion:  pair.BridgeVersion,
	}
	k.SetMintInfo(ctx, mintId, mintInfo)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMintToSign,
		sdk.NewAttribute(types.AttributeKeyMintId, mintId.Hex()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))

	return
}

func (k Keeper) vaultWithdraw(
	ctx sdk.Context, pair types.OrigPeggedPair, burnChainId uint64, ev *eth.PeggedTokenBridgeV2Burn, bridgeVersion uint32) (
	burnInfo types.BurnInfo, refundMsg string, err error) {

	withdrawAmt, baseFee, percFee := k.CalcAmountAndFees(ctx, pair, ev.Amount, false /* isPeggedDest */)
	if withdrawAmt.Sign() <= 0 {
		// in reason of too small burn amount, this burn would be refunded
		refundMsg = fmt.Sprintf("amount too small to cover fees, withdrawAmt %s baseFee %s percFee %s", withdrawAmt, baseFee, percFee)
		return
	}

	wdTokenAddr := pair.Orig.Address
	withdrawChainId := pair.Orig.ChainId
	var withdrawId eth.Hash
	// refChainId is burn chain ID, refId is burn ID
	if pair.VaultVersion == 0 {
		withdrawId = types.CalcWithdrawId(
			ev.ToAccount, eth.Hex2Addr(wdTokenAddr), withdrawAmt, ev.Account, burnChainId, ev.BurnId)
	} else {
		vaultAddr, found := k.GetVersionedVault(ctx, withdrawChainId, pair.VaultVersion)
		if !found {
			err = fmt.Errorf("versioned vault not found %d %d", withdrawChainId, pair.VaultVersion)
			return
		}
		withdrawId = types.CalcWithdrawIdV2(
			ev.ToAccount, eth.Hex2Addr(wdTokenAddr), withdrawAmt, ev.Account, burnChainId, ev.BurnId, vaultAddr)
	}

	// Record WithdrawInfo
	withdraw := types.WithdrawOnChain{
		// token is set later
		Receiver:    ev.ToAccount.Bytes(),
		Amount:      withdrawAmt.Bytes(),
		BurnAccount: ev.Account.Bytes(),
		RefChainId:  burnChainId,
		RefId:       ev.BurnId[:],
	}

	var withdrawProtoBytes []byte // to be serlized after set Token field
	// refChainId is burn chain ID, refId is burn ID
	if commontypes.IsFlowChain(withdrawChainId) {
		withdraw.Token = []byte(wdTokenAddr)
		withdraw.Receiver = withdraw.Receiver[12:] // only last 8 bytes for Flow account
		withdrawProtoBytes, _ = withdraw.Marshal()
		withdrawId = types.CalcFlowWithdrawId(withdrawProtoBytes)
	} else {
		withdraw.Token = eth.Hex2Bytes(wdTokenAddr)
		withdrawProtoBytes, _ = withdraw.Marshal()
	}

	wdInfo := types.WithdrawInfo{
		ChainId:            withdrawChainId,
		WithdrawProtoBytes: withdrawProtoBytes,
		Signatures:         make([]commontypes.Signature, 0),
		BaseFee:            baseFee.String(),
		PercentageFee:      percFee.String(),
		LastReqTime:        ctx.BlockTime().Unix(),
		VaultVersion:       pair.VaultVersion,
	}
	k.SetWithdrawInfo(ctx, withdrawId, wdInfo)

	// Mint fees to distribution module
	// NOTE: pegbridge fees are always claimed in the form of original tokens
	err = k.MintFeeAndSendToSyncer(ctx, pair.Orig, baseFee, burnChainId, withdrawChainId)
	if err != nil {
		return
	}
	_, err = k.MintFee(ctx, pair.Orig, percFee)
	if err != nil {
		return
	}

	burnInfo = types.BurnInfo{
		ChainId:       burnChainId,
		BurnId:        ev.BurnId[:],
		WithdrawId:    withdrawId.Bytes(),
		BridgeVersion: bridgeVersion,
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeWithdrawToSign,
		sdk.NewAttribute(types.AttributeKeyWithdrawId, withdrawId.Hex()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))

	log.Infof("x/pegbr applied: %s. withdrawId: %x withdrawAmt: %s baseFee: %s percFee: %s",
		ev.PrettyLog(burnChainId), withdrawId, withdrawAmt, baseFee, percFee)
	return
}

func (k Keeper) burnMint(
	ctx sdk.Context, pair types.OrigPeggedPair, burnChainId uint64, ev *eth.PeggedTokenBridgeV2Burn, bridgeVersion uint32) (
	burnInfo types.BurnInfo, refundMsg string, err error) {
	// first convert amount to original vault chain without fee
	// then apply logics for depositing at the original vault chain, and mint to another chain
	origChainAmt, _, _ := ConvertDestAmt(pair, ev.Amount, false)
	var mintId eth.Hash
	var mintAmount, baseFee, percFee *big.Int
	mintId, mintAmount, baseFee, percFee, refundMsg, err = k.pegMint(ctx, pair.Orig.ChainId, burnChainId, ev.ToChainId, ev.Account,
		ev.ToAccount, pair.Orig.Address, origChainAmt, ev.BurnId)
	if err != nil || refundMsg != "" {
		return
	}

	burnInfo = types.BurnInfo{
		ChainId:       burnChainId,
		BurnId:        ev.BurnId[:],
		MintId:        mintId.Bytes(),
		BridgeVersion: bridgeVersion,
	}
	log.Infof("x/pegbr applied: %s. mintId: %x mintAmt: %s baseFee: %s percFee: %s",
		ev.PrettyLog(burnChainId), mintId, mintAmount, baseFee, percFee)
	return
}

func (k Keeper) burnSupply(
	ctx sdk.Context, pair types.OrigPeggedPair, burnChainId uint64, token eth.Addr, amount *big.Int, version uint32) {
	// only update supply for latest version
	if pair.BridgeVersion == version {
		supplyCap := new(big.Int).SetInt64(0)
		if pair.SupplyCap != "" {
			// supply cap string was checked during config set
			supplyCap.SetString(pair.SupplyCap, 10)
		}
		// a zero supplyCap indicates infinite supply.
		// a negative supplyCap indicates a special mod of burn ONLY, NO mint.
		if supplyCap.Sign() == 0 {
			// do nothing
		} else {
			beforeBurnTotalSupply, found := k.GetTotalSupply(ctx, burnChainId, token)
			if !found {
				beforeBurnTotalSupply = new(big.Int).SetInt64(0)
			}
			afterBurnTotalSupply := new(big.Int).Sub(beforeBurnTotalSupply, amount)
			// if total supply after this burn would be negative, we'll reset it to zero instead of return an error.
			// this case would happen when total supply has not yet been set or is incorrectly set.
			// when pegbr on prod works well for a certain time and all pegged pairs' supplyCap and totalSupply
			// have been correctly set, we can remove the logic of resetting totalSupply to zero.
			if afterBurnTotalSupply.Sign() == -1 {
				afterBurnTotalSupply.SetInt64(0)
			}
			k.SetTotalSupply(ctx, burnChainId, token, afterBurnTotalSupply)
		}
	}
}

// for flow, we need depoToken string because ev.Token is empty
func (k Keeper) manageDataForDepositRefund(
	ctx sdk.Context, depositChainId uint64, ev *eth.OriginalTokenVaultV2Deposited, version uint32, depoToken string) {
	// Record a DepositInfo without mintId
	depositInfo := types.DepositInfo{
		ChainId:      depositChainId,
		DepositId:    ev.DepositId[:],
		MintId:       []byte{},
		VaultVersion: version,
	}
	k.SetDepositInfo(ctx, depositInfo)
	// Record a depositRefund: withdrawOnChain
	wdOnChain := types.WithdrawOnChain{
		Token:       ev.Token.Bytes(),
		Receiver:    ev.Depositor.Bytes(),
		Amount:      ev.Amount.Bytes(),
		BurnAccount: eth.ZeroAddr.Bytes(),
		RefChainId:  depositChainId,
		RefId:       ev.DepositId[:],
	}
	if commontypes.IsFlowChain(depositChainId) {
		wdOnChain.Token = []byte(depoToken) // use human string as is
	}
	k.SetDepositRefund(ctx, ev.DepositId, wdOnChain)
}

func (k Keeper) manageDataForBurnRefund(
	ctx sdk.Context, burnChainId uint64, ev *eth.PeggedTokenBridgeV2Burn, version uint32, burnToken string) {
	// Record a BurnInfo without withdrawId
	burnInfo := types.BurnInfo{
		ChainId:       burnChainId,
		BurnId:        ev.BurnId[:],
		WithdrawId:    []byte{},
		BridgeVersion: version,
	}
	k.SetBurnInfo(ctx, burnInfo)
	// Record a burnRefund: mintOnChain
	mintOnChain := types.MintOnChain{
		Token:      ev.Token.Bytes(),
		Account:    ev.Account.Bytes(),
		Amount:     ev.Amount.Bytes(),
		Depositor:  eth.ZeroAddr.Bytes(),
		RefChainId: burnChainId,
		RefId:      ev.BurnId[:],
	}
	if commontypes.IsFlowChain(burnChainId) {
		mintOnChain.Token = []byte(burnToken) // use human string as is
	}
	k.SetBurnRefund(ctx, ev.BurnId, mintOnChain)
}
