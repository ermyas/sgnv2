<!--
order: 1
-->

# State

## PendingUpdate

`PendingUpdate` ([Protobuf reference](https://github.com/celer-network/sgn-v2/blob/f9f76fb10d/proto/sgn/sync/v1/sync.proto#L20-L42)) records an update proposal waiting for validators' votes. It is identified by the `updateId` which automatically increase by one for each update.
```
0x01 | UpdateId -> PendingUpdate
0x11 -> NextUpdateId
```