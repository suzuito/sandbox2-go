package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreateFileUploaded(
	ctx context.Context,
	fileName string,
	fileType entity.FileType,
	input io.Reader,
) (*entity.FileUploaded, error) {
	fileID := entity.FileUploadedID(uuid.New().String())
	if fileName == "" {
		fileName = string(fileID)
	}
	file := entity.FileUploaded{
		ID:            fileID,
		Name:          fileName,
		Type:          fileType,
		ProcessStatus: entity.FileProcessStatusRegistered,
	}
	// TODO Impl transaction
	if err := t.RepositoryFileUploaded.Create(ctx, &file); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.StorageFileUploaded.Put(ctx, file.ID, input); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}
