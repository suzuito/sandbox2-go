package repository

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func (t *Impl) CreateUserCreationRequest(
	ctx context.Context,
	request *common_entity.UserCreationRequest,
) error {
	now := t.NowFunc()
	m := newModelUserCreationRequest(
		request,
	)
	m.CreatedAt = now
	if err := t.GormDB.WithContext(ctx).Create(&m).Error; err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) GetUserCreationRequest(
	ctx context.Context,
	id common_entity.UserCreationRequestID,
) (*common_entity.UserCreationRequest, error) {
	m := modelUserCreationRequest{}
	if err := t.GormDB.Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &common_repository.NoEntryError{
				EntryType: "UserCreationRequest",
				EntryID:   string(id),
			}
		}
		return nil, terrors.Wrap(err)
	}
	return m.ToEntity(), nil
}

func (t *Impl) GetUserCreationRequestByEmail(
	ctx context.Context,
	email string,
) (*common_entity.UserCreationRequest, error) {
	m := modelUserCreationRequest{}
	if err := t.GormDB.Where("email = ?", email).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &common_repository.NoEntryError{
				EntryType: "UserCreationRequest",
				EntryID:   string(email),
			}
		}
		return nil, terrors.Wrap(err)
	}
	return m.ToEntity(), nil
}

func (t *Impl) DeleteUserCreationRequest(
	ctx context.Context,
	userCreationRequestID common_entity.UserCreationRequestID,
) error {
	if err := t.GormDB.Where("id = ?", userCreationRequestID).Delete(&modelUserCreationRequest{}).Error; err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
