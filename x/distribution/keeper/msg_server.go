package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/armon/go-metrics"
	"github.com/gogo/protobuf/proto"

	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	ethproto "github.com/celer-network/sgn-v2/proto/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SetWithdrawAddress(goCtx context.Context, msg *types.MsgSetWithdrawAddress) (*types.MsgSetWithdrawAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	delegatorAddress := eth.Hex2Addr(msg.DelegatorAddress)
	withdrawAddress := eth.Hex2Addr(msg.WithdrawAddress)
	err := k.SetWithdrawAddr(ctx, delegatorAddress, withdrawAddress)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	)

	return &types.MsgSetWithdrawAddressResponse{}, nil
}

func (k msgServer) WithdrawDelegatorReward(goCtx context.Context, msg *types.MsgWithdrawDelegatorReward) (*types.MsgWithdrawDelegatorRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.withdrawDelegatorRewardForOneValidator(ctx, msg.DelegatorAddress, msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawDelegatorRewardResponse{}, nil
}

func (k msgServer) WithdrawValidatorCommission(goCtx context.Context, msg *types.MsgWithdrawValidatorCommission) (*types.MsgWithdrawValidatorCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr := eth.Hex2Addr(msg.ValidatorAddress)
	amount, err := k.Keeper.WithdrawValidatorCommission(ctx, valAddr)
	if err != nil {
		return nil, err
	}

	defer func() {
		for _, a := range amount {
			if a.Amount.IsInt64() {
				telemetry.SetGaugeWithLabels(
					[]string{"tx", "msg", "withdraw_commission"},
					float32(a.Amount.Int64()),
					[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
				)
			}
		}
	}()

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddress),
		),
	)

	return &types.MsgWithdrawValidatorCommissionResponse{}, nil
}

func (k msgServer) FundCommunityPool(goCtx context.Context, msg *types.MsgFundCommunityPool) (*types.MsgFundCommunityPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	if err := k.Keeper.FundCommunityPool(ctx, msg.Amount, depositor); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor),
		),
	)

	return &types.MsgFundCommunityPoolResponse{}, nil
}

func (k msgServer) ClaimAllStakingReward(goCtx context.Context, msg *types.MsgClaimAllStakingReward) (*types.MsgClaimAllStakingRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Check cooldown and update claim time
	claimInfo, err := k.checkCooldownAndUpdateClaimTime(ctx, msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	// 2. Withdraw reward for all validators
	delAddr := eth.Hex2Addr(msg.DelegatorAddress)
	err = k.withdrawAllDelegatorRewards(ctx, delAddr)
	if err != nil {
		return nil, err
	}

	// 3. Accumulate staking rewards into claimInfo
	err = k.accumulateStakingReward(ctx, delAddr, claimInfo)
	if err != nil {
		return nil, err
	}

	// 4. Emit claim_all event to trigger validators signing
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeClaimAllStakingReward,
		sdk.NewAttribute(types.AttributeKeyDelegatorAddress, msg.DelegatorAddress),
	))

	return &types.MsgClaimAllStakingRewardResponse{}, nil
}

func (k msgServer) checkCooldownAndUpdateClaimTime(ctx sdk.Context, addr string) (*types.StakingRewardClaimInfo, error) {
	// Reject if last claim is too recent
	blockTime := ctx.BlockTime()
	claimInfo, found := k.GetStakingRewardClaimInfo(ctx, eth.Hex2Addr(addr))
	if found && blockTime.Before(claimInfo.LastClaimTime.Add(k.GetClaimCooldown(ctx))) {
		return nil, types.ErrClaimCooldownNotPassed
	}
	// Initialize claimInfo if not present
	if !found {
		claimInfo = types.StakingRewardClaimInfo{
			Recipient: addr,
		}
	}
	claimInfo.LastClaimTime = blockTime
	k.SetStakingRewardClaimInfo(ctx, claimInfo)
	return &claimInfo, nil
}

// accumulateStakingReward updates StakingRewardClaimInfo
func (k msgServer) accumulateStakingReward(ctx sdk.Context, delAddr eth.Addr, claimInfo *types.StakingRewardClaimInfo) error {
	// 1. Update CumulativeRewardAmount
	withdrawAddr := k.GetDelegatorWithdrawAddr(ctx, delAddr)
	derivedRewardAccount := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, withdrawAddr)
	rewards := k.bankKeeper.GetAllBalances(ctx, derivedRewardAccount)
	if rewards.Empty() {
		// TODO: Check
		return errors.New("no reward")
	}
	for _, reward := range rewards {
		denom := reward.Denom
		// NOTE: Only accumulate staking reward token
		if denom == types.StakingRewardDenom {
			cumulativeReward := k.bankKeeper.GetBalance(ctx, derivedRewardAccount, denom)
			cumulativeRewardAmount := cumulativeReward.Amount
			// Set initial CumulativeRewardAmount
			if claimInfo.CumulativeRewardAmount == (sdk.DecCoin{}) {
				claimInfo.CumulativeRewardAmount = sdk.NewDecCoin(denom, sdk.ZeroInt())
			}
			existing := sdk.NewDecCoinFromDec(denom, claimInfo.CumulativeRewardAmount.Amount)
			updated := sdk.NewDecCoin(denom, cumulativeRewardAmount)
			if !existing.Amount.Equal(updated.Amount) {
				claimInfo.CumulativeRewardAmount =
					claimInfo.CumulativeRewardAmount.Sub(existing).Add(updated)
			}
			break
		}
	}
	// 2. Reconstruct RewardProtoBytes with updated CumulativeRewardAmount
	// Marshal RewardProtoBytes
	rewardProtoBytes, marshalErr := proto.Marshal(
		&ethproto.StakingReward{
			Recipient:              delAddr.Bytes(),
			CumulativeRewardAmount: claimInfo.CumulativeRewardAmount.Amount.RoundInt().BigInt().Bytes(),
		})
	if marshalErr != nil {
		return marshalErr
	}
	claimInfo.RewardProtoBytes = rewardProtoBytes
	// 3.1. Clear stale signatures
	claimInfo.Signatures = []commontypes.Signature{}
	// 3.2. Set RewardClaimInfo
	k.SetStakingRewardClaimInfo(ctx, *claimInfo)
	return nil
}

func (k msgServer) SignStakingReward(
	goCtx context.Context, msg *types.MsgSignStakingReward) (*types.MsgSignStakingRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAcct, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	validator, found := k.stakingKeeper.GetValidatorBySgnAddr(ctx, senderAcct)
	if !found {
		return nil, fmt.Errorf("sender is not a validator")
	}
	if !validator.IsBonded() {
		return nil, fmt.Errorf("validator is not bonded")
	}

	claimInfo, found := k.GetStakingRewardClaimInfo(ctx, eth.Hex2Addr(msg.DelegatorAddress))
	if !found {
		return nil, err
	}
	addSigErr := claimInfo.AddSig(
		msg.Signature,
		validator.GetSignerAddr().String(),
	)
	if addSigErr != nil {
		return nil, fmt.Errorf("failed to add sig: %s", addSigErr)
	}
	k.SetStakingRewardClaimInfo(ctx, claimInfo)
	return &types.MsgSignStakingRewardResponse{}, nil
}
