package service

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) PutArticleMarkdown(ctx context.Context, articleID entity.ArticleID, markdownBodyReader io.Reader) error {
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		return terrors.Wrapf("Document %s is not found", articleID)
	}
	if err := t.StorageArticle.PutArticle(ctx, articleID, markdownBodyReader); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
