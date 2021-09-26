package dal

import (
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (d *DAL) UpsertLP(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11) ON CONFLICT (usr_addr, chain_id, seq_num, lp_type) DO UPDATE
	SET status = $8, update_time = $6`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum)
	return sqldb.ChkExec(res, err, 1, "UpsertLP")
}

func (d *DAL) UpdateLPStatus(seqNum, status uint64) error {
	q := `UPDATE lp SET status=$2, update_time=$3 WHERE seq_num=$1`
	res, err := d.Exec(q, seqNum, status, now())
	return sqldb.ChkExec(res, err, 1, "UpdateLPStatus")
}

func (d *DAL) GetLPInfo(seqNum uint64) (uint64, string, uint64, bool, error) {
	var chainId, status uint64
	var txHash string
	q := `SELECT chain_id, tx_hash, status FROM lp WHERE seqNum = $1`
	err := d.QueryRow(q, seqNum).Scan(&chainId, &txHash, &status)
	found, err := sqldb.ChkQueryRow(err)
	return chainId, txHash, status, found, err
}

func (d *DAL) GetAllLpChainToken(usr string) ([]*types.ChainTokenAddrPair, error) {
	q := "SELECT chain_id, token_addr FROM lp WHERE usr_addr = $1 group by chain_id, token_addr"
	rows, err := d.Query(q, usr)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var tps []*types.ChainTokenAddrPair
	var chainId uint64
	var tokenAddr string
	for rows.Next() {
		err = rows.Scan(&chainId, &tokenAddr)
		if err != nil {
			return nil, err
		}

		tp := &types.ChainTokenAddrPair{
			ChainId:   chainId,
			TokenAddr: tokenAddr,
		}
		tps = append(tps, tp)
	}

	return tps, nil
}
