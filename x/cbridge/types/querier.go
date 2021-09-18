package types

const (
	QueryParams = "params"
	QueryRelay  = "relay"
)

func NewQueryRelayParams(xrefId []byte) *QueryRelayParams {
	return &QueryRelayParams{
		XrefId: xrefId,
	}
}

type QueryRelayParams struct {
	XrefId []byte `json:"xref_id,omitempty"`
}
