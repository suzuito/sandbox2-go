package repository

import (
	"encoding/json"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelPhotoStudioMemberWebPushSubscription struct {
	Endpoint            string                            `gorm:"primaryKey;not null"`
	PhotoStudioMemberID common_entity.PhotoStudioMemberID `gorm:"not null"`
	ExpirationTime      time.Time                         `gorm:"not null"`
	Value               string                            `gorm:"not null"`
	CreatedAt           time.Time                         `gorm:"not null"`
}

func (t *modelPhotoStudioMemberWebPushSubscription) TableName() string {
	return "photo_studio_members_web_push_subscriptions"
}

func (t *modelPhotoStudioMemberWebPushSubscription) ToEntity() (*entity.PhotoStudioMemberWebPushSubscription, error) {
	s := &webpush.Subscription{}
	if err := json.Unmarshal([]byte(t.Value), &s); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &entity.PhotoStudioMemberWebPushSubscription{
		Endpoint:            t.Endpoint,
		ExpirationTime:      t.ExpirationTime,
		PhotoStudioMemberID: t.PhotoStudioMemberID,
		Value:               s,
	}, nil
}

func newModelPhotoStudioMemberWebPushSubscription(s *entity.PhotoStudioMemberWebPushSubscription) (*modelPhotoStudioMemberWebPushSubscription, error) {
	v, err := json.Marshal(s.Value)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &modelPhotoStudioMemberWebPushSubscription{
		Endpoint:            s.Endpoint,
		ExpirationTime:      s.ExpirationTime,
		PhotoStudioMemberID: s.PhotoStudioMemberID,
		Value:               string(v),
	}, nil
}
