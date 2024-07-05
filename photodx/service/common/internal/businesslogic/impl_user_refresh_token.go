package businesslogic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateUserRefreshToken(
	ctx context.Context,
	userID entity.UserID,
) (string, error) {
	now := t.NowFunc()
	ttlDays := 7
	expiresAt := now.Add(time.Hour * time.Duration(ttlDays) * 24)
	claims := auth.JWTClaimsUserRefreshToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(userID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	tokenString, err := t.UserRefreshTokenJWTCreator.CreateJWTToken(
		ctx,
		&claims,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}

func (t *Impl) VerifyUserRefreshToken(
	ctx context.Context,
	refreshToken string,
) (entity.UserPrincipalRefreshToken, error) {
	claims, err := t.UserRefreshTokenJWTVerifier.VerifyJWTToken(ctx, refreshToken, &auth.JWTClaimsUserRefreshToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsRefreshToken, ok := claims.(*auth.JWTClaimsUserRefreshToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsUserRefreshToken")
	}
	return claimsRefreshToken, nil
}
