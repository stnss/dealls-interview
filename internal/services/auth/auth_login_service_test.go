package auth

import (
	"context"
	"errors"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/entity"
	"github.com/stnss/dealls-interview/internal/presentation"
	cryptx2 "github.com/stnss/dealls-interview/mocks/pkg/cryptx"
	jwtx2 "github.com/stnss/dealls-interview/mocks/pkg/jwtx"
	user2 "github.com/stnss/dealls-interview/mocks/repositories/user"
	"github.com/stnss/dealls-interview/pkg/jwtx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	// Mocks
	testCases := []struct {
		name             string
		param            *presentation.LoginRequest
		setupMocks       func(ctx context.Context, userRepo *user2.MockRepository, c *cryptx2.MockHelper, j *jwtx2.MockHelper)
		expectedResponse *presentation.LoginResponse
		expectedError    error
	}{
		{
			name: "successful login",
			param: &presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, userRepo *user2.MockRepository, c *cryptx2.MockHelper, j *jwtx2.MockHelper) {
				now := time.Now()
				c.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("plain-password", nil)
				userRepo.On("GetUserByEmail", ctx, "test@example.com").Return(&entity.User{
					ID:       "123",
					Name:     "Test User",
					Password: "hashed-password",
				}, nil)
				c.On("BcryptValidate", "hashed-password", "plain-password").Return(nil)
				j.On("GenerateJWT", jwtx.UserClaim{
					Uid:  "123",
					Name: "Test User",
				}, "access-secret", mock.Anything).Return("access-token", now.Add(15*time.Minute))
				j.On("GenerateJWT", jwtx.UserClaim{
					Uid: "123",
				}, "refresh-secret", mock.Anything).Return("refresh-token", now.Add(7*24*time.Hour))
			},
			expectedResponse: &presentation.LoginResponse{
				AccessToken:  "access-token",
				ExpiresIn:    899,
				RefreshToken: "refresh-token",
			},
			expectedError: nil,
		},
		{
			name: "password decryption fails",
			param: &presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, userRepo *user2.MockRepository, c *cryptx2.MockHelper, j *jwtx2.MockHelper) {
				c.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("", errors.New("decryption error"))
			},
			expectedResponse: nil,
			expectedError:    errors.New("decryption error"),
		},
		{
			name: "user not found",
			param: &presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, userRepo *user2.MockRepository, c *cryptx2.MockHelper, j *jwtx2.MockHelper) {
				c.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("plain-password", nil)
				userRepo.On("GetUserByEmail", ctx, "test@example.com").Return(nil, errors.New("user not found"))
			},
			expectedResponse: nil,
			expectedError:    errors.New("user not found"),
		},
		{
			name: "password validation fails",
			param: &presentation.LoginRequest{
				Email:    "test@example.com",
				Password: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, userRepo *user2.MockRepository, c *cryptx2.MockHelper, j *jwtx2.MockHelper) {
				c.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("plain-password", nil)
				userRepo.On("GetUserByEmail", ctx, "test@example.com").Return(&entity.User{
					ID:       "123",
					Name:     "Test User",
					Password: "hashed-password",
				}, nil)
				c.On("BcryptValidate", "hashed-password", "plain-password").Return(errors.New("password mismatch"))
			},
			expectedResponse: nil,
			expectedError:    errors.New("password mismatch"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(user2.MockRepository)
			c := new(cryptx2.MockHelper)
			j := new(jwtx2.MockHelper)

			svc := NewAuthService(&appctx.Config{
				JWT: appctx.JWT{
					AccessSecret:  "access-secret",
					RefreshSecret: "refresh-secret",
				},
			}, userRepo, j, c)

			ctx := context.Background()

			tc.setupMocks(ctx, userRepo, c, j)

			response, err := svc.Login(ctx, tc.param)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, response)
			}

			// Verify all expectations met
			userRepo.AssertExpectations(t)
			c.AssertExpectations(t)
			j.AssertExpectations(t)
		})
	}
}
