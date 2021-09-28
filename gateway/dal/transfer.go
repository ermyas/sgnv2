package dal

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"time"
)

func (d *DAL) GetTransfer(transferId string) (string, string, uint64, uint64, uint64, bool, error) {
	var usrAddr, tokenSymbol string
	var srcChainId, dsChainId, status uint64
	q := `SELECT usr_addr, token_symbol, src_chain_id, dst_chain_id,status FROM transfer WHERE transfer_id = $1`
	err := d.QueryRow(q, transferId).Scan(&usrAddr, &tokenSymbol, &srcChainId, &dsChainId, &status)
	found, err := sqldb.ChkQueryRow(err)
	return usrAddr, tokenSymbol, srcChainId, dsChainId, status, found, err
}

func (d *DAL) CheckTransferStatusNotIn(transferId string, statusList []uint64) bool {
	var status uint64
	q := `SELECT status FROM transfer WHERE transfer_id = $1`
	err := d.QueryRow(q, transferId).Scan(&status)
	found, err := sqldb.ChkQueryRow(err)
	if err != nil || !found {
		return false
	}
	statusList = append(statusList,
		uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED),
		uint64(types.TransferHistoryStatus_TRANSFER_FAILED),
		uint64(types.TransferHistoryStatus_TRANSFER_REFUNDED),
	)
	for _, s := range statusList {
		if s == status {
			return false
		}
	}
	return true
}

func (d *DAL) MarkTransferSend(transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, txHash string, srcChainId, dsChainId uint64, volume float64) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_SUBMITTING)
	q := `INSERT INTO transfer (transfer_id, dst_transfer_id, usr_addr, token_symbol, amt, received_amt, src_chain_id, dst_chain_id, status, volume, create_time, update_time, src_tx_hash)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	res, err := d.Exec(q, transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, srcChainId, dsChainId, status, volume, now(), now(), txHash)
	return sqldb.ChkExec(res, err, 1, "MarkTransferSend")
}

func (d *DAL) UpdateTransferStatus(transferId string, status uint64) error {
	var checked bool
	switch status {
	case
		uint64(types.TransferHistoryStatus_TRANSFER_REFUNDED), // relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_FAILED):   // UpdateTransferStatusInHistory
		checked = true // final status
	case
		uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE), // relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED),           // UpdateTransferStatusInHistory
		uint64(types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED):   // UpdateTransferStatusInHistory
		checked = true //todo CheckTransferStatusNotIn @Aric
	case
		uint64(types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND),      // MarkTransferRequestingRefund
		uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED),              //TransferCompleted called by relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND), // MarkTransferRefund called by user
		uint64(types.TransferHistoryStatus_TRANSFER_SUBMITTING):             //MarkTransferSend called by user
		checked = false // status changed by other api
	default:
		checked = false // unknown status
	}
	if !checked {
		return nil
	}
	q := `UPDATE transfer SET status=$2, update_time=$3 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, status, now())
	return sqldb.ChkExec(res, err, 1, "UpdateTransferStatus")
}

type Transfer struct {
	TransferId  string
	SrcChainId  uint64
	DstChainId  uint64
	Status      types.TransferHistoryStatus
	SrcTxHash   string
	DstTxHash   string
	SrcAmt      string
	DstAmt      string
	TokenSymbol string
	CT          time.Time
	Volume      float64
}

func (d *DAL) PaginateTransferList(sender string, end time.Time, size uint64) ([]*Transfer, int, time.Time, error) {
	q := "SELECT transfer_id, create_time, status, src_chain_id,dst_chain_id, src_tx_hash, dst_tx_hash, token_symbol, amt, received_amt FROM transfer WHERE usr_addr = $1 and create_time < $3 order by create_time desc limit $2"
	rows, err := d.Query(q, sender, size, end)
	if err != nil {
		log.Errorf("db error:%v", err)
		return nil, 0, time.Unix(0, 0), err
	}
	defer closeRows(rows)

	var tps []*Transfer
	var transferId, srcTxHash, dstTxHash, tokenSymbol, srcAmt, dstAmt string
	var srcChainId, status, dstChainId uint64
	var ct time.Time
	minTime := now()
	for rows.Next() {
		err = rows.Scan(&transferId, &ct, &status, &srcChainId, &dstChainId, &srcTxHash, &dstTxHash, &tokenSymbol, &srcAmt, &dstAmt)
		if err != nil {
			return nil, 0, time.Unix(0, 0), err
		}

		tp := &Transfer{
			TransferId:  transferId,
			SrcChainId:  srcChainId,
			DstChainId:  dstChainId,
			CT:          ct,
			SrcTxHash:   srcTxHash,
			DstTxHash:   dstTxHash,
			Status:      types.TransferHistoryStatus(int32(status)),
			TokenSymbol: tokenSymbol,
			SrcAmt:      srcAmt,
			DstAmt:      dstAmt,
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

func (d *DAL) GetTransferBySeqNum(seqNum uint64) (string, bool, error) {
	var transferId string
	q := `SELECT transfer_id FROM transfer WHERE refund_seq_num = $1`
	err := d.QueryRow(q, seqNum).Scan(&transferId)
	found, err := sqldb.ChkQueryRow(err)
	return transferId, found, err
}
func (d *DAL) TransferCompleted(transferId, txHash string) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED)
	q := `UPDATE transfer SET dst_tx_hash=$2, status=$3, update_time=$4 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, txHash, status, now())
	return sqldb.ChkExec(res, err, 1, "TransferCompleted")
}

func (d *DAL) MarkTransferRefund(transferId, txHash string, withdrawSeqNum uint64) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND)
	var statusList []uint64
	if !d.CheckTransferStatusNotIn(transferId, statusList) {
		return nil
	}
	q := `UPDATE transfer SET refund_tx=$2, status=$3, update_time=$4, refund_seq_num=$5 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, txHash, status, now(), withdrawSeqNum)
	return sqldb.ChkExec(res, err, 1, "MarkTransferRefund")
}

func (d *DAL) MarkTransferRequestingRefund(transferId string, withdrawSeqNum uint64) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND)
	var statusList []uint64
	if !d.CheckTransferStatusNotIn(transferId, statusList) {
		return nil
	}
	q := `UPDATE transfer SET status=$3, update_time=$4, refund_seq_num=$5 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, status, now(), withdrawSeqNum)
	return sqldb.ChkExec(res, err, 1, "MarkTransferRefund")
}

func (d *DAL) UpsertTransferOnSend(transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, sendTx string, srcChainId, dsChainId, status uint64, volume float64) error {
	q := `INSERT INTO transfer (transfer_id, dst_transfer_id, usr_addr, token_symbol, amt, received_amt, src_chain_id, dst_chain_id, status, volume, create_time, update_time, src_tx_hash)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (transfer_id) DO UPDATE
	SET received_amt=$6, status= $7, update_time=$12, src_tx_hash=$13`
	res, err := d.Exec(q, transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, srcChainId, dsChainId, status, volume, now(), now(), sendTx)
	return sqldb.ChkExec(res, err, 1, "UpsertTransferOnSend")
}

func (d *DAL) UpdateTransferOnRelay(dstTransferId, receivedAmt, relayTx string, status uint64) error {
	q := `UPDATE transfer set received_amt=$2, dst_tx_hash=$3, status= $4, update_time=$5 WHERE dst_transfer_id=$1`
	res, err := d.Exec(q, dstTransferId, receivedAmt, relayTx, status, now())
	return sqldb.ChkExec(res, err, 1, "UpdateTransferOnRelay")
}

func (d *DAL) UpsertSlippageSetting(addr string, slippage uint32) error {
	q := `INSERT INTO usr_slippage (addr, slippage) VALUES ($1, $2) ON CONFLICT (addr) DO UPDATE SET slippage=$2`
	res, err := d.Exec(q, addr, slippage)
	return sqldb.ChkExec(res, err, 1, "UpsertSlippageSetting")
}

func (d *DAL) GetSlippageSetting(addr string) (uint32, bool, error) {
	var slippage uint32
	q := `SELECT slippage FROM usr_slippage WHERE addr = $1`
	err := d.QueryRow(q, addr).Scan(&slippage)
	found, err := sqldb.ChkQueryRow(err)
	return slippage, found, err
}

func (d *DAL) Get24hTx() ([]*Transfer, error) {
	q := "SELECT dst_chain_id, token_symbol, volume, received_amt FROM transfer WHERE create_time > $1"
	rows, err := d.Query(q)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var tps []*Transfer
	var tokenSymbol, rcvAmt string
	var dstChainId uint64
	var volume float64
	for rows.Next() {
		err = rows.Scan(&dstChainId, &tokenSymbol, &volume, &rcvAmt)
		if err != nil {
			return nil, err
		}

		tp := &Transfer{
			DstChainId:  dstChainId,
			TokenSymbol: tokenSymbol,
			DstAmt:      rcvAmt,
			Volume:      volume,
		}
		tps = append(tps, tp)
	}

	return tps, nil
}
