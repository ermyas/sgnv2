package executor

import (
	"github.com/celer-network/goutils/eth/mon2"
	//"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/sqldb"
)

// implements WatchDAL interface methods
var _ mon2.DAL = Dal

//func (dal *DAL) InsertMonitor(event string, blockNum uint64, blockIdx int64, restart bool) error {
//	q := `INSERT INTO monitor_block (event, block_num, block_idx, restart)
//				VALUES ($1, $2, $3, $4)`
//	res, err := dal.Exec(q, event, blockNum, blockIdx, restart)
//	return sqldb.ChkExec(res, err, 1, "InsertMonitor")
//}

func (dal *DAL) GetMonitorBlock(key string) (blknum uint64, blkidx int64, found bool, err error) {
	q := `SELECT block_num, block_idx 
				FROM monitor_block 
				WHERE chid_addr = $1`
	err = dal.QueryRow(q, key).Scan(&blknum, &blkidx)
	found, err = sqldb.ChkQueryRow(err)
	return
}

func (dal *DAL) SetMonitorBlock(key string, blockNum uint64, blockIdx int64) error {
	q := `UPSERT INTO monitor_block (event, block_num, block_idx) 
				VALUES ($1, $2, $3)`
	_, err := dal.Exec(q, key, blockNum, blockIdx)
	return err
}

//func (dal *DAL) UpdateMonitorBlock(event string, blockNum uint64, blockIdx int64) error {
//	q := `UPDATE monitor_block
//				SET block_num = $2, block_idx = $3
//				WHERE event = $1`
//	res, err := dal.Exec(q, event, blockNum, blockIdx)
//	return sqldb.ChkExec(res, err, 1, "UpdateMonitorBlock")
//}
//
//func (dal *DAL) UpsertMonitorBlock(event string, blockNum uint64, blockIdx int64, restart bool) error {
//	q := `UPSERT INTO monitor_block (event, block_num, block_idx, restart)
//				VALUES ($1, $2, $3, $4)`
//	_, err := dal.Exec(q, event, blockNum, blockIdx, restart)
//	return err
//}
