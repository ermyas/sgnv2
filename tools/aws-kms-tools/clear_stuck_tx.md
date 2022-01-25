# Use aws-kms-tools to clear stuck transactions manually

1. Get the current pending nonce:

```sh
curl <chain-rpc-url> \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["<signer-address>", "pending"],"id":1}'
```

2. Send zero value transactions to the signer address itself one by one, up to `pending_nonce - 1`:

```sh
# From the sgn-v2 directory
go build -o $GOBIN/aws-kms-tools tools/aws-kms-tools/main.go
# <key-name> is like sgnv2-prod2-<node-id>
aws-kms-tools send-tx --region "us-west-2" --alias <key-name> --destination <signer-address> --chainid <chain-id> -nonce <nonce> --gasprice <gas-price-in-gwei>
```

Increase `gasprice` if you see error like "replacement transaction underpriced".

Increase `nonce` if you see error like "nonce too low".

(Optional) 3. If many txs have been stuck, you can write a script `clear_stuck_txs.sh` like:

```sh
#!/bin/bash
for nonce in $(eval echo {$1..$2})
do
   aws-kms-tools --region "us-west-2" --alias <key-name> --destination <signer-address> --chainid <chain-id> -nonce $nonce --gasprice <gas-price-in-gwei>
done
```

Run with:

```sh
chmod +x clear_stuck_txs.sh
./clear_stuck_txs <last_confirmed_nonce + 1> <pending_nonce - 1>
```
