package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type JWTClaimsRefreshToken struct {
	jwt.RegisteredClaims
	Hoge string `json:"hoge"`
}

func (t *JWTClaimsRefreshToken) GetPhotoStudioMemberID() entity.PhotoStudioMemberID {
	return entity.PhotoStudioMemberID(t.Subject)
}
