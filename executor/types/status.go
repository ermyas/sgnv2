package types

import (
	"fmt"

	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
)

type ExecuteRequest struct {
	EC         *msgtypes.ExecutionContext
	SS         ExecutionStatus
	RetryCount uint64
}

type ExecutionStatus uint64

const (
	ExecutionStatus_Unknown ExecutionStatus = iota
	// initial default status
	ExecutionStatus_Unexecuted

	// status branch: if the msg is the "refund" kind of message
	// executor needs to do InitWithdraw (if liq bridge) or ClaimRefund (if peg bridge) first before executing the message
	ExecutionStatus_Init_Refund_Executing
	// executor only executes the "refund" message at msgbus if the message is in this status
	ExecutionStatus_Init_Refund_Executed
	ExecutionStatus_Init_Refund_Failed

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
