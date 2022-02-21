package impl

import (
	"context"
	"fmt"
	"sort"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagAddr = "addr"
)

func PrintNonceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print-nonce",
		Short: "Prints the pending and latest nonces of the key",
		RunE: func(cmd *cobra.Command, args []string) error {
			chainId, err := cmd.Flags().GetUint64(FlagChainId)
			if err != nil {
				return err
			}
			rpc, err := cmd.Flags().GetString(FlagRpc)
			if err != nil {
				return err
			}
			if rpc == "" {
				rpc = commonChainIdRpcs[chainId]
			}
			ec, err := ethclient.Dial(rpc)
			if err != nil {
				return fmt.Errorf("dial rpc %s err %w", rpc, err)
			}

			addrstr, err := cmd.Flags().GetString(FlagAddr)
			if err != nil {
				return err
			}
			addr := eth.Hex2Addr(addrstr)
			if addr == eth.ZeroAddr {
				signer, err := getKmsSigner()
				if err != nil {
					return err
				}
				addr = signer.Addr
				fmt.Println("KMS signer eth address:", addr)
			}

			nonce, err := ec.NonceAt(context.Background(), addr, nil)
			if err != nil {
				return fmt.Errorf("NonceAt %w", err)
			}
			fmt.Println("Latest nonce:", nonce)

			pendingNonce, err := ec.PendingNonceAt(context.Background(), addr)
			if err != nil {
				return fmt.Errorf("PendingNonceAt %w", err)
			}
			fmt.Println("Pending nonce:", pendingNonce)

			return nil
		},
	}
	chainIds := make([]int, 0, len(commonChainIdRpcs))
	for chainId := range commonChainIdRpcs {
		chainIds = append(chainIds, int(chainId))
	}
	sort.Ints(chainIds)
	cmd.Flags().String(FlagRpc, "", fmt.Sprintf("JSON-RPC endpoint to use, optional if chain ID is in: %v", chainIds))
	cmd.Flags().Uint64P(FlagChainId, FlagChainIdShort, 0, "Optional chain ID")
	cmd.Flags().String(FlagAddr, "", "Optional eth address (if provided, no need for kms flags)")

	viper.BindPFlag(FlagRpc, cmd.Flags().Lookup(FlagRpc))
	viper.BindPFlag(FlagChainId, cmd.Flags().Lookup(FlagChainId))

	return cmd
}
