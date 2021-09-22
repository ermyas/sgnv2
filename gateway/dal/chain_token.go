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
	SET decimal = $4, addr = $2, contract=$5, max_amt=$6`
	res, err := d.Exec(q, symbol, addr, chainId, decimal, contract, maxAmt)
	return sqldb.ChkExec(res, err, 1, "UpsertTokenBaseInfo")
}

func (d *DAL) UpdateTokenUIInfo(symbol string, chainId uint64, name, icon string, price float64) error {
	q := `UPDATE transfer set name=$3, icon=$4, price=$5 where symbol=$1 AND chain_id=$2`
	res, err := d.Exec(q, symbol, chainId, name, icon, price)
	return sqldb.ChkExec(res, err, 1, "UpdateTokenUIInfo")
}

func (d *DAL) GetToken(symbol string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	var addr string
	var decimal uint64
	var name, icon, contract, amt string
	var price float64
	q := `SELECT address, decimal, name, icon, price, max_amt, contract FROM token WHERE symbol = $1 AND chain_id=$2`
	err := d.QueryRow(q, symbol, chainId).Scan(&addr, &decimal, &name, &icon, &price, &amt, &contract)
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
	q := `SELECT symbol, chain_id, address, decimal, name, icon, price, max_amt, contract FROM token`
	rows, err := d.Query(q)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var symbol, addr string
	var chainId, decimal uint64
	var name, icon, contract, amt string
	var price float64

	resp := make(map[uint64]*webapi.ChainTokenInfo)
	for rows.Next() {
		err = rows.Scan(&symbol, &chainId, &addr, &decimal, &name, &icon, &price, &amt, contract)
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
	q := fmt.Sprintf(`SELECT id, name, icon, price, max_amt, contract FROM chain where %s`, inClause)
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
