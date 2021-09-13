package common

import (
	"bufio"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/input"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/viper"
)

const (
	retryTimeout = 500 * time.Millisecond
)

func ParseTransactorAddrs(ts []string) ([]sdk.AccAddress, error) {
	var transactors []sdk.AccAddress
	for _, t := range ts {
		transactor, err := sdk.AccAddressFromBech32(t)
		if err != nil {
			return transactors, err
		}

		transactors = append(transactors, transactor)
	}

	return transactors, nil
}

func SetupUserPassword() error {
	buf := bufio.NewReader(os.Stdin)

	if viper.Get(FlagEthSignerPassphrase) == nil {
		pass, err := input.GetString("Enter eth keystore passphrase:", buf)
		if err != nil {
			return err
		}

		viper.Set(FlagEthSignerPassphrase, pass)
	}

	if viper.Get(FlagSgnPassphrase) == nil {
		pass, err := input.GetString("Enter sgn validator passphrase:", buf)
		if err != nil {
			return err
		}

		viper.Set(FlagSgnPassphrase, pass)
	}

	return nil
}

func RobustQuery(cliCtx client.Context, route string) ([]byte, error) {
	res, _, err := cliCtx.Query(route)
	if err != nil {
		time.Sleep(retryTimeout)
		res, _, err = cliCtx.Query(route)
		return res, err
	}

	return res, err
}

func RobustQueryWithData(cliCtx client.Context, route string, data []byte) ([]byte, error) {
	res, _, err := cliCtx.QueryWithData(route, data)
	if err != nil {
		time.Sleep(retryTimeout)
		res, _, err = cliCtx.QueryWithData(route, data)
		return res, err
	}

	return res, err
}

func VerifyAddressFormat(bz []byte) error {
	if len(bz) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "invalid address; cannot be empty")
	}
	if len(bz) != MaxAddrLen {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownAddress,
			"invalid address length; got: %d, max: %d", len(bz), MaxAddrLen,
		)
	}

	return nil
}
