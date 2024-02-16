package entity

type UserID string

type User struct {
	ID      UserID
	Avatar  Avatar
	Feeling Feeling
	Name    string
}

type FeelingID string

type Feeling struct {
	ID   FeelingID
	Name string
}

type AvatarID string

// Avatar is image of user.
// Selectable avatars are predefined in DB.
type Avatar struct {
	ID AvatarID
}
