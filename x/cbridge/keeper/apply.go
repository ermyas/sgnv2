package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ChainIdTokenAddr struct {
	ChId      uint64
	TokenAddr eth.Addr
}

// data is serialized OnChainEvent
func (k Keeper) ApplyEvent(ctx sdk.Context, data []byte) (bool, error) {
	onchev := new(types.OnChainEvent)
	err := onchev.Unmarshal(data)
	if err != nil {
		return false, err
	}
	elog := new(ethtypes.Log)
	err = json.Unmarshal(onchev.Elog, elog)
	if err != nil {
		return false, err
	}
	kv := ctx.KVStore(k.storeKey)
	cbrContract, _ := eth.NewBridgeFilterer(eth.ZeroAddr, nil)
	switch onchev.Evtype {
	case types.CbrEventLiqAdd:
		ev, err := cbrContract.ParseLiquidityAdded(*elog)
		if err != nil {
			return false, err
		}
		if HasEvLiqAdd(kv, onchev.Chainid, ev.Seqnum) {
			// already processed, return error
			return false, fmt.Errorf("already applied liq add event: chainid %d seq %d", onchev.Chainid, ev.Seqnum)
		}
		SetEvLiqAdd(kv, onchev.Chainid, ev.Seqnum)
		bal := k.ChangeLiquidity(ctx, kv, onchev.Chainid, ev.Token, ev.Provider, ev.Amount)

		log.Infoln("x/cbr applied addLiquidity", onchev.Chainid, eth.Addr2Hex(ev.Token), eth.Addr2Hex(ev.Provider), "Balance:", bal.String())
		return true, nil
	case types.CbrEventSend:
		ev, err := cbrContract.ParseSend(*elog)
		if err != nil {
			return false, err
		}
		if HasEvSend(kv, ev.TransferId) {
			return false, fmt.Errorf("already applied send event. chainid %d xferId %x", onchev.Chainid, ev.TransferId)
		}
		log.Infoln("x/cbr apply send", ev.String())
		// in case of bad_xxx, save info for later user refund, NO seqnum yet as it'll be set
		// when user calls InitWithdraw
		wdOnchain := &types.WithdrawOnchain{
			Chainid:  onchev.Chainid,
			Receiver: ev.Sender[:],
			Token:    ev.Token[:],
			Amount:   ev.Amount.Bytes(),
		}

		// must set to non-zero before return
		var sendStatus types.XferStatus
		defer func() {
			SetEvSendStatus(kv, ev.TransferId, sendStatus)
		}()

		src := &ChainIdTokenAddr{
			ChId:      onchev.Chainid,
			TokenAddr: ev.Token,
		}
		assetSym := GetAssetSymbol(kv, src)
		destToken := GetAssetInfo(kv, assetSym, ev.DstChainId)
		destTokenAddr := eth.Hex2Addr(destToken.Addr)
		dest := &ChainIdTokenAddr{
			ChId:      ev.DstChainId,
			TokenAddr: destTokenAddr,
		}
		// now we need to decide if this send can be completed by sgn, eg. has enough liquidity on dest chain etc
		destAmount := k.CalcEqualOnDestChain(src, dest, ev.Amount)
		if destAmount.Sign() == 0 { // avoid div by 0
			// define another enum?
			sendStatus = types.XferStatus_BAD_LIQUIDITY
			SetXferRefund(kv, ev.TransferId, wdOnchain)
			return true, nil
		}
		// check has enough liq on dest chain
		if !HasEnoughLiq(kv, dest, destAmount) {
			sendStatus = types.XferStatus_BAD_LIQUIDITY
			SetXferRefund(kv, ev.TransferId, wdOnchain)
			return true, nil
		}
		feeAmt := CalcFee(kv, src, dest, destAmount)
		// check slippage
		if feeAmt.Sign() == 1 {
			slippage := new(big.Int).Mul(feeAmt, big.NewInt(1e6))
			slippage.Div(slippage, destAmount)
			if slippage.Uint64() > uint64(ev.MaxSlippage) {
				sendStatus = types.XferStatus_BAD_SLIPPAGE
				SetXferRefund(kv, ev.TransferId, wdOnchain)
				return true, nil
			}
		}

		// pick LPs, minus each's destChain liquidity, add src liquidity
		randNum := new(big.Int).SetBytes(ev.TransferId[28:]).Uint64() // last 4B of xfer id
		k.PickLPsAndAdjustLiquidity(ctx, kv, src, dest, ev.Amount, destAmount, feeAmt, randNum)

		relayOnchain := &types.RelayOnChain{
			Sender:        ev.Sender[:],
			Receiver:      ev.Receiver[:],
			Token:         destTokenAddr[:],
			Amount:        new(big.Int).Sub(destAmount, feeAmt).Bytes(),
			SrcChainId:    onchev.Chainid,
			DstChainId:    ev.DstChainId,
			SrcTransferId: ev.TransferId[:],
		}
		relayRaw, _ := relayOnchain.Marshal()
		SetXferRelay(kv, ev.TransferId, &types.XferRelay{
			Relay: relayRaw,
		}, k.cdc)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventToSign,
			sdk.NewAttribute(types.EvAttrType, types.SignDataType_RELAY.String()),
			sdk.NewAttribute(types.EvAttrData, string(relayRaw)),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventToSign),
		))
		sendStatus = types.XferStatus_OK_TO_RELAY
	case types.CbrEventRelay:
		// relay happened on dest chain
		ev, err := cbrContract.ParseRelay(*elog)
		if err != nil {
			return false, err
		}
		log.Infoln("x/cbr apply relay", ev.String())
		SetEvSendStatus(kv, ev.SrcTransferId, types.XferStatus_SUCCESS)
		// only set value when apply event, relay xferid -> src xferid only for debugging
		SetEvRelay(kv, ev.TransferId, ev.SrcTransferId)
	case types.CbrEventWithdraw:
		ev, err := cbrContract.ParseWithdrawDone(*elog)
		if err != nil {
			return false, err
		}
		log.Infoln("x/cbr apply withdrawDone", ev.String())
		wdDetail := GetWithdrawDetail(kv, ev.Seqnum)
		if wdDetail == nil {
			// what to do if not found?
			return true, nil
		}
		wdDetail.Completed = true
		SaveWithdrawDetail(kv, ev.Seqnum, wdDetail)
		if wdDetail.XferId != nil {
			// this is a refund so we set xfer status to refund_done
			SetEvSendStatus(kv, eth.Bytes2Hash(wdDetail.XferId), types.XferStatus_REFUND_DONE)
		}
	case types.CbrEventSignersUpdated:
		ev, err := cbrContract.ParseSignersUpdated(*elog)
		if err != nil {
			return false, err
		}
		chainSigners := &types.ChainSigners{
			ChainId:      onchev.Chainid,
			SignersBytes: ev.CurSigners,
			CurrSigners:  &types.SortedSigners{},
		}
		chainSigners.CurrSigners.Unmarshal(ev.CurSigners)
		k.SetChainSigners(ctx, chainSigners)
		log.Infoln("x/cbr applied chainSigners:", chainSigners.String())
	}
	return true, nil
}
