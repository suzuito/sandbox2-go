package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type JWTClaimsUserAccessToken struct {
	jwt.RegisteredClaims
	Roles []*rbac.Role `json:"roles"`
}

func (t *JWTClaimsUserAccessToken) GetUserID() entity.UserID {
	return entity.UserID(t.Subject)
}

func (t *JWTClaimsUserAccessToken) GetPermissions() []*pbrbac.Permission {
	permissions := []*pbrbac.Permission{}
	for _, role := range t.Roles {
		permissions = append(permissions, role.Permissions...)
	}
	return permissions
}
