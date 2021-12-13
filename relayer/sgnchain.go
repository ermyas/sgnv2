package relayer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	pegbrcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	slashingcli "github.com/celer-network/sgn-v2/x/slashing/client/cli"
	slashingtypes "github.com/celer-network/sgn-v2/x/slashing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	"github.com/tendermint/tendermint/rpc/client/http"
	tm "github.com/tendermint/tendermint/types"
)

var (
	EventQuerySlash = tmquery.MustParse(
		fmt.Sprintf("tm.event='NewBlock' AND %s.%s EXISTS", slashingtypes.EventTypeSlash, slashingtypes.AttributeKeyNonce)).String()

	EventQueryCbridge = tmquery.MustParse(
		fmt.Sprintf("%s.%s='%s'", cbrtypes.EventTypeDataToSign, sdk.AttributeKeyModule, cbrtypes.ModuleName)).String()
	EventQueryPegMint = tmquery.MustParse(
		fmt.Sprintf("%s.%s='%s'", pegbrtypes.EventTypeMintToSign, sdk.AttributeKeyModule, pegbrtypes.ModuleName)).String()
	EventQueryPegWithdraw = tmquery.MustParse(
		fmt.Sprintf("%s.%s='%s'", pegbrtypes.EventTypeWithdrawToSign, sdk.AttributeKeyModule, pegbrtypes.ModuleName)).String()
	EventQueryFarmingClaimAll = tmquery.MustParse(
		fmt.Sprintf("tm.event='Tx' AND %s.%s EXISTS", farmingtypes.EventTypeClaimAll, farmingtypes.AttributeKeyAddress)).String()

	EventQueryDistributionClaimAllStakingReward = tmquery.MustParse(
		fmt.Sprintf("tm.event='Tx' AND %s.%s EXISTS",
			distrtypes.EventTypeClaimAllStakingReward, distrtypes.AttributeKeyDelegatorAddress)).String()
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

	res, err := client.Subscribe(context.Background(), "monitor", eventQuery)
	if err != nil {
		log.Errorln("ws client subscribe error", err)
		return
	}

	for e := range res {
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
			for _, nonceStr := range events[fmt.Sprintf("%s.%s", slashingtypes.EventTypeSlash, slashingtypes.AttributeKeyNonce)] {
				nonce, err := strconv.ParseUint(nonceStr, 10, 64)
				if err != nil {
					log.Errorln("Parse slash nonce error", err)
					return
				}

				slashEvent := NewSlashEvent(nonce)
				slash, err := slashingcli.QuerySlash(r.Transactor.CliCtx, slashEvent.Nonce)
				if err != nil {
					log.Errorf("Query slash %d err %s", slashEvent.Nonce, err)
					return
				}
				log.Infof("New slash to %x, reason %s, nonce %d", slash.SlashOnChain.Validator, slash.Reason, slashEvent.Nonce)

				dataToSign := slash.EncodeDataToSign(r.EthClient.ChainId, r.EthClient.Contracts.Staking.Address)
				sig, err := r.EthClient.SignEthMessage(dataToSign)
				if err != nil {
					log.Errorln("SignEthMessage err", err)
					return
				}

				msg := slashingtypes.NewMsgSignSlash(slashEvent.Nonce, sig, r.Transactor.Key.GetAddress())
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
			dataTypes := events[fmt.Sprintf("%s.%s", cbrtypes.EventTypeDataToSign, cbrtypes.AttributeKeyType)]
			dataArr := events[fmt.Sprintf("%s.%s", cbrtypes.EventTypeDataToSign, cbrtypes.AttributeKeyData)]
			for i, dataType := range dataTypes {
				data := eth.Hex2Bytes(dataArr[i])
				msg := &cbrtypes.MsgSendMySig{
					Data:    data,
					Creator: r.Transactor.Key.GetAddress().String(),
				}
				logmsg := fmt.Sprintf("Sign cBridge data, dataType: %s", dataType)
				switch dataType {
				case cbrtypes.SignDataType_RELAY.String():
					msg.Datatype = cbrtypes.SignDataType_RELAY
					relay := new(cbrtypes.RelayOnChain)
					err := relay.Unmarshal(data)
					if err != nil {
						log.Errorf("%s, failed to unmarshal RelayOnChain: %s", logmsg, err)
						continue
					}
					chain := r.cbrMgr[relay.DstChainId]
					if chain == nil {
						log.Errorf("%s, no cbrMgr %d found", logmsg, relay.DstChainId)
						continue
					}
					dataToSign := cbrtypes.EncodeRelayOnChainToSign(relay.DstChainId, chain.cbrContract.Address, data)
					sig, err := r.EthClient.SignEthMessage(dataToSign)
					if err != nil {
						log.Errorf("%s, sign msg err: %s", logmsg, err)
						return
					}
					msg.MySigs = append(msg.MySigs, &cbrtypes.MySig{ChainId: relay.DstChainId, Sig: sig})

					relayRequest := NewRelayRequest(relay.SrcTransferId, relay.DstChainId)
					err = r.dbSet(GetCbrXferKey(relayRequest.XferId, relay.DstChainId), relayRequest.MustMarshal())
					if err != nil {
						log.Errorf("%s, db Set err: %s", logmsg, err)
					}
					log.Infof("%s: %s", logmsg, relay.String())
				case cbrtypes.SignDataType_WITHDRAW.String():
					msg.Datatype = cbrtypes.SignDataType_WITHDRAW
					withdraw := new(cbrtypes.WithdrawOnchain)
					err := withdraw.Unmarshal(data)
					if err != nil {
						log.Errorf("%s, failed to unmarshal WithdrawOnchain: %s", logmsg, err)
						continue
					}
					chain := r.cbrMgr[withdraw.Chainid]
					if chain == nil {
						log.Errorf("%s, no cbrMgr %d found", logmsg, withdraw.Chainid)
						continue
					}
					dataToSign := cbrtypes.EncodeWithdrawOnchainToSign(withdraw.Chainid, chain.cbrContract.Address, data)
					sig, err := r.EthClient.SignEthMessage(dataToSign)
					if err != nil {
						log.Errorf("%s, sign msg err: %s", logmsg, err)
						continue
					}
					msg.MySigs = append(msg.MySigs, &cbrtypes.MySig{ChainId: withdraw.Chainid, Sig: sig})

					log.Infof("%s: %s", logmsg, withdraw.String())
				case cbrtypes.SignDataType_SIGNERS.String():
					msg.Datatype = cbrtypes.SignDataType_SIGNERS
					for chainId, c := range r.cbrMgr {
						dataToSign := cbrtypes.EncodeSignersUpdateToSign(chainId, c.cbrContract.Address, data)
						sig, err := r.EthClient.SignEthMessage(dataToSign)
						if err != nil {
							log.Errorf("%s, sign msg err: %s", logmsg, err)
							continue
						}
						msg.MySigs = append(msg.MySigs, &cbrtypes.MySig{ChainId: chainId, Sig: sig})
					}
					r.setCbrSsUpdating()
					log.Infof("%s", logmsg)
				}
				r.Transactor.AddTxMsg(msg)
			}
		})
}

func (r *Relayer) monitorSgnPegMintToSign() {
	MonitorTendermintEvent(
		r.Transactor.CliCtx.NodeURI,
		EventQueryPegMint,
		func(events map[string][]string) {
			if !r.isBonded() {
				return
			}
			tmEventType := events["tm.event"][0]
			if tmEventType != tm.EventTx && tmEventType != tm.EventNewBlock {
				return
			}

			mintIds := events[fmt.Sprintf("%s.%s", pegbrtypes.EventTypeMintToSign, pegbrtypes.AttributeKeyMintId)]
			mintChainIds := events[fmt.Sprintf("%s.%s", pegbrtypes.EventTypeMintToSign, pegbrtypes.AttributeKeyMintChainId)]
			for i, mintId := range mintIds {
				mintChainId, _ := strconv.ParseUint(mintChainIds[i], 10, 64)

				// sign data first
				mintInfo, err := pegbrcli.QueryMintInfo(r.Transactor.CliCtx, mintId)
				if err != nil {
					log.Error(err)
					return
				}

				mintOnChain := new(pegbrtypes.MintOnChain)
				err = mintOnChain.Unmarshal(mintInfo.MintProtoBytes)
				if err != nil {
					log.Errorf("Unmarshal mintInfo.MintProtoBytes err %s", err)
					return
				}

				sig, err := r.EthClient.SignEthMessage(mintInfo.EncodeDataToSign(r.cbrMgr[mintChainId].pegContracts.bridge.Address))
				if err != nil {
					log.Error(err)
					return
				}

				msg := &pegbrtypes.MsgSignMint{
					MintId:    mintId,
					Signature: sig,
					Sender:    r.Transactor.Key.GetAddress().String(),
				}
				r.Transactor.AddTxMsg(msg)
				mintRequest := NewMintRequest(eth.Hex2Bytes(mintId), mintChainId, mintOnChain.RefChainId, mintOnChain.RefId)
				err = r.dbSet(GetPegbrMintKey(mintChainId, mintRequest.DepositChainId, mintRequest.DepositId), mintRequest.MustMarshal())
				if err != nil {
					log.Errorf("db Set err: %s", err)
				}
			}
		})
}

func (r *Relayer) monitorSgnPegWithdrawToSign() {
	MonitorTendermintEvent(
		r.Transactor.CliCtx.NodeURI,
		EventQueryPegWithdraw,
		func(events map[string][]string) {
			if !r.isBonded() {
				return
			}
			tmEventType := events["tm.event"][0]
			if tmEventType != tm.EventTx && tmEventType != tm.EventNewBlock {
				return
			}

			wdIds := events[fmt.Sprintf("%s.%s", pegbrtypes.EventTypeWithdrawToSign, pegbrtypes.AttributeKeyWithdrawId)]
			wdChainIds := events[fmt.Sprintf("%s.%s", pegbrtypes.EventTypeWithdrawToSign, pegbrtypes.AttributeKeyWithdrawChainId)]
			for i, wdId := range wdIds {
				wdChainId, _ := strconv.ParseUint(wdChainIds[i], 10, 64)

				// sign data first
				wdInfo, err := pegbrcli.QueryWithdrawInfo(r.Transactor.CliCtx, wdId)
				if err != nil {
					log.Error(err)
					return
				}

				wdOnChain := new(pegbrtypes.WithdrawOnChain)
				err = wdOnChain.Unmarshal(wdInfo.WithdrawProtoBytes)
				if err != nil {
					log.Errorf("Unmarshal wdInfo.WithdrawProtoBytes err %s", err)
					return
				}

				sig, err := r.EthClient.SignEthMessage(wdInfo.EncodeDataToSign(r.cbrMgr[wdChainId].pegContracts.vault.Address))
				if err != nil {
					log.Error(err)
					return
				}

				msg := &pegbrtypes.MsgSignWithdraw{
					WithdrawId: wdId,
					Signature:  sig,
					Sender:     r.Transactor.Key.GetAddress().String(),
				}
				r.Transactor.AddTxMsg(msg)

				// RefChainId = 0 means fee claim, don't add a WithdrawRequest
				if wdOnChain.RefChainId == 0 {
					return
				}
				wdRequest := NewWithdrawRequest(eth.Hex2Bytes(wdId), wdChainId, wdOnChain.RefChainId, wdOnChain.RefId)
				err = r.dbSet(GetPegbrWdKey(wdChainId, wdOnChain.RefChainId, wdOnChain.RefId), wdRequest.MustMarshal())
				if err != nil {
					log.Errorf("db Set err: %s", err)
				}
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
					if details.ChainId != r.EthClient.ChainId {
						log.Errorf("Farming reward on chain %d not supported yet", details.ChainId)
						continue
					}
					dataToSign := details.EncodeDataToSign(r.EthClient.Contracts.FarmingRewards.Address)
					sig, err := r.EthClient.SignEthMessage(dataToSign)
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

func (r *Relayer) monitorSgnDistributionClaimAllStakingRewardEvent() {
	MonitorTendermintEvent(
		r.Transactor.CliCtx.NodeURI,
		EventQueryDistributionClaimAllStakingReward,
		func(events map[string][]string) {
			if !r.isBonded() {
				return
			}
			for _, addr := range events[fmt.Sprintf("%s.%s",
				distrtypes.EventTypeClaimAllStakingReward, distrtypes.AttributeKeyDelegatorAddress)] {
				queryClient := distrtypes.NewQueryClient(r.Transactor.CliCtx)
				stakingRewardClaimInfo, err := queryClient.StakingRewardClaimInfo(
					context.Background(),
					&distrtypes.QueryStakingRewardClaimInfoRequest{
						DelegatorAddress: addr,
					},
				)
				if err != nil {
					log.Errorf("Query StakingRewardClaimInfo err %s", err)
					return
				}
				dataToSign := stakingRewardClaimInfo.RewardClaimInfo.EncodeDataToSign(
					r.EthClient.ChainId, r.EthClient.Contracts.StakingReward.Address)
				sig, err := r.EthClient.SignEthMessage(dataToSign)
				if err != nil {
					log.Errorln("SignEthMessage err", err)
					return
				}
				msg := distrtypes.NewMsgSignStakingReward(eth.Hex2Addr(addr), r.Transactor.Key.GetAddress(), sig)
				r.Transactor.AddTxMsg(msg)
			}
		})
}
