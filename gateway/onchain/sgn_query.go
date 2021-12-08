package onchain

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app/params"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/spf13/viper"
)

var SGNTransactors *transactor.TransactorPool

func InitSGNTransactors(home string, encoding params.EncodingConfig) {
	txrAddrs := viper.GetStringSlice(common.FlagSgnTransactors)
	chainId := viper.GetString(common.FlagSgnChainId)
	nodeUri := viper.GetString(common.FlagSgnNodeURI)
	log.Infof("Initializing sgn transactors with args: home %s, chainId %s, nodeuri %s, addrs %v", home, chainId, nodeUri, txrAddrs)
	SGNTransactors = transactor.NewTransactorPool(home, chainId, encoding.Amino, encoding.Codec, encoding.InterfaceRegistry)
	err := SGNTransactors.AddTransactors(
		nodeUri,
		viper.GetString(common.FlagSgnPassphrase),
		txrAddrs)
	if err != nil {
		log.Fatalf("failed to add transactors: %s", err.Error())
	}
	log.Infof("Initialized %d sgn transactors", len(txrAddrs))
}

func getFeePerc(srcChainId, dstChainId uint64, tokenSymbol string) uint32 {
	perc := uint32(0)
	tr := SGNTransactors.GetTransactor()
	if tr != nil {
		_perc, err := cbrcli.QueryFeePerc(tr.CliCtx, &cbrtypes.GetFeePercentageRequest{
			Symbol:     tokenSymbol,
			SrcChainId: srcChainId,
			DstChainId: dstChainId,
		})
		if _perc == nil || err != nil {
			log.Warnf("get fee perc failed, srcChainId:%d, dsChainId:%d, will record 0 in db", srcChainId, dstChainId)
		} else {
			perc = _perc.FeePerc
		}
	}
	return perc
}

func getEstimatedAmt(srcChainId, dstChainId uint64, srcToken *webapi.TokenInfo, amt string) (string, error) {
	if !utils.IsValidAmt(amt) {
		return "0", fmt.Errorf("invalid amt, params checking failed")
	}
	tr := SGNTransactors.GetTransactor()

	getFeeRequest := &cbrtypes.GetFeeRequest{
		SrcChainId:   srcChainId,
		DstChainId:   dstChainId,
		SrcTokenAddr: srcToken.Token.GetAddress(),
		Amt:          amt,
	}
	feeInfo, err := cbrcli.QueryFee(tr.CliCtx, getFeeRequest)
	if err != nil {
		log.Warnf("cli.QueryFee error, srcChainId:%d, dstChainId:%d, srcTokenAddr:%s, amt:%s, err:%+v", srcChainId, dstChainId, srcToken.Token.GetAddress(), amt, err)
		return "0", err
	}
	if feeInfo == nil {
		return "0", fmt.Errorf("can not estimate fee")
	}
	eqValueTokenAmt := feeInfo.GetEqValueTokenAmt()
	percFee := feeInfo.GetPercFee()
	baseFee := feeInfo.GetBaseFee()
	feeAmt := new(big.Int).Add(common.Str2BigInt(percFee), common.Str2BigInt(baseFee))
	estimateReceivedAmt := new(big.Int).Sub(common.Str2BigInt(eqValueTokenAmt), feeAmt)
	if estimateReceivedAmt.Cmp(new(big.Int).SetInt64(0)) < 0 {
		return "0", fmt.Errorf("got invalid estimateReceivedAmt:%s,eqValueTokenAmt:%s, percFee:%s, baseFee:%s, use 0 instead", eqValueTokenAmt, percFee, baseFee, estimateReceivedAmt.String())
	}
	return estimateReceivedAmt.String(), nil
}
