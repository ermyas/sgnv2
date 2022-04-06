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
	// log.Infof("bridgeAddr:%x, msgToSign:%x", bridgeAddr, msgToSign)

	if commontypes.IsFlowChain(mintInfo.ChainId) {
		// TODO, only avoid call recover, find better way later.
		var exist bool
		for _, s := range mintInfo.Signatures {
			if s.Signer == validator.GetSignerAddr().String() {
				exist = true
				break
			}
		}
		if !exist {
			mintInfo.Signatures = append(mintInfo.Signatures,
				commontypes.NewSignature(validator.GetSignerAddr().String(), msg.Signature))
		}
	} else {
		addSigErr := mintInfo.AddSig(
			msgToSign,
			msg.Signature,
			validator.GetSignerAddr().String(),
		)
		if addSigErr != nil {
			return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
		}
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
	if commontypes.IsFlowChain(withdrawInfo.ChainId) {
		// TODO, only avoid call recover, find better way later.
		var exist bool
		for _, s := range withdrawInfo.Signatures {
			if s.Signer == validator.GetSignerAddr().String() {
				exist = true
				break
			}
		}
		if !exist {
			withdrawInfo.Signatures = append(withdrawInfo.Signatures,
				commontypes.NewSignature(validator.GetSignerAddr().String(), msg.Signature))
		}
	} else {
		addSigErr := withdrawInfo.AddSig(
			msgToSign,
			msg.Signature,
			validator.GetSignerAddr().String(),
		)
		if addSigErr != nil {
			return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
		}
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
	var withdrawId eth.Hash
	switch vaultVersion {
	case 0:
		withdrawId = types.CalcWithdrawId(
			withdrawAddr,
			tokenAddr,
			withdrawAmt,
			eth.Addr{}, /* burnAccount */
			withdrawOnChain.RefChainId,
			eth.Bytes2Hash(withdrawOnChain.RefId))
	case 2:
		vaultV2Addr, found := k.GetVersionedVault(ctx, msg.ChainId, 2)
		if !found {
			return nil, types.WrapErrNoOriginalTokenVaultFound(msg.ChainId)
		}
		withdrawId = types.CalcWithdrawIdV2(
			withdrawAddr,
			tokenAddr,
			withdrawAmt,
			eth.Addr{}, /* burnAccount */
			withdrawOnChain.RefChainId,
			eth.Bytes2Hash(withdrawOnChain.RefId),
			vaultV2Addr)
	default:
		return nil, fmt.Errorf("invalid vault version %d", vaultVersion)
	}
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
	_, isDeposit := k.GetDepositInfo(ctx, refId)
	var err error
	if isDeposit { //use deposit refund. Note that it is more efficient to pass depositInfo directly.
		//for maintainability, we only pass refId to keep the func claimDepositRefund clean
		err = k.claimDepositRefund(ctx, refId, msg)
	} else { //use burn refund
		err = k.claimBurnRefund(ctx, refId, msg)
	}
	if err != nil {
		return nil, err
	}
	return &types.MsgClaimRefundResponse{}, nil
}

func (k Keeper) claimDepositRefund(ctx sdk.Context, depositId eth.Hash, msg *types.MsgClaimRefund) error {
	depositInfo, found := k.GetDepositInfo(ctx, depositId)
	if !found {
		return fmt.Errorf("there is no refund for this deposit:%s", depositId.Hex())
	}
	if len(depositInfo.MintId) > 0 {
		// a non-empty mintId indicates a valid deposit.
		return fmt.Errorf("there is no refund for this deposit:%s", depositId.Hex())
	}
	// get depositRefund:withdrawOnChain
	withdraw, found := k.GetDepositRefund(ctx, depositId)
	if !found {
		// this refund has already been claimed.
		return fmt.Errorf("this deposit has already been refunded:%s", depositId.Hex())
	}
	var withdrawId eth.Hash
	if commontypes.IsFlowChain(depositInfo.ChainId) {
		wdRawData, err := withdraw.Marshal()
		if err != nil {
			return fmt.Errorf("fail to marshal flow WithdrawOnChain for refund, deposit:%v, err:%v", depositInfo, err)
		}
		withdrawId = types.CalcFlowWithdrawId(wdRawData)
	} else {
		switch depositInfo.GetVaultVersion() {
		case 0:
			withdrawId = types.CalcWithdrawId(
				eth.Bytes2Addr(withdraw.Receiver),
				eth.Bytes2Addr(withdraw.Token),
				new(big.Int).SetBytes(withdraw.Amount),
				eth.Bytes2Addr(withdraw.BurnAccount),
				withdraw.RefChainId,
				eth.Bytes2Hash(withdraw.RefId))
		case 2:
			vaultAddr, found := k.GetVersionedVault(ctx, depositInfo.ChainId, 2)
			if !found {
				return types.WrapErrNoOriginalTokenVaultFound(depositInfo.ChainId)
			}
			withdrawId = types.CalcWithdrawIdV2(
				eth.Bytes2Addr(withdraw.Receiver),
				eth.Bytes2Addr(withdraw.Token),
				new(big.Int).SetBytes(withdraw.Amount),
				eth.Bytes2Addr(withdraw.BurnAccount),
				withdraw.RefChainId,
				eth.Bytes2Hash(withdraw.RefId),
				vaultAddr)
		default:
			return fmt.Errorf("invalid vault version %d", depositInfo.GetVaultVersion())
		}
	}
	// record a withdrawInfo
	withdrawProtoBytes := k.cdc.MustMarshal(&withdraw)
	wdInfo := types.WithdrawInfo{
		ChainId:            withdraw.RefChainId,
		WithdrawProtoBytes: withdrawProtoBytes,
		Signatures:         make([]commontypes.Signature, 0),
		BaseFee:            "",
		PercentageFee:      "",
		LastReqTime:        ctx.BlockTime().Unix(),
		VaultVersion:       depositInfo.GetVaultVersion(),
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
	return nil
}

func (k Keeper) claimBurnRefund(ctx sdk.Context, burnId eth.Hash, msg *types.MsgClaimRefund) error {
	burnInfo, isBurn := k.GetBurnInfo(ctx, burnId)
	if !isBurn {
		return types.WrapErrNoInfoFound(burnId)
	}
	if len(burnInfo.WithdrawId) > 0 {
		// a non-empty withdrawId indicates a valid burn.
		return fmt.Errorf("there is no refund for this burn:%s", burnId.Hex())
	}
	// get burnRefund:mintOnChain
	mint, found := k.GetBurnRefund(ctx, burnId)
	if !found {
		// this refund has already been claimed.
		return fmt.Errorf("this burn has already been refunded:%s", burnId.Hex())
	}
	mintAmount := new(big.Int).SetBytes(mint.Amount)
	var mintId eth.Hash
	if commontypes.IsFlowChain(burnInfo.ChainId) {
		mintRawData, err := mint.Marshal()
		if err != nil {
			return fmt.Errorf("fail to marshal flow MintOnChain for refund, burn:%v, err:%v", burnInfo, err)
		}
		mintId = types.CalcFlowMintId(mintRawData)
	} else {
		switch burnInfo.GetBridgeVersion() {
		case 0:
			mintId = types.CalcMintId(eth.Bytes2Addr(mint.Account), eth.Bytes2Addr(mint.Token),
				mintAmount, eth.Bytes2Addr(mint.Depositor), mint.RefChainId, eth.Bytes2Hash(mint.RefId))
		case 2: //V2 peg bridge logic
			bridgeV2Addr, found := k.GetVersionedBridge(ctx, burnInfo.ChainId, 2)
			if !found {
				return types.WrapErrNoPeggedTokenBridgeFound(burnInfo.ChainId)
			}
			mintId = types.CalcMintIdV2(eth.Bytes2Addr(mint.Account), eth.Bytes2Addr(mint.Token),
				mintAmount, eth.Bytes2Addr(mint.Depositor), mint.RefChainId, eth.Bytes2Hash(mint.RefId), bridgeV2Addr)
		default:
			return fmt.Errorf("invalid bridge version %d", burnInfo.GetBridgeVersion())
		}
	}
	// record a mintInfo
	mintProtoBytes := k.cdc.MustMarshal(&mint)
	mintInfo := types.MintInfo{
		ChainId:        mint.RefChainId,
		MintProtoBytes: mintProtoBytes,
		Signatures:     make([]commontypes.Signature, 0),
		BaseFee:        "",
		PercentageFee:  "",
		LastReqTime:    ctx.BlockTime().Unix(),
		BridgeVersion:  burnInfo.GetBridgeVersion(),
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
	return nil
}
