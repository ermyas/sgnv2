package gatewaysvc

import (
	"context"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingcli "github.com/celer-network/sgn-v2/x/farming/client/cli"
	farmingkp "github.com/celer-network/sgn-v2/x/farming/keeper"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"math/big"

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
	historicalCumulativeRewards, totalVolunme, err := gs.getHistoricalCumulativeRewards(ctx, addr)
	if err != nil {
		log.Errorf("getHistoricalCumulativeRewards err:%+V", err)
	}
	return &webapi.RewardingDataResponse{
		TotalFarmingRewards:         totalVolunme,
		HistoricalCumulativeRewards: historicalCumulativeRewards,
		UnlockedCumulativeRewards:   unlockedCumulativeRewards,
	}, nil
}

func (gs *GatewayService) ClaimWithdrawReward(ctx context.Context, request *webapi.ClaimWithdrawRewardRequest) (*webapi.ClaimWithdrawRewardResponse, error) {
	tr := gs.TP.GetTransactor()
	_, err := farmingcli.ClaimAllRewards(tr, &farmingtypes.MsgClaimAllRewards{
		Address: eth.Addr2Hex(common.Hex2Addr(request.GetAddr())),
		Sender:  tr.Key.GetAddress().String(),
	})
	if err != nil {
		return &webapi.ClaimWithdrawRewardResponse{
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
		return &webapi.ClaimWithdrawRewardResponse{}, nil
	}
}

func (gs *GatewayService) ClaimRewardDetails(ctx context.Context, request *webapi.ClaimRewardDetailsRequest) (*webapi.ClaimRewardDetailsResponse, error) {
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
		return &webapi.ClaimRewardDetailsResponse{
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
			ChainId:          detail.GetChainId(),
			RewardProtoBytes: detail.GetRewardProtoBytes(),
			Signatures:       detail.GetSignatures(),
		})
	}
	return &webapi.ClaimRewardDetailsResponse{
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
	if res == nil || err != nil {
		log.Warnf("check failed, error:%+v", err)
	} else {
		rewardClaimInfo := res.GetRewardClaimInfo()
		records := make(map[string]rewardRecord)
		for _, detail := range rewardClaimInfo.GetRewardClaimDetailsList() {
			rewardAmts := detail.GetCumulativeRewardAmounts()
			for _, rewardAmt := range rewardAmts {
				token, amt, _, parseErr := gs.getInfoFromFarmingReward(rewardAmt)
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

func (gs *GatewayService) getHistoricalCumulativeRewards(ctx context.Context, address string) ([]*webapi.Reward, float64, error) {
	tr := gs.TP.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.AccountInfo(
		ctx,
		&farmingtypes.QueryAccountInfoRequest{
			Address: address,
		},
	)
	var rewards []*webapi.Reward
	sumVolume := 0.0
	if res == nil || err != nil {
		log.Warnf("check failed, error:%+v", err)
	} else {
		records := make(map[string]rewardRecord)
		accountInfo := res.GetAccountInfo()
		for _, reward := range accountInfo.GetCumulativeRewardAmounts() {
			token, amt, volume, parseErr := gs.getInfoFromFarmingReward(reward)
			if parseErr != nil {
				continue
			}
			r := records[token.Symbol]
			r.rwd += amt
			r.token = token
			records[token.Symbol] = r
			sumVolume += volume
		}
		for _, rcd := range records {
			rewards = append(rewards, &webapi.Reward{
				Amt:   rcd.rwd,
				Token: rcd.token,
			})

		}
	}
	return rewards, sumVolume, nil
}

func (gs *GatewayService) getInfoFromFarmingReward(reward sdk.DecCoin) (*types.Token, float64, float64, error) {
	chainId, tokenSymbol, parseErr := farmingkp.ParseERC20TokenDenom(reward.GetDenom())
	if parseErr != nil {
		log.Errorf("parse token denom error, denom:%s, err:%+v", reward.GetDenom(), parseErr)
	}
	tokenSymbol = common.GetSymbolFromFarmingToken(tokenSymbol)
	token, found, dbErr := dal.DB.GetRewardTokenBySymbol(tokenSymbol, chainId)
	rwd := 0.0
	amtInt := new(big.Int)
	if found && dbErr == nil {
		amt, parsed := new(big.Float).SetString(reward.Amount.String())
		if parsed {
			rwd, _ = new(big.Float).Quo(amt, new(big.Float).SetInt64(int64(token.GetDecimal()))).Float64()
			amtInt = common.FloatToBigInt(amt)
		}
	}
	return token, rwd, gs.F.GetUsdVolume(token, amtInt), parseErr
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
		return nil
	}
	farmingPools := make(map[uint64]map[string]float64) // map<chain_id, map<token_symbol, FarmingPool>>
	for _, pool := range res.GetPools() {
		farmingPool := make(map[string]float64)
		token := pool.GetStakeToken()
		tokenSymbol := common.GetSymbolFromFarmingToken(token.GetSymbol())
		totalStakedAmount := pool.TotalStakedAmount
		if totalStakedAmount.Amount.Equal(sdk.ZeroDec()) {
			log.Debugf("farming totalStakedAmount is 0 on chain:%d, token: %s", token.GetChainId(), tokenSymbol)
			farmingPool[tokenSymbol] = 0.0
		} else {
			totalReward := sdk.ZeroDec()
			for _, reward := range pool.GetRewardTokenInfos() {
				totalReward = totalReward.Add(reward.RewardAmountPerBlock)
			}

			// apy=totalReward/totalStakedAmount
			apy, _ := new(big.Float).Quo(new(big.Float).SetInt(totalReward.BigInt()), new(big.Float).SetInt(totalStakedAmount.Amount.BigInt())).Float64()
			farmingPool[tokenSymbol] = apy
		}
		farmingPools[token.GetChainId()] = farmingPool
	}
	return farmingPools
}
