package server

import (
	"context"
	"github.com/stnss/dealls-interview/internal/appctx"
)

type Server interface {
	Run(context.Context) error
	Done()
	Config() *appctx.Config
}
