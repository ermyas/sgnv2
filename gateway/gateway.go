package webapi

import (
	"context"
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"net"
	"net/http"
	"time"

	"github.com/celer-network/goutils/log"
	"google.golang.org/grpc"
)

var (
	port    = flag.Int("port", 8080, "Listening port")
	rpcPort = flag.Int("rpc", 10000, "Listening port for rpc")
)

func InitGateway() {
	flag.Parse()
	log.Infof("Starting gateway at rest:%d, grpc:%d", *port, *rpcPort)

	gs, err := NewGatewayService("127.0.0.1:26257")
	if err != nil {
		log.Fatalf("fail to init gateway server, err:%v", err)
		return
	}
	defer gs.Close()

	gs.StartChainTokenPolling(10 * time.Second)
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

	gwmux := runtime.NewServeMux()
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

	startListenAndServeByPort(*port, handler)
}

func startListenAndServeByPort(port int, handler http.Handler) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Infof("startListenAndServeByPort with err:%v", err)
	}
}
