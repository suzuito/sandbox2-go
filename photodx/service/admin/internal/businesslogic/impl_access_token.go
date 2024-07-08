package businesslogic

import (
	"context"

	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) VerifyAdminAccessToken(
	ctx context.Context,
	accessToken string,
) (entity.AdminPrincipalAccessToken, error) {
	return common_businesslogic.VerifyAdminAccessToken(
		ctx,
		t.AdminAccessTokenJWTVerifier,
		accessToken,
	)
}
