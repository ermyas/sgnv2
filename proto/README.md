## Generating Go bindings for Protobufs

1. Install [Starport](https://docs.starport.network/guide/install.html)
2. From the project root directory, run:

```sh
starport generate proto-go
```

# WARNING
For proto also used by solidity
- must use actual bytes instead of ascii/hex string. eg. big.Int.Bytes(), not big.Int.String()
- must make sure proto serialization is wire compatible so solidity can decode correctly

For sgn's own use eg. msg, saved struct in kv, ok to use string w/ gogo type option to avoid writing marshal/unmarshal code, with increased storage and traffic overhead