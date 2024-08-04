package repository

import (
	"time"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelUserCreationRequest struct {
	ID        common_entity.UserCreationRequestID `gorm:"not null"`
	Email     string                              `gorm:"not null"`
	Code      common_entity.UserCreationCode      `gorm:"not null"`
	CreatedAt time.Time                           `gorm:"not null"`
	ExpiredAt time.Time                           `gorm:"not null"`
}

func (t *modelUserCreationRequest) TableName() string {
	return "user_creation_requests"
}

func (t *modelUserCreationRequest) ToEntity() *common_entity.UserCreationRequest {
	return &common_entity.UserCreationRequest{
		ID:        t.ID,
		Email:     t.Email,
		Code:      t.Code,
		ExpiredAt: t.ExpiredAt,
	}
}

func newModelUserCreationRequest(s *common_entity.UserCreationRequest) *modelUserCreationRequest {
	return &modelUserCreationRequest{
		ID:        s.ID,
		Email:     s.Email,
		Code:      s.Code,
		ExpiredAt: s.ExpiredAt,
	}
}
