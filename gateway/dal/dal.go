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
	Name   string
}
type DAL struct {
	*sqldb.Db
	Prices      map[string]float64 // do not access this map with token symbol since its key is coingecko's tokenId
	AllTokenIds map[string]*TokenData
}

func NewDAL(driver, info string, poolSize int) (*DAL, error) {
	db, err := sqldb.NewDb(driver, info, poolSize)
	if err != nil {
		log.Errorf("fail with db init:%s, %s, %d, err:%+v", driver, info, poolSize, err)
		return nil, err
	}

	dal := &DAL{
		db,
		make(map[string]float64),
		make(map[string]*TokenData),
	}
	return dal, nil
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
	token, ok := d.AllTokenIds[tokenSymbol]
	if !ok {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
	}
	tokenId := token.Id
	if tokenId == "" {
		return 0, fmt.Errorf("unsupported token %s", tokenSymbol)
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
