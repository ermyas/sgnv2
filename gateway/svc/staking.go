package gatewaysvc

import (
	"context"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/distribution/client/cli"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/spf13/viper"
)

func (gs *GatewayService) StakingConfig(ctx context.Context, request *webapi.StakingConfigRequest) (*webapi.StakingConfigResponse, error) {
	return &webapi.StakingConfigResponse{
		ViewerContract:        viper.GetString(common.FlagEthContractViewer),
		StakingContract:       viper.GetString(common.FlagEthContractStaking),
		StakingRewardContract: viper.GetString(common.FlagEthContractStakingReward),
		CelrContract:          viper.GetString(common.FlagEthContractCelr),
	}, nil
}

func (gs *GatewayService) UnlockStakingReward(ctx context.Context, request *webapi.UnlockStakingRewardRequest) (*webapi.UnlockStakingRewardResponse, error) {
	tr := onchain.SGNTransactors.GetTransactor()
	if !utils.CheckUnlockStakingRewardParams(request.GetDelegatorAddress()) {
		log.Warnf("Unlock Staking Reward failed, param check failed")
		return &webapi.UnlockStakingRewardResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "params checking failed",
			},
		}, nil
	}
	_, err := cli.ClaimAllStakingReward(tr, &types.MsgClaimAllStakingReward{
		DelegatorAddress: eth.Addr2Hex(eth.Hex2Addr(request.GetDelegatorAddress())),
		Sender:           tr.Key.GetAddress().String(),
	})
	if err != nil {
		return &webapi.UnlockStakingRewardResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  err.Error(),
			},
		}, nil
	} else {
		return &webapi.UnlockStakingRewardResponse{}, nil
	}
}

func (gs *GatewayService) GetStakingRewardDetails(ctx context.Context, request *webapi.GetStakingRewardDetailsRequest) (*webapi.GetStakingRewardDetailsResponse, error) {
	tr := onchain.SGNTransactors.GetTransactor()
	queryClient := types.NewQueryClient(tr.CliCtx)
	res, err := queryClient.StakingRewardClaimInfo(
		ctx,
		&types.QueryStakingRewardClaimInfoRequest{
			DelegatorAddress: eth.Hex2Addr(request.GetDelegatorAddress()).String(),
		},
	)
	if res == nil || err != nil {
		log.Warnf("check failed, error:%+v", err)
		return &webapi.GetStakingRewardDetailsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "check failed",
			},
		}, nil
	}
	rewardClaimInfo := res.GetRewardClaimInfo()
	return &webapi.GetStakingRewardDetailsResponse{
		Detail: &types.StakingRewardClaimInfo{
			Recipient:              rewardClaimInfo.GetRecipient(),
			LastClaimTime:          rewardClaimInfo.GetLastClaimTime(),
			CumulativeRewardAmount: rewardClaimInfo.GetCumulativeRewardAmount(),
			RewardProtoBytes:       rewardClaimInfo.GetRewardProtoBytes(),
			Signatures:             rewardClaimInfo.GetSignatures(),
		},
	}, nil
}
