package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *Impl) APIGetPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	offset int,
) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error) {
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
	return &common_entity.ListResponse[*common_entity.ChatMessageWrapper]{
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
	if err := t.AuthUserBusinessLogic.PushNotification(ctx, t.L, userID, chatMessage.Text); err != nil {
		t.L.Warn("", "err", err)
	}
	if err := t.AuthBusinessLogic.PushNotificationToAllMembers(ctx, t.L, photoStudioID, chatMessage.Text); err != nil {
		t.L.Warn("", "err", err)
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
