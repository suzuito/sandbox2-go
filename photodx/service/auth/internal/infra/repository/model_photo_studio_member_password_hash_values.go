package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelPhotoStudioMemberPasswordHashValue struct {
	PhotoStudioMemberID entity.PhotoStudioMemberID `gorm:"primaryKey;not null"`
	Value               string                     `gorm:"not null"`
	CreatedAt           time.Time                  `gorm:"not null"`
	UpdatedAt           time.Time                  `gorm:"not null"`
}

func (t *modelPhotoStudioMemberPasswordHashValue) TableName() string {
	return "photo_studio_member_password_hash_values"
}
