package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type JWTClaimsUserRefreshToken struct {
	jwt.RegisteredClaims
}

func (t *JWTClaimsUserRefreshToken) GetUserID() entity.UserID {
	return entity.UserID(t.Subject)
}
