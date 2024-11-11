package cryptx

import (
	"golang.org/x/crypto/bcrypt"
)

func (c *cryptox) BcryptHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrBcryptHash
	}
	return string(hashedPassword), nil
}

func (c *cryptox) BcryptValidate(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
