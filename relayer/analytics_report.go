package relayer

import (
	"gopkg.in/resty.v1"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
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

func (r *Relayer) startReportSgnAnalytics() {
	endpoint := viper.GetString(common.FlagSgnLivenessReportEndpoint)
	if endpoint == "" {
		log.Info("report current block number disabled")
		return
	}
	go func() {
		// let gateway start upfront
		time.Sleep(15 * time.Second)
		log.Infoln("start Report Current Block Number,", viper.GetString(common.FlagSgnLivenessReportEndpoint))
		ticker := jitterbug.New(
			time.Minute*5,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			r.reportSgnAnalytics()
		}
	}()
}

func (r *Relayer) reportSgnAnalytics() {
	var report = &SgnAnalyticsReport{
		Timestamp:                    common.TsMilli(time.Now()),
		BlockNums:                    make(map[string]uint64),
		SgndVersion:                  version.Version,
		LpFeeEarningHistories:        getAndClearLpEarningFeeHistory(),
		BaseFeeDistributionHistories: getAndClearBaseFeeDistributionHistory(),
	}
	for chainId, oneChain := range r.cbrMgr {
		blockNumber := oneChain.mon.GetCurrentBlockNumber()
		report.BlockNums[strconv.Itoa(int(chainId))] = blockNumber.Uint64()
	}
	log.Debugln("try to report:", report)
	bytes, err := proto.Marshal(report)
	if err != nil {
		log.Warnln("fail to Marshal CurrentBlockNumberReport,", err)
		return
	}
	sig, err := r.EthClient.SignEthMessage(bytes)
	if err != nil {
		log.Warnln("fail to Sign CurrentBlockNumberReport,", err)
		return
	}
	req := &ReportSgnAnalyticsRequest{
		Report: bytes,
		Sig:    sig,
	}
	client := resty.New()
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(req)
	if err != nil {
		log.Warnln("failed to MarshalToString: err ", err)
		return
	}
	url := viper.GetString(common.FlagSgnLivenessReportEndpoint)
	if len(url) == 0 {
		return
	}
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(str).
		SetResult(&ReportSgnAnalyticsResponse{}).
		Post(url)
	if err != nil || response.StatusCode() != 200 {
		log.Warnln("fail to reportSgnAnalytics ", req, err, response)
		return
	}
	resp := response.Result().(*ReportSgnAnalyticsResponse)
	if resp.GetErr() != nil {
		log.Warnln("fail to reportSgnAnalytics ", req, err, response)
		return
	}
}

func AppendLpPickHistoryLog(lpAddr, tokenAddr eth.Addr, lpAmt *big.Int, dstChainId uint64, used, earnedFee *big.Int, start time.Time) {
	if viper.GetBool(common.FlagSgnReportLpFeeEarningFlag) {
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
	if viper.GetBool(common.FlagSgnReportLpFeeEarningFlag) {
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
	// only node 0 report
	if !viper.GetBool(common.FlagSgnReportLpFeeEarningFlag) {
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
	// only node 0 report
	if !viper.GetBool(common.FlagSgnReportLpFeeEarningFlag) {
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
