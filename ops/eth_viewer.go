package ops

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagBonded = "bonded"
)

func EthViewerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "view",
		Short:                      "View ethereum staking contract states",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		ViewParamsCommand(),
		ViewDelegatorCommand(),
		ViewValidatorsCommand(),
	)

	return cmd
}

func ViewParamsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Get staking contract params",
		RunE: func(cmd *cobra.Command, args []string) error {
			ethClient, err := common.NewEthClientFromConfig()
			if err != nil {
				return err
			}
			res, err := ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 2)
			if err != nil {
				return err
			}
			fmt.Printf("Unbounding period: %s eth block\n", res)

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 3)
			if err != nil {
				return err
			}
			fmt.Printf("Max bonded validators: %s\n", res)

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 4)
			if err != nil {
				return err
			}
			f := new(big.Float).Quo(new(big.Float).SetInt(res), new(big.Float).SetInt64(1e18))
			fmt.Printf("Min validator tokens: %s (%f CELR)\n", res, f)

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 5)
			if err != nil {
				return err
			}
			f = new(big.Float).Quo(new(big.Float).SetInt(res), new(big.Float).SetInt64(1e18))
			fmt.Printf("Min self delegation: %s (%f CELR)\n", res, f)

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 7)
			if err != nil {
				return err
			}
			fmt.Printf("Validator bonded interval: %s eth block\n", res)

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 8)
			if err != nil {
				return err
			}
			fmt.Printf("Max slash factor: %s (%.2f%%)\n", res, float32(res.Uint64())/float32(1e4))

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 0)
			if err != nil {
				return err
			}
			f = new(big.Float).Quo(new(big.Float).SetInt(res), new(big.Float).SetInt64(1e18))
			fmt.Printf("Proposal deposit: %s (%f CELR)\n", res, f)

			res, err = ethClient.Contracts.Staking.GetParamValue(&bind.CallOpts{}, 1)
			if err != nil {
				return err
			}
			fmt.Printf("Voting period: %s eth block\n", res)

			return nil
		},
	}
	return cmd
}

func ViewDelegatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator [delegator-eth-addr]",
		Short: "Get delegator info from eth staking contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ethClient, err := common.NewEthClientFromConfig()
			if err != nil {
				return err
			}
			delInfos, err := ethClient.Contracts.Viewer.GetDelegatorInfos(&bind.CallOpts{}, eth.Hex2Addr(args[0]))
			if err != nil {
				return err
			}
			bytes, err := json.MarshalIndent(&delInfos, "", "  ")
			if err != nil {
				return err
			}
			_, err = fmt.Println(string(bytes))
			return err
		},
	}
	return cmd
}

func ViewValidatorsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators",
		Short: "Get all validators info from eth staking contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getValidators(viper.GetBool(FlagBonded))
		},
	}
	cmd.Flags().Bool(FlagBonded, false, "Only bonded validators")
	return cmd
}

func getValidators(onlyBonded bool) error {
	ethClient, err := common.NewEthClientFromConfig()
	if err != nil {
		return err
	}
	var validators []eth.DataTypesValidatorInfo
	if onlyBonded {
		validators, err = ethClient.Contracts.Viewer.GetBondedValidatorInfos(&bind.CallOpts{})
	} else {
		validators, err = ethClient.Contracts.Viewer.GetValidatorInfos(&bind.CallOpts{})
	}
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(&validators, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(bytes))
	return err
}
