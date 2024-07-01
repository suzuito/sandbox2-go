package jwttoken

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/common/corgauth/entity"
)

type JWTClaimsRefreshToken struct {
	jwt.RegisteredClaims
}

func (t *JWTClaimsRefreshToken) Validate() error {
	return nil
}

func (t *JWTClaimsRefreshToken) GetPrincipalID() entity.PrincipalID {
	return entity.PrincipalID(t.Subject)
}
