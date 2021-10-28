package gatewaysvc

import (
	"context"
	"math"
	"math/big"
	"sort"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/spf13/viper"
)

func (gs *GatewayService) GetAdvancedInfo(ctx context.Context, request *webapi.GetAdvancedInfoRequest) (*webapi.GetAdvancedInfoResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	return &webapi.GetAdvancedInfoResponse{
		SlippageTolerance: GetSlippage(addr),
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
	sort.SliceStable(lps, func(i, j int) bool {
		if lps[i].HasFarmingSessions {
			if lps[j].HasFarmingSessions {
				return lps[i].GetVolume_24H() > lps[j].GetVolume_24H()
			} else {
				return true
			}
		} else {
			if lps[j].HasFarmingSessions {
				return false
			} else {
				return lps[i].GetVolume_24H() > lps[j].GetVolume_24H()
			}
		}
	})
	return &webapi.GetLPInfoListResponse{
		LpInfo: lps,
	}, nil
}

func (gs *GatewayService) getFarmingApy(ctx context.Context) map[uint64]map[string]float64 {
	cache := GetFarmingApyCache()
	if cache != nil {
		log.Debugf("farming apy cache used")
		return cache
	}
	log.Debugf("farming apy cache not used")
	tr := gs.TP.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.Pools(
		ctx,
		&farmingtypes.QueryPoolsRequest{},
	)
	if err != nil {
		log.Error("getFarmingApy error", err)
		return nil
	}
	apysByChainId := make(map[uint64]map[string]float64) // map<chain_id, map<token_symbol, apy>>
	for _, pool := range res.GetPools() {
		apy, calErr := gs.calcPoolApy(&pool)
		if calErr != nil {
			log.Error("getFarmingApy error", err)
			return nil
		}
		apysByToken := make(map[string]float64)
		stakeToken := pool.StakeToken
		stakeTokenSymbol := cbrtypes.GetSymbolFromStakeToken(stakeToken.GetSymbol())
		apysByToken[stakeTokenSymbol] = apy
		apysByChainId[stakeToken.GetChainId()] = apysByToken
	}
	SetFarmingApyCache(apysByChainId)
	return apysByChainId
}

// todo cache this  @aric
func (gs *GatewayService) get24hTx() map[uint64]map[string]*txData {
	cache := GetTx24hCache()
	if cache != nil {
		log.Debugf("24h tx cache used")
		return cache
	}
	log.Debugf("24h tx cache not used")
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

func GetSlippage(addr string) uint32 {
	slippageSetting, found, err := dal.DB.GetSlippageSetting(addr)
	if !found || err != nil {
		slippageSetting = 5000 //default 500
	}
	return slippageSetting
}
