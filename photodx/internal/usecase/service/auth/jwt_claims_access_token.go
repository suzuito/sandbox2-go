package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/rbac"
)

type JWTClaimsAccessToken struct {
	jwt.RegisteredClaims
	Hoge  string        `json:"hoge"`
	Roles []rbac.RoleID `json:"roles"`
}
