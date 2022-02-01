<!--
order: 6
-->

# Events

The distribution module emits the following events:

## BeginBlocker

| Type            | Attribute Key | Attribute Value    |
|-----------------|---------------|--------------------|
| proposer_reward | amount        | {proposerReward}   |
| proposer_reward | validator     | {validatorAddress} |
| commission      | amount        | {commissionAmount} |
| commission      | validator     | {validatorAddress} |
| rewards         | amount        | {rewardAmount}     |
| rewards         | validator     | {validatorAddress} |

## Handlers

### MsgSetWithdrawAddress

| Type                 | Attribute Key    | Attribute Value      |
|----------------------|------------------|----------------------|
| set_withdraw_address | withdraw_address | {withdrawAddress}    |
| message              | module           | distribution         |
| message              | action           | set_withdraw_address |
| message              | sender           | {senderAddress}      |

### MsgWithdrawDelegatorReward

| Type    | Attribute Key | Attribute Value           |
|---------|---------------|---------------------------|
| withdraw_rewards | amount        | {rewardAmount}            |
| withdraw_rewards | validator     | {validatorAddress}        |
| message          | module        | distribution              |
| message          | action        | withdraw_delegator_reward |
| message          | sender        | {senderAddress}           |

### MsgWithdrawValidatorCommission

| Type       | Attribute Key | Attribute Value               |
|------------|---------------|-------------------------------|
| withdraw_commission | amount        | {commissionAmount}            |
| message    | module        | distribution                  |
| message    | action        | withdraw_validator_commission |
| message    | sender        | {senderAddress}               |

### MsgClaimAllStakingReward

| Type        | Attribute Key | Attribute Value           |
|-------------|---------------|---------------------------|
| claim_all_staking_reward   | delegator_address       | {delegatorAddress}             |

## Bridge Fee Claims

These events are triggered by the messages in the corresponding modules.

### x/cbridge

| Type        | Attribute Key | Attribute Value           |
|-------------|---------------|---------------------------|
| claim_cbridge_fee_share | delegator_address       | {delegatorAddress}             |

### x/pegbridge

| Type        | Attribute Key | Attribute Value           |
|-------------|---------------|---------------------------|
| claim_pegbridge_fees | delegator_address       | {delegatorAddress}             |

### x/message

| Type        | Attribute Key | Attribute Value           |
|-------------|---------------|---------------------------|
| claim_message_fees | delegator_address       | {delegatorAddress}             |
