# Use kms-send to clear stuck transactions manually

1. Get the current pending nonce:

```sh
curl <chain-rpc-url> \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["<signer-address>", "pending"],"id":1}'
```

2. Send zero value transactions to the signer address itself one by one, up to the pending nonce:

```sh
# From the sgn-v2 directory
go build -o $GOBIN/kms-send tools/kms-send/main.go
# <key-name> is like sgnv2-prod2-<node-id>
kms-send -a <key-name> -d <signer-address> -i <chain-id> -nonce <nonce> -gasprice <gas-price-in-gwei> -zerovalue
```

Increase `gasprice` if you see error like "replacement transaction underpriced".

Increase `nonce` if you see error like "nonce too low".
