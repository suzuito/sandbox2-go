package infra

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

// Deprecated
func (t *StorageArticle) CreateArticle(ctx context.Context, articleID entity.ArticleID) error {
	w := t.Cli.Bucket(t.Bucket).Object(t.filePathMarkdown(articleID)).NewWriter(ctx)
	if err := w.Close(); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
