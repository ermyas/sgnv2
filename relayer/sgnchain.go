package relayer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	slashcli "github.com/celer-network/sgn-v2/x/slash/client/cli"
	slashtypes "github.com/celer-network/sgn-v2/x/slash/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	tm "github.com/tendermint/tendermint/types"
)

var (
	EventSlash   = fmt.Sprintf("%s.%s='%s'", slashingtypes.EventTypeSlash, sdk.AttributeKeyAction, slashtypes.ActionSlash)
	EventCbridge = fmt.Sprintf("%s.%s='%s'", cbrtypes.EventToSign, sdk.AttributeKeyAction, cbrtypes.EventToSign)
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

func (r *Relayer) monitorCbrToSign() {
	MonitorTendermintEvent(r.Transactor.CliCtx.NodeURI, EventCbridge, func(e abci.Event) {
		if e.Type != cbrtypes.EventToSign {
			return
		}
		log.Infoln("monitorCbrToSign, eventType: ", e.Type)
		if !r.isBonded() {
			return
		}
		event := sdk.StringifyEvent(e)
		// sign data first
		data := []byte(event.Attributes[1].Value)
		sig, err := r.EthClient.SignEthMessage(data)
		if err != nil {
			log.Error(err)
			return
		}
		msg := &cbrtypes.MsgSendMySig{
			Data:    data,
			MySig:   sig,
			Creator: r.Transactor.Key.GetAddress().String(),
		}
		logmsg := fmt.Sprintf("Sign cBridge data %s", event.Attributes[0])
		switch event.Attributes[0].Value {
		case cbrtypes.SignDataType_RELAY.String():
			msg.Datatype = cbrtypes.SignDataType_RELAY
			relay := new(cbrtypes.RelayOnChain)
			err = relay.Unmarshal(data)
			if err != nil {
				log.Errorf("%s, failed to unmarshal XrefRelay: %s", logmsg, err)
				return
			}
			relayEvent := NewRelayEvent(relay.SrcTransferId)
			err = r.dbSet(GetCbrXferKey(relayEvent.XferId), relayEvent.MustMarshal())
			if err != nil {
				log.Errorf("%s, db Set err: %s", logmsg, err)
			}
			log.Infof("%s: %s", logmsg, relay.String())
		case cbrtypes.SignDataType_WITHDRAW.String():
			msg.Datatype = cbrtypes.SignDataType_WITHDRAW
			log.Infof("%s", logmsg)
		case cbrtypes.SignDataType_SIGNERS.String():
			msg.Datatype = cbrtypes.SignDataType_SIGNERS
			//r.setCbrSsUpdating()
			ss := new(cbrtypes.SortedSigners)
			err = ss.Unmarshal(data)
			if err != nil {
				log.Errorf("%s, failed to unmarshal sorted signers: %s", logmsg, err)
				return
			}
			log.Infof("%s: %s", logmsg, ss.String())
		}
		r.Transactor.AddTxMsg(msg)
	})
}

func (r *Relayer) monitorSgnFarmingClaimAllEvent() {
	MonitorTendermintEvent(r.Transactor.CliCtx.NodeURI, farmingtypes.EventTypeClaimAll, func(e abci.Event) {
		if !r.isBonded() {
			return
		}
		event := sdk.StringifyEvent(e)
		r.handleFarmingClaimEvent(event.Attributes[1].Value)
	})
}

func (r *Relayer) handleFarmingClaimEvent(addr string) {
	queryClient := farmingtypes.NewQueryClient(r.Transactor.CliCtx)
	rewardClaimInfo, err := queryClient.RewardClaimInfo(
		context.Background(),
		&farmingtypes.QueryRewardClaimInfoRequest{
			Address: addr,
		},
	)
	if err != nil {
		log.Errorf("Query RewardClaimInfo err %s", err)
		return
	}
	var signatureDetailsList []farmingtypes.SignatureDetails
	for _, details := range rewardClaimInfo.RewardClaimInfo.RewardClaimDetailsList {
		sig, err := r.EthClient.SignEthMessage(details.RewardProtoBytes)
		if err != nil {
			log.Errorln("SignEthMessage err", err)
			return
		}
		signatureDetailsList = append(signatureDetailsList, farmingtypes.SignatureDetails{
			ChainId:   details.ChainId,
			Signature: sig,
		})
	}
	msg := farmingtypes.NewMsgSignRewards(eth.Hex2Addr(addr), r.Transactor.Key.GetAddress(), signatureDetailsList)
	r.Transactor.AddTxMsg(msg)
}
