package ops

import (
	"context"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagChain = "chain"
	flagToken = "token"
	flagDry   = "dry"
)

func SyncFarmingCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync-farming",
		Short: "Syncs all LP liquidity for a token on a chain with their stakes in the farming pool",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd)
		},
	}
	cmd.Flags().Uint32(flagChain, 0, "chain id")
	cmd.Flags().String(flagToken, "", "token address")
	cmd.Flags().Bool(flagDry, false, "run without sending message to sgn chain")
	cmd.MarkFlagRequired(flagChain)
	cmd.MarkFlagRequired(flagToken)
	return cmd
}

func run(cmd *cobra.Command) error {
	chid := viper.GetUint64(flagChain)
	token := viper.GetString(flagToken)

	cliCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return fmt.Errorf("GetClientQueryContext %v", err)
	}
	home, err := cmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return fmt.Errorf("get home flag %v", err)
	}
	cbrcli := types.NewQueryClient(cliCtx)
	req := &types.QueryLPsRequest{
		ChainId:   chid,
		TokenAddr: token,
	}
	res, err := cbrcli.QueryLPs(context.Background(), req)
	if err != nil {
		return fmt.Errorf("QueryLPs %v, clictx %v", err, cliCtx)
	}
	txr, err := transactor.NewCliTransactor(home, cliCtx.LegacyAmino, cliCtx.Codec, cliCtx.InterfaceRegistry)
	if err != nil {
		return fmt.Errorf("NewCliTransactor %v", err)
	}
	log.Infof("sync-farming: chain %d, token %s", chid, token)
	for _, addr := range res.GetLps() {
		msg := &types.MsgSyncFarming{
			LpAddress:    addr,
			ChainId:      chid,
			TokenAddress: token,
			Creator:      viper.GetString(common.FlagSgnValidatorAccount),
		}
		if !viper.GetBool(flagDry) {
			txr.AddTxMsg(msg)
		}
		log.Infof("LP %s", addr)
	}

	return nil
}
