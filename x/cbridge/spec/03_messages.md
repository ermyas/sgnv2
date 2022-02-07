# Messages
defined in proto/sgn/cbridge/v1/ folder. proto comments has more details.
## cbridge.proto
- Params: governable param using common govern param flow
- CbrConfig: cbridge specific configs
- CbrPrice: per asset and chain price to calculate base fee
- RelayGasCostParam: to calculate gas cost for relay (also needed for base fee calculation)
- AssetPrice: per token USD price, used in CbrPrice
- GasPrice: gas price (ie. in wei) per chain
- ChainAsset: per symbol, per chain detail about token, eg. addr, decimal etc
- ChainPair: a pair of chains for transfer token between them. Note to avoid duplicates, chid1 must be smaller than chid2
- PerChainPairAssetOverride: as name suggests, provide a method to override chainpair param for specific asset.
- CbrProposal: update CbrConfig, eg. add new chain, token
- RelayOnChain and WithdrawOnchain: serialized bytes will be submitted to onchain contract
- XferStatus: enum for possible transfer status
- XferRelay: saved in sgn kv for relay
- WithdrawDetail: saved in sgn kv for withdraw
- Signer, ChainSigners and LatestSigners: deal with contract signers list