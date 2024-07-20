package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"gorm.io/gorm/clause"
)

func (t *Impl) GetLatestUserWebPushSubscriptions(
	ctx context.Context,
	userID common_entity.UserID,
) ([]*entity.UserWebPushSubscription, error) {
	m := []*modelUserWebPushSubscription{}
	err := t.GormDB.WithContext(ctx).
		Order("created_at DESC").
		Limit(5).
		Where("user_id = ? && (expiration_time = null || expiration_time > ?)", userID, t.NowFunc()).
		Find(&m).
		Error
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	ret := arrayutil.
		Map(m, func(i *modelUserWebPushSubscription) *entity.UserWebPushSubscription {
			e, err := i.ToEntity()
			if err != nil {
				return nil
			}
			return e
		})
	ret = arrayutil.Filter(ret, func(i *entity.UserWebPushSubscription) bool { return i != nil })
	return ret, nil
}

func (t *Impl) UpdateOrCreateUserWebPushSubscription(
	ctx context.Context,
	s *entity.UserWebPushSubscription,
) (*entity.UserWebPushSubscription, error) {
	m, err := newModelUserWebPushSubscription(s)
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
