package cmd

import (
	"grpc-server/config"
	"grpc-server/gRPC/client"
	"grpc-server/network"
	"grpc-server/repository"
	"grpc-server/service"
)

type App struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient
	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	client.NewGRPCClient(cfg)

	if a.gRPCClient, err = client.NewGRPCClient(cfg); err != nil {
		panic(err)
	} else if a.repository, err = repository.NewRepository(cfg, a.gRPCClient); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service, a.gRPCClient); err != nil {
		panic(err)
	} else {
		a.network.StartServer()
	}
}
