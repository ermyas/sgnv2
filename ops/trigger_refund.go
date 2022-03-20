package ops

import (
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// this is to fix a a historical bug that EXCEED_MAX_OUT_AMOUNT transfer was not refund properly
func TriggerSetRefundCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trigger-set-refund",
		Short: "Trigger set refund of pool-based bridge send",
		RunE: func(cmd *cobra.Command, args []string) error {
			chainid := viper.GetUint64(FlagChainId)
			txhash := viper.GetString(FlagTxHash)

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

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

			cbr, txReceipt, err := setupCbr(chainid, txhash)
			if err != nil {
				return err
			}

			elog := eth.FindMatchContractEvent(eth.LiquidityBridge, cbrtypes.CbrEventSend, cbr.cbrContract.Address, txReceipt.Logs)
			if elog == nil {
				log.Errorln("no match event found in tx:", txhash)
				return fmt.Errorf("no match event found in tx: %s", txhash)
			}

			ev, err := cbr.cbrContract.ParseSend(*elog)
			if ev == nil {
				log.Errorf("not a valid bridge event tx: %s", txhash)
				return fmt.Errorf("not a valid bridge event tx: %s", txhash)
			}
			log.Info(ev.PrettyLog(chainid))

			xferId := eth.Hash(ev.TransferId)
			resp, err := cbrcli.QueryTransferStatus(cliCtx, &cbrtypes.QueryTransferStatusRequest{
				TransferId: []string{xferId.Hex()},
			})
			if err != nil {
				return fmt.Errorf("QueryTransferStatus err: %w", err)
			}
			status := resp.Status[xferId.Hex()].SgnStatus
			if status != cbrtypes.XferStatus_EXCEED_MAX_OUT_AMOUNT {
				return fmt.Errorf("invalid xfer %x status %s", xferId, status)
			}
			refund, err := cbrcli.QueryRefund(cliCtx, xferId.Bytes())
			if err != nil && !strings.Contains(err.Error(), "key not found") {
				return fmt.Errorf("QueryRefund err: %w", err)
			}
			if refund != nil {
				return fmt.Errorf("refund already exists for xfer %x", xferId)
			}

			msg := &cbrtypes.MsgTriggerSetRefund{
				SrcChainId: chainid,
				Sender:     eth.Addr2Hex(ev.Sender),
				Receiver:   eth.Addr2Hex(ev.Receiver),
				Token:      eth.Addr2Hex(ev.Token),
				Amount:     ev.Amount.String(),
				DstChainId: ev.DstChainId,
				Nonce:      ev.Nonce,
				Creator:    txr.Key.GetAddress().String(),
			}

			txr.CliSendTxMsgWaitMined(msg)
			return nil
		},
	}
	cmd.Flags().Uint64(FlagChainId, 0, "which chainid to query tx hash")
	cmd.Flags().String(FlagTxHash, "", "tx hash, will parse event with same ID as evname")
	cmd.MarkFlagRequired(FlagChainId)
	cmd.MarkFlagRequired(FlagTxHash)
	return cmd
}
