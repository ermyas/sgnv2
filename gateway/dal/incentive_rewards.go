package dal

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/common"
	"math"
	"math/big"
	"time"
)

type RetentionRewardsEvent struct {
	EventEndTime time.Time
	// key is level
	LevelConfig map[uint64]*RetentionRewardsLevelConfig
}

type RetentionRewardsLevelConfig struct {
	// in wei
	MaxReward         *big.Int
	MaxTransferVolume float64
}

var WeiMultiplier = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)

var RetentionRewardsConfig = map[uint64]*RetentionRewardsEvent{
	1: {
		EventEndTime: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
		LevelConfig: map[uint64]*RetentionRewardsLevelConfig{
			1: {
				MaxReward:         big.NewInt(0).Mul(WeiMultiplier, big.NewInt(1000)),
				MaxTransferVolume: 10000,
			},
			2: {
				MaxReward:         big.NewInt(0).Mul(WeiMultiplier, big.NewInt(2000)),
				MaxTransferVolume: 10000,
			},
			3: {
				MaxReward:         big.NewInt(0).Mul(WeiMultiplier, big.NewInt(3000)),
				MaxTransferVolume: 10000,
			},
		},
	},

	2: {
		EventEndTime: time.Date(2021, time.December, 6, 9, 0, 0, 0, time.UTC),
		LevelConfig: map[uint64]*RetentionRewardsLevelConfig{
			1: {
				MaxReward:         big.NewInt(0).Mul(WeiMultiplier, big.NewInt(1500)),
				MaxTransferVolume: 10000,
			},
			2: {
				MaxReward:         big.NewInt(0).Mul(WeiMultiplier, big.NewInt(2500)),
				MaxTransferVolume: 10000,
			},
			3: {
				MaxReward:         big.NewInt(0).Mul(WeiMultiplier, big.NewInt(3500)),
				MaxTransferVolume: 10000,
			},
		},
	},
}

type FeeRebateEvent struct {
	EventStartTime time.Time
	EventEndTime   time.Time
	// LevelDivisionUpperbound[0] <= tx volume <= LevelDivisionUpperbound[1] is level 1, LevelDivisionUpperbound[1] <= tx volume <= LevelDivisionUpperbound[2] is level 2
	LevelDivisionUpperbound []float64
	// key is level
	LevelConfig map[uint64]*FeeRebateLevelConfig
}

type FeeRebateLevelConfig struct {
	RebatePortion float64
	// in wei
	MaxReward *big.Int
}

var FeeRebateConfig = map[uint64]*FeeRebateEvent{
	10000: {
		EventStartTime:          time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
		EventEndTime:            time.Date(2021, time.December, 9, 0, 0, 0, 0, time.UTC),
		LevelDivisionUpperbound: []float64{100, 1000, 10000},
		LevelConfig: map[uint64]*FeeRebateLevelConfig{
			1: {
				RebatePortion: 0.1,
				MaxReward:     big.NewInt(0).Mul(WeiMultiplier, big.NewInt(1500)),
			},
			2: {
				RebatePortion: 0.2,
				MaxReward:     big.NewInt(0).Mul(WeiMultiplier, big.NewInt(2000)),
			},
			3: {
				RebatePortion: 0.3,
				MaxReward:     big.NewInt(0).Mul(WeiMultiplier, big.NewInt(2500)),
			},
		},
	},

	10001: {
		EventStartTime:          time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
		EventEndTime:            time.Date(2021, time.December, 7, 0, 0, 0, 0, time.UTC),
		LevelDivisionUpperbound: []float64{200, 2000, 20000},
		LevelConfig: map[uint64]*FeeRebateLevelConfig{
			1: {
				RebatePortion: 0.2,
				MaxReward:     big.NewInt(0).Mul(WeiMultiplier, big.NewInt(1500)),
			},
			2: {
				RebatePortion: 0.4,
				MaxReward:     big.NewInt(0).Mul(WeiMultiplier, big.NewInt(1500)),
			},
			3: {
				RebatePortion: 0.6,
				MaxReward:     big.NewInt(0).Mul(WeiMultiplier, big.NewInt(1500)),
			},
		},
	},
}

// use RetentionRewardEventId 1 for first time launch
var RetentionRewardEventId uint64 = 2

// use FeeRebateEventId 10000 for first time launch
var FeeRebateEventId uint64 = 10000

func (d *DAL) GetRetentionRewardsRecord(addr string, eventId uint64) (
	level uint64, rewardAmt *big.Int, claimTime time.Time, signature []byte, found bool, error error) {
	var amt string
	q := `SELECT group_level, reward_amt, claim_time, signature
          FROM retention_rewards_log
          WHERE usr_addr = $1
          and event_id = $2`
	err := d.QueryRow(q, addr, eventId).Scan(&level, &amt, &claimTime, &signature)
	found, err = sqldb.ChkQueryRow(err)
	rewardAmt, b := big.NewInt(0).SetString(amt, 10)
	if !b {
		rewardAmt = big.NewInt(0)
	}
	return level, rewardAmt, claimTime, signature, found, err
}

func (d *DAL) UpdateRetentionRewardsRecord(addr string, eventId uint64, rewardAmt *big.Int, signature []byte) error {
	q := `UPDATE retention_rewards_log
          SET reward_amt = $1,
              claim_time = $2,
              signature = $3
          WHERE usr_addr = $4
          and event_id = $5
          and reward_amt = '0'`
	res, err := d.Exec(q, rewardAmt.String(), now(), signature, addr, eventId)
	return sqldb.ChkExec(res, err, 1, "UpdateRetentionRewardsRecord")
}

func (d *DAL) GetFeeRebateRecord(addr string, eventId uint64) (
	rewardAmt *big.Int, rebatePortion float64, claimTime time.Time, signature []byte, totalFee float64, found bool, error error) {
	var amt string
	q := `SELECT reward_amt, rebate_portion, claim_time, signature, total_fee
          FROM fee_rebate_log
          WHERE usr_addr = $1
          and event_id = $2`
	err := d.QueryRow(q, addr, eventId).Scan(&amt, &rebatePortion, &claimTime, &signature, &totalFee)
	found, err = sqldb.ChkQueryRow(err)
	rewardAmt, b := big.NewInt(0).SetString(amt, 10)
	if !b {
		rewardAmt = big.NewInt(0)
	}
	return rewardAmt, rebatePortion, claimTime, signature, totalFee, found, err
}

func (d *DAL) AddFeeRebateFee(transferId string) {
	transfer, found, err := d.GetTransfer(transferId)
	if err != nil || !found {
		log.Errorln("not possible finding transfer:", transferId)
		return
	}
	srcToken, _, _ := d.GetTokenBySymbol(transfer.TokenSymbol, transfer.SrcChainId)
	dstToken, _, _ := d.GetTokenBySymbol(transfer.TokenSymbol, transfer.DstChainId)
	srcAmt := rmAmtDec(transfer.SrcAmt, int(srcToken.GetToken().Decimal))
	dstAmt := rmAmtDec(transfer.DstAmt, int(dstToken.GetToken().Decimal))
	if dstAmt < srcAmt {
		price, err := d.GetUsdPrice(transfer.TokenSymbol)
		if err != nil {
			log.Errorln("not possible GetUsdPrice:", transfer.TokenSymbol)
			return
		}
		feeUsdPrice := (srcAmt - dstAmt) * price
		event := FeeRebateConfig[FeeRebateEventId]
		if time.Now().Before(event.EventStartTime) || time.Now().After(event.EventEndTime) {
			return
		}
		err = d.upsertFeeRebateRecord(transfer.UsrAddr, FeeRebateEventId, feeUsdPrice)
		if err != nil {
			log.Errorln("failed to upsertFeeRebateRecord:", err)
		}
	}
}

func rmAmtDec(amt string, decimal int) float64 {
	f, _ := new(big.Float).Quo(new(big.Float).SetInt(common.Str2BigInt(amt)), big.NewFloat(math.Pow10(decimal))).Float64()
	return f
}

func (d *DAL) upsertFeeRebateRecord(addr string, eventId uint64, fee float64) error {
	q := `insert into fee_rebate_log
          (usr_addr, event_id, fee_rebate_log.total_fee)
          values($1, $2, $3)
          on conflict (usr_addr, event_id)
          DO UPDATE
	      SET fee_rebate_log.total_fee = fee_rebate_log.total_fee + $3`
	res, err := d.Exec(q, addr, eventId, fee)
	return sqldb.ChkExec(res, err, 1, "upsertFeeRebateRecord")
}

func (d *DAL) ClaimFeeRebateRecord(addr string, eventId uint64, rewardAmt *big.Int, rebatePortion float64,
	signature []byte) error {
	q := `update fee_rebate_log
          set rebate_portion = $1,
              reward_amt = $2,
              claim_time = $3,
              signature = $4
          where usr_addr = $5
          and event_id = $6`
	res, err := d.Exec(q, rebatePortion, rewardAmt.String(), now(), signature, addr, eventId)
	return sqldb.ChkExec(res, err, 1, "insertFeeRebateRecord")
}
