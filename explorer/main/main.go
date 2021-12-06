package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/explorer"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

var (
	configPath = flag.String("cf", "/Users/liuxiao/code/sgn-v2/sgn-v2/explorer/env/local/config.toml", "config file path")
)

func main() {
	flag.Parse()
	config, err := explorer.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("fail to init config from file, path:%s, err:%s", *configPath, err.Error())
	}

	es, err := explorer.NewExplorerServer(config)
	if err != nil {
		log.Fatalf("fail to init gateway server, err:%v", err)
		return
	}
	go es.StartScheduleJob()
	// start a rpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	explorer.RegisterExplorerServer(grpcServer, es)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// Add grpc web
	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool { return true }))
	grpcWebSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GrpcWebPort),
		Handler: wrappedServer,
	}
	if grpcWebSrvErr := grpcWebSrv.ListenAndServe(); grpcWebSrvErr != nil {
		log.Fatal(grpcWebSrvErr)
	}
}
