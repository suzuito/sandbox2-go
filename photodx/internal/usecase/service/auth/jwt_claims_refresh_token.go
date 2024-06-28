package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaimsRefreshToken struct {
	jwt.RegisteredClaims
	Hoge string `json:"hoge"`
}
