package dal

import (
	"fmt"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (d *DAL) UpsertTokenBaseInfo(symbol, addr string, chainId, decimal uint64, disabled bool) error {
	q := `INSERT INTO token (symbol, address, chain_id, decimal, update_time, disabled)
                VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (symbol, chain_id) DO UPDATE
	SET decimal = $4, address = $2, update_time=$5, disabled=$6`
	res, err := d.Exec(q, symbol, addr, chainId, decimal, now(), disabled)
	return sqldb.ChkExec(res, err, 1, "UpsertTokenBaseInfo")
}

func (d *DAL) UpsertRewardToken(symbol, addr string, chainId, decimal uint64) error {
	q := `INSERT INTO reward_token (symbol, address, chain_id, decimal, update_time)
                VALUES ($1, $2, $3, $4, $5) ON CONFLICT (symbol, chain_id) DO UPDATE
	SET decimal = $4, address = $2, update_time=$5`
	res, err := d.Exec(q, symbol, addr, chainId, decimal, now())
	return sqldb.ChkExec(res, err, 1, "UpsertRewardToken")
}

func (d *DAL) GetRewardTokenBySymbol(symbol string, chainId uint64) (*types.Token, bool, error) {
	var addr string
	var decimal uint64
	q := `SELECT address, decimal FROM reward_token WHERE symbol = $1 AND chain_id=$2`
	err := d.QueryRow(q, symbol, chainId).Scan(&addr, &decimal)
	found, err := sqldb.ChkQueryRow(err)
	return &types.Token{
		Symbol:  symbol,
		Address: addr,
		Decimal: int32(decimal),
	}, found, err
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
	q := `UPDATE token set name=$3, icon=$4 where symbol=$1 AND chain_id=$2`
	res, err := d.Exec(q, symbol, chainId, name, icon)
	return sqldb.ChkExec(res, err, 1, "UpdateTokenUIInfo")
}

func (d *DAL) GetTokenBySymbol(symbol string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	cache := GetTokenCacheBySymbol(chainId, symbol)
	if cache != nil {
		return cache, true, nil
	}
	token, found, err := d.getTokenBySymbol(symbol, chainId)
	if found && err == nil {
		SetTokenCache(chainId, token)
	}
	return token, found, err
}

func (d *DAL) getTokenBySymbol(symbol string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	q := `SELECT address, decimal, name, icon FROM token WHERE symbol = $1 AND chain_id=$2`
	return d.getTokenBySymbolWithQ(q, symbol, chainId)
}

func (d *DAL) GetTokenBySymbolForTransfer(symbol string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	q := `SELECT address, decimal, name, icon FROM token WHERE symbol = $1 AND chain_id=$2 AND disabled = false`
	return d.getTokenBySymbolWithQ(q, symbol, chainId)
}

func (d *DAL) getTokenBySymbolWithQ(q, symbol string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	var addr string
	var decimal uint64
	var name, icon string
	err := d.QueryRow(q, symbol, chainId).Scan(&addr, &decimal, &name, &icon)
	found, err := sqldb.ChkQueryRow(err)
	return &webapi.TokenInfo{
		Token: &types.Token{
			Symbol:  symbol,
			Address: addr,
			Decimal: int32(decimal),
		},
		Name: name,
		Icon: icon,
	}, found, err
}

func (d *DAL) GetTokenByAddr(addr string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	cache := GetTokenCacheByAddr(chainId, addr)
	if cache != nil {
		return cache, true, nil
	}
	token, found, err := d.getTokenByAddr(addr, chainId)
	if found && err == nil {
		SetTokenCache(chainId, token)
	}
	return token, found, err
}

func (d *DAL) getTokenByAddr(addr string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	var symbol string
	var decimal uint64
	var name, icon string
	q := `SELECT symbol, decimal, name, icon FROM token WHERE address = $1 AND chain_id=$2`
	err := d.QueryRow(q, addr, chainId).Scan(&symbol, &decimal, &name, &icon)
	found, err := sqldb.ChkQueryRow(err)
	return &webapi.TokenInfo{
		Token: &types.Token{
			Symbol:  symbol,
			Address: addr,
			Decimal: int32(decimal),
		},
		Name: name,
		Icon: icon,
	}, found, err
}

func (d *DAL) GetEnabledChainTokenList() (map[uint32]*webapi.ChainTokenInfo, error) {
	q := `SELECT symbol, chain_id, address, decimal, name, icon FROM token where disabled = false`
	return d.getChainTokenList(q)
}

func (d *DAL) GetChainTokenList() (map[uint32]*webapi.ChainTokenInfo, error) {
	q := `SELECT symbol, chain_id, address, decimal, name, icon FROM token`
	return d.getChainTokenList(q)
}

func (d *DAL) getChainTokenList(q string) (map[uint32]*webapi.ChainTokenInfo, error) {
	rows, err := d.Query(q)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var symbol, addr string
	var chainId, decimal uint32
	var name, icon string

	resp := make(map[uint32]*webapi.ChainTokenInfo)
	for rows.Next() {
		err = rows.Scan(&symbol, &chainId, &addr, &decimal, &name, &icon)
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
			Name: name,
			Icon: icon,
		}
		if !found {
			tps = &webapi.ChainTokenInfo{}
		}
		tps.Token = append(tps.GetToken(), tp)
		resp[chainId] = tps
	}
	return resp, nil
}

// GetAllChainAndGasToken return key is gas token symbol, value is chainId
func (d *DAL) GetAllChainAndGasToken() (symbol2chainIds map[string][]uint64, chainId2Symbol map[uint64]string,
	chainId2DropGas map[uint64]string, chainId2SuggestedBaseFee map[uint64]float64, error error) {
	symbol2chainIds = make(map[string][]uint64)
	chainId2Symbol = make(map[uint64]string)
	chainId2DropGas = make(map[uint64]string)
	chainId2SuggestedBaseFee = make(map[uint64]float64)
	q := `SELECT id, gas_token_symbol, drop_gas_amt, suggested_base_fee
          FROM chain
          WHERE gas_token_symbol is not null`
	rows, err := d.Query(q)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer closeRows(rows)
	var id uint64
	var gasTokenSymbol, dropGasAmt string
	var suggestedBaseFee float64
	for rows.Next() {
		err = rows.Scan(&id, &gasTokenSymbol, &dropGasAmt, &suggestedBaseFee)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if gasTokenSymbol == "" {
			continue
		}
		chainId2Symbol[id] = gasTokenSymbol
		chainId2DropGas[id] = dropGasAmt
		chainId2SuggestedBaseFee[id] = suggestedBaseFee
		if len(symbol2chainIds[gasTokenSymbol]) == 0 {
			symbol2chainIds[gasTokenSymbol] = []uint64{id}
		} else {
			symbol2chainIds[gasTokenSymbol] = append(symbol2chainIds[gasTokenSymbol], id)
		}
	}
	return symbol2chainIds, chainId2Symbol, chainId2DropGas, chainId2SuggestedBaseFee, nil
}
func (d *DAL) GetChainInfo(ids []uint32) ([]*webapi.Chain, error) {
	inClause := sqldb.InClause("id", len(ids), 1)
	q := fmt.Sprintf(`SELECT id, name, icon, block_delay, gas_token_symbol, explore_url, rpc_url, contract, drop_gas_amt, suggested_base_fee FROM chain WHERE %s`, inClause)
	var params []interface{}
	for _, v := range ids {
		params = append(params, v)
	}
	rows, err := d.Query(q, params...)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var id, blockDelay uint32
	var name, icon, gasTokenSymbol, exploreUrl, rpcUrl, contract, dropGasAmt string
	var suggestedBaseFee float64
	var tps []*webapi.Chain

	for rows.Next() {
		err = rows.Scan(&id, &name, &icon, &blockDelay, &gasTokenSymbol, &exploreUrl, &rpcUrl, &contract, &dropGasAmt, &suggestedBaseFee)
		if err != nil {
			return nil, err
		}
		tp := &webapi.Chain{
			Id:               id,
			Name:             name,
			Icon:             icon,
			BlockDelay:       blockDelay,
			GasTokenSymbol:   gasTokenSymbol,
			ExploreUrl:       exploreUrl,
			ContractAddr:     contract,
			DropGasAmt:       dropGasAmt,
			SuggestedBaseFee: suggestedBaseFee,
		}
		tps = append(tps, tp)
	}
	return tps, nil
}
func (d *DAL) UpsertChainUIInfo(id uint64, name, icon, url, gasTokenSymbol, exploreUrl, rpcUrl, dropGasAmt string, suggestedBaseFee float64) error {
	q := `INSERT INTO chain (id, name, icon, tx_url, gas_token_symbol, explore_url, rpc_url, drop_gas_amt, suggested_base_fee)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (id) DO UPDATE
	SET name=$2, icon=$3, tx_url=$4, gas_token_symbol=$5, explore_url=$6, rpc_url=$7, drop_gas_amt=$8, suggested_base_fee=9`
	res, err := d.Exec(q, id, name, icon, url, gasTokenSymbol, exploreUrl, rpcUrl, dropGasAmt, suggestedBaseFee)
	return sqldb.ChkExec(res, err, 1, "UpsertChainInfo")
}

func (d *DAL) UpsertChainBaseInfo(id uint64, blockDelay uint32, contractAddr string) error {
	q := `INSERT INTO chain (id, block_delay, contract)
                VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE
	SET block_delay=$2, contract=$3`
	res, err := d.Exec(q, id, blockDelay, contractAddr)
	return sqldb.ChkExec(res, err, 1, "UpsertChainBaseInfo")
}

func (d *DAL) GetChain(id uint64) (*webapi.Chain, string, bool, error) {
	cache, url := GetChainCache(id)
	if cache != nil {
		return cache, url, true, nil
	}
	var name, icon, txUrl, gasTokenSymbol, exploreUrl, rpcUrl, contract, dropGasAmt string
	var blockDelay uint32
	var suggestedBaseFee float64
	q := `SELECT name, icon, tx_url, block_delay, gas_token_symbol, explore_url, rpc_url, contract, drop_gas_amt, suggested_base_fee FROM chain where id=$1`
	err := d.QueryRow(q, id).Scan(&name, &icon, &txUrl, &blockDelay, &gasTokenSymbol, &exploreUrl, &rpcUrl, &contract, &dropGasAmt, &suggestedBaseFee)
	found, err := sqldb.ChkQueryRow(err)
	chain := &webapi.Chain{
		Id:               uint32(id),
		Name:             name,
		Icon:             icon,
		BlockDelay:       blockDelay,
		GasTokenSymbol:   gasTokenSymbol,
		ExploreUrl:       exploreUrl,
		ContractAddr:     contract,
		DropGasAmt:       dropGasAmt,
		SuggestedBaseFee: suggestedBaseFee,
	}
	if found && err == nil {
		SetChainCache(chain, txUrl)
	}
	return chain, txUrl, found, err
}

func (d *DAL) GetChainBlockDelay(id uint64) (uint32, bool, error) {
	var blockDelay uint32
	q := `SELECT block_delay FROM chain where id=$1`
	err := d.QueryRow(q, id).Scan(&blockDelay)
	found, err := sqldb.ChkQueryRow(err)
	return blockDelay, found, err
}
