package businesslogic

import (
	"context"

	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type ExposedBusinessLogic interface {
	GetPhotoStudio(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*common_entity.PhotoStudio, error)
}
