package utils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

// just to satisfy monitor interface requirement
type cbrContract struct {
	*eth.Bridge
	Address eth.Addr
}

// ethclient etc
type CbrOneChain struct {
	*ethclient.Client
	*ethutils.Transactor
	contract *cbrContract
}

func newOneChain(chainId uint64) (*CbrOneChain, error) {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Fatalln("fail to load multichain configs err:", err)
	}
	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	for _, cfg := range mcc {
		if cfg.ChainID == chainId {
			ec, err := ethclient.Dial(cfg.Gateway)
			if err != nil {
				log.Fatalln("dial", cfg.Gateway, "err:", err)
			}
			chid, err := ec.ChainID(context.Background())
			if err != nil {
				log.Fatalf("get chainid %d err: %s", cfg.ChainID, err)
			}
			if chid.Uint64() != cfg.ChainID {
				log.Fatalf("chainid mismatch! cfg has %d but onchain has %d", cfg.ChainID, chid.Uint64())
			}
			cbr, err := eth.NewBridge(eth.Hex2Addr(cfg.CBridge), ec)
			if err != nil {
				log.Fatalln("cbridge contract at", cfg.CBridge, "err:", err)
			}
			signer, addr, err := eth.CreateSigner(signerKey, signerPass, chid)
			if err != nil {
				log.Fatalln("CreateSigner err:", err)
			}
			transactor := ethutils.NewTransactorByExternalSigner(
				addr,
				signer,
				ec,
				big.NewInt(int64(cfg.ChainID)),
				ethutils.WithBlockDelay(cfg.BlkDelay),
				ethutils.WithPollingInterval(time.Duration(cfg.BlkInterval)*time.Second),
				ethutils.WithAddGasEstimateRatio(cfg.AddGasEstimateRatio),
				ethutils.WithAddGasGwei(cfg.AddGasGwei),
				ethutils.WithMaxFeePerGasGwei(cfg.MaxFeePerGasGwei),
			)
			if err != nil {
				log.Fatalln("NewTransactor err:", err)
			}
			c := &CbrOneChain{
				Client:     ec,
				Transactor: transactor,
				contract: &cbrContract{
					Bridge:  cbr,
					Address: eth.Hex2Addr(cfg.CBridge),
				},
			}
			return c, nil
		}
	}

	return nil, fmt.Errorf("chainId %d not exist", chainId)
}

func setupCbr(chainid uint64, txhash string) (cbr *CbrOneChain, txReceipt *ethtypes.Receipt, err error) {
	cbr, err = newOneChain(chainid)
	if err != nil {
		log.Fatal("newOneChain err:", err)
	}
	txReceipt, err = cbr.TransactionReceipt(context.Background(), eth.Hex2Hash(txhash))
	if err != nil {
		log.Errorln("TransactionReceipt err:", err)
		return
	}
	return
}

func GetELog(chainid uint64, txhash, evname string) (*ethtypes.Log, error) {
	cbr, txReceipt, err := setupCbr(chainid, txhash)
	if err != nil {
		return nil, err
	}

	elog := eth.FindMatchCbrEvent(evname, cbr.contract.Address, txReceipt.Logs)

	if elog == nil {
		log.Errorln("no match event found in tx:", txhash)
		return nil, fmt.Errorf("no match event found in tx: %s", txhash)
	}
	return elog, nil
}
