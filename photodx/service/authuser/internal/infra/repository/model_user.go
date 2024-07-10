package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelUser struct {
	ID              entity.UserID `gorm:"primaryKey;not null"`
	Name            string
	ProfileImageURL string
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
}

func (t *modelUser) TableName() string {
	return `users`
}

func (t *modelUser) ToEntity() *entity.User {
	return &entity.User{
		ID:              t.ID,
		Name:            t.Name,
		ProfileImageURL: t.ProfileImageURL,
	}
}

func NewModelUser(s *entity.User) *modelUser {
	return &modelUser{
		ID:              s.ID,
		Name:            s.Name,
		ProfileImageURL: s.ProfileImageURL,
	}
}
