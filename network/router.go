package network

import (
	"grpc-server/config"
	"grpc-server/service"
)

type Network struct {
	cfg    *config.Config
	servie *service.Service
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, servie: service}
	return r, nil
}
