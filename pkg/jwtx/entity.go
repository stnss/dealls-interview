package jwtx

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	Uid  string `json:"uid"`
	Name string `json:"name,omitempty"`
}

type Claims struct {
	UserClaim
	jwt.RegisteredClaims
}
