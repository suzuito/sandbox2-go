package repository

import (
	"time"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelUserPasswordHashValue struct {
	UserID    common_entity.UserID `gorm:"primaryKey;not null"`
	Value     string               `gorm:"not null"`
	CreatedAt time.Time            `gorm:"not null"`
	UpdatedAt time.Time            `gorm:"not null"`
}

func (t *modelUserPasswordHashValue) TableName() string {
	return "user_password_hash_values"
}
