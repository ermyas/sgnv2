package impl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func AddGenesisValidatorCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-validator [key_name] [amount]",
		Short: "Add a genesis validator to genesis.json",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			_, valPubKey, err := genutil.InitializeNodeValidatorFiles(serverCtx.Config)
			if err != nil {
				return fmt.Errorf("failed to initialize node validator files, %w", err)
			}
			pkAny, err := codectypes.NewAnyWithValue(valPubKey)
			if err != nil {
				return fmt.Errorf("failed to generate pkAny, %w", err)
			}

			name := args[0]
			key, err := clientCtx.Keyring.Key(name)
			if err != nil {
				return fmt.Errorf("failed to fetch '%s' from the keyring, %w", name, err)
			}
			inBuf := bufio.NewReader(cmd.InOrStdin())
			clientCtx = clientCtx.WithInput(inBuf).WithFromAddress(key.GetAddress())

			tokens, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return err
			}
			signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
			_, addr, err := eth.CreateSigner(signerKey, signerPass, big.NewInt(0))
			if err != nil {
				log.Fatalln("CreateSigner err:", err)
			}

			initialValidator := stakingtypes.Validator{
				EthAddress:      viper.GetString(common.FlagEthValidatorAddress),
				EthSigner:       eth.Addr2Hex(addr),
				SgnAddress:      key.GetAddress().String(),
				ConsensusPubkey: pkAny,
				Status:          stakingtypes.Bonded,
				Tokens:          tokens,
				DelegatorShares: tokens,
			}
			log.Infoln("validator: ", initialValidator.String())

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			stakingGenState := stakingtypes.GetGenesisStateFromAppState(cdc, appState)
			stakingGenState.Validators = append(stakingGenState.Validators, initialValidator)

			stakingGenStateBz, err := cdc.MarshalJSON(stakingGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal staking genesis state: %w", err)
			}
			appState[stakingtypes.ModuleName] = stakingGenStateBz

			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
