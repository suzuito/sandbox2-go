package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type ExposedBusinessLogic interface {
	CreateChatMessage(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		message *common_entity.ChatMessage,
	) (*common_entity.ChatMessage, error)
	GetChatMessages(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.ChatMessage, bool, error)
}
