package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryFileUploaded interface {
	Create(ctx context.Context, f *entity.FileUploaded) error
	StartProcess(
		ctx context.Context,
		fileID entity.FileUploadedID,
		f func() error,
	) (*entity.FileUploaded, error)
}
