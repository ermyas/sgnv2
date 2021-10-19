package common

import (
	"bufio"
	"encoding/hex"
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
	ec "github.com/ethereum/go-ethereum/common"
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

func DeriveSdkAccAddressFromEthAddress(namespace string, ethAddr eth.Addr) sdk.AccAddress {
	return sdk.AccAddress(sdkaddress.Module(fmt.Sprintf("eth-%s", namespace), ethAddr.Bytes()))
}

func TsMilliToTime(ms uint64) time.Time {
	sec := int64(ms / 1000)
	nsec := int64((ms % 1000) * 1000000)
	return time.Unix(sec, nsec).UTC()
}

func TsMilli(t time.Time) uint64 {
	ts := uint64(t.UnixNano())
	return ts / uint64(time.Millisecond)
}

func TsToTime(ts uint64) time.Time {
	return time.Unix(0, int64(ts*1000000))
}

func TsSecToTime(ts uint64) time.Time {
	return time.Unix(int64(ts), 0)
}

type Hash = ec.Hash
type Addr = ec.Address

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

// ========== Hex/Bytes ==========

// Hex2Bytes supports hex string with or without 0x prefix
// Calls hex.DecodeString directly and ignore err
// similar to ec.FromHex but better
func Hex2Bytes(s string) (b []byte) {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	// hex.DecodeString expects an even-length string
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, _ = hex.DecodeString(s)
	return b
}

// Bytes2Hex returns hex string without 0x prefix
func Bytes2Hex(b []byte) string {
	return hex.EncodeToString(b)
}

// ========== Address ==========

// Hex2Addr accepts hex string with or without 0x prefix and return Addr
func Hex2Addr(s string) Addr {
	return ec.BytesToAddress(Hex2Bytes(s))
}

// Addr2Hex returns hex without 0x
func Addr2Hex(a Addr) string {
	return Bytes2Hex(a[:])
}

// Bytes2Addr returns Address from b
// Addr.Bytes() does the reverse
func Bytes2Addr(b []byte) Addr {
	return ec.BytesToAddress(b)
}

func Bytes2Hash(b []byte) Hash {
	return ec.BytesToHash(b)
}

func GetSymbolFromFarmingToken(token string) string {
	return strings.Replace(token, "CB-", "", 1)
}

func IsValidTxHash(txHash string) bool {
	if txHash == "" {
		return false
	}
	num, err := strconv.Atoi(txHash)
	notNum := num == 0 && err != nil
	preFix0x := strings.HasPrefix(txHash, "0x")
	return notNum || preFix0x
}
