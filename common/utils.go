package common

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/input"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/viper"
)

// Lengths of hashes and addresses in bytes.
const (
	retryTimeout        = 500 * time.Millisecond
	ERC20DenomSeparator = "/" // NOTE: Cosmos SDK only accepts "/" or "-"
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
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

func DeriveSdkAccAddressFromEthAddress(namespace string, ethAddr eth.Addr) sdk.AccAddress {
	return sdk.AccAddress(sdkaddress.Module(fmt.Sprintf("eth-%s", namespace), ethAddr.Bytes()))
}

// DeriveERC20TokenDenom generates denoms of the form symbol/chainId
func DeriveERC20TokenDenom(chainId uint64, symbol string) string {
	return fmt.Sprintf("%s%s%d", symbol, ERC20DenomSeparator, chainId)
}

func ParseERC20TokenDenom(denom string) (chainId uint64, symbol string, err error) {
	splitted := strings.Split(denom, ERC20DenomSeparator)
	if len(splitted) != 2 {
		return 0, "", fmt.Errorf("invalid denom %s", denom)
	}
	chainIdInt64, err := strconv.ParseInt(splitted[1], 10, 64)
	if err != nil {
		return 0, "", err
	}
	return uint64(chainIdInt64), splitted[0], nil
}

func TsMilli(t time.Time) uint64 {
	ts := uint64(t.UnixNano())
	return ts / uint64(time.Millisecond)
}

func TsMilliToTime(ts uint64) time.Time {
	return time.Unix(0, int64(ts*1000000))
}

func TsSecToTime(ts uint64) time.Time {
	return time.Unix(int64(ts), 0)
}

// return is ALWAYS >= 0 ie. unsigned
func Bytes2Int(in []byte) *big.Int {
	return new(big.Int).SetBytes(in)
}

func Str2BigInt(str string) *big.Int {
	b := new(big.Int)
	b.SetString(str, 10)
	return b
}

func IsNegative(in *big.Int) bool {
	return in.Sign() == -1
}

// IsValidTxHash verifies whether a string can represent a valid hash or not.
func IsValidTxHash(txHash string) bool {
	if txHash == "" {
		return false
	}
	if has0xPrefix(txHash) {
		txHash = txHash[2:]
	}
	return len(txHash) == 2*HashLength && isHex(txHash)
}

// IsHexAddress verifies whether a string can represent a valid hex-encoded
// Ethereum address or not.
func IsHexAddress(s string) bool {
	if has0xPrefix(s) {
		s = s[2:]
	}
	return len(s) == 2*AddressLength && isHex(s)
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
