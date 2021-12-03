package explorer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/lrucache"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/spf13/viper"
)

var (
	defaultStatBeginTime = time.Unix(1626566400, 0) // 1637107200 2021-11-17 00:00:00 // 1626566400 2021-07-18 00:00:00
	transferStatDuration = time.Hour
	dailyDuration        = 24 * time.Hour
)

type ExplorerServerConfig struct {
	GatewayDbUrl       string
	ExplorerDbUrl      string
	V1DbUrl            string
	Prod2DbUrl         string
	GatewayRpcUrl      string
	GatewayProd2RpcUrl string
	GrpcPort           int
	GrpcWebPort        int
}

type explorerServer struct {
	config             *ExplorerServerConfig
	explorerDb         *DAL
	v1Db               *DAL
	gatewayDb          *dal.DAL             // prod1
	prod2Db            *dal.DAL             // prod2
	gatewayClient      *utils.GatewayClient // prod1
	gatewayProd2Client *utils.GatewayClient // prod2

	dailyLiqVolumeCache *lrucache.LRUCache
	dailyTxVolumeCache  *lrucache.LRUCache
	dailyTxCountCache   *lrucache.LRUCache

	totalStatCache *totalStatCache
}

type totalStatCache struct {
	expireTime time.Time
	totalData  *GetTotalStatsResponse
	lock       sync.RWMutex
}

func (tsc *totalStatCache) updateData(newData *GetTotalStatsResponse) {
	tsc.lock.Lock()
	defer tsc.lock.Unlock()
	tsc.totalData = newData
	tsc.expireTime = time.Now().Add(5 * time.Minute)
}

func (tsc *totalStatCache) getData() *GetTotalStatsResponse {
	tsc.lock.RLock()
	defer tsc.lock.RUnlock()
	return tsc.totalData
}

func NewExplorerServer(config *ExplorerServerConfig) (*explorerServer, error) {
	if config == nil {
		return nil, fmt.Errorf("nil config")
	}
	server := &explorerServer{
		config:              config,
		dailyLiqVolumeCache: lrucache.NewLRUCache(200, nil),
		dailyTxVolumeCache:  lrucache.NewLRUCache(200, nil),
		dailyTxCountCache:   lrucache.NewLRUCache(200, nil),
		totalStatCache:      &totalStatCache{},
	}
	var err error
	server.gatewayDb = dal.NewDAL(config.GatewayDbUrl)
	if err != nil {
		return nil, err
	}

	server.prod2Db = dal.NewDAL(config.Prod2DbUrl)
	if err != nil {
		return nil, err
	}

	server.explorerDb, err = NewDAL("postgres", fmt.Sprintf("postgresql://explorer@%s/explorer?sslmode=disable", config.ExplorerDbUrl), 20)
	if err != nil {
		return nil, err
	}

	// for testnet, we do not need v1 data source
	if config.V1DbUrl != "" {
		server.v1Db, err = NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", config.V1DbUrl), 20)
		if err != nil {
			return nil, err
		}
	}

	server.gatewayClient, err = utils.NewGatewayAPI(config.GatewayRpcUrl)
	if err != nil {
		return nil, err
	}
	server.gatewayProd2Client, err = utils.NewGatewayAPI(config.GatewayProd2RpcUrl)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (e *explorerServer) GetTotalStats(ctx context.Context, in *GetTotalStatsRequest) (*GetTotalStatsResponse, error) {
	resp := &GetTotalStatsResponse{}
	data := e.totalStatCache.getData()
	if data == nil {
		resp.Err = defaultExplorerErr("data is not ready now, please wait a moment")
		return resp, nil
	}
	resp = data
	return resp, nil
}

// it should be call be timer job, and set the data into a cache for frontend to call.
func (e *explorerServer) GetCurrentTotalStatData() (*GetTotalStatsResponse, error) {
	resp := &GetTotalStatsResponse{}
	nowTs := time.Now()
	overallBalance, err := e.gatewayClient.GetOverallBalanceInfo(context.Background())
	if err != nil {
		return nil, err
	}

	overallBalanceProd2, err := e.gatewayProd2Client.GetOverallBalanceInfo(context.Background())
	if err != nil {
		return nil, err
	}

	// TODO should get from db
	priceIdMap, err := GetUsdPrices()
	if err != nil {
		log.Errorf("fail to GetUsdPrices, err:%s", err.Error())
		return nil, err
	}
	// prod1
	for symbol, liqBalance := range overallBalance {
		id, foundId := tokenSymbolTokenIds[symbol]
		if !foundId {
			log.Errorf("can not get this symbol id:%s", symbol)
			continue
		}
		usdPrice, foundUsdPrice := priceIdMap[id]
		if !foundUsdPrice {
			log.Errorf("can not get this price symbol:%s, id:%s", symbol, id)
			continue
		}
		log.Infof("symbol:%s, usdPrice:%f", symbol, usdPrice)
		resp.TotalLiquidity += liqBalance * usdPrice
	}

	// prod2
	for symbol, liqBalance := range overallBalanceProd2 {
		id, foundId := tokenSymbolTokenIds[symbol]
		if !foundId {
			log.Errorf("can not get this symbol id:%s", symbol)
			continue
		}
		usdPrice, foundUsdPrice := priceIdMap[id]
		if !foundUsdPrice {
			log.Errorf("can not get this price symbol:%s, id:%s", symbol, id)
			continue
		}
		log.Infof("symbol:%s, usdPrice:%f", symbol, usdPrice)
		resp.TotalLiquidity += liqBalance * usdPrice
	}

	resp.TotalTransactionVolume, resp.TotalTransactionCount, err = e.explorerDb.GetSumTransferVolumeAndCountByDaily(defaultStatBeginTime, nowTs)
	if err != nil {
		return nil, err
	}

	resp.Last_24TotalTransactionVolume, resp.Last_24TotalTransactionCount, err = e.explorerDb.GetSumTransferVolumeAndCount(nowTs.Add(-25*time.Hour), nowTs)
	if err != nil {
		return nil, err
	}

	walletCount, err := e.explorerDb.GetWalletCount()
	if err != nil {
		return nil, err
	}
	resp.UniqueAddress = walletCount

	return resp, nil
}

func (e *explorerServer) GetDailyTotalLiquidity(ctx context.Context, in *GetDailyTotalLiquidityRequest) (*GetDailyTotalLiquidityResponse, error) {
	resp := &GetDailyTotalLiquidityResponse{}
	begin := fmtDailyTime(in.GetBegin())
	end := fmtDailyTime(in.GetEnd())
	paramErr := checkTimeRange(begin, end)
	if paramErr != nil {
		resp.Err = paramErr
		return resp, nil
	}

	tsNow := time.Now()
	cacheKey := fmt.Sprintf("%s-%s", begin.String(), end.String())
	val, found := e.dailyLiqVolumeCache.Get(cacheKey)
	if found {
		cacheData := val.(*GetDailyTotalLiquidityResponse)
		resp = cacheData
		return resp, nil
	}
	liqs, dbErr := e.explorerDb.GetDailyLiquidityStatByRange(begin, end)
	if dbErr != nil {
		resp.Err = defaultExplorerErr(dbErr.Error())
		return resp, nil
	}
	for _, liq := range liqs {
		resp.DailyLiquidity = append(resp.DailyLiquidity, &DailyTotalLiquidity{
			Time:           common.TsMilli(liq.begin),
			TotalLiquidity: liq.volume,
		})
	}
	// current day data is still updating, should not save in cache
	if end.Add(25 * time.Hour).Before(tsNow) {
		e.dailyLiqVolumeCache.Put(cacheKey, resp)
	}
	return resp, nil
}

func (e *explorerServer) GetDailyTransactionCount(ctx context.Context, in *GetDailyTransactionCountRequest) (*GetDailyTransactionCountResponse, error) {
	resp := &GetDailyTransactionCountResponse{}
	begin := fmtDailyTime(in.GetBegin())
	end := fmtDailyTime(in.GetEnd())
	paramErr := checkTimeRange(begin, end)
	if paramErr != nil {
		resp.Err = paramErr
		return resp, nil
	}
	tsNow := time.Now()
	cacheKey := fmt.Sprintf("%s-%s", begin.String(), end.String())
	val, found := e.dailyTxCountCache.Get(cacheKey)
	if found {
		cacheData := val.(*GetDailyTransactionCountResponse)
		resp = cacheData
		return resp, nil
	}
	txs, dbErr := e.explorerDb.GetDailyTransactionStat(begin, end)
	if dbErr != nil {
		resp.Err = defaultExplorerErr(dbErr.Error())
		return resp, nil
	}
	for _, tx := range txs {
		resp.DailyTransactionCount = append(resp.DailyTransactionCount, &DailyTransactionCount{
			Time:             common.TsMilli(tx.begin),
			TransactionCount: tx.count,
		})
	}
	if end.Add(25 * time.Hour).Before(tsNow) {
		e.dailyTxCountCache.Put(cacheKey, resp)
	}
	return resp, nil
}

func (e *explorerServer) GetDailyTransactionVolume(ctx context.Context, in *GetDailyTransactionVolumeRequest) (*GetDailyTransactionVolumeResponse, error) {
	resp := &GetDailyTransactionVolumeResponse{}
	begin := fmtDailyTime(in.GetBegin())
	end := fmtDailyTime(in.GetEnd())
	paramErr := checkTimeRange(begin, end)
	if paramErr != nil {
		resp.Err = paramErr
		return resp, nil
	}
	tsNow := time.Now()
	cacheKey := fmt.Sprintf("%s-%s", begin.String(), end.String())
	val, found := e.dailyTxVolumeCache.Get(cacheKey)
	if found {
		cacheData := val.(*GetDailyTransactionVolumeResponse)
		resp = cacheData
		return resp, nil
	}
	txs, dbErr := e.explorerDb.GetDailyTransactionStat(begin, end)
	if dbErr != nil {
		resp.Err = defaultExplorerErr(dbErr.Error())
		return resp, nil
	}
	for _, tx := range txs {
		resp.DailyTransactionVolume = append(resp.DailyTransactionVolume, &DailyTransactionVolume{
			Time:                   common.TsMilli(tx.begin),
			DailyTransactionVolume: tx.volume,
		})
	}
	if end.Add(25 * time.Hour).Before(tsNow) {
		e.dailyTxVolumeCache.Put(cacheKey, resp)
	}
	return resp, nil
}

func ParseConfig(path string) (*ExplorerServerConfig, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("read explorer failure %v", err)
		return nil, err
	}
	cfg := &ExplorerServerConfig{}
	err = viper.UnmarshalKey("ExplorerConf", cfg)
	if err != nil {
		log.Errorf("parse ExplorerConf failure %v", err)
		return nil, err
	}
	return cfg, nil
}

func (e *explorerServer) StartScheduleJob() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		log.Infoln("start StartScheduleJob")
		e.processDailyLiqStat()
		processTransferStatErr := e.processTransferStat()
		if processTransferStatErr == nil {
			e.processDailyTransactionsStat()
		} else {
			log.Errorf("we find err in processTransferStat, err:%s, so we should ignore processDailyTransactionsStat", processTransferStatErr.Error())
		}
	}
}

func (e *explorerServer) StartRefreshTotalStat() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		e.refreshTotalStat()
	}
}

func (e *explorerServer) refreshTotalStat() {
	log.Infoln("start refresh total stat")
	newData, err := e.GetCurrentTotalStatData()
	if err != nil {
		log.Errorf("fail to GetCurrentTotalStatData, err:%s", err.Error())
	} else {
		e.totalStatCache.updateData(newData)
	}
}

// by default, do query hourly stat
func (e *explorerServer) processTransferStat() error {
	log.Infoln("processTransferStat")
	latestStat, found, dbErr := e.explorerDb.GetLatestTransactionStat()
	if dbErr != nil {
		log.Errorf("fail to GetLatestTransactionStat, err:%s", dbErr.Error())
		return dbErr
	}
	begin := defaultStatBeginTime
	if found {
		begin = latestStat.begin
	}
	end := begin.Add(transferStatDuration)
	for {
		// add transfer addr and lp addr
		txAddrs, dbErr := e.gatewayDb.GetDistinctTransferAddrByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetDistinctAddrByTimeRange, err:%s", dbErr.Error())
			return dbErr
		}
		for _, addr := range txAddrs {
			dbErr = e.explorerDb.InsertDistinctAddr(addr)
			if dbErr != nil {
				log.Errorf("fail to InsertDistinctAddr, err:%s", dbErr.Error())
				return dbErr
			}
		}

		lpAddrs, dbErr := e.gatewayDb.GetDistinctLpAddrByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetDistinctAddrByTimeRange, err:%s", dbErr.Error())
			return dbErr
		}
		for _, addr := range lpAddrs {
			dbErr = e.explorerDb.InsertDistinctAddr(addr)
			if dbErr != nil {
				log.Errorf("fail to InsertDistinctAddr, err:%s", dbErr.Error())
				return dbErr
			}
		}

		txAddrsProd2, dbErr := e.prod2Db.GetDistinctTransferAddrByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetDistinctAddrByTimeRange prod2, err:%s", dbErr.Error())
			return dbErr
		}
		for _, addr := range txAddrsProd2 {
			dbErr = e.explorerDb.InsertDistinctAddr(addr)
			if dbErr != nil {
				log.Errorf("fail to InsertDistinctAddr prod2, err:%s", dbErr.Error())
				return dbErr
			}
		}

		lpAddrsProd2, dbErr := e.prod2Db.GetDistinctLpAddrByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetDistinctAddrByTimeRange prod2, err:%s", dbErr.Error())
			return dbErr
		}
		for _, addr := range lpAddrsProd2 {
			dbErr = e.explorerDb.InsertDistinctAddr(addr)
			if dbErr != nil {
				log.Errorf("fail to InsertDistinctAddr prod2, err:%s", dbErr.Error())
				return dbErr
			}
		}

		if e.v1Db != nil {
			v1Addrs, dbErr := e.v1Db.GetV1DistinctTransferAddrByTimeRange(begin, end)
			if dbErr != nil {
				log.Errorf("fail to GetV1DistinctTransferAddrByTimeRange, err:%s", dbErr.Error())
				return dbErr
			}
			for _, addr := range v1Addrs {
				dbErr = e.explorerDb.InsertDistinctAddr(addr)
				if dbErr != nil {
					log.Errorf("fail to InsertDistinctAddr, err:%s", dbErr.Error())
					return dbErr
				}
			}
		}

		// add transfer and lp volume, count
		txVolume, txCount, dbErr := e.gatewayDb.GetTxStatByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetTxStatByTimeRange, err:%s", dbErr.Error())
			return dbErr
		}
		lpVolume, lpCount, dbErr := e.gatewayDb.GetLpStatByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetLpStatByTimeRange, err:%s", dbErr.Error())
			return dbErr
		}

		txVolumeProd2, txCountProd2, dbErr := e.prod2Db.GetTxStatByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetTxStatByTimeRange prod2, err:%s", dbErr.Error())
			return dbErr
		}
		lpVolumeProd2, lpCountProd2, dbErr := e.prod2Db.GetLpStatByTimeRange(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetLpStatByTimeRange prod2, err:%s", dbErr.Error())
			return dbErr
		}

		var v1Volume float64
		var v1Count uint64
		if e.v1Db != nil {
			v1Volume, v1Count, dbErr = e.v1Db.GetV1TxStatByTimeRange(begin, end)
			if dbErr != nil {
				log.Errorf("fail to GetV1TxStatByTimeRange, err:%s", dbErr.Error())
				return dbErr
			}
		}
		dbErr = e.explorerDb.InsertHourlyTransactionStat(begin, end, txVolume+lpVolume+txVolumeProd2+lpVolumeProd2+v1Volume, txCount+lpCount+txCountProd2+lpCountProd2+v1Count)
		if dbErr != nil {
			log.Errorf("fail to InsertTransactionStat, err:%s", dbErr.Error())
			return dbErr
		}
		log.Infof("finish save transfer stat, begin [%s]", begin.String())

		begin = begin.Add(transferStatDuration)
		end = end.Add(transferStatDuration)
		if begin.After(time.Now()) {
			break
		}
	}
	return nil
}

// Do daily data collect
func (e *explorerServer) processDailyLiqStat() {
	log.Infoln("processDailyLiqStat")
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(dailyDuration)
	log.Infof("today:%s, tomorrow:%s", today.String(), tomorrow.String())
	overallBalance, err := e.gatewayClient.GetOverallBalanceInfo(context.Background())
	if err != nil {
		log.Errorf("fail to GetOverallBalanceInfo, err:%s", err.Error())
		return
	}

	overallBalanceProd2, err := e.gatewayProd2Client.GetOverallBalanceInfo(context.Background())
	if err != nil {
		log.Errorf("fail to GetOverallBalanceInfo, err:%s", err.Error())
		return
	}

	priceIdMap, err := GetUsdPrices()
	if err != nil {
		log.Errorf("fail to GetUsdPrices, err:%s", err.Error())
		return
	}
	var todayLiq float64
	for symbol, liqBalance := range overallBalance {
		id, foundId := tokenSymbolTokenIds[symbol]
		if !foundId {
			log.Errorf("can not get this symbol id:%s", symbol)
			continue
		}
		usdPrice, foundUsdPrice := priceIdMap[id]
		if !foundUsdPrice {
			log.Errorf("can not get this price symbol:%s, id:%s", symbol, id)
			continue
		}
		log.Infof("symbol:%s, usdPrice:%f", symbol, usdPrice)
		todayLiq += liqBalance * usdPrice
	}

	for symbol, liqBalance := range overallBalanceProd2 {
		id, foundId := tokenSymbolTokenIds[symbol]
		if !foundId {
			log.Errorf("can not get this symbol id:%s", symbol)
			continue
		}
		usdPrice, foundUsdPrice := priceIdMap[id]
		if !foundUsdPrice {
			log.Errorf("can not get this price symbol:%s, id:%s", symbol, id)
			continue
		}
		log.Infof("symbol:%s, usdPrice:%f", symbol, usdPrice)
		todayLiq += liqBalance * usdPrice
	}

	dbErr := e.explorerDb.InsertDailyLiquidityStat(today, todayLiq)
	if dbErr != nil {
		log.Errorf("fail to InsertDailyLiquidityStat, err:%s", dbErr.Error())
		return
	}
}

func (e *explorerServer) processDailyTransactionsStat() {
	log.Infoln("processDailyTransactionsStat")
	// transfer daily is different from liq daily
	latestDailyTransactionStat, found, dbErr := e.explorerDb.GetLatestDailyTransactionStat()
	if dbErr != nil {
		log.Errorf("fail to GetLatestDailyTransactionStat, err:%s", dbErr.Error())
		return
	}
	begin := defaultStatBeginTime
	if found {
		begin = latestDailyTransactionStat.begin
	}
	end := begin.Add(dailyDuration)
	for {
		// TODO, add status check, because gateway mark can be called by anyone, we should filter fake data.
		log.Infof("transfer daily range, [%s, %s]", begin.String(), end.String())
		transferVolume, transferCount, dbErr := e.explorerDb.GetSumTransferVolumeAndCount(begin, end)
		if dbErr != nil {
			log.Errorf("fail to GetSumTransferVolumeAndCount, err:%s", dbErr.Error())
			return
		}
		dbErr = e.explorerDb.InsertDailyTransactionStat(begin, transferVolume, transferCount)
		if dbErr != nil {
			log.Errorf("fail to InsertDailyLiquidityStat, err:%s", dbErr.Error())
			return
		}
		begin = begin.Add(dailyDuration)
		end = begin.Add(dailyDuration)
		if begin.After(time.Now()) {
			break
		}
	}
}
