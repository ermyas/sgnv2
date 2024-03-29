package transactor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/go-bip39"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	mnemonicEntropySize = 256

	passphraseFlag  = "passphrase"
	namePrefixFlag  = "prefix"
	countFlag       = "count"
	genesisCoinFlag = "coin"
	appName         = "sgn"
)

func AccountsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "Add accounts in batch",
		RunE: func(cmd *cobra.Command, args []string) error {
			addresses, err := addAccounts(cmd)
			if err != nil {
				return err
			}

			jsonString, err := keys.MarshalJSON(addresses)
			if err != nil {
				return err
			}
			log.Infoln("All addresses", string(jsonString))

			genesisCoin := viper.GetString(genesisCoinFlag)
			if genesisCoin != "" {
				for _, address := range addresses {
					cmd := exec.Command("sgnd", "add-genesis-account", address, genesisCoin)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Dir, _ = filepath.Abs(".")
					if err = cmd.Run(); err == nil {
						log.Infof("Add address %s to genesis", address)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().String(passphraseFlag, "12345678", "account passphrase")
	cmd.Flags().String(namePrefixFlag, "transactor", "account prefix")
	cmd.Flags().Int(countFlag, 1, "account count")
	cmd.Flags().String(genesisCoinFlag, "", "amount of coin adding to genesis for the account")
	cmd.Flags().String(flags.FlagKeyringBackend, cKeys.BackendFile, "Select keyring's backend (os|file|test)")
	return cmd
}

func addAccounts(cmd *cobra.Command) ([]string, error) {
	var addresses []string

	passphrase := viper.GetString(passphraseFlag)
	np := viper.GetString(namePrefixFlag)
	count := viper.GetInt(countFlag)
	kb, err := cKeys.New(appName,
		viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), strings.NewReader(passphrase))
	if err != nil {
		return addresses, err
	}

	for i := 0; i < count; i++ {
		name := fmt.Sprintf("%s_%d", np, i)
		info, err := kb.Key(name)
		if err == nil {
			log.Infof("Account %s has existed", name)
			printAccount(info)
			addresses = append(addresses, info.GetAddress().String())
			continue
		}

		entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
		if err != nil {
			return addresses, err
		}

		mnemonic, err := bip39.NewMnemonic(entropySeed[:])
		if err != nil {
			return addresses, err
		}

		info, err = kb.NewAccount(name, mnemonic, passphrase, "", hd.Secp256k1)
		if err != nil {
			return addresses, err
		}

		log.Infof("Account %s created", name)
		printAccount(info)
		addresses = append(addresses, info.GetAddress().String())
	}

	return addresses, nil
}

func printAccount(info cKeys.Info) {
	out, err := cKeys.MkAccKeyOutput(info)
	if err != nil {
		return
	}

	jsonString, err := keys.MarshalJSON(out)
	if err != nil {
		return
	}

	log.Infof(string(jsonString))
}
