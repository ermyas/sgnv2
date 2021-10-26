package gateway

import (
	"context"
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/relayer"
	"net"
	"net/http"
	"strings"
	"time"

	gatewaysvc "github.com/celer-network/sgn-v2/gateway/svc"
	"github.com/gogo/gateway"

	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"

	"github.com/celer-network/goutils/log"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 8081, "Listening port")
	rpcPort = flag.Int("rpc", 10000, "Listening port for rpc")
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
	_dbUrl string) {

	gatewaysvc.RootDir = _homeDir
	gatewaysvc.LegacyAmino = _legacyAmino
	gatewaysvc.Cdc = _cdc
	gatewaysvc.InterfaceRegistry = _interfaceRegistry
	gatewaysvc.SelfStart = _selfStart

	flag.Parse()
	log.Infof("Starting gateway at rest:%d, grpc:%d", *port, *rpcPort)

	gs, err := gatewaysvc.NewGatewayService(_dbUrl)
	if err != nil {
		log.Fatalf("fail to init gateway server, err:%v", err)
		return
	}
	defer gs.Close()
	log.Infof(" gateway svc started")

	err = gs.InitTransactors()
	if err != nil {
		log.Fatalf("fail to init transactor in gateway server, err:%v", err)
		return
	}

	gs.StartChainTokenPolling(10 * time.Second)
	log.Infof("chain token cached")

	gs.F = fee.NewTokenPriceCache(gs.TP.GetTransactor())
	log.Infof(" token price cached")

	gs.StartUpdateTokenPricePolling(relayer.Interval)

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
