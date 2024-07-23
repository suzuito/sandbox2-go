package businesslogic

import (
	"context"
	"log/slog"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type ExposedBusinessLogic interface {
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*common_entity.PhotoStudio, error)
	GetPhotoStudios(
		ctx context.Context,
		photoStudioIDs []common_entity.PhotoStudioID,
	) ([]*common_entity.PhotoStudio, error)
	GetPhotoStudioMembers(
		ctx context.Context,
		photoStudioMemberIDs []common_entity.PhotoStudioMemberID,
	) ([]*common_entity.PhotoStudioMemberWrapper, error)
	CreatePhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		name string,
	) (*common_entity.PhotoStudio, error)
	CreatePhotoStudioMember(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		email string,
		name string,
	) (*common_entity.PhotoStudioMember, []*rbac.Role, *common_entity.PhotoStudio, string, error)
	PushNotification(
		ctx context.Context,
		l *slog.Logger,
		photoStudioMemberID common_entity.PhotoStudioMemberID,
		message string,
	) error
	PushNotificationToAllMembers(
		ctx context.Context,
		l *slog.Logger,
		photoStudioID common_entity.PhotoStudioID,
		message string,
	) error
}
