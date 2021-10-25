package ops

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagChainId = "chainid"
	FlagTxHash  = "txhash"
)

var (
	chainid uint64
	txhash  string
	// receipt of the txhash
	txReceipt *ethtypes.Receipt
	cbr       *CbrOneChain
)

// GetSyncCmd
func GetSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync an event from onchain to sidechain",
		RunE:  client.ValidateCmd,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error
			cbr, err = newOneChain(chainid)
			if err != nil {
				log.Fatal("newOneChain err:", err)
			}
			txReceipt, err = cbr.TransactionReceipt(context.Background(), eth.Hex2Hash(txhash))
			if err != nil {
				log.Fatal("TransactionReceipt err:", err)
			}
		},
	}

	cmd.AddCommand(common.PostCommands(
		GetSyncSigners(),
		GetSyncEvent(),
	)...)

	cmd.PersistentFlags().Uint64Var(&chainid, FlagChainId, 0, "which chainid to query tx hash")
	cmd.PersistentFlags().StringVar(&txhash, FlagTxHash, "", "tx hash, will parse last event")
	return cmd
}

func GetSyncEvent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "event",
		Short: "Sync event from onchain, automatically figure out which event based on elog",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s ops sync event --chainid=883 --txhash="xxxxx"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			elog := *txReceipt.Logs[len(txReceipt.Logs)-1]
			evname, ev := parseEvAndName(cbr.contract, elog)
			log.Info(ev.PrettyLog(chainid))
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			err = verifyEvent(cliCtx, ev)
			if err != nil {
				log.Errorf("verifyEvent err: %s", err)
				return err
			}
			err = sendCbrOnchainEvent(cliCtx, chainid, evname, elog)
			if err != nil {
				log.Errorf("sendCbrOnchainEvent err: %s", err)
				return err
			}
			return nil
		},
	}
	return cmd
}

/*
another way is to use abi and event.Sig and compare to elog.Topics
cbrabi, _ := abi.JSON(strings.NewReader(eth.BridgeABI))
for evname, v := range cbrabi.Events {
	if eth.Hex2Hash(v.Sig) == elog.Topics[0] {
		// evname found
	}
}
*/
func parseEvAndName(cbr *cbrContract, elog ethtypes.Log) (string, hasPrettyLog) {
	var ev hasPrettyLog
	ev, err := cbr.ParseLiquidityAdded(elog)
	if err == nil {
		return cbrtypes.CbrEventLiqAdd, ev
	}
	ev, err = cbr.ParseSend(elog)
	if err == nil {
		return cbrtypes.CbrEventSend, ev
	}
	ev, err = cbr.ParseRelay(elog)
	if err == nil {
		return cbrtypes.CbrEventRelay, ev
	}
	ev, err = cbr.ParseWithdrawDone(elog)
	if err == nil {
		return cbrtypes.CbrEventWithdraw, ev
	}
	return "", nil
}

func verifyEvent(cliCtx client.Context, ev hasPrettyLog) error {
	switch e := ev.(type) {
	case *eth.BridgeLiquidityAdded:
		resp, err := cbrcli.QueryAddLiquidityStatus(cliCtx, &cbrtypes.QueryAddLiquidityStatusRequest{
			ChainId: chainid,
			SeqNum:  e.Seqnum,
		})
		if err != nil {
			return fmt.Errorf("QueryAddLiquidityStatus err: %s", err)
		}
		if resp.Status == cbrtypes.LPHistoryStatus_LP_COMPLETED {
			return fmt.Errorf("LiquidityAdded with seqNum %d on chain %d already synced", e.Seqnum, chainid)
		}
		return nil
	case *eth.BridgeSend:
		xferId := e.CalcXferId(chainid).Hex()
		resp, err := cbrcli.QueryTransferStatus(cliCtx, &cbrtypes.QueryTransferStatusRequest{
			TransferId: []string{xferId},
		})
		if err != nil {
			return fmt.Errorf("QueryAddLiquidityStatus err: %s", err)
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
		if resp.Status == cbrtypes.LPHistoryStatus_LP_COMPLETED {
			return fmt.Errorf("withdrawal with seqNum %d on chain %d already synced", e.Seqnum, chainid)
		}
		return nil
	}

	return nil
}

type hasPrettyLog interface {
	PrettyLog(uint64) string
}

func GetSyncSigners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signers",
		Short: "Sync signers from onchain",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s ops sync signers --chainid=883 --txhash="xxxxx"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			elog := *txReceipt.Logs[len(txReceipt.Logs)-1]
			ev, err := cbr.contract.ParseSignersUpdated(elog)
			if err != nil {
				log.Errorf("ParseSignersUpdated err: %s", err)
				return err
			}

			// check in store
			storedChainSigners, err := cbrcli.QueryChainSigners(cliCtx, chainid)
			if err != nil && !strings.Contains(err.Error(), "record not found") {
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

			err = sendCbrOnchainEvent(cliCtx, chainid, cbrtypes.CbrEventSignersUpdated, elog)
			if err != nil {
				log.Errorf("sendCbrOnchainEvent err: %s", err)
				return err
			}
			return nil
		},
	}
	return cmd
}

func sendCbrOnchainEvent(cliCtx client.Context, chainid uint64, evtype string, elog ethtypes.Log) error {
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

	elogJson, _ := json.Marshal(elog)
	onchev := &cbrtypes.OnChainEvent{
		Chainid: chainid,
		Evtype:  evtype,
		Elog:    elogJson,
	}
	data, _ := onchev.Marshal()
	msg.Updates = append(msg.Updates, &synctypes.ProposeUpdate{
		Type:    synctypes.DataType_CbrOnchainEvent,
		ChainId: chainid,
		Data:    data,
	})

	txr.CliSendTxMsgWaitMined(msg)
	return nil
}
