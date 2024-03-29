package relayer

import (
	"encoding/json"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	slashingcli "github.com/celer-network/sgn-v2/x/slashing/client/cli"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	mapset "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

const (
	maxSlashRetry = 5
)

func (r *Relayer) processSlashQueue() {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalSlash)) * time.Second
	log.Infoln("start process slashing queue, interval:", interval)
	for {
		time.Sleep(interval)

		if !r.isSyncer() {
			continue
		}

		var keys, vals [][]byte
		r.lock.RLock()
		iterator, err := r.db.Iterator(SlashKeyPrefix, storetypes.PrefixEndBytes(SlashKeyPrefix))
		if err != nil {
			log.Errorln("Create db iterator err", err)
			continue
		}
		for ; iterator.Valid(); iterator.Next() {
			keys = append(keys, iterator.Key())
			vals = append(vals, iterator.Value())
		}
		iterator.Close()
		r.lock.RUnlock()

		enableSlash := true
		params, err := slashingcli.QueryParams(r.Transactor.CliCtx)
		if err == nil {
			enableSlash = params.EnableSlash
		}

		for i, key := range keys {
			event := NewSlashEventFromBytes(vals[i])
			err = r.dbDelete(key)
			if err != nil {
				log.Errorln("db Delete err", err)
				continue
			}

			if !enableSlash {
				// do nothing and also don't need to requeue the deleted key
				continue
			}

			r.submitSlash(event)
		}
	}
}

func (r *Relayer) submitSlash(slashEvent SlashEvent) {
	log.Infoln("Process Slash", slashEvent.Nonce)

	used, err := r.EthClient.Contracts.Staking.SlashNonces(&bind.CallOpts{}, big.NewInt(int64(slashEvent.Nonce)))
	if err != nil {
		log.Errorln("Get slashNonces err", err)
		return
	}

	if used {
		log.Infof("Slash %d has been used", slashEvent.Nonce)
		return
	}

	slash, err := slashingcli.QuerySlash(r.Transactor.CliCtx, slashEvent.Nonce)
	if err != nil {
		log.Errorln("QuerySlash err", err)
		return
	}

	signedValidators := mapset.NewSet()
	for _, sig := range slash.Signatures {
		signedValidators.Add(sig.Signer)
	}
	pass := r.validateSigs(signedValidators)
	if !pass {
		log.Debugf("Slash %d does not have enough sigs", slashEvent.Nonce)
		r.requeueSlash(slashEvent)
		return
	}

	tx, err := r.EthClient.Transactor.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("Slash transaction %x succeeded", receipt.TxHash)
				} else {
					log.Errorf("Slash transaction %x failed", receipt.TxHash)
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Errorf("Slash transaction %x err: %s", tx.Hash(), err)
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return r.EthClient.Contracts.Staking.Slash(opts, slash.GetSlashBytes(), slash.GetSigsBytes())
		},
	)
	if err != nil {
		r.requeueSlash(slashEvent)
		log.Errorln("Slash err", err)
		return
	}
	log.Infoln("Slash tx submitted", tx.Hash().Hex())
}

func (r *Relayer) requeueSlash(slashEvent SlashEvent) {
	if slashEvent.RetryCount >= maxSlashRetry {
		log.Infof("Slash %d hits retry limit", slashEvent.Nonce)
		return
	}

	slashEvent.RetryCount = slashEvent.RetryCount + 1
	err := r.dbSet(GetSlashKey(slashEvent.Nonce), slashEvent.MustMarshal())
	if err != nil {
		log.Errorln("db Set err", err)
	}
}

type SlashEvent struct {
	Nonce      uint64 `json:"nonce"`
	RetryCount uint64 `json:"retry_count"`
}

func NewSlashEvent(nonce uint64) SlashEvent {
	return SlashEvent{
		Nonce:      nonce,
		RetryCount: 0,
	}
}

func NewSlashEventFromBytes(input []byte) SlashEvent {
	event := SlashEvent{}
	event.MustUnMarshal(input)
	return event
}

// Marshal event into json bytes
func (e SlashEvent) MustMarshal() []byte {
	res, err := json.Marshal(&e)
	if err != nil {
		panic(err)
	}

	return res
}

// Unmarshal json bytes to slash event
func (e *SlashEvent) MustUnMarshal(input []byte) {
	err := json.Unmarshal(input, e)
	if err != nil {
		panic(err)
	}
}
