package health_check

import "github.com/stnss/dealls-interview/internal/controller/contract"

type HealthCheck struct {
	Liveness contract.Controller
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		Liveness: newLivenessController(),
	}
}
