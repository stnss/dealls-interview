package services

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/providers"
	"github.com/stnss/dealls-interview/internal/repositories"
)

type (
	Services struct {
	}

	Dependency struct {
		Repository *repositories.Repository
		Provider   *providers.Provider
	}
)

func NewService(cfg *appctx.Config, dep *Dependency) *Services {
	return &Services{}
}
