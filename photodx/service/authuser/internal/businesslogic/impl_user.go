package businesslogic

import (
	"context"
	"errors"
	"fmt"

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
	user *entity.User,
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
	userID, err := t.UserIDGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	user.ID = entity.UserID(userID)
	user.Guest = false
	createdUser, err := t.Repository.CreateUserByResourceOwnerID(
		ctx,
		user,
		oauth2loginflow.ProviderID(providerID),
		oauth2loginflow.ResourceOwnerID(resourceOwnerID),
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return createdUser, nil
}

func (t *Impl) CreateGuestUser(
	ctx context.Context,
) (*common_entity.User, error) {
	userID, err := t.UserIDGenerator.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	user := common_entity.User{
		ID:              common_entity.UserID(userID),
		Name:            fmt.Sprintf("ゲストユーザー%d", t.NowFunc().Unix()),
		ProfileImageURL: "https://vos.line-scdn.net/chdev-console-static/default-profile.png",
		Guest:           true,
		Active:          true,
	}
	return t.Repository.CreateUser(ctx, &user)
}

func (t *Impl) GetUsers(
	ctx context.Context,
	userIDs []entity.UserID,
) ([]*entity.User, error) {
	return t.Repository.GetUsers(ctx, userIDs)
}

func (t *Impl) GetUser(
	ctx context.Context,
	userID common_entity.UserID,
) (*common_entity.User, error) {
	return t.Repository.GetUser(ctx, userID)
}
