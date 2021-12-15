package dal

import (
	"database/sql"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
)

type TokenData struct {
	Id     string
	Symbol string
}
type DAL struct {
	*sqldb.Db
	Prices      map[string]float64 // do not access this map with token symbol since its key is coingecko's tokenId
	AllTokenIds map[string]*TokenData
}

func NewDAL(url string) *DAL {
	db, err := sqldb.NewDb("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", url), 20)
	if err != nil {
		log.Fatalf("Failed to create db with url %s: %+v", url, err)
	}
	dal := &DAL{
		db,
		make(map[string]float64),
		make(map[string]*TokenData),
	}
	return dal
}

func (d *DAL) Close() {
	if d.Db != nil {
		d.Db.Close()
		d.Db = nil
	}
}

func (d *DAL) DB() *sqldb.Db {
	return d.Db
}

func (d *DAL) GetUsdVolume(tokenSymbol string, chainId uint64, amt *big.Int) (float64, error) {
	token, foundToken, dbErr := d.getTokenBySymbol(tokenSymbol, chainId)
	if dbErr != nil {
		return 0, dbErr
	}
	if !foundToken {
		return 0, fmt.Errorf("invalid token, symbol:%s, chainId:%d", tokenSymbol, chainId)
	}
	usdPrice, err := d.GetUsdPrice(tokenSymbol)
	if err != nil {
		return 0, err
	}
	tokenAmt, _ := new(big.Float).Quo(new(big.Float).SetInt(amt), big.NewFloat(math.Pow(10, float64(token.GetToken().GetDecimal())))).Float64()
	return tokenAmt * usdPrice, nil
}

// GetUsdPrice gets the token/USD price by token symbol. e.g. "ETH", "DAI", "USDT"
func (d *DAL) GetUsdPrice(tokenSymbol string) (float64, error) {
	if tokenSymbol == "WETH" {
		// will always use ETH price
		tokenSymbol = "ETH"
	}
	tokenId := d.GetTokenIdBySymbol(tokenSymbol)
	if tokenId == "" {
		price, mocked := GetMockedPrice(tokenSymbol) // try to use mocked price if token not found
		if mocked {
			return price, nil
		} else {
			return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
		}
	}
	price, ok := d.Prices[tokenId]
	if !ok {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
	}
	return price, nil
}

func now() time.Time {
	return time.Now().UTC()
}

func closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Warnln("closeRows: error:", err)
	}
}

func GetMockedPrice(symbol string) (float64, bool) {
	if symbol == "TCELR" || symbol == "LYRA" {
		// new token, mock price
		return 0.5, true
	}
	if symbol == "DOMI" {
		// new token, mock price
		return 0.15, true
	}
	return 0, false
}
