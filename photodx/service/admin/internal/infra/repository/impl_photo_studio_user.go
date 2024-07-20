package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"gorm.io/gorm/clause"
)

func (t *Impl) CreatePhotoStudioUser(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	user *common_entity.User,
) (*entity.PhotoStudioUser, error) {
	mPhotoStudioUser := newModelPhotoStudioUser(&entity.PhotoStudioUser{
		UserID:        user.ID,
		PhotoStudioID: photoStudioID,
	})
	if err := t.GormDB.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&mPhotoStudioUser).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mPhotoStudioUser.ToEntity(), nil
}

func (t *Impl) GetPhotoStudioUsers(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	q *cgorm.ListQuery,
) ([]*entity.PhotoStudioUser, bool, error) {
	db := t.GormDB
	db = db.WithContext(ctx)
	db = db.Where("photo_studio_id = ?", photoStudioID)
	db = q.Set(db)
	mPhotoStudioUsers := []modelPhotoStudioUser{}
	if err := db.Find(&mPhotoStudioUsers).Error; err != nil {
		return nil, false, terrors.Wrap(err)
	}
	ret := []*entity.PhotoStudioUser{}
	for _, a := range mPhotoStudioUsers {
		ret = append(ret, a.ToEntity())
	}
	return ret, len(mPhotoStudioUsers) >= q.Limit, nil
}

func (t *Impl) GetPhotoStudioUser(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
) (*entity.PhotoStudioUser, error) {
	db := t.GormDB
	db = db.WithContext(ctx)
	db = db.Where("photo_studio_id = ?", photoStudioID)
	db = db.Where("user_id = ?", userID)
	mPhotoStudioUser := modelPhotoStudioUser{}
	if err := db.First(&mPhotoStudioUser).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mPhotoStudioUser.ToEntity(), nil
}
