package auth

import (
	"context"
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
