package keeper

import (
	"fmt"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// inboxReceiver and inboxSender are only supplied for onchain withdrawal request.
func (k Keeper) initWithdraw(ctx sdk.Context, wdReq *types.WithdrawReq, userSig []byte, creator string, inboxReceiver, inboxSender eth.Addr) error {
	kv := ctx.KVStore(k.storeKey)
	var signer eth.Addr
	var err error
	// if sig is empty AND wd is a refund, we assume it's for a contract sender, so signer equals wdOnchain.Receiver
	// otherwise for refund, recovered signer must match wdOnchain.Receiver
	// wdOnchain.Receiver is saved when apply Send event, value is xfer sender
	if len(userSig) == 0 && wdReq.WithdrawType == types.RefundTransfer {
		// usersig is not set, assume contract refund, unfortunately we have to duplicate some
		// logic from k.refund for security reason
		xferId := eth.Bytes2Hash(eth.Hex2Bytes(wdReq.XferId))
		wdOnchain := GetXferRefund(kv, xferId)
		if wdOnchain == nil {
			return types.Error(types.ErrCode_XFER_NOT_REFUNDABLE, "xfer %x not refundable", xferId)
		}
		signer = eth.Bytes2Addr(wdOnchain.Receiver)
	} else if wdReq.WithdrawType == types.ValidatorClaimFeeShare {
		senderSgnAcct, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return types.Error(types.ErrCode_INVALID_REQ, "invalid creator accnt")
		}
		validator, found := k.stakingKeeper.GetValidatorBySgnAddr(ctx, senderSgnAcct)
		if !found {
			return types.Error(types.ErrCode_NOT_FOUND, "creator accnt %s not validator", senderSgnAcct)
		}
		signer = validator.GetEthAddr()
	} else if wdReq.WithdrawType == types.ContractRemoveLiquidity {
		// signer is used as the receiver's address
		signer = inboxReceiver
	} else {
		// check reqid, recover user addr, ensure no existing wdDetail-%x-%d
		signer, err = ethutils.RecoverSigner(k.cdc.MustMarshal(wdReq), userSig)
		if err != nil {
			return fmt.Errorf("recover signer err: %w", err)
		}
	}
	// note wdReq.ReqId could be 0, but as long as signer matches expected, we're ok.
	// note if someone sends in random sig data, it'll recover a random address so this
	// check will not stop duplicated withdraw, therefore further logic MUST check signer
	// is expected!!!
	if GetWithdrawDetail(kv, signer, wdReq.ReqId) != nil {
		// same reqid already exist
		return types.Error(types.ErrCode_DUP_REQID, "withdraw %x %d exists", signer, wdReq.ReqId)
	}
	var wdOnchain *types.WithdrawOnchain
	var xferIdBytes []byte
	switch wdReq.WithdrawType {
	case types.RemoveLiquidity:
		wdOnchain, err = k.withdrawLP(ctx, wdReq, signer, creator)
		if err != nil {
			return err
		}
	case types.RefundTransfer:
		xferIdBytes = eth.Hex2Bytes(wdReq.XferId)
		wdOnchain, err = k.refund(ctx, wdReq, signer, creator)
		if err != nil {
			return err
		}
	case types.ClaimFeeShare, types.ValidatorClaimFeeShare:
		wdOnchain, err = k.claimFeeShare(ctx, wdReq, signer, creator)
		if err != nil {
			return err
		}
	case types.ContractRemoveLiquidity:
		wdOnchain, err = k.withdrawLPFrom(ctx, wdReq, inboxReceiver, inboxSender)
		if err != nil {
			return err
		}
	default:
		return types.Error(types.ErrCode_INVALID_REQ, "invalid withdraw type %s", wdReq.WithdrawType)
	}

	// rate limit check
	assetInfo := GetAssetInfo(kv, GetAssetSymbol(kv, &ChainIdTokenAddr{
		ChId:      wdOnchain.Chainid,
		TokenAddr: eth.Bytes2Addr(wdOnchain.Token),
	}), wdOnchain.Chainid)
	if assetInfo.GetMaxOutAmt() != "" {
		maxSend, ok := new(big.Int).SetString(assetInfo.GetMaxOutAmt(), 10)
		if ok && isPos(maxSend) {
			wdAmt := new(big.Int).SetBytes(wdOnchain.Amount)
			if wdAmt.Cmp(maxSend) == 1 {
				return types.Error(types.ErrCode_WD_EXCEED_MAX_OUT_AMOUNT, "withdrawal amount %s exceeds allowance %s", wdAmt, maxSend)
			}
		}
	}

	wdOnChainRaw, _ := wdOnchain.Marshal()
	SaveWithdrawDetail(
		kv, signer, wdReq.ReqId,
		&types.WithdrawDetail{
			WdOnchain:   wdOnChainRaw, // only has what to send onchain now
			LastReqTime: ctx.BlockTime().Unix(),
			XferId:      xferIdBytes, // nil if not user refund
		})
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDataToSign,
		sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_WITHDRAW.String()),
		sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(wdOnChainRaw)),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	return nil
}

func (k Keeper) refund(ctx sdk.Context, wdReq *types.WithdrawReq, signer eth.Addr, creator string) (*types.WithdrawOnchain, error) {
	// we use non-zero reqid as a way to avoid duplicated refund request, so reqid MUST NOT be 0
	if wdReq.ReqId == 0 {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "reqid is 0")
	}
	kv := ctx.KVStore(k.storeKey)
	xferId := eth.Bytes2Hash(eth.Hex2Bytes(wdReq.XferId))
	wdOnchain := GetXferRefund(kv, xferId)
	if wdOnchain == nil {
		return nil, types.Error(types.ErrCode_XFER_NOT_REFUNDABLE, "xfer %x not refundable", xferId)
	}
	if wdOnchain.Seqnum != 0 {
		// already requested withdraw before
		return nil, types.Error(types.ErrCode_XFER_REFUND_STARTED, "xfer %x refund started", xferId)
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

func (k Keeper) withdrawLPFrom(ctx sdk.Context, wdReq *types.WithdrawReq, receiver, sender eth.Addr) (*types.WithdrawOnchain, error) {
	log.Infof("x/cbr handle contract lp withdrawal request, from(creator):%x, to:%x", sender, receiver)
	wdOnchain, err := k.withdrawLP(ctx, wdReq, sender, fmt.Sprintf("ETH:%x", sender))
	if err != nil {
		return wdOnchain, err
	}
	wdOnchain.Receiver = receiver.Bytes()
	return wdOnchain, nil
}

// creator would be an eth address if a contractLP withdrawal request is processing
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
		coin := k.distrKeeper.GetWithdrawableBalance(ctx, delAddr, denom)

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

func (k Keeper) Refund(ctx sdk.Context, xferId eth.Hash, nonce uint64) error {
	wdReq := &types.WithdrawReq{
		XferId:       xferId.Hex(),
		ReqId:        nonce,
		WithdrawType: types.RefundTransfer,
	}
	return k.initWithdraw(ctx, wdReq, nil, "", eth.ZeroAddr, eth.ZeroAddr)
}
