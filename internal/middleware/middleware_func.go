package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
)

type Func func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response

type FuncV2 func(Func) Func

func Wrap(handleFunc Func, mfs ...FuncV2) Func {
	fn := handleFunc
	for i := len(mfs) - 1; i >= 0; i-- {
		fn = mfs[i](fn)
	}

	return fn
}
