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
	"net/http"
)

type registrationController struct {
	svc auth.Service
}

func newRegistrationController(svc auth.Service) contract.Controller {
	return &registrationController{svc: svc}
}

func (ctrl *registrationController) EventName() string {
	return "controller.auth.registration"
}

func (ctrl *registrationController) Serve(data appctx.Data) appctx.Response {
	var (
		lf = logger.NewFields(
			logger.EventName(ctrl.EventName()),
		)

		ctx   = data.Ctx.UserContext()
		param presentation.RegistrationRequest
	)

	if err := data.Ctx.BodyParser(&param); err != nil {
		logger.ErrorWithContext(ctx, logger.MessageFormat("parse request body got err: %+v", err), lf...)
		return *appctx.NewResponse().
			WithStatusCode(fiber.StatusBadRequest).
			WithCode("US-001").
			WithMessage("Bad Request")
	}

	if err := param.Validate(); err != nil {
		logger.WarnWithContext(ctx, "create user got validation error", lf...)
		return *appctx.NewResponse().
			WithStatusCode(fiber.StatusUnprocessableEntity).
			WithMessage("validation error").
			WithErrors(err)
	}

	err := ctrl.svc.Registration(ctx, &param)
	if err != nil {
		switch {
		case errors.Is(err, rsa.ErrDecryption):
			logger.WarnWithContext(ctx, "store user got error decrypt", lf...)
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusBadRequest).
				WithCode("US-002").
				WithMessage("failed decrypt password or password_confirmation")
		case errors.Is(err, consts.ErrUniqueViolation):
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusBadRequest).
				WithCode("US-003").
				WithMessage("email already exists")
		case errors.Is(err, consts.ErrPasswordNotMatch):
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusBadRequest).
				WithCode("US-004").
				WithMessage("invalid password_confirmation")
		default:
			logger.ErrorWithContext(ctx, logger.MessageFormat("create user got error %+v", err), lf...)
			return *appctx.NewResponse().
				WithStatusCode(fiber.StatusInternalServerError).
				WithMessage(http.StatusText(fiber.StatusInternalServerError))
		}
	}

	return *appctx.NewResponse().
		WithStatusCode(fiber.StatusCreated).
		WithMessage("registration success")
}
