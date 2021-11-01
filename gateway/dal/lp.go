package dal

import (
	"github.com/celer-network/goutils/log"
	"time"

	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (d *DAL) InsertLPWithSeqNumAndMethodType(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum, methodType uint64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num, withdraw_method_type)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum, methodType)
	if err != nil {
		log.Errorf("db err:%+v", err)
	}
	return sqldb.ChkExec(res, err, 1, "UpsertLPWithMethodType")
}

func (d *DAL) UpsertLPWithSeqNum(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (usr_addr, chain_id, seq_num, lp_type) DO UPDATE
	SET status = $9, tx_hash=$6, update_time = $7`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum)
	if err != nil {
		log.Errorf("db err:%+v", err)
	}
	return sqldb.ChkExec(res, err, 1, "UpsertLPWithSeqNum")
}

func (d *DAL) UpsertLPWithTx(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (usr_addr, chain_id, tx_hash, lp_type) DO UPDATE
	SET status = $9, seq_num = $11, update_time = $7`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum)
	if err != nil {
		log.Errorf("db err, err:%+v", err)
	}
	return sqldb.ChkExec(res, err, 1, "UpsertLPWithTx")
}

func (d *DAL) UpdateWaitingForLPStatus(seqNum, lpType, chainId uint64, lpAddr, amt string, status uint64) error {
	q := `UPDATE lp SET status=$5, update_time=$6, amt=$7 WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	res, err := d.Exec(q, seqNum, chainId, lpAddr, lpType, status, now(), amt)
	if err != nil {
		log.Errorf("UpdateLPStatus error:%+v", err)
	}
	return sqldb.ChkExec(res, err, 1, "UpdateLPStatusForAdd")
}

func (d *DAL) UpdateLPStatus(seqNum, lpType, chainId uint64, lpAddr string, status uint64) error {
	q := `UPDATE lp SET status=$5, update_time=$6 WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	res, err := d.Exec(q, seqNum, chainId, lpAddr, lpType, status, now())
	if err != nil {
		log.Errorf("UpdateLPStatus error:%+v", err)
	}
	return sqldb.ChkExec(res, err, 1, "UpdateLPStatusForAdd")
}

func (d *DAL) UpdateLPStatusForWithdraw(chainId, seqNum, status uint64, lpAddr string) error {
	lpType := uint64(webapi.LPType_LP_TYPE_REMOVE)
	return d.UpdateLPStatus(seqNum, lpType, chainId, lpAddr, status)
}

func (d *DAL) GetLPInfoBySeqNum(seqNum, lpType, chainId uint64, lpAddr string) (string, uint64, time.Time, bool, error) {
	var status uint64
	var txHash string
	var ut time.Time
	q := `SELECT chain_id, tx_hash, status, update_time FROM lp WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err := d.QueryRow(q, seqNum, chainId, lpAddr, lpType).Scan(&chainId, &txHash, &status, &ut)
	found, err := sqldb.ChkQueryRow(err)
	return txHash, status, ut, found, err
}

func (d *DAL) GetLPInfoByHash(lpType, chainId uint64, lpAddr, txHash string) (uint64, uint64, time.Time, bool, error) {
	var status, seqNum uint64
	var ut time.Time
	q := `SELECT chain_id, seq_num, status, update_time FROM lp WHERE tx_hash = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err := d.QueryRow(q, txHash, chainId, lpAddr, lpType).Scan(&chainId, &seqNum, &status, &ut)
	found, err := sqldb.ChkQueryRow(err)
	return seqNum, status, ut, found, err
}

func (d *DAL) HasSeqNumUsedForWithdraw(seqNum uint64, lpAddr string) bool {
	var cnt uint64
	q := `SELECT count(1) FROM lp WHERE seq_num = $1 and usr_addr = $2 and lp_type = $3`
	err := d.QueryRow(q, seqNum, lpAddr, uint64(webapi.LPType_LP_TYPE_REMOVE)).Scan(&cnt)
	if err != nil {
		log.Errorf("run sql HasSeqNumUsedForWithdraw failed, err%+v", err)
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
	Status      types.LPHistoryStatus
	LpType      webapi.LPType
	SeqNum      uint64
	Addr        string
	MethodType  webapi.WithdrawMethodType
}

func (d *DAL) PaginateLpHistory(sender string, end time.Time, size uint64) ([]*LP, int, time.Time, error) {
	q := "SELECT chain_id, token_symbol, amt, tx_hash, create_time, status, lp_type, seq_num, usr_addr, withdraw_method_type FROM lp WHERE usr_addr = $1 and create_time < $3 and withdraw_method_type in (1,2) order by create_time desc limit $2"
	rows, err := d.Query(q, sender, size, end)
	if err != nil {
		return nil, 0, time.Unix(0, 0), err
	}
	defer closeRows(rows)

	var tps []*LP
	var txHash, tokenSymbol, amt, addr string
	var chainId, status, lpType, seqnum, methodType uint64
	var ct time.Time
	minTime := now()
	for rows.Next() {
		err = rows.Scan(&chainId, &tokenSymbol, &amt, &txHash, &ct, &status, &lpType, &seqnum, &addr, &methodType)
		if err != nil {
			return nil, 0, time.Unix(0, 0), err
		}

		tp := &LP{
			ChainId:     chainId,
			TokenSymbol: tokenSymbol,
			Amt:         amt,
			TxHash:      txHash,
			Ct:          ct,
			Status:      types.LPHistoryStatus(status),
			LpType:      webapi.LPType(lpType),
			SeqNum:      seqnum,
			Addr:        addr,
			MethodType:  webapi.WithdrawMethodType(methodType),
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
