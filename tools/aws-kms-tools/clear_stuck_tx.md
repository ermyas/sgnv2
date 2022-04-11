# Use aws-kms-tools to clear stuck transactions manually

1. Build binary

    ```sh
    # From the sgn-v2 directory
    go build -o $GOBIN/aws-kms-tools tools/aws-kms-tools/main.go
    ```

2. Get the current account nonce and pending nonce:

    Use sgnd config (cbridge.toml) or default rpc
    ```sh
    aws-kms-tools print-nonce -c <chain-id> --addr <signer-address>
    ```
    or use a specific rpc endpoint
    ```sh
    aws-kms-tools print-nonce --rpc <endpoint-url> --addr <signer-address>
    ```
    If `--addr` is not provided, the address and nonce of local kms signer will be printed.

3. Send zero value transactions to the signer address itself one by one, from `account_nonce` up to `pending_nonce - 1`:

    ```sh
    aws-kms-tools send-tx -c <chain-id> --nonce <nonce> --gasprice <gas-price-gwei>
    ```
    If `--gasprice` is not provided, the auto-suggested gas price will be used.

    Increase `gasprice` if you see error like "replacement transaction underpriced".

    Increase `nonce` if you see error like "nonce too low".

    If many txs have been stuck, you can write a script `clear_stuck_txs.sh` like:
    ```sh
    #!/bin/bash
    for nonce in $(eval echo {$1..$2})
    do
        aws-kms-tools send-tx -n $nonce -c <chain-id> --gasprice <gas-price-gwei>
    done
    ```

    Run with:
    ```sh
    chmod +x clear_stuck_txs.sh
    ./clear_stuck_txs <account_nonce> <pending_nonce - 1>
    ```
