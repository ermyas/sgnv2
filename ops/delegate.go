package ops

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func DelegateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate",
		Short: "Delegate tokens to a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Delegate()
		},
	}
	cmd.Flags().String(FlagKeystore, "", "Delegator keystore file")
	cmd.Flags().String(FlagPassphrase, "", "Delegator keystore passphrase")
	cmd.Flags().String(FlagValidator, "", "Validator ETH address")
	cmd.Flags().String(FlagAmount, "", "Delegate amount (integer in unit of CELR)")

	cmd.MarkFlagRequired(FlagKeystore)
	cmd.MarkFlagRequired(FlagValidator)
	cmd.MarkFlagRequired(FlagAmount)
	return cmd
}

func UndelegateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "undelegate",
		Short: "Undelegate tokens from a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Undelegate()
		},
	}
	cmd.Flags().String(FlagKeystore, "", "Delegator keystore file")
	cmd.Flags().String(FlagPassphrase, "", "Delegator keystore passphrase")
	cmd.Flags().String(FlagValidator, "", "Validator ETH address")
	cmd.Flags().String(FlagAmount, "", "Undelegate amount (integer in unit of CELR)")

	cmd.MarkFlagRequired(FlagKeystore)
	cmd.MarkFlagRequired(FlagValidator)
	cmd.MarkFlagRequired(FlagAmount)
	return cmd
}

func CompleteUndelegateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete-undelegate",
		Short: "Complete undelegate tokens from a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			return CompleteUndelegate()
		},
	}
	cmd.Flags().String(FlagKeystore, "", "Delegator keystore file")
	cmd.Flags().String(FlagPassphrase, "", "Delegator keystore passphrase")
	cmd.Flags().String(FlagValidator, "", "Validator ETH address")

	cmd.MarkFlagRequired(FlagKeystore)
	cmd.MarkFlagRequired(FlagValidator)
	return cmd
}

func Delegate() error {
	ethClient, err := newEthClient( /*useSigner*/ false)
	if err != nil {
		return err
	}
	amount := calcRawAmount(viper.GetString(FlagAmount))
	validator := eth.Hex2Addr(viper.GetString(FlagValidator))

	err = approveCelr(ethClient, ethClient.Contracts.Staking.Address, amount)
	if err != nil {
		return err
	}
	log.Infof("Delegating to validator %x with amount: %s", validator, amount)
	_, err = ethClient.Transactor.TransactWaitMined(
		"Delegate",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return ethClient.Contracts.Staking.Delegate(opts, validator, amount)
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func Undelegate() error {
	ethClient, err := newEthClient( /*useSigner*/ false)
	if err != nil {
		return err
	}
	amount := calcRawAmount(viper.GetString(FlagAmount))
	validator := eth.Hex2Addr(viper.GetString(FlagValidator))

	log.Infof("Undelegating from validator %x with amount: %s", validator, amount)
	_, err = ethClient.Transactor.TransactWaitMined(
		"Undelegate",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return ethClient.Contracts.Staking.UndelegateTokens(opts, validator, amount)
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func CompleteUndelegate() error {
	ethClient, err := newEthClient( /*useSigner*/ false)
	if err != nil {
		return err
	}
	validator := eth.Hex2Addr(viper.GetString(FlagValidator))

	log.Infof("Completing undelegate from validator %x", validator)
	_, err = ethClient.Transactor.TransactWaitMined(
		"CompleteUndelegate",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return ethClient.Contracts.Staking.CompleteUndelegate(opts, validator)
		},
	)
	if err != nil {
		return err
	}
	return nil
}
