package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
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
		user *entity.User,
		providerID oauth2loginflow.ProviderID,
		resourceOwnerID oauth2loginflow.ResourceOwnerID,
	) (*entity.User, error)
	GetUserByResourceOwnerID(
		ctx context.Context,
		providerID oauth2loginflow.ProviderID,
		resourceOwnerID oauth2loginflow.ResourceOwnerID,
	) (*entity.User, error)
}
