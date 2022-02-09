# State 
keys are defined in types/keys.go
## State keys
1. original token vault, `0x01 | ChainId` -> [ConctractInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/common/v1/common.proto#L17)
2. pegged token bridge, `0x02 | ChainId` -> [ConctractInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/common/v1/common.proto#L17)
3. pair infos of original and pegged tokens, `0x03 | SrcChainId-OrigTokenAddr-DstChainId` -> [OrigPeggedPair](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L40)
4. index of token infos, created by pegged token, `0x04 | DstChainId-PeggedTokenAddr` -> [PeggedOrigIndex](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L64)
5. deposit info, `0x05 | DepositId` -> [DepositInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L73)
6. withdraw info, `0x06 | WithdrawId` -> [WithdrawInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L126)
7. mint info, `0x07 | MintId` -> [MintInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L87)
8. burn info, `0x08 | BurnId` -> [BurnInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L112)
9. fee claim info, `0x09 | UserAddr | Nonce` -> [FeeClaimInfo](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L151)
10. total supply amount of pegged token, `0x0a | DstChainId-PeggedTokenAddr` -> big.Int bytes
11. refund, `0x0b | DepositId/BurnId` -> [WithdrawOnChain](https://github.com/celer-network/sgn-v2/blob/7083316f71a4e794c89a737cd09eb7c1ae38106f/proto/sgn/pegbridge/v1/pegbridge.proto#L187)
12. refund claim info, `0x0c | DepositId/BurnId` -> WithdrawId/MintId
