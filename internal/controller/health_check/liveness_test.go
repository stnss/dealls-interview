package health_check

import (
	"github.com/valyala/fasthttp"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stretchr/testify/assert"
)

func TestLivenessController_Serve(t *testing.T) {
	// Arrange
	ctrl := newLivenessController()
	mockCtx := fiber.New().AcquireCtx(&fasthttp.RequestCtx{})
	mockData := appctx.Data{
		Ctx:    mockCtx,
		Config: &appctx.Config{}, // Assuming Config is not relevant for this test
	}

	// Act
	response := ctrl.Serve(mockData)

	// Assert
	assert.Equal(t, "Perfectly Fine", response.Message)
	assert.Equal(t, 200, response.StatusCode)
}

func TestLivenessController_EventName(t *testing.T) {
	// Arrange
	ctrl := newLivenessController()

	// Act
	eventName := ctrl.EventName()

	// Assert
	assert.Equal(t, "controller.liveness", eventName)
}
