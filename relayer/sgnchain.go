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
	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	"github.com/tendermint/tendermint/rpc/client/http"
	tm "github.com/tendermint/tendermint/types"
)

var (
	EventQuerySlash = tmquery.MustParse(
		fmt.Sprintf("tm.event='NewBlock' AND %s.%s EXISTS", slashtypes.EventTypeSlash, slashtypes.AttributeKeyNonce)).String()
	EventQueryCbridge = tmquery.MustParse(
		fmt.Sprintf("%s.%s='%s'", cbrtypes.EventTypeDataToSign, sdk.AttributeKeyModule, cbrtypes.ModuleName)).String()
	EventQueryFarmingClaimAll = tmquery.MustParse(
		fmt.Sprintf("tm.event='Tx' AND %s.%s EXISTS", farmingtypes.EventTypeClaimAll, farmingtypes.AttributeKeyAddress)).String()
)

func MonitorTendermintEvent(nodeURI, eventQuery string, handleEvents func(events map[string][]string)) {
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

	txs, err := client.Subscribe(context.Background(), "monitor", eventQuery)
	if err != nil {
		log.Errorln("ws client subscribe error", err)
		return
	}

	for e := range txs {
		handleEvents(e.Events)
	}
}

func (r *Relayer) monitorSgnSlash() {
	MonitorTendermintEvent(
		r.Transactor.CliCtx.NodeURI,
		EventQuerySlash,
		func(events map[string][]string) {
			if !r.isBonded() {
				return
			}
			for _, nonceStr := range events[fmt.Sprintf("%s.%s", slashtypes.EventTypeSlash, slashtypes.AttributeKeyNonce)] {
				nonce, err :=
					strconv.ParseUint(nonceStr, 10, 64)
				if err != nil {
					log.Errorln("Parse slash nonce error", err)
					return
				}

				slashEvent := NewSlashEvent(nonce)
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

				err = r.dbSet(GetSlashKey(slashEvent.Nonce), slashEvent.MustMarshal())
				if err != nil {
					log.Errorln("db Set err", err)
				}
			}
		})
}

func (r *Relayer) monitorSgnCbrDataToSign() {
	MonitorTendermintEvent(
		r.Transactor.CliCtx.NodeURI,
		EventQueryCbridge,
		func(events map[string][]string) {
			if !r.isBonded() {
				return
			}
			tmEventType := events["tm.event"][0]
			if tmEventType != tm.EventTx && tmEventType != tm.EventNewBlock {
				return
			}
			sigTypes := events[fmt.Sprintf("%s.%s", cbrtypes.EventTypeDataToSign, cbrtypes.AttributeKeyType)]
			dataArr := events[fmt.Sprintf("%s.%s", cbrtypes.EventTypeDataToSign, cbrtypes.AttributeKeyData)]
			for i, sigType := range sigTypes {
				data := eth.Hex2Bytes(dataArr[i])
				// sign data first
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
				logmsg := fmt.Sprintf("Sign cBridge data, sigType: %s", sigType)
				switch sigType {
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
					r.setCbrSsUpdating()
					ss := new(cbrtypes.SortedSigners)
					err = ss.Unmarshal(data)
					if err != nil {
						log.Errorf("%s, failed to unmarshal sorted signers: %s", logmsg, err)
						return
					}
					log.Infof("%s: %s", logmsg, ss.String())
				}
				r.Transactor.AddTxMsg(msg)
			}
		})
}

func (r *Relayer) monitorSgnFarmingClaimAllEvent() {
	MonitorTendermintEvent(
		r.Transactor.CliCtx.NodeURI,
		EventQueryFarmingClaimAll,
		func(events map[string][]string) {
			if !r.isBonded() {
				return
			}
			for _, addr := range events[fmt.Sprintf("%s.%s", farmingtypes.EventTypeClaimAll, farmingtypes.AttributeKeyAddress)] {
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
		})
}
