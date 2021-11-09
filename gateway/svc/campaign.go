package gatewaysvc

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (gs *GatewayService) GetCampaignScores(ctx context.Context, request *webapi.GetCampaignScoresRequest) (*webapi.GetCampaignScoresResponse, error) {
	log.Infof("get campaign score data, data:%d, begin:%d, end:%d", request.GetDate(), request.GetBeginBlock(), request.GetEndBlock())
	stakingEthClient, err := common.NewEthClientFromConfig()
	if stakingEthClient == nil || err != nil {
		log.Errorf("get eth client failed, err:%+v", err)
		return nil, nil
	}
	beginTime := time.Date(2021, time.November, int(request.GetDate()), 0, 0, 0, 0, time.UTC)

	//time.Now().Truncate(time.day)
	cbridgeScore, err := dal.DB.CalcCampaignScore(beginTime)
	if err != nil {
		log.Warnf("cal campaign err:%+v", err)
		return nil, nil
	}

	scoreMap := score2Map(cbridgeScore)
	log.Infof("score Map:%+v", scoreMap)

	stakingMap, err := getStakingCampaignData(stakingEthClient, request.GetBeginBlock(), request.GetEndBlock())
	if err != nil {
		return &webapi.GetCampaignScoresResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get staking data failed: " + err.Error(),
			},
		}, nil
	}
	log.Infof("staking Map:%+v", stakingMap)
	rewardMap, err := getStakingRewardCampaignData(stakingEthClient, request.GetBeginBlock(), request.GetEndBlock())
	if err != nil {
		return &webapi.GetCampaignScoresResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get reward data failed: " + err.Error(),
			},
		}, nil
	}
	log.Infof("reward Map:%+v", rewardMap)

	mergeMap(scoreMap, stakingMap)
	mergeMap(scoreMap, rewardMap)
	return &webapi.GetCampaignScoresResponse{
		Scores: map2Score(scoreMap),
		Begin:  request.BeginBlock,
	}, nil
}

func mergeMap(a, b map[string]uint64) {
	for addr, score := range b {
		a[addr] = a[addr] + score
	}
}

func getStakingCampaignData(ec *eth.EthClient, start, end uint64) (map[string]uint64, error) {
	resp := make(map[string]uint64)
	if start == 0 && end == 0 {
		return resp, nil
	}
	if ec == nil {
		return resp, fmt.Errorf("eth.EthClient is nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	e := new(uint64)
	*e = end
	iterator, err := ec.Contracts.Staking.FilterDelegationUpdate(&bind.FilterOpts{
		Start:   start,
		End:     e,
		Context: ctx,
	}, nil, nil)
	if err != nil {
		return nil, err
	}
	for iterator.Next() {
		if utils.IsBot(iterator.Event.DelAddr.String()) {
			continue
		}
		if resp[iterator.Event.DelAddr.String()] < 100 {
			resp[iterator.Event.DelAddr.String()]++
		}
	}
	return resp, nil
}

func getStakingRewardCampaignData(ec *eth.EthClient, start, end uint64) (map[string]uint64, error) {
	resp := make(map[string]uint64, 0)
	if start == 0 && end == 0 {
		return resp, nil
	}
	if ec == nil {
		return resp, fmt.Errorf("eth.EthClient is nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	e := new(uint64)
	*e = end
	iterator, err := ec.Contracts.StakingReward.FilterStakingRewardClaimed(&bind.FilterOpts{
		Start:   start,
		End:     e,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}
	for iterator.Next() {
		if utils.IsBot(iterator.Event.Recipient.String()) {
			continue
		}
		if resp[iterator.Event.Recipient.String()] < 50 {
			resp[iterator.Event.Recipient.String()]++
		}
	}
	return resp, nil
}

func score2Map(s []*webapi.CampaignScore) map[string]uint64 {
	m := make(map[string]uint64)
	for _, score := range s {
		m[score.GetUsrAddr()] += score.GetScore()
	}
	return m
}

func map2Score(m map[string]uint64) []*webapi.CampaignScore {
	var s []*webapi.CampaignScore
	for addr, score := range m {
		s = append(s, &webapi.CampaignScore{
			UsrAddr: addr,
			Score:   score,
		})
	}
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Score > s[j].Score
	})
	return s
}
