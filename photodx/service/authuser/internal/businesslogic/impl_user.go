package businesslogic

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
)

func (t *Impl) CreateUserIfNotExistsFromResourceOwnerID(
	ctx context.Context,
	providerID string,
	resourceOwnerID string,
) (*entity.User, error) {
	existingUser, err := t.Repository.GetUserByResourceOwnerID(
		ctx,
		oauth2loginflow.ProviderID(providerID),
		oauth2loginflow.ResourceOwnerID(resourceOwnerID),
	)
	if err == nil {
		return existingUser, nil
	}
	var noEntryError *repository.NoEntryError
	if !errors.As(err, &noEntryError) {
		return nil, terrors.Wrap(err)
	}
	user := common_entity.User{}
	userID, err := t.UserIDGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	user.ID = entity.UserID(userID)
	user.Active = true
	user.InitializedByUser = false
	createdUser, err := t.Repository.CreateUserByResourceOwnerID(
		ctx,
		&user,
		oauth2loginflow.ProviderID(providerID),
		oauth2loginflow.ResourceOwnerID(resourceOwnerID),
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return createdUser, nil
}
