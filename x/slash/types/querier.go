package types

const (
	QuerySlash        = "slash"
	QuerySlashes      = "slashes"
	QuerySlashRequest = "slash-request"
	QueryParameters   = "parameters"
)

func NewQuerySlashParams(nonce uint64) *QuerySlashParams {
	return &QuerySlashParams{
		Nonce: nonce,
	}
}
