package eth

import "reflect"

// compare events, ignore TxIndex for now as some chain (e.g., Fantom) has inconsistent value for it.

func (ev *BridgeLiquidityAdded) Equal(e *BridgeLiquidityAdded) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *BridgeSend) Equal(e *BridgeSend) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *BridgeRelay) Equal(e *BridgeRelay) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *BridgeWithdrawDone) Equal(e *BridgeWithdrawDone) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *BridgeSignersUpdated) Equal(e *BridgeSignersUpdated) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *OriginalTokenVaultDeposited) Equal(e *OriginalTokenVaultDeposited) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *PeggedTokenBridgeMint) Equal(e *PeggedTokenBridgeMint) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *PeggedTokenBridgeBurn) Equal(e *PeggedTokenBridgeBurn) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *OriginalTokenVaultWithdrawn) Equal(e *OriginalTokenVaultWithdrawn) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *OriginalTokenVaultV2Deposited) Equal(e *OriginalTokenVaultV2Deposited) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *PeggedTokenBridgeV2Mint) Equal(e *PeggedTokenBridgeV2Mint) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *PeggedTokenBridgeV2Burn) Equal(e *PeggedTokenBridgeV2Burn) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *OriginalTokenVaultV2Withdrawn) Equal(e *OriginalTokenVaultV2Withdrawn) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *WithdrawInboxWithdrawalRequest) Equal(e *WithdrawInboxWithdrawalRequest) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *MessageBusMessageWithTransfer) Equal(e *MessageBusMessageWithTransfer) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *MessageBusMessage) Equal(e *MessageBusMessage) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}

func (ev *MessageBusExecuted) Equal(e *MessageBusExecuted) bool {
	ev.Raw.TxIndex = e.Raw.TxIndex
	return reflect.DeepEqual(ev, e)
}
