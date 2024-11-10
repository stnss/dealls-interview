package appctx

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var (
	onceRsp sync.Once
	rsp     *Response
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Timestamp  string      `json:"timestamp,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Code       string      `json:"code,omitempty"`
	Errors     any         `json:"errors,omitempty"`
}

func NewResponse() *Response {
	onceRsp.Do(func() {
		rsp = &Response{
			StatusCode: http.StatusOK,
			Timestamp:  time.Now().Format(time.RFC3339),
		}
	})
	x := *rsp
	return &x
}

func (r *Response) WithStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) WithTimestamp(t time.Time) *Response {
	r.Timestamp = t.Format(time.RFC3339)
	return r
}

func (r *Response) WithErrors(err any) *Response {
	r.Errors = err
	return r
}

func (r *Response) WithCode(code string) *Response {
	r.Code = code
	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) Byte() []byte {
	b, _ := json.Marshal(r)
	return b
}
