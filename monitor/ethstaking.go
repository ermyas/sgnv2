package monitor

import (
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (m *Monitor) monitorEthValidatorParamsUpdate() {
}

func (m *Monitor) monitorEthValidatorStatusUpdate() {
}

func (m *Monitor) monitorEthDelegationUpdate() {
}

func (m *Monitor) monitorEthSgnAddrUpdate() {
	_, err := m.ethMonitor.Monitor(
		&monitor.Config{
			EventName:     string(eth.EventSgnAddrUpdate),
			Contract:      m.EthClient.Contracts.Sgn,
			StartBlock:    m.startEthBlock,
			Reset:         true,
			CheckInterval: getEventCheckInterval(eth.EventSgnAddrUpdate),
		},
		func(cb monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
			log.Infof("Catch event SgnAddrUpdate, tx hash: %x", eLog.TxHash)
			event := eth.NewEvent(eth.EventSgnAddrUpdate, eLog)
			dberr := m.dbSet(GetPullerKey(eLog), event.MustMarshal())
			if dberr != nil {
				log.Errorln("db Set err", dberr)
			}
			if !m.isBonded() {
			}

			return false
		},
	)
	if err != nil {
		log.Fatal(err)
	}

}
