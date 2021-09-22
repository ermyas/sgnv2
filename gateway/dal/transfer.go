package dal

import (
	"github.com/celer-network/goutils/sqldb"
)

func (d *DAL) InsertTransfer(transferId, usrAddr, tokenSymbol string, srcChainId, dsChainId uint64) error {
	q := `INSERT INTO transfer (transfer_id, usr_addr, token_symbol, src_chain_id, dst_chain_id)
                VALUES ($1, $2, $3, $4, $5)`
	res, err := d.Exec(q, transferId, usrAddr, tokenSymbol, srcChainId, dsChainId)
	return sqldb.ChkExec(res, err, 1, "InsertTransfer")
}

func (d *DAL) GetTransfer(transferId string) (string, string, uint64, uint64, bool, error) {
	var usrAddr, tokenSymbol string
	var srcChainId, dsChainId uint64
	q := `SELECT usr_addr, token_symbol, src_chain_id, dst_chain_id FROM transfer WHERE transfer_id = $1`
	err := d.QueryRow(q, transferId).Scan(&usrAddr, &tokenSymbol, &srcChainId, &dsChainId)
	found, err := sqldb.ChkQueryRow(err)
	return usrAddr, tokenSymbol, srcChainId, dsChainId, found, err
}
