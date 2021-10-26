package relayer

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
)

const (
	Interval = time.Duration(5) * time.Minute
)

// sleep, check if syncer, if yes, check if update_epoch is newer
func (r *Relayer) pullPriceChange() {
	log.Infoln("start pull cbr price change, interval:", Interval)
	for {
		time.Sleep(Interval)
		if !r.isSyncer() {
			continue
		}
		msg := &synctypes.MsgProposeUpdates{
			Sender: r.Transactor.Key.GetAddress().String(),
		}
		cp, success := getCbrPriceFromUrl()
		log.Debugln("get valid cbr price.", cp, success)
		if !success {
			continue
		}
		data, _ := cp.Marshal()
		msg.Updates = append(msg.Updates, &synctypes.ProposeUpdate{
			Type: synctypes.DataType_CbrUpdateCbrPrice,
			Data: data,
		})
		r.Transactor.AddTxMsg(msg)
	}
}

func getCbrPriceFromUrl() (cp *types.CbrPrice, success bool) {
	cp = &types.CbrPrice{}
	url := viper.GetString(common.FlagSgnPriceUpdateUrl)
	client := resty.New()
	response, err := client.R().Get(url)
	if err != nil || response.StatusCode() != 200 {
		log.Errorln("fail to get price change json, ", url, " error:", err)
		return nil, false
	}
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	err = unmarshaler.Unmarshal(strings.NewReader(response.String()), cp)
	if err != nil {
		log.Errorln("fail to get price change json, ", url, " error:", err)
		return nil, false
	}
	// rough check
	if common.TsToTime(cp.GetUpdateEpoch()).Add(Interval).Before(time.Now()) {
		log.Errorln("seems like oracle stopped working, latest cbrPrice update time ", common.TsToTime(cp.GetUpdateEpoch()))
		return nil, false
	}
	return cp, true
}
