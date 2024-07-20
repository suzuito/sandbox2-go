package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func (t *Impl) CreateUserByResourceOwnerID(
	ctx context.Context,
	user *entity.User,
	providerID oauth2loginflow.ProviderID,
	resourceOwnerID oauth2loginflow.ResourceOwnerID,
) (*entity.User, error) {
	mUser := NewModelUser(user)
	if err := t.GormDB.WithContext(ctx).Where(user.ID).First(&modelUser{}).Error; err == nil {
		return nil, &repository.DuplicateEntryError{
			EntryType: repository.EntryTypeUser,
			EntryID:   string(user.ID),
		}
	}
	if err := t.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(mUser).Error; err != nil {
			return terrors.Wrap(err)
		}
		m := modelProviderResourceOwnersUsersMapping{
			ProviderID:      providerID,
			ResourceOwnerID: resourceOwnerID,
			UserID:          user.ID,
			CreatedAt:       t.NowFunc(),
		}
		if err := tx.Create(&m).Error; err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, terrors.Wrap(err)
	}
	return user, nil
}

func (t *Impl) GetUserByResourceOwnerID(
	ctx context.Context,
	providerID oauth2loginflow.ProviderID,
	resourceOwnerID oauth2loginflow.ResourceOwnerID,
) (*entity.User, error) {
	mProviderResourceOwnersUsersMapping := modelProviderResourceOwnersUsersMapping{}
	if err := t.GormDB.
		WithContext(ctx).
		Where(
			"provider_id = ? AND resource_owner_id = ?",
			providerID,
			resourceOwnerID,
		).
		First(&mProviderResourceOwnersUsersMapping).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: "ProviderResourceOwnersUsersMappings",
				EntryID:   fmt.Sprintf("%s.%s", providerID, resourceOwnerID),
			})
		}
		return nil, terrors.Wrap(err)
	}
	return t.GetUser(ctx, entity.UserID(mProviderResourceOwnersUsersMapping.UserID))
}

func (t *Impl) GetUser(
	ctx context.Context,
	userID entity.UserID,
) (*entity.User, error) {
	mUser := modelUser{}
	if err := t.GormDB.Where(userID).First(&mUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: repository.EntryTypeUser,
				EntryID:   string(userID),
			})
		}
		return nil, terrors.Wrap(err)
	}
	return mUser.ToEntity(), nil
}

func (t *Impl) GetUsers(
	ctx context.Context,
	userIDs []entity.UserID,
) ([]*entity.User, error) {
	mUsers := []*modelUser{}
	if err := t.GormDB.WithContext(ctx).Where("id in ?", userIDs).Find(&mUsers).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	ret := []*entity.User{}
	for _, u := range mUsers {
		ret = append(ret, u.ToEntity())
	}
	return ret, nil
}
