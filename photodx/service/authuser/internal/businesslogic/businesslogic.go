package businesslogic

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"

	webpush "github.com/SherClockHolmes/webpush-go"
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

	// impl_user.go
	GetUser(
		ctx context.Context,
		userID common_entity.UserID,
	) (*common_entity.User, error)

	// impl_user_creation.go
	CreateUserCreationRequest(
		ctx context.Context,
		email string,
		ttlSeconds int,
		frontURL *url.URL,
	) (*common_entity.UserCreationRequest, error)
	DeleteUserCreationRequest(
		ctx context.Context,
		userCreationRequestID common_entity.UserCreationRequestID,
	) error
	GetValidUserCreationRequestNotExpired(
		ctx context.Context,
		id common_entity.UserCreationRequestID,
		code common_entity.UserCreationCode,
	) (*common_entity.UserCreationRequest, error)
	GetUserCreationRequestNotExpired(
		ctx context.Context,
		id common_entity.UserCreationRequestID,
	) (*common_entity.UserCreationRequest, error)
	CreateUser(
		ctx context.Context,
		email string,
		planinPassword string,
	) (*common_entity.User, error)

	// impl_web_push.go
	GetWebPushVAPIDPublicKey(
		ctx context.Context,
		userID common_entity.UserID,
	) (string, error)
	CreateWebPushSubscription(
		ctx context.Context,
		subscription *webpush.Subscription,
		userID common_entity.UserID,
	) error
	PushNotification(
		ctx context.Context,
		l *slog.Logger,
		userID common_entity.UserID,
		notification *common_entity.Notification,
	) error
}
