package auth

import (
	"context"
	"errors"
	"github.com/stnss/dealls-interview/internal/appctx"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stnss/dealls-interview/internal/presentation"
	cryptx2 "github.com/stnss/dealls-interview/mocks/pkg/cryptx"
	jwtx2 "github.com/stnss/dealls-interview/mocks/pkg/jwtx"
	user2 "github.com/stnss/dealls-interview/mocks/repositories/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRegistration(t *testing.T) {
	testCases := []struct {
		name          string
		param         presentation.RegistrationRequest
		setupMocks    func(ctx context.Context, mockUserRepo *user2.MockRepository, mockCrypto *cryptx2.MockHelper)
		expectedError error
	}{
		{
			name: "successful registration",
			param: presentation.RegistrationRequest{
				Email:                "test@example.com",
				Password:             "encrypted-password",
				PasswordConfirmation: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, mockUserRepo *user2.MockRepository, mockCrypto *cryptx2.MockHelper) {
				mockCrypto.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("plain-password", nil).Twice()
				mockCrypto.On("BcryptHash", "plain-password").Return("hashed-password", nil)
				mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "decryption failure",
			param: presentation.RegistrationRequest{
				Email:                "test@example.com",
				Password:             "encrypted-password",
				PasswordConfirmation: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, mockUserRepo *user2.MockRepository, mockCrypto *cryptx2.MockHelper) {
				mockCrypto.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("", errors.New("decryption error"))
			},
			expectedError: errors.New("decryption error"),
		},
		{
			name: "passwords do not match",
			param: presentation.RegistrationRequest{
				Email:                "test@example.com",
				Password:             "encrypted-password1",
				PasswordConfirmation: "encrypted-password2",
			},
			setupMocks: func(ctx context.Context, mockUserRepo *user2.MockRepository, mockCrypto *cryptx2.MockHelper) {
				mockCrypto.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password1").Return("plain-password1", nil)
				mockCrypto.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password2").Return("plain-password2", nil)
			},
			expectedError: consts.ErrPasswordNotMatch,
		},
		{
			name: "bcrypt hashing failure",
			param: presentation.RegistrationRequest{
				Email:                "test@example.com",
				Password:             "encrypted-password",
				PasswordConfirmation: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, mockUserRepo *user2.MockRepository, mockCrypto *cryptx2.MockHelper) {
				mockCrypto.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("plain-password", nil).Twice()
				mockCrypto.On("BcryptHash", "plain-password").Return("", errors.New("hashing error"))
			},
			expectedError: errors.New("hashing error"),
		},
		{
			name: "database insertion error",
			param: presentation.RegistrationRequest{
				Email:                "test@example.com",
				Password:             "encrypted-password",
				PasswordConfirmation: "encrypted-password",
			},
			setupMocks: func(ctx context.Context, mockUserRepo *user2.MockRepository, mockCrypto *cryptx2.MockHelper) {
				mockCrypto.On("DecryptRSAWithBase64", mock.Anything, "encrypted-password").Return("plain-password", nil).Twice()
				mockCrypto.On("BcryptHash", "plain-password").Return("hashed-password", nil)
				mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*entity.User")).Return(errors.New("database error"))
			},
			expectedError: errors.New("database error"),
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

			tc.setupMocks(ctx, userRepo, c)

			// Call the Registration function
			err := svc.Registration(ctx, &tc.param)

			// Assert the expected error (if any)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			// Verify that all expectations are met
			c.AssertExpectations(t)
			userRepo.AssertExpectations(t)
		})
	}
}
