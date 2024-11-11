package cryptx

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrDecodeBase64            = Error("failed to decode Base64")
	ErrDecodeBase64CipherText  = Error("failed to decode Base64 cipher text")
	ErrParsePEMBlock           = Error("failed to parse private key PEM block")
	ErrBcryptHash              = Error("failed to hash bcrypt")
	ErrUnsupportedRSAKeyFormat = Error("unsupported rsa key format")
	ErrParsePrivateKey         = Error("failed to parse private key")
)
