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
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagChainId = "chainid"
	FlagTxHash  = "txhash"
)

// GetSyncCmd
func GetSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "sync",
		Short:                      "Sync a change from onchain to sidechain",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		GetSyncSigners(),
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

			chainId := viper.GetUint64(FlagChainId)
			txHash := viper.GetString(FlagChainId)

			c, err := newOneChain(chainId)
			if err != nil {
				return err
			}

			receipt, err := c.TransactionReceipt(context.Background(), ec.HexToHash(txHash))
			if err != nil {
				log.Errorf("TransactionReceipt err: %s", err)
				return err
			}

			elog := *receipt.Logs[len(receipt.Logs)-1]
			ev, err := c.contract.ParseSignersUpdated(elog)
			if err != nil {
				log.Errorf("ParseSignersUpdated err: %s", err)
				return err
			}

			// check in store
			storedChainSigners, err := cbrcli.QueryChainSigners(cliCtx, chainId)
			if err != nil {
				log.Errorf("QueryChainSigners err: %s", err)
				return err
			}

			if relayer.EqualSigners(storedChainSigners.GetSortedSigners(), ev) {
				log.Infof("Signers already updated")
				return nil
			}

			// check on chain
			ssHash, err := c.contract.SsHash(&bind.CallOpts{})
			if err != nil {
				log.Errorf("query ssHash err: %s", err)
				return err
			}
			curssHash := eth.Bytes2Hash(crypto.Keccak256(eth.SignerBytes(ev.Signers, ev.Powers)))
			if curssHash != ssHash {
				log.Errorf("curss hash %x not match onchain values: %x", curssHash, ssHash)
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
				log.Errorf("NewTransactor err: %s", err)
				return err
			}

			// find all events need to be sent out, batch into one msg
			msg := &synctypes.MsgProposeUpdates{
				Sender:  txr.Key.GetAddress().String(),
				Updates: make([]*synctypes.ProposeUpdate, 0),
			}

			elogJson, _ := json.Marshal(elog)
			onchev := &cbrtypes.OnChainEvent{
				Chainid: chainId,
				Evtype:  cbrtypes.CbrEventSignersUpdated,
				Elog:    elogJson,
			}
			data, _ := onchev.Marshal()
			msg.Updates = append(msg.Updates, &synctypes.ProposeUpdate{
				Type:       synctypes.DataType_CbrOnchainEvent,
				ChainId:    chainId,
				ChainBlock: 0, // why do we need this in ProposeUpdate?
				Data:       data,
			})

			txr.AddTxMsg(msg)

			return nil
		},
	}

	cmd.Flags().String(FlagChainId, "", "chain id")
	cmd.Flags().String(FlagTxHash, "", "the hash of the tx which triggered the SignersUpdated event")
	cmd.MarkFlagRequired(FlagChainId)
	cmd.MarkFlagRequired(FlagTxHash)

	return cmd
}
