package ops

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BondValidatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bond",
		Short: "Bond validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			return BondValidator()
		},
	}
	return cmd
}

func BondValidator() error {
	ethClient, err := newEthClient( /*useSigner*/ true)
	if err != nil {
		return err
	}
	valAddr := eth.Hex2Addr(viper.GetString(common.FlagEthValidatorAddress))
	shouldBond, err := ethClient.Contracts.Viewer.ShouldBondValidator(&bind.CallOpts{}, valAddr)
	if err != nil {
		return fmt.Errorf("check if should bond validator err: %w", err)
	}
	if !shouldBond {
		log.Info("Validator not ready to be bonded")
		return nil
	}
	sgnAddr, err := ethClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, valAddr)
	if err != nil {
		return fmt.Errorf("get sgn addr err: %w", err)
	}
	acctAddress, err := sdk.AccAddressFromBech32(viper.GetString(common.FlagSgnValidatorAccount))
	if err != nil {
		return err
	}
	if !sdk.AccAddress(sgnAddr).Equals(acctAddress) {
		return fmt.Errorf("sgn addr not match %s %s", acctAddress, sgnAddr)
	}
	_, err = ethClient.Transactor.TransactWaitMined(
		"BondValidator",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return ethClient.Contracts.Staking.BondValidator(opts)
		},
	)
	if err != nil {
		return err
	}
	return nil
}
