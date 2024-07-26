package entity

type NotificationID string
type NotificationType string

const (
	NotificationTypeChatMessage NotificationType = "PostChatMessage"
)

type Notification struct {
	ID                 NotificationID      `json:"id"`
	Type               NotificationType    `json:"type"`
	ChatMessageWrapper *ChatMessageWrapper `json:"chatMessage"`
}
