package executor

import (
	_ "embed"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/executor/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

type DAL struct {
	*sqldb.Db
}

//go:embed schema.sql
var schema string

var Dal *DAL

func NewDAL() *DAL {
	log.Infoln("Creating DB connection")
	url := viper.GetString(types.FlagExecutorDbUrl)
	log.Infoln(url)
	db, err := sqldb.NewDb("postgres", fmt.Sprintf("postgresql://root@%s/executor?sslmode=disable", url), 4)
	if err != nil {
		log.Fatalf("Failed to create db with url %s: %+v", url, err)
	}
	log.Infoln("Syncing DB schemas")
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln("failed to initialize tables", err)
	}
	Dal = &DAL{db}
	return Dal
}

func (dal *DAL) GetExecuteContext(id []byte) *types.ExecuteRequest {
	q := `SELECT exec_ctx, status, retry_count FROM execution_context WHERE id = $1`
	var execCtxBytes []byte
	var status uint64
	var retryCount uint64
	err := dal.Db.QueryRow(q, id).Scan(&execCtxBytes, &status, &retryCount)
	if err != nil {
		log.Errorln("failed to scan result", err)
		return nil
	}
	execCtx := &msgtypes.ExecutionContext{}
	err = proto.Unmarshal(execCtxBytes, execCtx)
	if err != nil {
		log.Errorln("failed to unmarshal execution context", err)
		return nil
	}
	return &types.ExecuteRequest{
		EC:         execCtx,
		SS:         types.ExecutionStatus(status),
		RetryCount: retryCount,
	}
}

func (dal *DAL) GetExecutionContextsToExecute() []*types.ExecuteRequest {
	q := `SELECT exec_ctx, status, retry_count FROM execution_context WHERE status in ($1, $2)`
	rows, err := dal.Db.Query(q, types.ExecutionStatus_Unexecuted, types.ExecutionStatus_Init_Refund_Executed)
	if err != nil {
		log.Errorf("failed to get execution context with status %d and %d: %v", types.ExecutionStatus_Unexecuted,
			types.ExecutionStatus_Init_Refund_Executed, err)
	}
	requests := make([]*types.ExecuteRequest, 0)
	for rows.Next() {
		var execCtxBytes []byte
		var status uint64
		var retryCount uint64
		err = rows.Scan(&execCtxBytes, &status, &retryCount)
		if err != nil {
			log.Errorln("failed to scan result", err)
			continue
		}
		execCtx := &msgtypes.ExecutionContext{}
		err = proto.Unmarshal(execCtxBytes, execCtx)
		if err != nil {
			log.Errorln("failed to unmarshal execution context", err)
			continue
		}
		request := &types.ExecuteRequest{
			EC:         execCtx,
			SS:         types.ExecutionStatus(status),
			RetryCount: retryCount,
		}
		requests = append(requests, request)
	}
	return requests
}

func (dal *DAL) SaveExecutionContexts(execCtxs []*msgtypes.ExecutionContext) {
	errmsg := "could not save execution context:"
	for i := range execCtxs {
		execCtx := execCtxs[i]
		execCtxBytes, err := execCtx.Marshal()
		if err != nil {
			log.Errorln(errmsg+"failed to serialize", err)
			continue
		}
		q := `INSERT INTO execution_context (id, exec_ctx) VALUES ($1, $2) ON CONFLICT DO NOTHING`
		_, err = dal.Db.Exec(q, execCtx.MessageId, execCtxBytes)
		if err != nil {
			log.Errorln(errmsg+"failed to exec", err)
			continue
		}
	}
}

// UpdateStatus updates the execution status of a message.
// `id` is computed differently for each type of message:
// message associated with peg mint: id = mintId = hash(account, token, amount, depositor, refChainId, refId)
// message associated with peg withdraw: id = mintId, same as peg mint
// message associated with liquidity send: id = dstTransferId = hash(sender, receiver, token, amount, srcChainId, dstChainId, srcTransferId)
// message associated with liquidity withdraw: id = wdId = hash(chainid, seqnum, receiver, token, amount)
// all above ids are again hashed with bridge address id = hash(bridgeAddr, id)
// no transfer associated: id = messageId
func (dal *DAL) UpdateStatus(id []byte, status types.ExecutionStatus) error {
	q := `SELECT status FROM execution_context WHERE id = $1`
	var oldStatus uint64
	err := dal.Db.QueryRow(q, id).Scan(&oldStatus)
	if err != nil {
		return err
	}
	if uint64(status) <= oldStatus {
		log.Infof("Skipping updating execution_context (id %x) because old status %d >= new status %d", id, oldStatus, status)
		return nil
	}
	q = `UPSERT INTO execution_context (id, status) VALUES ($1, $2)`
	_, err = dal.Db.Exec(q, id, status)
	if err != nil {
		return err
	}
	log.Debugf("execution_context (id %x) status changed from %v to %v", id, types.ExecutionStatus(oldStatus), status)
	return nil
}

func (dal *DAL) RevertStatus(id []byte, status types.ExecutionStatus) error {
	if status != types.ExecutionStatus_Unexecuted && status != types.ExecutionStatus_Init_Refund_Executed {
		return fmt.Errorf("revert status to %d is forbidden", status)
	}
	log.Infof("message (id %x) status reverted to %d", status)
	q := `UPDATE execution_context SET status = $1 where id = $2`
	res, err := dal.Db.Exec(q, status, id)
	return sqldb.ChkExec(res, err, 1, "RevertStatus")
}

func (dal *DAL) IncreaseRetryCount(id []byte) (newCount uint64) {
	q := `SELECT retry_count FROM execution_context WHERE id = $1`
	var oldCount uint64
	err := dal.Db.QueryRow(q, id).Scan(&oldCount)
	if err != nil {
		log.Errorf("cannot increase message (id %x) retry count: %v", id, err)
		return 0
	}
	newCount = oldCount + 1
	q = `UPDATE execution_context SET retry_count = $1 where id = $2`
	res, err := dal.Db.Exec(q, newCount, id)
	if e := sqldb.ChkExec(res, err, 1, "IncreaseRetryCount"); e != nil {
		log.Errorf("cannot increase message (id %x) retry count: %v", id, e)
		return 0
	}
	return newCount
}

func (dal *DAL) UpdateRetryCount(id []byte, retryCount uint64) {
	q := `UPDATE execution_context SET retry_count = $1 where id = $2`
	res, err := dal.Db.Exec(q, retryCount, id)
	if e := sqldb.ChkExec(res, err, 1, "UpdateRetryCount"); e != nil {
		log.Errorf("cannot update message (id %x) retry count: %v", id, e)
	}
}

func (dal *DAL) SaveTransfer(id []byte) error {
	q := `INSERT INTO transfer (id) VALUES ($1)`
	res, err := dal.Db.Exec(q, id)
	return sqldb.ChkExec(res, err, 1, "SaveTransfer")
}

func (dal *DAL) HasTransfer(id []byte) (bool, error) {
	return dal.queryAndCheckCount("transfer", id)
}

func (dal *DAL) queryAndCheckCount(table string, id []byte) (bool, error) {
	q := fmt.Sprintf("SELECT count(1) FROM %s WHERE id = $1", table)
	var count int
	err := dal.Db.QueryRow(q, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
