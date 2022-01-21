# Pool based cBridge flow
Code under sgn-v2 repo x/cbridge folder is like solidity contract on Ethereum, except they are run by SGN nodes. Given same input, they must behave exactly the same.

High level:
User calls source chain `send` which transfer his own asset into source chain cBridge contract, SGN will monitor and process the event and one SGN node will call relay on dest chain to send asset (after deduct fee) to receiver address.

The only way assets can be removed from cBridge contract is by sending a message with enough SGN nodes’ signatures. This includes both relay and withdraw.
## Transfer asset across chains
- User calls source chain cBridge contract send, specify token address/amount and dest chain id, receiver address etc.
- When SGN node monitor the send event, it’ll broadcast msg to other sgn nodes saying I saw a send event, do you agree
- After seeing this msg, every SGN node will do its own verify onchain and vote yes
- Now SGN nodes agree there is a user request to transfer asset, and sgn-v2 repo x/cbridge/keeper/apply.go code will be called to process this request:
  - If request is invalid due to unsupported token, bad amount or liquidity etc, we’ll save a withdraw message for user to refund later onchain by himself
  - If request can go through, code will create a relay message and ask all SGN nodes to sign the message
- Each SGN node sees the relay message will add its own signature and send to x/cbridge via SendMySig grpc call.
- When it’s a SGN node’s turn to be the syncer (ie. responsible for onchain stuff), it will periodically check if x/cbridge has relay message that has enough signatures and call relay on dest chain
## LP (Liquidity Provider) Add assets
In cBridge V2, our smart contract will lock assets from liquidity providers. They put their own assets into the system to earn reward and fee. Add assets to cbridge contract is done by calling addLiquidity onchain. If someone wants to add liquidity on different chains, he needs to switch metamask network and do it separately.
## User refund or LP withdraw
Either user’s send can’t complete or an LP wants to withdraw his asset from cbridge contract, it’s the same withdraw flow.
- user/lp requests refund or withdraw via gateway. If x/cbridge verifies request is valid, it’ll create a withdraw message for SGN nodes to sign
- SGN nodes see the withdraw msg and add its own signature, similar to the relay msg above. But SGN node won’t do anything onchain, it’s user’s job
- User will query and get the ready to submit onchain withdraw message including SGN signatures, and call onchain withdraw directly

# Pegged (customized) token bridge
Goal: Token T exists on chain A but not on chain B, and we would like to support it on chain B. 

Approach: Deploy a PeggedToken on chain B with zero initial supply, and config SGN (through gov) to mark it as 1:1 pegged to the chain A’s original token. Anyone can lock original token T on chain A’s OriginigalTokenVaults contract and mint pegged token T’ on chain B through the PeggedTokenBridge contract accordingly.
## Move chain A’s original token to chain B
- User calls deposit on chain A to lock original tokens in chain A’s vault contract.
- SGN relayers sync the deposit event to the x/pegbridge module
- x/pegbridge process the event and generate Mint proto msg for validators to sign. The mint amount will be less than the deposit amount because SGN charges fees.
- SGN syncer call mint function on chain B with the signed Mint proto msg
- User receive pegged tokens as the result of step 4, and can do anything with the newly minted tokens afterwards.
## Move chain B’s pegged token back to chain A
- User calls burn on chain B to burn the pegged token.
- SGN relayers sync the burn event to the x/pegbridge module
- x/pegbridge process the event and generate Withdraw proto msg for validators to sign.  The withdrawal amount will be less than the burn amount because SGN charges fees.
- SGN syncer call withdraw function on chain A with the signed Withdraw proto msg
## SGN delegators claim fee shares on chain A
- Delegator requests withdrawal of fee shares via gateway.
- SGN validates the request and generates co-signed Withdraw proto msg.
- Delegator get the Withdraw proto msg from gateway and call withdraw function on chain A with it.
