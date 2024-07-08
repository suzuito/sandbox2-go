package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	// impl_user_access_token.go
	VerifyUserAccessToken(ctx context.Context,
		accessToken string,
	) (entity.UserPrincipalAccessToken, error)
}
