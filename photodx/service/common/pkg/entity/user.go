package entity

type UserID string
type User struct {
	ID                UserID `json:"id"`
	Name              string `json:"name"`
	ProfileImageURL   string `json:"profileImageUrl"`
	Active            bool   `json:"active"`
	InitializedByUser bool   `json:"initializeByUser"`
}
