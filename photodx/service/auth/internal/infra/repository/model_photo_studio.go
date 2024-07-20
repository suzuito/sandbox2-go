package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelPhotoStudio struct {
	ID        entity.PhotoStudioID `gorm:"primaryKey"`
	Name      string               `gorm:"not null"`
	Active    bool                 `gorm:"not null;default:false"`
	CreatedAt time.Time            `gorm:"not null"`
	UpdatedAt time.Time            `gorm:"not null"`
}

func (t *modelPhotoStudio) TableName() string {
	return "photo_studios"
}

func (t *modelPhotoStudio) ToEntity() *entity.PhotoStudio {
	return &entity.PhotoStudio{
		ID:     t.ID,
		Name:   t.Name,
		Active: t.Active,
	}
}

func newModelPhotoStudio(s *entity.PhotoStudio) *modelPhotoStudio {
	return &modelPhotoStudio{
		ID:     s.ID,
		Name:   s.Name,
		Active: s.Active,
	}
}
