package service

import (
	"context"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

func (t *Impl) PutArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	title *string,
	published *bool,
	publishedAt *time.Time,
) (*entity.Article, error) {
	return t.RepositoryArticle.UpdateArticle(ctx, articleID, title, published, publishedAt)
}
