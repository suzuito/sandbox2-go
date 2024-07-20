package repository

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func (t *Impl) GetLineLinkInfo(ctx context.Context, photoStudioID common_entity.PhotoStudioID) (*entity.LineLinkInfo, error) {
	mLineLinkInfo := modelLineLinkInfo{}
	if err := t.GormDB.
		WithContext(ctx).
		Where("photo_studio_id = ?", photoStudioID).
		First(&mLineLinkInfo).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &common_repository.NoEntryError{
				EntryType: "LineLinkInfo",
				EntryID:   string(photoStudioID),
			}
		}
		return nil, terrors.Wrap(err)
	}
	return mLineLinkInfo.ToEntity(), nil
}

func (t *Impl) CreateLineLinkInfo(ctx context.Context, info *entity.LineLinkInfo) (*entity.LineLinkInfo, error) {
	if err := t.GormDB.WithContext(ctx).First(newModelLineLinkInfo(info)).Error; err == nil {
		return nil, &common_repository.DuplicateEntryError{
			EntryType: "LineLinkInfo",
			EntryID:   string(info.PhotoStudioID),
		}
	}
	mLineLinkInfo := newModelLineLinkInfo(info)
	mLineLinkInfo.CreatedAt = t.NowFunc()
	mLineLinkInfo.UpdatedAt = t.NowFunc()
	if err := t.GormDB.WithContext(ctx).Create(&mLineLinkInfo).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mLineLinkInfo.ToEntity(), nil
}

func (t *Impl) DeleteLineLinkInfo(ctx context.Context, photoStudioID common_entity.PhotoStudioID) error {
	if err := t.GormDB.WithContext(ctx).Where("photo_studio_id = ?", photoStudioID).Delete(&modelLineLinkInfo{}).Error; err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) setLineLinkInfo(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	f func(m *modelLineLinkInfo),
) (*entity.LineLinkInfo, error) {
	if err := t.GormDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		mLineLinkInfo := modelLineLinkInfo{}
		if err := tx.Where("photo_studio_id = ?", photoStudioID).First(&mLineLinkInfo).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &common_repository.NoEntryError{
					EntryType: "LineLinkInfo",
					EntryID:   string(photoStudioID),
				}
			}
			return terrors.Wrap(err)
		}
		f(&mLineLinkInfo)
		mLineLinkInfo.UpdatedAt = t.NowFunc()
		if err := tx.Save(&mLineLinkInfo).Error; err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, terrors.Wrap(err)
	}
	return t.GetLineLinkInfo(ctx, photoStudioID)
}

func (t *Impl) SetLineLinkInfoActive(ctx context.Context, photoStudioID common_entity.PhotoStudioID, active bool) (*entity.LineLinkInfo, error) {
	return t.setLineLinkInfo(ctx, photoStudioID, func(m *modelLineLinkInfo) {
		m.Active = active
	})
}

func (t *Impl) SetLineLinkInfo(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	arg *repository.SetLineLinkInfoArgument,
) (*entity.LineLinkInfo, error) {
	return t.setLineLinkInfo(ctx, photoStudioID, func(m *modelLineLinkInfo) {
		if arg.MessagingAPIChannelSecret != nil {
			m.MessagingAPIChannelSecret = *arg.MessagingAPIChannelSecret
		}
		if arg.LongAccessToken != nil {
			m.LongAccessToken = *arg.LongAccessToken
		}
	})
}
