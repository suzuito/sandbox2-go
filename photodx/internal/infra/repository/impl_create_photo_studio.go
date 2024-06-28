package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) CreatePhotoStudio(
	ctx context.Context,
	photoStudio *entity.PhotoStudio,
) (*entity.PhotoStudio, error) {
	if err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		if _, err := csql.ExecContext(
			ctx,
			tx,
			"INSERT INTO `photo_studios`(`id`, `name`, `active`, `created_at`, `updated_at`) VALUES (?,?,?,NOW(), NOW())",
			photoStudio.ID,
			photoStudio.Name,
			photoStudio.Active,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return nil, terrors.Wrap(err)
	}
	created, err := t.GetPhotoStudio(ctx, photoStudio.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return created, nil
}
