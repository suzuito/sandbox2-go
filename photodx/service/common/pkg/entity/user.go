package entity

import "time"

type UserID string
type User struct {
	ID              UserID `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerified   bool   `json:"email_verified"`
	ProfileImageURL string `json:"profileImageUrl"`
	Active          bool   `json:"active"`
	Guest           bool   `json:"guest"`
}

type UserCreationRequestID string
type UserCreationCode string

type UserCreationRequest struct {
	ID        UserCreationRequestID `json:"id"`
	Email     string                `json:"email"`
	Code      UserCreationCode      `json:"code"`
	ExpiredAt time.Time             `json:"expiredAt"`
}
