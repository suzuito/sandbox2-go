package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	GenerateUserFromLINEProfile(
		ctx context.Context,
		lineLinkInfo *entity.LineLinkInfo,
		lineUserID string,
	) (*common_entity.User, error)
	CreatePhotoStudioUser(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		user *common_entity.User,
	) (*entity.PhotoStudioUser, error)
	GetPhotoStudioUsers(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		q *cgorm.ListQuery,
	) ([]*entity.PhotoStudioUser, bool, error)
	GetPhotoStudioUser(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
	) (*entity.PhotoStudioUser, error)

	GetPhotoStudioChats(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.ChatRoom, bool, error)
	GetPhotoStudioChat(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
	) (*common_entity.ChatRoom, error)

	CreateChatMessage(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		message *common_entity.ChatMessage,
	) (*common_entity.ChatMessage, error)
	GetChatMessages(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.ChatMessage, bool, error)
}
