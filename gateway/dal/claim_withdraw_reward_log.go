package dal

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
)

type CampaignScore struct {
	UsrAddr string
	Score   uint64
}

func (d *DAL) InsertClaimWithdrawRewardLog(addr string) error {
	q := `INSERT INTO claim_withdraw_reward_log (usr_addr, create_time)
                VALUES ($1, $2)`
	res, err := d.Exec(q, addr, now())
	return sqldb.ChkExec(res, err, 1, "insertClaimWithdrawRewardLog")
}

func (d *DAL) CalcCampaignScore() ([]*CampaignScore, error) {
	q := `select usr_addr, sum(score)
             from (select usr_addr, sum(daily_total) as score
                   from (select usr_addr, date_trunc('day', create_time), if(count(*) > 1000, 1000, count(*)) as daily_total
                         from transfer
                         where status = 5
                         group by usr_addr, 2) transfer
                   group by usr_addr
                   union
                   select usr_addr, sum(daily_total) as score
                   from (select usr_addr, date_trunc('day', create_time), if(count(*) > 20, 20, count(*)) as daily_total
                         from lp
                         where status = 4
                         group by usr_addr, 2) lp
                   group by usr_addr
                   union
                   select usr_addr, sum(daily_total) as score
                   from (select usr_addr, date_trunc('day', create_time), if(count(*) > 5, 5, count(*)) as daily_total
                         from claim_withdraw_reward_log
                         group by usr_addr, 2) claim
                   group by usr_addr) score
             group by usr_addr
             order by 2 desc`
	rows, err := d.Query(q)
	if err != nil {
		log.Errorf("db error:%v", err)
		return nil, err
	}
	defer closeRows(rows)

	var res []*CampaignScore
	var usrAddr string
	var score uint64
	for rows.Next() {
		err = rows.Scan(&usrAddr, &score)
		if err != nil {
			return nil, err
		}
		r := &CampaignScore{
			UsrAddr: usrAddr,
			Score:   score,
		}
		res = append(res, r)
	}
	return res, nil
}
