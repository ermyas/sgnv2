package keeper

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k Keeper) refund(ctx sdk.Context, wdReq *types.WithdrawReq, signer eth.Addr, creator string) (*types.WithdrawOnchain, error) {
	kv := ctx.KVStore(k.storeKey)
	xferId := eth.Bytes2Hash(common.Hex2Bytes(wdReq.XferId))
	wdOnchain := GetXferRefund(kv, xferId)
	if wdOnchain == nil {
		return nil, types.Error(types.ErrCode_XFER_NOT_REFUNDABLE, "xfer %d not refundable", xferId)
	}
	if wdOnchain.Seqnum != 0 {
		// already requested withdraw before
		return nil, types.Error(types.ErrCode_XFER_REFUND_STARTED, "xfer %d refund started", xferId)
	}
	// now make sure address match
	if eth.Bytes2Addr(wdOnchain.Receiver) != signer {
		return nil, types.Error(types.ErrCode_INVALID_SIG, "")
	}
	wdOnchain.Seqnum = wdReq.ReqId
	log.Infof("x/cbr handle refund xferId %x, reqId %d, wdOnChain %s, creator %s",
		xferId, wdReq.ReqId, wdOnchain.String(), creator)
	// save this back to avoid dup initwithdraw for refund
	SetXferRefund(kv, xferId, wdOnchain)
	return wdOnchain, nil
}

func (k Keeper) withdrawLP(ctx sdk.Context, wdReq *types.WithdrawReq, lpAddr eth.Addr, creator string) (*types.WithdrawOnchain, error) {
	kv := ctx.KVStore(k.storeKey)
	if len(wdReq.Withdraws) == 0 {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "empty withdraw list")
	}
	var recvToken eth.Addr
	logmsg := fmt.Sprintf("lp:%x request_id:%d exit_chain_id:%d", lpAddr, wdReq.ReqId, wdReq.ExitChainId)
	var wdmsgs string
	reqAmt := big.NewInt(0)
	recvAmt := big.NewInt(0)
	for _, wd := range wdReq.Withdraws {
		wdmsg := fmt.Sprintf("from_chain_id:%d token_addr:%s ratio:%f max_slippage:%f",
			wd.FromChainId, eth.FormatAddrHex(wd.TokenAddr), float32(wd.Ratio)/types.WithdrawPercentageBase, float32(wd.MaxSlippage)/1000000)
		token := eth.Hex2Addr(wd.TokenAddr)
		balance := GetLPBalance(kv, wd.FromChainId, token, lpAddr)
		if balance.Sign() <= 0 {
			return nil, types.Error(types.ErrCode_BAL_NOT_ENOUGH, "%s %s zero balance", logmsg, wdmsg)
		}
		if wd.Ratio > types.WithdrawPercentageBase || wd.Ratio == 0 {
			return nil, types.Error(types.ErrCode_INVALID_REQ, "%s %s invalid ratio", logmsg, wdmsg)
		}
		var destToken eth.Addr
		amt := new(big.Int).Div(new(big.Int).Mul(balance, big.NewInt(int64(wd.Ratio))), big.NewInt(int64(types.WithdrawPercentageBase)))
		wdmsg = fmt.Sprintf("%s req_amt:%s", wdmsg, amt)
		reqAmt.Add(reqAmt, amt)
		if wd.FromChainId == wdReq.ExitChainId {
			recvAmt.Add(recvAmt, amt)
			destToken = token
		} else {
			randBytes := crypto.Keccak256Hash([]byte(fmt.Sprintf("%x-%d-%d", lpAddr, wdReq.ReqId, ctx.BlockTime().Unix())))
			status, recvAmount, destTk, _, _ := k.transfer(
				ctx, token, amt, wd.FromChainId, wdReq.ExitChainId, wd.MaxSlippage, lpAddr, randBytes.Bytes()[0:4])
			wdmsg = fmt.Sprintf("%s recv_amt:%s", wdmsg, recvAmount)
			if status != types.XferStatus_OK_TO_RELAY {
				return nil, types.Error(types.ErrCode_WD_INTERNAL_XFER_FAILURE, "%s %s internal transfer failed %s", logmsg, wdmsg, status)
			}
			recvAmt.Add(recvAmt, recvAmount)
			destToken = destTk
		}
		if recvToken == eth.ZeroAddr {
			recvToken = destToken
		} else if recvToken != destToken {
			return nil, types.Error(types.ErrCode_INVALID_REQ, "%s %s inconsistent exit token %x %x", logmsg, wdmsg, recvToken, destToken)
		}
		wdmsgs += fmt.Sprintf("<%s> ", wdmsg)
	}
	logmsg = fmt.Sprintf("%s %srecv_token:%x, total_req_amt:%s total_recv_amt:%s", logmsg, wdmsgs, recvToken, reqAmt, recvAmt)
	log.Infof("x/cbr handle lp withdraw: %s creator:%s", logmsg, creator)
	negAmt := new(big.Int).Neg(recvAmt)
	k.ChangeLiquidity(ctx, kv, wdReq.ExitChainId, recvToken, lpAddr, negAmt) // remove amt from lp map
	// also remove liq from liqsum
	ChangeLiqSum(kv, wdReq.ExitChainId, recvToken, negAmt)
	return &types.WithdrawOnchain{
		Chainid:  wdReq.ExitChainId,
		Receiver: lpAddr.Bytes(),
		Token:    recvToken.Bytes(),
		Amount:   recvAmt.Bytes(),
		Seqnum:   wdReq.ReqId,
	}, nil
}

func (k Keeper) claimFeeShare(ctx sdk.Context, wdReq *types.WithdrawReq, delAddr eth.Addr, creator string) (*types.WithdrawOnchain, error) {
	kv := ctx.KVStore(k.storeKey)
	if len(wdReq.Withdraws) != 1 {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "only support claiming a single fee")
	}
	if wdReq.ExitChainId != wdReq.Withdraws[0].FromChainId {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "only support claiming fee on the same chain")
	}
	// 1. Claim cBridge fee share in distribution module
	err := k.distrKeeper.ClaimCBridgeFeeShare(ctx, delAddr)
	if err != nil {
		return nil, err
	}
	logmsg := fmt.Sprintf("claimFeeShare:%x request_id:%d exit_chain_id:%d", delAddr, wdReq.ReqId, wdReq.ExitChainId)
	var wdmsgs string
	wd := wdReq.Withdraws[0]
	wdmsg := wd.String()
	feeTokenAddr := eth.Hex2Addr(wd.TokenAddr)
	// 2. Take the fee balance in the distribution module and generate a WithdrawOnchain
	symbol := GetAssetSymbol(kv, &ChainIdTokenAddr{wd.FromChainId, feeTokenAddr})
	denom := fmt.Sprintf("%s%s/%d", types.CBridgeFeeDenomPrefix, symbol, wd.FromChainId)
	coin := k.distrKeeper.GetWithdrawableCBridgeFeeShare(ctx, delAddr, sdk.NewCoin(denom, sdk.ZeroInt()))
	err = k.distrKeeper.BurnCBridgeFeeShare(ctx, delAddr, coin)
	if err != nil {
		return nil, err
	}
	amount := coin.Amount.BigInt()
	wdmsg = fmt.Sprintf("%sreqAmt:%s", wdmsg, amount)
	wdmsgs += fmt.Sprintf("<%s> ", wdmsg)
	logmsg = fmt.Sprintf("%s %sfee_token:%x, amt:%s", logmsg, wdmsgs, feeTokenAddr, amount)
	log.Infof("x/cbr handle claim fee share: %s creator:%s", logmsg, creator)
	return &types.WithdrawOnchain{
		Chainid:  wdReq.ExitChainId,
		Receiver: delAddr.Bytes(),
		Token:    feeTokenAddr.Bytes(),
		Amount:   amount.Bytes(),
		Seqnum:   wdReq.ReqId,
	}, nil
}
