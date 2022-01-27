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
	url := viper.GetString(FlagExecutorDbUrl)
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

func (dal *DAL) GetExecutionContextsToExecute() ([]*msgtypes.ExecutionContext, []types.ExecutionStatus) {
	q := `SELECT exec_ctx, status FROM execution_context WHERE status in ($1, $2)`
	rows, err := dal.Db.Query(q, types.ExecutionStatus_Unexecuted, types.ExecutionStatus_Init_Refund_Executed)
	if err != nil {
		log.Errorf("failed to get execution context with status %d: %v", types.ExecutionStatus_Unexecuted, err)
	}
	execCtxs := []*msgtypes.ExecutionContext{}
	statuses := []types.ExecutionStatus{}
	for rows.Next() {
		var execCtxBytes []byte
		var status uint64
		err = rows.Scan(&execCtxBytes, &status)
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
		execCtxs = append(execCtxs, execCtx)
		statuses = append(statuses, types.ExecutionStatus(status))
	}
	return execCtxs, statuses
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
	log.Debugf("execution_context (id %x) status changed from %d to %d", id, oldStatus, uint64(status))
	return nil
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
