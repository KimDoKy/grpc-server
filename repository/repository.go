package repository

import (
	"grpc-server/config"
	"grpc-server/gRPC/client"
	auth "grpc-server/gRPC/proto"
)

type Repository struct {
	cfg        *config.Config
	gRPCClient *client.GRPCClient
}

func NewRepository(cfg *config.Config, gRPCClient *client.GRPCClient) (*Repository, error) {
	r := &Repository{cfg: cfg, gRPCClient: gRPCClient}

	return r, nil
}

func (r *Repository) CreateAuth(name string) (*auth.AuthData, error) {
	return r.gRPCClient.CreateAuth(name)
}
