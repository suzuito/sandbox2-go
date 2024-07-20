package auth

import (
	"context"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

type JWTVerifierRS256 struct {
	PublicKey *rsa.PublicKey
}

var signingMethodRS = jwt.SigningMethodRS256

func (t *JWTVerifierRS256) VerifyJWTToken(
	ctx context.Context,
	tokenString string,
	zeroClaims jwt.Claims,
) (jwt.Claims, error) {
	return verifyJWTToken(ctx, tokenString, signingMethodRS, t.PublicKey, zeroClaims)
}

type JWTCreatorRS256 struct {
	PrivateKey *rsa.PrivateKey
}

func (t *JWTCreatorRS256) CreateJWTToken(
	ctx context.Context,
	claims jwt.Claims,
) (string, error) {
	return createJWTToken(ctx, claims, signingMethodRS, t.PrivateKey)
}
