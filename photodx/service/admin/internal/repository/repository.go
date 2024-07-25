package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/cgorm"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type SetLineLinkInfoArgument struct {
	MessagingAPIChannelSecret *string `json:"messagingApiChannelSecret"`
	LongAccessToken           *string `json:"longAccessToken"`
}

type Repository interface {
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

	CreatePhotoStudioUserChatRoomIFNotExists(
		ctx context.Context,
		room *common_entity.ChatRoom,
	) (*common_entity.ChatRoom, error)
	GetChatRoomByPhotoStudioIDANDUserID(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		userID common_entity.UserID,
	) (*common_entity.ChatRoom, error)
	GetPhotoStudioChats(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.ChatRoom, bool, error)

	CreateChatMessage(
		ctx context.Context,
		roomID common_entity.ChatRoomID,
		message *common_entity.ChatMessage,
	) (*common_entity.ChatMessage, error)
	GetChatMessages(
		ctx context.Context,
		roomID common_entity.ChatRoomID,
		listQuery *cgorm.ListQuery,
	) ([]*common_entity.ChatMessage, bool, error)
	GetChatMessagesByTimeRange(
		ctx context.Context,
		roomID common_entity.ChatRoomID,
		offset *common_entity.ChatMessageOffset,
		limit int,
		isGetOlder bool,
	) ([]*common_entity.ChatMessage, bool, error)
	GetOlderChatMessages(
		ctx context.Context,
		roomID common_entity.ChatRoomID,
		offset int,
		limit int,
	) ([]*common_entity.ChatMessage, bool, int, error)
}
