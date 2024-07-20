package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Repository interface {
	CreateOAuth2State(
		ctx context.Context,
		state *oauth2loginflow.State,
	) (*oauth2loginflow.State, error)
	GetAndDeleteOAuth2State(
		ctx context.Context,
		stateCode oauth2loginflow.StateCode,
	) (*oauth2loginflow.State, error)

	CreateUserByResourceOwnerID(
		ctx context.Context,
		user *common_entity.User,
		providerID oauth2loginflow.ProviderID,
		resourceOwnerID oauth2loginflow.ResourceOwnerID,
	) (*common_entity.User, error)
	GetUserByResourceOwnerID(
		ctx context.Context,
		providerID oauth2loginflow.ProviderID,
		resourceOwnerID oauth2loginflow.ResourceOwnerID,
	) (*common_entity.User, error)
	GetUsers(
		ctx context.Context,
		userIDs []common_entity.UserID,
	) ([]*common_entity.User, error)
	GetUser(
		ctx context.Context,
		userID common_entity.UserID,
	) (*common_entity.User, error)

	UpdateOrCreateUserWebPushSubscription(
		ctx context.Context,
		s *entity.UserWebPushSubscription,
	) (*entity.UserWebPushSubscription, error)
	GetLatestUserWebPushSubscriptions(
		ctx context.Context,
		userID common_entity.UserID,
	) ([]*entity.UserWebPushSubscription, error)
}
