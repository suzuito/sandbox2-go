package infra

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type StorageFile struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageFile) filePath(file *entity.File) string {
	return fmt.Sprintf("%s%s", file.ID, file.ExtIncludingDot())
}

func (t *StorageFile) filePathThumbnail(file *entity.FileThumbnail) string {
	return fmt.Sprintf("%s_thumbnail%s", file.ID, file.ExtIncludingDot())
}

func (t *StorageFile) put(
	ctx context.Context,
	filePath string,
	mediaType string,
	r io.Reader,
) error {
	w := t.Cli.Bucket(t.Bucket).Object(filePath).NewWriter(ctx)
	if mediaType != "" {
		w.ContentType = mediaType
	}
	_, err := io.Copy(w, r)
	if err != nil {
		return terrors.Wrapf("io.Copy is failed on %s/%s : %w", t.Bucket, filePath, err)
	}
	if err := w.Close(); err != nil {
		return terrors.Wrapf("w.Close is failed on %s/%s : %w", t.Bucket, filePath, err)
	}
	return nil
}

func (t *StorageFile) Put(
	ctx context.Context,
	file *entity.File,
	r io.Reader,
) error {
	return terrors.Wrap(t.put(
		ctx,
		t.filePath(file),
		file.MediaType,
		r,
	))
}

func (t *StorageFile) PutThumbnail(
	ctx context.Context,
	file *entity.FileThumbnail,
	r io.Reader,
) error {
	return terrors.Wrap(t.put(
		ctx,
		t.filePathThumbnail(file),
		file.MediaType,
		r,
	))
}
