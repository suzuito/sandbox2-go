package businesslogic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateAdminRefreshToken(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (string, error) {
	now := t.NowFunc()
	ttlDays := 7
	expiresAt := now.Add(time.Hour * time.Duration(ttlDays) * 24)
	claims := auth.JWTClaimsAdminRefreshToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(photoStudioMemberID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	tokenString, err := t.AdminRefreshTokenJWTCreator.CreateJWTToken(
		ctx,
		&claims,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}

func (t *Impl) VerifyAdminRefreshToken(
	ctx context.Context,
	refreshToken string,
) (entity.AdminPrincipalRefreshToken, error) {
	claims, err := t.AdminRefreshTokenJWTVerifier.VerifyJWTToken(ctx, refreshToken, &auth.JWTClaimsAdminRefreshToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsRefreshToken, ok := claims.(*auth.JWTClaimsAdminRefreshToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsAccessToken")
	}
	return claimsRefreshToken, nil
}
