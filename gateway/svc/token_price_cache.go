package gatewaysvc

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/lthibault/jitterbug"
	"gopkg.in/resty.v1"
)

const priceApiUrl = "https://api.coingecko.com/api/v3/simple/price"

// VsTokenPrices vsToken -> price
type VsTokenPrices map[string]float64

type TokenPrices map[string]VsTokenPrices

type TokenPriceCache struct {
	vsTokenIds []string
	Prices     map[string]float64 // do not access this map with token symbol since its key is coingecko's tokenId
}

// NewTokenPriceCache builds a new instance of TokenPriceCache with an empty Prices map.
// Note you still have to manually start polling loop using StartTokenPricePolling()
func NewTokenPriceCache(tr *transactor.Transactor) *TokenPriceCache {
	var vsTokenIds = []string{"usd"}

	feeSvr := &TokenPriceCache{
		vsTokenIds: vsTokenIds,
		Prices:     make(map[string]float64),
	}
	feeSvr.StartTokenPricePolling(tr, 1*time.Minute)
	log.Infof("token price cached")
	return feeSvr
}

// StartTokenPricePolling starts a loop with the given interval and 3s stdev for polling price
func (t *TokenPriceCache) StartTokenPricePolling(tr *transactor.Transactor, interval time.Duration) {
	go func() {
		time.Sleep(3 * time.Second)
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {

			err := t.refreshCache(tr)
			if err != nil {
				log.Errorln("failed to refresh token price cache:", err)
			} else {
				// also update price cache in dal
				if dal.DB.AllTokenIds == nil || len(dal.DB.AllTokenIds) == 0 {
					newAllTokenIds := make(map[string]*dal.TokenData)
					tokenIds, err := dal.DB.GetAllTokenIds()
					if err != nil {
						for _, tokenInfo := range tokenIds {
							newAllTokenIds[tokenInfo.Symbol] = &dal.TokenData{
								Id:     tokenInfo.Id,
								Symbol: tokenInfo.Symbol,
							}
						}
					}

					dal.DB.AllTokenIds = newAllTokenIds
				}
				dal.DB.Prices = t.Prices
			}
		}
	}()
	time.Sleep(2 * time.Second)
}

// GetUsdPrice gets the token/USD price by token symbol. e.g. "ETH", "DAI", "USDT"
func (t *TokenPriceCache) GetUsdPrice(tokenSymbol string) (float64, error) {
	if tokenSymbol == "WETH" {
		// will always use ETH price
		tokenSymbol = "ETH"
	}
	tokenId := dal.DB.GetTokenIdBySymbol(tokenSymbol)
	if tokenId == "" {
		price, mocked := dal.GetMockedPrice(tokenSymbol) // try to use mocked price if token not found
		if mocked {
			return price, nil
		} else {
			return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
		}
	}
	price, ok := t.Prices[tokenId]
	if !ok {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
	}
	return price, nil
}

// GetUsdPrice gets the token/USD price by token symbol. e.g. "ETH", "DAI", "USDT"
func (t *TokenPriceCache) GetUsdVolume(token *cbrtypes.Token, amt *big.Int) float64 {
	tokenPrize, err := t.GetUsdPrice(token.GetSymbol())
	if err != nil {
		return 0
	}
	tokenAmt, _ := new(big.Float).Quo(new(big.Float).SetInt(amt), big.NewFloat(math.Pow(10, float64(token.GetDecimal())))).Float64()
	return tokenAmt * tokenPrize
}

func (t *TokenPriceCache) GetTokenPrice(token *cbrtypes.Token, chainToken *cbrtypes.Token, chainTokenAmt *big.Int) (*big.Int, error) {
	tokenPrize, err := t.GetUsdPrice(token.GetSymbol())
	if err != nil {
		return big.NewInt(0), err
	}
	chainTokenPrice, err := t.GetUsdPrice(chainToken.GetSymbol())
	if err != nil {
		return big.NewInt(0), err
	}
	priceRate := new(big.Float).Mul(big.NewFloat(chainTokenPrice/tokenPrize), new(big.Float).SetInt(chainTokenAmt))
	tokenAmount := new(big.Float).Mul(priceRate, big.NewFloat(math.Pow(10, float64(token.GetDecimal()))/math.Pow(10, float64(chainToken.GetDecimal()))))
	ret := new(big.Int)
	tokenAmount.Int(ret)
	log.Info("special log, tokenPrize:", tokenPrize, " chainTokenPrice:", chainTokenPrice, " ret:", ret)
	return ret, nil
}

func (t *TokenPriceCache) refreshCache(tr *transactor.Transactor) error {
	resp, err := cli.QueryChainTokensConfig(tr.CliCtx, &cbrtypes.ChainTokensConfigRequest{})
	if err != nil {
		log.Errorln("we will use mocked chain tokens failed to load basic token info:", err)
	}
	chainTokens := resp.GetChainTokens()
	tokenMap := make(map[string]uint)
	for _, tokens := range chainTokens {
		for _, token := range tokens.Tokens {
			tokenMap[token.GetSymbol()] = 1
		}
	}

	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	farmingPools, err := queryClient.Pools(
		ctx,
		&farmingtypes.QueryPoolsRequest{},
	)
	if err != nil {
		log.Error(err)
	}
	if farmingPools != nil {
		for _, pool := range farmingPools.GetPools() {
			for _, erc20Token := range pool.GetRewardTokens() {
				tokenSymbol := cbrtypes.GetSymbolFromStakeToken(erc20Token.GetSymbol())
				tokenMap[tokenSymbol] = 1
			}
		}
	}

	var tokenIds []string

	for symbol := range tokenMap {
		tokenId := dal.DB.GetTokenIdBySymbol(symbol)
		if tokenId != "" {
			tokenIds = append(tokenIds, tokenId)
		} else {
			_, mocked := dal.GetMockedPrice(symbol)
			if !mocked {
				log.Errorf("token %s not found in db, table: token_id, please add token_id from www.coingecko.com", symbol)
			}
		}
	}

	symbol2chainIds, _, _, _, err := dal.DB.GetAllChainAndGasToken()
	if err != nil {
		return fmt.Errorf("failed to GetAllChainAndGasToken")
	}
	for sym := range symbol2chainIds {
		tokenId := dal.DB.GetTokenIdBySymbol(sym)
		if tokenId != "" {
			tokenIds = append(tokenIds, tokenId)
		} else {
			_, mocked := dal.GetMockedPrice(sym)
			if !mocked {
				log.Errorf("token %s not found in db, table: token_id, please add token_id from www.coingecko.com", sym)
			}
		}
	}

	if len(tokenIds) == 0 || len(t.vsTokenIds) == 0 {
		return fmt.Errorf("tokenIds and vsTokenIds are required")
	}
	qs := fmt.Sprintf(
		"ids=%s&vs_currencies=%s",
		strings.Join(tokenIds, ","),
		strings.Join(t.vsTokenIds, ","))
	client := resty.New()
	r, err := client.R().SetQueryString(qs).SetResult(&TokenPrices{}).Get(priceApiUrl)
	if err != nil {
		return fmt.Errorf("failed to refresh token price cache: err %s", err)
	}
	if r.StatusCode() != 200 {
		return fmt.Errorf("failed to refresh token price cache: status code %d", r.StatusCode())
	}
	tokenPrices := r.Result().(*TokenPrices)
	newPrices := make(map[string]float64)
	// flatten the nested map since we only care about USD Prices
	for tokenId, vsTokenPrices := range *tokenPrices {
		newPrices[tokenId] = vsTokenPrices["usd"]
	}
	t.Prices = newPrices
	return nil
}
