module github.com/celer-network/sgn-v2

go 1.16

require (
	github.com/allegro/bigcache v1.2.1
	github.com/celer-network/goutils v0.1.33
	github.com/cosmos/cosmos-sdk v0.43.0
	github.com/cosmos/go-bip39 v1.0.0
	github.com/ethereum/go-ethereum v1.10.8
	github.com/gammazero/deque v0.1.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0 // indirect
	github.com/iancoleman/strcase v0.1.0
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
