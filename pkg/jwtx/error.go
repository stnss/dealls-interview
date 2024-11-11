package jwtx

const (
	ErrUnexpectedSigningMethod = Error("unexpected signing method")
	ErrAccessTokenExpired      = Error("access token expired")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
