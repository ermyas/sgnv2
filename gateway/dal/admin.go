package dal

import (
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
)

func (d *DAL) InsertClaimWithdrawRewardLog(addr string) error {
	q := `INSERT INTO claim_withdraw_reward_log (usr_addr, create_time)
                VALUES ($1, $2)`
	res, err := d.Exec(q, addr, now())
	return sqldb.ChkExec(res, err, 1, "insertClaimWithdrawRewardLog")
}

func (d *DAL) CalcCampaignScore(begin time.Time) ([]*webapi.CampaignScore, error) {
	q := `select usr_addr, sum(score)
             from (select usr_addr, sum(daily_total) as score, 'transfer' as name
                   from (select usr_addr, if(count(*) > 5000, 5000, count(*)) as daily_total
                         from transfer
                         where status = 5 and create_time>$1 and create_time<$2
                         group by usr_addr) transfer
                   group by usr_addr
                   union
                   select usr_addr, sum(daily_total) as score, 'lp' as name
                   from (select usr_addr, if(count(*) > 1000, 1000, count(*)) as daily_total
                         from lp
                         where status = 4 and withdraw_method_type < 3 and create_time>$1 and create_time<$2
                         group by usr_addr) lp
                   group by usr_addr
                   union
                   select usr_addr, sum(daily_total) as score, 'staking_reward' as name
                   from (select usr_addr, if(count(*) > 50, 50, count(*)) as daily_total
                         from lp
                         where status = 4 and withdraw_method_type = 3 and create_time>$1 and create_time<$2 
                         group by usr_addr) lp
                   group by usr_addr
                   union
                   select usr_addr, sum(daily_total) as score， 'farming_reward' as name
                   from (select usr_addr, if(count(*) > 50, 50, count(*)) as daily_total
                         from claim_withdraw_reward_log
						 where create_time>$1 and create_time<$2
                         group by usr_addr) claim
                   group by usr_addr) score
             group by usr_addr
             order by 2 desc`
	rows, err := d.Query(q, begin, begin.Add(24*time.Hour))
	if err != nil {
		log.Errorf("db error:%v", err)
		return nil, err
	}
	defer closeRows(rows)

	var res []*webapi.CampaignScore
	var usrAddr string
	var score uint64
	for rows.Next() {
		err = rows.Scan(&usrAddr, &score)
		if err != nil {
			return nil, err
		}
		if utils.IsBot(usrAddr) {
			continue
		}
		r := &webapi.CampaignScore{
			UsrAddr: usrAddr,
			Score:   score,
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *DAL) IsAdminAddrValid(addr string) bool {
	var cnt uint64
	q := `SELECT count(1) FROM admin_addr WHERE addr = $1`
	err := d.QueryRow(q, addr).Scan(&cnt)
	if err != nil {
		log.Errorf("run sql HasAdminAddr failed, err%+v", err)
		return false
	}
	return cnt > 0
}

func (d *DAL) LatestApyUpdateTime() time.Time {
	var ct time.Time
	q := `SELECT create_time FROM apy order by create_time desc limit 1`
	err := d.QueryRow(q).Scan(&ct)
	found, err := sqldb.ChkQueryRow(err)
	if !found || err != nil {
		return time.Now().Add(-2 * time.Hour)
	}
	return ct
}

// return map[chainId]map[token]apy
func (d *DAL) GetApyList(n int64) []string {
	q := `SELECT apy FROM apy order by create_time desc limit $1`
	rows, err := d.Query(q, n)
	if err != nil {
		return nil
	}
	defer closeRows(rows)
	var apyList []string
	var apy string
	for rows.Next() {
		err = rows.Scan(&apy)
		if err != nil {
			continue
		}
		apyList = append(apyList, apy)
	}
	return apyList
}

func (d *DAL) InsertApy(apyJson string) error {
	q := `INSERT INTO apy (apy, create_time)
                VALUES ($1, $2)`
	res, err := d.Exec(q, apyJson, now())
	return sqldb.ChkExec(res, err, 1, "InsertApy")
}
