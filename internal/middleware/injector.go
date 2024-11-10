package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/consts"
	"time"
)

func injector(next Func) Func {
	return func(c *fiber.Ctx, conf *appctx.Config) appctx.Response {
		reqId := c.Get(consts.HeaderXRequestID)
		if reqId == "" {
			uid, _ := uuid.NewV7()
			reqId = uid.String()
		}
		// Set new id to response header
		c.Set(consts.HeaderXRequestID, reqId)

		ctx := context.WithValue(c.UserContext(), consts.HeaderXRequestID, reqId)
		ctx = context.WithValue(ctx, consts.ContextKeyStartTime, time.Now())
		ctx = context.WithValue(ctx, consts.ContextKeyIP, c.IP())
		ctx = context.WithValue(ctx, consts.ContextKeyPath, c.Path())
		ctx = context.WithValue(ctx, consts.ContextKeyMethod, c.Method())

		c.SetUserContext(ctx)

		return next(c, conf)
	}
}
