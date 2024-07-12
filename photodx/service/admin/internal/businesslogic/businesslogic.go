package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	GetLineLink(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*entity.LineLinkInfo, error)
	ActivateLineLink(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) (*entity.LineLinkInfo, error)
	DeactivateLineLink(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
	) error
	SetLineLinkInfoMessagingAPIChannelSecret(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		secret string,
	) (*entity.LineLinkInfo, error)
}
