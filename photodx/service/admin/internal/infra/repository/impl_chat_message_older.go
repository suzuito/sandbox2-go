package repository

import (
	"context"
	"errors"
	"time"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
)

var olderSortColumns = cgorm.SortColumns{
	{Name: "posted_at", Type: cgorm.Desc},
	{Name: "id", Type: cgorm.Desc},
}

func (t *Impl) GetOlderChatMessagesOffsetByID(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	chatMessageID common_entity.ChatMessageID,
) (int, error) {
	mChatMessage := modelChatMessage{}
	db := t.GormDB.
		WithContext(ctx).
		Where(
			"chat_room_id = ? AND id = ?",
			roomID,
			chatMessageID,
		).
		First(&mChatMessage)
	if err := db.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return -1, &common_repository.NoEntryError{
				EntryType: "ChatMessage",
				EntryID:   string(chatMessageID),
			}
		}
		return -1, terrors.Wrap(err)
	}
	return t.getOlderChatMessagesOffset(
		ctx,
		roomID,
		mChatMessage.PostedAt,
		mChatMessage.ID,
	)
}

func (t *Impl) getOlderChatMessagesOffset(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	postedAt time.Time,
	chatMessageID common_entity.ChatMessageID,
) (int, error) {
	count := int64(0)
	db := t.GormDB.
		WithContext(ctx).
		Model(&modelChatMessage{}).
		Where(
			"chat_room_id = ? AND posted_at > ? AND id > ?",
			roomID,
			postedAt,
			chatMessageID,
		)
	db = olderSortColumns.Set(db)
	db = db.Count(&count)
	if err := db.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return -1, &common_repository.NoEntryError{
				EntryType: "ChatMessage",
				EntryID:   string(chatMessageID),
			}
		}
		return -1, terrors.Wrap(err)
	}
	return int64ToInt(count)
}

func (t *Impl) GetOlderChatMessages(
	ctx context.Context,
	roomID common_entity.ChatRoomID,
	offset int,
	limit int,
) ([]*common_entity.ChatMessage, bool, int, bool, int, error) {
	listQuery := cgorm.ListQuery{
		Offset:      offset,
		Limit:       limit,
		SortColumns: olderSortColumns,
	}
	return t.getChatMessages(ctx, roomID, &listQuery)
}
