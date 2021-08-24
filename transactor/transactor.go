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
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
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

var gasErrCode = fmt.Errorf("code 11")

type Transactor struct {
	TxBuilder  client.TxBuilder
	CliCtx     client.Context
	Key        keyring.Info
	passphrase string
	msgQueue   deque.Deque
	gpe        *GasPriceEstimator
}

func NewTransactor(cliHome, chainID, nodeURI, accAddr, passphrase string, cdc codec.Codec, gpe *GasPriceEstimator) (*Transactor, error) {
	kb, err := keyring.New(appName,
		viper.GetString(common.FlagSgnKeyringBackend), cliHome, strings.NewReader(passphrase))
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

	fees, err := sdk.ParseCoinsNormalized(viper.GetString(flags.FlagFees))
	if err != nil {
		panic(err)
	}

	txConfig := tx.NewTxConfig((cdc).(*codec.ProtoCodec), tx.DefaultSignModes)
	cliCtx := client.Context{}.
		WithCodec(cdc).
		WithFromAddress(key.GetAddress()).
		WithFromName(key.GetName()).
		WithNodeURI(nodeURI).
		WithKeyring(kb).
		WithChainID(chainID).
		WithBroadcastMode(flags.BroadcastSync).
		WithTxConfig(txConfig)

	txBldr := cliCtx.TxConfig.NewTxBuilder()
	txBldr.SetFeeAmount(fees)
	txBldr.SetGasLimit(common.DefaultSgnGasLimit)

	transactor := &Transactor{
		TxBuilder:  txBldr,
		CliCtx:     cliCtx,
		Key:        key,
		passphrase: passphrase,
		gpe:        gpe,
	}

	return transactor, nil
}

func NewCliTransactor(cdc codec.Codec, cliHome string) (*Transactor, error) {
	return NewTransactor(
		cliHome,
		viper.GetString(common.FlagSgnChainID),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		cdc,
		nil,
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
		// for _, msg := range msgs {
		// 	logEntry.MsgType[msg.Type()] = logEntry.MsgType[msg.Type()] + 1
		// }
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
		//logEntry.GasWanted = int64(stdSignMsg.Fee.Gas)
		seal.CommitTransactorLog(logEntry)

		// wait till transaction is mined or failed
		txResponse, err = t.waitMined(txResponse.TxHash)
		if txResponse != nil {
			logEntry.GasWanted = txResponse.GasWanted
			logEntry.GasUsed = txResponse.GasUsed
		}
		if err != nil {
			if errors.Is(err, gasErrCode) && retryNum < maxGasRetry {
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
	// var txResponseErr error
	// for try := 0; try < maxTxRetry; try++ {
	// 	txBytes, stdSignMsg, err := t.buildAndSignTx(msgs, gas)
	// 	if err != nil {
	// 		return nil, nil, fmt.Errorf("buildAndSignTx err: %w", err)
	// 	}
	// 	txResponse, err := t.CliCtx.BroadcastTx(txBytes)
	// 	if err != nil {
	// 		return nil, stdSignMsg, fmt.Errorf("BroadcastTx err: %w", err)
	// 	}

	// 	if txResponse.Code == sdkerrors.SuccessABCICode {
	// 		return &txResponse, stdSignMsg, nil
	// 	}

	// 	txResponseErr = fmt.Errorf("BroadcastTx failed with code: %d, rawLog: %s, stdSignMsg chainId: %s acct: %s accnum: %d seq: %d",
	// 		txResponse.Code, txResponse.RawLog, stdSignMsg.ChainID, t.Key.GetAddress(), stdSignMsg.AccountNumber, stdSignMsg.Sequence)
	// 	if txResponse.Code == sdkerrors.ErrUnauthorized.ABCICode() {
	// 		log.Warnln(txResponseErr.Error(), "retrying")
	// 		time.Sleep(txRetryDelay)
	// 	} else {
	// 		return &txResponse, stdSignMsg, txResponseErr
	// 	}
	// }
	return nil, nil
}

// func (t *Transactor) buildAndSignTx(msgs []sdk.Msg, gas uint64) ([]byte, *types.StdSignMsg, error) {
// 	if t.gpe != nil {
// 		t.TxBuilder = t.TxBuilder.WithGasPrices(t.gpe.GetGasPrice())
// 	}

// 	txBldr, err := utils.PrepareTxBuilder(t.TxBuilder, t.CliCtx)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("PrepareTxBuilder err: %w", err)
// 	}

// 	if gas != 0 {
// 		txBldr = txBldr.WithGas(gas)
// 	} else if txBldr.SimulateAndExecute() || t.CliCtx.Simulate {
// 		txBldr, err = utils.EnrichWithGas(txBldr, t.CliCtx, msgs)
// 		if err != nil {
// 			return nil, nil, fmt.Errorf("EnrichWithGas err: %w", err)
// 		}
// 	}

// 	var txBytes []byte
// 	var stdSignMsg types.StdSignMsg
// 	for try := 0; try < maxSignRetry; try++ {
// 		stdSignMsg, err = txBldr.BuildSignMsg(msgs)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		txBytes, err = txBldr.Sign(t.Key.GetName(), t.passphrase, stdSignMsg)
// 		if err == nil {
// 			return txBytes, &stdSignMsg, nil
// 		}
// 		if !strings.Contains(err.Error(), "resource temporarily unavailable") {
// 			break
// 		}
// 		if try != maxSignRetry-1 {
// 			log.Debugln("Failed to call txBldr.BuildAndSign. Will retry it.")
// 			time.Sleep(signRetryDelay)
// 		}
// 	}
// 	return nil, nil, nil
// }

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
			return txResponse, fmt.Errorf("tx failed with %w, %s", gasErrCode, txResponse.RawLog)
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
