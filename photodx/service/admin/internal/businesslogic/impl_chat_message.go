package businesslogic

import (
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) CreateChatMessage(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	message *common_entity.ChatMessage,
) (*common_entity.ChatMessage, error) {
	rooms, _, err := t.Repository.GetChatRooms(
		ctx,
		photoStudioID,
		&cgorm.ListQuery{
			Offset: 0,
			Limit:  100,
			SortColumns: []cgorm.SortColumn{
				{Name: "created_at", Type: cgorm.Desc},
			},
		},
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if len(rooms) != 1 {
		return nil, terrors.Wrap(fmt.Errorf("invalid chat room state"))
	}
	messageID, err := t.GenerateChatMessageID.Gen()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	message.ID = common_entity.ChatMessageID(messageID)
	return t.Repository.CreateChatMessage(
		ctx,
		rooms[0].ID,
		message,
	)
}

func (t *Impl) GetChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	listQuery *cgorm.ListQuery,
) ([]*common_entity.ChatMessage, bool, error) {
	rooms, _, err := t.Repository.GetChatRooms(
		ctx,
		photoStudioID,
		&cgorm.ListQuery{
			Offset: 0,
			Limit:  100,
			SortColumns: []cgorm.SortColumn{
				{Name: "created_at", Type: cgorm.Desc},
			},
		},
	)
	if err != nil {
		return nil, false, terrors.Wrap(err)
	}
	if len(rooms) != 1 {
		return nil, false, terrors.Wrap(fmt.Errorf("invalid chat room state"))
	}
	return t.Repository.GetChatMessages(
		ctx,
		rooms[0].ID,
		listQuery,
	)
}
