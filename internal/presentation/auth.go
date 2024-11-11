package presentation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	RegistrationRequest struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}
)

func (param *LoginRequest) Validate() error {
	return validation.ValidateStruct(
		param,
		validation.Field(&param.Email, validation.Required, is.EmailFormat),
		validation.Field(&param.Password, validation.Required, is.Base64),
	)
}

func (param *RegistrationRequest) Validate() error {
	return validation.ValidateStruct(
		param,
		validation.Field(&param.Name, validation.Required, is.Alpha),
		validation.Field(&param.Email, validation.Required, is.EmailFormat),
		validation.Field(&param.Password, validation.Required, is.Base64),
		validation.Field(&param.PasswordConfirmation, validation.Required, is.Base64),
	)
}
