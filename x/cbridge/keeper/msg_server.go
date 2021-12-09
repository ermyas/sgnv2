package keeper

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

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
		xferIdBytes = eth.Hex2Bytes(wdReq.XferId)
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
		log.Infof("x/cbr sign again Withdraw user %x reqId %d", req.UserAddr, req.ReqId)
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
		log.Infof("x/cbr sign again Relay %x", xferId)
	case types.SignDataType_SIGNERS:
		latestSigners, found := k.GetLatestSigners(sdkCtx)
		if !found {
			return nil, types.Error(types.ErrCode_NOT_FOUND, "latest signers not found")
		}
		now := sdkCtx.BlockTime()
		if now.Before(common.TsSecToTime(latestSigners.LastSignTime).Add(k.Keeper.GetSignAgainCoolDownDuration(sdkCtx))) {
			return nil, types.Error(types.ErrCode_REQ_TOO_SOON, "")
		}
		sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDataToSign,
			sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_SIGNERS.String()),
			sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(latestSigners.SignersBytes)),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))
		latestSigners.LastSignTime = uint64(now.Unix())
		k.SetLatestSigners(sdkCtx, &latestSigners)
		log.Infof("x/cbr sign again UpdateSigners")
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

	ret := &types.MsgSendMySigResp{}
	if msg.Datatype == types.SignDataType_RELAY {
		relay := new(types.RelayOnChain)
		err = relay.Unmarshal(msg.Data)
		if err != nil {
			return nil, fmt.Errorf("%s err %w", logmsg, err)
		}

		// add sig
		xferId := eth.Bytes2Hash(relay.SrcTransferId)
		logmsg = fmt.Sprintf("%s, xferId %x", logmsg, xferId)

		if len(msg.MySigs) != 1 {
			return nil, fmt.Errorf("%s invalid sig num %d", logmsg, len(msg.MySigs))
		}
		mySig := msg.MySigs[0].Sig

		contractAddr, found := GetCbrContract(kv, relay.GetDstChainId())
		if !found {
			return nil, fmt.Errorf("%s contract not found for chain %d", logmsg, relay.GetDstChainId())
		}
		dataToSign := types.EncodeRelayOnChainToSign(relay.GetDstChainId(), contractAddr, msg.Data)
		signer, err := ethutils.RecoverSigner(dataToSign, mySig)
		if err != nil {
			return nil, fmt.Errorf("%s err %w", logmsg, err)
		}
		if signer != validator.GetSignerAddr() {
			return nil, fmt.Errorf("%s invalid signature", logmsg)
		}
		logmsg = fmt.Sprintf("%s, signer %x", logmsg, signer)

		xferRelay := GetXferRelay(kv, xferId)
		if xferRelay == nil {
			return nil, fmt.Errorf("%s xfer not found", logmsg)
		}
		// SortedSigs will be modified in place
		xferRelay.SortedSigs = UpdateSortedSigs(xferRelay.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  mySig,
		})
		SetXferRelay(kv, xferId, xferRelay)
		return ret, nil
	} else if msg.Datatype == types.SignDataType_WITHDRAW {
		withdraw := new(types.WithdrawOnchain)
		err = withdraw.Unmarshal(msg.Data)
		if err != nil {
			return nil, fmt.Errorf("%s, unmarshal %x to WithdrawOnchain fail %w", logmsg, msg.Data, err)
		}
		logmsg = fmt.Sprintf("%s, seqnum %d", logmsg, withdraw.Seqnum)

		if len(msg.MySigs) != 1 {
			return nil, fmt.Errorf("%s invalid sig num %d", logmsg, len(msg.MySigs))
		}
		mySig := msg.MySigs[0].Sig

		contractAddr, found := GetCbrContract(kv, withdraw.Chainid)
		if !found {
			return nil, fmt.Errorf("%s contract not found for chain %d", logmsg, withdraw.Chainid)
		}
		dataToSign := types.EncodeWithdrawOnchainToSign(withdraw.Chainid, contractAddr, msg.Data)
		signer, err := ethutils.RecoverSigner(dataToSign, mySig)
		if err != nil {
			return nil, fmt.Errorf("%s err %w", logmsg, err)
		}
		if signer != validator.GetSignerAddr() {
			return nil, fmt.Errorf("%s invalid signature", logmsg)
		}
		logmsg = fmt.Sprintf("%s, signer %x", logmsg, signer)

		usrAddr := eth.Bytes2Addr(withdraw.Receiver)
		wdDetail := GetWithdrawDetail(kv, usrAddr, withdraw.Seqnum)
		if wdDetail == nil {
			return nil, fmt.Errorf("%s, withdraw seq not found", logmsg)
		}
		wdDetail.SortedSigs = UpdateSortedSigs(wdDetail.SortedSigs, &types.AddrSig{
			Addr: signer[:],
			Sig:  mySig,
		})
		SaveWithdrawDetail(kv, usrAddr, withdraw.Seqnum, wdDetail)
	} else if msg.Datatype == types.SignDataType_SIGNERS {

		latestSigners, found := k.GetLatestSigners(sdkCtx)
		if !found {
			return nil, fmt.Errorf("%s, latest signers not found", logmsg)
		}
		// validate signature
		if !bytes.Equal(latestSigners.GetSignersBytes(), msg.Data) {
			log.Errorf("%s, signed data not match stored data", logmsg)
			return ret, nil
		}
		for _, mySig := range msg.MySigs {
			// validate sig
			contractAddr, found := GetCbrContract(kv, mySig.GetChainId())
			if !found {
				log.Warnf("%s contract not found for chain %d", logmsg, mySig.GetChainId())
				continue
			}
			dataToSign := types.EncodeSignersUpdateToSign(mySig.GetChainId(), contractAddr, msg.Data)
			signer, err := ethutils.RecoverSigner(dataToSign, mySig.Sig)
			if err != nil {
				return nil, fmt.Errorf("%s err %w", logmsg, err)
			}
			if signer != validator.GetSignerAddr() {
				return nil, fmt.Errorf("%s invalid signature", logmsg)
			}
			logmsg = fmt.Sprintf("%s, signer %x", logmsg, signer)

			chainSigners, found := k.GetChainSigners(sdkCtx, mySig.ChainId)
			if !found {
				log.Errorf("%s, signerSigs for chainId %d not found", logmsg, mySig.ChainId)
				continue
			}
			chainSigners.SortedSigs = UpdateSortedSigs(chainSigners.SortedSigs, &types.AddrSig{
				Addr: signer.Bytes(),
				Sig:  mySig.Sig,
			})
			k.SetChainSigners(sdkCtx, &chainSigners)
		}
		logmsg = fmt.Sprintf("%s, latestSigners %s", logmsg, latestSigners.String())
	}
	log.Info(logmsg)
	return ret, nil
}

func (k msgServer) UpdateLatestSigners(ctx context.Context, msg *types.MsgUpdateLatestSigners) (*types.MsgUpdateLatestSignersResp, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil msg")
	}
	logmsg := "x/cbr handle UpdateLatestSigners"
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

// Add newSig and return updated curSigs. if newSig.Addr equals one already in curSigs, only update sig
func UpdateSortedSigs(curSigs []*types.AddrSig, newSig *types.AddrSig) []*types.AddrSig {
	for i, addrSig := range curSigs {
		if bytes.Equal(addrSig.Addr, newSig.Addr) {
			// Overwriting existing sig
			addrSig.Sig = newSig.Sig
			return curSigs
		}
		// NOTE: We must compare full 20 bytes, otherwise if address has leading 00, it may be put in the wrong order
		if bytes.Compare(eth.Pad20Bytes(newSig.Addr), eth.Pad20Bytes(addrSig.Addr)) == -1 {
			// Found the spot, do insertion
			newSigs := append(curSigs[:i+1], curSigs[i:]...)
			newSigs[i] = newSig
			return newSigs
		}
	}
	// Address larger than all existing signers, append to the end
	return append(curSigs, newSig)
}

// SyncFarming attempts to sync the liquidity of a (chainID, token) for an LP
// with their stake in the farming module.
func (k msgServer) SyncFarming(goCtx context.Context, msg *types.MsgSyncFarming) (*types.MsgSyncFarmingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenAddr := eth.Hex2Addr(msg.TokenAddress)
	lpAddr := eth.Hex2Addr(msg.LpAddress)
	liqMapKey := types.LiqMapKey(msg.ChainId, tokenAddr, lpAddr)
	kv := ctx.KVStore(k.storeKey)
	liqBytes := kv.Get(liqMapKey)
	liquidity := new(big.Int).SetBytes(liqBytes)

	logMsg := "x/cbr handle SyncFarming"
	symbol := GetAssetSymbol(kv, &ChainIdTokenAddr{msg.ChainId, tokenAddr})
	err := k.Keeper.SyncFarming(ctx, symbol, msg.ChainId, lpAddr, liquidity)
	if err != nil {
		return nil, fmt.Errorf("%s err %w", logMsg, err)
	}
	log.Info(logMsg)
	return &types.MsgSyncFarmingResponse{}, nil
}
