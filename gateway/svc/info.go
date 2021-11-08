package gatewaysvc

import (
	"context"
	"math"
	"math/big"
	"sort"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/spf13/viper"
)

func (gs *GatewayService) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	chainTokenList, err := dal.DB.GetEnabledChainTokenList()
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
		enrichChainUiInfo(chain)
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

func (gs *GatewayService) GetTokenInfo(ctx context.Context, request *webapi.GetTokenInfoRequest) (*webapi.GetTokenInfoResponse, error) {
	chainId := uint64(request.GetChainId())
	tokenInfo, found, err := dal.DB.GetTokenBySymbol(request.GetTokenSymbol(), chainId)
	if tokenInfo != nil && found && err == nil {
		return &webapi.GetTokenInfoResponse{
			TokenInfo: tokenInfo,
		}, nil
	}
	// if bridge token not found, try to find reward token
	token, found, err := dal.DB.GetRewardTokenBySymbol(request.GetTokenSymbol(), chainId)
	if token != nil && found && err == nil {
		tokenInfo = &webapi.TokenInfo{
			Token: token,
			Name:  token.Symbol,
			Icon:  "",
		}
		enrichUnknownToken(tokenInfo)
		return &webapi.GetTokenInfoResponse{
			TokenInfo: tokenInfo,
		}, nil
	}
	return &webapi.GetTokenInfoResponse{
		Err: &webapi.ErrMsg{
			Code: webapi.ErrCode_ERROR_CODE_COMMON,
			Msg:  "token not found",
		}}, nil
}

func (gs *GatewayService) GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error) {
	userAddr := common.Hex2Addr(request.GetAddr()).String()
	chainTokenInfos, err := dal.DB.GetChainTokenList()
	if err != nil || len(chainTokenInfos) == 0 {
		return &webapi.GetLPInfoListResponse{}, nil
	}
	var chainTokens []*cbrtypes.ChainTokenAddrPair
	for chainId, tokens := range chainTokenInfos {
		for _, tokenInfo := range tokens.Token {
			chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
				ChainId:   uint64(chainId),
				TokenAddr: tokenInfo.GetToken().Address,
			})
		}
	}
	var lps []*webapi.LPInfo

	userDetailMap := make(map[uint64]map[string]*cbrtypes.LiquidityDetail)
	hasUsr := request.GetAddr() != ""
	if hasUsr {
		tr := gs.TP.GetTransactor()
		detailList, detailErr := cbrcli.QueryLiquidityDetailList(tr.CliCtx, &cbrtypes.LiquidityDetailListRequest{
			LpAddr:     userAddr,
			ChainToken: chainTokens,
		})
		if detailList == nil || detailErr != nil {
			var emptyLiquidityDetail []*cbrtypes.LiquidityDetail
			detailList = &cbrtypes.LiquidityDetailListResponse{LiquidityDetail: emptyLiquidityDetail}
		}
		for _, detail := range detailList.GetLiquidityDetail() {
			chainId := detail.GetChainId()
			tokenWithAddr := detail.GetToken() // only has addr field
			token, found, dbErr := dal.DB.GetTokenByAddr(common.Hex2Addr(tokenWithAddr.GetAddress()).String(), chainId)
			if !found || dbErr != nil {
				log.Debugf("data, token not found in lp list, token addr:%s, chainId:%d", tokenWithAddr.GetAddress(), chainId)
				continue
			}
			detail.Token = token.Token
			chainInfo, found := userDetailMap[chainId]
			if !found {
				chainInfo = make(map[string]*cbrtypes.LiquidityDetail)
			}
			chainInfo[token.Token.Symbol] = detail
			userDetailMap[chainId] = chainInfo
		}
	}

	farmingApyMap := gs.getFarmingApy(ctx)
	data24h := gs.get24hTx()

	for chainId32, chainToken := range chainTokenInfos {
		chainId := uint64(chainId32)
		for _, token := range chainToken.Token {
			tokenSymbol := token.Token.Symbol
			totalLiquidity := "0"
			usrLpFeeEarning := "0"
			usrLiquidity := "0"
			_, found1 := userDetailMap[chainId]
			if found1 {
				detail, found2 := userDetailMap[chainId][tokenSymbol]
				if found2 {
					totalLiquidity = detail.GetTotalLiquidity()
					usrLpFeeEarning = detail.GetUsrLpFeeEarning()
					usrLiquidity = detail.GetUsrLiquidity()
				}
			}

			enrichUnknownToken(token)
			chain, _, found, dbErr := dal.DB.GetChain(chainId)
			if !found || dbErr != nil {
				chain = unknownChain(chainId32)
			} else {
				chain = enrichChainUiInfo(chain)
			}

			data := data24h[chainId][tokenSymbol]
			lpFeeEarningApy := 0.0
			volume24h := 0.0
			if data != nil {
				if common.Str2BigInt(totalLiquidity).Cmp(new(big.Int).SetInt64(0)) > 0 {
					rate, _ := new(big.Float).Quo(new(big.Float).SetInt(data.fee), new(big.Float).SetInt(common.Str2BigInt(totalLiquidity))).Float64()
					lpFeeEarningApy = math.Pow(1+rate, 365) - 1
				}
				volume24h = data.volume
			}
			farmingApy, hasSession := farmingApyMap[chainId][token.Token.GetSymbol()]
			lp := &webapi.LPInfo{
				Chain:              chain,
				Token:              token,
				Liquidity:          gs.F.GetUsdVolume(token.Token, common.Str2BigInt(usrLiquidity)),
				LiquidityAmt:       usrLiquidity,
				HasFarmingSessions: hasSession,
				LpFeeEarning:       gs.F.GetUsdVolume(token.Token, common.Str2BigInt(usrLpFeeEarning)),
				Volume_24H:         volume24h,
				TotalLiquidity:     gs.F.GetUsdVolume(token.Token, common.Str2BigInt(totalLiquidity)),
				TotalLiquidityAmt:  totalLiquidity,
				LpFeeEarningApy:    lpFeeEarningApy,
				FarmingApy:         farmingApy,
			}
			lps = append(lps, lp)
		}
	}
	sortLpList(lps)
	return &webapi.GetLPInfoListResponse{
		LpInfo: lps,
	}, nil
}

func (gs *GatewayService) GetTotalLiquidityProviderTokenBalance(ctx context.Context, request *webapi.GetTotalLiquidityProviderTokenBalanceRequest) (*webapi.GetTotalLiquidityProviderTokenBalanceResponse, error) {
	tokenSymbol := request.GetTokenSymbol()
	chainIds := request.GetChainIds()
	ret := make(map[uint64]string)
	if len(chainIds) == 0 {
		// all chains
		chainTokenInfos, err := dal.DB.GetChainTokenList()
		if err == nil && len(chainTokenInfos) > 0 {
			for chainId, chainToken := range chainTokenInfos {
				for _, token := range chainToken.GetToken() {
					if token.GetToken().GetSymbol() == tokenSymbol {
						ret[uint64(chainId)] = gs.getLiquidityOnChainToken(uint64(chainId), token.GetToken().GetAddress())
						break
					}
				}
			}
		}
	} else {
		for _, chainId := range chainIds {
			token, tokenFound, dberr := dal.DB.GetTokenBySymbol(tokenSymbol, uint64(chainId))
			if tokenFound && dberr == nil {
				ret[uint64(chainId)] = gs.getLiquidityOnChainToken(uint64(chainId), token.GetToken().GetAddress())
			}
		}
	}
	return &webapi.GetTotalLiquidityProviderTokenBalanceResponse{
		TotalLiq: ret,
	}, nil
}

func (gs *GatewayService) getLiquidityOnChainToken(chainId uint64, tokenAddr string) string {
	tr := gs.TP.GetTransactor()
	resp, err := cbrcli.QueryTotalLiquidity(tr.CliCtx, &cbrtypes.QueryTotalLiquidityRequest{
		ChainId:   chainId,
		TokenAddr: tokenAddr,
	})
	if err != nil {
		log.Warnf("getLiquidityOnChainToken err, chain:%d, token:%s, err:%+v", chainId, tokenAddr, err)
		return "0"
	}
	return resp.GetTotalLiq()
}

func (gs *GatewayService) getFarmingApy(ctx context.Context) map[uint64]map[string]float64 {
	cache := GetFarmingApyCache()
	if cache != nil {
		return cache
	}
	tr := gs.TP.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.Pools(
		ctx,
		&farmingtypes.QueryPoolsRequest{},
	)
	if err != nil {
		log.Warnf("getFarmingApy error:%+v", err)
		return nil
	}
	apysByChainId := make(map[uint64]map[string]float64) // map<chain_id, map<token_symbol, apy>>
	for _, pool := range res.GetPools() {
		apy, calErr := gs.calcPoolApy(&pool)
		if calErr != nil {
			continue
		}
		stakeToken := pool.StakeToken
		apysByToken, exists := apysByChainId[stakeToken.GetChainId()]
		if !exists {
			apysByToken = make(map[string]float64)
		}
		stakeTokenSymbol := cbrtypes.GetSymbolFromStakeToken(stakeToken.GetSymbol())
		apysByToken[stakeTokenSymbol] = apy
		apysByChainId[stakeToken.GetChainId()] = apysByToken
	}
	SetFarmingApyCache(apysByChainId)
	return apysByChainId
}

func (gs *GatewayService) get24hTx() map[uint64]map[string]*txData {
	cache := GetTx24hCache()
	if cache != nil {
		return cache
	}
	txs, err := dal.DB.Get24hTx()
	resp := make(map[uint64]map[string]*txData) // map<chain_id, map<token_symbol, txData>>
	if err == nil {
		for _, tx := range txs {
			tokenSymbol := tx.TokenSymbol
			dstToken, found, dbErr := dal.DB.GetTokenBySymbol(tokenSymbol, tx.DstChainId)
			if !found || dbErr != nil {
				continue
			}
			dstChainId := tx.DstChainId
			data, found := resp[dstChainId]
			if !found || data == nil {
				data = make(map[string]*txData)
			}
			d, found := data[tokenSymbol]
			if !found || d == nil {
				d = &txData{
					volume:   0,
					fee:      new(big.Int),
					dstToken: dstToken.Token,
				}
			}
			feePerc := float64(tx.FeePerc) / 1e6
			feeAmt := new(big.Float).Mul(new(big.Float).SetInt(common.Str2BigInt(tx.DstAmt)), new(big.Float).SetFloat64(feePerc))
			feeAmtInt := new(big.Int)
			feeAmt.Int(feeAmtInt)
			d.fee = new(big.Int).Add(d.fee, feeAmtInt)
			d.volume += tx.Volume
			data[tokenSymbol] = d
			resp[tx.DstChainId] = data
		}
	}
	SetTx24hCache(resp)
	return resp
}

func sortLpList(lps []*webapi.LPInfo) {
	sort.SliceStable(lps, func(i, j int) bool {
		if lps[i].HasFarmingSessions {
			if lps[j].HasFarmingSessions {
				if lps[i].GetVolume_24H() == lps[j].GetVolume_24H() {
					return cmpChainToken(lps[i], lps[j])
				}
				return lps[i].GetVolume_24H() > lps[j].GetVolume_24H()
			} else {
				return true
			}
		} else {
			if lps[j].HasFarmingSessions {
				return false
			} else {
				if lps[i].GetVolume_24H() == lps[j].GetVolume_24H() {
					return cmpChainToken(lps[i], lps[j])
				}
				return lps[i].GetVolume_24H() > lps[j].GetVolume_24H()
			}
		}
	})
}

func cmpChainToken(lp1, lp2 *webapi.LPInfo) bool {
	cmpChain := strings.Compare(lp1.GetChain().GetName(), lp2.GetChain().GetName())
	cmpToken := strings.Compare(lp1.GetToken().GetName(), lp2.GetToken().GetName())
	if cmpChain == 0 {
		return cmpToken < 0
	} else {
		return cmpChain < 0
	}
}
