// Package logger
package logger

import (
	"context"
	"fmt"
	"github.com/stnss/dealls-interview/internal/consts"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	conf = &Config{}
)

// Field log object
type Field struct {
	Key   string
	Value interface{}
}

// FieldFunc log
type FieldFunc func(key string, value interface{}) *Field

type Fields []Field

// NewFields create instance new field
func NewFields(p ...Field) Fields {
	x := Fields{}

	for i := 0; i < len(p); i++ {
		x.Append(p[i])
	}

	return x
}

// Append new field
func (f *Fields) Append(p Field) {
	*f = append(*f, p)
}

// Any log
func Any(k string, v interface{}) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

// String log
func String(k string, v string) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

// EventName log
func EventName(v interface{}) Field {
	return Field{
		Key:   EventNameKey,
		Value: v,
	}
}

// MessageFormat message with custom argument
func MessageFormat(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func extract(args ...Field) map[string]interface{} {
	if len(args) == 0 {
		return nil
	}

	data := map[string]interface{}{}
	for _, fl := range args {
		data[fl.Key] = fl.Value
	}
	return data
}

func AddHook(hook logrus.Hook) {
	logrus.AddHook(hook)
}

// Error log
func Error(arg interface{}, fl ...Field) {
	logrus.WithFields(
		addField(extract(fl...)),
	).Error(arg)

}

func Info(arg interface{}, fl ...Field) {
	logrus.WithFields(
		addField(extract(fl...)),
	).Info(arg)
}

func Debug(arg interface{}, fl ...Field) {
	logrus.WithFields(
		addField(extract(fl...)),
	).Debug(arg)
}

// Fatal log
func Fatal(arg interface{}, fl ...Field) {
	logrus.WithFields(
		addField(extract(fl...)),
	).Fatal(arg)
}

// Warn log
func Warn(arg interface{}, fl ...Field) {
	logrus.WithFields(
		addField(extract(fl...)),
	).Warn(arg)
}

// Trace log
func Trace(arg interface{}, fl ...Field) {
	logrus.WithFields(
		addField(extract(fl...)),
	).Trace(arg)
}

// InfoWithContext log info with context
func InfoWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(
		fieldFromContext(ctx, extract(fl...)),
	).WithContext(ctx).Info(arg)
}

// WarnWithContext log warn with context
func WarnWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(
		fieldFromContext(ctx, extract(fl...)),
	).WithContext(ctx).Warn(arg)
}

// ErrorWithContext log error with context
func ErrorWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(
		fieldFromContext(ctx, extract(fl...)),
	).WithContext(ctx).Error(arg)
}

// DebugWithContext log debug with context
func DebugWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(
		fieldFromContext(ctx, extract(fl...)),
	).WithContext(ctx).Debug(arg)
}

// TraceWithContext log trace with context
func TraceWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(
		fieldFromContext(ctx, extract(fl...)),
	).WithContext(ctx).Trace(arg)
}

func addField(f logrus.Fields) logrus.Fields {
	if f == nil {
		f = logrus.Fields{}
	}
	if Environment(conf.Environment) != "" {
		f[EnvironmentKey] = Environment(conf.Environment)
	}

	if conf.ServiceName != "" {
		f[ServiceKey] = conf.ServiceName
	}
	return f
}

func fieldFromContext(ctx context.Context, logField map[string]any) map[string]any {
	if logField == nil {
		logField = logrus.Fields{}
	}

	t, tOk := ctx.Value(consts.ContextKeyStartTime).(time.Time)
	if tOk {
		logField["process_time"] = time.Since(t).Seconds()
		logField["process_time_unit"] = "second"
		logField[StartTimeField] = t
	}

	logField[RequestIpField] = ctx.Value(consts.ContextKeyIP)
	logField[RequestPathField] = ctx.Value(consts.ContextKeyPath)
	logField[RequestMethodField] = ctx.Value(consts.ContextKeyMethod)
	logField[EnvironmentKey] = Environment(conf.Environment)

	return logField
}
