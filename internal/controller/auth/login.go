package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stnss/dealls-interview/internal/controller/contract"
	"github.com/stnss/dealls-interview/internal/presentation"
	"github.com/stnss/dealls-interview/internal/services/auth"
	"github.com/stnss/dealls-interview/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type login struct {
	svc auth.Service
}

func newLoginController(svc auth.Service) contract.Controller {
	return &login{svc: svc}
}

func (ctrl *login) EventName() string {
	return "controller.auth.login"
}

func (ctrl *login) Serve(data appctx.Data) appctx.Response {
	var (
		lf = logger.NewFields(
			logger.EventName(ctrl.EventName()),
		)
		ctx = data.Ctx.UserContext()

		param presentation.LoginRequest
	)

	logger.InfoWithContext(ctx, "authorizing user...", lf...)

	if err := data.Ctx.BodyParser(&param); err != nil {
		logger.ErrorWithContext(ctx, logger.MessageFormat("parse request body got err: %+v", err), lf...)
		return *appctx.NewResponse().
			WithStatusCode(fiber.StatusBadRequest).
			WithCode("AU-001").
			WithMessage("Bad Request")
	}

	if err := param.Validate(); err != nil {
		logger.WarnWithContext(ctx, "authorized user got validation error", lf...)
		return *appctx.NewResponse().
			WithStatusCode(fiber.StatusUnprocessableEntity).
			WithMessage("validation error").
			WithErrors(err)
	}

	resp, err := ctrl.svc.Login(ctx, &param)
	if err != nil {
		switch {
		case errors.Is(err, rsa.ErrDecryption):
			logger.WarnWithContext(ctx, "authorized user got error decrypt", lf...)
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusBadRequest).
				WithCode("AU-002").
				WithMessage("failed decrypt password")
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword), errors.Is(err, consts.ErrNoRowsFound):
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusBadRequest).
				WithCode("AU-003").
				WithMessage("email or password invalid")
		default:
			logger.ErrorWithContext(ctx, logger.MessageFormat("authorized user got error %+v", err), lf...)
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusInternalServerError).
				WithMessage(http.StatusText(fiber.StatusInternalServerError))
		}
	}

	return *appctx.NewResponse().
		WithStatusCode(fiber.StatusOK).
		WithMessage("authorized successfully").
		WithData(resp)
}
