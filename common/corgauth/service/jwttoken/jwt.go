package jwttoken

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

type JWTVerifier interface {
	VerifyJWTToken(
		ctx context.Context,
		tokenString string,
		zeroClaims jwt.Claims,
	) (jwt.Claims, error)
}

func verifyJWTToken(
	ctx context.Context,
	tokenString string,
	signingMethod jwt.SigningMethod,
	key any,
	zeroClaims jwt.Claims,
) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		zeroClaims,
		func(token *jwt.Token) (any, error) {
			if token.Method.Alg() != signingMethod.Alg() {
				return nil, terrors.Wrapf("unexpected signing method: %v", token.Method.Alg())
			}
			return key, nil
		},
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return token.Claims, nil
}
