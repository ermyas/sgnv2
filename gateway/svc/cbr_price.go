package gatewaysvc

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
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
	resp, err := cli.QueryChainTokensConfig(gs.TP.GetTransactor().CliCtx, &types.ChainTokensConfigRequest{})
	if err != nil {
		log.Errorln("we will use mocked chain tokens failed to load basic token info:", err)
		return
	}
	chainTokens := resp.GetChainTokens()
	tokenMap := make(map[string]bool)
	for _, tokens := range chainTokens {
		for _, token := range tokens.Tokens {
			tokenMap[token.GetSymbol()] = true
		}
	}
	symbol2chainIds, chainId2Symbol, err := dal.DB.GetAllChainAndGasToken()
	if err != nil {
		log.Errorln("failed to GetAllChainAndGasToken: err ", err)
		return
	}

	c := &types.CbrPrice{
		UpdateEpoch: uint64(time.Now().UnixNano() / 1000000),
		AssetPrice:  gs.PrepareAssetPrice(tokenMap, symbol2chainIds),
		GasPrice:    gs.PrepareGasPrice(chainId2Symbol),
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
	return
}

func (gs *GatewayService) PrepareGasPrice(chainId2Symbol map[uint64]string) (gp []*types.GasPrice) {
	for chainId := range chainId2Symbol {
		var price *big.Int
		var err error
		switch chainId {
		case 1:
			// Ethereum mainnet
			price = GetEthGasPrice()
		case 10:
			// Optimistic Ethereum
			price, err = gs.calcOptimismEffectiveGasPrice(1, 10)
			if err != nil {
				continue
			}
		case 69:
			// Optimistic Kovan
			// NOTE: Requires Kovan to be in the list of clients
			price, err = gs.calcOptimismEffectiveGasPrice(42, 69)
			if err != nil {
				continue
			}
		case 42161:
			// Arbitrum
			// SuggestGasPrice tends to overestimate gasPriceUsed by a factor of 2
			client := gs.EC[chainId]
			price, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Errorln("failed to SuggestGasPrice: chainId: ", chainId, ", error:", err)
				continue
			}
			price.Div(price, big.NewInt(2))
		default:
			client := gs.EC[chainId]
			price, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Errorln("failed to SuggestGasPrice: chainId: ", chainId, ", error:", err)
				continue
			}
		}
		gp = append(gp, &types.GasPrice{
			ChainId: chainId,
			Price:   price.String(),
		})
	}
	return gp
}

// calcOptimismEffectiveGasPrice calculates the effective gas price using the heuristic
// effectiveGasPrice = L1GasPrice / 14 + L2GasPrice
func (gs *GatewayService) calcOptimismEffectiveGasPrice(l1ChainId uint64, l2ChainId uint64) (*big.Int, error) {
	l1Client := gs.EC[l1ChainId]
	l1Price, err := l1Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorln("failed to SuggestGasPrice: chainId: ", l1ChainId, ", error:", err)
		return big.NewInt(0), err
	}
	l2Client := gs.EC[l2ChainId]
	l2Price, err := l2Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorln("failed to SuggestGasPrice: chainId: ", l2ChainId, ", error:", err)
		return big.NewInt(0), err
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

func GetEthGasPrice() *big.Int {
	qs := fmt.Sprintf(
		"confidenceLevels=99")
	client := resty.New()
	r, err := client.R().SetHeader("Authorization", viper.GetString(common.FlagBlockNativeApiKey)).
		SetQueryString(qs).SetResult(&BlockNativeResp{}).Get("https://api.blocknative.com/gasprices/blockprices")
	if err != nil || r.StatusCode() != 200 {
		log.Errorln("fail to get eth gas price from https://api.blocknative.com/gasprices/blockprices. ", err, r)
		return new(big.Int)
	}
	resp := r.Result().(*BlockNativeResp)
	if len(resp.BlockPrices) == 0 || len(resp.BlockPrices[0].EstimatedPrices) == 0 {
		log.Errorln("fail to get eth gas price from https://api.blocknative.com/gasprices/blockprices. ", r)
		return new(big.Int)
	}
	// blocknative return gas price in gwei
	return new(big.Int).Mul(big.NewInt(int64(resp.BlockPrices[0].EstimatedPrices[0].Price)), big.NewInt(1e9))
}
