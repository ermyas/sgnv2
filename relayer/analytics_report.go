package relayer

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"gopkg.in/resty.v1"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
)

var lpFeeEarningHistoryMap = make(map[uint64]*LPFeeEarningHistory)
var lpFeeEarningHistoryLock sync.RWMutex

var baseFeeDistributionHistoryMap = make(map[uint64]*BaseFeeDistributionHistory)
var baseFeeDistributionHistoryLock sync.RWMutex

func (r *Relayer) startValidatorNodeAnalyticsReport() {
	endpoint := viper.GetString(common.FlagSgnLivenessReportEndpoint)
	if endpoint == "" {
		log.Info("sgn validator node analytics report disabled")
		return
	}
	go func() {
		time.Sleep(15 * time.Second)
		log.Infoln("start sgn validator node analytics report,", viper.GetString(common.FlagSgnLivenessReportEndpoint))
		ticker := jitterbug.New(
			time.Minute*5,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			r.reportValidatorNodeAnalytics()
		}
	}()
}

func startConsensusLogReport() {
	endpoint := viper.GetString(common.FlagSgnConsensusLogReportEndpoint)
	if endpoint == "" {
		log.Info("sgn consensus log report disabled")
		return
	}
	go func() {
		time.Sleep(30 * time.Second)
		log.Infoln("start sgn consensus log report,", viper.GetString(common.FlagSgnConsensusLogReportEndpoint))
		ticker := jitterbug.New(
			time.Minute*5,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			reportConsensusLog()
		}
	}()
}

func reportConsensusLog() {
	report := &SgnConsensusLogReport{
		LpFeeEarningHistories:        getAndClearLpEarningFeeHistory(),
		BaseFeeDistributionHistories: getAndClearBaseFeeDistributionHistory(),
	}
	log.Debugln("try to report:", report)
	url := viper.GetString(common.FlagSgnConsensusLogReportEndpoint)
	if len(url) == 0 {
		return
	}
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(report)
	if err != nil {
		log.Warnln("failed to MarshalToString: err ", err)
		return
	}
	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(str).
		SetResult(&ReportSgnAnalyticsResponse{}).
		Post(url)
	if err != nil || response.StatusCode() != 200 {
		chainConfigReported = false
		log.Warnln("fail to reportConsensusLog ", report, err, response)
		return
	}
	resp := response.Result().(*ReportSgnAnalyticsResponse)
	if resp.GetErr() != nil {
		chainConfigReported = false
		log.Warnln("fail to reportConsensusLog ", report, err, response)
		return
	}
}

func (r *Relayer) reportValidatorNodeAnalytics() {
	var report = &SgnAnalyticsReport{
		Timestamp:    common.TsMilli(time.Now()),
		BlockNums:    make(map[string]uint64),
		SgndVersion:  version.Version,
		ChainConfigs: r.getChainConfig(),
		BlockTimes:   make(map[string]uint64),
	}
	for chainId, oneChain := range r.cbrMgr {
		var blkNum, blkTm uint64
		if types.IsFlowChain(chainId) {
			blkNum = oneChain.fcc.GetBlkNum()

			blkInfo, err := oneChain.fcc.QueryLatestBlock(context.Background(), true)
			if nil != err {
				log.Warnf("chain %d QueryLatestBlock err: %v", chainId, err)
			} else {
				blkTm = uint64(blkInfo.Timestamp.Unix())
			}
		} else {
			blkNum = oneChain.mon.GetBlkNum()

			head, err := oneChain.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Warnf("chain %d HeaderByNumber err: %v", chainId, err)
			} else {
				blkTm = head.Time
			}
		}
		report.BlockNums[strconv.Itoa(int(chainId))] = blkNum
		report.BlockTimes[strconv.Itoa(int(chainId))] = blkTm
	}
	log.Debugln("try to report:", report)
	bytes, err := proto.Marshal(report)
	if err != nil {
		chainConfigReported = false
		log.Warnln("fail to Marshal CurrentBlockNumberReport,", err)
		return
	}
	sig, err := r.EthClient.SignEthMessage(bytes)
	if err != nil {
		chainConfigReported = false
		log.Warnln("fail to Sign CurrentBlockNumberReport,", err)
		return
	}
	req := &ReportSgnAnalyticsRequest{
		Report: bytes,
		Sig:    sig,
	}
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(req)
	if err != nil {
		chainConfigReported = false
		log.Warnln("failed to MarshalToString: err ", err)
		return
	}
	url := viper.GetString(common.FlagSgnLivenessReportEndpoint)
	if len(url) == 0 {
		return
	}
	response, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(str).
		SetResult(&ReportSgnAnalyticsResponse{}).
		Post(url)
	if err != nil || response.StatusCode() != 200 {
		chainConfigReported = false
		log.Warnln("fail to reportValidatorNodeAnalytics ", req, err, response)
		return
	}
	resp := response.Result().(*ReportSgnAnalyticsResponse)
	if resp.GetErr() != nil {
		chainConfigReported = false
		log.Warnln("fail to reportValidatorNodeAnalytics ", req, err, response)
		return
	}
}

func AppendLpPickHistoryLog(lpAddr, tokenAddr eth.Addr, lpAmt *big.Int, dstChainId uint64, used, earnedFee *big.Int, start time.Time) {
	if len(viper.GetString(common.FlagSgnConsensusLogReportEndpoint)) > 0 {
		lpFeeEarningHistoryLock.Lock()
		defer lpFeeEarningHistoryLock.Unlock()
		t := uint64(start.UnixNano())
		if lpFeeEarningHistoryMap[t] == nil {
			lpFeeEarningHistoryMap[t] = &LPFeeEarningHistory{
				DstTokenAddr: tokenAddr.Hex(),
				DstChainId:   dstChainId,
				Logs: map[string]*LPFeeEarningLog{
					lpAddr.Hex(): {
						DstChainLiquidityUsed:     used.String(),
						EarnedFee:                 earnedFee.String(),
						DstChainLiquidityRemained: lpAmt.String(),
					},
				},
			}
		} else {
			lpFeeEarningHistoryMap[t].Logs[lpAddr.Hex()] = &LPFeeEarningLog{
				DstChainLiquidityUsed:     used.String(),
				EarnedFee:                 earnedFee.String(),
				DstChainLiquidityRemained: lpAmt.String(),
			}
		}
	}
}

func ReportBaseFeeDistribution(bridgeType BridgeType, syncerAddr eth.Addr, start time.Time, baseFee *big.Int, tokenSymbol string, tokenDecimal uint32, srcChainId, dstChainId uint64) {
	if len(viper.GetString(common.FlagSgnConsensusLogReportEndpoint)) > 0 {
		baseFeeDistributionHistoryLock.Lock()
		defer baseFeeDistributionHistoryLock.Unlock()
		t := uint64(start.UnixNano())
		baseFeeDistributionHistoryMap[t] = &BaseFeeDistributionHistory{
			BridgeType:          bridgeType,
			BaseFeeReceiverAddr: syncerAddr.Hex(),
			BaseFeeAmt:          baseFee.String(),
			TokenSymbol:         tokenSymbol,
			TokenDecimal:        tokenDecimal,
			SrcChainId:          srcChainId,
			DstChainId:          dstChainId,
		}
	}
}

func getAndClearLpEarningFeeHistory() map[uint64]*LPFeeEarningHistory {
	// only witness node report
	if len(viper.GetString(common.FlagSgnConsensusLogReportEndpoint)) == 0 {
		return make(map[uint64]*LPFeeEarningHistory)
	}
	lpFeeEarningHistoryLock.Lock()
	defer lpFeeEarningHistoryLock.Unlock()
	t := make(map[uint64]*LPFeeEarningHistory)
	for nanoTs, history := range lpFeeEarningHistoryMap {
		t[nanoTs] = history
	}
	lpFeeEarningHistoryMap = make(map[uint64]*LPFeeEarningHistory)
	return t
}

func getAndClearBaseFeeDistributionHistory() map[uint64]*BaseFeeDistributionHistory {
	// only witness node report
	if len(viper.GetString(common.FlagSgnConsensusLogReportEndpoint)) == 0 {
		return make(map[uint64]*BaseFeeDistributionHistory)
	}
	baseFeeDistributionHistoryLock.Lock()
	defer baseFeeDistributionHistoryLock.Unlock()
	t := make(map[uint64]*BaseFeeDistributionHistory)
	for nanoTs, history := range baseFeeDistributionHistoryMap {
		t[nanoTs] = history
	}
	baseFeeDistributionHistoryMap = make(map[uint64]*BaseFeeDistributionHistory)
	return t
}

var chainConfigReported = false

func (r *Relayer) getChainConfig() map[string]*ChainConfig {
	if chainConfigReported {
		return nil
	}
	m := make(map[string]*ChainConfig)
	for chainId, oneChain := range r.cbrMgr {
		if types.IsFlowChain(chainId) {
			m[fmt.Sprintf("%d", chainId)] = &ChainConfig{
				CbridgeContractAddr:            oneChain.ContractAddr, // all flow contracts under same account
				OriginalTokenVaultContractAddr: oneChain.ContractAddr,
				PeggedTokenBridgeContractAddr:  oneChain.ContractAddr,
				// no msgbus on flow yet
			}
			continue
		}
		m[fmt.Sprintf("%d", chainId)] = &ChainConfig{
			CbridgeContractAddr:            oneChain.cbrContract.GetAddr().Hex(),
			OriginalTokenVaultContractAddr: oneChain.pegContracts.vault.GetAddr().Hex(),
			PeggedTokenBridgeContractAddr:  oneChain.pegContracts.bridge.GetAddr().Hex(),
			MsgBusContractAddr:             oneChain.msgContract.GetAddr().Hex(),
		}
	}
	chainConfigReported = true
	return m
}
