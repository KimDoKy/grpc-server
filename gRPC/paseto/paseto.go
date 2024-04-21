package paseto

import (
	"grpc-server/config"
	auth "grpc-server/gRPC/proto"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg *config.Config) *PasetoMaker {
	return &PasetoMaker{Pt: paseto.NewV2(), Key: []byte(cfg.Paseto.Key)}
}

func (m *PasetoMaker) CreateNewToken(auth auth.AuthData) (string, error) {
	return m.Pt.Encrypt(m.Key, auth, nil)
}

func (m *PasetoMaker) VerifyToken(token string) error {
	var auth auth.AuthData
	return m.Pt.Decrypt(token, m.Key, &auth, nil)
}
