package ops

import (
	"errors"
	"fmt"
	"strings"

	ethutils "github.com/celer-network/goutils/eth"
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
	FlagXferId       = "xferid"
	FlagGasPriceGwei = "gasprice"
)

func SubmitRelayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-relay",
		Short: "Submit cbridge relay using source transferId",
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

			xferId := eth.Hex2Bytes(viper.GetString(FlagXferId))
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

			relay, err := cbrcli.QueryRelay(txr.CliCtx, xferId)
			if err != nil {
				return fmt.Errorf("QueryRelay err: %w", err)
			}

			relayOnChain := new(cbrtypes.RelayOnChain)
			err = relayOnChain.Unmarshal(relay.Relay)
			if err != nil {
				return fmt.Errorf("Unmarshal relay.Relay err %w", err)
			}

			chainSigners, err := cbrcli.QueryChainSigners(cliCtx, relayOnChain.GetDstChainId())
			if err != nil && !errors.Is(err, sdkerrors.ErrKeyNotFound) {
				return fmt.Errorf("QueryChainSigners err: %w", err)
			}
			curss := chainSigners.GetSortedSigners()
			pass, sigsBytes := cbrtypes.ValidateSigQuorum(relay.SortedSigs, curss)
			if !pass {
				return fmt.Errorf("Not have enough sigs %s, curss %x", relay.SignersStr(), curss)
			}
			relayId := relayOnChain.GetRelayOnChainTransferId()
			cbr, err := newOneChain(relayOnChain.GetDstChainId())
			if err != nil {
				return fmt.Errorf("newOneChain err: %w", err)
			}
			existRelay, err := cbr.cbrContract.BridgeCaller.Transfers(&bind.CallOpts{}, relayId)
			if err != nil {
				return fmt.Errorf("fail to query transfer err: %w", err)
			} else if existRelay {
				log.Infof("dest transfer %x already exist on chain, skip it", relayId)
				return nil
			}
			logmsg := fmt.Sprintf("chain %d->%d transfer %x->%x",
				relayOnChain.GetSrcChainId(), relayOnChain.GetDstChainId(), xferId, relayId)
			tx, err := cbr.Transactor.Transact(
				nil,
				func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
					signers, powers := cbrtypes.SignersToEthArrays(curss)
					return cbr.cbrContract.Relay(opts, relay.Relay, sigsBytes, signers, powers)
				},
				ethutils.WithForceGasGwei(viper.GetUint64(FlagGasPriceGwei)),
			)
			if err != nil {
				return err
			}
			log.Infof("%s. tx hash %x, gas %s", logmsg, tx.Hash(), tx.GasPrice())
			return nil
		},
	}

	cmd.Flags().String(FlagXferId, "", "transferId, used to retry submit relay")
	cmd.Flags().Uint64(FlagGasPriceGwei, 0, "force gas price in gwei")
	cmd.MarkFlagRequired(FlagXferId)

	return cmd
}
