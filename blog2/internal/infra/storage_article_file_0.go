package infra

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type StorageArticleFile struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageArticleFile) filePath(articleID entity.ArticleID, fileID entity.ArticleFileUploadedID) string {
	return fmt.Sprintf("%s/%s", articleID, fileID)
}

func (t *StorageArticleFile) Get(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, w io.Writer) error {
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

func (t *StorageArticleFile) Put(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, r io.Reader) error {
	w := t.Cli.Bucket(t.Bucket).Object(t.filePath(articleID, fileID)).NewWriter(ctx)
	_, err := io.Copy(w, r)
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := w.Close(); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
