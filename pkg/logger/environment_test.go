package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Valid inputs with different cases and trimming spaces
		{name: "Production exact", input: "production", expected: "production"},
		{name: "Production uppercase", input: "PRODUCTION", expected: "production"},
		{name: "Production with spaces", input: " production ", expected: "production"},
		{name: "Prod short form", input: "prod", expected: "production"},
		{name: "Prd short form", input: "prd", expected: "production"},

		{name: "Staging exact", input: "staging", expected: "staging"},
		{name: "Staging uppercase", input: "STAGING", expected: "staging"},
		{name: "Stg short form", input: "stg", expected: "staging"},

		{name: "Development exact", input: "development", expected: "development"},
		{name: "Development uppercase", input: "DEVELOPMENT", expected: "development"},
		{name: "Dev short form", input: "dev", expected: "development"},

		{name: "Local exact", input: "local", expected: "local"},
		{name: "Test exact", input: "test", expected: "test"},
		{name: "Green exact", input: "green", expected: "green"},
		{name: "Blue exact", input: "blue", expected: "blue"},

		// Invalid inputs
		{name: "Invalid environment", input: "invalid", expected: ""},
		{name: "Empty string", input: "", expected: ""},
		{name: "Only spaces", input: "   ", expected: ""},
		{name: "Non-existent key", input: "unknownEnv", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Environment(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
