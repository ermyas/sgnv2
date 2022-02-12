package ops

import (
	"fmt"

	"github.com/celer-network/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/client/flags"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pvm "github.com/tendermint/tendermint/privval"
)

func ValidatorAddressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "Query validator addresses",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			fmt.Printf("Addresses of the validator configured in %s:\n\n", home)
			fmt.Printf("validator eth address: %s\n\n", viper.GetString(common.FlagEthValidatorAddress))
			ethClient, err := newEthClient( /*useSigner*/ true)
			if err != nil {
				return err
			}
			fmt.Printf("signer eth address: %x\n\n", ethClient.Address)
			sgnAddr := viper.GetString(common.FlagSgnValidatorAccount)
			acctAddress, err := sdk.AccAddressFromBech32(sgnAddr)
			if err != nil {
				return err
			}
			fmt.Printf("sgn acct address in bech32: %s\n", sgnAddr)
			fmt.Printf("sgn acct address in hex: 0x%x\n\n", acctAddress.Bytes())

			tmCfg := server.GetServerContextFromCmd(cmd).Config
			privValidator := pvm.LoadFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile())
			tmPubKey, err := privValidator.GetPubKey()
			if err != nil {
				return err
			}
			sdkPubKey, err := cryptocodec.FromTmPubKeyInterface(tmPubKey)
			if err != nil {
				return err
			}
			consAddr := sdk.GetConsAddress(sdkPubKey)
			fmt.Printf("sgn consensus address: %s\n\n", consAddr)
			return nil
		},
	}
	return cmd
}
