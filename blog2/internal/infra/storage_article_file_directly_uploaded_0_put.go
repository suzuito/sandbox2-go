package infra

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *StorageArticleFileDirectlyUploaded) Put(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileDirectlyUploadedID, r io.Reader) error {
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
