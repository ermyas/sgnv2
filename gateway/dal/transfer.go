package dal

import (
	"database/sql"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

const (
	BridgeTypeUnknown = iota
	BridgeTypeSendRelay
	BridgeTypeDepositMint
	BridgeTypeBurnWithDraw
)

func (d *DAL) GetTransfer(transferId string) (*Transfer, bool, error) {
	q := `SELECT create_time, update_time, status, src_chain_id, dst_chain_id, src_tx_hash, dst_tx_hash, token_symbol, amt, received_amt, refund_seq_num, usr_addr, refund_tx, bridge_type FROM transfer WHERE transfer_id = $1`
	var srcTxHash, dstTxHash, tokenSymbol, srcAmt, dstAmt, usrAddr, refundTx string
	var srcChainId, status, dstChainId, refundSeqNum uint64
	var bridgeType int
	var ct, ut time.Time
	err := d.QueryRow(q, transferId).Scan(&ct, &ut, &status, &srcChainId, &dstChainId, &srcTxHash, &dstTxHash, &tokenSymbol, &srcAmt, &dstAmt, &refundSeqNum, &usrAddr, &refundTx, &bridgeType)
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
		RefundTx:     refundTx,
		BridgeType:   bridgeType,
	}, found, err
}

func (d *DAL) GetTransferBySrcTxHash(srcTxHash string, chainId uint32) (*Transfer, bool, error) {
	q := `SELECT create_time, update_time, status, src_chain_id, dst_chain_id, transfer_id, dst_tx_hash, token_symbol, amt, received_amt, refund_seq_num, usr_addr, refund_tx, bridge_type FROM transfer WHERE src_tx_hash = $1 and src_chain_id=$2`
	var transferId, dstTxHash, tokenSymbol, srcAmt, dstAmt, usrAddr, refundTx string
	var srcChainId, status, dstChainId, refundSeqNum uint64
	var bridgeType int
	var ct, ut time.Time
	err := d.QueryRow(q, srcTxHash, chainId).Scan(&ct, &ut, &status, &srcChainId, &dstChainId, &transferId, &dstTxHash, &tokenSymbol, &srcAmt, &dstAmt, &refundSeqNum, &usrAddr, &refundTx, &bridgeType)
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
		RefundTx:     refundTx,
		BridgeType:   bridgeType,
	}, found, err
}

func (d *DAL) GetTransferByDstTransferId(dstTransferId string) (*Transfer, bool, error) {
	q := `SELECT transfer_id, create_time, update_time, status, src_chain_id, dst_chain_id, src_tx_hash, dst_tx_hash, token_symbol, amt, received_amt, refund_seq_num, usr_addr, refund_tx, bridge_type FROM transfer WHERE dst_transfer_id = $1`
	var transferId, srcTxHash, dstTxHash, tokenSymbol, srcAmt, dstAmt, usrAddr, refundTx string
	var srcChainId, status, dstChainId, refundSeqNum uint64
	var bridgeType int
	var ct, ut time.Time
	err := d.QueryRow(q, dstTransferId).Scan(&transferId, &ct, &ut, &status, &srcChainId, &dstChainId, &srcTxHash, &dstTxHash, &tokenSymbol, &srcAmt, &dstAmt, &refundSeqNum, &usrAddr, &refundTx, &bridgeType)
	found, err := sqldb.ChkQueryRow(err)
	return &Transfer{
		TransferId:    transferId,
		DstTransferId: dstTransferId,
		SrcChainId:    srcChainId,
		DstChainId:    dstChainId,
		CT:            ct,
		UT:            ut,
		SrcTxHash:     srcTxHash,
		DstTxHash:     dstTxHash,
		Status:        types.TransferHistoryStatus(int32(status)),
		TokenSymbol:   tokenSymbol,
		SrcAmt:        srcAmt,
		DstAmt:        dstAmt,
		RefundSeqNum:  refundSeqNum,
		UsrAddr:       usrAddr,
		RefundTx:      refundTx,
		BridgeType:    bridgeType,
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

func (d *DAL) UpdateTransferStatus(transferId string, status uint64) error {
	if pass := checkTransferStatus(status); !pass {
		return nil
	}
	q := `UPDATE transfer SET status=$2, update_time=$3 WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, status, now())
	return sqldb.ChkExec(res, err, 1, "UpdateTransferStatus")
}

func (d *DAL) UpdateTransferStatusByFrom(transferId string, from, to int32) error {
	q := `UPDATE transfer SET status=$1, update_time=$2 WHERE transfer_id=$3 and status=$4`
	_, err := d.Exec(q, to, now(), transferId, from)
	return err
}

func (d *DAL) UpdateTransferForRefund(transferId string, status uint64, refundId string, refundTx string) error {
	q := `UPDATE transfer SET status=$2, refund_id=$3, refund_tx=$4, update_time=now() WHERE transfer_id=$1`
	res, err := d.Exec(q, transferId, status, refundId, refundTx)
	return sqldb.ChkExec(res, err, 1, "UpdateTransferForRefund")
}

func checkTransferStatus(status uint64) bool {
	var pass bool
	switch status {
	case
		uint64(types.TransferHistoryStatus_TRANSFER_REFUNDED), // relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_FAILED):   // UpdateTransferStatusInHistory
		pass = true // final status
	case
		uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION), // send event
		uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE),     // relayer event
		uint64(types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED),               // UpdateTransferStatusInHistory
		uint64(types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED),       // UpdateTransferStatusInHistory
		uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED),                    //TransferCompleted called by relayer event, and update if relay event missing
		uint64(types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND):            // 1. UpdateTransferStatusInHistory when signAgainWithdraw; 2. MarkTransferRequestingRefund
		pass = true
	case
		uint64(types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND), // MarkTransferRefund called by user
		uint64(types.TransferHistoryStatus_TRANSFER_SUBMITTING):             //MarkTransferSend called by user
		pass = false // status changed by other api
	default:
		pass = false // unknown status
	}
	return pass
}

type Transfer struct {
	TransferId    string
	DstTransferId string
	SrcChainId    uint64
	DstChainId    uint64
	Status        types.TransferHistoryStatus
	SrcTxHash     string
	DstTxHash     string
	SrcAmt        string
	DstAmt        string
	TokenSymbol   string
	CT            time.Time
	UT            time.Time
	Volume        float64
	RefundSeqNum  uint64
	UsrAddr       string
	RefundTx      string
	BridgeType    int
	FeePerc       uint32
}

func (d *DAL) PaginateTransferList(sender string, end time.Time, size uint64) ([]*Transfer, int, time.Time, error) {
	q := "SELECT transfer_id, create_time, status, src_chain_id,dst_chain_id, src_tx_hash, dst_tx_hash, token_symbol, amt, received_amt, refund_tx, bridge_type FROM transfer WHERE usr_addr = $1 and create_time < $3 order by create_time desc limit $2"
	rows, err := d.Query(q, sender, size, end)
	if err != nil {
		log.Errorf("db error:%v", err)
		return nil, 0, time.Unix(0, 0), err
	}
	defer closeRows(rows)

	var tps []*Transfer
	var transferId, srcTxHash, dstTxHash, tokenSymbol, srcAmt, dstAmt, refundTx string
	var srcChainId, status, dstChainId uint64
	var bridgeType int
	var ct time.Time
	minTime := now()
	for rows.Next() {
		err = rows.Scan(&transferId, &ct, &status, &srcChainId, &dstChainId, &srcTxHash, &dstTxHash, &tokenSymbol, &srcAmt, &dstAmt, &refundTx, &bridgeType)
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
			RefundTx:    refundTx,
			BridgeType:  bridgeType,
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

func (d *DAL) GetTransferByRefundSeqNum(chainId, seqNum uint64, addr string) (string, bool, error) {
	var transferId string
	q := `SELECT transfer_id FROM transfer WHERE src_chain_id=$1 and refund_seq_num = $2 and usr_addr=$3`
	err := d.QueryRow(q, chainId, seqNum, addr).Scan(&transferId)
	found, err := sqldb.ChkQueryRow(err)
	return transferId, found, err
}

func (d *DAL) ExistsTransferWithRefundId(refundId string) (bool, error) {
	cnt := 0
	q := `select count(1) from transfer where refund_id = $1`
	err := d.QueryRow(q, refundId).Scan(&cnt)
	return cnt > 0, err
}

func (d *DAL) UpsertTransferOnSend(transferId, usrAddr string, token *webapi.TokenInfo, amt, receivedAmt, sendTxHash string, srcChainId, dsChainId uint64, volume float64, feePerc uint32, bridgeType int) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION)
	q := `INSERT INTO transfer (transfer_id, usr_addr, token_symbol, amt, src_chain_id, dst_chain_id, status, create_time, update_time, src_tx_hash, volume, fee_perc, received_amt, bridge_type)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) ON CONFLICT (transfer_id) DO UPDATE
	SET amt=$4, update_time=$9, src_tx_hash=$10, volume=$11, fee_perc=$12`
	res, err := d.Exec(q, transferId, usrAddr, token.Token.Symbol, amt, srcChainId, dsChainId, status, now(), now(), sendTxHash, volume, feePerc, receivedAmt, bridgeType)
	return sqldb.ChkExec(res, err, 1, "UpsertTransferOnSend")
}

func (d *DAL) UpsertTransferOnRelay(transferId, dstTransferId, usrAddr string, token *webapi.TokenInfo, receivedAmt, txHash string, srcChainId, dsChainId uint64, isDelayed bool, bridgeType int) error {
	status := uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED)
	if isDelayed {
		status = uint64(types.TransferHistoryStatus_TRANSFER_DELAYED)
	}
	q := `INSERT INTO transfer (transfer_id, usr_addr, token_symbol, received_amt, src_chain_id, dst_chain_id, status, create_time, update_time, dst_tx_hash, dst_transfer_id, bridge_type)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) ON CONFLICT (transfer_id) DO UPDATE
	SET received_amt=$4, status= $7, update_time=$9, dst_tx_hash=$10, dst_transfer_id=$11`
	res, err := d.Exec(q, transferId, usrAddr, token.Token.Symbol, receivedAmt, srcChainId, dsChainId, status, now(), now(), txHash, dstTransferId, bridgeType)
	return sqldb.ChkExec(res, err, 1, "UpsertTransferOnRelay")
}

func (d *DAL) UpdateTransferStatusByDstTransferId(dstXferId string, status types.TransferHistoryStatus, txHash string) error {
	q := `UPDATE transfer SET status=$2, dst_tx_hash=$3, update_time=now() WHERE dst_transfer_id=$1`
	res, err := d.Exec(q, dstXferId, uint64(status), txHash)
	return sqldb.ChkExec(res, err, 1, "UpdateTransferStatusByTransferId")
}

func (d *DAL) UpdateTransferStatusByRefundId(refundId string, status types.TransferHistoryStatus, txHash string) error {
	q := `UPDATE transfer SET status=$2, refund_tx=$3, update_time=now() WHERE refund_id=$1`
	res, err := d.Exec(q, refundId, uint64(status), txHash)
	return sqldb.ChkExec(res, err, 1, "UpdateTransferStatusByRefundId")
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

func (d *DAL) Get24hTx() ([]*Transfer, error) {
	q := "SELECT dst_chain_id, token_symbol, volume, received_amt, fee_perc FROM transfer WHERE create_time > $1"
	rows, err := d.Query(q, now().Add(-24*time.Hour))
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var tps []*Transfer
	var tokenSymbol, rcvAmt string
	var dstChainId uint64
	var volume float64
	var feePerc uint32
	for rows.Next() {
		err = rows.Scan(&dstChainId, &tokenSymbol, &volume, &rcvAmt, &feePerc)
		if err != nil {
			return nil, err
		}

		tp := &Transfer{
			DstChainId:  dstChainId,
			TokenSymbol: tokenSymbol,
			DstAmt:      rcvAmt,
			Volume:      volume,
			FeePerc:     feePerc,
		}
		tps = append(tps, tp)
	}

	return tps, nil
}

func (d *DAL) GetTxStatByTimeRange(begin, end time.Time) (float64, uint64, error) {
	var volume sql.NullFloat64
	var count sql.NullInt64
	q := "SELECT sum(volume),count(1) FROM transfer WHERE create_time >= $1 and create_time < $2"
	err := d.QueryRow(q, begin, end).Scan(&volume, &count)
	if err != nil {
		return 0, 0, err
	}
	return volume.Float64, uint64(count.Int64), nil
}

func (d *DAL) GetDistinctTransferAddrByTimeRange(begin, end time.Time) ([]string, error) {
	var addrs []string
	q := "SELECT distinct(usr_addr) FROM transfer WHERE create_time >= $1 and create_time < $2"
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

func (d *DAL) GetCompletedVolumeBetween(addr string, startTime, endTime time.Time) (float64, error) {
	var totolVolume sql.NullFloat64
	status := uint64(types.TransferHistoryStatus_TRANSFER_COMPLETED)
	q := `SELECT sum(volume)
		  FROM transfer
		  WHERE usr_addr = $1
          and status = $2
		  and update_time between $3 and $4`
	err := d.QueryRow(q, addr, status, startTime, endTime).Scan(&totolVolume)
	if err != nil {
		return 0, err
	}
	if totolVolume.Valid {
		return totolVolume.Float64, nil
	}
	return 0, nil
}

func (d *DAL) GetTransfersWithStatus(status types.TransferHistoryStatus, startTime, endTime time.Time) ([]*utils.StatusAlertInfo, error) {
	q := "SELECT src_chain_id, src_tx_hash, update_time, bridge_type FROM transfer WHERE status=$1 AND update_time > $2 AND update_time < $3"
	rows, err := d.Query(q, uint64(status), startTime, endTime)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	var tps []*utils.StatusAlertInfo
	var srcTxHash string
	var srcChainId, bridgeType uint64
	var ut time.Time
	for rows.Next() {
		err = rows.Scan(&srcChainId, &srcTxHash, &ut, &bridgeType)
		if err != nil {
			return nil, err
		}

		tp := &utils.StatusAlertInfo{
			ChainId:  srcChainId,
			TxHash:   srcTxHash,
			Ut:       ut,
			IsPegged: bridgeType == 2 || bridgeType == 3,
		}
		tps = append(tps, tp)
	}
	return tps, nil
}
