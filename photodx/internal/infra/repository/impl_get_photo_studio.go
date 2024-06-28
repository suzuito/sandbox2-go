package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) GetPhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
) (*entity.PhotoStudio, error) {
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id`, `name` FROM `photo_studios` WHERE `id` = ?",
		photoStudioID,
	)
	photoStudio := entity.PhotoStudio{}
	if err := row.Scan(
		&photoStudio.ID,
		&photoStudio.Name,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, terrors.Wrap(&repository.NoEntryError{
				EntryType: "PhotoStudio",
				EntryID:   string(photoStudioID),
			})
		}
		return nil, terrors.Wrap(err)
	}
	return &photoStudio, nil
}
