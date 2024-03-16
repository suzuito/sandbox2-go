package infra

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *StorageArticle) PutArticle(ctx context.Context, articleID entity.ArticleID, r io.Reader) error {
	w := t.Cli.Bucket(t.Bucket).Object(t.filePathMarkdown(articleID)).NewWriter(ctx)
	_, err := io.Copy(w, r)
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := w.Close(); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
