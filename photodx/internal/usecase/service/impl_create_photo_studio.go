package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) CreatePhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
	name string,
) (*entity.PhotoStudio, error) {
	photoStudio := entity.PhotoStudio{
		ID:     photoStudioID,
		Name:   name,
		Active: false,
	}
	if _, err := t.Repository.GetPhotoStudio(ctx, photoStudioID); err == nil {
		return nil, terrors.Wrap(&repository.DuplicateEntryError{EntryType: "PhotoStudio"})
	}
	created, err := t.Repository.CreatePhotoStudio(ctx, &photoStudio)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}
