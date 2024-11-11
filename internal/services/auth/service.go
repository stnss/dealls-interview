package auth

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/repositories/user"
	"github.com/stnss/dealls-interview/pkg/cryptx"
	"github.com/stnss/dealls-interview/pkg/jwtx"
)

type service struct {
	cfg      *appctx.Config
	userRepo user.Repository
	jwt      jwtx.Helper
	crypto   cryptx.Helper
}

func NewAuthService(cfg *appctx.Config, userRepo user.Repository, jwt jwtx.Helper, crypto cryptx.Helper) Service {
	return &service{cfg: cfg, userRepo: userRepo, jwt: jwt, crypto: crypto}
}
