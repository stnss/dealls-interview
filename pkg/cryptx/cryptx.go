package cryptx

type Helper interface {
	BcryptHash(password string) (string, error)
	BcryptValidate(hashedPassword, password string) error
	DecryptRSAWithBase64(base64PrivateKey string, base64Ciphertext string) (string, error)
}

type cryptox struct{}

func NewCryptox() Helper {
	return &cryptox{}
}
