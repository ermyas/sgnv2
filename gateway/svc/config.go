package gatewaysvc

import (
	"context"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
)

func (gs *GatewayService) UpdateChain(ctx context.Context, request *webapi.UpdateChainRequest) (*webapi.UpdateChainResponse, error) {
	chainInput := request.GetChain()
	chainId := uint64(chainInput.GetId())

	if chainId == 0 {
		return &webapi.UpdateChainResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "chainId is 0",
			},
		}, nil
	}
	var name, icon, gasTokenSymbol, exploreUrl, rpcUrl string
	chainInDb, url, chainFound, err := dal.DB.GetChain(chainId)
	if chainInDb != nil && chainFound && err == nil {
		name = chainInDb.GetName()
		icon = chainInDb.GetIcon()
		gasTokenSymbol = chainInDb.GetGasTokenSymbol()
		exploreUrl = chainInDb.GetExploreUrl()
		rpcUrl = chainInDb.GetRpcUrl()
	}
	if chainInput.GetName() != "" {
		name = chainInput.GetName()
	}
	if chainInput.GetIcon() != "" {
		icon = chainInput.GetIcon()
	}
	if request.GetTxUrlPrefix() != "" {
		url = request.GetTxUrlPrefix()
	}
	if chainInput.GetGasTokenSymbol() != "" {
		gasTokenSymbol = chainInput.GetGasTokenSymbol()
	}
	if chainInput.GetExploreUrl() != "" {
		exploreUrl = chainInput.GetExploreUrl()
	}
	if chainInput.GetRpcUrl() != "" {
		rpcUrl = chainInput.GetRpcUrl()
	}

	dal.DB.UpsertChainUIInfo(chainId, name, icon, url, gasTokenSymbol, exploreUrl, rpcUrl)
	chainInDb, url, _, _ = dal.DB.GetChain(chainId)
	return &webapi.UpdateChainResponse{
		Chain:       chainInDb,
		TxUrlPrefix: url,
	}, nil
}

func (gs *GatewayService) UpdateToken(ctx context.Context, request *webapi.UpdateTokenRequest) (*webapi.UpdateTokenResponse, error) {
	chainId := uint64(request.GetChainId())
	tokenSymbol := request.GetTokenSymbol()
	if chainId == 0 || tokenSymbol == "" {
		return &webapi.UpdateTokenResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "invalid input, check chainId and tokenSymbol",
			},
		}, nil
	}
	var name, icon string
	tokenInDb, found, err := dal.DB.GetTokenBySymbol(tokenSymbol, chainId)
	if tokenInDb != nil && found && err == nil {
		name = tokenInDb.GetName()
		icon = tokenInDb.GetIcon()
	} else {
		return &webapi.UpdateTokenResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "token not found, can not update UI info now",
			},
		}, nil
	}
	if request.GetTokenName() != "" {
		name = request.GetTokenName()
	}
	if request.GetTokenIcon() != "" {
		icon = request.GetTokenIcon()
	}
	dal.DB.UpdateTokenUIInfo(tokenSymbol, chainId, name, icon)
	tokenInDb, _, _ = dal.DB.GetTokenBySymbol(tokenSymbol, chainId)
	return &webapi.UpdateTokenResponse{
		Token: tokenInDb,
	}, nil
}