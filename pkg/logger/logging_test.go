package logger

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func captureLogOutput(f func()) string {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	defer logrus.SetOutput(nil) // reset to default output
	f()
	return buf.String()
}

func TestAddField(t *testing.T) {
	conf = &Config{
		Environment: "test",
		ServiceName: "service",
	}

	tests := []struct {
		name     string
		input    logrus.Fields
		expected logrus.Fields
	}{
		{
			name:     "Add environment and service fields",
			input:    logrus.Fields{},
			expected: logrus.Fields{EnvironmentKey: "test", ServiceKey: "service"},
		},
		{
			name:     "Existing fields with environment and service added",
			input:    logrus.Fields{"key": "value"},
			expected: logrus.Fields{"key": "value", EnvironmentKey: "test", ServiceKey: "service"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addField(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFieldFromContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.ContextKeyStartTime, time.Now().Add(-time.Second))
	ctx = context.WithValue(ctx, consts.ContextKeyIP, "127.0.0.1")
	ctx = context.WithValue(ctx, consts.ContextKeyPath, "/test")
	ctx = context.WithValue(ctx, consts.ContextKeyMethod, "GET")

	result := fieldFromContext(ctx, nil)

	assert.NotNil(t, result["process_time"], "expected process_time")
	assert.Equal(t, "127.0.0.1", result[RequestIpField], "expected IP")
	assert.Equal(t, "/test", result[RequestPathField], "expected path")
	assert.Equal(t, "GET", result[RequestMethodField], "expected method")
}

func TestFieldFunctions(t *testing.T) {
	lf := NewFields(EventName("lf"))
	assert.Equal(t, EventNameKey, lf[0].Key)
	assert.Equal(t, "lf", lf[0].Value)

	lf.Append(Any("key", "value"))
	assert.Equal(t, "key", lf[1].Key)
	assert.Equal(t, "value", lf[1].Value)

	field := Any("key", "value")
	assert.Equal(t, "key", field.Key)
	assert.Equal(t, "value", field.Value)

	stringField := String("key", "stringValue")
	assert.Equal(t, "key", stringField.Key)
	assert.Equal(t, "stringValue", stringField.Value)

	eventField := EventName("eventName")
	assert.Equal(t, EventNameKey, eventField.Key)
	assert.Equal(t, "eventName", eventField.Value)
}

func TestExtract(t *testing.T) {
	fields := extract(
		Any("key1", "value1"),
		String("key2", "value2"),
	)
	expected := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	assert.Equal(t, expected, fields)
}

func TestAddHook(t *testing.T) {
	//hook := &logrus.
	//AddHook(hook)
	//assert.Contains(t, logrus.StandardLogger().Hooks[logrus.InfoLevel], hook)
}

func TestLoggingFunctions(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	tests := []struct {
		name    string
		logFunc func(interface{}, ...Field)
		level   logrus.Level
		message string
	}{
		{
			name:    "Error",
			logFunc: Error,
			level:   logrus.ErrorLevel,
			message: "error message",
		},
		{
			name:    "Info",
			logFunc: Info,
			level:   logrus.InfoLevel,
			message: "info message",
		},
		{
			name:    "Debug",
			logFunc: Debug,
			level:   logrus.DebugLevel,
			message: "debug message",
		},
		{
			name:    "Warn",
			logFunc: Warn,
			level:   logrus.WarnLevel,
			message: "warn message",
		},
		{
			name:    "Trace",
			logFunc: Trace,
			level:   logrus.TraceLevel,
			message: "trace message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			logrus.SetLevel(tt.level)
			tt.logFunc(tt.message)
			assert.Contains(t, buf.String(), tt.message)
		})
	}
}

func TestLoggingWithContextFunctions(t *testing.T) {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	ctx := context.WithValue(context.Background(), consts.ContextKeyStartTime, time.Now().Add(-time.Second))

	tests := []struct {
		name    string
		logFunc func(context.Context, interface{}, ...Field)
		message string
	}{
		{
			name:    "InfoWithContext",
			logFunc: InfoWithContext,
			message: "info message",
		},
		{
			name:    "WarnWithContext",
			logFunc: WarnWithContext,
			message: "warn message",
		},
		{
			name:    "ErrorWithContext",
			logFunc: ErrorWithContext,
			message: "error message",
		},
		{
			name:    "DebugWithContext",
			logFunc: DebugWithContext,
			message: "debug message",
		},
		{
			name:    "TraceWithContext",
			logFunc: TraceWithContext,
			message: "trace message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.logFunc(ctx, tt.message)
			assert.Contains(t, buf.String(), tt.message)
		})
	}
}

func TestMessageFormat(t *testing.T) {
	msg := MessageFormat("hello %s", "world")
	assert.Equal(t, "hello world", msg)
}

// Additional edge case for Any, String, EventName
func TestFieldFunctions_Empty(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() Field
		expected Field
	}{
		{
			name: "Test Any",
			fn: func() Field {
				return Any("key1", "value1")
			},
			expected: Field{
				Key:   "key1",
				Value: "value1",
			},
		},
		{
			name: "Test String",
			fn: func() Field {
				return String("key2", "value2")
			},
			expected: Field{
				Key:   "key2",
				Value: "value2",
			},
		},
		{
			name: "Test EventName",
			fn: func() Field {
				return EventName("eventName")
			},
			expected: Field{
				Key:   EventNameKey,
				Value: "eventName",
			},
		},
		{
			name: "Test Any with empty key and value",
			fn: func() Field {
				return Any("", nil)
			},
			expected: Field{
				Key:   "",
				Value: nil,
			},
		},
		{
			name: "Test String with empty key and value",
			fn: func() Field {
				return String("", "")
			},
			expected: Field{
				Key:   "",
				Value: "",
			},
		},
		{
			name: "Test EventName with empty value",
			fn: func() Field {
				return EventName("")
			},
			expected: Field{
				Key:   EventNameKey,
				Value: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.fn())
		})
	}
}

// Edge case for fieldFromContext
func TestFieldFromContext_NilValues(t *testing.T) {
	conf = &Config{
		Environment: "",
		ServiceName: "service",
	}
	ctx := context.Background()
	result := fieldFromContext(ctx, nil)
	assert.Equal(t, "", result[EnvironmentKey], "expected environment to be empty")

	ctx = context.WithValue(context.Background(), consts.ContextKeyIP, nil)
	result = fieldFromContext(ctx, nil)
	assert.Equal(t, nil, result[RequestIpField], "expected IP to be nil")
}

// Test extract with duplicate keys
func TestExtract_DuplicateKeys(t *testing.T) {
	fields := []Field{
		{Key: "key1", Value: "value1"},
		{Key: "key1", Value: "value2"},
	}
	expected := map[string]interface{}{
		"key1": "value2", // last value should be retained
	}
	assert.Equal(t, expected, extract(fields...))
}
