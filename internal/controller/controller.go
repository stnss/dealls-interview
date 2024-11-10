package controller

import (
	"github.com/stnss/dealls-interview/internal/controller/contract"
	healthcheck "github.com/stnss/dealls-interview/internal/controller/health_check"
)

type Controller struct {
	HealthCheck *healthcheck.HealthCheck
}

func NewController(dep *contract.Dependency) *Controller {
	return &Controller{
		HealthCheck: healthcheck.NewHealthCheck(),
	}
}
