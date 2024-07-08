package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	VerifyAdminAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.AdminPrincipalAccessToken, error)
}
