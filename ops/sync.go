package ops

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/privval"
)

const (
	FlagChainId = "chainid"
	FlagTxHash  = "txhash"
	FlagEvName  = "evname"
	FlagValAddr = "valaddr"
	FlagDelAddr = "deladdr"
)

// GetSyncCmd
func GetSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "sync",
		Short:                      "Sync an event from onchain to sidechain",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		GetSyncSigners(),
		GetSyncCbrEvent(),
		GetSyncStaking(),
	)...)

	return cmd
}

func GetSyncSigners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signers",
		Short: "Sync signers from onchain",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s ops sync signers --chainid=883 --txhash="0xxx"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			chainid := viper.GetUint64(FlagChainId)
			txhash := viper.GetString(FlagTxHash)
			cbr, txReceipt, err := setupCbr(chainid, txhash)
			if err != nil {
				return err
			}

			elog := eth.FindMatchCbrEvent(cbrtypes.CbrEventSignersUpdated, cbr.contract.Address, txReceipt.Logs)

			if elog == nil {
				log.Errorln("no match event found in tx:", txhash)
				return fmt.Errorf("no match event found in tx: %s", txhash)
			}
			ev, err := cbr.contract.ParseSignersUpdated(*elog)
			if err != nil {
				log.Errorf("ParseSignersUpdated err: %s", err)
				return err
			}

			// check in store
			storedChainSigners, err := cbrcli.QueryChainSigners(cliCtx, chainid)
			if err != nil && !errors.Is(err, sdkerrors.ErrKeyNotFound) {
				log.Errorf("QueryChainSigners err: %s", err)
				return err
			}

			if storedChainSigners != nil && relayer.EqualSigners(storedChainSigners.GetSortedSigners(), ev) {
				log.Infof("Signers already updated")
				return nil
			}

			// check on chain
			ssHash, err := cbr.contract.SsHash(&bind.CallOpts{})
			if err != nil {
				log.Errorf("query ssHash err: %s", err)
				return err
			}
			curssHash := eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(ev.Signers, ev.Powers)))
			if curssHash != ssHash {
				log.Errorf("curss hash %x not match onchain values: %x", curssHash, ssHash)
				return err
			}

			err = sendCbrOnchainEvent(cliCtx, chainid, cbrtypes.CbrEventSignersUpdated, *elog)
			if err != nil {
				log.Errorf("sendCbrOnchainEvent err: %s", err)
				return err
			}
			return nil
		},
	}

	cmd.Flags().Uint64(FlagChainId, 0, "which chainid to query tx hash")
	cmd.Flags().String(FlagTxHash, "", "tx hash, will parse event with same ID as SignersUpdated")
	cmd.MarkFlagRequired(FlagChainId)
	cmd.MarkFlagRequired(FlagTxHash)

	return cmd
}

func GetSyncCbrEvent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "event",
		Short: "Sync bridge event from onchain",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s ops sync event --chainid=883 --txhash="0xxx" --evname="Send"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			chainid := viper.GetUint64(FlagChainId)
			txhash := viper.GetString(FlagTxHash)
			evname := viper.GetString(FlagEvName)

			return SyncCbrEvent(cliCtx, chainid, txhash, evname)
		},
	}

	cmd.Flags().Uint64(FlagChainId, 0, "which chainid to query tx hash")
	cmd.Flags().String(FlagTxHash, "", "tx hash, will parse event with same ID as evname")
	cmd.Flags().String(FlagEvName, "", "ev name, the name of the parsed event")
	cmd.MarkFlagRequired(FlagChainId)
	cmd.MarkFlagRequired(FlagTxHash)
	cmd.MarkFlagRequired(FlagEvName)

	return cmd
}

func SyncCbrEvent(cliCtx client.Context, chainid uint64, txhash string, evname string) error {
	cbr, txReceipt, err := setupCbr(chainid, txhash)
	if err != nil {
		return err
	}

	elog := eth.FindMatchCbrEvent(evname, cbr.contract.Address, txReceipt.Logs)

	if elog == nil {
		log.Errorln("no match event found in tx:", txhash)
		return fmt.Errorf("no match event found in tx: %s", txhash)
	}

	ev := parseCbrEv(cbr.contract, *elog, evname)
	if ev == nil {
		log.Errorf("not a valid bridge event tx: %s", txhash)
		return fmt.Errorf("not a valid bridge event tx: %s", txhash)
	}
	log.Info(ev.PrettyLog(chainid))

	err = verifyEvent(cliCtx, ev, chainid)
	if err != nil {
		log.Errorf("verifyEvent err: %s", err)
		return err
	}
	err = sendCbrOnchainEvent(cliCtx, chainid, evname, *elog)
	if err != nil {
		log.Errorf("sendCbrOnchainEvent err: %s", err)
		return err
	}
	return nil
}

func GetSyncStaking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking",
		Short: "Sync latest staking info from onchain of the given validator addr. Delegation info will be synced also if delegator addr is provided.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s ops sync staking --valaddr="0xxx" --deladdr="0xxx"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			valAddr := viper.GetString(FlagValAddr)

			ethClient, err := newEthClient()
			if err != nil {
				log.Fatal("newEthClient err:", err)
			}

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			updates := make([]*synctypes.ProposeUpdate, 0)

			// 1. Check validator SGN address
			log.Infoln("Check SGN address")
			sgnAddr, err := ethClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, common.Hex2Addr(valAddr))
			if err != nil {
				log.Errorf("Failed to query contract SGN address err: %s", err)
				return err
			}
			exist, err := stakingcli.QuerySgnAccount(cliCtx, sdk.AccAddress(sgnAddr).String())
			if err != nil {
				if strings.Contains(err.Error(), "sgn account not found") {
					exist = false
				} else {
					log.Errorf("Failed to query SGN account address err: %s", err)
					return err
				}
			}
			if !exist {
				log.Infoln("SGN address needs update")
				updateVal := &stakingtypes.Validator{
					EthAddress: valAddr,
					SgnAddress: sdk.AccAddress(sgnAddr).String(),
				}
				updates = append(updates, &synctypes.ProposeUpdate{
					Type: synctypes.DataType_ValidatorSgnAddr,
					Data: cliCtx.Codec.MustMarshal(updateVal),
				})
			} else {
				log.Infoln("SGN address needs no update")
			}

			if len(updates) > 0 {
				err = sendSgnTxMsg(cliCtx, updates)
				if err != nil {
					log.Errorf("sendSgnTxMsg err: %s", err)
					return err
				}
				log.Infoln("Sync validator SGN address request submitted, please wait a little bit and try to execute command again to sync others...")
				return nil
			}

			storeVal, err := stakingcli.QueryValidator(cliCtx, valAddr)
			if err != nil && !strings.Contains(err.Error(), "validator not found") {
				return err
			}

			if (err != nil && strings.Contains(err.Error(), "validator not found")) || storeVal == nil {
				serverCtx := server.GetServerContextFromCmd(cmd)
				filePV := privval.LoadOrGenFilePV(serverCtx.Config.PrivValidatorKeyFile(), serverCtx.Config.PrivValidatorStateFile())
				tmValPubKey, err := filePV.GetPubKey()
				if err != nil {
					return err
				}
				valPubKey, err := cryptocodec.FromTmPubKeyInterface(tmValPubKey)
				if err != nil {
					return err
				}
				pkAny, err := codectypes.NewAnyWithValue(valPubKey)
				if err != nil {
					return fmt.Errorf("failed to generate pkAny, %w", err)
				}

				storeVal = &stakingtypes.Validator{
					ConsensusPubkey: pkAny,
				}
			}

			// 2. Check validator params
			log.Infoln("Check validator params")
			ethVal, err := ethClient.Contracts.Staking.Validators(&bind.CallOpts{}, common.Hex2Addr(valAddr))
			if err != nil {
				log.Errorf("Failed to query contract validator info: %s", err)
				return err
			}
			if eth.Addr2Hex(ethVal.Signer) != storeVal.EthSigner ||
				!sdk.NewDec(int64(ethVal.CommissionRate)).QuoInt64(eth.CommissionRateBase).Equal(storeVal.CommissionRate) {
				log.Infoln("Validator params need update")
				updateVal := &stakingtypes.Validator{
					EthAddress:      valAddr,
					EthSigner:       eth.Addr2Hex(ethVal.Signer),
					SgnAddress:      sdk.AccAddress(sgnAddr).String(),
					ConsensusPubkey: storeVal.ConsensusPubkey,
					CommissionRate:  sdk.NewDec(int64(ethVal.CommissionRate)).QuoInt64(eth.CommissionRateBase),
				}

				updates = append(updates, &synctypes.ProposeUpdate{
					Type: synctypes.DataType_ValidatorParams,
					Data: cliCtx.Codec.MustMarshal(updateVal),
				})
			} else {
				log.Infoln("Validator params need no update")
			}

			// 3. Check validator states
			log.Infoln("Check validator states")
			if stakingtypes.BondStatus(ethVal.Status) != storeVal.Status ||
				!sdk.NewIntFromBigInt(ethVal.Tokens).Equal(storeVal.Tokens) ||
				!sdk.NewIntFromBigInt(ethVal.Shares).Equal(storeVal.DelegatorShares) {
				log.Infoln("Validator states need update")
				updateVal := &stakingtypes.Validator{
					EthAddress:      valAddr,
					Status:          stakingtypes.BondStatus(ethVal.Status),
					Tokens:          sdk.NewIntFromBigInt(ethVal.Tokens),
					DelegatorShares: sdk.NewIntFromBigInt(ethVal.Shares),
				}

				updates = append(updates, &synctypes.ProposeUpdate{
					Type: synctypes.DataType_ValidatorStates,
					Data: cliCtx.Codec.MustMarshal(updateVal),
				})
			} else {
				log.Infoln("Validator states need no update")
			}

			// 4. Check delegation info
			delAddr := viper.GetString(FlagDelAddr)
			if delAddr != "" {
				log.Infoln("Check delegator shares")
				ethDel, err := ethClient.Contracts.Staking.GetDelegatorInfo(&bind.CallOpts{}, common.Hex2Addr(valAddr), common.Hex2Addr(delAddr))
				if err != nil {
					log.Errorf("Failed to query delegator info err: %s", err)
					return nil
				}

				updateDel := &stakingtypes.Delegation{
					DelegatorAddress: delAddr,
					ValidatorAddress: valAddr,
					Shares:           sdk.NewIntFromBigInt(ethDel.Shares),
				}

				storeDel, _ := stakingcli.QueryDelegation(cliCtx, valAddr, delAddr)

				if storeDel == nil || !updateDel.Shares.Equal(storeDel.Shares) {
					log.Infoln("Delegator shares need update")
					updates = append(updates, &synctypes.ProposeUpdate{
						Type: synctypes.DataType_DelegatorShares,
						Data: cliCtx.Codec.MustMarshal(updateDel),
					})
				} else {
					log.Infoln("Delegator shares need no update")
				}
			}

			if len(updates) > 0 {
				err = sendSgnTxMsg(cliCtx, updates)
				if err != nil {
					log.Errorf("sendSgnTxMsg err: %s", err)
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().String(FlagValAddr, "", "validator address, required")
	cmd.Flags().String(FlagDelAddr, "", "delagator address, optional. must be presented when sync delegation info")
	cmd.MarkFlagRequired(FlagValAddr)

	return cmd
}

func setupCbr(chainid uint64, txhash string) (cbr *CbrOneChain, txReceipt *ethtypes.Receipt, err error) {
	cbr, err = newOneChain(chainid)
	if err != nil {
		log.Fatal("newOneChain err:", err)
	}
	txReceipt, err = cbr.TransactionReceipt(context.Background(), eth.Hex2Hash(txhash))
	if err != nil {
		log.Errorln("TransactionReceipt err:", err)
		return
	}
	return
}

func parseCbrEv(cbr *cbrContract, elog ethtypes.Log, evname string) hasPrettyLog {
	var ev hasPrettyLog
	switch evname {
	case cbrtypes.CbrEventLiqAdd:
		ev, _ = cbr.ParseLiquidityAdded(elog)
	case cbrtypes.CbrEventSend:
		ev, _ = cbr.ParseSend(elog)
	case cbrtypes.CbrEventRelay:
		ev, _ = cbr.ParseRelay(elog)
	case cbrtypes.CbrEventWithdraw:
		ev, _ = cbr.ParseWithdrawDone(elog)
	default:
		ev = nil
	}

	return ev
}

func verifyEvent(cliCtx client.Context, ev hasPrettyLog, chainid uint64) error {
	switch e := ev.(type) {
	case *eth.BridgeLiquidityAdded:
		resp, err := cbrcli.QueryAddLiquidityStatus(cliCtx, &cbrtypes.QueryAddLiquidityStatusRequest{
			ChainId: chainid,
			SeqNum:  e.Seqnum,
		})
		if err != nil {
			return fmt.Errorf("QueryAddLiquidityStatus err: %s", err)
		}
		if resp.Status == cbrtypes.WithdrawStatus_WD_COMPLETED {
			return fmt.Errorf("LiquidityAdded with seqNum %d on chain %d already synced", e.Seqnum, chainid)
		}
		return nil
	case *eth.BridgeSend:
		xferId := e.CalcXferId(chainid).Hex()
		resp, err := cbrcli.QueryTransferStatus(cliCtx, &cbrtypes.QueryTransferStatusRequest{
			TransferId: []string{xferId},
		})
		if err != nil {
			return fmt.Errorf("QueryTransferStatus err: %s", err)
		}
		if resp.Status[xferId].SgnStatus != cbrtypes.XferStatus_UNKNOWN {
			return fmt.Errorf("xfer with xferId %s from src chain %d already synced", xferId, chainid)
		}
		return nil
	case *eth.BridgeRelay:
		return nil
	case *eth.BridgeWithdrawDone:
		resp, err := cbrcli.QueryWithdrawLiquidityStatus(cliCtx, &cbrtypes.QueryWithdrawLiquidityStatusRequest{
			SeqNum:  e.Seqnum,
			UsrAddr: e.Receiver.String(),
		})
		if err != nil {
			return fmt.Errorf("QueryWithdrawLiquidityStatus err: %s", err)
		}
		if resp.Status == cbrtypes.WithdrawStatus_WD_COMPLETED {
			return fmt.Errorf("withdrawal with seqNum %d on chain %d already synced", e.Seqnum, chainid)
		}
		return nil
	}

	return nil
}

type hasPrettyLog interface {
	PrettyLog(uint64) string
}

func sendCbrOnchainEvent(cliCtx client.Context, chainid uint64, evtype string, elog ethtypes.Log) error {
	elogJson, _ := json.Marshal(elog)
	onchev := &cbrtypes.OnChainEvent{
		Chainid: chainid,
		Evtype:  evtype,
		Elog:    elogJson,
	}
	data, _ := onchev.Marshal()
	return sendSgnTxMsg(cliCtx, []*synctypes.ProposeUpdate{{
		Type:    synctypes.DataType_CbrOnchainEvent,
		ChainId: chainid,
		Data:    data,
	}})
}

func sendSgnTxMsg(cliCtx client.Context, updates []*synctypes.ProposeUpdate) error {
	txr, err := transactor.NewTransactor(
		cliCtx.HomeDir,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		cliCtx.LegacyAmino,
		cliCtx.Codec,
		cliCtx.InterfaceRegistry,
	)
	if err != nil {
		return fmt.Errorf("NewTransactor err: %w", err)
	}

	// find all events need to be sent out, batch into one msg
	msg := &synctypes.MsgProposeUpdates{
		Sender:  txr.Key.GetAddress().String(),
		Updates: make([]*synctypes.ProposeUpdate, 0),
	}

	msg.Updates = append(msg.Updates, updates...)

	txr.CliSendTxMsgWaitMined(msg)
	return nil
}
