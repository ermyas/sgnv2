package dal

import (
	"fmt"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (d *DAL) UpsertTokenBaseInfo(symbol, addr, contract, maxAmt string, chainId, decimal uint64) error {
	q := `INSERT INTO token (symbol, address, chain_id, decimal, max_amt, contract)
                VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (symbol, chain_id) DO UPDATE
	SET decimal = $4, address = $2, contract=$5, max_amt=$6`
	res, err := d.Exec(q, symbol, addr, chainId, decimal, contract, maxAmt)
	return sqldb.ChkExec(res, err, 1, "UpsertTokenBaseInfo")
}

func (d *DAL) GetTokenSymbols() ([]string, error) {
	q := `select distinct symbol from token`
	rows, err := d.Query(q)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var symbol string

	var resp []string
	for rows.Next() {
		err = rows.Scan(&symbol)
		if err != nil {
			return nil, err
		}
		resp = append(resp, symbol)
	}
	return resp, nil
}

func (d *DAL) UpdateTokenUIInfo(symbol string, chainId uint64, name, icon string) error {
	q := `UPDATE transfer set name=$3, icon=$4 where symbol=$1 AND chain_id=$2`
	res, err := d.Exec(q, symbol, chainId, name, icon)
	return sqldb.ChkExec(res, err, 1, "UpdateTokenUIInfo")
}

func (d *DAL) GetTokenBySymbol(symbol string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	var addr string
	var decimal uint64
	var name, icon, contract, amt string
	q := `SELECT address, decimal, name, icon, max_amt, contract FROM token WHERE symbol = $1 AND chain_id=$2`
	err := d.QueryRow(q, symbol, chainId).Scan(&addr, &decimal, &name, &icon, &amt, &contract)
	found, err := sqldb.ChkQueryRow(err)
	return &webapi.TokenInfo{
		Token: &types.Token{
			Symbol:  symbol,
			Address: addr,
			Decimal: int32(decimal),
		},
		Name:         name,
		Icon:         icon,
		MaxAmt:       amt,
		ContractAddr: contract,
	}, found, err
}

func (d *DAL) GetTokenByAddr(addr string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	var symbol string
	var decimal uint64
	var name, icon, contract, amt string
	q := `SELECT symbol, decimal, name, icon, max_amt, contract FROM token WHERE addr = $1 AND chain_id=$2`
	err := d.QueryRow(q, addr, chainId).Scan(&symbol, &decimal, &name, &icon, &amt, &contract)
	found, err := sqldb.ChkQueryRow(err)
	return &webapi.TokenInfo{
		Token: &types.Token{
			Symbol:  symbol,
			Address: addr,
			Decimal: int32(decimal),
		},
		Name:         name,
		Icon:         icon,
		MaxAmt:       amt,
		ContractAddr: contract,
	}, found, err
}

func (d *DAL) UpsertChain(id uint64, name, icon string) error {
	q := `INSERT INTO chain (name, icon)
               VALUES ($2, $3) WHERE id=$1`
	res, err := d.Exec(q, id, name, icon)
	return sqldb.ChkExec(res, err, 1, "UpsertChain")
}

func (d *DAL) GetChainTokenList() (map[uint64]*webapi.ChainTokenInfo, error) {
	q := `SELECT symbol, chain_id, address, decimal, name, icon, max_amt, contract FROM token`
	rows, err := d.Query(q)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var symbol, addr string
	var chainId, decimal uint64
	var name, icon, contract, amt string

	resp := make(map[uint64]*webapi.ChainTokenInfo)
	for rows.Next() {
		err = rows.Scan(&symbol, &chainId, &addr, &decimal, &name, &icon, &amt, &contract)
		if err != nil {
			return nil, err
		}
		tps, found := resp[chainId]
		tp := &webapi.TokenInfo{
			Token: &types.Token{
				Symbol:  symbol,
				Address: addr,
				Decimal: int32(decimal),
			},
			Name:         name,
			Icon:         icon,
			MaxAmt:       amt,
			ContractAddr: contract,
		}
		if !found {
			tps = &webapi.ChainTokenInfo{}
		}
		tps.Token = append(tps.GetToken(), tp)
		resp[chainId] = tps
	}
	return resp, nil
}

func (d *DAL) GetChainInfo(ids []uint64) ([]*webapi.Chain, error) {
	inClause := sqldb.InClause("id", len(ids), 1)
	q := fmt.Sprintf(`SELECT id, name, icon FROM chain WHERE %s`, inClause)
	var params []interface{}
	for _, v := range ids {
		params = append(params, v)
	}
	rows, err := d.Query(q, params...)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var id uint64
	var name, icon string

	var tps []*webapi.Chain
	for rows.Next() {
		err = rows.Scan(&id, &name, &icon)
		if err != nil {
			return nil, err
		}
		tp := &webapi.Chain{
			Id:   id,
			Name: name,
			Icon: icon,
		}
		tps = append(tps, tp)
	}
	return tps, nil
}
func (d *DAL) UpsertChainInfo(id uint64, name, icon, url string) error {
	q := `INSERT INTO chain (id, name, icon, tx_url)
                VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO UPDATE
	SET name=$2, icon=$3, tx_url=$4`
	res, err := d.Exec(q, id, name, icon, url)
	return sqldb.ChkExec(res, err, 1, "UpsertChainInfo")
}

func (d *DAL) GetChain(id uint64) (*webapi.Chain, string, bool, error) {
	var name, icon, url string
	q := `SELECT name, icon, tx_url FROM chain where id=$1`
	err := d.QueryRow(q, id).Scan(&name, &icon, &url)
	found, err := sqldb.ChkQueryRow(err)
	return &webapi.Chain{
		Id:   id,
		Name: name,
		Icon: icon,
	}, url, found, err
}
