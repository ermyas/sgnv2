# Params

The cbridge module contains the following parameters:

- signer_update_duration: how often we update the signers list
- sign_again_cool_down_duration: time interval between valid `SignAgain` requests, to avoid DoS attack

| Key                 | Type             | Example       |
|---------------------|------------------|---------------|
| signerupdateduration | string (time ns) | "864000000000000" |
| signagaincooldownduration | string (time ns) | "60000000000" |
