package ops

import (
	"encoding/json"
	"fmt"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
)

func EthViewerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "view",
		Short:                      "Operation subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		ViewDelegatorCommand(),
		ViewValidatorsCommand(),
		ViewBondedValidatorsCommand(),
	)

	return cmd
}

func ViewDelegatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator [delegator-eth-addr]",
		Short: "Get delegator info from eth chain",
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
		Short: "Get all validators info from eth chain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getValidators(false)
		},
	}
	return cmd
}

func ViewBondedValidatorsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bonded-validators",
		Short: "Get all bonded validators info from eth chain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getValidators(true)
		},
	}
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
