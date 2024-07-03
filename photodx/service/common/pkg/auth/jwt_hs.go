package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type JWTHS struct {
	PrivateKey []byte
}

var signingMethodHS = jwt.SigningMethodHS256

func (t *JWTHS) VerifyJWTToken(
	ctx context.Context,
	tokenString string,
	zeroClaims jwt.Claims,
) (jwt.Claims, error) {
	return verifyJWTToken(ctx, tokenString, signingMethodHS, t.PrivateKey, zeroClaims)
}

func (t *JWTHS) CreateJWTToken(
	ctx context.Context,
	claims jwt.Claims,
) (string, error) {
	return createJWTToken(ctx, claims, signingMethodHS, t.PrivateKey)
}
