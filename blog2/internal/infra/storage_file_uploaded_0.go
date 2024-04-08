package infra

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type StorageFileUploaded struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageFileUploaded) filePath(fileID entity.FileUploadedID) string {
	return fmt.Sprintf("%s", fileID)
}

func (t *StorageFileUploaded) Put(ctx context.Context, fileID entity.FileUploadedID, r io.Reader) error {
	w := t.Cli.Bucket(t.Bucket).Object(t.filePath(fileID)).NewWriter(ctx)
	_, err := io.Copy(w, r)
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := w.Close(); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *StorageFileUploaded) Get(ctx context.Context, fileID entity.FileUploadedID, w io.Writer) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(t.filePath(fileID)).NewReader(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer reader.Close()
	if _, err := io.Copy(w, reader); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *StorageFileUploaded) GetReader(ctx context.Context, fileID entity.FileUploadedID, f func(r io.Reader) error) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(t.filePath(fileID)).NewReader(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer reader.Close()
	return f(reader)
}
