package infra

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type StorageFileThumbnail struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageFileThumbnail) filePath(file *entity.FileThumbnail) string {
	return string(file.ID)
}

func (t *StorageFileThumbnail) Put(
	ctx context.Context,
	file *entity.FileThumbnail,
	r io.Reader,
) error {
	filePath := t.filePath(file)
	w := t.Cli.Bucket(t.Bucket).Object(filePath).NewWriter(ctx)
	w.ContentType = file.MediaType
	_, err := io.Copy(w, r)
	if err != nil {
		return terrors.Wrapf("io.Copy is failed on %s/%s : %w", t.Bucket, filePath, err)
	}
	if err := w.Close(); err != nil {
		return terrors.Wrapf("w.Close is failed on %s/%s : %w", t.Bucket, filePath, err)
	}
	return nil
}

func (t *StorageFileThumbnail) Get(
	ctx context.Context,
	file *entity.FileThumbnail,
	w io.Writer,
) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(t.filePath(file)).NewReader(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer reader.Close()
	if _, err := io.Copy(w, reader); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
