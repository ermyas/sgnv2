package keeper

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/celer-network/sgn-v2/common"
	"github.com/tendermint/tendermint/libs/rand"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
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
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)
	// check reqid, recover user addr, ensure no existing wdDetail-%x-%d
	signAddr, err := ethutils.RecoverSigner(eth.ToPadBytes(req.ReqId), req.UserSig)
	if err != nil {
		return nil, fmt.Errorf("recover signer err: %w", err)
	}
	if GetWithdrawDetail(kv, signAddr, req.ReqId) != nil {
		// same reqid already exist
		return nil, types.WdErr_DUP_REQID
	}
	var wdOnchain *types.WithdrawOnchain
	if req.XferId != nil { // user refund
		xferId := eth.Bytes2Hash(req.XferId)
		wdOnchain = GetXferRefund(kv, xferId)
		if wdOnchain == nil {
			return nil, types.WdErr_XFER_NOT_REFUNDABLE
		}
		if wdOnchain.Seqnum != 0 {
			// already requested withdraw before
			return nil, types.WdErr_XFER_HAS_WITHDRAW
		}
		// now make sure address match
		if eth.Bytes2Addr(wdOnchain.Receiver) != signAddr {
			return nil, types.WdErr_INVALID_SIG
		}
		wdOnchain.Seqnum = req.ReqId
		log.Infof("x/cbr handle refund xferId %x, reqId %d, creator %s, wdOnChain %s",
			xferId, req.ReqId, req.Creator, wdOnchain.String())
	} else { // LP withdraw liquidity
		lpAddr := eth.Bytes2Addr(req.LpAddr)
		if lpAddr != signAddr {
			return nil, types.WdErr_INVALID_SIG
		}
		token := eth.Bytes2Addr(req.Token)
		amt := new(big.Int).SetBytes(req.Amount)
		balance := GetLPBalance(kv, req.Chainid, token, lpAddr)
		log.Infof("x/cbr handle lp withdraw: %s, lp balance %s", req.String()[9:], balance)
		if balance.Cmp(amt) < 0 {
			// balance not enough, return error
			return nil, types.WdErr_LP_BAL_NOT_ENOUGH
		}
		negAmt := new(big.Int).Neg(amt)
		k.Keeper.ChangeLiquidity(sdkCtx, kv, req.Chainid, token, lpAddr, negAmt) // remove amt from lp map
		// also remove liq from liqsum
		ChangeLiqSum(kv, req.Chainid, token, negAmt)
		wdOnchain = &types.WithdrawOnchain{
			Chainid:  req.Chainid,
			Receiver: req.LpAddr,
			Token:    req.Token,
			Amount:   req.Amount,
			Seqnum:   req.ReqId,
		}
	}
	if req.XferId != nil {
		// save this back to avoid dup initwithdraw for refund
		SetXferRefund(kv, eth.Bytes2Hash(req.XferId), wdOnchain)
	}
	wdOnChainRaw, _ := wdOnchain.Marshal()
	SaveWithdrawDetail(kv, signAddr, req.ReqId, &types.WithdrawDetail{
		WdOnchain:   wdOnChainRaw, // only has what to send onchain now
		LastReqTime: sdkCtx.BlockTime().Unix(),
		XferId:      req.XferId, // nil if not user refund
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
	// resp.errmsg is nil if accepted
	usrAddr := eth.Bytes2Addr(req.UserAddr)
	wdDetail := GetWithdrawDetail(kv, usrAddr, req.ReqId)
	if wdDetail == nil {
		// (addr, reqid) not found
		return nil, types.WdErr_NOT_FOUND
	}
	if wdDetail.Completed {
		return nil, types.WdErr_ALREADY_DONE
	}
	nowTime := sdkCtx.BlockTime()
	now := nowTime.Unix()
	if nowTime.Before(common.TsToTime(uint64(wdDetail.LastReqTime)).Add(k.Keeper.GetSignAgainCoolDownDuration(sdkCtx))) {
		return nil, types.WdErr_REQ_TOO_SOON
	}
	// remove all previous sigs
	wdDetail.SortedSigs = nil
	wdDetail.LastReqTime = now
	SaveWithdrawDetail(kv, usrAddr, req.ReqId, wdDetail)
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDataToSign,
		sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_WITHDRAW.String()),
		sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(wdDetail.WdOnchain)),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	return new(types.MsgSignAgainResp), nil
}

// send my sig for data, so it can be later submitted onchain
func (k msgServer) SendMySig(ctx context.Context, msg *types.MsgSendMySig) (*types.MsgSendMySigResp, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil msg")
	}
	logmsg := fmt.Sprintf("x/cbr handle MsgSendMySig type %s", msg.Datatype.String())
	if msg == nil {
		return nil, fmt.Errorf("%s, nil msg", logmsg)
	}
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

		xferRelay := GetXferRelay(kv, xferId, k.cdc)
		if xferRelay == nil {
			return nil, fmt.Errorf("%s xfer not found", logmsg)
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
		if bytes.Compare(latestSigners.GetSignersBytes(), msg.Data) != 0 {
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

func (k msgServer) InternalTransfer(ctx context.Context, req *types.MsgInternalTransfer) (*types.MsgInternalTransferResp, error) {
	ret := &types.MsgInternalTransferResp{}
	xfer := new(types.InternalTransfer)
	err := xfer.Unmarshal(req.GetTransfer())
	if err != nil {
		return nil, err
	}
	sender, err := ethutils.RecoverSigner(req.GetTransfer(), req.GetSig())
	if err != nil {
		return nil, fmt.Errorf("recover signer err: %w", err)
	}
	log.Infof("internal xfer %s, sender %x", xfer.String(), sender)

	amount, ok := new(big.Int).SetString(xfer.GetAmount(), 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount %s", xfer.GetAmount())
	}
	// TODO: pre-processing: check for replay, validated src balance, etc.
	k.Transfer(sdk.UnwrapSDKContext(ctx),
		sender, eth.Hex2Addr(xfer.Token), amount, xfer.SrcChainId, xfer.DstChainId, xfer.MaxSlippage, big.NewInt(int64(rand.Uint32())).Bytes())
	// TODO: post-processing, record status, etc.
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
