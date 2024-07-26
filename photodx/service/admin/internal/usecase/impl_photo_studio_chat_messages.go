package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_usecase "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/usecase"
)

func (t *Impl) APIGetOlderPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	offset int,
) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error) {
	chatMessages, hasNext, nextOffset, hasPrev, prevOffset, err := t.BusinessLogic.GetOlderChatMessagesForFront(
		ctx,
		photoStudioID,
		userID,
		offset,
		30,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return t.buildChatMessageWrapper(ctx, chatMessages, hasNext, nextOffset, hasPrev, prevOffset)
}

func (t *Impl) APIGetOlderPhotoStudioChatMessagesByID(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	chatMessageID common_entity.ChatMessageID,
) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error) {
	chatMessages, hasNext, nextOffset, hasPrev, prevOffset, err := t.BusinessLogic.GetOlderChatMessagesForFrontByID(
		ctx,
		photoStudioID,
		userID,
		chatMessageID,
		30,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return t.buildChatMessageWrapper(ctx, chatMessages, hasNext, nextOffset, hasPrev, prevOffset)
}

func (t *Impl) buildChatMessageWrapper(
	ctx context.Context,
	chatMessages []*common_entity.ChatMessage,
	hasNext bool,
	nextOffset int,
	hasPrev bool,
	prevOffset int,
) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error) {
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
		NextOffset: nextOffset,
		HasPrev:    hasPrev,
		PrevOffset: prevOffset,
	}, nil
}

func (t *Impl) APIPostPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	photoStudioMemberID common_entity.PhotoStudioMemberID,
	text string,
	skipPushMessage bool,
) (*common_entity.ChatMessageWrapper, error) {
	return common_usecase.PostChatMessage(
		ctx,
		t.L,
		t.NowFunc,
		t.AuthBusinessLogic,
		t.AuthUserBusinessLogic,
		t.BusinessLogic,
		string(photoStudioMemberID),
		common_entity.ChatMessagePostedByTypePhotoStudioMember,
		photoStudioID,
		userID,
		text,
		!skipPushMessage,
	)
}
