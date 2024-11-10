package consts

const (
	// HeaderXRequestID const
	HeaderXRequestID = `X-Request-ID`

	// HeaderXForwardedFor const
	HeaderXForwardedFor = `X-Forwarded-For`

	// HeaderXRealIP const
	HeaderXRealIP = `X-Real-IP`
)

const (
	ContextKeyRequestID = iota
	ContextKeyStartTime
	ContextKeyIP
	ContextKeyPath
	ContextKeyMethod
)
