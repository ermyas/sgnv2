package gatewaysvc

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	"github.com/lthibault/jitterbug"
	"time"
)

func (gs *GatewayService) StartFailedGasOnArrivalMonitor() {
	go func() {
		time.Sleep(time.Minute)
		ticker := jitterbug.New(
			time.Hour,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			gs.doStartMonitor()
		}
	}()
}

func (gs *GatewayService) doStartMonitor() {
	logs, err := dal.DB.FindFailedGasOnArrivalLog(time.Now().Add(-24 * time.Hour))
	if err != nil {
		log.Errorln("failed to FindFailedGasOnArrivalLog,", err)
		return
	}
	if len(logs) == 0 {
		log.Infoln("no failed gas on arrival to be sent.")
		return
	}
	log.Infoln(len(logs), " addresses waiting to be sent gas on arrival")
	for _, arrivalLog := range logs {
		transfer, found, err := dal.DB.GetTransfer(arrivalLog.TransferId)
		if err != nil || !found {
			log.Errorln("can't find transfer info at gateway, ", arrivalLog.TransferId, err)
			continue
		}
		ethClient := gs.Chains.GetEthClient(arrivalLog.ChainId)
		err = onchain.SendGasOnArrival(ethClient, transfer)
		if err != nil {
			log.Errorln("can't SendGasOnArrival, ", arrivalLog.TransferId, err)
			continue
		}
		err = dal.DB.UpdateGasOnArrivalLogToSuccess(arrivalLog.TransferId)
		if err != nil {
			log.Errorln("can't UpdateGasOnArrivalLogToSuccess, ", arrivalLog.TransferId, err)
			continue
		}
		log.Infoln("auto retry send gas on arrival to ", arrivalLog.UsrAddr, ", transferId:", arrivalLog.TransferId)
	}
}
