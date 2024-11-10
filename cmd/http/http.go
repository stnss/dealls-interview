package http

import (
	"context"
	"github.com/stnss/dealls-interview/internal/server"
	"github.com/stnss/dealls-interview/pkg/logger"
)

func Start(ctx context.Context) {
	httpServer := server.NewHttpServer()
	defer httpServer.Done()

	logger.Info(logger.MessageFormat("starting %s services... %d", httpServer.Config().App.Name, httpServer.Config().App.Port))

	if err := httpServer.Run(ctx); err != nil {
		logger.Fatal(logger.MessageFormat("http httpServer start got error: %v", err))
	}
}
