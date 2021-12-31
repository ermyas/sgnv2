package keeper

import (
	"context"
	"errors"
	"fmt"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	log "github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
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

func (k msgServer) SignMint(goCtx context.Context, msg *types.MsgSignMint) (*types.MsgSignMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator, err := k.isSenderBondedValidator(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}
	mindId := eth.Hex2Hash(msg.MintId)
	mintInfo, found := k.GetMintInfo(ctx, mindId)
	if !found {
		return nil, types.WrapErrNoInfoFound(mindId)
	}
	bridge, found := k.GetPeggedTokenBridge(ctx, mintInfo.ChainId)
	if !found {
		return nil, types.WrapErrNoPeggedTokenBridgeFound(mintInfo.ChainId)
	}
	msgToSign := mintInfo.EncodeDataToSign(eth.Hex2Addr(bridge.Address))
	addSigErr := mintInfo.AddSig(
		msgToSign,
		msg.Signature,
		validator.GetSignerAddr().String(),
	)
	if addSigErr != nil {
		return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
	}
	k.SetMintInfo(ctx, mindId, mintInfo)
	log.Infof("x/pegbridge SignMintInfo add sig mintId:%x signer:%x :sender:%s", mindId, validator.GetSignerAddr(), msg.Sender)
	return &types.MsgSignMintResponse{}, nil
}

func (k msgServer) SignWithdraw(goCtx context.Context, msg *types.MsgSignWithdraw) (*types.MsgSignWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator, err := k.isSenderBondedValidator(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}
	withdrawId := eth.Hex2Hash(msg.WithdrawId)
	withdrawInfo, found := k.GetWithdrawInfo(ctx, withdrawId)
	if !found {
		return nil, types.WrapErrNoInfoFound(withdrawId)
	}
	vaults, found := k.GetOriginalTokenVault(ctx, withdrawInfo.ChainId)
	if !found {
		return nil, types.WrapErrNoOriginalTokenVaultFound(withdrawInfo.ChainId)
	}
	msgToSign := withdrawInfo.EncodeDataToSign(eth.Hex2Addr(vaults.Address))
	addSigErr := withdrawInfo.AddSig(
		msgToSign,
		msg.Signature,
		validator.GetSignerAddr().String(),
	)
	if addSigErr != nil {
		return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
	}
	k.SetWithdrawInfo(ctx, withdrawId, withdrawInfo)
	log.Infof("x/pegbridge SignWithdrawInfo add sig withdrawId:%x signer:%x sender:%s", withdrawId, validator.GetSignerAddr(), msg.Sender)
	return &types.MsgSignWithdrawResponse{}, nil
}

func (k msgServer) TriggerSignMint(goCtx context.Context, msg *types.MsgTriggerSignMint) (*types.MsgTriggerSignMintResponse, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil request")
	}
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	mindId := eth.Hex2Hash(msg.MintId)
	mintInfo, found := k.GetMintInfo(sdkCtx, mindId)
	if !found {
		return nil, types.WrapErrNoInfoFound(mindId)
	}
	if mintInfo.Success {
		return nil, fmt.Errorf("mint %x already completed", mindId)
	}
	now := sdkCtx.BlockTime()
	if now.Before(common.TsSecToTime(uint64(mintInfo.LastReqTime)).Add(k.Keeper.GetTriggerSignCooldown(sdkCtx))) {
		return nil, fmt.Errorf("request too soon")
	}
	// remove all previous sigs
	mintInfo.Signatures = nil
	mintInfo.LastReqTime = now.Unix()
	k.SetMintInfo(sdkCtx, mindId, mintInfo)
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMintToSign,
		sdk.NewAttribute(types.AttributeKeyMintId, msg.MintId),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	log.Infof("x/pegbr trigger sign mint, mintId %x", mindId)

	return &types.MsgTriggerSignMintResponse{}, nil
}

func (k msgServer) TriggerSignWithdraw(goCtx context.Context, msg *types.MsgTriggerSignWithdraw) (*types.MsgTriggerSignWithdrawResponse, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil request")
	}
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	withdrawId := eth.Hex2Hash(msg.WithdrawId)
	wdInfo, found := k.GetWithdrawInfo(sdkCtx, withdrawId)
	if !found {
		return nil, types.WrapErrNoInfoFound(withdrawId)
	}
	if wdInfo.Success {
		return nil, fmt.Errorf("withdraw %x already completed", withdrawId)
	}
	now := sdkCtx.BlockTime()
	if now.Before(common.TsSecToTime(uint64(wdInfo.LastReqTime)).Add(k.Keeper.GetTriggerSignCooldown(sdkCtx))) {
		return nil, fmt.Errorf("request too soon")
	}
	// remove all previous sigs
	wdInfo.Signatures = nil
	wdInfo.LastReqTime = now.Unix()
	k.SetWithdrawInfo(sdkCtx, withdrawId, wdInfo)
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeWithdrawToSign,
		sdk.NewAttribute(types.AttributeKeyWithdrawId, msg.WithdrawId),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	log.Infof("x/pegbr trigger sign Withdraw, withdrawId %x", withdrawId)

	return &types.MsgTriggerSignWithdrawResponse{}, nil
}

func (k msgServer) ClaimFee(goCtx context.Context, msg *types.MsgClaimFee) (*types.MsgClaimFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	logMsg := "x/pegbridge handle ClaimFee"
	withdrawAddr := eth.Hex2Addr(msg.DelegatorAddress)
	if k.HasFeeClaimInfo(ctx, withdrawAddr, msg.Nonce) {
		return nil, errors.New("fee claim nonce used")
	}

	signer, err := ethutils.RecoverSigner(msg.EncodeDataToSignByDelegator(), msg.Signature)
	if err != nil {
		return nil, fmt.Errorf("recover signer err: %w", err)
	}
	if signer != withdrawAddr {
		return nil, fmt.Errorf("%s invalid signature", logMsg)
	}

	tokenAddr := eth.Hex2Addr(msg.TokenAddress)
	withdrawAmt, withdrawOnChain, err := k.claimFee(ctx, withdrawAddr, msg.ChainId, tokenAddr, msg.Nonce, msg.Sender)
	if err != nil {
		return nil, err
	}
	withdrawProtoBytes, err := withdrawOnChain.Marshal()
	if err != nil {
		return nil, err
	}
	withdrawId := types.CalcWithdrawId(
		withdrawAddr,
		tokenAddr,
		withdrawAmt,
		eth.Addr{}, /* burnAccount */
		withdrawOnChain.RefChainId,
		eth.Bytes2Hash(withdrawOnChain.RefId))
	withdrawInfo := types.WithdrawInfo{
		ChainId:            msg.ChainId,
		WithdrawProtoBytes: withdrawProtoBytes,
		LastReqTime:        ctx.BlockTime().Unix(),
	}
	// Record WithdrawInfo and FeeClaimInfo
	k.SetWithdrawInfo(ctx, withdrawId, withdrawInfo)
	feeClaimInfo := types.FeeClaimInfo{
		WithdrawId: withdrawId.Bytes(),
	}
	k.SetFeeClaimInfo(ctx, withdrawAddr, msg.Nonce, feeClaimInfo)

	// Emit event for validators to sign
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeWithdrawToSign,
		sdk.NewAttribute(types.AttributeKeyWithdrawId, withdrawId.Hex()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))

	return &types.MsgClaimFeeResponse{}, nil
}

func (k msgServer) isSenderBondedValidator(ctx sdk.Context, sender string) (stakingtypes.ValidatorI, error) {
	senderAcct, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %s", sender)
	}
	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(ctx, senderAcct)
	if !found {
		return nil, fmt.Errorf("sender is not a validator")
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("validator is not bonded")
	}
	return validator, nil
}

func (k msgServer) ClaimRefund(goCtx context.Context, msg *types.MsgClaimRefund) (*types.MsgClaimRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	depositId := eth.Hex2Hash(msg.DepositId)
	depositInfo, found := k.GetDepositInfo(ctx, depositId)
	if !found {
		return nil, types.WrapErrNoInfoFound(depositId)
	}
	if len(depositInfo.MintId) > 0 {
		// a non-empty mintId indicates a valid deposit.
		return nil, fmt.Errorf("there is no refund for this deposit:%s", msg.DepositId)
	}
	// get depositRefund:withdrawOnChain
	withdraw, found := k.GetDepositRefund(ctx, depositId)
	if !found {
		// this refund has already been claimed.
		return nil, fmt.Errorf("failed to fetch deposit refund, no withdrawOnChain found for %s", msg.DepositId)
	}
	// check claimRefund requester is withdraw.Receiver
	signer, err := ethutils.RecoverSigner(eth.Hex2Bytes(msg.DepositId), msg.Signature)
	if err != nil {
		return nil, fmt.Errorf("recover signer err: %w", err)
	}
	if signer != eth.Bytes2Addr(withdraw.Receiver) {
		return nil, fmt.Errorf("invalid signature")
	}
	withdrawId := types.CalcWithdrawId(eth.Bytes2Addr(withdraw.Receiver), eth.Bytes2Addr(withdraw.Token),
		new(big.Int).SetBytes(withdraw.Amount), eth.Bytes2Addr(withdraw.BurnAccount), withdraw.RefChainId, eth.Bytes2Hash(withdraw.RefId))
	// record a withdrawInfo
	withdrawProtoBytes := k.cdc.MustMarshal(&withdraw)
	wdInfo := types.WithdrawInfo{
		ChainId:            withdraw.RefChainId,
		WithdrawProtoBytes: withdrawProtoBytes,
		Signatures:         make([]commontypes.Signature, 0),
		BaseFee:            "",
		PercentageFee:      "",
		LastReqTime:        ctx.BlockTime().Unix(),
	}
	k.SetWithdrawInfo(ctx, withdrawId, wdInfo)
	// record a refundClaimInfo
	// Although the invalid depositId is stored deeply in withdrawInfo.
	// We'd like to keep a forward link from invalid depositId to its corresponding withdrawId.
	k.SetRefundClaimInfo(ctx, depositId, withdrawId)
	// delete depositRefund:withdrawOnChain in case of double refunding
	k.DeleteDepositRefund(ctx, depositId)
	// emit event for validators to sign
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeWithdrawToSign,
		sdk.NewAttribute(types.AttributeKeyWithdrawId, withdrawId.Hex()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	log.Infof("x/pegbr claim refund, depositId: %x, withdrawId: %x, sender: %s",
		depositId, withdrawId, msg.Sender)
	return &types.MsgClaimRefundResponse{}, nil
}
