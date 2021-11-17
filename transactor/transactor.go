package transactor

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/seal"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gammazero/deque"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

const (
	maxTxRetry         = 15
	maxTxQueryRetry    = 30
	txRetryDelay       = 1 * time.Second
	maxSignRetry       = 10
	signRetryDelay     = 100 * time.Millisecond
	maxWaitMinedRetry  = 5
	maxRawMsgBytesInTx = 500000
)

var errGasCode = fmt.Errorf("code %d", sdkerrors.ErrOutOfGas.ABCICode())
var errSeqCode = fmt.Errorf("code %d", sdkerrors.ErrWrongSequence.ABCICode())

type Transactor struct {
	TxFactory  clienttx.Factory
	CliCtx     client.Context
	Key        keyring.Info
	passphrase string
	msgQueue   deque.Deque
	lock       sync.Mutex
}

func NewTransactor(
	homeDir, chainID, nodeURI, accAddr, passphrase string,
	legacyAmino *codec.LegacyAmino,
	cdc codec.Codec,
	interfaceRegistry codectypes.InterfaceRegistry,
) (*Transactor, error) {
	reader := strings.NewReader(passphrase + "\n")
	kb, err := keyring.New(appName,
		viper.GetString(common.FlagSgnKeyringBackend), homeDir, reader)
	if err != nil {
		return nil, err
	}

	addr, err := sdk.AccAddressFromBech32(accAddr)
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
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithInterfaceRegistry(interfaceRegistry).
		WithClient(cli).
		WithHomeDir(homeDir)

	f := clienttx.Factory{}.
		WithKeybase(cliCtx.Keyring).
		WithTxConfig(cliCtx.TxConfig).
		WithAccountRetriever(cliCtx.AccountRetriever).
		WithAccountNumber(viper.GetUint64(flags.FlagAccountNumber)).
		WithSequence(viper.GetUint64(flags.FlagSequence)).
		WithGas(common.DefaultSgnGasLimit).
		WithGasAdjustment(gasAdjustment).
		WithChainID(chainID).
		WithMemo(viper.GetString(flags.FlagNote)).
		WithFees(viper.GetString(flags.FlagFees)).
		WithGasPrices(viper.GetString(flags.FlagGasPrices)).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithSimulateAndExecute(true)

	transactor := &Transactor{
		TxFactory:  f,
		CliCtx:     cliCtx,
		Key:        key,
		passphrase: passphrase,
	}

	return transactor, nil
}

func NewCliTransactor(homeDir string, legacyAmino *codec.LegacyAmino, cdc codec.Codec, interfaceRegistry codectypes.InterfaceRegistry) (*Transactor, error) {
	return NewTransactor(
		homeDir,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		legacyAmino,
		cdc,
		interfaceRegistry,
	)
}

func (t *Transactor) Run() {
	go t.start()
}

// AddTxMsg add msg into a queue before actual broadcast
func (t *Transactor) AddTxMsg(msg sdk.Msg) {
	t.checkSigner([]sdk.Msg{msg})
	t.msgQueue.PushBack(msg)
}

// Poll tx queue and send msgs in batch
func (t *Transactor) start() {
	for {
		if t.msgQueue.Len() == 0 {
			time.Sleep(time.Second)
			continue
		}

		t.consumeTxMsgQueue()
	}
}

func (t *Transactor) consumeTxMsgQueue() {
	msgs := make([]sdk.Msg, 0)
	var msgsBytesLen int
	var msgType string
	for t.msgQueue.Len() != 0 {
		msg := t.msgQueue.PopFront().(sdk.Msg)

		msgBytes, _ := proto.Marshal(msg)
		msgsBytesLen += len(msgBytes)
		if msgsBytesLen > maxRawMsgBytesInTx {
			if len(msgs) != 0 {
				t.msgQueue.PushFront(msg) // adds back to the queue, if it's not the first msg, otherwise, drop the msg as single one cannot be processed
			} else { // for first msg, try to split the msg into smaller ones
				switch m := msg.(type) {
				case *synctypes.MsgProposeUpdates:
					tmp := *m //copy one
					for {
						half := len(tmp.Updates) / 2
						if half == 0 {
							break
						}
						tmp.Updates = tmp.Updates[:half] //get half
						tmpMsgBytes, _ := proto.Marshal(&tmp)
						if len(tmpMsgBytes) <= maxRawMsgBytesInTx {
							msgs = append(msgs, &tmp)
							m.Updates = m.Updates[half:]
							t.msgQueue.PushFront(m) // push back the left
							break
						}
					}
				case *synctypes.MsgVoteUpdates:
					tmp := *m //copy one
					for {
						half := len(tmp.Votes) / 2
						if half == 0 {
							break
						}
						tmp.Votes = tmp.Votes[:half] //get half
						tmpMsgBytes, _ := proto.Marshal(&tmp)
						if len(tmpMsgBytes) <= maxRawMsgBytesInTx {
							msgs = append(msgs, &tmp)
							m.Votes = m.Votes[half:]
							t.msgQueue.PushFront(m) // push back the left
							break
						}
					}
				default:
					msgType = reflect.TypeOf(msg).String()
				}
			}
			break
		}

		msgs = append(msgs, msg)
	}

	if len(msgs) > 0 {
		t.SendTxMsgsWaitMined(msgs)
		log.Debugln("Current msgs count in one tx:", len(msgs))
	} else {
		log.Errorf("Single msg too large, msg type is: %s!", msgType)
	}
}

func (t *Transactor) SendTxMsgsWaitMined(msgs []sdk.Msg) (*sdk.TxResponse, error) {
	t.checkSigner(msgs)

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
			//logEntry.Msgs = fmt.Sprintf("%s", msgs)
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
			if errors.Is(err, errGasCode) && retryNum < maxWaitMinedRetry {
				gas = uint64(txResponse.GasUsed) * 2
				logEntry.Warn = append(logEntry.Warn, err.Error()+". will retry...")
				retry = true
			} else if errors.Is(err, errSeqCode) && retryNum < maxWaitMinedRetry {
				logEntry.Warn = append(logEntry.Warn, err.Error()+". will retry...")
				retry = true
			} else {
				//logEntry.Msgs = fmt.Sprintf("%s", msgs)
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
		log.Debugln("tx msg bytes size:", len(txBytes))
		txResponse, err := t.CliCtx.BroadcastTx(txBytes)
		if err != nil {
			return nil, fmt.Errorf("BroadcastTx err: %w", err)
		}

		if txResponse.Code == sdkerrors.SuccessABCICode {
			return txResponse, nil
		}

		txResponseErr = fmt.Errorf("BroadcastTx failed with code: %d, rawLog: %s, acct: %s",
			txResponse.Code, txResponse.RawLog, t.Key.GetAddress())
		if txResponse.Code == sdkerrors.ErrWrongSequence.ABCICode() && try < maxTxRetry-1 {
			log.Debugln(txResponseErr.Error(), "will retry")
			time.Sleep(txRetryDelay)
		} else {
			return txResponse, txResponseErr
		}
	}
	return nil, txResponseErr
}

// send a single msg so one fail won't affect others. this is only intended
// for initwithdraw/signagain request, if x/cbr return err, we return it immediately (wrapped in sendTxMsgs fmt.Errorf)
// if nil err, caller should query later. No waitmine
// note due to inherent async of estimategas and actual include in block, it's
// possible even this returns nil err, x/cbr still fails
// lock to ensure req are serialized even gateway handle concurrent initwithdraw from clients
func (t *Transactor) LockSendTx(msg sdk.Msg) (*sdk.TxResponse, error) {
	t.checkSigner([]sdk.Msg{msg})
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.sendTxMsgs([]sdk.Msg{msg}, 0) // 0 gas so estimate will be called
}

func (t *Transactor) buildAndSignTx(msgs []sdk.Msg, gas uint64) ([]byte, error) {
	txf := t.TxFactory
	txf, err := prepareFactory(t.CliCtx, txf)
	if err != nil {
		return nil, err
	}

	if gas != 0 {
		txf = txf.WithGas(gas)
	} else if txf.SimulateAndExecute() || t.CliCtx.Simulate {
		for try := 0; try < maxTxRetry; try++ {
			_, adjusted, err := clienttx.CalculateGas(t.CliCtx, txf, msgs...)
			if err != nil {
				if strings.Contains(err.Error(), "account sequence mismatch") && try < maxTxRetry-1 {
					log.Debugln(err, "increment seq and retry")
					txf = txf.WithSequence(txf.Sequence() + 1)
					continue
				}
				return nil, fmt.Errorf("CalculateGas err: %w", err)
			}
			txf = txf.WithGas(adjusted)
			break
		}
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
		if txResponse.Code == sdkerrors.ErrOutOfGas.ABCICode() { // out of gas
			return txResponse, fmt.Errorf("tx failed with %w, %s", errGasCode, txResponse.RawLog)
		} else if txResponse.Code == sdkerrors.ErrWrongSequence.ABCICode() {
			return txResponse, fmt.Errorf("tx failed with %w, %s", errSeqCode, txResponse.RawLog)
		} else {
			return txResponse, fmt.Errorf("tx failed with code %d, %s", txResponse.Code, txResponse.RawLog)
		}
	}
	return txResponse, nil
}

func (t *Transactor) CliSendTxMsgWaitMined(msg sdk.Msg) {
	t.checkSigner([]sdk.Msg{msg})
	t.cliSendTxMsgsWaitMined([]sdk.Msg{msg})
}

func (t *Transactor) cliSendTxMsgsWaitMined(msgs []sdk.Msg) {
	res, err := t.SendTxMsgsWaitMined(msgs)
	t.CliCtx.OutputFormat = "text"
	if err == nil {
		t.CliCtx.PrintProto(res)
	} else {
		t.CliCtx.PrintString(err.Error())
	}
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

func (t *Transactor) checkSigner(msgs []sdk.Msg) {
	for _, msg := range msgs {
		if t.Key.GetAddress().String() != msg.GetSigners()[0].String() {
			log.Fatal("tx msg signer is not the transactor")
		}
	}
}
