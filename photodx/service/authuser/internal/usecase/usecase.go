package usecase

import (
	"context"
	"net/url"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
)

type Usecase interface {
	GetAuthorizeURLLINE(
		ctx context.Context,
		oauth2RedirectURL *url.URL,
	) (*DTOGetAuthorizeLINE, error)
	GetCallback(
		ctx context.Context,
		code string,
		stateCode oauth2loginflow.StateCode,
	) (*DTOGetCallback, error)

	MiddlewareRefreshTokenAuthe(
		ctx context.Context,
		refreshToken string,
	) (*DTOMiddlewareRefreshTokenAuthe, error)
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)
	PostRefreshAccessToken(
		ctx context.Context,
		principal entity.UserPrincipalRefreshToken,
	) (*DTOPostRefreshAccessToken, error)
}
