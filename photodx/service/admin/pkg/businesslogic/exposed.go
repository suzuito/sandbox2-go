package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
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
	GetChatMessages(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.ChatMessage, bool, error)
}
