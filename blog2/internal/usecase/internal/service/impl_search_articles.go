package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

func (t *Impl) SearchArticles(ctx context.Context, q *entity.ArticleSearchQuery) ([]*entity.Article, *int, *int, error) {
	return t.RepositoryArticle.SearchArticles(ctx, q)
}
