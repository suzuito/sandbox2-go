package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth"
)

func (t *Impl) CreateAccessToken(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
	roles []rbac.RoleID,
) (string, error) {
	now := t.NowFunc()
	ttlMinutes := 5
	expiresAt := now.Add(time.Second * time.Duration(ttlMinutes) * 60)
	claims := auth.JWTClaimsAccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(photoStudioMemberID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Roles: roles,
	}
	tokenString, err := t.AccessTokenJWTCreator.CreateJWTToken(
		ctx,
		&claims,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}

func (t *Impl) VerifyAccessToken(
	ctx context.Context,
	accessToken string,
) (entity.Principal, error) {
	claims, err := t.AccessTokenJWTVerifier.VerifyJWTToken(ctx, accessToken, &auth.JWTClaimsAccessToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsAccessToken, ok := claims.(*auth.JWTClaimsAccessToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsAccessToken")
	}
	return claimsAccessToken, nil
}
