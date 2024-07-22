package businesslogic

import (
	"context"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateChatRoom(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	chatRoomID common_entity.ChatRoomID,
) (*common_entity.ChatRoom, error) {
	root := common_entity.ChatRoom{
		ID:            chatRoomID,
		PhotoStudioID: photoStudioID,
	}
	return t.Repository.CreateChatRoom(ctx, &root)
}
