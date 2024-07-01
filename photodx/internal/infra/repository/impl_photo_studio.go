package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) CreatePhotoStudio(
	ctx context.Context,
	photoStudio *entity.PhotoStudio,
) (*entity.PhotoStudio, error) {
	photoStudio, err := t.GetPhotoStudio(ctx, photoStudio.ID)
	if err == nil {
		return nil, terrors.Wrap(&repository.DuplicateEntryError{
			EntryType: repository.EntryTypePhotoStudio,
			EntryID:   string(photoStudio.ID),
		})
	}
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

func (t *Impl) GetPhotoStudio(
	ctx context.Context,
	photoStudioID entity.PhotoStudioID,
) (*entity.PhotoStudio, error) {
	photoStudios, err := getPhotoStudios(ctx, t.Pool, []entity.PhotoStudioID{photoStudioID})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if len(photoStudios) <= 0 {
		return nil, terrors.Wrap(&repository.NoEntryError{
			EntryType: repository.EntryTypePhotoStudio,
			EntryID:   string(photoStudioID),
		})
	}
	return photoStudios[0], nil
}

func getPhotoStudios(
	ctx context.Context,
	txOrDB csql.TxOrDB,
	photoStudioIDs []entity.PhotoStudioID,
) ([]*entity.PhotoStudio, error) {
	rows, err := csql.QueryContext(
		ctx,
		txOrDB,
		"SELECT `id`, `name`, `active` FROM `photo_studios` WHERE "+csql.SqlIn("id", photoStudioIDs),
		csql.ToAnySlice(photoStudioIDs)...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rows.Close()
	photoStudios := []*entity.PhotoStudio{}
	for rows.Next() {
		photoStudio := entity.PhotoStudio{}
		if err := rows.Scan(
			&photoStudio.ID,
			&photoStudio.Name,
			&photoStudio.Active,
		); err != nil {
			return nil, terrors.Wrap(err)
		}
		photoStudios = append(photoStudios, &photoStudio)
	}
	return photoStudios, nil
}
