package cli

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
)

const (
	flagCheckMainchain = "check-mainchain"
)

func GetQueryCmd() *cobra.Command {
	validatorQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the validator module",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	validatorQueryCmd.AddCommand(common.GetCommands(
		GetCmdValidator(types.StoreKey),
		GetCmdValidators(types.StoreKey),
		GetCmdDelegator(types.StoreKey),
		GetCmdDelegators(types.StoreKey),
		GetCmdSdkValidator(sdk_staking.StoreKey),
		GetCmdSdkValidators(sdk_staking.StoreKey),
		GetCmdSyncer(types.StoreKey),
		GetCmdQueryParams(types.StoreKey),
	)...)
	return validatorQueryCmd
}

// GetCmdSyncer queries syncer info
func GetCmdSyncer(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "syncer",
		Short: "query syncer info",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			cliCtx := common.NewQueryCLIContext(&clientCtx.Codec)
			syncer, err := QuerySyncer(cliCtx, queryRoute)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(&syncer)
		},
	}
}

// Query syncer info
func QuerySyncer(cliCtx client.Context, queryRoute string) (syncer types.Syncer, err error) {
	route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QuerySyncer)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.Codec.UnmarshalJSON(res, &syncer)
	return
}

// GetCmdDelegator queries request info
func GetCmdDelegator(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "delegator [validator-eth-addr] [delegator-eth-addr]",
		Short: "query delegator info by validator and delegator ETH addresses",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			cliCtx := common.NewQueryCLIContext(&clientCtx.Codec)
			delegator, err := QueryDelegator(cliCtx, queryRoute, args[0], args[1])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(&delegator)
		},
	}
}

func QueryDelegator(cliCtx client.Context, queryRoute, validatorAddress, delegatorAddress string) (delegator types.Delegator, err error) {
	return
}

// GetCmdValidator queries request info
func GetCmdValidator(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "validator [validator-eth-addr]",
		Short: "query validator info by validator ETH address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			cliCtx := common.NewQueryCLIContext(&clientCtx.Codec)
			validator, err := QueryValidator(cliCtx, queryRoute, args[0])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(&validator)
		},
	}
}

func QueryValidator(cliCtx client.Context, queryRoute, ethAddress string) (validator types.Validator, err error) {
	return
}

func GetCmdValidators(queryRoute string) *cobra.Command {
	return &cobra.Command{}
}

func QueryValidators(cliCtx client.Context, queryRoute string) (validators []types.Validator, err error) {
	return
}

// GetCmdDelegators queries request info
// TODO: support pagination
func GetCmdDelegators(queryRoute string) *cobra.Command {
	return &cobra.Command{}
}

func QueryDelegators(cliCtx client.Context, queryRoute, ethAddress string) (delegators []types.Delegator, err error) {
	return
}

// GetCmdValidator queries validator info
func GetCmdSdkValidator(queryRoute string) *cobra.Command {
	return &cobra.Command{}
}

// GetCmdValidator queries validator info
func GetCmdSdkValidators(queryRoute string) *cobra.Command {
	return &cobra.Command{}
}

// QuerySdkValidators is an interface for convenience to query (all) validators in staking module
func QuerySdkValidators(cliCtx client.Context, storeName string) (validators sdk_staking.Validators, err error) {
	return
}

// QueryBondedValidators is an interface for convenience to query bonded validators in staking module
func QueryBondedSdkValidators(cliCtx client.Context, storeName string) (validators sdk_staking.Validators, err error) {
	return
}

// addrStr should be bech32 sgn account address with prefix sgn
func QuerySdkValidator(cliCtx client.Context, storeName string, addrStr string) (validator sdk_staking.Validator, err error) {
	addr, err := sdk.AccAddressFromBech32(addrStr)
	if err != nil {
		log.Error(err)
		return
	}

	res, _, err := cliCtx.QueryStore(sdk_staking.GetValidatorKey(sdk.ValAddress(addr)), storeName)
	if err != nil {
		return
	}

	if len(res) == 0 {
		err = fmt.Errorf("%w for address %s", common.ErrRecordNotFound, addr)
		return
	}

	validator = sdk_staking.MustUnmarshalValidator(cliCtx.Codec, res)
	return
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams(queryRoute string) *cobra.Command {
	return &cobra.Command{}
}

// Query params info
func QueryParams(cliCtx client.Context, queryRoute string) (params types.Params, err error) {
	return
}

// ----------------------- CLI print-friendly output --------------------
