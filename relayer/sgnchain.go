package relayer

import (
	"context"

	"github.com/celer-network/goutils/log"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	tm "github.com/tendermint/tendermint/types"
)

func MonitorTendermintEvent(nodeURI, eventTag string, handleEvent func(event abci.Event)) {
	client, err := http.New(nodeURI, "/websocket")
	if err != nil {
		log.Errorln("Fail to start create http client", err)
		return
	}

	err = client.Start()
	if err != nil {
		log.Errorln("Fail to start ws client", err)
		return
	}
	defer client.Stop()

	txs, err := client.Subscribe(context.Background(), "monitor", eventTag)
	if err != nil {
		log.Errorln("ws client subscribe error", err)
		return
	}

	for e := range txs {
		switch data := e.Data.(type) {
		case tm.EventDataNewBlock:
			for _, event := range data.ResultBeginBlock.Events {
				handleEvent(event)
			}
			for _, event := range data.ResultEndBlock.Events {
				handleEvent(event)
			}
		case tm.EventDataTx:
			for _, event := range data.TxResult.Result.Events {
				handleEvent(event)
			}
		}
	}
}

func (r *Relayer) monitorSgnchainCreateValidator() {
}
