package logger

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// Capture output from logrus logger for verification in tests.
func captureOutput(f func()) string {
	var buf bytes.Buffer
	logrus.SetOutput(&buf) // Temporarily set logrus output to buf

	f()

	logrus.SetOutput(nil) // Restore default
	return buf.String()
}

// Table-driven test for SetJSONFormatter function.
func TestSetJSONFormatter(t *testing.T) {
	SetJSONFormatter()
	assert.IsType(t, &Formatter{}, logrus.StandardLogger().Formatter)

	tests := []struct {
		name        string
		formatter   *Formatter
		expectMsg   bool
		expectFile  bool
		expectPanic bool
		expectBase  bool
	}{
		{
			name: "default configuration",
			formatter: &Formatter{
				ChildFormatter: &logrus.JSONFormatter{
					FieldMap: logrus.FieldMap{
						logrus.FieldKeyMsg: "msg",
					},
				},
				Line:         true,
				File:         true,
				BaseNameOnly: false,
				Package:      true,
			},
			expectMsg:  true,
			expectFile: true,
			expectBase: false,
		},
		{
			name: "file with base name and line info",
			formatter: &Formatter{
				ChildFormatter: &logrus.JSONFormatter{
					FieldMap: logrus.FieldMap{
						logrus.FieldKeyMsg: "msg",
					},
				},
				Line:         true,
				File:         true,
				BaseNameOnly: true,
			},
			expectMsg:  true,
			expectFile: true,
			expectBase: true,
		},
		{
			name: "missing file and line info",
			formatter: &Formatter{
				ChildFormatter: &logrus.JSONFormatter{
					FieldMap: logrus.FieldMap{
						logrus.FieldKeyMsg: "msg",
					},
				},
				Line:         false,
				File:         false,
				BaseNameOnly: false,
			},
			expectMsg:  true,
			expectFile: false,
			expectBase: false,
		},
		{
			name:        "invalid formatter (nil)",
			formatter:   nil,
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				assert.Panics(t, func() {
					logrus.SetFormatter(tt.formatter)
					logrus.Info("test message")
				}, "expected panic but got none")
				return
			}

			// Set the formatter.
			logrus.SetFormatter(tt.formatter)
			output := captureOutput(func() {
				logrus.Info("test message")
			})

			if tt.expectMsg {
				assert.Contains(t, output, `"msg":"test message"`)
			} else {
				assert.NotContains(t, output, `"msg":"test message"`)
			}

			if tt.expectFile {
				assert.Contains(t, output, `"file":`)
			} else {
				assert.NotContains(t, output, `"file":`)
			}

			if tt.expectBase {
				assert.Contains(t, output, `"file":"testing.go"`)
			} else {
				assert.NotContains(t, output, `"file":"testing.go"`)
			}
		})
	}
}

// Table-driven test for Setup function.
func TestSetup(t *testing.T) {
	tests := []struct {
		name     string
		cfg      Config
		expected logrus.Level
	}{
		{
			name:     "debug enabled",
			cfg:      Config{Debug: true, Level: "info", Environment: "development", ServiceName: "test-service"},
			expected: logrus.TraceLevel, // Debug true sets TraceLevel
		},
		{
			name:     "valid log level info",
			cfg:      Config{Debug: false, Level: "info", Environment: "production", ServiceName: "test-service"},
			expected: logrus.InfoLevel,
		},
		{
			name:     "valid log level warn",
			cfg:      Config{Debug: false, Level: "warn", Environment: "production", ServiceName: "test-service"},
			expected: logrus.WarnLevel,
		},
		{
			name:     "valid log level error",
			cfg:      Config{Debug: false, Level: "error", Environment: "staging", ServiceName: "test-service"},
			expected: logrus.ErrorLevel,
		},
		{
			name:     "invalid log level, fallback to info",
			cfg:      Config{Debug: false, Level: "invalid", Environment: "development", ServiceName: "test-service"},
			expected: logrus.InfoLevel, // Invalid level defaults to InfoLevel
		},
		{
			name:     "empty log level, fallback to info",
			cfg:      Config{Debug: false, Level: "", Environment: "production", ServiceName: "test-service"},
			expected: logrus.InfoLevel, // Empty level defaults to InfoLevel
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setup(tt.cfg)

			// Assert that the log level is correctly set
			assert.Equal(t, tt.expected, logrus.GetLevel())
		})
	}
}
