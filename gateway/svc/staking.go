package gatewaysvc

import (
	"context"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/webapi"
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
