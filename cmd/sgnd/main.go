package main

import (
	"context"

	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/cmd"
	"github.com/celer-network/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = false
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
	config.Seal()

	// prepare and add flags
	encodingConfig := app.MakeEncodingConfig()
	executor := cmd.GetSgndExecutor(encodingConfig)
	srvCtx := server.NewDefaultContext()
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &client.Context{})
	ctx = context.WithValue(ctx, server.ServerContextKey, srvCtx)
	err := executor.ExecuteContext(ctx)
	if err != nil {
		panic(err)
	}
}
