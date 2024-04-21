package network

import (
	"grpc-server/config"
	"grpc-server/gRPC/client"
	"grpc-server/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg *config.Config

	service    *service.Service
	gRPCClient *client.GRPCClient

	engin *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service, gRPCClient *client.GRPCClient) (*Network, error) {
	r := &Network{cfg: cfg, service: service, engin: gin.New(), gRPCClient: gRPCClient}

	r.engin.POST("/login", r.login)

	r.engin.GET("/verify", r.verifyLogin(), r.verify)

	return r, nil
}

func (n *Network) StartServer() {
	n.engin.Run(":8080")
}
