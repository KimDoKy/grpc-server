package network

import (
	"grpc-server/config"
	"grpc-server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg    *config.Config
	servie *service.Service
	engin  *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, servie: service, engin: gin.New()}
	return r, nil
}

func (n *Network) StartServer() {
	n.engin.Run(":8080")
}
