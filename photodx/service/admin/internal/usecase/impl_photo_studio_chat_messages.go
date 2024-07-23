package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIGetPhotoStudioChatMessages struct {
	Results    []*common_entity.ChatMessageWrapper
	HasNext    bool
	HasPrev    bool
	NextOffset int
	PrevOffset int
}

func (t *Impl) APIGetPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	offset int,
) (*DTOAPIGetPhotoStudioChatMessages, error) {
	listQuery := cgorm.ListQuery{
		Offset: offset,
		Limit:  30,
	}
	chatMessages, hasNext, err := t.BusinessLogic.GetChatMessages(
		ctx,
		photoStudioID,
		userID,
		&listQuery,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	chatMessageWrappers, err := common_entity.BuildChatMessageWrapper(
		ctx,
		chatMessages,
		t.AuthUserBusinessLogic.GetUsers,
		t.AuthBusinessLogic.GetPhotoStudioMembers,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetPhotoStudioChatMessages{
		Results:    chatMessageWrappers,
		HasNext:    hasNext,
		HasPrev:    listQuery.HasPrev(),
		NextOffset: listQuery.NextOffset(),
		PrevOffset: listQuery.PrevOffset(),
	}, nil
}

func (t *Impl) APIPostPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	photoStudioMemberID common_entity.PhotoStudioMemberID,
	text string,
) (*common_entity.ChatMessageWrapper, error) {
	message := common_entity.ChatMessage{
		Type:         common_entity.ChatMessageTypeText,
		Text:         text,
		PostedBy:     string(photoStudioMemberID),
		PostedByType: common_entity.ChatMessagePostedByTypePhotoStudioMember,
		PostedAt:     common_entity.WTime(t.NowFunc()),
	}
	chatMessage, err := t.BusinessLogic.CreateChatMessage(
		ctx,
		photoStudioID,
		userID,
		&message,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	chatMessageWrappers, err := common_entity.BuildChatMessageWrapper(
		ctx,
		[]*common_entity.ChatMessage{
			chatMessage,
		},
		t.AuthUserBusinessLogic.GetUsers,
		t.AuthBusinessLogic.GetPhotoStudioMembers,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return chatMessageWrappers[0], nil
}
