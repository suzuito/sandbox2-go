package repository

import (
	"time"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelChatMessage struct {
	ID           common_entity.ChatMessageID           `gorm:"primaryKey;not null;"`
	ChatRoomID   common_entity.ChatRoomID              `gorm:"not null"`
	Type         common_entity.ChatMessageType         `gorm:"not null"`
	Text         string                                `gorm:"not null"`
	PostedBy     string                                `gorm:"not null"`
	PostedByType common_entity.ChatMessagePostedByType `gorm:"not null"`
	PostedAt     time.Time                             `gorm:"not null"`
	CreatedAt    time.Time                             `gorm:"not null"`
	UpdatedAt    time.Time                             `gorm:"not null"`

	// Associationsの種類によってforeignKeyの意味が変わるからややこしい。belongsToとhasMany
	// (belongsTo)ChatMessageはChatRoomに属する。
	// `gorm:"foreignKey:ChatRoomID"`のChatRoomIDは、modelChatMessage(このテーブル)のChatRoomIDを指す。
	ChatRoom *modelChatRoom `gorm:"foreignKey:ChatRoomID"`
}

func (t *modelChatMessage) TableName() string {
	return "chat_messages"
}

func (t *modelChatMessage) ToEntity() *common_entity.ChatMessage {
	return &common_entity.ChatMessage{
		ID:           t.ID,
		Type:         t.Type,
		Text:         t.Text,
		PostedBy:     t.PostedBy,
		PostedByType: t.PostedByType,
		PostedAt:     common_entity.WTime(t.PostedAt),
	}
}

func newModelChatMessage(s *common_entity.ChatMessage) *modelChatMessage {
	return &modelChatMessage{
		ID:           s.ID,
		Type:         s.Type,
		Text:         s.Text,
		PostedBy:     s.PostedBy,
		PostedByType: s.PostedByType,
		PostedAt:     time.Time(s.PostedAt),
	}
}
