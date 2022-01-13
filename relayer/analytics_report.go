package relayer

import (
	"gopkg.in/resty.v1"
	"strconv"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
)

var LiquidityProviderFeeEarningLogList = make([]*LiquidityProviderFeeEarningLog, 0)

func (r *Relayer) startReportCurrentBlockNumber() {
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
			r.reportCurrentBlockNumber()
		}
	}()
}

func (r *Relayer) reportCurrentBlockNumber() {
	var report = &CurrentBlockNumberReport{
		Timestamp:        common.TsMilli(time.Now()),
		BlockNums:        make(map[string]uint64),
		SgndVersion:      version.Version,
		LpFeeEarningLogs: getAndClearLpEarningFeeLogList(),
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
	req := &ReportCurrentBlockNumberRequest{
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
		SetResult(&ReportCurrentBlockNumberResponse{}).
		Post(url)
	if err != nil || response.StatusCode() != 200 {
		log.Warnln("fail to reportCurrentBlockNumber ", req, err, response)
		return
	}
	resp := response.Result().(*ReportCurrentBlockNumberResponse)
	if resp.GetErr() != nil {
		log.Warnln("fail to reportCurrentBlockNumber ", req, err, response)
		return
	}
}

func getAndClearLpEarningFeeLogList() []*LiquidityProviderFeeEarningLog {
	// only node 0 report
	if !viper.GetBool(common.FlagSgnReportLpFeeEarningFlag) {
		return make([]*LiquidityProviderFeeEarningLog, 0)
	}
	t := make([]*LiquidityProviderFeeEarningLog, len(LiquidityProviderFeeEarningLogList))
	copy(t, LiquidityProviderFeeEarningLogList)
	LiquidityProviderFeeEarningLogList = make([]*LiquidityProviderFeeEarningLog, 0)
	return t
}
