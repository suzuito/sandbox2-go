package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateChatRoom(
	ctx context.Context,
	room *common_entity.ChatRoom,
) (*common_entity.ChatRoom, error) {
	mRoom := newModelChatRoom(room)
	mRoom.CreatedAt = t.NowFunc()
	mRoom.UpdatedAt = mRoom.CreatedAt
	db := t.GormDB.WithContext(ctx)
	db = db.Create(mRoom)
	if err := db.Error; err != nil {
		return nil, terrors.Wrap(err)
	}
	return mRoom.ToEntity(), nil
}

func (t *Impl) GetChatRooms(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	q *cgorm.ListQuery,
) ([]*common_entity.ChatRoom, bool, error) {
	db := t.GormDB
	db = db.WithContext(ctx)
	db = db.Where("photo_studio_id = ?", photoStudioID)
	db = q.Set(db)
	mChatRooms := []modelChatRoom{}
	if err := db.Find(&mChatRooms).Error; err != nil {
		return nil, false, terrors.Wrap(err)
	}
	ret := []*common_entity.ChatRoom{}
	for _, a := range mChatRooms {
		ret = append(ret, a.ToEntity())
	}
	return ret, len(mChatRooms) >= q.Limit, nil
}
