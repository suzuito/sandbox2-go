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
	) (*entity.User, error)
}
