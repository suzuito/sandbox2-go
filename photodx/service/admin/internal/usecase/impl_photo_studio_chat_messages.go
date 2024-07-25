package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_usecase "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/usecase"
)

func (t *Impl) APIGetPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	offset int,
) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error) {
	listQuery := cgorm.ListQuery{
		Offset: offset,
		// Limit:  30000,
		Limit: 30, // debug
		SortColumns: []cgorm.SortColumn{
			{
				Name: "posted_at",
				Type: cgorm.Desc,
			},
			{
				Name: "id",
				Type: cgorm.Desc,
			},
		},
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

func (t *Impl) APIGetOlderPhotoStudioChatMessages(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
	userID common_entity.UserID,
	offset int,
) (*common_entity.ListResponse2[*common_entity.ChatMessageWrapper], error) {
	chatMessages, hasNext, nextOffset, err := t.BusinessLogic.GetOlderChatMessages(
		ctx,
		photoStudioID,
		userID,
		offset,
		30,
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
	return &common_entity.ListResponse2[*common_entity.ChatMessageWrapper]{
		Results:    chatMessageWrappers,
		HasNext:    hasNext,
		NextOffset: nextOffset,
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
