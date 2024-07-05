package businesslogic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

func (t *Impl) CreateUserAccessToken(
	ctx context.Context,
	userID entity.UserID,
) (string, error) {
	now := t.NowFunc()
	ttlMinutes := 5
	expiresAt := now.Add(time.Second * time.Duration(ttlMinutes) * 60)
	claims := auth.JWTClaimsUserAccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(userID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Roles: []*rbac.Role{
			&rbac.RoleUser,
		},
	}
	tokenString, err := t.UserAccessTokenJWTCreator.CreateJWTToken(
		ctx,
		&claims,
	)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	return tokenString, nil
}

func (t *Impl) VerifyUserAccessToken(ctx context.Context,
	accessToken string,
) (entity.UserPrincipal, error) {
	claims, err := t.UserAccessTokenJWTVerifier.VerifyJWTToken(ctx, accessToken, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	claimsAccessToken, ok := claims.(*auth.JWTClaimsUserAccessToken)
	if !ok {
		return nil, terrors.Wrapf("cannot convert JWTClaims to JWTClaimsUserAccessToken")
	}
	return claimsAccessToken, nil
}
