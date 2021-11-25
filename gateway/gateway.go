package gateway

import (
	"context"
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/fee"
	gatewaysvc "github.com/celer-network/sgn-v2/gateway/svc"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 8081, "Listening port")
	rpcWebPort = flag.Int("rpcweb", 8082, "Listening port for rpc web")
	rpcPort    = flag.Int("rpc", 10000, "Listening port for rpc")
)

// CustomGRPCHeaderMatcher for mapping request headers to
// GRPC metadata.
// HTTP headers that start with 'Grpc-Metadata-' are automatically mapped to
// gRPC metadata after removing prefix 'Grpc-Metadata-'. We can use this
// CustomGRPCHeaderMatcher if headers don't start with `Grpc-Metadata-`
func CustomGRPCHeaderMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case grpctypes.GRPCBlockHeightHeader:
		return grpctypes.GRPCBlockHeightHeader, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func InitGateway(
	_homeDir string,
	_legacyAmino *codec.LegacyAmino,
	_cdc codec.Codec,
	_interfaceRegistry codectypes.InterfaceRegistry,
	_selfStart bool,
	_dbUrl string,
	_env string) {

	gatewaysvc.RootDir = _homeDir
	gatewaysvc.LegacyAmino = _legacyAmino
	gatewaysvc.Cdc = _cdc
	gatewaysvc.InterfaceRegistry = _interfaceRegistry
	gatewaysvc.SelfStart = _selfStart
	gatewaysvc.Env = _env

	flag.Parse()
	log.Infof("Starting gateway at rest:%d, grpc:%d", *port, *rpcPort)

	gs, err := gatewaysvc.NewGatewayService(_dbUrl)
	if err != nil {
		log.Fatalf("fail to init gateway server, err:%v", err)
		return
	}
	defer gs.Close()
	log.Infof(" gateway svc started, env:%s", gatewaysvc.Env)

	err = gs.InitTransactors()
	if err != nil {
		log.Fatalf("fail to init transactor in gateway server, err:%v", err)
		return
	}

	signerKey, signerPass := viper.GetString(common.FlagGatewayIncentiveRewardsKeystore), viper.GetString(common.FlagGatewayIncentiveRewardsPassphrase)
	signer, addr, err := eth.CreateSigner(signerKey, signerPass, nil)
	if err != nil {
		log.Fatalf("fail to CreateSigner in gateway server, err:%v", err)
		return
	}
	gs.S = &gatewaysvc.IncentiveRewardsSigner{
		Signer: &signer,
		Addr:   &addr,
	}

	gs.StartChainTokenPolling(1 * time.Hour)
	log.Infof("chain token cached")

	gs.F = fee.NewTokenPriceCache(gs.TP.GetTransactor())
	log.Infof(" token price cached")

	gs.StartUpdateTokenPricePolling(time.Duration(viper.GetInt32(common.FlagSgnCheckIntervalCbrPrice)) * time.Second)
	gs.StartAvgLpFeeEarningPolling(10 * time.Minute)
	gs.StartAbnormalBalanceCheckPolling(1 * time.Hour)

	// start a rpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *rpcPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	webapi.RegisterWebServer(grpcServer, gs)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// Add grpc web
	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool { return true }))
	grpcWebSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *rpcWebPort),
		Handler: wrappedServer,
	}
	go func() {
		if grpcWebSrvErr := grpcWebSrv.ListenAndServe(); grpcWebSrvErr != nil {
			log.Fatal(grpcWebSrvErr)
		}
	}()

	// The default JSON marshaller used by the gRPC-Gateway is unable to marshal non-nullable non-scalar fields.
	// Using the gogo/gateway package with the gRPC-Gateway WithMarshaler option fixes the scalar field marshalling issue.
	marshalerOption := &gateway.JSONPb{
		EmitDefaults: true,
		EnumsAsInts:  true,
		Indent:       "  ",
		OrigName:     true,
		AnyResolver:  _interfaceRegistry,
	}

	gwmux := runtime.NewServeMux(
		// Custom marshaler option is required for gogo proto
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshalerOption),

		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),

		// Custom header matcher for mapping request headers to
		// GRPC metadata
		runtime.WithIncomingHeaderMatcher(CustomGRPCHeaderMatcher),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = webapi.RegisterWebHandlerFromEndpoint(context.Background(), gwmux, fmt.Sprintf(":%d", *rpcPort), opts)
	if err != nil {
		log.Fatal(err)
	}
	handler := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(gwmux)

	log.Infof("gateway started successfully")
	startListenAndServeByPort(*port, handler)
}

func startListenAndServeByPort(port int, handler http.Handler) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Infof("startListenAndServeByPort with err:%v", err)
	}
}
