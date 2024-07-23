package businesslogic

import (
	"context"

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
