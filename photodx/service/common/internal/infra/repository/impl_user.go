package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/oauth2loginflow"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

var deleteMeUser *entity.User

func (t *Impl) CreateUser(
	ctx context.Context,
	user *entity.User,
) (*entity.User, error) {
	deleteMeUser = user
	return user, nil
}

func (t *Impl) GetUserByResourceOwnerID(
	ctx context.Context,
	providerID oauth2loginflow.ProviderID,
	resourceOwnerID oauth2loginflow.ResourceOwnerID,
) (*entity.User, error) {
	if deleteMeUser == nil {
		return nil, &common_repository.NoEntryError{}
	}
	return deleteMeUser, nil
}
