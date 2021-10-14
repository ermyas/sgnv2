package dal

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"time"
)

func (d *DAL) GetTransfer(transferId string) (*Transfer, bool, error) {
	q := `SELECT create_time, update_time, status, src_chain_id, dst_chain_id, src_tx_hash, dst_tx_hash, token_symbol, amt, received_amt, refund_seq_num, usr_addr FROM transfer WHERE transfer_id = $1`
	var srcTxHash, dstTxHash, tokenSymbol, srcAmt, dstAmt, usrAddr string
	var srcChainId, status, dstChainId, refundSeqNum uint64
	var ct, ut time.Time
	err := d.QueryRow(q, transferId).Scan(&ct, &ut, &status, &srcChainId, &dstChainId, &srcTxHash, &dstTxHash, &tokenSymbol, &srcAmt, &dstAmt, &refundSeqNum, &usrAddr)
	found, err := sqldb.ChkQueryRow(err)
	return &Transfer{
		TransferId:   transferId,
		SrcChainId:   srcChainId,
		DstChainId:   dstChainId,
		CT:           ct,
		UT:           ut,
		SrcTxHash:    srcTxHash,
		DstTxHash:    dstTxHash,
		Status:       types.TransferHistoryStatus(int32(status)),
		TokenSymbol:  tokenSymbol,
		SrcAmt:       srcAmt,
		DstAmt:       dstAmt,
		RefundSeqNum: refundSeqNum,
		UsrAddr:      usrAddr,
	}, found, err
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

func (d *DAL) MarkTransferSend(transferId, usrAddr, tokenSymbol, amt, receivedAmt, txHash string, srcChainId, dsChainId uint64, volume float64) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_SUBMITTING)
	q := `INSERT INTO transfer (transfer_id, usr_addr, token_symbol, amt, received_amt, src_chain_id, dst_chain_id, status, volume, create_time, update_time, src_tx_hash)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	res, err := d.Exec(q, transferId, usrAddr, tokenSymbol, amt, receivedAmt, srcChainId, dsChainId, status, volume, now(), now(), txHash)
	return sqldb.ChkExec(res, err, 1, "MarkTransferSend")
}

func (d *DAL) UpdateTransferRelayedStatus(transferId, txHash string) error {
	log.Debugf("UpdateTransferRelayedStatus transferId:%s, txHash:%s", transferId, txHash)
	status := uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE)
	q := `UPDATE transfer SET status=$2, update_time=$3, dst_tx_hash=$4 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, status, now(), txHash)
	return sqldb.ChkExec(res, err, 1, "UpdateTransferRelayedStatus")
}

func (d *DAL) UpdateTransferStatus(transferId string, status uint64) error {
	var checked bool
	switch status {
	case
		uint64(types.TransferHistoryStatus_TRANSFER_REFUNDED), // relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_FAILED):   // UpdateTransferStatusInHistory
		checked = true // final status
	case
		uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION), // send event
		uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE),     // relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED),               // UpdateTransferStatusInHistory
		uint64(types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED):       // UpdateTransferStatusInHistory
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
	TransferId   string
	SrcChainId   uint64
	DstChainId   uint64
	Status       types.TransferHistoryStatus
	SrcTxHash    string
	DstTxHash    string
	SrcAmt       string
	DstAmt       string
	TokenSymbol  string
	CT           time.Time
	UT           time.Time
	Volume       float64
	RefundSeqNum uint64
	UsrAddr      string
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

func (d *DAL) GetTransferByRefundSeqNum(seqNum uint64) (string, bool, error) {
	var transferId string
	q := `SELECT transfer_id FROM transfer WHERE refund_seq_num = $1`
	err := d.QueryRow(q, seqNum).Scan(&transferId)
	found, err := sqldb.ChkQueryRow(err)
	return transferId, found, err
}

func (d *DAL) UpsertTransferOnSend(transferId, usrAddr, tokenAddr, amt, sendTxHash string, srcChainId, dsChainId uint64) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION)
	token, tokenFound, tokenErr := GetTokenByAddr(tokenAddr, srcChainId)
	if token == nil || !tokenFound || tokenErr != nil {
		log.Errorf("token not found on send event, tokenAddr:%s, chainId:%d", tokenAddr, dsChainId)
		updateErr := d.UpdateTransferStatus(transferId, status)
		if updateErr != nil {
			log.Errorf("try update transfer status but failed for transfer:%s, status:%d", transferId, status)
		}
		return updateErr
	}
	q := `INSERT INTO transfer (transfer_id, usr_addr, token_symbol, amt, src_chain_id, dst_chain_id, status, create_time, update_time, src_tx_hash)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (transfer_id) DO UPDATE
	SET amt=$4, status= $7, update_time=$9, src_tx_hash=$10`
	res, err := d.Exec(q, transferId, usrAddr, token.Token.Symbol, amt, srcChainId, dsChainId, status, now(), now(), sendTxHash)
	return sqldb.ChkExec(res, err, 1, "UpsertTransferOnSend")
}
func (d *DAL) TransferCompleted(transferId, txHash, dstTransferId, receivedAmt string) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED)
	q := `UPDATE transfer SET dst_tx_hash=$2, status=$3, update_time=$4, dst_transfer_id=$5, received_amt=$6 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, txHash, status, now(), dstTransferId, receivedAmt)
	return sqldb.ChkExec(res, err, 1, "TransferCompleted")
}

func (d *DAL) MarkTransferRefund(transferId, txHash string) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND)
	var statusList []uint64
	if !d.CheckTransferStatusNotIn(transferId, statusList) {
		return nil
	}
	q := `UPDATE transfer SET refund_tx=$2, status=$3, update_time=$4 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, txHash, status, now())
	return sqldb.ChkExec(res, err, 1, "MarkTransferRefund")
}

func (d *DAL) MarkTransferRequestingRefund(transferId string, withdrawSeqNum uint64) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND)
	var statusList []uint64
	if !d.CheckTransferStatusNotIn(transferId, statusList) {
		return nil
	}
	q := `UPDATE transfer SET status=$2, update_time=$3, refund_seq_num=$4 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, status, now(), withdrawSeqNum)
	return sqldb.ChkExec(res, err, 1, "MarkTransferRefund")
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
	rows, err := d.Query(q, now().Add(-24*time.Hour))
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
