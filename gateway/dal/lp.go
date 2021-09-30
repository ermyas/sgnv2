package dal

import (
	"time"

	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (d *DAL) UpsertLP(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	q := `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (usr_addr, chain_id, seq_num, lp_type) DO UPDATE
	SET status = $9, update_time = $7`
	res, err := d.Exec(q, usrAddr, chainId, tokenSymbol, tokenAddr, amt, txHash, now(), now(), status, lpType, seqNum)
	return sqldb.ChkExec(res, err, 1, "UpsertLP")
}

func (d *DAL) UpdateLPStatus(seqNum, status uint64) error {
	q := `UPDATE lp SET status=$2, update_time=$3 WHERE seq_num=$1`
	res, err := d.Exec(q, seqNum, status, now())
	return sqldb.ChkExec(res, err, 1, "UpdateLPStatus")
}

func (d *DAL) GetLPInfo(seqNum, lpType, chainId uint64, lpAddr string) (string, uint64, bool, error) {
	var status uint64
	var txHash string
	q := `SELECT chain_id, tx_hash, status FROM lp WHERE seq_num = $1 and chain_id = $2 and usr_addr = $3 and lp_type = $4`
	err := d.QueryRow(q, seqNum, chainId, lpAddr, lpType).Scan(&chainId, &txHash, &status)
	found, err := sqldb.ChkQueryRow(err)
	return txHash, status, found, err
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
}

func (d *DAL) PaginateLpHistory(sender string, end time.Time, size uint64) ([]*LP, int, time.Time, error) {
	q := "SELECT chain_id, token_symbol, amt, tx_hash, create_time, status, lp_type, seq_num, usr_addr FROM lp WHERE usr_addr = $1 and create_time < $3 order by create_time desc limit $2"
	rows, err := d.Query(q, sender, size, end)
	if err != nil {
		return nil, 0, time.Unix(0, 0), err
	}
	defer closeRows(rows)

	var tps []*LP
	var txHash, tokenSymbol, amt, addr string
	var chainId, status, lpType, seqnum uint64
	var ct time.Time
	minTime := now()
	for rows.Next() {
		err = rows.Scan(&chainId, &tokenSymbol, &amt, &txHash, &ct, &status, &lpType, &seqnum, &addr)
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
