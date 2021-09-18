package relayer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	slashcli "github.com/celer-network/sgn-v2/x/slash/client/cli"
	slashtypes "github.com/celer-network/sgn-v2/x/slash/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	tm "github.com/tendermint/tendermint/types"
)

var (
	EventSlash = fmt.Sprintf("%s.%s='%s'", slashingtypes.EventTypeSlash, sdk.AttributeKeyAction, slashtypes.ActionSlash)
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

func (r *Relayer) monitorSgnchainSlash() {
	MonitorTendermintEvent(r.Transactor.CliCtx.NodeURI, EventSlash, func(e abci.Event) {
		if !r.isBonded() {
			return
		}

		event := sdk.StringifyEvent(e)

		if event.Attributes[0].Value == slashtypes.ActionSlash {
			nonce, err := strconv.ParseUint(event.Attributes[1].Value, 10, 64)
			if err != nil {
				log.Errorln("Parse slash nonce error", err)
				return
			}

			slashEvent := NewSlashEvent(nonce)
			r.handleSlash(slashEvent)
			err = r.dbSet(GetSlashKey(slashEvent.Nonce), slashEvent.MustMarshal())
			if err != nil {
				log.Errorln("db Set err", err)
			}
		}
	})
}

func (r *Relayer) handleSlash(slashEvent SlashEvent) {
	slash, err := slashcli.QuerySlash(r.Transactor.CliCtx, slashtypes.StoreKey, slashEvent.Nonce)
	if err != nil {
		log.Errorf("Query slash %d err %s", slashEvent.Nonce, err)
		return
	}
	log.Infof("New slash to %s, reason %s, nonce %d", slash.Validator, slash.Reason, slashEvent.Nonce)

	sig, err := r.EthClient.SignEthMessage(slash.EthSlashBytes)
	if err != nil {
		log.Errorln("SignEthMessage err", err)
		return
	}

	msg := slashtypes.NewMsgSignSlash(slashEvent.Nonce, sig, r.Transactor.Key.GetAddress())
	r.Transactor.AddTxMsg(&msg)
}

// zhihua
func (r *Relayer) monitorCbrToSign() {
	// cbr event to sign msg
	// see cbridge/types/types.go for event attributes
	// msg to send is MsgSendMySig
	MonitorTendermintEvent(r.Transactor.CliCtx.NodeURI, cbrtypes.EventToSign, func(e abci.Event) {})
}
