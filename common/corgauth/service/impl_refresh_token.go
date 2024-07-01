package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
	"github.com/suzuito/sandbox2-go/common/corgauth/service/jwttoken"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreateRefreshToken(
	ctx context.Context,
	principalID entity.PrincipalID,
) (string, error) {
	now := t.NowFunc()
	ttlDays := 7
	expiresAt := now.Add(time.Hour * time.Duration(ttlDays) * 24)
	claims := jwttoken.JWTClaimsRefreshToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(principalID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	if err := claims.Validate(); err != nil {
		return "", terrors.Wrap(err)
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
) (entity.ClaimsRefreshToken, error) {
	claims, err := t.RefreshTokenJWTVerifier.VerifyJWTToken(
		ctx, refreshToken, &jwttoken.JWTClaimsRefreshToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsRefreshToken, ok := claims.(*jwttoken.JWTClaimsRefreshToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsAccessToken")
	}
	return claimsRefreshToken, nil
}
