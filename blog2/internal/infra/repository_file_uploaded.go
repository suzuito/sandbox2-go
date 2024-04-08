package infra

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type RepositoryFileUploaded struct {
	Cli *firestore.Client
}

func (t *RepositoryFileUploaded) Create(
	ctx context.Context,
	file *entity.FileUploaded,
) error {
	return t.Cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		dr := t.Cli.Collection("Blog2Files").Doc(string(file.ID))
		if err := tx.Set(dr, file); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	})
}

func (t *RepositoryFileUploaded) Get(
	ctx context.Context,
	fileID entity.FileUploadedID,
) (*entity.FileUploaded, error) {
	dr := t.Cli.Collection("Blog2Files").Doc(string(fileID))
	snp, err := dr.Get(ctx)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	file := entity.FileUploaded{}
	if err := snp.DataTo(&file); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}

func (t *RepositoryFileUploaded) StartProcess(
	ctx context.Context,
	fileID entity.FileUploadedID,
	f func() error,
) (*entity.FileUploaded, error) {
	dr := t.Cli.Collection("Blog2Files").Doc(string(fileID))
	snp, err := dr.Get(ctx)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	file := entity.FileUploaded{}
	if err := snp.DataTo(&file); err != nil {
		return nil, terrors.Wrap(err)
	}
	if errF := f(); errF != nil {
		file.ProcessStatus = entity.FileProcessStatusError
		if _, err := dr.Set(ctx, &file); err != nil {
			return nil, terrors.Wrap(err)
		}
		return nil, terrors.Wrap(errF)
	}
	file.ProcessStatus = entity.FileProcessStatusDone
	if _, err := dr.Set(ctx, &file); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}
