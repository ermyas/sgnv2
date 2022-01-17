package keeper

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k Keeper) refund(ctx sdk.Context, wdReq *types.WithdrawReq, signer eth.Addr, creator string) (*types.WithdrawOnchain, error) {
	// we use non-zero reqid as a way to avoid duplicated refund request, so reqid MUST NOT be 0
	if wdReq.ReqId == 0 {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "reqid is 0")
	}
	kv := ctx.KVStore(k.storeKey)
	xferId := eth.Bytes2Hash(eth.Hex2Bytes(wdReq.XferId))
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
	wdOnchain.Refid = eth.Hex2Bytes(wdReq.XferId)
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
	reqAmt := big.NewInt(0)  // total request amount to be withdrawn from all chains
	recvAmt := big.NewInt(0) // total actually recvd amount at exit chain, may be different from reqAmt due to slippage and fees
	// process each withdrawLq request
	for _, wd := range wdReq.Withdraws {
		token := eth.Hex2Addr(wd.TokenAddr)
		wdmsg := fmt.Sprintf("from_chain_id:%d token_addr:%x ratio:%f max_slippage:%f",
			wd.FromChainId, token, float32(wd.Ratio)/types.WithdrawPercentageBase, float32(wd.MaxSlippage)/1e6)
		// get the LP's balance at the "from" chain
		balance := GetLPBalance(kv, wd.FromChainId, token, lpAddr)
		if balance.Sign() <= 0 {
			return nil, types.Error(types.ErrCode_BAL_NOT_ENOUGH, "%s %s zero balance", logmsg, wdmsg)
		}
		if wd.Ratio > types.WithdrawPercentageBase || wd.Ratio == 0 {
			return nil, types.Error(types.ErrCode_INVALID_REQ, "%s %s invalid ratio", logmsg, wdmsg)
		}
		var destToken eth.Addr
		// compute the amount to be withdrawn from this chain
		amt := new(big.Int).Div(new(big.Int).Mul(balance, big.NewInt(int64(wd.Ratio))), big.NewInt(int64(types.WithdrawPercentageBase)))
		wdmsg = fmt.Sprintf("%s req_amt:%s", wdmsg, amt)
		reqAmt.Add(reqAmt, amt)
		if wd.FromChainId == wdReq.ExitChainId {
			// if this is also the exit chain, directly withdraw
			recvAmt.Add(recvAmt, amt)
			destToken = token
		} else {
			// if this is not the exit chain, simulate the following behavior:
			// 1. withdraw exact amt from this chain (no onchain submission)
			// 2. transfer the withdrawn amt to the exit chain (similar to send/relay flow)
			// 3. add the transfer recv amt (after slippage and fee) at exit chain to the total recvAmt
			randBytes := crypto.Keccak256Hash([]byte(fmt.Sprintf("%x-%d-%d", lpAddr, wdReq.ReqId, ctx.BlockTime().Unix())))
			status, recvAmount, destTk, _, _, err := k.transfer(
				ctx, token, amt, wd.FromChainId, wdReq.ExitChainId, wd.MaxSlippage, lpAddr, randBytes.Bytes()[0:4])
			wdmsg = fmt.Sprintf("%s recv_amt:%s", wdmsg, recvAmount)
			if err != nil {
				wdmsg = fmt.Sprintf("%s err:%s", wdmsg, err)
			}
			if status != types.XferStatus_OK_TO_RELAY {
				return nil, types.Error(types.ErrCode_WD_INTERNAL_XFER_FAILURE, "%s %s internal transfer failed %s", logmsg, wdmsg, status)
			}
			// add to total receive amount
			recvAmt.Add(recvAmt, recvAmount)
			destToken = destTk
		}
		negAmt := new(big.Int).Neg(amt)
		// remove amt from lp map at this withdraw_from chain
		k.ChangeLiquidity(ctx, kv, wd.FromChainId, token, lpAddr, negAmt)
		// also remove liq from liqsum at this withdraw_from chain
		ChangeLiqSum(kv, wd.FromChainId, token, negAmt)

		if recvToken == eth.ZeroAddr {
			recvToken = destToken
		} else if recvToken != destToken {
			return nil, types.Error(types.ErrCode_INVALID_REQ, "%s %s inconsistent exit token %x %x", logmsg, wdmsg, recvToken, destToken)
		}
		wdmsgs += fmt.Sprintf("<%s> ", wdmsg)
	}
	logmsg = fmt.Sprintf("%s %srecv_token:%x, total_req_amt:%s total_recv_amt:%s", logmsg, wdmsgs, recvToken, reqAmt, recvAmt)
	log.Infof("x/cbr handle lp withdraw: %s creator:%s", logmsg, creator)
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
	if len(wdReq.Withdraws) < 1 {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "no Withdraw in WithdrawReq")
	}
	// 1. Claim cBridge fee share in distribution module
	err := k.distrKeeper.ClaimCBridgeFeeShare(ctx, delAddr)
	if err != nil {
		return nil, err
	}
	logmsg := fmt.Sprintf("EthAddr %x ReqId %d ExitChainId %d", delAddr, wdReq.ReqId, wdReq.ExitChainId)
	var wdmsgs string
	var destToken common.Address
	totalRecvAmt := big.NewInt(0)
	// 2. Take the fee balance in the distribution module for each chain and generate a WithdrawOnchain
	for _, wd := range wdReq.Withdraws {
		token := eth.Hex2Addr(wd.TokenAddr)
		wdmsg := wd.String()
		symbol := GetAssetSymbol(kv, &ChainIdTokenAddr{wd.FromChainId, token})
		denom := fmt.Sprintf("%s%s/%d", types.CBridgeFeeDenomPrefix, symbol, wd.FromChainId)
		coin := k.distrKeeper.GetWithdrawableBalance(ctx, delAddr, sdk.NewCoin(denom, sdk.ZeroInt()))
		err = k.BurnFeeShare(ctx, delAddr, coin)
		if err != nil {
			return nil, err
		}
		amt := coin.Amount.BigInt()
		if wd.FromChainId == wdReq.ExitChainId {
			totalRecvAmt = totalRecvAmt.Add(totalRecvAmt, amt)
			destToken = token
		} else {
			randBytes := crypto.Keccak256Hash([]byte(fmt.Sprintf("%x-%d-%d", delAddr, wdReq.ReqId, ctx.BlockTime().Unix())))
			status, recvAmt, destTk, _, _, err := k.transfer(
				ctx, token, amt, wd.FromChainId, wdReq.ExitChainId, wd.MaxSlippage, delAddr, randBytes.Bytes()[0:4])
			wdmsg = fmt.Sprintf("%sRecvAmt %s ", wdmsg, recvAmt)
			if err != nil {
				wdmsg = fmt.Sprintf("%s err %s", wdmsg, err)
			}
			if status != types.XferStatus_OK_TO_RELAY {
				return nil, types.Error(types.ErrCode_WD_INTERNAL_XFER_FAILURE, "%s %s internal transfer failed %s", logmsg, wdmsg, status)
			}
			totalRecvAmt = totalRecvAmt.Add(totalRecvAmt, recvAmt)
			destToken = destTk
		}
		wdmsg = fmt.Sprintf("%sreqAmt %s", wdmsg, amt)
		wdmsgs += fmt.Sprintf("<%s> ", wdmsg)
	}

	logmsg = fmt.Sprintf("%s %s FeeToken %x, TotalRecvAmt %s", logmsg, wdmsgs, destToken, totalRecvAmt)
	log.Infof("x/cbr handle claim fee share: %s creator:%s", logmsg, creator)
	// Use 0x1 to represent fee share claims. Must be of length 32.
	refId := eth.Hash{}
	refId[31] = 1
	return &types.WithdrawOnchain{
		Chainid:  wdReq.ExitChainId,
		Receiver: delAddr.Bytes(),
		Token:    destToken.Bytes(),
		Amount:   totalRecvAmt.Bytes(),
		Seqnum:   wdReq.ReqId,
		Refid:    refId[:],
	}, nil
}
