package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_usecase "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/usecase"
)

type InputAPIPostPhotoStudioMessages struct {
	Text string `json:"text"`
}

func (t *Impl) APIPostPhotoStudioMessages(
	ctx context.Context,
	principal common_entity.UserPrincipalAccessToken,
	photoStudioID common_entity.PhotoStudioID,
	input *InputAPIPostPhotoStudioMessages,
	skipPushMessage bool,
) (*common_entity.ChatMessageWrapper, error) {
	return common_usecase.PostChatMessage(
		ctx,
		t.L,
		t.NowFunc,
		t.AuthBusinessLogic,
		t.AuthUserBusinessLogic,
		t.AdminBusinessLogic,
		string(principal.GetUserID()),
		common_entity.ChatMessagePostedByTypeUser,
		photoStudioID,
		principal.GetUserID(),
		input.Text,
		!skipPushMessage,
	)
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
	offset int,
) (*DTOAPIGetPhotoStudioMessages, error) {
	listQuery := cgorm.ListQuery{
		Offset: offset,
		Limit:  30,
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
	messages, hasNext, err := t.AdminBusinessLogic.GetChatMessages(
		ctx,
		photoStudioID,
		principal.GetUserID(),
		&listQuery,
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
