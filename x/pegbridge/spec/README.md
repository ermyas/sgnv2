# Pegbridge

## Overview
Pegbridge module is designed to implement cross-chain token transferring in form of mint/burn instead of liquidity pool which is used by cbridge module. This would mean that all original tokens transferred by pegbridge would be locked temporally on the chain where they exist(which we call source chain), then nealy equal amount of pegged tokens(deduct transfer fee) would be minted on user specified chain(which we call destination chain). User can claim back their original tokens simply by submitting a burn transaction on destination chain, transfer fee would be deducted of course.

## Concepts

1. **[Concepts](01_concepts.md)**
   - [Original token](01_concepts.md#original-token)
   - [Pegged token](01_concepts.md#pegged-token)
   - [Original token vault](01_concepts.md#original-token-vault)
   - [Pegged token bridge](01_concepts.md#pegged-token-bridge)
   - [Workflow](01_concepts.md#workflow)
   - [Supply cap](01_concepts.md#supply-cap)
   - [Refund](01_concepts.md#refund)
2. **[State](02_state.md)**
3. **[Messages](03_messages.md)**
   - [MsgSignMint](03_messages.md#msgsignmint)
   - [MsgSignWithddraw](03_messages.md#msgsignwithdraw)
   - [MsgTriggerSignMint](03_messages.md#msgtriggersignmint)
   - [MsgTriggerSignWithdraw](03_messages.md#msgtriggersignwithdraw)
   - [MsgClaimFee](03_messages.md#msgclaimfee)
   - [MsgClaimRefund](03_messages.md#msgclaimrefund)
4. **[Events](04_events.md)**
   - [mint_to_sign](04_events.md#mint_to_sign)
   - [withdraw_to_sign](04_events.md#withdraw_to_sign)
5. **[Parameters](05_params.md)**
