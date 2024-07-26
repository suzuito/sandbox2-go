package businesslogic

import (
	"context"

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

func beforeAndAfterRange(offset, limit int) (int, int) {
	d := limit / 2
	offset = offset - d
	if offset < 0 {
		offset = 0
	}
	return offset, limit
}

func (t *Impl) GetOlderChatMessagesForFront(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	offset int,
	limit int,
) ([]*common_entity.ChatMessage, bool, int, bool, int, error) {
	room, err := t.Repository.GetChatRoomByPhotoStudioIDANDUserID(
		ctx,
		photoStudioID,
		userID,
	)
	if err != nil {
		return nil, false, 0, false, 0, terrors.Wrap(err)
	}
	return t.Repository.GetOlderChatMessages(
		ctx,
		room.ID,
		offset,
		limit,
	)
}

func (t *Impl) GetOlderChatMessagesForFrontByID(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	chatMessageID common_entity.ChatMessageID,
	limit int,
) ([]*common_entity.ChatMessage, bool, int, bool, int, error) {
	room, err := t.Repository.GetChatRoomByPhotoStudioIDANDUserID(
		ctx,
		photoStudioID,
		userID,
	)
	if err != nil {
		return nil, false, 0, false, 0, terrors.Wrap(err)
	}
	offset, err := t.Repository.GetOlderChatMessagesOffsetByID(
		ctx,
		room.ID,
		chatMessageID,
	)
	if err != nil {
		return nil, false, 0, false, 0, terrors.Wrap(err)
	}
	offset, limit = beforeAndAfterRange(offset, limit)
	return t.Repository.GetOlderChatMessages(
		ctx,
		room.ID,
		offset,
		limit,
	)
}
