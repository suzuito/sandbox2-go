package repository

import (
	"time"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type modelChatRoom struct {
	ID            common_entity.ChatRoomID    `gorm:"primaryKey;not null"`
	PhotoStudioID common_entity.PhotoStudioID `gorm:"not null"`
	CreatedAt     time.Time                   `gorm:"not null"`
	UpdatedAt     time.Time                   `gorm:"not null"`
}

func (t *modelChatRoom) TableName() string {
	return "chat_rooms"
}

func (t *modelChatRoom) ToEntity() *common_entity.ChatRoom {
	return &common_entity.ChatRoom{
		ID:            t.ID,
		PhotoStudioID: t.PhotoStudioID,
	}
}

func newModelChatRoom(s *common_entity.ChatRoom) *modelChatRoom {
	return &modelChatRoom{
		ID:            s.ID,
		PhotoStudioID: s.PhotoStudioID,
	}
}
