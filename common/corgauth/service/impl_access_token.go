package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
	"github.com/suzuito/sandbox2-go/common/corgauth/service/jwttoken"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreateAccessToken(
	ctx context.Context,
	principalID entity.PrincipalID,
) (string, error) {
	principal, roles, err := t.Repository.GetPrincipal(
		ctx,
		principalID,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	roleIDs := []entity.RoleID{}
	for _, role := range roles {
		roleIDs = append(roleIDs, role.ID)
	}
	now := t.NowFunc()
	ttlMinutes := 5
	expiresAt := now.Add(time.Second * time.Duration(ttlMinutes) * 60)
	claims := jwttoken.JWTClaimsAccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(principal.ID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Roles: roleIDs,
	}
	if err := claims.Validate(); err != nil {
		return "", terrors.Wrap(err)
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
) (entity.ClaimsAccessToken, error) {
	claims, err := t.AccessTokenJWTVerifier.VerifyJWTToken(ctx, accessToken, &jwttoken.JWTClaimsAccessToken{})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsAccessToken, ok := claims.(*jwttoken.JWTClaimsAccessToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsAccessToken")
	}
	return claimsAccessToken, nil
}
