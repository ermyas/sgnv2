package gateway

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"

	"github.com/celer-network/goutils/log"
	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 8081, "Listening port")
	rpcPort = flag.Int("rpc", 10000, "Listening port for rpc")
)

func InitGateway(
	_homeDir string,
	_legacyAmino *codec.LegacyAmino,
	_cdc codec.Codec,
	_interfaceRegistry codectypes.InterfaceRegistry,
	_selfStart bool,
	_dbUrl string) {

	rootDir = _homeDir
	legacyAmino = _legacyAmino
	cdc = _cdc
	interfaceRegistry = _interfaceRegistry
	selfStart = _selfStart

	flag.Parse()
	log.Infof("Starting gateway at rest:%d, grpc:%d", *port, *rpcPort)

	gs, err := NewGatewayService(_dbUrl)
	if err != nil {
		log.Fatalf("fail to init gateway server, err:%v", err)
		return
	}
	defer gs.Close()
	log.Infof(" gateway svc started")

	err = gs.initTransactors()
	if err != nil {
		log.Fatalf("fail to init transactor in gateway server, err:%v", err)
		return
	}

	gs.StartChainTokenPolling(10 * time.Second)
	log.Infof("chain token cached")

	gs.f = fee.NewTokenPriceCache(gs.tp.GetTransactor())
	log.Infof(" token price cached")

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

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				EmitDefaults: true,
				OrigName:     true,
				EnumsAsInts:  true,
			}}))
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
