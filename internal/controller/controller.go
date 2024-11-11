package controller

import (
	"github.com/stnss/dealls-interview/internal/controller/auth"
	"github.com/stnss/dealls-interview/internal/controller/contract"
	healthcheck "github.com/stnss/dealls-interview/internal/controller/health_check"
)

type Controller struct {
	HealthCheck    *healthcheck.HealthCheck
	Authentication *auth.Authentication
}

func NewController(dep *contract.Dependency) *Controller {
	return &Controller{
		HealthCheck:    healthcheck.NewHealthCheck(),
		Authentication: auth.NewAuthentication(dep.Services),
	}
}
