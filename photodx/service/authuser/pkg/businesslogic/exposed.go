package businesslogic

import (
	"context"
	"log/slog"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type ExposedBusinessLogic interface {
	CreateUserIfNotExistsFromResourceOwnerID(
		ctx context.Context,
		providerID string,
		resourceOwnerID string,
		user *common_entity.User,
	) (*common_entity.User, error)
	GetUsers(
		ctx context.Context,
		userIDs []common_entity.UserID,
	) ([]*common_entity.User, error)
	PushNotification(
		ctx context.Context,
		l *slog.Logger,
		userID common_entity.UserID,
		notification *common_entity.Notification,
	) error
}
