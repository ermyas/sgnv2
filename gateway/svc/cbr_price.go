package gatewaysvc

import (
	"context"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/celer-network/sgn-v2/eth"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	"github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
)

func (gs *GatewayService) StartUpdateTokenPricePolling(interval time.Duration) {
	if viper.GetString(common.FlagGatewayAwsKey) == "" {
		log.Warn("aws key not configured, no price upload")
		return
	}
	go func() {
		// let fee model run 90 sec upfront
		time.Sleep(90 * time.Second)
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 5 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			gs.UpdateTokenPrice2S3()
		}
	}()
	time.Sleep(2 * time.Second)
}

func (gs *GatewayService) UpdateTokenPrice2S3() {
	resp, err := cli.QueryChainTokensConfig(onchain.SGNTransactors.GetTransactor().CliCtx, &types.ChainTokensConfigRequest{})
	if err != nil {
		log.Warnln("we will use cached chain tokens failed to load basic token info:", err)
		return
	}
	chainTokens := resp.GetChainTokens()
	tokenMap := make(map[string]bool)
	for _, tokens := range chainTokens {
		for _, token := range tokens.Tokens {
			tokenMap[token.GetSymbol()] = true
		}
	}
	symbol2chainIds, chainId2Symbol, chainId2DropGas, chainId2SuggestedBaseFee, err := dal.DB.GetAllChainAndGasToken()
	if err != nil {
		log.Errorln("failed to GetAllChainAndGasToken: err ", err)
		return
	}

	c := &types.CbrPrice{
		UpdateEpoch: uint64(time.Now().UnixNano() / 1000000),
		AssetPrice:  gs.PrepareAssetPrice(tokenMap, symbol2chainIds),
		GasPrice:    gs.PrepareGasPrice(chainId2Symbol, chainId2DropGas, chainId2SuggestedBaseFee),
	}

	marshaler := jsonpb.Marshaler{}
	m, err := marshaler.MarshalToString(c)
	if err != nil {
		log.Errorln("failed to UploadFile: err ", err)
		return
	}
	UploadFile(m)
}

func UploadFile(content string) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(viper.GetString(common.FlagGatewayAwsKey), viper.GetString(common.FlagGatewayAwsSecret), ""),
		Region:      aws.String(viper.GetString(common.FlagGatewayAwsS3Region))},
	)
	parse, err := url.Parse(viper.GetString(common.FlagSgnPriceUpdateUrl))
	if err != nil {
		log.Errorln("fail to parse viper config FlagSgnPriceUpdateUrl,", viper.GetString(common.FlagSgnPriceUpdateUrl), err)
		return
	}

	s3up := s3manager.NewUploader(sess)
	poi := &s3manager.UploadInput{
		ACL:    aws.String("public-read"), // so every one can ready the file
		Bucket: aws.String(viper.GetString(common.FlagGatewayAwsS3Bucket)),
		Key:    aws.String(parse.RequestURI()),
		Body:   strings.NewReader(content),
	}
	poi.ContentType = aws.String("application/json")
	_, err = s3up.Upload(poi)
	if err != nil {
		log.Errorln("fail to UploadFile,", err)
		return
	}
	log.Infoln("success upload cbr price file, ", content)
	return
}

func (gs *GatewayService) PrepareGasPrice(chainId2Symbol map[uint64]string, chainId2DropGas map[uint64]string,
	chainId2SuggestedBaseFee map[uint64]float64) (gp []*types.GasPrice) {
	for chainId, symbol := range chainId2Symbol {
		var price *big.Int
		var err error
		switch chainId {
		case 10, 69:
			// Optimistic
			price, err = gs.calcOptimismGasPrice(chainId)
			if err != nil {
				log.Errorln("failed to calcOptimismGasPrice: chainId: ", chainId, ", error:", err)
				continue
			}
		case 42161:
			// Arbitrum
			// SuggestGasPrice tends to overestimate gasPriceUsed by a factor of 2
			client := gs.Chains.GetEthClient(chainId)
			price, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Errorln("failed to SuggestGasPrice: chainId: ", chainId, ", error:", err)
				continue
			}
			price.Div(price, big.NewInt(2))
		default:
			client := gs.Chains.GetEthClient(chainId)
			price, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Errorln("failed to SuggestGasPrice: chainId: ", chainId, ", error:", err)
				continue
			}
		}
		gp = append(gp, &types.GasPrice{
			ChainId: chainId,
			Price:   gs.calcGasPriceShouldRaiseDueToGasDrop(chainId, symbol, price, chainId2DropGas, chainId2SuggestedBaseFee).String(),
		})
	}
	return gp
}

func (gs *GatewayService) calcGasPriceShouldRaiseDueToGasDrop(chainId uint64, gasTokenSymbol string, originalGasPrice *big.Int,
	chainId2DropGas map[uint64]string, chainId2SuggestedBaseFee map[uint64]float64) *big.Int {
	gasTokenPrice, err := gs.F.GetUsdPrice(gasTokenSymbol)
	if err != nil {
		log.Errorln("fail to GetUsdPrice,", err)
		return originalGasPrice
	}
	droppedGasTokenAmtStr := chainId2DropGas[chainId]
	droppedGasTokenAmt, b := big.NewFloat(0).SetString(droppedGasTokenAmtStr)
	if !b || droppedGasTokenAmt.Cmp(big.NewFloat(0)) <= 0 {
		return originalGasPrice
	}
	droppedGasTokenAmt.Quo(droppedGasTokenAmt, big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
	droppedGasTokenAmt.Mul(droppedGasTokenAmt, big.NewFloat(gasTokenPrice))
	originalBaseFee := big.NewFloat(chainId2SuggestedBaseFee[chainId])
	// quo = originalBaseFee / originalGasPrice
	quo := big.NewFloat(0).Quo(originalBaseFee, big.NewFloat(0).SetInt(originalGasPrice))
	// add = originalBaseFee + droppedGas
	add := big.NewFloat(0).Add(originalBaseFee, droppedGasTokenAmt)
	// manipulatedGasPrice = (originalBaseFee + droppedGas) / originalBaseFee * originalGasPrice
	manipulatedGasPrice := big.NewFloat(0).Quo(add, quo)
	u, _ := manipulatedGasPrice.Uint64()
	newGasPrice := big.NewInt(0).SetUint64(u)
	log.Infoln("raise ", chainId, " gas price due to gas drop on arrival, before:", originalGasPrice.String(), ", after:", newGasPrice.String())
	return newGasPrice
}

// calcOptimismEffectiveGasPrice calculates the effective gas price using the heuristic
// effectiveGasPrice = L1GasPrice / 14 + L2GasPrice
func (gs *GatewayService) calcOptimismGasPrice(chainId uint64) (*big.Int, error) {
	caller, err := eth.NewOVMGasPriceOracleCaller(eth.Hex2Addr("0x420000000000000000000000000000000000000F"), gs.Chains.GetEthClient(chainId))
	if err != nil {
		return nil, err
	}
	l1Price, err := caller.L1BaseFee(nil)
	if err != nil {
		return nil, err
	}
	l2Price, err := caller.GasPrice(nil)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Add(new(big.Int).Div(l1Price, big.NewInt(14)), l2Price), nil
}

func (gs *GatewayService) PrepareAssetPrice(tokenMap map[string]bool, symbol2chainIds map[string][]uint64) (ap []*types.AssetPrice) {
	// for token can be transferred
	for symbol := range tokenMap {
		price, err := gs.F.GetUsdPrice(symbol)
		if err != nil {
			log.Errorln("fail to GetUsdPrice,", err)
			price = 0
		}
		var chainIds []uint64
		if symbol2chainIds[symbol] != nil {
			chainIds = symbol2chainIds[symbol]
		}
		ap = append(ap, &types.AssetPrice{
			Symbol:   symbol,
			ChainIds: chainIds,
			Price:    uint32(price * 1e4),
		})
	}
	// for token which is only native gas token but not used for transferred, e.g. BNB
	for symbol, chainIds := range symbol2chainIds {
		if tokenMap[symbol] == false {
			price, err := gs.F.GetUsdPrice(symbol)
			if err != nil {
				log.Errorln("fail to GetUsdPrice, token:", symbol, ", error:", err)
				price = 0
			}
			ap = append(ap, &types.AssetPrice{
				Symbol:   symbol,
				ChainIds: chainIds,
				Price:    uint32(price * 1e4),
			})
		}
	}
	return ap
}

type BlockNativeResp struct {
	BlockPrices []BlockPrices `json:"blockPrices"`
}

type BlockPrices struct {
	EstimatedPrices []EstimatedPrices `json:"estimatedPrices"`
}

type EstimatedPrices struct {
	Price uint64 `json:"price"`
}
