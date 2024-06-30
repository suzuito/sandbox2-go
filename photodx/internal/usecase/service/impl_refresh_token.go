package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/auth"
)

func (t *Impl) CreateRefreshToken(
	ctx context.Context,
	photoStudioMemberID entity.PhotoStudioMemberID,
) (string, error) {
	now := t.NowFunc()
	ttlDays := 7
	expiresAt := now.Add(time.Hour * time.Duration(ttlDays) * 24)
	claims := auth.JWTClaimsRefreshToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(photoStudioMemberID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	tokenString, err := t.RefreshTokenJWTCreator.CreateJWTToken(
		ctx,
		&claims,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}

func (t *Impl) VerifyRefreshToken(
	ctx context.Context,
	refreshToken string,
) (entity.PrincipalRefreshToken, error) {
	claims, err := t.RefreshTokenJWTVerifier.VerifyJWTToken(ctx, refreshToken, &auth.JWTClaimsRefreshToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsRefreshToken, ok := claims.(*auth.JWTClaimsRefreshToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsAccessToken")
	}
	return claimsRefreshToken, nil
}
