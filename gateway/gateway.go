package gateway

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	gatewaysvc "github.com/celer-network/sgn-v2/gateway/svc"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

var (
	port       = flag.Int("port", 18081, "Listening port")
	rpcWebPort = flag.Int("rpcweb", 18082, "Listening port for rpc web")
	rpcPort    = flag.Int("rpc", 20000, "Listening port for rpc")
	home       = flag.String("home", os.ExpandEnv("$HOME/.gateway"), "config path")
)

// CustomGRPCHeaderMatcher for mapping request headers to
// GRPC metadata.
// HTTP headers that start with 'Grpc-Metadata-' are automatically mapped to
// gRPC metadata after removing prefix 'Grpc-Metadata-'. We can use this
// CustomGRPCHeaderMatcher if headers don't start with Grpc-Metadata-
func CustomGRPCHeaderMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case grpctypes.GRPCBlockHeightHeader:
		return grpctypes.GRPCBlockHeightHeader, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func InitGateway() {
	flag.Parse()
	log.Infof("Starting gateway at rest: %d, grpc: %d, grpc-web: %d", *port, *rpcPort, *rpcWebPort)

	setupConfigDefaults()
	loadConfigs()
	setupEnv()

	db := dal.NewDAL(viper.GetString(common.FlagGatewayDbUrl))
	gs := gatewaysvc.NewGatewayService(db)
	defer gs.Close()

	encoding := app.MakeEncodingConfig()
	onchain.InitSGNTransactors(*home, encoding)
	gs.Chains = onchain.InitChainMgr(db)

	gs.StartChainTokenPolling(2 * time.Minute)
	gs.F = gatewaysvc.NewTokenPriceCache(onchain.SGNTransactors.GetTransactor())
	gs.StartUpdateTokenPricePolling(time.Duration(viper.GetInt32(common.FlagSgnCheckIntervalCbrPrice)) * time.Second)
	gs.StartAvgLpFeeEarningPolling(10 * time.Minute)
	gs.StartAbnormalBalanceCheckPolling(1 * time.Hour)

	grpcSvr := startGrpcServer(gs)
	startGrpcWebServer(grpcSvr)

	// The default JSON marshaller used by the gRPC-Gateway is unable to marshal non-nullable non-scalar fields.
	// Using the gogo/gateway package with the gRPC-Gateway WithMarshaler option fixes the scalar field marshalling issue.
	marshalerOption := &gateway.JSONPb{
		EmitDefaults: true,
		EnumsAsInts:  true,
		Indent:       "  ",
		OrigName:     true,
		AnyResolver:  encoding.InterfaceRegistry,
	}
	startGrpcGatewaySvr(marshalerOption)
}

func startGrpcServer(gs *gatewaysvc.GatewayService) *grpc.Server {
	// start a rpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *rpcPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcSvr := grpc.NewServer()
	webapi.RegisterWebServer(grpcSvr, gs)
	go func() {
		if err = grpcSvr.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	return grpcSvr
}

func startGrpcWebServer(grpcSvr *grpc.Server) {
	wrappedServer := grpcweb.WrapServer(grpcSvr, grpcweb.WithOriginFunc(func(origin string) bool { return true }))
	grpcWebSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *rpcWebPort),
		Handler: wrappedServer,
	}
	go func() {
		if grpcWebSrvErr := grpcWebSrv.ListenAndServe(); grpcWebSrvErr != nil {
			log.Fatal(grpcWebSrvErr)
		}
	}()
}

func startGrpcGatewaySvr(opt *gateway.JSONPb) {
	gwmux := runtime.NewServeMux(
		// Custom marshaler option is required for gogo proto
		runtime.WithMarshalerOption(runtime.MIMEWildcard, opt),
		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
		// Custom header matcher for mapping request headers to
		// GRPC metadata
		runtime.WithIncomingHeaderMatcher(CustomGRPCHeaderMatcher),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := webapi.RegisterWebHandlerFromEndpoint(context.Background(), gwmux, fmt.Sprintf(":%d", *rpcPort), opts)
	if err != nil {
		log.Fatal(err)
	}
	handler := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(gwmux)

	log.Infof("gateway started successfully")
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), handler)
	if err != nil {
		log.Infof("startListenAndServeByPort with err:%v", err)
	}
}

func setupConfigDefaults() {
	// sets account address prefix for transactors
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	sdkConfig.Seal()

	// sets home dir for the convenience of all later file loadings
	viper.SetDefault(flags.FlagHome, *home)

	// setup db url
	viper.SetDefault(common.FlagGatewayDbUrl, "127.0.0.1:26257")
}

func loadConfigs() {
	log.Infoln("Loading configs...")
	cbrCfgFile := filepath.Join(*home, "config", "cbridge.toml")
	viper.SetConfigFile(cbrCfgFile)
	if err := viper.MergeInConfig(); err != nil {
		log.Errorf("failed to read in cbridge configuration: %s", err.Error())
		return
	}
	log.Infoln("Loaded cbridge.toml")
	sgnCfgFile := filepath.Join(*home, "config", "sgn.toml")
	viper.SetConfigFile(sgnCfgFile)
	if err := viper.MergeInConfig(); err != nil {
		log.Errorf("failed to read in SGN configuration: %s", err.Error())
		return
	}
	log.Infoln("Loaded sgn.toml")
}

func setupEnv() {
	// setup env var
	env := "local"
	chainEnv := viper.GetString(common.FlagSgnChainId)
	if chainEnv == "sgn-2" || chainEnv == "sgn-3" {
		env = "prod"
	} else if chainEnv == "sgn-testnet-4000" {
		env = "test"
	}
	viper.SetDefault("env", env)
	log.Infoln("Setup env:", viper.GetString("env"))
}
