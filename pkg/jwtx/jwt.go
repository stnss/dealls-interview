package jwtx

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtHelper struct{}

func NewJwtHelper() Helper {
	return &jwtHelper{}
}

func (j *jwtHelper) GenerateJWT(data UserClaim, secret string, expiration time.Duration) (string, time.Time) {
	expirationTime := time.Now().Add(expiration)
	claims := &Claims{
		UserClaim: data,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString, expirationTime
}

func (j *jwtHelper) ClaimJWT(accessToken, secret string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrUnexpectedSigningMethod, token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrAccessTokenExpired
	}
	return &claims, nil
}
