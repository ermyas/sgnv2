package gatewaysvc

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	types3 "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/viper"
)

func (gs *GatewayService) GetRetentionRewardsInfo(ctx context.Context, request *webapi.GetRetentionRewardsInfoRequest) (*webapi.GetRetentionRewardsInfoResponse, error) {
	level, rewardAmt, claimTime, signature, found, err := dal.DB.GetRetentionRewardsRecord(request.GetAddr(), dal.RetentionRewardEventId)
	if err != nil {
		log.Errorln("failed to GetUserRetentionRewardsLevel:", request, " error:", err)
		return &webapi.GetRetentionRewardsInfoResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "db query fail",
			},
		}, nil
	}
	if !found {
		return &webapi.GetRetentionRewardsInfoResponse{}, nil
	}
	event := dal.RetentionRewardsConfig[dal.RetentionRewardEventId]
	levelConfig := event.LevelConfig[level]
	price, err := gs.F.GetUsdPrice("CELR")
	if err != nil {
		log.Errorln("failed to Get celr UsdPrice:", err)
	}

	if event.EventEndTime.Add(7 * 24 * time.Hour).Before(time.Now()) {
		return &webapi.GetRetentionRewardsInfoResponse{}, nil
	} else if claimTime.After(time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC)) {
		return &webapi.GetRetentionRewardsInfoResponse{
			EventId:           dal.RetentionRewardEventId,
			EventEndTime:      common.TsMilli(event.EventEndTime),
			MaxReward:         levelConfig.MaxReward.String(),
			MaxTransferVolume: levelConfig.MaxTransferVolume,
			CurrentReward:     rewardAmt.String(),
			CelrUsdPrice:      price,
			ClaimTime:         common.TsMilli(claimTime),
			Signature: types3.Signature{
				Signer:   gs.S.Addr.String(),
				SigBytes: signature,
			},
		}, nil
	} else if event.EventEndTime.Add(7 * 24 * time.Hour).After(time.Now()) {
		currentReward, err := calcRetentionReward(request.GetAddr(), event, levelConfig)
		if err != nil {
			log.Errorln("failed to calcRetentionReward:", request, " error:", err)
			return &webapi.GetRetentionRewardsInfoResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "db query fail",
				},
			}, nil
		}
		return &webapi.GetRetentionRewardsInfoResponse{
			EventId:           dal.RetentionRewardEventId,
			EventEndTime:      common.TsMilli(event.EventEndTime),
			MaxReward:         levelConfig.MaxReward.String(),
			MaxTransferVolume: levelConfig.MaxTransferVolume,
			CurrentReward:     currentReward.String(),
			CelrUsdPrice:      price,
		}, nil
	}
	return &webapi.GetRetentionRewardsInfoResponse{}, nil
}

func calcRetentionReward(addr string, event *dal.RetentionRewardsEvent, levelConfig *dal.RetentionRewardsLevelConfig) (*big.Int, error) {
	volume, err := dal.DB.GetCompletedVolumeBetween(addr, event.EventEndTime.Add(-30*24*time.Hour), event.EventEndTime)
	if err != nil {
		return nil, err
	}
	wei := big.NewFloat(0).Mul(big.NewFloat(math.Min(volume/levelConfig.MaxTransferVolume, 1)), big.NewFloat(0).SetInt(levelConfig.MaxReward))
	ret := big.NewInt(0)
	wei.Int(ret)
	return ret, nil
}

func (gs *GatewayService) ClaimRetentionRewards(ctx context.Context, request *webapi.ClaimRetentionRewardsRequest) (*webapi.ClaimRetentionRewardsResponse, error) {
	event := dal.RetentionRewardsConfig[dal.RetentionRewardEventId]
	endTime := event.EventEndTime
	if endTime.After(time.Now()) || endTime.Add(7*24*time.Hour).Before(time.Now()) {
		return &webapi.ClaimRetentionRewardsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "not allow to claim",
			},
		}, nil
	}
	level, rewardAmt, claimTime, oldSig, found, err := dal.DB.GetRetentionRewardsRecord(request.GetAddr(), dal.RetentionRewardEventId)
	if err != nil {
		log.Errorln("failed to GetUserRetentionRewardsLevel:", request, " error:", err)
		return &webapi.ClaimRetentionRewardsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "db query fail",
			},
		}, nil
	}
	if !found {
		return &webapi.ClaimRetentionRewardsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "this user is not allow to participate",
			},
		}, nil
	}
	levelConfig := event.LevelConfig[level]
	if claimTime.After(time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC)) {
		return &webapi.ClaimRetentionRewardsResponse{
			EventId:       dal.RetentionRewardEventId,
			CurrentReward: rewardAmt.String(),
			Signature: types3.Signature{
				Signer:   gs.S.Addr.String(),
				SigBytes: oldSig,
			},
		}, nil
	}
	currentReward, err := calcRetentionReward(request.GetAddr(), event, levelConfig)
	if err != nil {
		log.Errorln("failed to calcRetentionReward:", request, " error:", err)
		return &webapi.ClaimRetentionRewardsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "db query fail",
			},
		}, nil
	}
	sig, err := gs.SignIncentiveReward(request.GetAddr(), dal.RetentionRewardEventId, currentReward)
	if err != nil {
		log.Errorln("failed to Sign RetentionRewards:", request, " error:", err)
		return &webapi.ClaimRetentionRewardsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "failed to Sign",
			},
		}, nil
	}
	log.Infoln("sign RetentionRewards, sig:", sig, ", addr:", request.GetAddr(), ", retentionRewardEventId:", dal.RetentionRewardEventId, ", reward:", currentReward)
	err = dal.DB.UpdateRetentionRewardsRecord(request.GetAddr(), dal.RetentionRewardEventId, currentReward, sig)
	if err != nil {
		log.Errorln("failed to claim:", request, " error:", err)
		return &webapi.ClaimRetentionRewardsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "failed to claim",
			},
		}, nil
	}
	return &webapi.ClaimRetentionRewardsResponse{
		EventId:       dal.RetentionRewardEventId,
		CurrentReward: currentReward.String(),
		Signature: types3.Signature{
			Signer:   gs.S.Addr.String(),
			SigBytes: sig,
		},
	}, nil
}

func (gs *GatewayService) SignIncentiveReward(addr string, eventId uint64, rewardAmt *big.Int) ([]byte, error) {
	data := solsha3.Pack(
		// types
		[]string{"uint256", "address", "string", "address", "uint256", "uint256"},
		// values
		[]interface{}{
			fmt.Sprintf("%d", viper.GetInt64(common.FlagGatewayIncentiveRewardsBscChainId)),
			viper.GetString(common.FlagGatewayIncentiveRewardsBscContractAddress),
			"IncentiveRewardClaim",
			addr,
			fmt.Sprintf("%d", eventId),
			rewardAmt.String(),
		},
	)
	sig, err := (*gs.S.Signer).SignEthMessage(data)
	if err != nil {
		return nil, err
	}
	if sig[64] <= 1 {
		// Use 27/28 for v to be compatible with openzeppelin ECDSA lib
		sig[64] = sig[64] + 27
	}
	return sig, nil
}

func (gs *GatewayService) GetFeeRebateInfo(ctx context.Context, request *webapi.GetFeeRebateInfoRequest) (*webapi.GetFeeRebateInfoResponse, error) {
	event := dal.FeeRebateConfig[dal.FeeRebateEventId]
	if time.Now().Before(event.EventStartTime) || time.Now().After(event.EventEndTime.Add(7*24*time.Hour)) {
		return &webapi.GetFeeRebateInfoResponse{}, nil
	}
	price, err := gs.F.GetUsdPrice("CELR")
	if err != nil {
		log.Errorln("failed to Get celr UsdPrice:", err)
	}
	rewardAmt, portion, claimTime, signature, totalFee, found, err := dal.DB.GetFeeRebateRecord(request.GetAddr(), dal.FeeRebateEventId)
	if err != nil {
		log.Errorln("failed to GetFeeRebateRecord:", request, " error:", err)
		return &webapi.GetFeeRebateInfoResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "db query fail",
			},
		}, nil
	}
	if found && claimTime.After(time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC)) {
		return &webapi.GetFeeRebateInfoResponse{
			EventId:       dal.FeeRebateEventId,
			EventEndTime:  common.TsMilli(event.EventEndTime),
			RebatePortion: portion,
			Reward:        rewardAmt.String(),
			CelrUsdPrice:  price,
			ClaimTime:     common.TsMilli(claimTime),
			Signature: types3.Signature{
				Signer:   gs.S.Addr.String(),
				SigBytes: signature,
			},
		}, nil
	}
	portion, celrAmt, err := calcFeeRebatePortionAndReward(request.GetAddr(), price, totalFee, event)
	if err != nil {
		log.Errorln("failed to calcFeeRebatePortionAndReward:", request, " error:", err)
		return &webapi.GetFeeRebateInfoResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "calcFeeRebatePortionAndReward fail",
			},
		}, nil
	}
	return &webapi.GetFeeRebateInfoResponse{
		EventId:       dal.FeeRebateEventId,
		EventEndTime:  common.TsMilli(event.EventEndTime),
		RebatePortion: portion,
		Reward:        celrAmt.String(),
		CelrUsdPrice:  price,
	}, nil
}

func (gs *GatewayService) ClaimFeeRebate(ctx context.Context, request *webapi.ClaimFeeRebateRequest) (*webapi.ClaimFeeRebateResponse, error) {
	event := dal.FeeRebateConfig[dal.FeeRebateEventId]
	if event.EventEndTime.After(time.Now()) || event.EventEndTime.Add(7*24*time.Hour).Before(time.Now()) {
		return &webapi.ClaimFeeRebateResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "not allow to claim",
			},
		}, nil
	}
	rewardAmt, _, claimTime, signature, totalFee, found, err := dal.DB.GetFeeRebateRecord(request.GetAddr(), dal.FeeRebateEventId)
	if err != nil {
		log.Errorln("failed to GetFeeRebateRecord:", request, " error:", err)
		return &webapi.ClaimFeeRebateResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "db query fail",
			},
		}, nil
	}
	if found && claimTime.After(time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC)) {
		return &webapi.ClaimFeeRebateResponse{
			EventId: dal.FeeRebateEventId,
			Reward:  rewardAmt.String(),
			Signature: types3.Signature{
				Signer:   gs.S.Addr.String(),
				SigBytes: signature,
			},
		}, nil
	}
	price, err := gs.F.GetUsdPrice("CELR")
	if err != nil {
		log.Errorln("failed to Get celr UsdPrice:", err)
	}
	portion, celrAmt, err := calcFeeRebatePortionAndReward(request.GetAddr(), price, totalFee, event)
	if err != nil {
		log.Errorln("failed to calcFeeRebatePortionAndReward:", request, " error:", err)
		return &webapi.ClaimFeeRebateResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "calcFeeRebatePortionAndReward fail",
			},
		}, nil
	}
	sig, err := gs.SignIncentiveReward(request.GetAddr(), dal.FeeRebateEventId, celrAmt)
	if err != nil {
		log.Errorln("failed to Sign FeeRebate:", request, " error:", err)
		return &webapi.ClaimFeeRebateResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "failed to Sign",
			},
		}, nil
	}
	log.Infoln("sign FeeRebate, sig:", sig, ", addr:", request.GetAddr(), ", FeeRebateEventId:", dal.FeeRebateEventId, ", reward:", celrAmt)
	err = dal.DB.ClaimFeeRebateRecord(request.GetAddr(), dal.FeeRebateEventId, celrAmt, portion, sig)
	if err != nil {
		log.Errorln("failed to claim:", request, " error:", err)
		return &webapi.ClaimFeeRebateResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "failed to claim",
			},
		}, nil
	}
	return &webapi.ClaimFeeRebateResponse{
		EventId: dal.FeeRebateEventId,
		Reward:  celrAmt.String(),
		Signature: types3.Signature{
			Signer:   gs.S.Addr.String(),
			SigBytes: sig,
		},
	}, nil
}

func calcFeeRebatePortionAndReward(addr string, price, totalFeeInUsd float64, event *dal.FeeRebateEvent) (portion float64, celrAmt *big.Int, err error) {
	volume, err := dal.DB.GetCompletedVolumeBetween(addr, event.EventStartTime, event.EventEndTime)
	if err != nil {
		log.Errorln("failed to GetCompletedVolumeBetween, addr:", addr, " error:", err)
		return 0, nil, fmt.Errorf("db query fail")
	}
	// lower than event.levelDivisionUpperbound[0] get no reward
	if event.LevelDivisionUpperbound[0] > volume {
		return 0, big.NewInt(0), nil
	}
	i := 1
	for {
		if event.LevelDivisionUpperbound[i-1] <= volume && volume < event.LevelDivisionUpperbound[i] {
			break
		}
		i++
		if i >= len(event.LevelDivisionUpperbound) {
			break
		}
	}
	portion = event.LevelConfig[uint64(i)].RebatePortion
	quo := big.NewFloat(0).Quo(big.NewFloat(totalFeeInUsd), big.NewFloat(price))
	quo.Mul(quo, big.NewFloat(portion))
	quo.Mul(quo, big.NewFloat(0).SetInt(dal.WeiMultiplier))
	wei := big.NewInt(0)
	quo.Int(wei)
	if wei.Cmp(event.LevelConfig[uint64(i)].MaxReward) > 0 {
		wei = event.LevelConfig[uint64(i)].MaxReward
	}
	return portion, wei, nil
}
