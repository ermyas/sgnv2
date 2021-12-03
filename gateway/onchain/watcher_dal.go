package onchain

import (
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/dal"
)

// MUST have proper lock/unlock as one watcherDAL is shared by all chains
type watcherDAL struct {
	db *dal.DAL
}

func NewWatcherDAL(db *dal.DAL) *watcherDAL {
	return &watcherDAL{db: db}
}

func (dal *watcherDAL) InsertMonitor(event string, blockNum uint64, blockIdx int64, restart bool) error {
	q := `INSERT INTO monitor_block (event, block_num, block_idx, restart) 
				VALUES ($1, $2, $3, $4)`
	res, err := dal.db.Exec(q, event, blockNum, blockIdx, restart)
	return sqldb.ChkExec(res, err, 1, "InsertMonitor")
}

func (dal *watcherDAL) GetMonitorBlock(event string) (blknum uint64, blkidx int64, found bool, err error) {
	q := `SELECT block_num, block_idx 
				FROM monitor_block 
				WHERE event = $1`
	err = dal.db.QueryRow(q, event).Scan(&blknum, &blkidx)
	found, err = sqldb.ChkQueryRow(err)
	return
}

func (dal *watcherDAL) UpdateMonitorBlock(event string, blockNum uint64, blockIdx int64) error {
	q := `UPDATE monitor_block 
				SET block_num = $2, block_idx = $3
				WHERE event = $1`
	res, err := dal.db.Exec(q, event, blockNum, blockIdx)
	return sqldb.ChkExec(res, err, 1, "UpdateMonitorBlock")
}

func (dal *watcherDAL) UpsertMonitorBlock(event string, blockNum uint64, blockIdx int64, restart bool) error {
	q := `UPSERT INTO monitor_block (event, block_num, block_idx, restart) 
				VALUES ($1, $2, $3, $4)`
	_, err := dal.db.Exec(q, event, blockNum, blockIdx, restart)
	return err
}
