package businesslogic

import "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"

type Impl struct {
	UserAccessTokenJWTVerifier auth.JWTVerifier
}
