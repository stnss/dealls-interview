package auth

import (
	"context"
	"github.com/stnss/dealls-interview/internal/presentation"
)

type Service interface {
	Login(ctx context.Context, param *presentation.LoginRequest) (*presentation.LoginResponse, error)
	Registration(ctx context.Context, param *presentation.RegistrationRequest) error
}
