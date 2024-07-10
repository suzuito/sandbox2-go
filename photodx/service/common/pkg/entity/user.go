package entity

type UserID string
type User struct {
	ID              UserID
	Name            string
	ProfileImageURL string
}
