package businesslogic

import (
	"context"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type ExposedBusinessLogic interface {
	CreatePhotoStudioUserChatRoomIFNotExists(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
	) (*common_entity.ChatRoom, error)
	CreateChatMessage(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		message *common_entity.ChatMessage,
	) (*common_entity.ChatMessage, error)
	GetOlderChatMessagesForFront(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		offset int,
		limit int,
	) ([]*common_entity.ChatMessage, bool, int, bool, int, error)
	GetOlderChatMessagesForFrontByID(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		chatMessageID common_entity.ChatMessageID,
		limit int,
	) ([]*common_entity.ChatMessage, bool, int, bool, int, error)
}
