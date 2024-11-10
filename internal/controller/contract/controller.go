package contract

import "github.com/stnss/dealls-interview/internal/appctx"

type Controller interface {
	EventName() string
	Serve(appctx.Data) appctx.Response
}
