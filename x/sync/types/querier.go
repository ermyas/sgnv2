package types

const (
	QueryParams  = "params"
	QueryUpdate  = "update"
	QueryUpdates = "updates"
)

type QueryUpdateParams struct {
	Id uint64
}

// NewQueryChangeParams creates a new instance of QueryChangeParams
func NewQueryUpdateParams(id uint64) QueryUpdateParams {
	return QueryUpdateParams{
		Id: id,
	}
}
