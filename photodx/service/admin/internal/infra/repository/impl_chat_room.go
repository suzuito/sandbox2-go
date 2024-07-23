package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_repository "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (t *Impl) CreatePhotoStudioUserChatRoomIFNotExists(
	ctx context.Context,
	room *common_entity.ChatRoom,
) (*common_entity.ChatRoom, error) {
	mRoom := newModelChatRoom(room)
	mRoom.CreatedAt = t.NowFunc()
	mRoom.UpdatedAt = mRoom.CreatedAt
	db := t.GormDB.WithContext(ctx)
	db = db.Clauses(clause.OnConflict{
		DoNothing: true,
	})
	db = db.Create(mRoom)
	if err := db.Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mRoom.ToEntity(), nil
}

func (t *Impl) GetChatRoomByPhotoStudioIDANDUserID(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
) (*common_entity.ChatRoom, error) {
	db := t.GormDB
	db = db.WithContext(ctx)
	db = db.Where("photo_studio_id = ? && user_id = ?", photoStudioID, userID)
	mChatRoom := modelChatRoom{}
	if err := db.First(&mChatRoom).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &common_repository.NoEntryError{
				EntryType: "ChatRoom",
				EntryID:   fmt.Sprintf("%s-%s", photoStudioID, userID),
			}
		}
		return nil, terrors.Wrap(err)
	}
	return mChatRoom.ToEntity(), nil
}
