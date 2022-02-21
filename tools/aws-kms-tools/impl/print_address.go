package impl

import (
	"fmt"

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
	fmt.Println("Ethereum address:", signer.Addr)
	return nil
}
