package infra

import (
	"context"
	"database/sql"
	"strings"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type RepositoryArticle struct {
	Pool *sql.DB
}

func (t *RepositoryArticle) selectArticleByID(
	ctx context.Context,
	articleIDs ...entity.ArticleID,
) ([]*entity.Article, error) {
	articleID := ""
	title := ""
	published := false
	publishedAt := sql.NullTime{}
	whereStatement := "WHERE `id` IN (?" + strings.Repeat(",?", len(articleIDs)-1) + ")"
	anyArticleIDs := make([]any, len(articleIDs))
	for i, aID := range articleIDs {
		anyArticleIDs[i] = aID
	}
	rows, err := t.Pool.QueryContext(
		ctx,
		"SELECT `id`,`title`,`published`,`published_at` FROM `articles` "+whereStatement,
		anyArticleIDs...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	articles := []*entity.Article{}
	for rows.Next() {
		article := entity.Article{}
		if err := rows.Scan(&articleID, &title, &published, &publishedAt); err != nil {
			return nil, terrors.Wrap(err)
		}
		article.ID = entity.ArticleID(articleID)
		article.Title = title
		article.Published = published
		if publishedAt.Valid {
			article.PublishedAt = &publishedAt.Time
		}
		articles = append(articles, &article)
	}
	if err := rows.Err(); err != nil {
		return nil, terrors.Wrap(err)
	}

	return articles, nil
}
