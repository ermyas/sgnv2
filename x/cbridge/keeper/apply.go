package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
			// already processed, could happen if two syncers both propose
			log.Infof("skip already applied liq add event: chainid %d seq %d", onchev.Chainid, ev.Seqnum)
			return false, nil
		}
		// record the lp's chain when he first adds liquidity
		SetLPOrigin(kv, ev.Provider, onchev.Chainid)
		// note we don't check if config has this chid,token so in case someone addLiq *before* sgn supports it, it'll
		// be accounted correctly (but can't be used for transfer as that requires asset info)
		SetEvLiqAdd(kv, onchev.Chainid, ev.Seqnum)
		bal := k.ChangeLiquidity(ctx, kv, onchev.Chainid, ev.Token, ev.Provider, ev.Amount)
		ChangeLiqSum(kv, onchev.Chainid, ev.Token, ev.Amount)
		log.Infoln("x/cbr applied:", ev.PrettyLog(onchev.Chainid), "balance:", bal.String())
		return true, nil
	case types.CbrEventSend:
		ev, err := cbrContract.ParseSend(*elog)
		if err != nil {
			return false, err
		}
		if HasEvSend(kv, ev.TransferId) {
			log.Infof("skip already applied send event. chainid %d xferId %x", onchev.Chainid, ev.TransferId)
			return false, nil
		}
		// in case of bad_xxx, save info for later user refund, NO seqnum yet as it'll be set
		// when user calls InitWithdraw
		wdOnchain := &types.WithdrawOnchain{
			Chainid:  onchev.Chainid,
			Receiver: ev.Sender[:],
			Token:    ev.Token[:],
			Amount:   ev.Amount.Bytes(),
		}
		randBytes := crypto.Keccak256Hash([]byte(fmt.Sprintf("%x-%d", ev.TransferId, ctx.BlockTime().Unix())))
		// must set to non-zero before return
		sendStatus, userReceive, destTokenAddr, percFee, baseFee, err :=
			k.transfer(ctx, ev.Token, ev.Amount, onchev.Chainid, ev.DstChainId, ev.MaxSlippage, eth.ZeroAddr, randBytes.Bytes()[0:4])

		defer func() {
			logmsg := fmt.Sprintf("x/cbr applied: %s, status: %s, recv %s, fee perc %s base %s",
				ev.PrettyLog(onchev.Chainid), sendStatus, userReceive, percFee, baseFee)
			if err != nil {
				logmsg = fmt.Sprintf("%s, err: %s", logmsg, err)
			}
			log.Info(logmsg)
			SetEvSendStatus(kv, ev.TransferId, sendStatus)
		}()

		// Check and set refund
		switch sendStatus {
		case types.XferStatus_OK_TO_RELAY:
			// no-op
		case types.XferStatus_BAD_LIQUIDITY,
			types.XferStatus_BAD_SLIPPAGE,
			types.XferStatus_BAD_XFER_DISABLED,
			types.XferStatus_BAD_DEST_CHAIN,
			types.XferStatus_EXCEED_MAX_OUT_AMOUNT:
			SetXferRefund(kv, ev.TransferId, wdOnchain)
			return true, nil
		default:
			// Just return for non-refundable failure cases
			return true, nil
		}

		relayOnchain := &types.RelayOnChain{
			Sender:        ev.Sender[:],
			Receiver:      ev.Receiver[:],
			Token:         destTokenAddr[:],
			Amount:        userReceive.Bytes(),
			SrcChainId:    onchev.Chainid,
			DstChainId:    ev.DstChainId,
			SrcTransferId: ev.TransferId[:],
		}
		relayRaw, _ := relayOnchain.Marshal()
		SetXferRelay(kv, ev.TransferId,
			&types.XferRelay{
				Relay:       relayRaw,
				LastReqTime: ctx.BlockTime().Unix(),
				PercFee:     percFee.Bytes(),
				BaseFee:     baseFee.Bytes(),
			})
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDataToSign,
			sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_RELAY.String()),
			sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(relayRaw)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))
	case types.CbrEventRelay:
		// relay happened on dest chain
		ev, err := cbrContract.ParseRelay(*elog)
		if err != nil {
			return false, err
		}
		SetEvSendStatus(kv, ev.SrcTransferId, types.XferStatus_SUCCESS)
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

		latestSigners, _ := k.GetLatestSigners(ctx)
		if types.EqualSortedSigners(latestSigners.GetSortedSigners(), chainSigners.GetSortedSigners()) {
			// updated to latest
			chainSigners.SortedSigs = []*types.AddrSig{}
		}

		k.SetChainSigners(ctx, chainSigners)
		log.Infoln("x/cbr applied chainSigners:", chainSigners.String())
	case types.CbrEventWithdrawalRequest:
		wdiContract, _ := eth.NewWithdrawInboxFilterer(eth.ZeroAddr, nil)
		ev, err := wdiContract.ParseWithdrawalRequest(*elog)
		if err != nil {
			return false, err
		}
		log.Infoln("x/cbr apply withdrawalRequest", ev.String())
		origin := GetLPOrigin(kv, ev.Sender)
		if onchev.Chainid != origin {
			//WithdrawInbox contract should be called on the chain where lp first added their liquidity.
			log.Errorf("%d(chainid of this event) mismatches %d(chainid recorded on sgn when sender first add liq)", onchev.Chainid, origin)
			return false, nil
		}
		deadline := common.TsSecToTime(ev.Deadline.Uint64())
		if ctx.BlockTime().After(deadline) {
			log.Errorf("This withdrawal request has passed the deadline %s.", deadline.Format("2006.01.02 15:04:05"))
			return false, nil
		}
		//construct a withdraw request for initiating withdraw
		var wds []*types.WithdrawLq
		for i, chain := range ev.FromChains {
			wd := &types.WithdrawLq{
				FromChainId: chain,
				TokenAddr:   eth.Addr2Hex(ev.Tokens[i]),
				Ratio:       ev.Ratios[i],
				MaxSlippage: ev.Slippages[i],
			}
			wds = append(wds, wd)
		}
		wdReq := &types.WithdrawReq{
			Withdraws:    wds,
			ExitChainId:  ev.ToChain,
			ReqId:        ev.SeqNum,
			WithdrawType: types.ContractRemoveLiquidity,
		}
		err = k.initWithdraw(ctx, wdReq, nil, "", ev.Receiver, ev.Sender)
		if err != nil {
			return false, err
		}
		log.Infof("x/cbr applied withdrawalRequest: from %x to %x on chain %d", ev.Sender, ev.Receiver, ev.ToChain)
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

func (k Keeper) ApplyUpdateCbrPrice(ctx sdk.Context, data []byte) (bool, error) {
	price := new(types.CbrPrice)
	err := price.Unmarshal(data)
	if err != nil {
		return false, err
	}
	k.SetCbrPrice(ctx, price)
	log.Infoln("x/cbr applied UpdateCbrPrice:", price)
	return true, nil
}
