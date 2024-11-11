package jwtx

import "time"

type Helper interface {
	GenerateJWT(data UserClaim, secret string, expiration time.Duration) (string, time.Time)
	ClaimJWT(accessToken, secret string) (*Claims, error)
}
