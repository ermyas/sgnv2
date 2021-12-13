package explorer

import (
	"fmt"

	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	gatewaysvc "github.com/celer-network/sgn-v2/gateway/svc"
	"gopkg.in/resty.v1"
)

func defaultExplorerErr(msg string) *ErrMsg {
	return &ErrMsg{
		ErrCode_ERROR_CODE_UNDEFINED,
		msg,
	}
}

func checkTimeRange(begin, end time.Time) *ErrMsg {
	if begin.After(end) {
		return &ErrMsg{
			ErrCode_ERROR_CODE_INVALID_TIME_PARAM,
			"begin time is after end time",
		}
	}
	/*if begin.Before(defaultStatBeginTime) {
		return &ErrMsg{
			ErrCode_ERROR_CODE_INVALID_TIME_PARAM,
			fmt.Sprintf("begin time should not earlier than %s", defaultStatBeginTime.String()),
		}
	}*/
	return nil
}

// input Milliseconds
func fmtDailyTime(ts uint64) time.Time {
	return common.TsMilliToTime(ts).Truncate(24 * time.Hour)
}

// TODO, should find better way here
var tokenSymbolTokenIds = map[string]string{
	"ETH":   "ethereum",
	"WETH":  "weth",
	"USDC":  "usd-coin",
	"USDT":  "tether",
	"DAI":   "dai",
	"BUSD":  "binance-usd",
	"BNB":   "binancecoin",
	"MATIC": "matic-network",
	"IF":    "impossible-finance",
	"XDAI":  "xdai",
	"OKT":   "exchain",
	"AVAX":  "avalanche-2",
	"FTM":   "fantom",
	"ONE":   "harmony",
	"CELR":  "celer-network",
	"HT":    "huobi-token",
	"WBTC":  "wrapped-bitcoin",
	"DODO":  "dodo",
	"MCB":   "mcdex",
	"CELO":  "celo",
	"LYRA":  "scrypta",
	"IMX":   "impermax",
}

func GetUsdPrices() (map[string]float64, error) {
	var ids []string
	for _, v := range tokenSymbolTokenIds {
		ids = append(ids, v)
	}
	qs := fmt.Sprintf(
		"ids=%s&vs_currencies=%s",
		strings.Join(ids, ","),
		strings.Join([]string{"usd"}, ","))
	client := resty.New()
	r, err := client.R().SetQueryString(qs).SetResult(&gatewaysvc.TokenPrices{}).Get("https://api.coingecko.com/api/v3/simple/price")
	if err != nil {
		return nil, err
	}
	if r.StatusCode() != 200 {
		return nil, fmt.Errorf("fail to get usd price")
	}
	tokenPrices := r.Result().(*gatewaysvc.TokenPrices)
	newPrices := make(map[string]float64)
	// flatten the nested map since we only care about USD Prices
	for tokenId, vsTokenPrices := range *tokenPrices {
		newPrices[tokenId] = vsTokenPrices["usd"]
	}
	// TODO here is fo hard code usd price
	newPrices["LYRA"] = 0.5
	log.Infof("new prices:%+v", newPrices)
	return newPrices, nil
}
