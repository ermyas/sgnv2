# Use aws-kms-tools to clear stuck transactions manually

1. Build binary

    ```sh
    # From the sgn-v2 directory
    go build -o $GOBIN/aws-kms-tools tools/aws-kms-tools/main.go
    ```

2. Get the current pending nonce:

    ```sh
    aws-kms-tools print-nonce --chainid <chain-id> --addr <signer-address>
    ```
    or use a more reliable private rpc endpoint
    ```sh
    aws-kms-tools print-nonce --rpc <endpoint-url> --addr <signer-address>
    ```

3. Send zero value transactions to the signer address itself one by one, up to `pending_nonce - 1`:

    ```sh
    # <key-name> is like sgnv2-prod2-<node-id>
    aws-kms-tools send-tx --region "us-west-2" --alias <key-name> --destination <signer-address> --chainid <chain-id> -nonce <nonce> --gasprice <gas-price-in-gwei>
    ```

    Increase `gasprice` if you see error like "replacement transaction underpriced".

    Increase `nonce` if you see error like "nonce too low".

    If many txs have been stuck, you can write a script `clear_stuck_txs.sh` like:
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

    **Note**: replace `--chainid <chain-id>` with private `--rpc <endpoint-url>` for better reliablity.
