<!--
order: 3
-->

# Messages

In this section we describe the processing of messages for the `slashing` module.

## SignSlash

Upon an infraction is discovered and slashing is enabled, SignSlash message is triggered to collect signatures, which will be then submitted onchain to slash the failed validator.

```go
type MsgSignSlash struct {
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Sig   []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	// bech32-encoded sgn address (sdk.AccAddress)
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}
```
