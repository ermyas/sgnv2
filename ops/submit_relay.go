package ops

import (
	"errors"
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagXFerId = "xferid"
)

func SubmitRelayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-relay",
		Short: "submit relay using XferId",
		Long: strings.TrimSpace(
			fmt.Sprintf(`
Example:
$ %s ops submit-relay --xferid=xxxxx"
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			xFerId := eth.Hex2Bytes(viper.GetString(FlagXFerId))
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

			// logic below is very much like relayer/cbr_puller.go:88 submitRelay
			logmsg := fmt.Sprintf("Process relay srcId %x", xFerId)

			relay, err := cbrcli.QueryRelay(txr.CliCtx, xFerId)
			if err != nil {
				return fmt.Errorf("%s. QueryRelay err: %s", logmsg, err)
			}

			relayOnChain := new(cbrtypes.RelayOnChain)
			err = relayOnChain.Unmarshal(relay.Relay)
			if err != nil {
				return fmt.Errorf("%s. Unmarshal relay.Relay err %s", logmsg, err)
			}

			chainSigners, err := cbrcli.QueryChainSigners(cliCtx, relayOnChain.GetDstChainId())
			if err != nil && !errors.Is(err, sdkerrors.ErrKeyNotFound) {
				log.Errorf("QueryChainSigners err: %s", err)
				return err
			}
			curss := chainSigners.GetSortedSigners()
			pass, sigsBytes := cbrtypes.ValidateSigQuorum(relay.SortedSigs, curss)
			if !pass {
				log.Warnf("%s. Not have enough sigs %s, curss %x", logmsg, relay.SignersStr(), curss)
				return nil
			}
			relayTransferId := relayOnChain.GetRelayOnChainTransferId()
			logmsg = fmt.Sprintf("%s dstId %x", logmsg, relayTransferId)
			cbr, err := newOneChain(relayOnChain.GetDstChainId())
			if err != nil {
				log.Fatal("newOneChain err:", err)
			}
			existRelay, existRelayErr := cbr.cbrContract.BridgeCaller.Transfers(&bind.CallOpts{}, relayTransferId)
			if existRelayErr != nil {
				// if fail to query, continue to send this relay, because we can not make sure whether the relay already exist.
				log.Warnln("fail to query transfer err:", existRelayErr)
			} else if existRelay {
				log.Infof("%s. dest transfer already exist on chain, skip it", logmsg)
				return nil
			}
			logmsg = fmt.Sprintf("srcXferId %x chain %d->%d", relayOnChain.GetSrcTransferId(), relayOnChain.GetSrcChainId(), relayOnChain.GetDstChainId())
			re, err := cbr.Transactor.TransactWaitMined(
				"cli submit relay",
				func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
					signers, powers := cbrtypes.SignersToEthArrays(curss)
					return cbr.cbrContract.Relay(opts, relay.Relay, sigsBytes, signers, powers)
				},
			)
			if err != nil {
				return err
			}
			log.Infof("%s. tx hash %s", logmsg, re.TxHash.Hex())
			return nil
		},
	}

	cmd.Flags().String(FlagXFerId, "", "transferId, used to retry submit relay")
	cmd.MarkFlagRequired(FlagXFerId)

	return cmd
}
