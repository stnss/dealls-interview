package auth

import (
	"github.com/stnss/dealls-interview/internal/controller/contract"
	"github.com/stnss/dealls-interview/internal/services"
)

type Authentication struct {
	Login        contract.Controller
	Registration contract.Controller
}

func NewAuthentication(svc *services.Services) *Authentication {
	return &Authentication{
		Login:        newLoginController(svc.Auth),
		Registration: newRegistrationController(svc.Auth),
	}
}
