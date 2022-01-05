package impl

import (
	"github.com/celer-network/goutils/log"

	"github.com/spf13/cobra"
)

func PrintAddressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print-address",
		Short: "Prints the Ethereum address of the key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return PrintAddress()
		},
	}
	return cmd
}

func PrintAddress() error {
	signer, err := getKmsSigner()
	if err != nil {
		return err
	}
	log.Infoln("Ethereum address:", signer.Addr)
	return nil
}
