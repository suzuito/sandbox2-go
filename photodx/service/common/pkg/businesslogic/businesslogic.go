package businesslogic

import (
	"context"

	internal_businesslogic "github.com/suzuito/sandbox2-go/photodx/service/common/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/auth"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	VerifyAdminAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.AdminPrincipalAccessToken, error)
	VerifyUserAccessToken(
		ctx context.Context,
		accessToken string,
	) (entity.UserPrincipalAccessToken, error)
}

func NewBusinessLogic(
	AdminAccessTokenJWTVerifier auth.JWTVerifier,
	UserAccessTokenJWTVerifier auth.JWTVerifier,
) BusinessLogic {
	return &internal_businesslogic.Impl{
		AdminAccessTokenJWTVerifier: AdminAccessTokenJWTVerifier,
		UserAccessTokenJWTVerifier:  UserAccessTokenJWTVerifier,
	}
}
