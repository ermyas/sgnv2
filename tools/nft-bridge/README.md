# NFT Bridge
## New Multi-Chain Native NFT
- MCN NFT means same nft contract on all chains, only apply to new projects
- No orig concept and no deposit, only has burn/mint
- product also requires user interacts w/ MCN NFT contract directly instead of NFT bridge
- another new requirement is to fail invalid nft bridge request, so we add 2 maps into nft bridge and ops need to setup. **this also applies to orig/peg tokens**

### deploy NFTMCN
similar to NFTPeg
- make sure `NFT_BRIDGE_ADDR` is set to correct NFTBridge address, also .env must have `NFT_SYM` and `NFT_NAME` set
- run `hardhat deploy --network goerli --tags MCNNFT`

### set dest bridge map
saving per chain id NFT Bridge address in contract simplifes deposit/burn args.
`setDestBridge` set one chainid, nft bridge address. `setDestBridges` can set multiple in batch

### set dest NFT map(also required for orig/peg tokens)
for each nft contract we support on this chain, must set corresponding dest chain ids and NFT address.
`setDestNFT` and `setDestNFTs` for single and batch set

## Onchain Contracts
### deploy NFTBridge contract
- go to sgn-v2-contracts repo, make sure .env file has `MSG_BUS_ADDR=0x....`, msg bus address on different chains can be found at https://im-docs.celer.network/developer/contract-addresses-and-rpc-info#contracts
- run `hardhat deploy --network goerli --tags NFTBridge`
- record deployed contract address and set in .evn file as `NFT_BRIDGE_ADDR=0x...`, will be needed by NFTPeg token contract
- see above [set dest bridge map](#set-dest-bridge-map) to set chainid and new nft bridge address to all existing nft bridges

### deploy NFTPeg contract
- make sure `NFT_BRIDGE_ADDR` is set to correct NFTBridge address, also .env must have `NFT_SYM` and `NFT_NAME` set to the same as original NFT
- run `hardhat deploy --network goerli --tags PegNFT`

### deploy our own NFTOrig for test purpose
- only need to set `NFT_SYM` and `NFT_NAME` in .env file
- `hardhat deploy --network goerli --tags OrigNFT`

### update testnet.json on S3
file is at `s3://getcelerapp/nftbridge/testnet.json`. this json file has very simple schema: bridges is a list of nftbridge's chainid and address, nfts is a list of NFTs, each has one orig and a list of pegged. See [go code](./cfg.go#L53) for details. Note for MCN NFTs, orig is not set, only pegs.

invalidate cloudfront cache so that https://get.celer.app/nftbridge/testnet.json will return latest version. go to https://us-east-1.console.aws.amazon.com/cloudfront/v3/home?region=us-west-2#/distributions/E1IWV8QKIXDYBL/invalidations to create invalidation, path is `/nftbridge/*`

## Go executor and history server
### toml config
[nftbr_test.toml](./nftbr_test.toml) has example of all fields needed

### cockroach DB
- database and user are both `nftbr`
- tables: nftxfer for each bridging record, monitor for persistent event using goutils/eth/mon2. see [schema](./dal/schema.sql) for comments

### code behavior
- monitoring nftbridge contracts (chain id and address are from json), when see event `Sent`, insert a new row into nftxfer table
- when see event `Received`, update the row's status to complete by dest chain tx
- periodically polling sgn for nftbridges on all chains, if found msg to send, decode to figure out user, nft and tokenid info, try to find a match pending record in nftxfer table. Note if nftbridge saw the Sent event after sgn, it's possible we don't find a record first. next sgn poll will try again.
- if found a pending record, send onchain tx on dest chain and update the row's dest tx and status

### contract go binding
- in sgn-v2-contracts repo, after hardhat compile or deploy, cd to `artifacts/contracts/message/apps/NFTBridge.sol` folder and run
`jq .abi NFTBridge.json | abigen -abi - -type NFTBridge -pkg main -out bind_nftbr.go`, move bind_nftbr.go file here
- for execute message binding: (I didn't use sgn-v2/eth as this is supposed to be outside binary)
`jq .abi MessageBusReceiver.json | abigen -abi - -type MsgBusRecv -pkg main -out bind_msgbus_recv.go`

### nginx.conf
- add upstream for nftbr, optional but make later rule in server block more obvious
```
# nft bridge history server
upstream nftbr {
    server localhost:8888;
}    
```
- add new rule to server 80
```
# nft bridge related added by junda
location /nft {
    proxy_pass http://nftbr;
}
```
