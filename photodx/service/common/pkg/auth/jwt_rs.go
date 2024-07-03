package auth

import (
	"context"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

type JWTVerifierRS struct {
	PublicKey *rsa.PublicKey
}

var signingMethodRS = jwt.SigningMethodRS256

func (t *JWTVerifierRS) VerifyJWTToken(
	ctx context.Context,
	tokenString string,
	zeroClaims jwt.Claims,
) (jwt.Claims, error) {
	return verifyJWTToken(ctx, tokenString, signingMethodRS, t.PublicKey, zeroClaims)
}

type JWTCreatorRS struct {
	PrivateKey *rsa.PrivateKey
}

func (t *JWTCreatorRS) CreateJWTToken(
	ctx context.Context,
	claims jwt.Claims,
) (string, error) {
	return createJWTToken(ctx, claims, signingMethodRS, t.PrivateKey)
}
