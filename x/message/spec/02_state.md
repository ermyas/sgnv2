<!--
order: 2
-->

# State

## MessageBus

`0x01 | ChainId -> MessageBusAddress`

The addresses of `MessageBus` contracts at each chain.

## Messages

`0x02 | MessageId -> ProtocolBuffer(Message)`

`Message` is a part of `ExecutionContext`. It contains important information such as routing info, signatures and the message body.

## Transfers

`0x03 | MessageId -> ProtocolBuffer(Transfer)`

`Transfer` is a part of `ExecutionContext`. If a message is sent with an accompanied transfer, along with the message, a transfer is also saved.

## ActiveMessageIds

`0x04 | DstChainId | ReceiverAppAddress -> MessageId`

A `messageId` is added to the active message id list when the on-chain events `Message` or `MessageWithTransfer` is being applied. When a message is executed on-chain, meaning the `Executed` event is received, its id is removed from the active message id list.

## MessageRefundNonce

`0x05 -> nonce`

A nonce is required by cBridge to initiate a refund for failed transfers. The message module assigns a unique nonce to every failed "MessageWithTransfer".

## MessageRefund

`0x06 | SourceTransferId -> ProtocolBuffer(ExecutionContext)`

If a message is deemed needing for refund, a refund `ExecutionContext` is built and is saved and keyed by the source transfer id. The main purposes are:

1. to allow the relayer and the message module to determine whether a transfer's refund has already applied.

2. for the extensibility in the future if it would be decided that refund `ExecutionContext`s need to be queried using source transfer id as key (allows the end users to query their refund).

## FeeClaimInfo

TODO