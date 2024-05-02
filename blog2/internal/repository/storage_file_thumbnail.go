package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageFileThumbnal interface {
	Put(
		ctx context.Context,
		file *entity.FileThumbnail,
		r io.Reader,
	) error
	Get(
		ctx context.Context,
		file *entity.FileThumbnail,
		w io.Writer,
	) error
}
