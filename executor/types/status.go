package types

import "fmt"

type ExecutionStatus uint64

const (
	// initial default status
	ExecutionStatus_Unexecuted ExecutionStatus = iota

	// status branch: if the msg is the "withdraw" kind of message
	// executor needs to do InitWithdraw first before executing the message
	ExecutionStatus_WD_Executing
	// executor only executes the "refund" message at msgbus if the message is in this status
	ExecutionStatus_WD_Executed

	ExecutionStatus_Executing

	// statuses after execution at msgbus
	ExecutionStatus_Executed
	ExecutionStatus_Succeeded
	ExecutionStatus_Fallback
	ExecutionStatus_Failed
)

// txStatus is MessageReceiver's enum TxStatus
func NewExecutionStatus(txStatus uint8) (ExecutionStatus, error) {
	switch txStatus {
	case 1:
		return ExecutionStatus_Succeeded, nil
	case 2:
		return ExecutionStatus_Failed, nil
	case 3:
		return ExecutionStatus_Fallback, nil
	default:
		return ExecutionStatus_Unexecuted, fmt.Errorf("cannot map TxStatus (%d) to ExecutionStatus", txStatus)
	}
}
