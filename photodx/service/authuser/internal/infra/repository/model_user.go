package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelUser struct {
	ID                entity.UserID `gorm:"primaryKey;not null"`
	Name              string
	ProfileImageURL   string
	Active            bool      `gorm:"not null"`
	Guest             bool      `gorm:"not null"`
	InitializedByUser bool      `gorm:"not null"`
	CreatedAt         time.Time `gorm:"not null"`
	UpdatedAt         time.Time `gorm:"not null"`
}

func (t *modelUser) TableName() string {
	return `users`
}

func (t *modelUser) ToEntity() *entity.User {
	return &entity.User{
		ID:                t.ID,
		Name:              t.Name,
		ProfileImageURL:   t.ProfileImageURL,
		Active:            t.Active,
		Guest:             t.Guest,
		InitializedByUser: t.InitializedByUser,
	}
}

func NewModelUser(s *entity.User) *modelUser {
	return &modelUser{
		ID:                s.ID,
		Name:              s.Name,
		ProfileImageURL:   s.ProfileImageURL,
		Active:            s.Active,
		Guest:             s.Guest,
		InitializedByUser: s.InitializedByUser,
	}
}
