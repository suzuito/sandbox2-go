package businesslogic

import (
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
)

type Impl struct {
	AdminAccessTokenJWTVerifier auth.JWTVerifier
	UserAccessTokenJWTVerifier  auth.JWTVerifier
}
