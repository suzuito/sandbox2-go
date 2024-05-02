package infra

import (
	"context"
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
	return string(file.ID)
}

func (t *StorageFile) Put(
	ctx context.Context,
	file *entity.File,
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

func (t *StorageFile) Get(
	ctx context.Context,
	file *entity.File,
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
