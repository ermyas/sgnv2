package ops

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagMinSelfDelegation = "min-self-delegation"
	FlagCommissionRate    = "commission-rate"
)

func InitValidatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a validator (if validator keystore is on local machine)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return InitValidator()
		},
	}
	cmd.Flags().String(FlagKeystore, "", "Validator keystore file")
	cmd.Flags().String(FlagPassphrase, "", "Validator keystore passphrase")
	cmd.Flags().String(FlagMinSelfDelegation, "", "Minimum self-delegated stake (integer in unit of CELR)")
	cmd.Flags().Float64(FlagCommissionRate, 0, "Commission rate (e.g., 0.1 is 10%)")

	cmd.MarkFlagRequired(FlagKeystore)
	cmd.MarkFlagRequired(FlagMinSelfDelegation)
	cmd.MarkFlagRequired(FlagCommissionRate)
	return cmd
}

func InitValidator() error {
	ethClient, err := newEthClient( /*useSigner*/ false)
	if err != nil {
		return err
	}
	minSelfDelegation := calcRawAmount(viper.GetString(FlagMinSelfDelegation))
	commissionRate := viper.GetFloat64(FlagCommissionRate)

	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	_, signerAddr, err := eth.CreateSigner(signerKey, signerPass, big.NewInt(0))
	if err != nil {
		log.Fatalln("CreateSigner err:", err)
	}

	stakingContract := ethClient.Contracts.Staking
	info, err := stakingContract.Validators(&bind.CallOpts{}, ethClient.Address)
	if err != nil {
		return err
	}
	if info.Status == 0 {
		err = approveCelr(ethClient, ethClient.Contracts.Staking.Address, minSelfDelegation)
		if err != nil {
			return err
		}
		log.Infof(
			"Initializing validator %x with signer %x minSelfDelegation: %s, commissionRate: %f",
			ethClient.Address,
			signerAddr,
			minSelfDelegation,
			commissionRate,
		)
		_, err = ethClient.Transactor.TransactWaitMined(
			"InitializeValidator",
			func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
				return stakingContract.InitializeValidator(
					opts, signerAddr, minSelfDelegation, eth.CommissionRate(commissionRate))
			},
		)
		if err != nil {
			return err
		}
	}

	acctAddress, err := sdk.AccAddressFromBech32(viper.GetString(common.FlagSgnValidatorAccount))
	if err != nil {
		return err
	}
	log.Infof("Calling updateSgnAddr for %s", acctAddress)
	_, err = ethClient.Transactor.TransactWaitMined(
		"UpdateSgnAddr",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return ethClient.Contracts.Sgn.UpdateSgnAddr(opts, acctAddress.Bytes())
		},
	)
	if err != nil {
		return err
	}
	return nil
}
