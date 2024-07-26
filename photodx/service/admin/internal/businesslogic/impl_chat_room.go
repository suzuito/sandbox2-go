package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreatePhotoStudioUserChatRoomIFNotExists(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	userID entity.UserID,
) (*common_entity.ChatRoom, error) {
	root := common_entity.ChatRoom{
		ID:            common_entity.ChatRoomID(photoStudioID + "-" + common_entity.PhotoStudioID(userID)),
		PhotoStudioID: photoStudioID,
		UserID:        userID,
	}
	return t.Repository.CreatePhotoStudioUserChatRoomIFNotExists(ctx, &root)
}

func (t *Impl) GetPhotoStudioChats(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	listQuery *cgorm.ListQuery,
) ([]*common_entity.ChatRoom, bool, error) {
	return t.Repository.GetPhotoStudioChats(ctx, photoStudioID, listQuery)
}

func (t *Impl) GetPhotoStudioChat(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
) (*common_entity.ChatRoom, error) {
	return t.Repository.GetChatRoomByPhotoStudioIDANDUserID(ctx, photoStudioID, userID)
}
