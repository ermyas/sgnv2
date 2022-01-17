package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/version"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

func (c *CbrOneChain) setCurss(ss []*cbrtypes.Signer) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs, c.curss.powers = cbrtypes.SignersToEthArrays(ss)
}

func (c *CbrOneChain) setCurssByEvent(e *eth.BridgeSignersUpdated) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs = make([]eth.Addr, len(e.Signers))
	c.curss.powers = make([]*big.Int, len(e.Powers))
	for i, addr := range e.Signers {
		c.curss.addrs[i] = addr
		c.curss.powers[i] = e.Powers[i]
	}
}

func (c *CbrOneChain) getCurss() currentSigners {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.curss
}

// each event's key is name-blkNum-index, value is json marshaled elog
func (c *CbrOneChain) saveEvent(name string, elog ethtypes.Log) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key := fmt.Sprintf("%s-%d-%d", name, elog.BlockNumber, elog.Index)
	val, _ := json.Marshal(elog)
	return c.db.Set([]byte(key), val)
}

func (c *CbrOneChain) delEvent(name string, blknum, idx uint64) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.db.Delete([]byte(fmt.Sprintf("%s-%d-%d", name, blknum, idx)))
}

func (r *Relayer) startReportCurrentBlockNumber(interval time.Duration) {
	endpoint := viper.GetString(common.FlagSgnLivenessReportEndpoint)
	if endpoint == "" {
		log.Info("report current block number disabled")
		return
	}
	go func() {
		// let gateway start upfront
		time.Sleep(15 * time.Second)
		log.Infoln("report current block number to", endpoint)
		ticker := jitterbug.New(
			interval,
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
		Timestamp:   common.TsMilli(time.Now()),
		BlockNums:   make(map[string]uint64),
		SgndVersion: version.Version,
	}
	for chainId, oneChain := range r.cbrMgr {
		blockNumber := oneChain.mon.GetCurrentBlockNumber()
		report.BlockNums[strconv.Itoa(int(chainId))] = blockNumber.Uint64()
	}
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
