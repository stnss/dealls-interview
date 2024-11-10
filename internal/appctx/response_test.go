package appctx

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name         string
		statusCode   int
		message      string
		timestamp    time.Time
		errors       interface{}
		code         string
		data         interface{}
		expectedResp Response
	}{
		{
			name:       "Default Response",
			statusCode: http.StatusOK,
			expectedResp: Response{
				StatusCode: http.StatusOK,
				Timestamp:  time.Now().Format(time.RFC3339),
			},
		},
		{
			name:       "With StatusCode",
			statusCode: http.StatusCreated,
			expectedResp: Response{
				StatusCode: http.StatusCreated,
				Timestamp:  time.Now().Format(time.RFC3339),
			},
		},
		{
			name:    "With Message",
			message: "Success",
			expectedResp: Response{
				StatusCode: http.StatusOK,
				Timestamp:  time.Now().Format(time.RFC3339),
				Message:    "Success",
			},
		},
		{
			name:      "With Timestamp",
			timestamp: now,
			expectedResp: Response{
				StatusCode: http.StatusOK,
				Timestamp:  now.Format(time.RFC3339),
			},
		},
		{
			name:   "With Errors",
			errors: "Error occurred",
			expectedResp: Response{
				StatusCode: http.StatusOK,
				Timestamp:  time.Now().Format(time.RFC3339),
				Errors:     "Error occurred",
			},
		},
		{
			name: "With Code",
			code: "200",
			expectedResp: Response{
				StatusCode: http.StatusOK,
				Timestamp:  time.Now().Format(time.RFC3339),
				Code:       "200",
			},
		},
		{
			name: "With Data",
			data: map[string]interface{}{"key": "value"},
			expectedResp: Response{
				StatusCode: http.StatusOK,
				Timestamp:  time.Now().Format(time.RFC3339),
				Data:       map[string]interface{}{"key": "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := NewResponse()

			if tt.statusCode != 0 {
				resp = resp.WithStatusCode(tt.statusCode)
			}

			if tt.message != "" {
				resp = resp.WithMessage(tt.message)
			}

			if !tt.timestamp.IsZero() {
				resp = resp.WithTimestamp(tt.timestamp)
			}

			if tt.errors != nil {
				resp = resp.WithErrors(tt.errors)
			}

			if tt.code != "" {
				resp = resp.WithCode(tt.code)
			}

			if tt.data != nil {
				resp = resp.WithData(tt.data)
			}

			assert.Equal(t, tt.expectedResp.StatusCode, resp.StatusCode)
			assert.Equal(t, tt.expectedResp.Message, resp.Message)
			assert.Equal(t, tt.expectedResp.Timestamp, resp.Timestamp)
			assert.Equal(t, tt.expectedResp.Errors, resp.Errors)
			assert.Equal(t, tt.expectedResp.Code, resp.Code)
			assert.Equal(t, tt.expectedResp.Data, resp.Data)
		})
	}
}

func TestResponse_Byte(t *testing.T) {
	resp := NewResponse().WithMessage("Success").WithStatusCode(http.StatusOK)
	expected := `{"status_code":200,"timestamp":"` + resp.Timestamp + `","message":"Success"}`

	assert.JSONEq(t, expected, string(resp.Byte()))
}
