package entity

type ChatRoomID string

type ChatRoom struct {
	ID            ChatRoomID    `json:"id"`
	PhotoStudioID PhotoStudioID `json:"photoStudioId"`
	UserID        UserID        `json:"userId"`
}

type ChatRoomWrapper struct {
	*ChatRoom
	PhotoStudio *PhotoStudio `json:"photoStudio"`
	User        *User        `json:"user"`
}
