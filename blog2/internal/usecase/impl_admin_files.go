package usecase

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOPostAdminFiles struct {
}

func (t *Impl) PostAdminFiles(
	ctx context.Context,
	fileType entity.FileType,
	input io.Reader,
) (*DTOPostAdminFiles, error) {
	file := entity.FileUploaded{
		ID:            entity.FileUploadedID(uuid.New().String()),
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
	if err := t.FunctionTriggerStartFileUploadedProcess.Publish(ctx, &file); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPostAdminFiles{}, nil
}
