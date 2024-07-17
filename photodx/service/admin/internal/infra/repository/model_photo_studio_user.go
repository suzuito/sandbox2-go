package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelPhotoStudioUser struct {
	PhotoStudioID common_entity.PhotoStudioID `gorm:"not null"`
	UserID        common_entity.UserID        `gorm:"not null"`
	CreatedAt     time.Time                   `gorm:"not null"`
	UpdatedAt     time.Time                   `gorm:"not null"`
}

func (t *modelPhotoStudioUser) TableName() string {
	return "photo_studio_users"
}

func (t *modelPhotoStudioUser) ToEntity() *entity.PhotoStudioUser {
	return &entity.PhotoStudioUser{
		PhotoStudioID: t.PhotoStudioID,
		UserID:        t.UserID,
	}
}

func newModelPhotoStudioUser(s *entity.PhotoStudioUser) *modelPhotoStudioUser {
	return &modelPhotoStudioUser{
		PhotoStudioID: s.PhotoStudioID,
		UserID:        s.UserID,
	}
}
