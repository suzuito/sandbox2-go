package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"gorm.io/gorm/clause"
)

func (t *Impl) GetLatestPhotoStudioMemberWebPushSubscriptions(
	ctx context.Context,
	photoStudioMemberID common_entity.PhotoStudioMemberID,
) ([]*entity.PhotoStudioMemberWebPushSubscription, error) {
	m := []*modelPhotoStudioMemberWebPushSubscription{}
	err := t.GormDB.WithContext(ctx).
		Order("created_at DESC").
		Limit(5).
		Where("photo_studio_member_id = ? && (expiration_time = null || expiration_time > ?)", photoStudioMemberID, t.NowFunc()).
		Find(&m).
		Error
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	ret := arrayutil.
		Map(m, func(i *modelPhotoStudioMemberWebPushSubscription) *entity.PhotoStudioMemberWebPushSubscription {
			e, err := i.ToEntity()
			if err != nil {
				return nil
			}
			return e
		})
	ret = arrayutil.Filter(ret, func(i *entity.PhotoStudioMemberWebPushSubscription) bool { return i != nil })
	return ret, nil
}

func (t *Impl) UpdateOrCreateUserWebPushSubscription(
	ctx context.Context,
	s *entity.PhotoStudioMemberWebPushSubscription,
) (*entity.PhotoStudioMemberWebPushSubscription, error) {
	m, err := newModelPhotoStudioMemberWebPushSubscription(s)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	m.CreatedAt = t.NowFunc()
	if err := t.GormDB.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(m).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return s, nil
}
