package eth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/x/staking/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrPeersNotMatch = errors.New("channel peers not match")
)

func IsBonded(stakingValidatorStatus uint8) bool {
	return stakingValidatorStatus == Bonded
}

func ParseStatus(stakingValidatorStatus uint8) sdk.BondStatus {
	switch stakingValidatorStatus {
	case Bonded:
		return sdk.Bonded
	case Unbonding:
		return sdk.Unbonding
	case Unbonded:
		return sdk.Unbonded
	}

	return sdk.Unbonded
}

// GetEventSignature accepts the string of an event signature and return the hex
func GetEventSignature(eventSigStr string) HashType {
	return crypto.Keccak256Hash([]byte(eventSigStr))
}

// GetTxSender returns the sender address of the given transaction
func GetTxSender(ec *ethclient.Client, txHashStr string) (string, error) {
	tx, _, err := ec.TransactionByHash(context.Background(), Hex2Hash(txHashStr))
	if err != nil {
		return "", fmt.Errorf("failed to get tx: %w", err)
	}
	msg, err := tx.AsMessage(ethtypes.NewEIP155Signer(tx.ChainId()), nil) //TODO: base fee
	if err != nil {
		return "", fmt.Errorf("failed to get msg: %w", err)
	}
	return Addr2Hex(msg.From()), nil
}

func GetAddressFromKeystore(ksBytes []byte) (string, error) {
	type ksStruct struct {
		Address string
	}
	var ks ksStruct
	if err := json.Unmarshal(ksBytes, &ks); err != nil {
		return "", err
	}
	return ks.Address, nil
}
