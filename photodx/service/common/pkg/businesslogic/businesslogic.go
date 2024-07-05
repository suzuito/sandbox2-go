package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	// impl_admin_access_token.go
	VerifyAdminAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.AdminPrincipal, error)
}
