package types

import (
	"github.com/celer-network/sgn-v2/eth"
)

const (
	QueryParams           = "params"
	QueryPool             = "pool"
	QueryPools            = "pools"
	QueryEarnings         = "earnings"
	QueryStakeInfo        = "stake-info"
	QueryStakedPools      = "account"
	QueryAccountsStakedIn = "accounts-staked-in"
	QueryNumPools         = "num-pools"
)

// QueryPoolParams defines the params for the following queries:
// - 'custom/farming/pool'
// - 'custom/farming/accounts-staked-in'
type QueryPoolParams struct {
	PoolName string
}

// NewQueryPoolParams creates a new instance of QueryPoolParams
func NewQueryPoolParams(poolName string) QueryPoolParams {
	return QueryPoolParams{
		PoolName: poolName,
	}
}

// QueryPoolsParams defines the params for the following queries:
// - 'custom/farming/pools'
type QueryPoolsParams struct {
	Page, Limit int
}

// NewQueryPoolsParams creates a new instance of QueryPoolsParams
func NewQueryPoolsParams(page, limit int) QueryPoolsParams {
	return QueryPoolsParams{
		Page:  page,
		Limit: limit,
	}
}

// QueryPoolAccountParams defines the params for the following queries:
// - 'custom/farming/earnings'
// - 'custom/farming/stake-info'
type QueryPoolAccountParams struct {
	PoolName string
	Address  eth.Addr
}

// NewQueryPoolAccountParams creates a new instance of QueryPoolAccountParams
func NewQueryPoolAccountParams(poolName string, addr eth.Addr) QueryPoolAccountParams {
	return QueryPoolAccountParams{
		PoolName: poolName,
		Address:  addr,
	}
}

// QueryStakedPoolsParams defines the params for the following queries:
// - 'custom/farming/staked-pools'
type QueryStakedPoolsParams struct {
	Address eth.Addr
}

// NewQueryStakedPoolsParams creates a new instance of QueryStakedPoolsParams
func NewQueryStakedPoolsParams(addr eth.Addr) QueryStakedPoolsParams {
	return QueryStakedPoolsParams{
		Address: addr,
	}
}
