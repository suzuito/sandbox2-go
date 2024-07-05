package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type DTOAPIMiddlewarePhotoStudio struct {
	PhotoStudio *entity.PhotoStudio
}

func (t *Impl) APIMiddlewarePhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
) (*DTOAPIMiddlewarePhotoStudio, error) {
	/*
		photoStudio, err := t.B.GetPhotoStudio(ctx, photoStudioID)
		if err != nil {
			t.L.Error("", "err", err)
			return nil, terrors.Wrap(err)
		}
		return &DTOAPIMiddlewarePhotoStudio{PhotoStudio: photoStudio}, nil
	*/
	return nil, terrors.Wrapf("not impl")
}
