package auth

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stnss/dealls-interview/internal/presentation"
	"github.com/stnss/dealls-interview/mocks/services/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestRegistration_Serve(t *testing.T) {
	testCases := []struct {
		name            string
		param           any
		setupMocks      func(ctx context.Context, mockService *auth.MockService)
		expectedStatus  int
		expectedCode    string
		expectedMessage string
	}{
		{
			name: "successful registration",
			param: presentation.RegistrationRequest{
				Name:                 "Example Name",
				Email:                "test@example.com",
				Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Registration", mock.Anything, &presentation.RegistrationRequest{
					Name:                 "Example Name",
					Email:                "test@example.com",
					Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
					PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				}).Return(nil)
			},
			expectedStatus:  fiber.StatusCreated,
			expectedMessage: "registration success",
		},
		{
			name:            "failed to parse request body",
			param:           "a",
			setupMocks:      func(ctx context.Context, mockService *auth.MockService) {}, // No setup needed for parsing error
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "US-001",
			expectedMessage: "Bad Request",
		},
		{
			name: "validation error",
			param: presentation.RegistrationRequest{
				Email:                "",
				Password:             "",
				PasswordConfirmation: "",
			},
			setupMocks:      func(ctx context.Context, mockService *auth.MockService) {}, // Validation fails before reaching service
			expectedStatus:  fiber.StatusUnprocessableEntity,
			expectedMessage: "validation error",
		},
		{
			name: "decryption error",
			param: presentation.RegistrationRequest{
				Name:                 "Example Name",
				Email:                "test@example.com",
				Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Registration", mock.Anything, &presentation.RegistrationRequest{
					Name:                 "Example Name",
					Email:                "test@example.com",
					Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
					PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				}).Return(rsa.ErrDecryption)
			},
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "US-002",
			expectedMessage: "failed decrypt password or password_confirmation",
		},
		{
			name: "email already exists",
			param: presentation.RegistrationRequest{
				Name:                 "Example Name",
				Email:                "duplicate@example.com",
				Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Registration", mock.Anything, &presentation.RegistrationRequest{
					Name:                 "Example Name",
					Email:                "duplicate@example.com",
					Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
					PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				}).Return(consts.ErrUniqueViolation)
			},
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "US-003",
			expectedMessage: "email already exists",
		},
		{
			name: "password confirmation mismatch",
			param: presentation.RegistrationRequest{
				Name:                 "Example Name",
				Email:                "test@example.com",
				Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
			},
			setupMocks: func(ctx context.Context, mockService *auth.MockService) {
				mockService.On("Registration", mock.Anything, &presentation.RegistrationRequest{
					Name:                 "Example Name",
					Email:                "test@example.com",
					Password:             "ZW5jcnlwdGVkLXBhc3N3b3Jk",
					PasswordConfirmation: "ZW5jcnlwdGVkLXBhc3N3b3Jk",
				}).Return(consts.ErrPasswordNotMatch)
			},
			expectedStatus:  fiber.StatusBadRequest,
			expectedCode:    "US-004",
			expectedMessage: "invalid password_confirmation",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(auth.MockService)
			controller := newRegistrationController(mockService)
			ctx := context.Background()

			tc.setupMocks(ctx, mockService)

			handler := func() fiber.Handler {
				return func(c *fiber.Ctx) error {
					resp := controller.Serve(appctx.Data{Ctx: c})
					return c.Status(resp.StatusCode).Send(resp.Byte())
				}
			}

			app := fiber.New()
			app.Post("/registration", handler())

			payload, _ := json.Marshal(tc.param)
			req := httptest.NewRequest("POST", "/registration", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")
			res, err := app.Test(req, -1)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			var responseBody map[string]interface{}
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			fmt.Println(responseBody)
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
