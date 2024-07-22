package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Usecase interface {
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)
	MiddlewareRefreshTokenAuthe(
		ctx context.Context,
		refreshToken string,
	) (*DTOMiddlewareRefreshTokenAuthe, error)
	AuthPostLogin(
		ctx context.Context,
		photoStudioID entity.PhotoStudioID,
		email string,
		password string,
	) (*DTOAuthPostLogin, error)
	AuthPostRefresh(
		ctx context.Context,
		photoStudioMemberID entity.PhotoStudioMemberID,
	) (*DTOAuthPostRefresh, error)

	AuthGetInit(
		ctx context.Context,
		principal entity.AdminPrincipalAccessToken,
	) (*DTOAuthGetInit, error)
}
