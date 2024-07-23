package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateChatMessage(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	message *common_entity.ChatMessage,
) (*common_entity.ChatMessage, error) {
	room, err := t.Repository.GetChatRoomByPhotoStudioIDANDUserID(
		ctx,
		photoStudioID,
		userID,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	messageID, err := t.GenerateChatMessageID.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	message.ID = common_entity.ChatMessageID(messageID)
	return t.Repository.CreateChatMessage(
		ctx,
		room.ID,
		message,
	)
}

func (t *Impl) GetChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	listQuery *cgorm.ListQuery,
) ([]*common_entity.ChatMessage, bool, error) {
	room, err := t.Repository.GetChatRoomByPhotoStudioIDANDUserID(
		ctx,
		photoStudioID,
		userID,
	)
	if err != nil {
		return nil, false, terrors.Wrap(err)
	}
	return t.Repository.GetChatMessages(
		ctx,
		room.ID,
		listQuery,
	)
}
