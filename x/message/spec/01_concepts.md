<!--
order: 1
-->

# Concepts

## Message & Execution Context

"Message" and "Execution Context" basically mean the same thing in this module and will be used interchangably.

Execution Context is a data structure that carries all the necessary information of a message and is used as a body of communication with systems outside of this module. This context can be queried by external parties (notably the third-party-run executor) through the `ExecutionContexts` querier.

## Active Execution Context

An execution context is said to be "active" and is thus queriable through the `ExecutionContexts` rpc if it has an `activeMessageId` **AND** if its signatures reaches quorum.

# Flows

Update proposals of on-chain events (MessageWithTransfer, Message) emitted from MessageBus are built into execution contexts and are saved to the application state. x/message assigns the execution context an `activeMessageId` and emits `EventTypeMessageToSign`. Once enough signatures are appplied, execution contexts can be queried through `ExecutionContexts` rpc.

## Message With Transfer

For event MessageWithTransfer, the module does a little bit extra work when applying an update proposal. It calls the corresponding bridge keeper (x/cbridge and x/pegbridge) to verify if the transfer is in the required state and then builds an execution context.

## Refund

If x/message finds the transfer in a failed state, it instead builds an execution contexts for a refund message. Note: a refund message is not different from a normal message but its destination info is the same as the source info. It can also be queried through the `ExecutionContexts` rpc.
