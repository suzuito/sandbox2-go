package entity

type UserID string
type User struct {
	ID                UserID `json:"id"`
	Name              string `json:"name"`
	ProfileImageURL   string `json:"profileImageUrl"`
	Active            bool   `json:"active"`
	Guest             bool   `json:"guest"`
	InitializedByUser bool   `json:"initializeByUser"`
}
