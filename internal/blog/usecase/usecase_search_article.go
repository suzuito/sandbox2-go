package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func (t *UsecaseImpl) SearchArticles(
	ctx context.Context,
	query SearchArticlesQuery,
	articles *[]entity.Article,
	hasNext *bool,
) error {
	articlePrimaryKeys := []entity.ArticlePrimaryKey{}
	if err := t.RepositoryArticle.SearchArticles(ctx, query, &articlePrimaryKeys, hasNext); err != nil {
		return err
	}
	if len(articlePrimaryKeys) <= 0 {
		return nil
	}
	if err := t.RepositoryArticle.GetArticlesByPrimaryKey(
		ctx, articlePrimaryKeys, query.SortField, query.SortOrder, articles); err != nil {
		return err
	}
	return nil
}
