package businesslogic

import (
	"context"
	"net/http"
	"net/url"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	// impl_user_access_token.go
	CreateUserAccessToken(
		ctx context.Context,
		userID common_entity.UserID,
	) (string, error)

	// impl_user_refresh_token.go
	CreateUserRefreshToken(
		ctx context.Context,
		userID common_entity.UserID,
	) (string, error)
	VerifyUserRefreshToken(
		ctx context.Context,
		refreshToken string,
	) (entity.UserPrincipalRefreshToken, error)

	// impl_oauth2loginflow.go
	CreateOAuth2State(
		ctx context.Context,
		providerID oauth2loginflow.ProviderID,
		callbackURL *url.URL,
		oauth2RedirectURL *url.URL,
	) (*oauth2loginflow.State, error)
	CallbackVerifyState(
		ctx context.Context,
		stateCode oauth2loginflow.StateCode,
	) (*oauth2loginflow.State, error)
	FetchAccessToken(
		ctx context.Context,
		req *http.Request,
	) (string, error)
	FetchProfileAndCreateUserIfNotExists(
		ctx context.Context,
		accessToken string,
		providerID oauth2loginflow.ProviderID,
		fetchProfile oauth2loginflow.Oauth2FetchProfileFunc,
	) (*common_entity.User, error)
}
