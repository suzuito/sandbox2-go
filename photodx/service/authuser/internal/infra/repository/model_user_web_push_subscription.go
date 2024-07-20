package repository

import (
	"encoding/json"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelUserWebPushSubscription struct {
	Endpoint       string               `gorm:"primaryKey;not null"`
	UserID         common_entity.UserID `gorm:"not null"`
	ExpirationTime time.Time            `gorm:"not null"`
	Value          string               `gorm:"not null"`
	CreatedAt      time.Time            `gorm:"not null"`
}

func (t *modelUserWebPushSubscription) TableName() string {
	return "users_web_push_subscriptions"
}

func (t *modelUserWebPushSubscription) ToEntity() (*entity.UserWebPushSubscription, error) {
	s := &webpush.Subscription{}
	if err := json.Unmarshal([]byte(t.Value), &s); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &entity.UserWebPushSubscription{
		Endpoint:       t.Endpoint,
		ExpirationTime: t.ExpirationTime,
		UserID:         t.UserID,
		Value:          s,
	}, nil
}

func newModelUserWebPushSubscription(s *entity.UserWebPushSubscription) (*modelUserWebPushSubscription, error) {
	v, err := json.Marshal(s.Value)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &modelUserWebPushSubscription{
		Endpoint:       s.Endpoint,
		ExpirationTime: s.ExpirationTime,
		UserID:         s.UserID,
		Value:          string(v),
	}, nil
}
