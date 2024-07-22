package entity

type ChatRoomID string

type ChatRoom struct {
	ID            ChatRoomID    `json:"id"`
	PhotoStudioID PhotoStudioID `json:"photoStudioId"`
}
