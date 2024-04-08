package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageFileUploaded interface {
	Put(ctx context.Context, fileID entity.FileUploadedID, r io.Reader) error
	Get(
		ctx context.Context,
		fileID entity.FileUploadedID,
		w io.Writer,
	) error
	GetReader(
		ctx context.Context,
		fileID entity.FileUploadedID,
		f func(r io.Reader) error,
	) error
}
