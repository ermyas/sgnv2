package ops

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagFile  = "file"
	flagReqId = "reqid"
	flagQuery = "query"
)

func WithdrawValidatorFeeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-cbr-fee",
		Short: "validator query and submit cbridge fee withdrawal request",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = transactor.NewTransactor(
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

			chainId, err := cmd.Flags().GetUint64(flagChain)
			if err != nil {
				return err
			}
			reqId, err := cmd.Flags().GetUint64(flagReqId)
			if err != nil {
				return err
			}
			file, err := cmd.Flags().GetString(flagFile)
			if err != nil {
				return err
			}
			queryOnly, err := cmd.Flags().GetBool(flagQuery)
			if err != nil {
				return err
			}
			chainReqIds := make(map[uint64]([]uint64))
			if chainId != 0 && reqId != 0 {
				chainReqIds[chainId] = append(chainReqIds[chainId], reqId)
			} else if file != "" {
				f, err := os.Open(file)
				if err != nil {
					return err
				}
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					line := scanner.Text()
					fields := strings.Fields(line)
					if len(fields) != 3 {
						return fmt.Errorf("invalid file input: %s", line)
					}
					reqId, err := strconv.Atoi(fields[0])
					if err != nil {
						return err
					}
					chainId, err := strconv.Atoi(fields[1])
					if err != nil {
						return err
					}
					chainReqIds[uint64(chainId)] = append(chainReqIds[uint64(chainId)], uint64(reqId))
				}
			} else {
				return fmt.Errorf("invalid flag inputs")
			}

			valAddr := viper.GetString(common.FlagEthValidatorAddress)
			for chainId, reqIds := range chainReqIds {
				fmt.Printf("chainId: %d\n\n", chainId)
				cbr, err := newOneChain(chainId)
				if err != nil {
					return fmt.Errorf("newOneChain %d err: %w", chainId, err)
				}
				chainSigners, err := cbrcli.QueryChainSigners(cliCtx, chainId)
				if err != nil && !errors.Is(err, sdkerrors.ErrKeyNotFound) {
					return fmt.Errorf("QueryChainSigners err: %s", err)
				}
				curss := chainSigners.GetSortedSigners()
				signers, powers := cbrtypes.SignersToEthArrays(curss)
				for _, reqId := range reqIds {
					req := &cbrtypes.QueryWithdrawLiquidityStatusRequest{
						SeqNum:  reqId,
						UsrAddr: valAddr,
					}
					resp, err := cbrcli.QueryWithdrawLiquidityStatus(cliCtx, req)
					if err != nil {
						return err
					}
					withdrawOnChain := new(cbrtypes.WithdrawOnchain)
					err = withdrawOnChain.Unmarshal(resp.Detail.WdOnchain)
					if err != nil {
						return err
					}
					fmt.Printf("reqId: %d\n", reqId)
					fmt.Printf("status: %s\n", resp.Status)
					fmt.Printf("withdraw message: %s, %s, last req time %s \n",
						withdrawOnChain.String(), resp.Detail.SignersStr(), common.TsSecToTime(uint64(resp.Detail.LastReqTime)))

					if !queryOnly && resp.Status != cbrtypes.WithdrawStatus_WD_COMPLETED {
						tx, err := cbr.Transactor.Transact(
							nil,
							func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
								return cbr.contract.Withdraw(opts, resp.Detail.WdOnchain, resp.Detail.GetSortedSigsBytes(), signers, powers)
							},
						)
						if err != nil {
							fmt.Printf("submit withdraw chain %d reqid %d. err: %s\n\n", chainId, reqId, err)
							continue
						}
						fmt.Printf("submit withdraw chain %d reqid %d. tx hash %x\n", chainId, reqId, tx.Hash())
						time.Sleep(time.Second)
					}
					fmt.Println()
				}
			}

			return nil
		},
	}
	cmd.Flags().Uint64(flagChain, 0, "chain id")
	cmd.Flags().Uint64(flagReqId, 0, "withdraw request id")
	cmd.Flags().String(flagFile, "", "file that contains a list of <reqid, chainid, token> tuples")
	cmd.Flags().Bool(flagQuery, false, "only query the withdrawals")
	return cmd
}
