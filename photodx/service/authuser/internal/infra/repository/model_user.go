package repository

import (
	"time"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelUser struct {
	ID              entity.UserID `gorm:"primaryKey;not null"`
	Name            string
	Email           string
	EmailVerified   bool
	ProfileImageURL string
	Active          bool      `gorm:"not null"`
	Guest           bool      `gorm:"not null"`
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
		Email:           t.Email,
		EmailVerified:   t.EmailVerified,
		ProfileImageURL: t.ProfileImageURL,
		Active:          t.Active,
		Guest:           t.Guest,
	}
}

func NewModelUser(s *entity.User) *modelUser {
	return &modelUser{
		ID:              s.ID,
		Name:            s.Name,
		Email:           s.Email,
		EmailVerified:   s.EmailVerified,
		ProfileImageURL: s.ProfileImageURL,
		Active:          s.Active,
		Guest:           s.Guest,
	}
}
