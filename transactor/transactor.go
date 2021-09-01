package transactor

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/seal"
	vtypes "github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gammazero/deque"
	"github.com/spf13/viper"
)

const (
	maxTxRetry      = 15
	maxTxQueryRetry = 30
	txRetryDelay    = 1 * time.Second
	maxSignRetry    = 10
	signRetryDelay  = 100 * time.Millisecond
	maxGasRetry     = 5
)

var errGasCode = fmt.Errorf("code 11")

type Transactor struct {
	TxFactory  clienttx.Factory
	CliCtx     client.Context
	Key        keyring.Info
	passphrase string
	msgQueue   deque.Deque
}

func NewTransactor(cliHome, chainID, nodeURI, accAddr, passphrase string, cdc codec.Codec, legacyAmino *codec.LegacyAmino) (*Transactor, error) {
	reader := strings.NewReader(passphrase + "\n")
	kb, err := keyring.New(appName,
		viper.GetString(common.FlagSgnKeyringBackend), cliHome, reader)
	if err != nil {
		return nil, err
	}

	addr, err := vtypes.SdkAccAddrFromSgnBech32(accAddr)
	if err != nil {
		return nil, err
	}

	// may run into "resource temporarily unavailable" error if directly run it
	// retry when get this issue to avoid failure.
	var key keyring.Info
	for try := 0; try < maxSignRetry; try++ {
		key, err = kb.KeyByAddress(addr)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "resource temporarily unavailable") {
			log.Errorln("kb.GetByAddress error:", err)
			return nil, err
		}
		if try != maxSignRetry-1 {
			log.Debugln("retry kb.GetByAddress due to error:", err)
			time.Sleep(signRetryDelay)
		}
	}

	gasAdjustment := viper.GetFloat64(common.FlagSgnGasAdjustment)
	if gasAdjustment == 0 {
		gasAdjustment = common.DefaultSgnGasAdjustment
	}

	txConfig := tx.NewTxConfig((cdc).(*codec.ProtoCodec), tx.DefaultSignModes)
	cli, err := client.NewClientFromNode(nodeURI)
	if err != nil {
		log.Errorln("client.NewClientFromNode error:", err)
		return nil, err
	}
	cliCtx := client.Context{}.
		WithCodec(cdc).
		WithFromAddress(key.GetAddress()).
		WithFromName(key.GetName()).
		WithNodeURI(nodeURI).
		WithKeyring(kb).
		WithChainID(chainID).
		WithBroadcastMode(flags.BroadcastSync).
		WithTxConfig(txConfig).
		WithLegacyAmino(legacyAmino).
		WithClient(cli).
		WithAccountRetriever(types.AccountRetriever{})

	f := clienttx.Factory{}.
		WithKeybase(cliCtx.Keyring).
		WithTxConfig(cliCtx.TxConfig).
		WithAccountNumber(viper.GetUint64(flags.FlagAccountNumber)).
		WithSequence(viper.GetUint64(flags.FlagSequence)).
		WithGas(common.DefaultSgnGasLimit).
		WithGasAdjustment(gasAdjustment).
		WithChainID(chainID).
		WithMemo(viper.GetString(flags.FlagNote)).
		WithFees(viper.GetString(flags.FlagFees)).
		WithGasPrices(viper.GetString(flags.FlagGasPrices)).
		WithSignMode(signing.SignMode_SIGN_MODE_TEXTUAL).
		WithSimulateAndExecute(true)

	transactor := &Transactor{
		TxFactory:  f,
		CliCtx:     cliCtx,
		Key:        key,
		passphrase: passphrase,
	}

	return transactor, nil
}

func NewCliTransactor(cdc codec.Codec, cliHome string, legacyAmino *codec.LegacyAmino) (*Transactor, error) {
	return NewTransactor(
		cliHome,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		cdc,
		legacyAmino,
	)
}

func (t *Transactor) Run() {
	go t.start()
}

// AddTxMsg add msg into a queue before actual broadcast
func (t *Transactor) AddTxMsg(msg sdk.Msg) {
	t.msgQueue.PushBack(msg)
}

// Poll tx queue and send msgs in batch
func (t *Transactor) start() {
	for {
		if t.msgQueue.Len() == 0 {
			time.Sleep(time.Second)
			continue
		}

		t.drainTxMsgQueue()
	}
}

func (t *Transactor) drainTxMsgQueue() {
	var msgs []sdk.Msg
	for t.msgQueue.Len() != 0 {
		msg := t.msgQueue.PopFront().(sdk.Msg)
		msgs = append(msgs, msg)
	}

	t.sendTxMsgsWaitMined(msgs)
}

func (t *Transactor) sendTxMsgsWaitMined(msgs []sdk.Msg) (*sdk.TxResponse, error) {
	var txResponse *sdk.TxResponse
	var err error
	var gas uint64
	var retryNum int

	msgsId := uint32(time.Now().UnixNano() / 1000000)
	for {
		var retry bool
		logEntry := seal.NewTransactorLog(t.Key.GetAddress().String())
		logEntry.MsgNum = uint32(len(msgs))
		logEntry.MsgsId = msgsId
		txResponse, err = t.sendTxMsgs(msgs, gas)
		if txResponse != nil {
			logEntry.TxHash = txResponse.TxHash
		}
		if err != nil {
			logEntry.Error = append(logEntry.Error, err.Error())
			logEntry.Status = seal.TxMsgStatus_FAILED
			seal.CommitTransactorLog(logEntry)
			return txResponse, err
		}
		logEntry.Status = seal.TxMsgStatus_SENT
		seal.CommitTransactorLog(logEntry)

		// wait till transaction is mined or failed
		txResponse, err = t.waitMined(txResponse.TxHash)
		if txResponse != nil {
			logEntry.GasWanted = txResponse.GasWanted
			logEntry.GasUsed = txResponse.GasUsed
		}
		if err != nil {
			if errors.Is(err, errGasCode) && retryNum < maxGasRetry {
				gas = uint64(txResponse.GasUsed) * 2
				logEntry.Warn = append(logEntry.Warn, err.Error()+". retry...")
				retry = true
			} else {
				logEntry.Error = append(logEntry.Error, err.Error())
			}
			logEntry.Status = seal.TxMsgStatus_FAILED
		} else {
			logEntry.Status = seal.TxMsgStatus_SUCCEED
		}
		seal.CommitTransactorLog(logEntry)

		if retry {
			retryNum++
		} else {
			break
		}
	}
	return txResponse, err
}

func (t *Transactor) sendTxMsgs(msgs []sdk.Msg, gas uint64) (*sdk.TxResponse, error) {
	var txResponseErr error
	for try := 0; try < maxTxRetry; try++ {
		txBytes, err := t.buildAndSignTx(msgs, gas)
		if err != nil {
			return nil, fmt.Errorf("buildAndSignTx err: %w", err)
		}
		txResponse, err := t.CliCtx.BroadcastTx(txBytes)
		if err != nil {
			return nil, fmt.Errorf("BroadcastTx err: %w", err)
		}

		if txResponse.Code == sdkerrors.SuccessABCICode {
			return txResponse, nil
		}

		txResponseErr = fmt.Errorf("BroadcastTx failed with code: %d, rawLog: %s, acct: %s",
			txResponse.Code, txResponse.RawLog, t.Key.GetAddress())
		if txResponse.Code == sdkerrors.ErrUnauthorized.ABCICode() {
			log.Warnln(txResponseErr.Error(), "retrying")
			time.Sleep(txRetryDelay)
		} else {
			return txResponse, txResponseErr
		}
	}
	return nil, txResponseErr
}

func (t *Transactor) buildAndSignTx(msgs []sdk.Msg, gas uint64) ([]byte, error) {
	txf := t.TxFactory

	// txf, err := prepareFactory(t.CliCtx, txf)
	// if err != nil {
	// 	return nil, err
	// }

	if gas != 0 {
		txf = txf.WithGas(gas)
	} else if txf.SimulateAndExecute() || t.CliCtx.Simulate {
		_, adjusted, err := clienttx.CalculateGas(t.CliCtx, txf, msgs...)
		if err != nil {
			return nil, err
		}

		txf = txf.WithGas(adjusted)
	}

	tx, err := clienttx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return nil, err
	}

	tx.SetFeeGranter(t.CliCtx.GetFeeGranterAddress())

	err = clienttx.Sign(txf, t.CliCtx.GetFromName(), tx, true)
	if err != nil {
		return nil, err
	}

	return t.CliCtx.TxConfig.TxEncoder()(tx.GetTx())
}

func (t *Transactor) waitMined(txHash string) (*sdk.TxResponse, error) {
	var err error
	mined := false
	var txResponse *sdk.TxResponse
	for try := 0; try < maxTxQueryRetry; try++ {
		time.Sleep(txRetryDelay)
		if txResponse, err = tx.QueryTx(t.CliCtx, txHash); err == nil {
			mined = true
			break
		}
	}
	if !mined {
		return txResponse, fmt.Errorf("tx not mined, err: %w", err)
	} else if txResponse.Code != sdkerrors.SuccessABCICode {
		if txResponse.Code == 11 { // out of gas
			return txResponse, fmt.Errorf("tx failed with %w, %s", errGasCode, txResponse.RawLog)
		} else {
			return txResponse, fmt.Errorf("tx failed with code %d, %s", txResponse.Code, txResponse.RawLog)
		}
	}
	return txResponse, nil
}

func (t *Transactor) CliSendTxMsgWaitMined(msg sdk.Msg) {
	t.CliSendTxMsgsWaitMined([]sdk.Msg{msg})
}

func (t *Transactor) CliSendTxMsgsWaitMined(msgs []sdk.Msg) {
	res, _ := t.sendTxMsgsWaitMined(msgs)
	t.CliCtx.PrintProto(res)
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func prepareFactory(clientCtx client.Context, txf clienttx.Factory) (clienttx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}
