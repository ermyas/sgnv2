package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

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
		pair, found := k.GetOrigPeggedPair(ctx, depositChainId, ev.Token, ev.MintChainId)
		if !found {
			return false, fmt.Errorf("pegged pair not exists")
		}

		mintAmount, baseFee, percFee := k.CalcAmountAndFees(ctx, pair, ev.Amount, true /* isPeggedDest */)
		if mintAmount.Sign() <= 0 {
			// TODO: Trigger refund, or just ignore?
			return false, fmt.Errorf("amount too small to cover fees, mintAmount %s baseFee %s percFee %s", mintAmount, baseFee, percFee)
		}

		mintTokenAddr := pair.Pegged.Address
		// refChainId is deposit chain ID, refId is deposit ID
		mintId := types.CalcMintId(
			ev.MintAccount, eth.Hex2Addr(mintTokenAddr), mintAmount, ev.Depositor, depositChainId, ev.DepositId)

		if k.HasMintInfo(ctx, mintId) {
			log.Infof("skip already applied pegbr deposit event. chainid %d depositId %x", depositChainId, ev.DepositId)
			return false, nil
		}

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
		}
		k.SetMintInfo(ctx, mintId, mintInfo)

		// Mint fees to distribution module
		// NOTE: pegbridge fees are always claimed in the form of original tokens
		k.MintFee(ctx, pair.Orig, new(big.Int).Add(baseFee, percFee))

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeMintToSign,
			sdk.NewAttribute(types.AttributeKeyMintId, mintId.Hex()),
			sdk.NewAttribute(types.AttributeKeyMintChainId, strconv.FormatUint(ev.MintChainId, 10)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))

		log.Infof("x/pegbr applied: %s. mintId: %x", ev.PrettyLog(onchev.Chainid), mintId)
		return true, nil
	case types.PegbrEventBurn:
		ptbContract, _ := eth.NewPeggedTokenBridgeFilterer(eth.ZeroAddr, nil)
		ev, err := ptbContract.ParseBurn(*elog)
		if err != nil {
			return false, err
		}

		burnChainId := onchev.Chainid
		pair, pairFound := k.GetOrigPeggedPairByPegged(ctx, burnChainId, ev.Token)
		if !pairFound {
			return false, fmt.Errorf("pegged pair not exists")
		}

		withdrawAmt, baseFee, percFee := k.CalcAmountAndFees(ctx, pair, ev.Amount, false /* isPeggedDest */)
		if withdrawAmt.Sign() <= 0 {
			// TODO: Trigger refund, or just ignore?
			return false, fmt.Errorf("amount too small to cover fees, withdrawAmt %s baseFee %s percFee %s", withdrawAmt, baseFee, percFee)
		}

		wdTokenAddr := pair.Orig.Address
		withdrawChainId := pair.Orig.ChainId
		// refChainId is burn chain ID, refId is burn ID
		withdrawId := types.CalcWithdrawId(
			ev.WithdrawAccount, eth.Hex2Addr(wdTokenAddr), withdrawAmt, ev.Account, burnChainId, ev.BurnId)

		if k.HasWithdrawInfo(ctx, withdrawId) {
			log.Infof("skip already applied pegbr burn event. chainId %d burnId %x", burnChainId, ev.BurnId)
			return false, nil
		}

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
		}
		k.SetWithdrawInfo(ctx, withdrawId, wdInfo)

		// Mint fees to distribution module
		// NOTE: pegbridge fees are always claimed in the form of original tokens
		k.MintFee(ctx, pair.Orig, new(big.Int).Add(baseFee, percFee))

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeWithdrawToSign,
			sdk.NewAttribute(types.AttributeKeyWithdrawId, withdrawId.Hex()),
			sdk.NewAttribute(types.AttributeKeyWithdrawChainId, strconv.FormatUint(withdrawChainId, 10)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))

		log.Infof("x/pegbr applied: %s. withdrawId: %x", ev.PrettyLog(onchev.Chainid), withdrawId)
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
