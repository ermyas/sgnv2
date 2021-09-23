package dal

import (
	"github.com/celer-network/goutils/sqldb"
)

func (d *DAL) GetTransfer(transferId string) (string, string, uint64, uint64, bool, error) {
	var usrAddr, tokenSymbol string
	var srcChainId, dsChainId uint64
	q := `SELECT usr_addr, token_symbol, src_chain_id, dst_chain_id FROM transfer WHERE transfer_id = $1`
	err := d.QueryRow(q, transferId).Scan(&usrAddr, &tokenSymbol, &srcChainId, &dsChainId)
	found, err := sqldb.ChkQueryRow(err)
	return usrAddr, tokenSymbol, srcChainId, dsChainId, found, err
}

func (d *DAL) MarkTransfer(transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt string, srcChainId, dsChainId, status uint64, volume float64) error {
	q := `INSERT INTO transfer (transfer_id, dst_transfer_id, usr_addr, token_symbol, amt, received_amt, src_chain_id, dst_chain_id, status, volume, create_time, update_time)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	res, err := d.Exec(q, transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, srcChainId, dsChainId, status, volume, now(), now())
	return sqldb.ChkExec(res, err, 1, "MarkTransfer")
}

func (d *DAL) UpsertTransferOnSend(transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, sendTx string, srcChainId, dsChainId, status uint64, volume float64) error {
	q := `INSERT INTO transfer (transfer_id, dst_transfer_id, usr_addr, token_symbol, amt, received_amt, src_chain_id, dst_chain_id, status, volume, create_time, update_time, src_tx_hash)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ON CONFLICT (transfer_id) DO UPDATE
	SET received_amt=$6, status= $7, update_time=$12, src_tx_hash=$13`
	res, err := d.Exec(q, transferId, dstTransferId, usrAddr, tokenSymbol, amt, receivedAmt, srcChainId, dsChainId, status, volume, now(), now(), sendTx)
	return sqldb.ChkExec(res, err, 1, "MarkTransfer")
}

func (d *DAL) UpdateTransferOnRelay(dstTransferId, receivedAmt, relayTx string, status uint64) error {
	q := `UPDATE transfer set received_amt=$2, dst_tx_hash=$3, status= $4, update_time=$5 WHERE dst_transfer_id=$1`
	res, err := d.Exec(q, dstTransferId, receivedAmt, relayTx, status, now())
	return sqldb.ChkExec(res, err, 1, "InsertTransfer")
}
