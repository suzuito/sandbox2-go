package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	common_entity "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOMiddlewarePhotoStudio struct {
	PhotoStudio *common_entity.PhotoStudio
}

func (t *Impl) MiddlewarePhotoStudio(
	ctx context.Context,
	photoStudioID common_entity.PhotoStudioID,
) (*DTOMiddlewarePhotoStudio, error) {
	photoStudio, err := t.AuthBusinessLogic.GetPhotoStudio(ctx, photoStudioID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOMiddlewarePhotoStudio{
		PhotoStudio: photoStudio,
	}, nil
}
