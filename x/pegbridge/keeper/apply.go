package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

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
	elog := new(ethtypes.Log)
	err = json.Unmarshal(onchev.Elog, elog)
	if err != nil {
		return false, err
	}
	switch onchev.Evtype {
	case types.PegbrEventDeposited:
		otvContract, _ := eth.NewOriginalTokenVaultFilterer(eth.ZeroAddr, nil)
		ev, err := otvContract.ParseDeposited(*elog)
		if err != nil {
			return false, err
		}
		depositChainId := onchev.Chainid
		if k.HasDepositInfo(ctx, ev.DepositId) {
			log.Infof("skip already applied pegbr deposit event. chainid %d depositId %x", depositChainId, ev.DepositId)
			return false, nil
		}

		pair, found := k.GetOrigPeggedPair(ctx, depositChainId, ev.Token, ev.MintChainId)
		if !found {
			// in reason of invalid params, this deposit would be refunded
			k.manageDataForRefund(ctx, depositChainId, ev)
			return true, fmt.Errorf("pegged pair not exists")
		}
		mintAmount, baseFee, percFee := k.CalcAmountAndFees(ctx, pair, ev.Amount, true /* isPeggedDest */)
		if mintAmount.Sign() <= 0 {
			// TODO: Trigger refund, or just ignore?
			return false, fmt.Errorf("amount too small to cover fees, mintAmount %s baseFee %s percFee %s", mintAmount, baseFee, percFee)
		}
		// get supplyCap
		supplyCap := new(big.Int).SetInt64(0)
		if pair.SupplyCap != "" {
			// supply cap string was checked during config set
			supplyCap.SetString(pair.SupplyCap, 10)
		}
		// get totalSupply
		beforeMintTotalSupply, found := k.GetTotalSupply(ctx, depositChainId, ev.MintChainId, eth.Hex2Addr(pair.Pegged.Address))
		if !found {
			beforeMintTotalSupply = new(big.Int).SetInt64(0)
		}
		// check if mint would exceed supply cap
		afterMintTotalSupply := new(big.Int).Add(beforeMintTotalSupply, mintAmount)
		// a zero supplyCap means infinite supply. a negative supplyCap may mean a special mod of ONLY burn NO mint
		if supplyCap.Sign() != 0 && supplyCap.Cmp(afterMintTotalSupply) < 0 {
			// in reason of big mint amount that would exceed the supply cap, this deposit would be refunded
			k.manageDataForRefund(ctx, depositChainId, ev)
			return true, fmt.Errorf("ongoing mint would exceed supply cap, mintAmount %s current totalSupply %s supplyCap %s", mintAmount, beforeMintTotalSupply, supplyCap)
		}
		// reset totalSupply
		k.SetTotalSupply(ctx, depositChainId, ev.MintChainId, eth.Hex2Addr(pair.Pegged.Address), afterMintTotalSupply)

		mintTokenAddr := pair.Pegged.Address
		// refChainId is deposit chain ID, refId is deposit ID
		mintId := types.CalcMintId(
			ev.MintAccount, eth.Hex2Addr(mintTokenAddr), mintAmount, ev.Depositor, depositChainId, ev.DepositId)

		// Record DepositInfo
		depositInfo := types.DepositInfo{
			ChainId:   depositChainId,
			DepositId: ev.DepositId[:],
			MintId:    mintId.Bytes(),
		}
		k.SetDepositInfo(ctx, ev.DepositId, depositInfo)

		// Record MintInfo
		mint := types.MintOnChain{
			Token:      eth.Hex2Bytes(mintTokenAddr),
			Account:    ev.MintAccount.Bytes(),
			Amount:     mintAmount.Bytes(),
			Depositor:  ev.Depositor.Bytes(),
			RefChainId: depositChainId,
			RefId:      ev.DepositId[:],
		}
		mintProtoBytes, _ := mint.Marshal()
		mintInfo := types.MintInfo{
			ChainId:        ev.MintChainId,
			MintProtoBytes: mintProtoBytes,
			Signatures:     make([]commontypes.Signature, 0),
			BaseFee:        baseFee.String(),
			PercentageFee:  percFee.String(),
			LastReqTime:    ctx.BlockTime().Unix(),
		}
		k.SetMintInfo(ctx, mintId, mintInfo)

		// Mint fees to distribution module
		// NOTE: pegbridge fees are always claimed in the form of original tokens
		k.MintFee(ctx, pair.Orig, new(big.Int).Add(baseFee, percFee))

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeMintToSign,
			sdk.NewAttribute(types.AttributeKeyMintId, mintId.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))

		log.Infof("x/pegbr applied: %s. mintId: %x mintAmt: %s baseFee: %s percFee: %s",
			ev.PrettyLog(onchev.Chainid), mintId, mintAmount, baseFee, percFee)
		return true, nil
	case types.PegbrEventBurn:
		ptbContract, _ := eth.NewPeggedTokenBridgeFilterer(eth.ZeroAddr, nil)
		ev, err := ptbContract.ParseBurn(*elog)
		if err != nil {
			return false, err
		}
		burnChainId := onchev.Chainid
		if k.HasBurnInfo(ctx, ev.BurnId) {
			log.Infof("skip already applied pegbr burn event. chainId %d burnId %x", burnChainId, ev.BurnId)
			return false, nil
		}

		pair, pairFound := k.GetOrigPeggedPairByPegged(ctx, burnChainId, ev.Token)
		if !pairFound {
			return false, fmt.Errorf("pegged pair not exists")
		}

		withdrawAmt, baseFee, percFee := k.CalcAmountAndFees(ctx, pair, ev.Amount, false /* isPeggedDest */)
		if withdrawAmt.Sign() <= 0 {
			// TODO: Trigger refund, or just ignore?
			return false, fmt.Errorf("amount too small to cover fees, withdrawAmt %s baseFee %s percFee %s", withdrawAmt, baseFee, percFee)
		}

		// update totalSupply
		burnAmt := ev.Amount
		beforeBurnTotalSupply, found := k.GetTotalSupply(ctx, pair.Orig.ChainId, burnChainId, eth.Hex2Addr(pair.Pegged.Address))
		if !found {
			beforeBurnTotalSupply = new(big.Int).SetInt64(0)
		}
		afterBurnTotalSupply := new(big.Int).Sub(beforeBurnTotalSupply, burnAmt)
		// if total supply after this burn would be negative, we'll reset it to zero instead of return an error.
		// this case would happen when total supply has not yet been set or is incorrectly set.
		if afterBurnTotalSupply.Sign() == -1 {
			afterBurnTotalSupply.SetInt64(0)
		}
		k.SetTotalSupply(ctx, pair.Orig.ChainId, burnChainId, eth.Hex2Addr(pair.Pegged.Address), afterBurnTotalSupply)

		wdTokenAddr := pair.Orig.Address
		withdrawChainId := pair.Orig.ChainId
		// refChainId is burn chain ID, refId is burn ID
		withdrawId := types.CalcWithdrawId(
			ev.WithdrawAccount, eth.Hex2Addr(wdTokenAddr), withdrawAmt, ev.Account, burnChainId, ev.BurnId)

		// Record BurnInfo
		burnInfo := types.BurnInfo{
			ChainId:    burnChainId,
			BurnId:     ev.BurnId[:],
			WithdrawId: withdrawId.Bytes(),
		}
		k.SetBurnInfo(ctx, ev.BurnId, burnInfo)

		// Record WithdrawInfo
		withdraw := types.WithdrawOnChain{
			Token:       eth.Hex2Bytes(wdTokenAddr),
			Receiver:    ev.WithdrawAccount.Bytes(),
			Amount:      withdrawAmt.Bytes(),
			BurnAccount: ev.Account.Bytes(),
			RefChainId:  burnChainId,
			RefId:       ev.BurnId[:],
		}
		withdrawProtoBytes, _ := withdraw.Marshal()
		wdInfo := types.WithdrawInfo{
			ChainId:            withdrawChainId,
			WithdrawProtoBytes: withdrawProtoBytes,
			Signatures:         make([]commontypes.Signature, 0),
			BaseFee:            baseFee.String(),
			PercentageFee:      percFee.String(),
			LastReqTime:        ctx.BlockTime().Unix(),
		}
		k.SetWithdrawInfo(ctx, withdrawId, wdInfo)

		// Mint fees to distribution module
		// NOTE: pegbridge fees are always claimed in the form of original tokens
		k.MintFee(ctx, pair.Orig, new(big.Int).Add(baseFee, percFee))

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeWithdrawToSign,
			sdk.NewAttribute(types.AttributeKeyWithdrawId, withdrawId.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))

		log.Infof("x/pegbr applied: %s. withdrawId: %x withdrawAmt: %s baseFee: %s percFee: %s",
			ev.PrettyLog(onchev.Chainid), withdrawId, withdrawAmt, baseFee, percFee)
		return true, nil
	case types.PegbrEventMint:
		ptbContract, _ := eth.NewPeggedTokenBridgeFilterer(eth.ZeroAddr, nil)
		ev, err := ptbContract.ParseMint(*elog)
		if err != nil {
			return false, err
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
		otvContract, _ := eth.NewOriginalTokenVaultFilterer(eth.ZeroAddr, nil)
		ev, err := otvContract.ParseWithdrawn(*elog)
		if err != nil {
			return false, err
		}
		wdInfo, found := k.GetWithdrawInfo(ctx, ev.WithdrawId)
		if !found {
			log.Errorln("x/pegbr withdraw info not found", ev.PrettyLog(onchev.Chainid))
			return false, nil
		}
		wdInfo.Success = true
		k.SetWithdrawInfo(ctx, ev.WithdrawId, wdInfo)
		log.Infoln("x/pegbr applied:", ev.PrettyLog(onchev.Chainid))
		return true, nil
	}

	return true, nil
}

func (k Keeper) manageDataForRefund(ctx sdk.Context, depositChainId uint64, ev *eth.OriginalTokenVaultDeposited) {
	// Record a DepositInfo without mintId
	depositInfo := types.DepositInfo{
		ChainId:   depositChainId,
		DepositId: ev.DepositId[:],
		MintId:    []byte{},
	}
	k.SetDepositInfo(ctx, ev.DepositId, depositInfo)
	// Record a depositRefund: withdrawOnChain
	wdOnChain := types.WithdrawOnChain{
		Token:       ev.Token.Bytes(),
		Receiver:    ev.Depositor.Bytes(),
		Amount:      ev.Amount.Bytes(),
		BurnAccount: eth.ZeroAddr.Bytes(),
		RefChainId:  depositChainId,
		RefId:       ev.DepositId[:],
	}
	k.SetDepositRefund(ctx, ev.DepositId, wdOnChain)
}
