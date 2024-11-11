package services

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/dependencies"
	"github.com/stnss/dealls-interview/internal/providers"
	"github.com/stnss/dealls-interview/internal/repositories"
	"github.com/stnss/dealls-interview/internal/services/auth"
)

type (
	Services struct {
		Auth auth.Service
	}

	Dependency struct {
		Repository *repositories.Repository
		Provider   *providers.Provider
	}
)

func NewService(cfg *appctx.Config, pkgs *dependencies.Dependency, dep *Dependency) *Services {
	return &Services{
		Auth: auth.NewAuthService(cfg, dep.Repository.UserRepository, pkgs.Jwt, pkgs.Crypto),
	}
}
