package dal

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

type LPInfo struct {
	UsrAddr            string
	ChainId            uint64
	TokenSymbol        string
	TokenAddr          string
	Amt                string
	TxHash             string
	UpdateTime         time.Time
	CreateTime         time.Time
	Status             uint64
	LpType             uint64
	SeqNum             uint64
	WithdrawMethodType uint64
	WithdrawId         sql.NullString
}

func (d *DAL) InsertLPWithSeqNumAndMethodType(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum, methodType uint64, volume float64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num, withdraw_method_type, volume)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum, methodType, volume)
	if err != nil {
		log.Errorf("InsertLPWithSeqNumAndMethodType db err, usrAddr:%s, hash:%s, chainId:%d, seqNum:%d, lpType:%d, err:%+v", usrAddr, txHash, chainId, seqNum, lpType, err)
	}
	return sqldb.ChkExec(res, err, 1, "InsertLPWithSeqNumAndMethodType")
}

func (d *DAL) UpsertLPWithTx(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum, nonce uint64, volume float64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num, volume, nonce)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (usr_addr, chain_id, tx_hash, lp_type) DO UPDATE
	SET status = $9, seq_num = $11, update_time = $7`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum, volume, nonce)
	if err != nil {
		log.Errorf("UpsertLPWithTx db err, usrAddr:%s, hash:%s, chainId:%d, seqNum:%d, lpType:%d, err:%+v", usrAddr, txHash, chainId, seqNum, lpType, err)
	}
	return sqldb.ChkExec(res, err, 1, "UpsertLPWithTx")
}

func (d *DAL) UpdateWaitingForLPStatus(seqNum, lpType, chainId uint64, lpAddr, amt string, status uint64) error {
	q := `UPDATE lp SET status=$5, update_time=$6, amt=$7 WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	res, err := d.Exec(q, seqNum, chainId, lpAddr, lpType, status, now(), amt)
	if err != nil {
		log.Errorf("UpdateWaitingForLPStatus db err, usrAddr:%s, status:%d, chainId:%d, seqNum:%d, lpType:%d, err:%+v", lpAddr, status, chainId, seqNum, lpType, err)
	}
	return sqldb.ChkExec(res, err, 1, "UpdateWaitingForLPStatus")
}

func (d *DAL) UpdateLPStatus(seqNum, lpType, chainId uint64, lpAddr string, status uint64) error {
	q := `UPDATE lp SET status=$5, update_time=$6 WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	res, err := d.Exec(q, seqNum, chainId, lpAddr, lpType, status, now())
	if err != nil {
		log.Errorf("UpdateLPStatus db err, usrAddr:%s, chainId:%d, seqNum:%d, lpType:%d, status:%d, err:%+v", lpAddr, chainId, seqNum, lpType, status, err)
	}
	return sqldb.ChkExec(res, err, 1, "UpdateLPStatus")
}

func (d *DAL) UpdateLPStatusByWithdrawId(wdid string, status types.WithdrawStatus) error {
	q := `UPDATE lp SET status=$2, update_time=now() WHERE withdraw_id=$1`
	res, err := d.Exec(q, wdid, status)
	if err != nil {
		log.Errorf("UpdateLPStatus db err, wdid %s, status %d, err:%+v", wdid, uint64(status), err)
	}
	return sqldb.ChkExec(res, err, 1, "UpdateLPStatusByWithdrawId")
}

func (d *DAL) UpdateLPStatusForWithdraw(chainId, seqNum, status uint64, lpAddr string) error {
	lpType := uint64(webapi.LPType_LP_TYPE_REMOVE)
	return d.UpdateLPStatus(seqNum, lpType, chainId, lpAddr, status)
}

func (d *DAL) GetLPInfoBySeqNum(seqNum, lpType, chainId uint64, lpAddr string) (txHash string, status uint64, ut time.Time, found bool, err error) {
	q := `SELECT chain_id, tx_hash, status, update_time FROM lp WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err = d.QueryRow(q, seqNum, chainId, lpAddr, lpType).Scan(&chainId, &txHash, &status, &ut)
	found, err = sqldb.ChkQueryRow(err)
	return
}

func (d *DAL) GetLPInfo(seq, lptype, chid uint64, lpaddr string) (*LPInfo, bool, error) {
	l := &LPInfo{}
	q := `select 
	usr_addr,chain_id,token_symbol,token_addr,amt,tx_hash,update_time,create_time,status,lp_type,seq_num,withdraw_method_type,withdraw_id
	from lp where seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err := d.QueryRow(q, seq, chid, lpaddr, lptype).Scan(
		&l.UsrAddr, &l.ChainId, &l.TokenSymbol, &l.TokenAddr, &l.Amt, &l.TxHash, &l.UpdateTime,
		&l.CreateTime, &l.Status, &l.LpType, &l.SeqNum, &l.WithdrawMethodType, &l.WithdrawId)
	if err == sqldb.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	return l, true, nil
}

func (d *DAL) UpdateLP(chid, seq, status uint64, addr, wdid, tx string) error {
	t := uint64(webapi.LPType_LP_TYPE_REMOVE)
	q := `UPDATE lp SET status=$5, update_time=now(), withdraw_id=$6, tx_hash=$7 WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	res, err := d.Exec(q, seq, chid, addr, t, status, wdid, tx)
	if err != nil {
		return fmt.Errorf("unable to exec sql on lp with chid %d, seq %d, status %d, addr %s, wdid %s, tx %s: %s", chid, seq, status, addr, wdid, tx, err.Error())
	}
	return sqldb.ChkExec(res, err, 1, "UpdateLP")
}

func (d *DAL) ExistsLPInfoWithWithdrawId(wdid string) (bool, error) {
	cnt := 0
	q := `SELECT count(1) FROM lp WHERE withdraw_id = $1`
	err := d.QueryRow(q, wdid).Scan(&cnt)
	_, err = sqldb.ChkQueryRow(err)
	return cnt > 0, err
}

func (d *DAL) GetLPInfoByHash(lpType, chainId uint64, lpAddr, txHash string) (uint64, uint64, time.Time, bool, error) {
	var status, seqNum uint64
	var ut time.Time
	q := `SELECT chain_id, seq_num, status, update_time FROM lp WHERE tx_hash = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err := d.QueryRow(q, txHash, chainId, lpAddr, lpType).Scan(&chainId, &seqNum, &status, &ut)
	found, err := sqldb.ChkQueryRow(err)
	return seqNum, status, ut, found, err
}

func (d *DAL) GetCsInfoByHash(lpType, chainId uint64, lpAddr, txHash string) (webapi.WithdrawMethodType, string, string) {
	var wdType uint64
	var tokenSymbol, amt string
	q := `SELECT withdraw_method_type, token_symbol, amt FROM lp WHERE tx_hash = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err := d.QueryRow(q, txHash, chainId, lpAddr, lpType).Scan(&wdType, &tokenSymbol, &amt)
	found, err := sqldb.ChkQueryRow(err)
	if !found || err != nil {
		wdType = 0
	}
	return webapi.WithdrawMethodType(wdType), tokenSymbol, amt
}

func (d *DAL) HasSeqNumUsedForWithdraw(seqNum uint64, lpAddr string) bool {
	var cnt uint64
	q := `SELECT count(1) FROM lp WHERE seq_num = $1 and usr_addr = $2 and lp_type = $3`
	err := d.QueryRow(q, seqNum, lpAddr, uint64(webapi.LPType_LP_TYPE_REMOVE)).Scan(&cnt)
	if err != nil {
		return true
	}
	return cnt > 0
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

type LP struct {
	ChainId     uint64
	TokenSymbol string
	Amt         string
	TxHash      string
	Ct          time.Time
	Status      types.WithdrawStatus
	LpType      webapi.LPType
	SeqNum      uint64
	Addr        string
	MethodType  webapi.WithdrawMethodType
	Nonce       uint64
}

func (d *DAL) PaginateLpHistory(sender string, end time.Time, size uint64) ([]*LP, int, time.Time, error) {
	q := "SELECT chain_id, token_symbol, amt, tx_hash, create_time, status, lp_type, seq_num, usr_addr, withdraw_method_type,nonce FROM lp WHERE usr_addr = $1 and create_time < $3 and withdraw_method_type in (1,2) order by create_time desc limit $2"
	rows, err := d.Query(q, sender, size, end)
	if err != nil {
		return nil, 0, time.Unix(0, 0), err
	}
	defer closeRows(rows)

	var tps []*LP
	var txHash, tokenSymbol, amt, addr string
	var chainId, status, lpType, seqnum, methodType, nonce uint64
	var ct time.Time
	minTime := now()
	for rows.Next() {
		err = rows.Scan(&chainId, &tokenSymbol, &amt, &txHash, &ct, &status, &lpType, &seqnum, &addr, &methodType, &nonce)
		if err != nil {
			return nil, 0, time.Unix(0, 0), err
		}

		tp := &LP{
			ChainId:     chainId,
			TokenSymbol: tokenSymbol,
			Amt:         amt,
			TxHash:      txHash,
			Ct:          ct,
			Status:      types.WithdrawStatus(status),
			LpType:      webapi.LPType(lpType),
			SeqNum:      seqnum,
			Addr:        addr,
			MethodType:  webapi.WithdrawMethodType(methodType),
			Nonce:       nonce,
		}
		if minTime.After(ct) {
			minTime = ct
		}
		tps = append(tps, tp)
	}
	if len(tps) == 0 {
		minTime = time.Unix(0, 0)
	}

	return tps, len(tps), minTime, nil
}

func (d *DAL) GetAllLpHistoryForBalance(sender string) ([]*LP, error) {
	q := "SELECT chain_id, token_symbol, amt, lp_type, status FROM lp WHERE usr_addr = $1 and withdraw_method_type in (1,2)"
	rows, err := d.Query(q, sender)
	if err != nil {
		log.Warnf("db err:%+v", err)
		return nil, err
	}
	defer closeRows(rows)

	var tps []*LP
	var tokenSymbol, amt string
	var chainId, lpType, status uint64
	for rows.Next() {
		err = rows.Scan(&chainId, &tokenSymbol, &amt, &lpType, &status)
		if err != nil {
			return nil, err
		}

		isSubmitAdd := lpType == uint64(webapi.LPType_LP_TYPE_ADD) && status == uint64(types.WithdrawStatus_WD_SUBMITTING)
		validBalance := status != uint64(types.WithdrawStatus_WD_FAILED) && !isSubmitAdd
		if validBalance {
			tp := &LP{
				ChainId:     chainId,
				TokenSymbol: tokenSymbol,
				Amt:         amt,
				LpType:      webapi.LPType(lpType),
			}
			tps = append(tps, tp)
		}
	}

	return tps, nil
}

func (d *DAL) AllLpAmtForBalance() ([]*LP, error) {
	q := "SELECT usr_addr, chain_id, token_symbol, amt, lp_type, status FROM lp WHERE withdraw_method_type in (1,2)"
	rows, err := d.Query(q)
	if err != nil {
		log.Warnf("db err:%+v", err)
		return nil, err
	}
	defer closeRows(rows)

	var tps []*LP
	var tokenSymbol, amt, addr string
	var chainId, lpType, status uint64
	for rows.Next() {
		err = rows.Scan(&addr, &chainId, &tokenSymbol, &amt, &lpType, &status)
		if err != nil {
			return nil, err
		}

		isSubmitAdd := lpType == uint64(webapi.LPType_LP_TYPE_ADD) && status == uint64(types.WithdrawStatus_WD_SUBMITTING)
		validBalance := status != uint64(types.WithdrawStatus_WD_FAILED) && !isSubmitAdd
		if validBalance {
			tp := &LP{
				ChainId:     chainId,
				TokenSymbol: tokenSymbol,
				Amt:         amt,
				Addr:        addr,
				LpType:      webapi.LPType(lpType),
			}
			tps = append(tps, tp)
		}
	}

	return tps, nil
}

func (d *DAL) GetLpStatByTimeRange(begin, end time.Time) (float64, uint64, error) {
	var volume sql.NullFloat64
	var count sql.NullInt64
	q := "SELECT sum(volume),count(1) FROM lp WHERE create_time >= $1 and create_time < $2"
	err := d.QueryRow(q, begin, end).Scan(&volume, &count)
	if err != nil {
		return 0, 0, err
	}
	return volume.Float64, uint64(count.Int64), nil
}

func (d *DAL) GetDistinctLpAddrByTimeRange(begin, end time.Time) ([]string, error) {
	var addrs []string
	q := "SELECT distinct(usr_addr) FROM lp WHERE create_time >= $1 and create_time < $2"
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
