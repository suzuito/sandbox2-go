package businesslogic

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

func (t *BusinessLogicImpl) GetPhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
) (*entity.PhotoStudio, error) {
	return t.Repository.GetPhotoStudio(ctx, photoStudioID)
}

func (t *BusinessLogicImpl) CreatePhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	name string,
) (*entity.PhotoStudio, error) {
	photoStudio := entity.PhotoStudio{
		ID:     photoStudioID,
		Name:   name,
		Active: false,
	}
	created, err := t.Repository.CreatePhotoStudio(ctx, &photoStudio)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}
