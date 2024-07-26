package usecase

import (
	"context"

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
