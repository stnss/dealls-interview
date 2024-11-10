package health_check

import (
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/controller/contract"
	"github.com/stnss/dealls-interview/pkg/logger"
)

type livenessController struct{}

func newLivenessController() contract.Controller {
	return &livenessController{}
}

func (ctrl *livenessController) EventName() string {
	return "controller.liveness"
}

func (ctrl *livenessController) Serve(data appctx.Data) appctx.Response {
	const (
		LivenessMessage = `Perfectly Fine`
	)

	var (
		lf = logger.NewFields(
			logger.EventName(ctrl.EventName()),
		)
		ctx = data.Ctx.UserContext()
	)

	logger.InfoWithContext(ctx, "Liveness Check", lf...)

	return *appctx.NewResponse().
		WithMessage(LivenessMessage)
}
