package jwttoken

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
)

type JWTClaimsAccessToken struct {
	jwt.RegisteredClaims
	Roles []entity.RoleID
}

func (t *JWTClaimsAccessToken) Validate() error {
	return nil
}

func (t *JWTClaimsAccessToken) GetPrincipalID() entity.PrincipalID {
	return entity.PrincipalID(t.Subject)
}

func (t *JWTClaimsAccessToken) GetRoles() []entity.RoleID {
	return t.Roles
}
