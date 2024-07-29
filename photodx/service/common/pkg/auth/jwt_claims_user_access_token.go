package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type JWTClaimsUserAccessToken struct {
	jwt.RegisteredClaims
	Roles       []rbac.RoleID `json:"roles"`
	IsGuestUser bool
}

func (t *JWTClaimsUserAccessToken) GetUserID() entity.UserID {
	return entity.UserID(t.Subject)
}

func (t *JWTClaimsUserAccessToken) IsGuest() bool {
	return t.IsGuestUser
}

func (t *JWTClaimsUserAccessToken) GetPermissions() []*pbrbac.Permission {
	permissions := []*pbrbac.Permission{}
	roles := rbac.GetAvailablePredefinedRolesFromRoleID(t.Roles)
	for _, role := range roles {
		permissions = append(permissions, role.Permissions...)
	}
	return permissions
}
