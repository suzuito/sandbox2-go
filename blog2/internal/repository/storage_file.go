package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageFile interface {
	Put(
		ctx context.Context,
		file *entity.File,
		r io.Reader,
	) error
	PutThumbnail(
		ctx context.Context,
		file *entity.FileThumbnail,
		r io.Reader,
	) error
}
