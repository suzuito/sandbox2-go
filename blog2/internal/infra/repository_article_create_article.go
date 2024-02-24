package infra

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) CreateArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*entity.Article, error) {
	_, err := t.Pool.ExecContext(
		ctx,
		"INSERT INTO `articles`(`id`, `title`, `published`, `published_at`, `created_at`, `updated_at`) VALUES (?, ?, false, NULL, NOW(), NOW())",
		articleID,
		"",
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	articles, err := t.selectArticleByID(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return articles[0], nil
}
