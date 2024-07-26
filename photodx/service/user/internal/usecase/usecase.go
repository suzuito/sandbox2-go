package usecase

import (
	"context"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type Usecase interface {
	MiddlewareAccessTokenAuthe(
		ctx context.Context,
		accessToken string,
	) (*DTOMiddlewareAccessTokenAuthe, error)
	MiddlewarePhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*DTOMiddlewarePhotoStudio, error)
	APIPostPhotoStudioMessages(
		ctx context.Context,
		principal common_entity.UserPrincipalAccessToken,
		photoStudioID common_entity.PhotoStudioID,
		input *InputAPIPostPhotoStudioMessages,
		skipPushMessage bool,
	) (*common_entity.ChatMessageWrapper, error)
	APIGetOlderPhotoStudioChatMessages(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		offset int,
	) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error)
	APIGetOlderPhotoStudioChatMessagesByID(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		chatMessageID common_entity.ChatMessageID,
	) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error)
}
