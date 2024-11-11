package auth

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stnss/dealls-interview/internal/presentation"
	"github.com/stnss/dealls-interview/mocks/services/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"testing"
)

func TestServe(t *testing.T) {
	testCases := []struct {
		name            string
		param           any
		setupMocks      func(ctx context.Context, mockService *auth.MockService)
		expectedStatus  int
		expectedCode    string
		expectedMessage string
	}{
		{
			name: "successful login",
			param: presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "password",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Login", mock.Anything, &presentation.LoginRequest{
					Email:    "test@example.com",
					Password: "password",
				}).Return(&presentation.LoginResponse{
					AccessToken:  "access-token",
					RefreshToken: "refresh-token",
				}, nil)
			},
			expectedStatus:  fiber.StatusOK,
			expectedMessage: "authorized successfully",
		},
		{
			name:            "failed to parse request body",
			param:           "a",
			setupMocks:      func(ctx context.Context, mockService *auth.MockService) {}, // No setup needed as BodyParser fails before reaching service
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "AU-001",
			expectedMessage: "Bad Request",
		},
		{
			name: "validation error",
			param: presentation.LoginRequest{
				Email:    "",
				Password: "",
			},
			setupMocks:      func(ctx context.Context, mockService *auth.MockService) {}, // Validation fails before reaching service
			expectedStatus:  fiber.StatusUnprocessableEntity,
			expectedMessage: "validation error",
		},
		{
			name: "failed decryption",
			param: presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Login", mock.Anything, &presentation.LoginRequest{
					Email:    "test@example.com",
					Password: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				}).Return(nil, rsa.ErrDecryption)
			},
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "AU-002",
			expectedMessage: "failed decrypt password",
		},
		{
			name: "mismatched password",
			param: presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Login", mock.Anything, &presentation.LoginRequest{
					Email:    "test@example.com",
					Password: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				}).Return(nil, bcrypt.ErrMismatchedHashAndPassword)
			},
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "AU-003",
			expectedMessage: "email or password invalid",
		},
		{
			name: "user not found",
			param: presentation.LoginRequest{
				Email:    "notfound@example.com",
				Password: "password",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Login", mock.Anything, &presentation.LoginRequest{
					Email:    "notfound@example.com",
					Password: "password",
				}).Return(nil, consts.ErrNoRowsFound)
			},
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "AU-003",
			expectedMessage: "email or password invalid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(auth.MockService)
			controller := newLoginController(mockService)
			ctx := context.Background()

			tc.setupMocks(ctx, mockService)

			handler := func() fiber.Handler {
				return func(c *fiber.Ctx) error {
					resp := controller.Serve(appctx.Data{Ctx: c})
					return c.Status(resp.StatusCode).Send(resp.Byte())
				}
			}

			app := fiber.New()
			app.Post("/login", handler())

			// Create a new request with the payload
			payload, _ := json.Marshal(tc.param)
			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")
			res, err := app.Test(req, -1)
			assert.NoError(t, err)

			// Perform assertions
			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			var responseBody map[string]interface{}
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.NoError(t, err)

			if tc.expectedCode != "" {
				assert.Equal(t, tc.expectedCode, responseBody["code"])
			}
			if tc.expectedMessage != "" {
				assert.Equal(t, tc.expectedMessage, responseBody["message"])
			}

			mockService.AssertExpectations(t)
		})
	}
}
