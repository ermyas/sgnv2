package keeper

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
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

	validator, err := k.stakingKeeper.CheckSenderBondedValidator(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}
	mindId := eth.Hex2Hash(msg.MintId)
	mintInfo, found := k.GetMintInfo(ctx, mindId)
	if !found {
		return nil, types.WrapErrNoInfoFound(mindId)
	}
	bridgeAddr, found := k.GetPeggedBridge(ctx, mintInfo.ChainId, mintInfo.BridgeVersion)
	if !found {
		return nil, types.WrapErrNoPeggedTokenBridgeFound(mintInfo.ChainId)
	}
	msgToSign := mintInfo.EncodeDataToSign(bridgeAddr)
	addSigErr := mintInfo.AddSig(
		msgToSign,
		msg.Signature,
		validator.GetSignerAddr().String(),
	)
	if addSigErr != nil {
		return nil, fmt.Errorf("failed to add sig for bridge %x: %s", bridgeAddr, addSigErr)
	}
	k.SetMintInfo(ctx, mindId, mintInfo)
	log.Infof("x/pegbridge SignMintInfo add sig mintId:%x signer:%x :sender:%s", mindId, validator.GetSignerAddr(), msg.Sender)
	return &types.MsgSignMintResponse{}, nil
}

func (k msgServer) SignWithdraw(goCtx context.Context, msg *types.MsgSignWithdraw) (*types.MsgSignWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator, err := k.stakingKeeper.CheckSenderBondedValidator(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}
	withdrawId := eth.Hex2Hash(msg.WithdrawId)
	withdrawInfo, found := k.GetWithdrawInfo(ctx, withdrawId)
	if !found {
		return nil, types.WrapErrNoInfoFound(withdrawId)
	}
	vaultAddr, found := k.GetOriginalVault(ctx, withdrawInfo.ChainId, withdrawInfo.VaultVersion)
	if !found {
		return nil, types.WrapErrNoOriginalTokenVaultFound(withdrawInfo.ChainId)
	}
	msgToSign := withdrawInfo.EncodeDataToSign(vaultAddr)
	addSigErr := withdrawInfo.AddSig(
		msgToSign,
		msg.Signature,
		validator.GetSignerAddr().String(),
	)
	if addSigErr != nil {
		return nil, fmt.Errorf("failed to add sig for vault %x: %s", vaultAddr, addSigErr)
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
	var withdrawAddr eth.Addr
	if msg.IsValidator {
		senderSgnAcct, err := sdk.AccAddressFromBech32(msg.Sender)
		if err != nil {
			return nil, fmt.Errorf("invalid sender accnt")
		}
		validator, found := k.stakingKeeper.GetValidatorBySgnAddr(ctx, senderSgnAcct)
		if !found {
			return nil, fmt.Errorf("sender accnt %s not validator", senderSgnAcct)
		}
		withdrawAddr = validator.GetEthAddr()
	} else {
		withdrawAddr = eth.Hex2Addr(msg.DelegatorAddress)
		signer, err := ethutils.RecoverSigner(msg.EncodeDataToSignByDelegator(), msg.Signature)
		if err != nil {
			return nil, fmt.Errorf("recover signer err: %w", err)
		}
		if signer != withdrawAddr {
			return nil, fmt.Errorf("%s invalid signature", logMsg)
		}
	}

	if k.HasFeeClaimInfo(ctx, withdrawAddr, msg.Nonce) {
		return nil, errors.New("fee claim nonce used")
	}

	tokenAddr := eth.Hex2Addr(msg.TokenAddress)
	withdrawAmt, withdrawOnChain, vaultVersion, err := k.claimFee(ctx, withdrawAddr, msg.ChainId, tokenAddr, msg.Nonce, msg.Sender)
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
		VaultVersion:       vaultVersion,
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

func (k msgServer) ClaimRefund(goCtx context.Context, msg *types.MsgClaimRefund) (*types.MsgClaimRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	refId := eth.Hex2Hash(msg.RefId)
	depositInfo, isDeposit := k.GetDepositInfo(ctx, refId)
	if !isDeposit {
		burnInfo, isBurn := k.GetBurnInfo(ctx, refId)
		if !isBurn {
			return nil, types.WrapErrNoInfoFound(refId)
		}
		burnId := refId
		if len(burnInfo.WithdrawId) > 0 {
			// a non-empty withdrawId indicates a valid burn.
			return nil, fmt.Errorf("there is no refund for this burn:%s", burnId.Hex())
		}
		// get burnRefund:mintOnChain
		mint, found := k.GetBurnRefund(ctx, burnId)
		if !found {
			// this refund has already been claimed.
			return nil, fmt.Errorf("this burn has already been refunded:%s", burnId.Hex())
		}
		mintAmount := new(big.Int).SetBytes(mint.Amount)

		mintId := types.CalcMintId(eth.Bytes2Addr(mint.Account), eth.Bytes2Addr(mint.Token),
			mintAmount, eth.Bytes2Addr(mint.Depositor), mint.RefChainId, eth.Bytes2Hash(mint.RefId))
		// record a mintInfo
		mintProtoBytes := k.cdc.MustMarshal(&mint)
		mintInfo := types.MintInfo{
			ChainId:        mint.RefChainId,
			MintProtoBytes: mintProtoBytes,
			Signatures:     make([]commontypes.Signature, 0),
			BaseFee:        "",
			PercentageFee:  "",
			LastReqTime:    ctx.BlockTime().Unix(),
		}
		k.SetMintInfo(ctx, mintId, mintInfo)
		// record a refundClaimInfo
		// Although the invalid burnId is stored deeply in mintInfo.
		// We'd still like to keep a forward link from invalid burnId to its corresponding mintId.
		k.SetRefundClaimInfo(ctx, burnId, mintId)
		// delete burnRefund:mintOnChain in case of double refunding
		k.DeleteRefund(ctx, burnId)
		// emit event for validators to sign
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeMintToSign,
			sdk.NewAttribute(types.AttributeKeyMintId, mintId.Hex()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		))
		log.Infof("x/pegbr claim refund, burnId: %x, mintId: %x, sender: %s",
			burnId, mintId, msg.Sender)
		return &types.MsgClaimRefundResponse{}, nil
	} else {
		depositId := refId
		if len(depositInfo.MintId) > 0 {
			// a non-empty mintId indicates a valid deposit.
			return nil, fmt.Errorf("there is no refund for this deposit:%s", depositId.Hex())
		}
		// get depositRefund:withdrawOnChain
		withdraw, found := k.GetDepositRefund(ctx, depositId)
		if !found {
			// this refund has already been claimed.
			return nil, fmt.Errorf("this deposit has already been refunded:%s", depositId.Hex())
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
		// We'd still like to keep a forward link from invalid depositId to its corresponding withdrawId.
		k.SetRefundClaimInfo(ctx, depositId, withdrawId)
		// delete depositRefund:withdrawOnChain in case of double refunding
		k.DeleteRefund(ctx, depositId)
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
}
