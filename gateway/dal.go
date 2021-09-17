package webapi

import (
	"database/sql"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
)

type DAL struct {
	*sqldb.Db
}

func NewDAL(driver, info string, poolSize int) (*DAL, error) {
	db, err := sqldb.NewDb(driver, info, poolSize)
	if err != nil {
		log.Errorf("fail with db init:%s, %s, %d, err:%+v", driver, info, poolSize, err)
		return nil, err
	}

	dal := &DAL{
		db,
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

func now() time.Time {
	return time.Now().UTC()
}

func closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Warnln("closeRows: error:", err)
	}
}

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
