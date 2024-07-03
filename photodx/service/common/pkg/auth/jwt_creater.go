package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type JWTCreator interface {
	CreateJWTToken(
		ctx context.Context,
		claims jwt.Claims,
	) (string, error)
}

func createJWTToken(
	ctx context.Context,
	claims jwt.Claims,
	signingMethod jwt.SigningMethod,
	key any,
) (string, error) {
	token := jwt.NewWithClaims(
		signingMethod,
		claims,
	)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}
