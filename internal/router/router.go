package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/bootstrap"
	"github.com/stnss/dealls-interview/internal/controller"
	"github.com/stnss/dealls-interview/internal/controller/contract"
	"github.com/stnss/dealls-interview/internal/dependencies"
	"github.com/stnss/dealls-interview/internal/middleware"
	"github.com/stnss/dealls-interview/internal/providers"
	"github.com/stnss/dealls-interview/internal/repositories"
	"github.com/stnss/dealls-interview/internal/services"
	"github.com/stnss/dealls-interview/pkg/logger"
	"net/http"
	"runtime/debug"
)

type router struct {
	cfg   *appctx.Config
	fiber *fiber.App
}

func NewRouter(cfg *appctx.Config, fiber *fiber.App) Router {
	bootstrap.RegistryLogger(cfg)
	return &router{cfg: cfg, fiber: fiber}
}

func (rtr *router) handleRoute(hfn httpHandleFunc, ctrl contract.Controller, mdws ...middleware.FuncV2) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(logger.MessageFormat("got panic error: %v", err))
				logger.Error(logger.MessageFormat("got panic: stack trace: %v", string(debug.Stack())))
				res := *appctx.NewResponse().
					WithStatusCode(http.StatusInternalServerError).
					WithMessage(http.StatusText(http.StatusInternalServerError))

				_ = rtr.response(c, res)
			}
		}()

		adapter := func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
			resp := hfn(xCtx, ctrl, conf)
			return resp
		}

		fn := middleware.Wrap(adapter, mdws...)

		resp := fn(c, rtr.cfg)
		return rtr.response(c, resp)
	}
}

func (rtr *router) handleGroup(mdws ...middleware.FuncV2) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(logger.MessageFormat("got panic error: %v", err))
				logger.Error(logger.MessageFormat("got panic: stack trace: %v", string(debug.Stack())))
				res := *appctx.NewResponse().
					WithStatusCode(http.StatusInternalServerError).
					WithMessage(http.StatusText(http.StatusInternalServerError))

				_ = rtr.response(c, res)
			}
		}()

		if len(mdws) < 1 {
			return c.Next()
		}

		fn := func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
			return *appctx.NewResponse().WithStatusCode(fiber.StatusOK)
		}
		for _, mdw := range mdws {
			fn = mdw(fn)
		}

		if resp := fn(c, rtr.cfg); resp.StatusCode != http.StatusOK {
			return rtr.response(c, resp)
		}

		return c.Next()
	}
}

func (rtr *router) response(c *fiber.Ctx, resp appctx.Response) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(resp.StatusCode).Send(resp.Byte())
}

func (rtr *router) Route() {
	dep := dependencies.NewDependency(rtr.cfg)
	repo := repositories.NewRepository(dep)
	provider := providers.NewProvider(rtr.cfg)
	svcs := services.NewService(rtr.cfg, dep, &services.Dependency{
		Repository: repo,
		Provider:   provider,
	})
	// controllers
	controllers := controller.NewController(&contract.Dependency{
		Core: contract.Core{
			Cfg:      rtr.cfg,
			Services: svcs,
		},
	})
	m := middleware.NewMiddleware(repo)

	rtr.fiber.Use(rtr.handleGroup(
		m.Injector,
	))

	rtr.fiber.Get("/ruok", rtr.handleRoute(
		HttpRequest,
		controllers.HealthCheck.Liveness,
	))

	external := rtr.fiber.Group("external")
	exV1 := external.Group("v1")

	rtr.AuthRoute(exV1, controllers)
}

func (rtr *router) AuthRoute(r fiber.Router, controllers *controller.Controller) {
	r = r.Group("auth")
	r.Post("login", rtr.handleRoute(
		HttpRequest,
		controllers.Authentication.Login,
	))
	r.Post("registration", rtr.handleRoute(
		HttpRequest,
		controllers.Authentication.Registration,
	))
}
