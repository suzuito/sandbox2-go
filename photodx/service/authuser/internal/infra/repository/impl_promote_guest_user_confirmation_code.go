package repository

import (
	"context"
	"errors"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (t *Impl) CreatePromoteGuestUserConfirmationCode(
	ctx context.Context,
	userID common_entity.UserID,
	email string,
	code string,
	ttlSeconds int,
) error {
	now := t.NowFunc()
	m := modelPromoteGuestUserConfirmationCode{
		UserID:    userID,
		Email:     email,
		Code:      code,
		CreatedAt: now,
		ExpiredAt: now.Add(time.Duration(ttlSeconds) * time.Second),
	}
	if err := t.GormDB.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		},
	).Create(&m).Error; err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) GetPromoteGuestUserConfirmationCodeNotExpired(
	ctx context.Context,
	userID common_entity.UserID,
) (*common_entity.PromoteGuestUserConfirmationCode, error) {
	m := modelPromoteGuestUserConfirmationCode{}
	if err := t.GormDB.Where("user_id = ? AND expired_at >= ?", userID, t.NowFunc()).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &common_repository.NoEntryError{
				EntryType: "PromoteGuestUserConfirmationCode",
				EntryID:   string(userID),
			}
		}
		return nil, terrors.Wrap(err)
	}
	return m.ToEntity(), nil
}
