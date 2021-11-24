package explorer

import (
	"database/sql"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
)

type DAL struct {
	*sqldb.Db
}

type TransactionStat struct {
	begin  time.Time
	volume float64
	count  uint64
}

type LiqStat struct {
	begin  time.Time
	volume float64
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

func (d *DAL) InsertDailyLiquidityStat(begin time.Time, volume float64) error {
	q := `UPSERT INTO daily_liquidity (datetime,total_liquidity) VALUES ($1, $2)`
	_, err := d.Exec(q, begin, volume)
	return err
}

func (d *DAL) GetDailyLiquidityStatByRange(begin, end time.Time) ([]*LiqStat, error) {
	q := `select datetime,total_liquidity from daily_liquidity where datetime >= $1 and datetime <= $2 order by datetime asc`
	rows, err := d.Query(q, begin, end)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)
	var liqs []*LiqStat
	for rows.Next() {
		liq := &LiqStat{}
		if err := rows.Scan(&liq.begin, &liq.volume); err != nil {
			return nil, err
		}
		liqs = append(liqs, liq)
	}
	return liqs, nil
}

func (d *DAL) InsertDailyTransactionStat(begin time.Time, volume float64, count uint64) error {
	q := `UPSERT INTO daily_transaction_stat (datetime,transaction_volume,transaction_count) VALUES ($1, $2, $3)`
	_, err := d.Exec(q, begin, volume, count)
	return err
}

func (d *DAL) GetLatestDailyTransactionStat() (*TransactionStat, bool, error) {
	data := &TransactionStat{}
	q := "SELECT datetime,transaction_volume,transaction_count FROM daily_transaction_stat order by datetime desc limit 1"
	err := d.QueryRow(q).Scan(&data.begin, &data.volume, &data.count)
	found, err := sqldb.ChkQueryRow(err)
	if err != nil {
		return nil, false, err
	}
	return data, found, nil
}

func (d *DAL) GetDailyTransactionStat(begin, end time.Time) ([]*TransactionStat, error) {
	q := `select datetime,transaction_volume,transaction_count from daily_transaction_stat where datetime >= $1 and datetime <= $2 order by datetime asc`
	rows, err := d.Query(q, begin, end)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)
	var txs []*TransactionStat
	for rows.Next() {
		tx := &TransactionStat{}
		if err := rows.Scan(&tx.begin, &tx.volume, &tx.count); err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (d *DAL) GetSumTransferVolumeAndCountByDaily(begin, end time.Time) (float64, uint64, error) {
	var volume sql.NullFloat64
	var count sql.NullInt64
	q := "SELECT sum(transaction_volume),sum(transaction_count) FROM daily_transaction_stat where datetime >= $1 and datetime < $2"
	err := d.QueryRow(q, begin, end).Scan(&volume, &count)
	if err != nil {
		return 0, 0, err
	}
	return volume.Float64, uint64(count.Int64), nil
}

// hourly transactionsStat
func (d *DAL) InsertHourlyTransactionStat(begin, end time.Time, volume float64, count uint64) error {
	q := `UPSERT INTO transaction_stat (begin_datetime,end_datetime,transaction_volume,transaction_count) VALUES ($1, $2, $3, $4)`
	_, err := d.Exec(q, begin, end, volume, count)
	return err
}

func (d *DAL) GetSumTransferVolumeAndCount(begin, end time.Time) (float64, uint64, error) {
	var volume sql.NullFloat64
	var count sql.NullInt64
	q := "SELECT sum(transaction_volume),sum(transaction_count) FROM transaction_stat where begin_datetime >= $1 and begin_datetime < $2"
	err := d.QueryRow(q, begin, end).Scan(&volume, &count)
	if err != nil {
		return 0, 0, err
	}
	return volume.Float64, uint64(count.Int64), nil
}

func (d *DAL) GetLatestTransactionStat() (*TransactionStat, bool, error) {
	data := &TransactionStat{}
	q := "SELECT begin_datetime,transaction_volume,transaction_count FROM transaction_stat order by begin_datetime desc limit 1"
	err := d.QueryRow(q).Scan(&data.begin, &data.volume, &data.count)
	found, err := sqldb.ChkQueryRow(err)
	if err != nil {
		return nil, false, err
	}
	return data, found, nil
}

func (d *DAL) InsertDistinctAddr(addr string) error {
	q := `INSERT INTO wallet (addr,create_time) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := d.Exec(q, addr, time.Now())
	return err
}

func (d *DAL) GetWalletCount() (uint64, error) {
	var count sql.NullInt64
	q := "SELECT count(addr) FROM wallet"
	err := d.QueryRow(q).Scan(&count)
	if err != nil {
		return 0, err
	}
	return uint64(count.Int64), nil
}

// for v1 select
func (d *DAL) GetV1TxStatByTimeRange(begin, end time.Time) (float64, uint64, error) {
	var volume sql.NullFloat64
	var count sql.NullInt64
	q := "SELECT sum(volume),count(1) FROM transfer WHERE transfer_type = 1 and create_time >= $1 and create_time < $2"
	err := d.QueryRow(q, begin, end).Scan(&volume, &count)
	if err != nil {
		return 0, 0, err
	}
	return volume.Float64, uint64(count.Int64), nil
}

func (d *DAL) GetV1DistinctTransferAddrByTimeRange(begin, end time.Time) ([]string, error) {
	var addrs []string
	q := "SELECT distinct(sender) FROM transfer WHERE transfer_type = 1 and create_time >= $1 and create_time < $2"
	rows, err := d.Query(q, begin, end)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var addr string
		err = rows.Scan(&addr)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}

func (d *DAL) Close() {
	if d.Db != nil {
		d.Db.Close()
		d.Db = nil
	}
}

func closeRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Warnln("closeRows: error:", err)
	}
}

func (d *DAL) DB() *sqldb.Db {
	return d.Db
}
