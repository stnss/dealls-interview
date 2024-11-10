package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/controller/contract"
)

type httpHandleFunc func(*fiber.Ctx, contract.Controller, *appctx.Config) appctx.Response

func HttpRequest(xCtx *fiber.Ctx, svc contract.Controller, conf *appctx.Config) appctx.Response {
	data := appctx.Data{
		Ctx:    xCtx,
		Config: conf,
	}
	return svc.Serve(data)
}

type Router interface {
	Route()
}
