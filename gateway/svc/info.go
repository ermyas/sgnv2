package gatewaysvc

import (
	"context"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/spf13/viper"
)

func (gs *GatewayService) GetAdvancedInfo(ctx context.Context, request *webapi.GetAdvancedInfoRequest) (*webapi.GetAdvancedInfoResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	slippageSetting, found, err := dal.DB.GetSlippageSetting(addr)
	if !found || err != nil {
		log.Errorf("GetAdvancedInfo failed, err:%+v", err)
		return &webapi.GetAdvancedInfoResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "AdvancedInfo not found",
			},
		}, nil
	}
	return &webapi.GetAdvancedInfoResponse{
		SlippageTolerance: slippageSetting,
	}, nil
}

func (gs *GatewayService) SetAdvancedInfo(ctx context.Context, request *webapi.SetAdvancedInfoRequest) (*webapi.SetAdvancedInfoResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	err := dal.DB.UpsertSlippageSetting(addr, request.GetSlippageTolerance())
	if err == nil {
		return &webapi.SetAdvancedInfoResponse{}, nil
	} else {
		return &webapi.SetAdvancedInfoResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "update setting failed",
			},
		}, nil
	}
}

func (gs *GatewayService) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	chainTokenList, err := dal.DB.GetChainTokenList()
	if err != nil {
		return &webapi.GetTransferConfigsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get chain_token failed",
			},
			Chains:     nil,
			ChainToken: nil,
		}, nil
	}
	var chainIds []uint32
	for key := range chainTokenList {
		chainIds = append(chainIds, key)
	}
	chains, err := dal.DB.GetChainInfo(chainIds)
	chainFound := make(map[uint32]bool)
	for _, chain := range chains {
		chainFound[chain.Id] = true
	}
	for chainId, tokens := range chainTokenList {
		_, found := chainFound[chainId]
		if !found {
			chains = append(chains, unknownChain(chainId))
		}
		for _, token := range tokens.Token {
			enrichUnknownToken(token)
		}
	}
	if err != nil {
		return &webapi.GetTransferConfigsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get chain info failed",
			},
			Chains:     chains,
			ChainToken: chainTokenList,
		}, nil
	}
	return &webapi.GetTransferConfigsResponse{
		Err:                       nil,
		Chains:                    chains,
		ChainToken:                chainTokenList,
		FarmingRewardContractAddr: viper.GetString(common.FlagEthContractFarmingRewards),
	}, nil
}
