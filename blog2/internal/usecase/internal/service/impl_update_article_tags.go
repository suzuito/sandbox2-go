package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

func (t *Impl) UpdateArticleTags(
	ctx context.Context,
	articleID entity.ArticleID,
	add []entity.TagID,
	delete []entity.TagID,
) (*entity.Article, error) {
	return t.RepositoryArticle.UpdateArticleTags(ctx, articleID, add, delete)
}
