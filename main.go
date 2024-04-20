package main

import (
	"flag"
	"grpc-server/cmd"
	"grpc-server/config"
	"grpc-server/gRPC/server"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)

	if err := server.NewGRPCServer(cfg); err != nil {
		panic(err)
	} else {
		time.Sleep(1e9)
		cmd.NewApp(cfg)
	}
}
