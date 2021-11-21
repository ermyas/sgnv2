package keeper

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	wdReq := new(types.WithdrawReq)
	err := wdReq.Unmarshal(req.GetWithdrawReq())
	if err != nil {
		return nil, types.Error(types.ErrCode_INVALID_REQ, "fail to unmarshal")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)
	// check reqid, recover user addr, ensure no existing wdDetail-%x-%d
	signer, err := ethutils.RecoverSigner(req.WithdrawReq, req.UserSig)
	if err != nil {
		return nil, fmt.Errorf("recover signer err: %w", err)
	}
	if GetWithdrawDetail(kv, signer, wdReq.ReqId) != nil {
		// same reqid already exist
		return nil, types.Error(types.ErrCode_DUP_REQID, "withdraw %x %d exists", signer, wdReq.ReqId)
	}
	var wdOnchain *types.WithdrawOnchain
	var xferIdBytes []byte
	switch wdReq.WithdrawType {
	case types.RemoveLiquidity:
		wdOnchain, err = k.withdrawLP(sdkCtx, wdReq, signer, req.Creator)
		if err != nil {
			return nil, err
		}
	case types.RefundTransfer:
		xferIdBytes = common.Hex2Bytes(wdReq.XferId)
		wdOnchain, err = k.refund(sdkCtx, wdReq, signer, req.Creator)
		if err != nil {
			return nil, err
		}
	case types.ClaimFeeShare:
		wdOnchain, err = k.claimFeeShare(sdkCtx, wdReq, signer, req.Creator)
		if err != nil {
			return nil, err
		}
	default:
		return nil, types.Error(types.ErrCode_INVALID_REQ, "invalid withdraw type %d", wdReq.WithdrawType)
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
				return nil, types.Error(types.ErrCode_WD_EXCEED_MAX_OUT_AMOUNT, "withdrawal amount %s exceeds allowance %s", wdAmt, maxSend)
			}
		}
	}

	wdOnChainRaw, _ := wdOnchain.Marshal()
	SaveWithdrawDetail(
		kv, signer, wdReq.ReqId,
		&types.WithdrawDetail{
			WdOnchain:   wdOnChainRaw, // only has what to send onchain now
			LastReqTime: sdkCtx.BlockTime().Unix(),
			XferId:      xferIdBytes, // nil if not user refund
		})
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDataToSign,
		sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_WITHDRAW.String()),
		sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(wdOnChainRaw)),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	return new(types.MsgInitWithdrawResp), nil
}

// user can request to sign a previous withdraw again
// to mitigate dos attack, we could be smart and re-use sigs if
// they are still valid.
func (k msgServer) SignAgain(ctx context.Context, req *types.MsgSignAgain) (*types.MsgSignAgainResp, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)
	switch req.DataType {
	case types.SignDataType_WITHDRAW:
		// resp.errmsg is nil if accepted
		usrAddr := eth.Bytes2Addr(req.UserAddr)
		wdDetail := GetWithdrawDetail(kv, usrAddr, req.ReqId)
		if wdDetail == nil {
			// (addr, reqid) not found
			return nil, types.Error(types.ErrCode_NOT_FOUND, "withdraw %x %d not found", usrAddr, req.ReqId)
		}
		if wdDetail.Completed {
			return nil, types.Error(types.ErrCode_INVALID_STATUS, "withdraw  %x %d  already completed", usrAddr, req.ReqId)
		}
		now := sdkCtx.BlockTime()
		if now.Before(common.TsSecToTime(uint64(wdDetail.LastReqTime)).Add(k.Keeper.GetSignAgainCoolDownDuration(sdkCtx))) {
			return nil, types.Error(types.ErrCode_REQ_TOO_SOON, "")
		}
		// remove all previous sigs
		wdDetail.SortedSigs = nil
		wdDetail.LastReqTime = now.Unix()
		SaveWithdrawDetail(kv, usrAddr, req.ReqId, wdDetail)
		sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDataToSign,
			sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_WITHDRAW.String()),
			sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(wdDetail.WdOnchain)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))
	case types.SignDataType_RELAY:
		xferId := eth.Bytes2Hash(req.XferId)
		xferStatus := GetEvSendStatus(kv, xferId)
		if xferStatus != types.XferStatus_OK_TO_RELAY {
			return nil, types.Error(types.ErrCode_INVALID_STATUS, "invalid transfer %x status %s", xferId, xferStatus)
		}
		relay := GetXferRelay(kv, xferId)
		if relay == nil {
			// this should never happen
			return nil, types.Error(types.ErrCode_NOT_FOUND, "xfer %x not found", xferId)
		}
		now := sdkCtx.BlockTime()
		if now.Before(common.TsSecToTime(uint64(relay.LastReqTime)).Add(k.Keeper.GetSignAgainCoolDownDuration(sdkCtx))) {
			return nil, types.Error(types.ErrCode_REQ_TOO_SOON, "")
		}
		// remove all previous sigs
		relay.SortedSigs = nil
		relay.LastReqTime = now.Unix()
		SetXferRelay(kv, xferId, relay)
		sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDataToSign,
			sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_RELAY.String()),
			sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(relay.Relay)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))
	default:
		return nil, types.Error(types.ErrCode_INVALID_REQ, "invalid sign data type %d", req.DataType)
	}

	return new(types.MsgSignAgainResp), nil
}

// send my sig for data, so it can be later submitted onchain
func (k msgServer) SendMySig(ctx context.Context, msg *types.MsgSendMySig) (*types.MsgSendMySigResp, error) {
	// TODO: use ErrMsg
	if msg == nil {
		return nil, fmt.Errorf("nil msg")
	}
	logmsg := fmt.Sprintf("x/cbr handle MsgSendMySig type %s", msg.Datatype.String())
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)

	// check basics like sig, creator
	senderAcct, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("%s err %w", logmsg, err)
	}
	logmsg = fmt.Sprintf("%s, creator %s", logmsg, senderAcct.String())

	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(sdkCtx, senderAcct)
	if !found {
		return nil, fmt.Errorf("%s, sender is not a validator", logmsg)
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("%s, validator is not bonded", logmsg)
	}

	// validate sig
	signer, err := ethutils.RecoverSigner(msg.Data, msg.MySig)
	if err != nil {
		return nil, fmt.Errorf("%s err %w", logmsg, err)
	}
	if signer != validator.GetSignerAddr() {
		err = fmt.Errorf("mismatch signer address %s %s", signer, validator.GetSignerAddr())
		return nil, fmt.Errorf("%s err %w", logmsg, err)
	}
	logmsg = fmt.Sprintf("%s, signer %x", logmsg, signer)

	ret := &types.MsgSendMySigResp{}
	if msg.Datatype == types.SignDataType_RELAY {
		relay := new(types.RelayOnChain)
		err := relay.Unmarshal(msg.Data)
		if err != nil {
			return nil, fmt.Errorf("%s err %w", logmsg, err)
		}

		// add sig
		xferId := eth.Bytes2Hash(relay.SrcTransferId)
		logmsg = fmt.Sprintf("%s, xferId %x", logmsg, xferId)

		xferRelay := GetXferRelay(kv, xferId)
		if xferRelay == nil {
			return nil, fmt.Errorf("%s xfer not found", logmsg)
		}
		// SortedSigs will be modified in place
		xferRelay.SortedSigs = UpdateSortedSigs(xferRelay.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  msg.MySig,
		})
		SetXferRelay(kv, xferId, xferRelay)
		return ret, nil
	} else if msg.Datatype == types.SignDataType_WITHDRAW {
		onchain := new(types.WithdrawOnchain)
		err := onchain.Unmarshal(msg.Data)
		if err != nil {
			return nil, fmt.Errorf("%s, unmarshal %x to WithdrawOnchain fail %w", logmsg, msg.Data, err)
		}
		logmsg = fmt.Sprintf("%s, seqnum %d", logmsg, onchain.Seqnum)
		usrAddr := eth.Bytes2Addr(onchain.Receiver)
		wdDetail := GetWithdrawDetail(kv, usrAddr, onchain.Seqnum)
		if wdDetail == nil {
			return nil, fmt.Errorf("%s, withdraw seq not found", logmsg)
		}
		wdDetail.SortedSigs = UpdateSortedSigs(wdDetail.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  msg.MySig,
		})
		SaveWithdrawDetail(kv, usrAddr, onchain.Seqnum, wdDetail)
	} else if msg.Datatype == types.SignDataType_SIGNERS {
		latestSigners, found := k.GetLatestSigners(sdkCtx)
		if !found {
			return nil, fmt.Errorf("%s, latest signers not found", logmsg)
		}
		if !bytes.Equal(latestSigners.GetSignersBytes(), msg.Data) {
			return nil, fmt.Errorf("%s, signed latest signers not match stored data", logmsg)
		}
		latestSigners.SortedSigs = UpdateSortedSigs(latestSigners.SortedSigs, &types.AddrSig{
			Addr: signer.Bytes(),
			Sig:  msg.MySig,
		})
		logmsg = fmt.Sprintf("%s, latestSigners %s", logmsg, latestSigners.String())
		k.SetLatestSigners(sdkCtx, &latestSigners)
	}
	log.Info(logmsg)
	return ret, nil
}

func (k msgServer) UpdateLatestSigners(ctx context.Context, msg *types.MsgUpdateLatestSigners) (*types.MsgUpdateLatestSignersResp, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil msg")
	}
	logmsg := "x/cbr handle MsgSendMySig"
	senderAcct, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("%s err %w", logmsg, err)
	}
	logmsg = fmt.Sprintf("%s, creator %s", logmsg, senderAcct.String())

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(sdkCtx, senderAcct)
	if !found {
		return nil, fmt.Errorf("%s, sender is not a validator", logmsg)
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("%s, validator is not bonded", logmsg)
	}

	updated := k.Keeper.UpdateLatestSigners(sdkCtx, false)
	if !updated {
		return nil, fmt.Errorf("%s, signers not (need to be) updated", logmsg)
	}

	return &types.MsgUpdateLatestSignersResp{}, nil
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
