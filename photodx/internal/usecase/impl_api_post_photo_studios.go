package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

type DTOAPIPostPhotoStudios struct {
	PhotoStudio *entity.PhotoStudio
}

func (t *Impl) APIPostPhotoStudios(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	name string,
) (*DTOAPIPostPhotoStudios, error) {
	photoStudio, err := t.S.CreatePhotoStudio(ctx, photoStudioID, name)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostPhotoStudios{PhotoStudio: photoStudio}, nil
}
