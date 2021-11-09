package cli

import (
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/slash/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	slashQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the slash module",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	slashQueryCmd.AddCommand(common.GetCommands(
		GetCmdSlash(),
		GetCmdSlashes(),
		GetCmdSlashRequest(),
		GetCmdQueryParams(),
	)...)
	return slashQueryCmd
}

// GetCmdSlash queries slash info
func GetCmdSlash() *cobra.Command {
	return &cobra.Command{
		Use:   "slash [nonce]",
		Short: "query slash info by nonce",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			nonce, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			slash, err := QuerySlash(cliCtx, nonce)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(&slash)
		},
	}
}

// Query slash info
func QuerySlash(cliCtx client.Context, nonce uint64) (slash types.Slash, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(types.NewQuerySlashParams(nonce))
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySlash)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &slash)
	return
}

// GetCmdSlashes queries slash info
func GetCmdSlashes() *cobra.Command {
	return &cobra.Command{
		Use:   "slashes",
		Short: "query slashes info",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			slashes, err := QuerySlashes(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintObjectLegacy(slashes)
		},
	}
}

// Query slashes info
func QuerySlashes(cliCtx client.Context) (slashes []types.Slash, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySlashes)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &slashes)
	return
}

// GetCmdSlashRequest queries slash request proto
func GetCmdSlashRequest() *cobra.Command {
	return &cobra.Command{
		Use:   "slash-request [nonce]",
		Short: "query slash request proto by nonce",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			nonce, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			slashRequest, err := QuerySlashRequest(cliCtx, nonce)
			if err != nil {
				return err
			}

			return cliCtx.PrintBytes(slashRequest)
		},
	}
}

// Query slash info
func QuerySlashRequest(cliCtx client.Context, nonce uint64) (slashRequest []byte, err error) {
	slash, err := QuerySlash(cliCtx, nonce)
	if err != nil {
		return
	}

	slashRequest = slash.GetEthSlashBytes()
	return
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current slash parameters information",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			params, err := QueryParams(cliCtx)
			if err != nil {
				return err
			}

			return cliCtx.PrintProto(params)
		},
	}
}

func QueryParams(cliCtx client.Context) (*types.Params, error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParameters)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return nil, err
	}

	params := new(types.Params)
	cliCtx.LegacyAmino.MustUnmarshalJSON(res, params)

	return params, nil
}
