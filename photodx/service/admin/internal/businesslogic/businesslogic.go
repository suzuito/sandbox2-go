package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/admin/internal/repository"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type BusinessLogic interface {
	GetActiveLineLink(
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
	) (*entity.LineLinkInfo, error)
	SetLineLinkInfo(
		ctx context.Context,
		photoStudioID common_entity.PhotoStudioID,
		arg *repository.SetLineLinkInfoArgument,
	) (*entity.LineLinkInfo, error)
}
