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

type ChainIdTokenDecimal struct {
	*ChainIdTokenAddr
	Decimal uint32
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
		// note we don't check if config has this chid,token so in case someone addLiq *before* sgn supports it, it'll
		// be accounted correctly (but can't be used for transfer as that requires asset info)
		SetEvLiqAdd(kv, onchev.Chainid, ev.Seqnum)
		bal := k.ChangeLiquidity(ctx, kv, onchev.Chainid, ev.Token, ev.Provider, ev.Amount)
		log.Infoln("x/cbr applied:", ev.PrettyLog(onchev.Chainid), "balance:", bal.String())
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
			log.Infoln("x/cbr applied:", ev.PrettyLog(onchev.Chainid), "status:", sendStatus.String())
			SetEvSendStatus(kv, ev.TransferId, sendStatus)
		}()

		sendStatus, destAmount, feeAmt, destTokenAddr :=
			k.Transfer(ctx, eth.ZeroAddr, ev.Token, ev.Amount, onchev.Chainid, ev.DstChainId, ev.MaxSlippage, ev.TransferId[28:]) // last 4B of xfer id

		if sendStatus != types.XferStatus_OK_TO_RELAY {
			if sendStatus == types.XferStatus_BAD_LIQUIDITY || sendStatus == types.XferStatus_BAD_SLIPPAGE {
				SetXferRefund(kv, ev.TransferId, wdOnchain)
			}
			return true, nil
		}
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
		SetXferRelay(kv, ev.TransferId, &types.XferRelay{Relay: relayRaw}, k.cdc)
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDataToSign,
			sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_RELAY.String()),
			sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(relayRaw)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
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
		// relay-%x %x is srcTransferId
		log.Infoln("x/cbr applied:", ev.PrettyLog(onchev.Chainid))
	case types.CbrEventWithdraw:
		ev, err := cbrContract.ParseWithdrawDone(*elog)
		if err != nil {
			return false, err
		}
		log.Infoln("x/cbr apply withdrawDone", ev.String())
		wdDetail := GetWithdrawDetail(kv, ev.Receiver, ev.Seqnum)
		if wdDetail == nil {
			// what to do if not found?
			return true, nil
		}
		wdDetail.Completed = true
		SaveWithdrawDetail(kv, ev.Receiver, ev.Seqnum, wdDetail)
		if wdDetail.XferId != nil {
			// this is a refund so we set xfer status to refund_done
			SetEvSendStatus(kv, eth.Bytes2Hash(wdDetail.XferId), types.XferStatus_REFUND_DONE)
		}
		log.Infoln("x/cbr applied:", ev.PrettyLog(onchev.Chainid))
	case types.CbrEventSignersUpdated:
		ev, err := cbrContract.ParseSignersUpdated(*elog)
		if err != nil {
			return false, err
		}
		chainSigners := &types.ChainSigners{
			ChainId: onchev.Chainid,
		}
		chainSigners.SetByEvent(ev)
		k.SetChainSigners(ctx, chainSigners)
		log.Infoln("x/cbr applied chainSigners:", chainSigners.String())
	}
	return true, nil
}

// check slippage, be careful with decimal diff between src and dest
// user's max slippage represents our promise for at least how many dest token will be relayed
// promised = srcAmt * (1e6 - maxslip) / 1e6 * 1e(dstdecimal - srcdecimal)
// ev.MaxSlippage is slippage * 1e6, eg. 0.1% -> 1000
// if MaxSlippage >= 1e6, likely an attack, return 0 meaning no promise
func calcPromised(maxslip, srcDeci, destDeci uint32, srcAmt *big.Int) *big.Int {
	promised := new(big.Int)
	if maxslip >= 1e6 {
		return promised
	}
	e6 := big.NewInt(1e6)
	promised.Sub(e6, big.NewInt(int64(maxslip)))
	promised.Mul(promised, srcAmt)
	promised.Div(promised, e6)
	if destDeci > srcDeci {
		// dest amt is larger due to more decimals. note we lose some precision due to Div e6 first
		// but it's ok
		upScale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(destDeci-srcDeci)), nil)
		promised.Mul(promised, upScale)
	} else if destDeci < srcDeci {
		downScale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(srcDeci-destDeci)), nil)
		promised.Div(promised, downScale)
	} // if decimal equal, return directly without scaling
	return promised
}
