package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type JWTClaimsAdminRefreshToken struct {
	jwt.RegisteredClaims
	Hoge string `json:"hoge"`
}

func (t *JWTClaimsAdminRefreshToken) GetPhotoStudioMemberID() entity.PhotoStudioMemberID {
	return entity.PhotoStudioMemberID(t.Subject)
}
