package keeper

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	SignAgainCoolDownSec = 600 // if last sign within 600s, don't sign again
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// validate withdraw request, update kv then emit to sign event. if req is invalid, return error
func (k msgServer) InitWithdraw(ctx context.Context, req *types.MsgInitWithdraw) (*types.MsgInitWithdrawResp, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)
	// todo: do we need to check creator sig? or it doesn't matter anyway
	var wdOnchain *types.WithdrawOnchain
	resp := &types.MsgInitWithdrawResp{
		ReqId: req.ReqId,
	}
	// we have to emit event to let caller know the response, seqnum or errmsg will set before return
	defer emitWdResp(sdkCtx, resp)
	if req.XferId != nil { // user refund
		xferId := eth.Bytes2Hash(req.XferId)
		wdOnchain = GetXferRefund(kv, xferId)
		if wdOnchain == nil {
			resp.Errmsg = &types.ErrMsg{
				Code: types.ErrCode_XFER_NOT_REFUNDABLE,
			}
			return resp, fmt.Errorf("xfer %x not valid for refund", xferId)
		}
		if wdOnchain.Seqnum != 0 {
			// already requested withdraw before
			resp.Errmsg = &types.ErrMsg{
				Code: types.ErrCode_XFER_HAS_SEQNUM,
				Msg:  fmt.Sprintf("%d", wdOnchain.Seqnum),
			}
			return nil, fmt.Errorf("xfer %x already has withdraw seqnum %d, use SignAgain", xferId, wdOnchain.Seqnum)
		}
	} else { // LP withdraw liquidity
		lpAddr := eth.Bytes2Addr(req.LpAddr)
		token := eth.Bytes2Addr(req.Token)
		amt := new(big.Int).SetBytes(req.Amount)
		balance := GetLPBalance(kv, req.Chainid, token, lpAddr)
		if balance.Cmp(amt) < 0 {
			// balance not enough, return error
			resp.Errmsg = &types.ErrMsg{
				Code: types.ErrCode_LP_BAL_NOT_ENOUGH,
				Msg:  balance.String(),
			}
			return nil, fmt.Errorf("lp balance %s < %s", balance, amt)
		}
		ChangeLiquidity(kv, req.Chainid, token, lpAddr, new(big.Int).Neg(amt)) // remove amt from lp map
		wdOnchain = &types.WithdrawOnchain{
			Chainid:  req.Chainid,
			Receiver: req.LpAddr,
			Token:    req.Token,
			Amount:   req.Amount,
		}
	}
	// can withdraw, errmsg is nil, only set seqnum. seqnum start from 1
	newseq := IncrWithdrawSeq(kv)
	resp.Seqnum = newseq
	wdOnchain.Seqnum = newseq
	if req.XferId != nil {
		// save this back to avoid dup initwithdraw for refund
		SetXferRefund(kv, eth.Bytes2Hash(req.XferId), wdOnchain)
	}
	wdOnChainRaw, _ := wdOnchain.Marshal()
	SaveWithdrawDetail(kv, newseq, &types.WithdrawDetail{
		WdOnchain:   wdOnChainRaw, // only has what to send onchain now
		LastReqTime: sdkCtx.BlockTime().Unix(),
		XferId:      req.XferId, // nil if not user refund
	})
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventToSign,
		sdk.NewAttribute(types.EvAttrType, types.SignDataType_WITHDRAW.String()),
		sdk.NewAttribute(types.EvAttrData, string(wdOnChainRaw))))
	return resp, nil
}

// helper util to emit EventWdResp event
func emitWdResp(sdkCtx sdk.Context, wdresp *types.MsgInitWithdrawResp) {
	raw, _ := wdresp.Marshal()
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventMsgResp,
		sdk.NewAttribute(types.EvAttrMsgType, "MsgInitWithdrawResp"),
		sdk.NewAttribute(types.EvAttrResp, string(raw))))
}

func emitSignAgainResp(sdkCtx sdk.Context, resp *types.MsgSignAgainResp) {
	raw, _ := resp.Marshal()
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventMsgResp,
		sdk.NewAttribute(types.EvAttrMsgType, "MsgSignAgainResp"),
		sdk.NewAttribute(types.EvAttrResp, string(raw))))
}

// user can request to sign a previous withdraw again
// to mitigate dos attack, we could be smart and re-use sigs if
// they are still valid. we should also deny if withdraw already
// completed
func (k msgServer) SignAgain(ctx context.Context, req *types.MsgSignAgain) (*types.MsgSignAgainResp, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)
	// resp.errmsg is nil if accepted
	resp := &types.MsgSignAgainResp{
		ReqId: req.ReqId,
	}
	defer emitSignAgainResp(sdkCtx, resp)

	wdDetail := GetWithdrawDetail(kv, req.Seqnum)
	if wdDetail == nil {
		// not found
		resp.Errmsg = &types.ErrMsg{
			Code: types.ErrCode_SEQ_NOT_FOUND,
		}
		return nil, fmt.Errorf("withdraw seq %d not found", req.Seqnum)
	}
	if wdDetail.Completed {
		resp.Errmsg = &types.ErrMsg{
			Code: types.ErrCode_ALREADY_DONE,
		}
		return nil, fmt.Errorf("withdraw seq %d already completed", req.Seqnum)
	}
	now := sdkCtx.BlockTime().Unix()
	if now-wdDetail.LastReqTime < SignAgainCoolDownSec {
		resp.Errmsg = &types.ErrMsg{
			Code: types.ErrCode_REQ_TOO_SOON,
		}
		return nil, fmt.Errorf("withdraw seq %d sig was last requested at %d, try again after 10min", req.Seqnum, wdDetail.LastReqTime)
	}
	// remove all previous sigs
	wdDetail.SortedSigs = nil
	wdDetail.LastReqTime = now
	SaveWithdrawDetail(kv, req.Seqnum, wdDetail)
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventToSign,
		sdk.NewAttribute(types.EvAttrType, types.SignDataType_WITHDRAW.String()),
		sdk.NewAttribute(types.EvAttrData, string(wdDetail.WdOnchain))))
	return nil, nil
}

// send my sig for data, so it can be later submitted onchain
func (k msgServer) SendMySig(ctx context.Context, msg *types.MsgSendMySig) (*types.MsgSendMySigResp, error) {
	if msg == nil {
		return nil, fmt.Errorf("sendMySig could not be nil")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)

	// check basics like sig, creator
	senderAcct, _ := sdk.AccAddressFromBech32(msg.Creator)
	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(sdkCtx, senderAcct)
	if !found {
		return nil, fmt.Errorf("sender is not a validator")
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("validator is not bonded")
	}

	// validate sig
	signer, err := ethutils.RecoverSigner(msg.Data, msg.MySig)
	if err != nil {
		return nil, err
	}
	if signer != validator.GetSignerAddr() {
		err = fmt.Errorf("mismatch signer address %s %s", signer, validator.GetSignerAddr())
		return nil, err
	}
	ret := &types.MsgSendMySigResp{}
	if msg.Datatype == types.SignDataType_RELAY {
		relay := new(types.RelayOnChain)
		err := relay.Unmarshal(msg.Data)
		if err != nil {
			return nil, err
		}

		// add sig
		xferId := eth.Bytes2Hash(relay.SrcTransferId)
		xferRelay := GetXferRelay(kv, xferId, k.cdc)
		if xferRelay == nil {
			return nil, fmt.Errorf("xfer %x not found", xferId)
		}
		// SortedSigs will be modified in place
		xferRelay.SortedSigs = UpdateSortedSigs(xferRelay.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  msg.MySig,
		})
		SetXferRelay(kv, xferId, xferRelay, k.cdc)
		return ret, nil
	} else if msg.Datatype == types.SignDataType_WITHDRAW {
		onchain := new(types.WithdrawOnchain)
		err := onchain.Unmarshal(msg.Data)
		if err != nil {
			return nil, fmt.Errorf("unmarshal %x to WithdrawOnchain fail %w", msg.Data, err)
		}
		wdDetail := GetWithdrawDetail(kv, onchain.Seqnum)
		if wdDetail == nil {
			return nil, fmt.Errorf("withdraw seq %d not found", onchain.Seqnum)
		}
		wdDetail.SortedSigs = UpdateSortedSigs(wdDetail.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  msg.MySig,
		})
		SaveWithdrawDetail(kv, onchain.Seqnum, wdDetail)
	} else if msg.Datatype == types.SignDataType_SIGNERS {
		latestSigners, found := k.GetLatestSigners(sdkCtx)
		if !found {
			return nil, fmt.Errorf("latest signers not found")
		}
		if bytes.Compare(latestSigners.GetSignersBytes(), msg.Data) != 0 {
			return nil, fmt.Errorf("signed latest signers not match stored data")
		}
		latestSigners.SortedSigs = UpdateSortedSigs(latestSigners.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  msg.MySig,
		})
		k.SetLatestSigners(sdkCtx, &latestSigners)
	}
	return ret, nil
}

// sort curSigs in place and return it. if newsig.Addr equals one already in curSigs, only update sig
func UpdateSortedSigs(curSigs []*types.AddrSig, newsig *types.AddrSig) []*types.AddrSig {
	foundSameAddr := false
	for _, addrSig := range curSigs {
		if bytes.Equal(addrSig.Addr, newsig.Addr) {
			addrSig.Sig = newsig.Sig
			foundSameAddr = true
		}
	}
	if foundSameAddr {
		return curSigs
	}
	// new addr, add then sort by addr
	curSigs = append(curSigs, newsig)
	sort.Slice(curSigs, func(i, j int) bool {
		// note we must compare full 20 bytes, otherwise if address has leading 00, it may be put in the wrong order
		return bytes.Compare(eth.Pad20Bytes(curSigs[i].Addr), eth.Pad20Bytes(curSigs[j].Addr)) == -1
	})
	return curSigs
}
