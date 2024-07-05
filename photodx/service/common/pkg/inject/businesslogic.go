package inject

import (
	businesslogic_internal "github.com/suzuito/sandbox2-go/photodx/service/common/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
)

func NewBusinessLogic(
	adminAccessTokenJWTVerifier auth.JWTVerifier,
) businesslogic.BusinessLogic {
	return &businesslogic_internal.Impl{
		AdminAccessTokenJWTVerifier: adminAccessTokenJWTVerifier,
	}
}
