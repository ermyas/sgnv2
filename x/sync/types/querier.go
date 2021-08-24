package types

const (
	QueryParams         = "params"
	QueryPendingUpdate  = "update"
	QueryPendingUpdates = "updates"
)

type QueryPendingUpdateParams struct {
	UpdateId uint64
}

// NewQueryChangeParams creates a new instance of QueryChangeParams
func NewQueryPendingUpdateParams(id uint64) QueryPendingUpdateParams {
	return QueryPendingUpdateParams{
		UpdateId: id,
	}
}
