package impl

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagData = "data"
)

func SignMessageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-message",
		Short: "Signs an Ethereum message",
		RunE: func(cmd *cobra.Command, args []string) error {
			return SignMessage()
		},
	}
	cmd.Flags().String(FlagData, "", "Hex data to sign")

	cmd.MarkFlagRequired(FlagData)

	viper.BindPFlag(FlagData, cmd.Flags().Lookup(FlagData))

	return cmd
}

func SignMessage() error {
	signer, err := getKmsSigner()
	if err != nil {
		return err
	}
	sig, err := signer.SignEthMessage(eth.Hex2Bytes(viper.GetString(FlagData)))
	if err != nil {
		return err
	}
	log.Infoln("sig:", eth.Bytes2Hex(sig))
	return nil
}
