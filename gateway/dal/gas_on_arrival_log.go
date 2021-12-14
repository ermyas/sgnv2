package dal

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
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
	UpdateTime time.Time
	CreateTime time.Time
}

func (d *DAL) FindFailedGasOnArrivalLog(beginTime time.Time) ([]*GasOnArrivalLog, error) {
	q := `SELECT transfer_id, usr_addr, chain_id, status, update_time, create_time
		  FROM gas_on_arrival_log 
          WHERE create_time > $1 AND status = $2`
	rows, err := d.Query(q, beginTime, GasOnArrivalStatusFail)
	if err != nil {
		log.Errorf("db error:%v", err)
		return nil, err
	}
	defer closeRows(rows)

	var transferId, userAddr string
	var chainId, status uint64
	var createTime, updateTime time.Time
	var res []*GasOnArrivalLog
	for rows.Next() {
		err = rows.Scan(&transferId, &userAddr, &chainId, &status, &createTime, &updateTime)
		if err != nil {
			return nil, err
		}

		l := &GasOnArrivalLog{
			TransferId: transferId,
			UsrAddr:    userAddr,
			ChainId:    chainId,
			Status:     status,
			UpdateTime: createTime,
			CreateTime: updateTime,
		}
		res = append(res, l)
	}
	return res, nil
}

func (d *DAL) NewGasOnArrivalLog(transferId, userAddr string, chainId uint64) error {
	q := `INSERT INTO gas_on_arrival_log (transfer_id, usr_addr, chain_id, status)
                VALUES ($1, $2, $3, $4)`
	res, err := d.Exec(q, transferId, userAddr, chainId, GasOnArrivalStatusNew)
	return sqldb.ChkExec(res, err, 1, "NewGasOnArrivalLog")
}

func (d *DAL) UpdateGasOnArrivalLogToFail(transferId string) error {
	return d.UpdateGasOnArrivalLogToStatus(transferId, GasOnArrivalStatusFail)
}

func (d *DAL) UpdateGasOnArrivalLogToSuccess(transferId string) error {
	return d.UpdateGasOnArrivalLogToStatus(transferId, GasOnArrivalStatusSuccess)
}

func (d *DAL) UpdateGasOnArrivalLogToStatus(transferId string, toStatus int) error {
	q := `UPDATE gas_on_arrival_log 
          SET status=$2, update_time=now()
          WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, toStatus)
	return sqldb.ChkExec(res, err, 1, "UpdateGasOnArrivalLogToStatus")
}
