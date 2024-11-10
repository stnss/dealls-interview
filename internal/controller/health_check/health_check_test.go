package health_check

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHealthCheck(t *testing.T) {
	// Act
	healthCheck := NewHealthCheck()

	// Assert
	assert.NotNil(t, healthCheck.Liveness, "Liveness controller should not be nil")
	assert.IsType(t, &livenessController{}, healthCheck.Liveness, "Liveness should be of type *livenessController")
	assert.Equal(t, "controller.liveness", healthCheck.Liveness.EventName(), "Liveness controller event name should match")
}
