package entity

type UserID string
type User struct {
	ID                UserID `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	EmailVerified     bool   `json:"email_verified"`
	ProfileImageURL   string `json:"profileImageUrl"`
	Active            bool   `json:"active"`
	Guest             bool   `json:"guest"`
	InitializedByUser bool   `json:"initializeByUser"`
}

type PromoteGuestUserConfirmationCode struct {
	UserID UserID `json:"user_id"`
	Email  string `json:"email"`
	Code   string `json:"code"`
}
