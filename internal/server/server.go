package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/router"
	"github.com/stnss/dealls-interview/pkg/logger"
	"net/http"
	"time"
)

type httpServer struct {
	config *appctx.Config
	app    *fiber.App
	router router.Router
}

func NewHttpServer() Server {
	config := appctx.NewConfig()

	fiberConfig := fiber.Config{
		AppName:      config.App.Name,
		ReadTimeout:  config.App.ReadTimeout,
		WriteTimeout: config.App.WriteTimeout,
	}

	app := fiber.New(fiberConfig)

	return &httpServer{
		config: config,
		app:    app,
		router: router.NewRouter(config, app),
	}
}

func (s *httpServer) Run(ctx context.Context) error {
	var err error

	go func() {
		s.router.Route()
		err = s.app.Listen(fmt.Sprintf("0.0.0.0:%d", s.config.App.Port))
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(logger.MessageFormat("http server error: %v", err))
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.app.ShutdownWithContext(ctxShutDown)
	if err != nil {
		logger.Fatal(logger.MessageFormat("http server shutdown got error: %v", err))
	}

	logger.Info("server exited properly")

	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}

	return err
}

func (s *httpServer) Done() {
	logger.Info("service has stopped")
}

func (s *httpServer) Config() *appctx.Config {
	return s.config
}
