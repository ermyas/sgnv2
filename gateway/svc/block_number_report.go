package gatewaysvc

import (
	"context"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/gogo/protobuf/proto"
	"github.com/lthibault/jitterbug"
	"time"
)

func (gs *GatewayService) ReportCurrentBlockNumber(ctx context.Context, request *webapi.ReportCurrentBlockNumberRequest) (*webapi.ReportCurrentBlockNumberResponse, error) {
	report := &webapi.CurrentBlockNumberReport{}
	err := proto.Unmarshal(request.GetReport(), report)
	if err != nil {
		log.Warnln("failed to Unmarshal report proto, ", request, err)
		return &webapi.ReportCurrentBlockNumberResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "failed to Unmarshal",
			},
		}, nil
	}
	signer, err := ethutils.RecoverSigner(request.GetReport(), request.GetSig())
	if err != nil {
		log.Warnln("failed to RecoverSigner, ", request, err)
		return &webapi.ReportCurrentBlockNumberResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "failed to RecoverSigner",
			},
		}, nil
	}
	if gs.V.C[signer] != nil && report.GetTimestamp() < gs.V.C[signer].GetTimestamp() {
		log.Warnln(signer, " report outdated, ", request)
		return &webapi.ReportCurrentBlockNumberResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "report outdated",
			},
		}, nil
	}

	if gs.V.C[signer] == nil && len(gs.V.C) > 1000 {
		log.Errorln("report more than 1000 distinct addrs. will give up receiving more. ", request)
		return &webapi.ReportCurrentBlockNumberResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "report outdated",
			},
		}, nil
	}
	if gs.V.C[signer] != nil {
		// 5 min interval for validator to report, so if blk num at any chain will be regarded as problematic
		flag := false
		for chainId, blockNum := range report.GetBlockNums() {
			if gs.V.C[signer].GetBlockNums()[chainId] >= blockNum {
				gs.V.P[signer] = true
				flag = true
				break
			}
		}
		if flag {
			delete(gs.V.P, signer)
		}
	}

	gs.V.C[signer] = report
	log.Infoln(signer, " report current block number. now:", gs.V.C[signer])
	return &webapi.ReportCurrentBlockNumberResponse{}, nil
}

func (gs *GatewayService) GetCurrentBlockNumberByNode(ctx context.Context, request *webapi.GetCurrentBlockNumberByNodeRequest) (*webapi.GetCurrentBlockNumberByNodeResponse, error) {
	if len(gs.V.C) == 0 {
		return &webapi.GetCurrentBlockNumberByNodeResponse{}, nil
	}
	m := make(map[string]*webapi.CurrentBlockNumberReport)
	p := make([]string, 0)
	for addr, report := range gs.V.C {
		m[addr.String()] = report
		if common.TsMilliToTime(report.GetTimestamp()).Add(30 * time.Minute).Before(time.Now()) {
			p = append(p, addr.String())
		} else if gs.V.P[addr] {
			p = append(p, addr.String())
		}
	}
	return &webapi.GetCurrentBlockNumberByNodeResponse{
		Reports:          m,
		ProblematicAddrs: p,
	}, nil
}

func (gs *GatewayService) StartProblematicCurrentBlockNumberAddrMonitor() {
	go func() {
		time.Sleep(time.Minute)
		ticker := jitterbug.New(
			time.Minute*30,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			gs.doStartProblematicCurrentBlockNumberAddrMonitor()
		}
	}()
}

func (gs *GatewayService) doStartProblematicCurrentBlockNumberAddrMonitor() {
	p := make([]string, 0)
	for addr, report := range gs.V.C {
		if common.TsMilliToTime(report.GetTimestamp()).Add(30 * time.Minute).Before(time.Now()) {
			p = append(p, addr.String())
		} else if gs.V.P[addr] {
			p = append(p, addr.String())
		}
	}
	if len(p) > 0 {
		utils.SendProblematicBlockNumberAlert(p)
	}
}
