// Package logger
package logger

const (
	// LogTypePrint this log type for print
	LogTypePrint = "print"

	// LogTypeFile this log type for file
	LogTypeFile = "file"

	// LogFormatJSON this log format json
	LogFormatJSON = "json"

	EventKey     = "event"
	EventNameKey = "name"
	EventIdKey   = "id"

	// ServiceKey holds the service field
	ServiceKey = "service"

	// EnvironmentKey holds the environment field
	EnvironmentKey = "env"

	LogDriverLoki    = "loki"
	LogDriverGraylog = "graylog"

	StartTimeKey   = "start-time"
	StartTimeField = "start_time"

	RequestIpKey       = `request-ip`
	RequestIpField     = `request_ip`
	RequestMethodKey   = `request-method`
	RequestMethodField = `request_method`
	RequestPathKey     = `request-path`
	RequestPathField   = `request_path`
)
