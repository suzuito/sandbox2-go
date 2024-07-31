package repository

import (
	"time"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelPromoteGuestUserConfirmationCode struct {
	UserID    common_entity.UserID `gorm:"not null"`
	Email     string               `gorm:"not null"`
	Code      string               `gorm:"not null"`
	CreatedAt time.Time            `gorm:"not null"`
	ExpiredAt time.Time            `gorm:"not null"`
}

func (t *modelPromoteGuestUserConfirmationCode) TableName() string {
	return "promote_guest_user_confirmation_codes"
}

func (t *modelPromoteGuestUserConfirmationCode) ToEntity() *common_entity.PromoteGuestUserConfirmationCode {
	return &common_entity.PromoteGuestUserConfirmationCode{
		UserID: t.UserID,
		Email:  t.Email,
		Code:   t.Code,
	}
}

func newModelPromoteGuestUserConfirmationCode(s *common_entity.PromoteGuestUserConfirmationCode) *modelPromoteGuestUserConfirmationCode {
	return &modelPromoteGuestUserConfirmationCode{
		UserID: s.UserID,
		Email:  s.Email,
		Code:   s.Code,
	}
}
