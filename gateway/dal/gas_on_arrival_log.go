package dal

import (
	"database/sql"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"math/big"
	"time"
)

const (
	GasOnArrivalStatusUnknown = iota
	GasOnArrivalStatusNew
	GasOnArrivalStatusFail
	GasOnArrivalStatusSuccess
)

type GasOnArrivalLog struct {
	TransferId string
	UsrAddr    string
	ChainId    uint64
	Status     uint64
	TxHash     string
	DropGasAmt *big.Int
	UpdateTime time.Time
	CreateTime time.Time
}

func (d *DAL) FindFailedGasOnArrivalLog(beginTime time.Time) ([]*GasOnArrivalLog, error) {
	q := `SELECT transfer_id, usr_addr, chain_id, status, tx_hash, drop_gas_amt, update_time, create_time
		  FROM gas_on_arrival_log 
          WHERE create_time > $1 AND status = $2`
	rows, err := d.Query(q, beginTime, GasOnArrivalStatusFail)
	if err != nil {
		log.Errorf("db error:%v", err)
		return nil, err
	}
	defer closeRows(rows)

	var transferId, userAddr, txHash string
	var txHashStr, dropGasAmtStr sql.NullString
	var chainId, status uint64
	var createTime, updateTime time.Time
	var res []*GasOnArrivalLog
	for rows.Next() {
		err = rows.Scan(&transferId, &userAddr, &chainId, &status, &txHashStr, &dropGasAmtStr, &updateTime, &createTime)
		if err != nil {
			return nil, err
		}
		dropGasAmt := big.NewInt(0)
		if dropGasAmtStr.Valid {
			dropGasAmt, _ = big.NewInt(0).SetString(dropGasAmtStr.String, 10)
		}
		if txHashStr.Valid {
			txHash = txHashStr.String
		}

		l := &GasOnArrivalLog{
			TransferId: transferId,
			UsrAddr:    userAddr,
			ChainId:    chainId,
			Status:     status,
			TxHash:     txHash,
			DropGasAmt: dropGasAmt,
			UpdateTime: updateTime,
			CreateTime: createTime,
		}
		res = append(res, l)
	}
	return res, nil
}

func (d *DAL) FindGasOnArrivalLog(transferId string) (*GasOnArrivalLog, bool, error) {
	q := `SELECT transfer_id, usr_addr, chain_id, status, tx_hash, drop_gas_amt, update_time, create_time
		  FROM gas_on_arrival_log 
          WHERE transfer_id = $1`
	var userAddr, txHash string
	var txHashStr, dropGasAmtStr sql.NullString
	var chainId, status uint64
	var createTime, updateTime time.Time
	err := d.QueryRow(q, transferId).
		Scan(&transferId, &userAddr, &chainId, &status, &txHashStr, &dropGasAmtStr, &updateTime, &createTime)
	found, err := sqldb.ChkQueryRow(err)
	dropGasAmt := big.NewInt(0)
	if dropGasAmtStr.Valid {
		dropGasAmt, _ = big.NewInt(0).SetString(dropGasAmtStr.String, 10)
	}
	if txHashStr.Valid {
		txHash = txHashStr.String
	}
	return &GasOnArrivalLog{
		TransferId: transferId,
		UsrAddr:    userAddr,
		ChainId:    chainId,
		Status:     status,
		TxHash:     txHash,
		DropGasAmt: dropGasAmt,
		UpdateTime: updateTime,
		CreateTime: createTime,
	}, found, nil
}

func (d *DAL) NewGasOnArrivalLog(transferId, userAddr string, chainId uint64, dropGasAmt *big.Int) error {
	q := `INSERT INTO gas_on_arrival_log (transfer_id, usr_addr, chain_id, status, drop_gas_amt)
                VALUES ($1, $2, $3, $4, $5)`
	res, err := d.Exec(q, transferId, userAddr, chainId, GasOnArrivalStatusNew, dropGasAmt.String())
	return sqldb.ChkExec(res, err, 1, "NewGasOnArrivalLog")
}

func (d *DAL) UpdateGasOnArrivalLogToFail(transferId string) error {
	q := `UPDATE gas_on_arrival_log 
          SET status=$2, update_time=now()
          WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, GasOnArrivalStatusFail)
	return sqldb.ChkExec(res, err, 1, "UpdateGasOnArrivalLogToFail")
}

func (d *DAL) UpdateGasOnArrivalLogToSuccess(transferId, txHash string) error {
	q := `UPDATE gas_on_arrival_log 
          SET status=$2, tx_hash=$3, update_time=now()
          WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, GasOnArrivalStatusSuccess, txHash)
	return sqldb.ChkExec(res, err, 1, "UpdateGasOnArrivalLogToSuccess")
}
