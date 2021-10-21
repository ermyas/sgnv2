package gatewaysvc

import (
	"context"
	"math"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingcli "github.com/celer-network/sgn-v2/x/farming/client/cli"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type rewardRecord struct {
	rwd   float64
	token *types.Token
}

func (gs *GatewayService) RewardingData(ctx context.Context, request *webapi.RewardingDataRequest) (*webapi.RewardingDataResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	unlockedCumulativeRewards, err := gs.getUnlockedCumulativeRewards(ctx, addr)
	if err != nil {
		log.Errorf("getUnlockedCumulativeRewards err:%+V", err)
	}
	historicalCumulativeRewards, usdPriceMap, err := gs.getHistoricalCumulativeRewards(ctx, addr)
	if err != nil {
		log.Errorf("getHistoricalCumulativeRewards err:%+V", err)
	}
	return &webapi.RewardingDataResponse{
		UsdPrice:                    usdPriceMap,
		HistoricalCumulativeRewards: historicalCumulativeRewards,
		UnlockedCumulativeRewards:   unlockedCumulativeRewards,
	}, nil
}

func (gs *GatewayService) UnlockFarmingReward(ctx context.Context, request *webapi.UnlockFarmingRewardRequest) (*webapi.UnlockFarmingRewardResponse, error) {
	tr := gs.TP.GetTransactor()
	_, err := farmingcli.ClaimAllRewards(tr, &farmingtypes.MsgClaimAllRewards{
		Address: eth.Addr2Hex(common.Hex2Addr(request.GetAddr())),
		Sender:  tr.Key.GetAddress().String(),
	})
	if err != nil {
		return &webapi.UnlockFarmingRewardResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  err.Error(),
			},
		}, nil
	} else {
		dberr := dal.DB.InsertClaimWithdrawRewardLog(request.GetAddr())
		if dberr != nil {
			log.Errorf("InsertClaimWithdrawRewardLog failed, error:%+v", dberr)
		}
		return &webapi.UnlockFarmingRewardResponse{}, nil
	}
}

func (gs *GatewayService) GetFarmingRewardDetails(ctx context.Context, request *webapi.GetFarmingRewardDetailsRequest) (*webapi.GetFarmingRewardDetailsResponse, error) {
	tr := gs.TP.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.RewardClaimInfo(
		ctx,
		&farmingtypes.QueryRewardClaimInfoRequest{
			Address: common.Hex2Addr(request.GetAddr()).String(),
		},
	)
	if res == nil || err != nil {
		log.Warnf("check failed, error:%+v", err)
		return &webapi.GetFarmingRewardDetailsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "check failed",
			},
		}, nil
	}
	rewardClaimInfo := res.GetRewardClaimInfo()
	var claimDetails []*farmingtypes.RewardClaimDetails
	for _, detail := range rewardClaimInfo.GetRewardClaimDetailsList() {
		claimDetails = append(claimDetails, &farmingtypes.RewardClaimDetails{
			ChainId:                 detail.GetChainId(),
			CumulativeRewardAmounts: detail.GetCumulativeRewardAmounts(),
			RewardProtoBytes:        detail.GetRewardProtoBytes(),
			Signatures:              detail.GetSignatures(),
		})
	}
	return &webapi.GetFarmingRewardDetailsResponse{
		Details: claimDetails,
	}, nil
}

// ================================= internal method below =====================================

func (gs *GatewayService) getUnlockedCumulativeRewards(ctx context.Context, address string) ([]*webapi.Reward, error) {
	tr := gs.TP.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.RewardClaimInfo(
		ctx,
		&farmingtypes.QueryRewardClaimInfoRequest{
			Address: common.Hex2Addr(address).String(),
		},
	)
	var rewards []*webapi.Reward
	// TODO: Properly handle the case of no unlocked amounts
	if res == nil {
		if err != nil {
			log.Warnf("check failed, error:%+v", err)
		}
		// Populate with 0 amounts
		rewards, _, err = gs.getHistoricalCumulativeRewards(ctx, address)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		for _, reward := range rewards {
			reward.Amt = 0
		}
	} else {
		rewardClaimInfo := res.GetRewardClaimInfo()
		records := make(map[string]rewardRecord)
		for _, detail := range rewardClaimInfo.GetRewardClaimDetailsList() {
			rewardAmts := detail.GetCumulativeRewardAmounts()
			for _, rewardAmt := range rewardAmts {
				token, amt, parseErr := gs.getInfoFromFarmingReward(rewardAmt)
				if parseErr != nil {
					continue
				}
				r := records[token.Symbol]
				r.rwd += amt
				r.token = token
				records[token.Symbol] = r
			}
		}
		for _, rcd := range records {
			rewards = append(rewards, &webapi.Reward{
				Amt:   rcd.rwd,
				Token: rcd.token,
			})
		}
	}
	return rewards, nil
}

func (gs *GatewayService) getHistoricalCumulativeRewards(ctx context.Context, address string) ([]*webapi.Reward, map[string]float64, error) {
	tr := gs.TP.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.AccountInfo(
		ctx,
		&farmingtypes.QueryAccountInfoRequest{
			Address: address,
		},
	)
	var rewards []*webapi.Reward
	usdPriceMap := make(map[string]float64)
	if res == nil || err != nil {
		log.Warnf("check failed, error:%+v", err)
	} else {
		records := make(map[string]rewardRecord)
		accountInfo := res.GetAccountInfo()
		for _, reward := range accountInfo.GetCumulativeRewardAmounts() {
			token, amt, parseErr := gs.getInfoFromFarmingReward(reward)
			if parseErr != nil {
				continue
			}
			r := records[token.Symbol]
			r.rwd += amt
			r.token = token
			records[token.Symbol] = r
			usdPriceMap[token.Symbol], _ = gs.F.GetUsdPrice(token.Symbol)
		}
		for _, rcd := range records {
			rewards = append(rewards, &webapi.Reward{
				Amt:   rcd.rwd,
				Token: rcd.token,
			})

		}
	}
	return rewards, usdPriceMap, nil
}

func (gs *GatewayService) getInfoFromFarmingReward(reward sdk.DecCoin) (*types.Token, float64, error) {
	chainId, tokenSymbol, parseErr := common.ParseERC20TokenDenom(reward.GetDenom())
	if parseErr != nil {
		log.Errorf("parse token denom error, denom:%s, err:%+v", reward.GetDenom(), parseErr)
	}
	tokenSymbol = common.GetSymbolFromFarmingToken(tokenSymbol)
	token, found, dbErr := dal.DB.GetRewardTokenBySymbol(tokenSymbol, chainId)
	rewardFloat64 := 0.0
	if found && dbErr == nil {
		rewardFloat64 = formatDecimals(token, reward.Amount.MustFloat64())
	}
	return token, rewardFloat64, parseErr
}

// todo cache this @aric
func (gs *GatewayService) getFarmingApy(ctx context.Context) map[uint64]map[string]float64 {
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
		apy, err := gs.calcPoolApy(&pool)
		if err != nil {
			log.Error("getFarmingApy error", err)
			return nil
		}
		apysByToken := make(map[string]float64)
		stakeToken := pool.StakeToken
		stakeTokenSymbol := common.GetSymbolFromFarmingToken(stakeToken.GetSymbol())
		apysByToken[stakeTokenSymbol] = apy
		apysByChainId[stakeToken.GetChainId()] = apysByToken
	}
	return apysByChainId
}

// calcPoolApy calculates USD-based APY with the formula (1 + r)^n - 1, assuming 5 seconds block time and daily compounding.
// The returned APY is the sum from all the reward tokens of the pool.
func (gs *GatewayService) calcPoolApy(pool *farmingtypes.FarmingPool) (float64, error) {
	const n = 365
	const secondsPerDay = 86400

	// Calculate staked USD value
	stakeToken := pool.StakeToken
	stakeTokenSymbol := common.GetSymbolFromFarmingToken(pool.StakeToken.Symbol)
	totalStakedUsd, err := gs.calcUsdValue(stakeTokenSymbol, int(stakeToken.Decimals), pool.TotalStakedAmount.Amount.MustFloat64())
	if err != nil {
		log.Errorf("calcUsdValue %s error %s", stakeToken.Symbol, err)
		return 0.0, err
	}

	// Calculate apy for each reward token, and sum them up
	const sgnBlockTime = 5
	totalApy := 0.0
	if totalStakedUsd != 0 {
		for i, info := range pool.RewardTokenInfos {
			rewardPerBlock := info.RewardAmountPerBlock.MustFloat64()
			rewardPerDay := rewardPerBlock * secondsPerDay / float64(sgnBlockTime)
			rewardToken := &pool.RewardTokens[i]
			rewardUsdPerDay, err := gs.calcUsdValue(rewardToken.Symbol, int(rewardToken.Decimals), rewardPerDay)
			if err != nil {
				log.Errorf("calcUsdValue %s error %s", rewardToken.Symbol, err)
				return 0.0, err
			}
			apyForToken := math.Pow(1+rewardUsdPerDay/totalStakedUsd, n) - 1
			if apyForToken >= 9999999 { // limit the max to make it more sense and also to avoid +Inf in case
				apyForToken = 9999999
			}
			totalApy += apyForToken
		}
	} else {
		log.Debugf("farming totalStakedUsd is 0 on chain:%d, token: %s", pool.StakeToken.ChainId, pool.StakeToken.Symbol)
	}
	return totalApy, nil
}

func (gs *GatewayService) calcUsdValue(symbol string, decimals int, amount float64) (float64, error) {
	usdPrice, err := gs.F.GetUsdPrice(symbol)
	if err != nil {
		log.Errorf("unable to get price of token %s from token price cache: %s", symbol, err)
		return 0, err
	}
	usdValue := amount * usdPrice / math.Pow10((decimals))
	return usdValue, nil
}

func formatDecimals(token *types.Token, amount float64) float64 {
	return amount / math.Pow10(int(token.Decimal))
}
