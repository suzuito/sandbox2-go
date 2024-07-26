package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/arrayutil"
	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateChatMessage(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	message *common_entity.ChatMessage,
) (*common_entity.ChatMessage, error) {
	mMessage := newModelChatMessage(message)
	mMessage.CreatedAt = t.NowFunc()
	mMessage.ChatRoomID = roomID
	if err := t.GormDB.WithContext(ctx).Create(mMessage).Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mMessage.ToEntity(), nil
}

func (t *Impl) getChatMessages(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	listQuery *cgorm.ListQuery,
) ([]*common_entity.ChatMessage, bool, int, bool, int, error) {
	db := t.GormDB.WithContext(ctx)
	db = db.Where(
		"chat_room_id = ?",
		roomID,
	)
	db = listQuery.Set(db)
	mMessages := []*modelChatMessage{}
	if err := db.Find(&mMessages).Error; err != nil {
		return nil, false, 0, false, 0, terrors.Wrap(err)
	}
	messages := arrayutil.Map(mMessages, func(v *modelChatMessage) *common_entity.ChatMessage {
		return v.ToEntity()
	})
	hasNext := len(mMessages) >= listQuery.Limit
	return messages, hasNext, listQuery.NextOffset(), listQuery.HasPrev(), listQuery.PrevOffset(), nil
}
