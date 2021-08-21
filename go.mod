module github.com/celer-network/sgn-v2

go 1.16

require (
	github.com/allegro/bigcache v1.2.1
	github.com/celer-network/goutils v0.1.32
	github.com/cosmos/cosmos-sdk v0.43.0
	github.com/ethereum/go-ethereum v1.10.7
	github.com/gammazero/deque v0.1.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/protobuf v1.26.0 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
