package keeper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	disttypes "github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/celer-network/sgn-v2/x/message/types"
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

func (k Keeper) SignMessage(goCtx context.Context, msg *types.MsgSignMessage) (*types.MsgSignMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator, err := k.stakingKeeper.CheckSenderBondedValidator(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}
	messageId := eth.Hex2Hash(msg.MessageId)
	message, found := k.GetMessage(ctx, messageId)
	if !found {
		return nil, types.WrapErrNoMessageFound(messageId)
	}
	bus, found := k.GetMessageBus(ctx, message.DstChainId)
	if !found {
		return nil, types.WrapErrNoMessageBusFound(message.DstChainId)
	}
	msgToSign := message.EncodeDataToSign(messageId, eth.Hex2Addr(bus.ContractInfo.Address))
	addSigErr := message.AddSig(
		msgToSign,
		msg.Signature,
		validator.GetSignerAddr(),
	)
	if addSigErr != nil {
		return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
	}
	k.SetMessage(ctx, messageId, &message)
	log.Infof("x/message SignMessage add sig messageId:%x signer:%x sender:%s", messageId, validator.GetSignerAddr(), msg.Sender)
	return &types.MsgSignMessageResponse{}, nil
}

func (k msgServer) TriggerSignMessage(
	goCtx context.Context, msg *types.MsgTriggerSignMessage) (*types.MsgTriggerSignMessageResponse, error) {
	return nil, nil
}

func (k msgServer) ClaimAllFees(
	goCtx context.Context, msg *types.MsgClaimAllFees) (*types.MsgClaimAllFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Check cooldown and update claim time
	claimInfo, err := k.checkCooldownAndUpdateClaimTime(ctx, msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	// 2. Claim message fees in distribution module, emitting event
	delAddr := eth.Hex2Addr(msg.DelegatorAddress)
	err = k.distrKeeper.ClaimMessageFees(ctx, delAddr)
	if err != nil {
		return nil, err
	}

	// 3. Accumulate fees into claimInfo
	err = k.accumulateFees(ctx, delAddr, claimInfo)
	if err != nil {
		return nil, err
	}

	return &types.MsgClaimAllFeesResponse{}, nil
}

func (k msgServer) SignFees(
	goCtx context.Context, msg *types.MsgSignFees) (*types.MsgSignFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAcct, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, fmt.Errorf("invalid sender %s", msg.Sender)
	}
	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(ctx, senderAcct)
	if !found {
		return nil, fmt.Errorf("sender is not a validator")
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("validator is not bonded")
	}

	claimInfo, found := k.GetFeeClaimInfo(ctx, eth.Hex2Addr(msg.Address))
	if !found {
		return nil, types.WrapErrNoClaimInfoFound(msg.Address)
	}
	if len(claimInfo.FeeClaimDetailsList) == 0 {
		return nil, errors.New("empty fee claim details list")
	}

	chainIdToFeeClaimDetails := make(map[uint64]*types.FeeClaimDetails)
	for i := 0; i < len(claimInfo.FeeClaimDetailsList); i++ {
		detail := &claimInfo.FeeClaimDetailsList[i]
		chainIdToFeeClaimDetails[detail.ChainId] = detail
	}
	for _, signatureDetails := range msg.SignatureDetailsList {
		messageBus, found := k.GetMessageBus(ctx, signatureDetails.ChainId)
		if !found {
			return nil, fmt.Errorf("MessageBus contract for chain %d not found", signatureDetails.ChainId)
		}
		claimDetails := chainIdToFeeClaimDetails[signatureDetails.ChainId]
		dataToSign := claimDetails.EncodeDataToSign(eth.Hex2Addr(messageBus.ContractInfo.Address), eth.Hex2Addr(msg.Address))
		addSigErr := claimDetails.AddSig(dataToSign, signatureDetails.Signature, validator.GetSignerAddr().String())
		if addSigErr != nil {
			return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
		}
	}
	k.SetFeeClaimInfo(ctx, claimInfo)
	log.Infof("x/message SignFees add sig address:%s signer:%x :sender:%s", msg.Address, validator.GetSignerAddr(), msg.Sender)
	return &types.MsgSignFeesResponse{}, nil
}

func (k msgServer) checkCooldownAndUpdateClaimTime(ctx sdk.Context, addr string) (*types.FeeClaimInfo, error) {
	// Reject if last claim is too recent
	blockTime := ctx.BlockTime()
	claimInfo, found := k.GetFeeClaimInfo(ctx, eth.Hex2Addr(addr))
	if found && blockTime.Before(claimInfo.LastClaimTime.Add(k.GetTriggerSignCooldown(ctx))) {
		return nil, types.WrapErrClaimCooldownNotPassed(claimInfo.LastClaimTime)
	}
	// Initialize claimInfo if not present
	if !found {
		claimInfo = types.FeeClaimInfo{
			Recipient: addr,
		}
	}
	claimInfo.LastClaimTime = blockTime
	k.SetFeeClaimInfo(ctx, claimInfo)
	return &claimInfo, nil
}

// accumulateFees updates FeeClaimInfo
func (k msgServer) accumulateFees(ctx sdk.Context, delAddr eth.Addr, claimInfo *types.FeeClaimInfo) error {
	// 1. Collect chainIds
	chainIdToDetails := make(map[uint64]*types.FeeClaimDetails)
	for _, detail := range claimInfo.FeeClaimDetailsList {
		chainIdToDetails[detail.ChainId] = &detail
	}
	// 2. Update CumulativeFeeAmount in details
	derivedFeeAccount := common.DeriveSdkAccAddressFromEthAddress(disttypes.ModuleName, delAddr)
	rewards := k.bankKeeper.GetAllBalances(ctx, derivedFeeAccount)
	log.Debugf("accumulateFees: delAddr %x, drived account %s, rewards %s", delAddr, derivedFeeAccount.String(), rewards.Sort().String())
	if rewards.Empty() {
		// TODO: Check
		return errors.New("no reward")
	}
	for _, reward := range rewards {
		denom := reward.Denom
		if !strings.HasPrefix(denom, types.MessageFeeDenomPrefix) {
			continue
		}
		chainId, _, err := common.ParseERC20TokenDenom(denom)
		if err != nil {
			return err
		}
		details, found := chainIdToDetails[chainId]
		if !found {
			// Create details if not existent
			details = &types.FeeClaimDetails{
				ChainId:             chainId,
				CumulativeFeeAmount: sdk.NewInt64DecCoin(denom, 0),
			}
			chainIdToDetails[chainId] = details
		}
		existing := sdk.NewDecCoinFromDec(denom, details.CumulativeFeeAmount.Amount)
		updated := sdk.NewDecCoin(denom, reward.Amount)
		if existing.Amount.LT(updated.Amount) {
			details.CumulativeFeeAmount = updated
		}
	}
	// 3.1. Append FeeClaimDetails
	// TODO: 1. Avoid copying 2. Sort by ascending chain IDs?
	claimInfo.FeeClaimDetailsList = []types.FeeClaimDetails{}
	for _, details := range chainIdToDetails {
		claimInfo.FeeClaimDetailsList = append(claimInfo.FeeClaimDetailsList, *details)
	}
	// 3.2. Clear stale signatures
	for i := 0; i < len(claimInfo.FeeClaimDetailsList); i++ {
		detail := &claimInfo.FeeClaimDetailsList[i]
		detail.Signatures = []commontypes.Signature{}
	}
	// 3.3. Set FeeClaimInfo
	k.SetFeeClaimInfo(ctx, *claimInfo)
	log.Infoln("x/message accumulateFees set FeeClaimInfo", claimInfo.LogStr())
	return nil
}
