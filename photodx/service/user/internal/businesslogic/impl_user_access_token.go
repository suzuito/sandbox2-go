package businesslogic

import (
	"context"

	common_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) VerifyUserAccessToken(
	ctx context.Context,
	accessToken string,
) (entity.UserPrincipalAccessToken, error) {
	return common_businesslogic.VerifyUserAccessToken(
		ctx,
		t.UserAccessTokenJWTVerifier,
		accessToken,
	)
}
