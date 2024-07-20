package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type ExposedBusinessLogic interface {
	CreateUserIfNotExistsFromResourceOwnerID(
		ctx context.Context,
		providerID string,
		resourceOwnerID string,
		user *entity.User,
	) (*entity.User, error)
	GetUsers(
		ctx context.Context,
		userIDs []entity.UserID,
	) ([]*entity.User, error)
}
