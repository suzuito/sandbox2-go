package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIGetPhotoStudioChats struct {
	Results    []*common_entity.ChatRoomWrapper
	HasNext    bool
	HasPrev    bool
	NextOffset int
	PrevOffset int
}

func (t *Impl) APIGetPhotoStudioChats(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	offset int,
) (*DTOAPIGetPhotoStudioChats, error) {
	if offset < 0 {
		offset = 0
	}
	listQuery := cgorm.ListQuery{
		Offset: offset,
		Limit:  30,
	}
	chatRooms, hasNext, err := t.BusinessLogic.GetPhotoStudioChats(
		ctx,
		photoStudioID,
		&listQuery,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	chatRoomWrappers, err := common_entity.BuildChatRoomWrappers(
		ctx,
		chatRooms,
		t.AuthUserBusinessLogic.GetUsers,
		t.AuthBusinessLogic.GetPhotoStudios,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetPhotoStudioChats{
		Results:    chatRoomWrappers,
		HasNext:    hasNext,
		HasPrev:    listQuery.HasPrev(),
		PrevOffset: listQuery.PrevOffset(),
		NextOffset: listQuery.NextOffset(),
	}, nil
}

func (t *Impl) APIGetPhotoStudioChat(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
) (*common_entity.ChatRoomWrapper, error) {
	chatRoom, err := t.BusinessLogic.GetPhotoStudioChat(
		ctx,
		photoStudioID,
		userID,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	chatRoomWrappers, err := common_entity.BuildChatRoomWrappers(
		ctx,
		[]*common_entity.ChatRoom{chatRoom},
		t.AuthUserBusinessLogic.GetUsers,
		t.AuthBusinessLogic.GetPhotoStudios,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return chatRoomWrappers[0], nil
}
