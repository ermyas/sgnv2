# State
keys are defined in types/keys.go
## State keys
1. original token vault, `0x01 | ChainId` -> [ContractInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/common/v1/common.proto#L17-L22)

2. pegged token bridge, `0x02 | ChainId` -> [ContractInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/common/v1/common.proto#L17-L22)

3. pair infos of original and pegged tokens, `0x03 | SrcChainId-OrigTokenAddr-DstChainId` -> [OrigPeggedPair](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L46-L71)

4. index of token infos, created by pegged token, `0x04 | DstChainId-PeggedTokenAddr` -> [PeggedOrigIndex](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L74-L80)

5. deposit info, `0x05 | DepositId` -> [DepositInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L83-L96)

6. withdraw info, `0x06 | WithdrawId` -> [WithdrawInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L145-L169)

7. mint info, `0x07 | MintId` -> [MintInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L99-L123)

8. burn info, `0x08 | BurnId` -> [BurnInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L126-L142)

9. fee claim info, `0x09 | UserAddr | Nonce` -> [FeeClaimInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L172-L177)

10. total supply amount of pegged token, `0x0a | DstChainId-PeggedTokenAddr` -> big.Int bytes

11. refund, `0x0b | DepositId/BurnId` -> [WithdrawOnChain](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L210-L239)

12. refund claim info, `0x0c | DepositId/BurnId` -> WithdrawId/MintId

13. versioned token vault, `0x11 | ChainId | Version` -> [ContractInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L25-L29)

14. versioned token bridge, `0x12 | ChainId | Version` -> [ContractInfo](https://github.com/celer-network/sgnv2/blob/d25f4280b1/proto/sgn/pegbridge/v1/pegbridge.proto#L25-L29)

15. vault version, `0x13 | ChainId | ContractAddr` -> Version

16. vault version, `0x14 | ChainId | ContractAddr` -> Version
