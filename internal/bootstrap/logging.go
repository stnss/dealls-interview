package bootstrap

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/pkg/logger"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: logger.Environment(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.App.Name,
	})
}
