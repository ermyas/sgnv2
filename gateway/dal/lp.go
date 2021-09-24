package dal

import (
	"github.com/celer-network/goutils/sqldb"
)

func (d *DAL) UpsertLP(usrAddr, tokenSymbol, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (usr_addr, chain_id, seq_num, lp_type) DO UPDATE
	SET status = $8, update_time = $6`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, amt, txHash, now(), now(), status, lpType, seqNum)
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
