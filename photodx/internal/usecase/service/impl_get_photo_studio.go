package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) GetPhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
) (*entity.PhotoStudio, error) {
	return t.Repository.GetPhotoStudio(ctx, photoStudioID)
}
