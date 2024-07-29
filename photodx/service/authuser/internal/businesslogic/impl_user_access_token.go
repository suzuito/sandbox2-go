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
	isGuest bool,
) (string, error) {
	now := t.NowFunc()
	ttlMinutes := 5
	expiresAt := now.Add(time.Second * time.Duration(ttlMinutes) * 60)
	roles := []rbac.RoleID{}
	if isGuest {
		roles = append(roles, rbac.RoleGuest.ID)
	} else {
		roles = append(roles, rbac.RoleUser.ID)
	}
	claims := auth.JWTClaimsUserAccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   string(userID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Roles:       roles,
		IsGuestUser: isGuest,
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
