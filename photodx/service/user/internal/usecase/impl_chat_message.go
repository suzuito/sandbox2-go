package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type InputAPIPostPhotoStudioMessages struct {
	Text string `json:"text"`
}

func (t *Impl) APIPostPhotoStudioMessages(
	ctx context.Context,
	principal common_entity.UserPrincipalAccessToken,
	photoStudioID common_entity.PhotoStudioID,
	input *InputAPIPostPhotoStudioMessages,
) (*common_entity.ChatMessageWrapper, error) {
	msg := common_entity.ChatMessage{
		Type:         common_entity.ChatMessageTypeText,
		Text:         input.Text,
		PostedBy:     string(principal.GetUserID()),
		PostedByType: common_entity.ChatMessagePostedByTypeUser,
		PostedAt:     common_entity.WTime(t.NowFunc()),
	}
	if _, err := t.AdminBusinessLogic.CreatePhotoStudioUserChatRoomIFNotExists(
		ctx,
		photoStudioID,
		principal.GetUserID(),
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	created, err := t.AdminBusinessLogic.CreateChatMessage(
		ctx,
		photoStudioID,
		principal.GetUserID(),
		&msg,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.AuthUserBusinessLogic.PushNotification(ctx, t.L, principal.GetUserID(), created.Text); err != nil {
		t.L.Warn("", "err", err)
	}
	if err := t.AuthBusinessLogic.PushNotificationToAllMembers(ctx, t.L, photoStudioID, created.Text); err != nil {
		t.L.Warn("", "err", err)
	}
	a, err := common_entity.BuildChatMessageWrapper(
		ctx,
		[]*common_entity.ChatMessage{created},
		t.AuthUserBusinessLogic.GetUsers,
		t.AuthBusinessLogic.GetPhotoStudioMembers,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return a[0], nil
}

type DTOAPIGetPhotoStudioMessages struct {
	Results    []*common_entity.ChatMessageWrapper `json:"results"`
	HasNext    bool                                `json:"hasNext"`
	HasPrev    bool                                `json:"hasPrev"`
	NextOffset int                                 `json:"nextOffset"`
	PrevOffset int                                 `json:"prevOffset"`
}

func (t *Impl) APIGetPhotoStudioMessages(
	ctx context.Context,
	principal common_entity.UserPrincipalAccessToken,
	photoStudioID common_entity.PhotoStudioID,
	listQuery *cgorm.ListQuery,
) (*DTOAPIGetPhotoStudioMessages, error) {
	listQuery.Limit = 30
	if listQuery.Offset < 0 {
		listQuery.Offset = 0
	}
	listQuery.SortColumns = []cgorm.SortColumn{
		{Name: "posted_at", Type: cgorm.Asc},
	}
	messages, hasNext, err := t.AdminBusinessLogic.GetChatMessages(
		ctx,
		photoStudioID,
		principal.GetUserID(),
		listQuery,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	wmessages, err := common_entity.BuildChatMessageWrapper(
		ctx,
		messages,
		t.AuthUserBusinessLogic.GetUsers,
		t.AuthBusinessLogic.GetPhotoStudioMembers,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetPhotoStudioMessages{
		Results:    wmessages,
		HasNext:    hasNext,
		HasPrev:    listQuery.HasPrev(),
		NextOffset: listQuery.NextOffset(),
		PrevOffset: listQuery.PrevOffset(),
	}, nil
}
