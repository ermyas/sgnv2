package dal

import (
	"database/sql"
)

// Delayed Operation Type
// table: delayed_op
// column: type
type DelayedOpType uint64

const (
	DelayedOpUnknown DelayedOpType = iota
	DelayedOpTransfer
	DelayedOpRefund
	DelayedOpWithdraw
)

func (dal *DAL) GetDelayedOp(id string) (t uint64, found bool, err error) {
	q := `select type from delayed_op where id = $1`
	err = dal.Db.QueryRow(q, id).Scan(&t)
	if err == sql.ErrNoRows {
		found = false
		err = nil
		return
	}
	if err != nil {
		found = false
		return
	}
	found = true
	return
}

func (dal *DAL) InsertDelayedOp(id, txhash string, t DelayedOpType) error {
	q := `insert into delayed_op (id, tx_hash) values ($1, $2)`
	_, err := dal.Db.Exec(q, id, txhash)
	if err != nil {
		return err
	}
	return nil
}

func (dal *DAL) UpdateDelayedOpType(id string, t DelayedOpType) error {
	q := `update delayed_op set type = $2 where id = $1`
	_, err := dal.Db.Exec(q, id, t)
	if err != nil {
		return err
	}
	return nil
}
