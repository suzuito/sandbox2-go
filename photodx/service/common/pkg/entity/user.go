package entity

import (
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type UserID string
type User struct {
	ID              UserID `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	EmailVerified   bool   `json:"email_verified"`
	ProfileImageURL string `json:"profileImageUrl" validate:"required"`
	Active          bool   `json:"active"`
}

func (t *User) Validate() error {
	return terrors.Wrap(validate.Struct(t))
}

type UserCreationRequestID string
type UserCreationCode string

type UserCreationRequest struct {
	ID        UserCreationRequestID `json:"id"`
	Email     string                `json:"email"`
	Code      UserCreationCode      `json:"code"`
	ExpiredAt time.Time             `json:"expiredAt"`
}
