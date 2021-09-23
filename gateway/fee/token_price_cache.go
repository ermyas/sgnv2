package fee

import (
	"encoding/json"
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/lthibault/jitterbug"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"math"
	"math/big"
	"strings"
	"time"
)

const priceApiUrl = "https://api.coingecko.com/api/v3/simple/price"

// VsTokenPrices vsToken -> price
type VsTokenPrices map[string]float64

type TokenPrices map[string]VsTokenPrices

type TokenPriceCache struct {
	tokenIds    []string
	vsTokenIds  []string
	Prices      map[string]float64 // do not access this map with token symbol since its key is coingecko's tokenId
	allTokenIds map[string]AToken
}

// NewTokenPriceCache builds a new instance of TokenPriceCache with an empty Prices map.
// Note you still have to manually start polling loop using StartTokenPricePolling()
func NewTokenPriceCache() *TokenPriceCache {
	var tokenIds []string
	var vsTokenIds = []string{"usd"}

	feeSvr := &TokenPriceCache{
		tokenIds:   tokenIds,
		vsTokenIds: vsTokenIds,
		Prices:     make(map[string]float64),
	}
	err := feeSvr.cacheTokenData()
	if err != nil {
		log.Error("NewTokenPriceCache error", err)
	}
	feeSvr.StartTokenPricePolling(1 * time.Minute)
	return feeSvr
}

// StartTokenPricePolling starts a loop with the given interval and 3s stdev for polling price
func (t *TokenPriceCache) StartTokenPricePolling(interval time.Duration) {
	go func() {
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {

			err := t.refreshCache()
			if err != nil {
				log.Errorln("failed to refresh token price cache:", err)
			}
		}
	}()
	time.Sleep(2 * time.Second)
}

// GetUsdPrice gets the token/USD price by token symbol. e.g. "ETH", "DAI", "USDT"
func (t *TokenPriceCache) GetUsdPrice(tokenSymbol string) (float64, error) {
	token, ok := t.allTokenIds[tokenSymbol]
	if !ok {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
	}
	tokenId, ok := token.Data["id"]
	if !ok {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
	}
	price, ok := t.Prices[tokenId]
	if !ok {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
	}
	return price, nil
}

func (t *TokenPriceCache) GetTokenPrice(token *types.Token, chainToken *types.Token, chainTokenAmt *big.Int) (*big.Int, error) {
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

func (t *TokenPriceCache) refreshCache() error {
	var tokenIds []string
	for symbol := range t.Prices {
		token, found := t.allTokenIds[symbol]
		if found {
			tokenIds = append(tokenIds, token.Data["id"])
		} else {
			log.Errorf("token %s not found in json file", symbol)
		}
	}
	if len(t.tokenIds) == 0 || len(t.vsTokenIds) == 0 {
		return fmt.Errorf("tokenIds and vsTokenIds are required")
	}
	qs := fmt.Sprintf(
		"ids=%s&vs_currencies=%s",
		strings.Join(t.tokenIds, ","),
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

type AToken struct {
	//Data Id     string `json:"id"`
	//Data Symbol string `json:"symbol"`
	//Data Name   string `json:"name"`
	Data map[string]string
}

type Tokens struct {
	tokenData []AToken
}

func (t *TokenPriceCache) cacheTokenData() error {
	var tokens Tokens
	file, err := ioutil.ReadFile("token_info.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &tokens)
	if err != nil {
		log.Fatal(err)
	}
	resp := make(map[string]AToken)
	for _, token := range tokens.tokenData {
		resp[token.Data["symbol"]] = token
	}
	t.allTokenIds = resp
	return err
}
