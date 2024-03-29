package impl

import (
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagDestination      = "destination"
	FlagDestinationShort = "d"
	FlagValue            = "value"
	FlagValueShort       = "v"
	FlagNonce            = "nonce"
	FlagNonceShort       = "n"
	FlagGasLimit         = "gaslimit"
	FlagGasPrice         = "gasprice"
	FlagSendAll          = "sendall"
	FlagMinBalance       = "minbalance"
)

func SendTxCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-tx",
		Short: "Sends a transaction",
		RunE: func(cmd *cobra.Command, args []string) error {
			return SendTx()
		},
	}
	cmd.Flags().StringP(FlagDestination, FlagDestinationShort, "", "Destination to send the tx")

	chainIds := make([]int, 0, len(commonChainIdRpcs))
	for chainId := range commonChainIdRpcs {
		chainIds = append(chainIds, int(chainId))
	}
	sort.Ints(chainIds)
	cmd.Flags().String(FlagRpc, "", fmt.Sprintf("JSON-RPC endpoint to use, optional if chain ID is in: %v", chainIds))
	cmd.Flags().Uint64P(FlagChainId, FlagChainIdShort, 0, "chain ID")
	cmd.Flags().Int64P(FlagNonce, FlagNonceShort, 0, "Nonce override, if -1 use account nonce")
	cmd.Flags().Uint64(FlagGasPrice, 0, "Gas price override, NOTE: indicates maxPriorityFeePerGas if chain supports EIP-1559")
	cmd.Flags().Uint64(FlagGasLimit, 21000, "Gas limit override")
	cmd.Flags().Uint64(FlagValue, 0, "Native gas value to send with the tx")
	cmd.Flags().String(FlagData, "", "Hex data to send with the tx")
	cmd.Flags().Bool(FlagSendAll, false, "Whether to send the entire balance minus gas fee. Overrides value if set to true")
	cmd.Flags().Uint64(FlagMinBalance, 0, "Minimal balance required on the address")

	viper.BindPFlag(FlagDestination, cmd.Flags().Lookup(FlagDestination))
	viper.BindPFlag(FlagRpc, cmd.Flags().Lookup(FlagRpc))
	viper.BindPFlag(FlagChainId, cmd.Flags().Lookup(FlagChainId))
	viper.BindPFlag(FlagNonce, cmd.Flags().Lookup(FlagNonce))
	viper.BindPFlag(FlagGasPrice, cmd.Flags().Lookup(FlagGasPrice))
	viper.BindPFlag(FlagGasLimit, cmd.Flags().Lookup(FlagGasLimit))
	viper.BindPFlag(FlagValue, cmd.Flags().Lookup(FlagValue))
	viper.BindPFlag(FlagData, cmd.Flags().Lookup(FlagData))
	viper.BindPFlag(FlagSendAll, cmd.Flags().Lookup(FlagSendAll))
	viper.BindPFlag(FlagMinBalance, cmd.Flags().Lookup(FlagMinBalance))

	return cmd
}

func SendTx() error {
	rpc := viper.GetString(FlagRpc)
	if rpc == "" {
		rpc = getCfgRpc(viper.GetUint64(FlagChainId))
		if rpc != "" {
			log.Debugln("use cfg rpc", rpc)
		}
	}
	if rpc == "" {
		rpc = commonChainIdRpcs[viper.GetUint64(FlagChainId)]
		if rpc != "" {
			log.Debugln("use common rpc", rpc)
		}
	}
	ec, err := ethclient.Dial(rpc)
	if err != nil {
		return fmt.Errorf("dial err %w", err)
	}
	signer, err := getKmsSigner()
	if err != nil {
		return fmt.Errorf("getKmsSigner err %w", err)
	}
	balance, err := ec.BalanceAt(bgCtx, signer.Addr, nil)
	if err != nil {
		return fmt.Errorf("BalanceAt err %w", err)
	}
	log.Infoln("Balance:", balance)
	minBalance := new(big.Int).SetUint64(viper.GetUint64(FlagMinBalance))
	if balance.Cmp(minBalance) <= 0 {
		return errors.New("skip sending due to balance less than min")
	}
	destination := signer.Addr // sent to self by default
	if viper.GetString(FlagDestination) != "" {
		destination = eth.Hex2Addr(viper.GetString(FlagDestination))
	}
	// Now build tx and send
	return sendTx(ec, signer.Addr, destination, balance, signer.SignerFn)
}

func sendTx(ec *ethclient.Client, from, to eth.Addr, bal *big.Int, signer bind.SignerFn) error {
	var rawTx *types.Transaction
	head, err := ec.HeaderByNumber(bgCtx, nil)
	if err != nil {
		return fmt.Errorf("HeaderByNumber err %w", err)
	}
	var nonce uint64
	nonceFlag := viper.GetInt64(FlagNonce)
	if nonceFlag > 0 {
		nonce = uint64(nonceFlag)
	} else if nonceFlag == -1 {
		nonce, err = ec.NonceAt(bgCtx, from, nil)
		if err != nil {
			return fmt.Errorf("NonceAt %w", err)
		}
	} else {
		nonce, err = ec.PendingNonceAt(bgCtx, from)
		if err != nil {
			return fmt.Errorf("PendingNonceAt err %w", err)
		}
	}
	var gasPrice *big.Int
	gasPriceFlag := viper.GetUint64(FlagGasPrice)
	if gasPriceFlag != 0 {
		gasPrice = new(big.Int).SetUint64(gasPriceFlag * 1e9)
	} else {
		gasPrice, err = ec.SuggestGasPrice(bgCtx)
		if err != nil {
			return fmt.Errorf("SuggestGasPrice err %w", err)
		}
		log.Infoln("suggested gas price:", gasPrice.Uint64()/1e9)
	}
	var value *big.Int
	sendAll := viper.GetBool(FlagSendAll)
	valueFlag := viper.GetUint64(FlagValue)
	gasLimit := viper.GetUint64(FlagGasLimit)
	var data []byte
	dataFlag := viper.GetString(FlagData)
	if dataFlag != "" {
		data = eth.Hex2Bytes(dataFlag)
	}
	if head.BaseFee != nil {
		// EIP-1559, new dynamic tx, per spec we should do:
		// maxPriorityFeePerGas: eth_gasPrice - base_fee or just use the eth_maxPriorityFeePerGas rpc
		// maxFeePerGas: maxPriorityFeePerGas + 2 * base_fee = eth_gasPrice + base_fee
		// note if we calculate sendamt based on maxFeePerGas, it will leave one base_fee*gas residual
		// assume maxPriorityFee is way smaller than base fee, we could do following:
		// GasTipCap := eth_maxPriorityFeePerGas and GasFeeCap := eth_gasPrice + GasTipCap
		// but the risk is if eth becomes busy, our tx may pending for a long time. as here our gas is only 21K, we are ok w/ base_fee*gas residual
		gasFeeCap := new(big.Int).Add(gasPrice, head.BaseFee)
		gasCost := new(big.Int).Mul(gasFeeCap, new(big.Int).SetUint64(gasLimit))
		if sendAll {
			// NOTE: for EIP-1559, it's possible we still has some left b/c we only set cap
			value = new(big.Int).Sub(bal, gasCost)
		} else {
			value = new(big.Int).SetUint64(valueFlag)
		}
		rawTx = types.NewTx(&types.DynamicFeeTx{
			Nonce:     nonce,
			To:        &to,
			Gas:       gasLimit,
			GasTipCap: new(big.Int).Sub(gasPrice, head.BaseFee),
			GasFeeCap: gasFeeCap,
			Value:     value,
			Data:      data,
		})

	} else {
		if sendAll {
			value = new(big.Int).Sub(bal, new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit)))
		} else {
			value = new(big.Int).SetUint64(valueFlag)
		}
		rawTx = types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       &to,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Value:    value,
			Data:     data,
		})
	}
	signedTx, err := signer(from, rawTx)
	if err != nil {
		return fmt.Errorf("SignerFn err %w", err)
	}
	log.Infoln("txHash:", signedTx.Hash())
	err = ec.SendTransaction(bgCtx, signedTx)
	if err != nil {
		return fmt.Errorf("SendTransaction err %w", err)
	}
	return nil
}
