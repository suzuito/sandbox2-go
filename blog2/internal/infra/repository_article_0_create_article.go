package infra

import (
	"context"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) CreateArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	now time.Time,
) (*entity.Article, error) {
	var article *entity.Article
	err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		_, err := execContext(
			ctx,
			tx,
			"INSERT INTO `articles`(`id`, `title`, `published`, `published_at`, `created_at`, `updated_at`) VALUES (?, ?, false, NULL, FROM_UNIXTIME(?), FROM_UNIXTIME(?))",
			articleID,
			"",
			now.Unix(),
			now.Unix(),
		)
		if err != nil {
			return terrors.Wrap(err)
		}
		articles, err := getArticles(ctx, tx, articleID)
		if err != nil {
			return terrors.Wrap(err)
		}
		article = articles[0]
		if err := updateSearchIndex(ctx, tx, article.ID); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return article, nil
}
