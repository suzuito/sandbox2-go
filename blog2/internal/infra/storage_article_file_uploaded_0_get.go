package infra

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *StorageArticleFileUploaded) Get(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, w io.Writer) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(t.filePath(articleID, fileID)).NewReader(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer reader.Close()
	if _, err := io.Copy(w, reader); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *StorageArticleFileUploaded) GetReader(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, f func(r io.Reader) error) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(t.filePath(articleID, fileID)).NewReader(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer reader.Close()
	return f(reader)
}
