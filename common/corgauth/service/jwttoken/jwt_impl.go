package jwttoken

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

// 複数のJWT検証アルゴリズムを許容する
// JWTの切り替えなど発生した際に、移行期が発生すると思われるため、作成した
type JWTVerifiers struct {
	Verifiers []JWTVerifier
}

func (t *JWTVerifiers) VerifyJWTToken(
	ctx context.Context,
	tokenString string,
	zeroClaims jwt.Claims,
) (jwt.Claims, error) {
	for _, verifier := range t.Verifiers {
		claims, err := verifier.VerifyJWTToken(ctx, tokenString, zeroClaims)
		if err == nil {
			return claims, nil
		}
		fmt.Printf("%+v\n", err)
	}
	return nil, terrors.Wrapf("cannot verify")
}

// 共通鍵認証を用いたJWT
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

// 公開鍵認証を用いたJWT
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
