package contract

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/services"
)

type Core struct {
	Cfg      *appctx.Config
	Services *services.Services
}
