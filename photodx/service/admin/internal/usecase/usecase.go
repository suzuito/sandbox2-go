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

	APIGetPhotoStudioUsers(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
		offset int,
	) (*DTOAPIGetPhotoStudioUsers, error)
	APIGetPhotoStudioUser(
		ctx context.Context,
		principal common_entity.AdminPrincipalAccessToken,
		userID common_entity.UserID,
	) (*DTOPhotoStudioUser, error)

	APIGetPhotoStudioChats(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		offset int,
	) (*common_entity.ListResponse[*common_entity.ChatRoomWrapper], error)
	APIGetPhotoStudioChat(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
	) (*common_entity.ChatRoomWrapper, error)

	APIGetPhotoStudioChatMessages(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		offset int,
	) (*common_entity.ListResponse[*common_entity.ChatMessageWrapper], error)
	APIPostPhotoStudioChatMessages(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
		text string,
	) (*common_entity.ChatMessageWrapper, error)

	APIPostSuperInit(ctx context.Context) (*DTOAPIPostSuperInit, error)
}
