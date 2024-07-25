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

func (t *Impl) GetChatMessages(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	listQuery *cgorm.ListQuery,
) ([]*common_entity.ChatMessage, bool, error) {
	db := t.GormDB.WithContext(ctx)
	db = db.Where(
		"chat_room_id = ?",
		roomID,
	)
	db = listQuery.Set(db)
	mMessages := []*modelChatMessage{}
	if err := db.Find(&mMessages).Error; err != nil {
		return nil, false, terrors.Wrap(err)
	}
	hasNext := len(mMessages) >= listQuery.Limit
	messages := arrayutil.Map(mMessages, func(v *modelChatMessage) *common_entity.ChatMessage {
		return v.ToEntity()
	})
	return messages, hasNext, nil
}

func (t *Impl) GetOlderChatMessages(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	offset int,
	limit int,
) ([]*common_entity.ChatMessage, bool, int, error) {
	listQuery := cgorm.ListQuery{
		Offset: offset,
		Limit:  limit,
		SortColumns: []cgorm.SortColumn{
			{Name: "posted_at", Type: cgorm.Desc},
			{Name: "id", Type: cgorm.Desc},
		},
	}
	db := t.GormDB.WithContext(ctx)
	db = db.Where(
		"chat_room_id = ?",
		roomID,
	)
	db = listQuery.Set(db)
	mMessages := []*modelChatMessage{}
	if err := db.Find(&mMessages).Error; err != nil {
		return nil, false, 0, terrors.Wrap(err)
	}
	hasNext := len(mMessages) >= listQuery.Limit
	messages := arrayutil.Map(mMessages, func(v *modelChatMessage) *common_entity.ChatMessage {
		return v.ToEntity()
	})
	return messages, hasNext, listQuery.NextOffset(), nil
}

func (t *Impl) GetChatMessagesByTimeRange(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	offset *common_entity.ChatMessageOffset,
	limit int,
	isGetOlder bool,
) ([]*common_entity.ChatMessage, bool, error) {
	boundType := cgorm.ListQuery2BoundTypeUpper
	dir := cgorm.ListQuery2KeyDirectionAsc
	if isGetOlder {
		boundType = cgorm.ListQuery2BoundTypeLower
		dir = cgorm.ListQuery2KeyDirectionDesc
	}
	listQuery := cgorm.ListQuery2{
		Limit: limit,
		Range: cgorm.ListQuery2KeyRange{
			Bounds: []cgorm.ListQuery2Bound{
				{
					Type:      boundType,
					KeyName:   "posted_at",
					Direction: dir,
					Open:      true,
					Value:     offset.PostedAt,
				},
				{
					Type:      boundType,
					KeyName:   "id",
					Direction: dir,
					Open:      true,
					Value:     offset.ID,
				},
			},
		},
	}
	db := t.GormDB.WithContext(ctx)
	db = db.Where(
		"chat_room_id = ?",
		roomID,
	)
	db = listQuery.Set(db)
	mMessages := []*modelChatMessage{}
	if err := db.Find(&mMessages).Error; err != nil {
		return nil, false, terrors.Wrap(err)
	}
	hasNext := len(mMessages) >= listQuery.Limit
	messages := arrayutil.Map(mMessages, func(v *modelChatMessage) *common_entity.ChatMessage {
		return v.ToEntity()
	})
	return messages, hasNext, nil
}
