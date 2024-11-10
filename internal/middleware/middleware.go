package middleware

import (
	"github.com/stnss/dealls-interview/internal/repositories"
)

type Middleware struct {
	Injector FuncV2
}

func NewMiddleware(repo *repositories.Repository) *Middleware {
	return &Middleware{
		Injector: injector,
	}
}
