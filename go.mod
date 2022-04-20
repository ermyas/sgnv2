module github.com/celer-network/sgn-v2

go 1.16

require (
	github.com/allegro/bigcache v1.2.1
	github.com/armon/go-metrics v0.3.10
	github.com/celer-network/cbridge-flow v0.0.1
	github.com/celer-network/endpoint-proxy v0.2.3
	github.com/celer-network/goutils v0.1.53
	github.com/cockroachdb/cockroach-go/v2 v2.2.8
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/cosmos/go-bip39 v1.0.0
	github.com/deckarep/golang-set v1.8.0
	github.com/ethereum/go-ethereum v1.10.16
	github.com/gammazero/deque v0.1.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.10.2
	github.com/lthibault/jitterbug v2.0.0+incompatible
	github.com/miguelmota/go-solidity-sha3 v0.1.1
	github.com/onflow/flow-go-sdk v0.24.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
