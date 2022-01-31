<!--
order: 4
-->

# Events

## Tendermint Events

The message module emits the following events:

### MessageToSign
  
Emitted to notify the validators that there is a message to sign.

## Handlers

### MsgSignMessage

| Type | Attribute Key | Attribute Value |
| ---- | ------------- | --------------- |
| hash | message_id    | {messageId}     |
