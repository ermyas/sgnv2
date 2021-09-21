package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

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
		newliq := ChangeLiquidity(kv, onchev.Chainid, ev.Token, ev.Provider, ev.Amount)

		k.Logger(ctx).Info("Applied LP add_liquidity", "LQKey", types.LiqMapKey(onchev.Chainid, ev.Token, ev.Provider), "NewAmt", newliq.String())
		return true, nil
	case types.CbrEventSend:
		ev, err := cbrContract.ParseSend(*elog)
		if err != nil {
			return false, err
		}
		if HasEvSend(kv, ev.TransferId) {
			return false, fmt.Errorf("already applied send event. chainid %d xferId %x", onchev.Chainid, ev.TransferId)
		}
		// TODO: SetEvSend
		src := &ChainIdTokenAddr{
			ChId:      onchev.Chainid,
			TokenAddr: ev.Token,
		}
		assetSym := k.GetAssetSymbol(src)
		destTokenAddr := k.GetTokenAddr(assetSym, ev.DstChainId)
		dest := &ChainIdTokenAddr{
			ChId:      ev.DstChainId,
			TokenAddr: destTokenAddr,
		}
		// now we need to decide if this send can be completed by sgn, eg. has enough liquidity on dest chain etc
		destAmount := k.CalcEqualOnDestChain(src, dest, ev.Amount)
		if destAmount.Sign() == 0 {
			// what do we do?
			// must not continue to avoid div by zero in slippage calc
		}
		// check has enough liq on dest chain
		if !HasEnoughLiq(kv, dest, destAmount) {
			// bad transfer, what to do?
		}
		userGet := k.CalcUserGet(src, dest, destAmount)
		// check slippage
		lessAmount := new(big.Int).Sub(ev.Amount, destAmount)
		if lessAmount.Sign() == 1 {
			slippage := new(big.Int).Mul(lessAmount, big.NewInt(1e6))
			slippage.Div(slippage, destAmount)
			if slippage.Uint64() > uint64(ev.MaxSlippage) {
				// slippage too big, bad transfer
			}
		}
		// pick LPs, minus each's destChain liquidity, return how much to add on src chain into kv, but not
		// add src liquidity yet, must wait till Relay event
		k.PickLPsAndAdjustLiquidity(src, dest, ev.Amount, destAmount, userGet)

		relayOnchain := &types.RelayOnChain{
			// todo: fill fields
		}
		relayRaw, _ := relayOnchain.Marshal()
		// Save transfer detail, including relayRaw
		// Add to torelay xfer id list
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventToSign,
			sdk.NewAttribute(types.EvAttrType, types.SignDataType_RELAY.String()),
			sdk.NewAttribute(types.EvAttrData, string(relayRaw)),
		))
	case types.CbrEventWithdraw:
		ev, err := cbrContract.ParseWithdrawDone(*elog)
		if err != nil {
			return false, err
		}
		wdDetail := GetWithdrawDetail(kv, ev.Seqnum)
		if wdDetail == nil {
			// what to do if not found?
			return true, nil
		}
		wdDetail.Completed = true
		SaveWithdrawDetail(kv, ev.Seqnum, wdDetail)
	}
	return true, nil
}
