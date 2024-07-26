package repository

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

func (t *Impl) CreatePhotoStudio(
	ctx context.Context,
	photoStudio *entity.PhotoStudio,
) (*entity.PhotoStudio, error) {
	mPhotoStudio := newModelPhotoStudio(photoStudio)
	mPhotoStudio.CreatedAt = t.NowFunc()
	mPhotoStudio.UpdatedAt = t.NowFunc()
	if err := t.GormDB.
		WithContext(ctx).
		Where(photoStudio.ID).
		First(&modelPhotoStudio{}).Error; err == nil {
		return nil, terrors.Wrap(&repository.DuplicateEntryError{
			EntryType: repository.EntryTypePhotoStudio,
			EntryID:   string(photoStudio.ID),
		})
	}
	if err := t.GormDB.Create(mPhotoStudio).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return photoStudio, nil
}

func (t *Impl) GetPhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
) (*entity.PhotoStudio, error) {
	mPhotoStudio := modelPhotoStudio{}
	if err := t.GormDB.WithContext(ctx).Where(photoStudioID).First(&mPhotoStudio).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: repository.EntryTypePhotoStudio,
				EntryID:   string(photoStudioID),
			})
		}
		return nil, terrors.Wrap(err)
	}
	return mPhotoStudio.ToEntity(), nil
}

func (t *Impl) GetPhotoStudios(
	ctx context.Context,
	photoStudioIDs []entity.PhotoStudioID,
) ([]*entity.PhotoStudio, error) {
	mPhotoStudios := []*modelPhotoStudio{}
	db := t.GormDB.WithContext(ctx)
	db = db.Where("id IN ?", photoStudioIDs)
	db = db.Find(&mPhotoStudios)
	if err := db.Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return arrayutil.Map(mPhotoStudios, func(v *modelPhotoStudio) *entity.PhotoStudio { return v.ToEntity() }), nil
}
