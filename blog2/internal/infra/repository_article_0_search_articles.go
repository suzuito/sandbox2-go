package infra

import (
	"context"
	"strings"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) SearchArticles(ctx context.Context, q *entity.ArticleSearchQuery) ([]*entity.Article, *int, *int, error) {
	offset := 0
	if q.Offset != nil {
		offset = *q.Offset
	}
	limit := 10
	if q.Limit != nil {
		limit = *q.Limit
	}
	where := []string{}
	args := []any{}
	if q.TagID != nil {
		where = append(where, "MATCH(`tags`) AGAINST(? IN BOOLEAN MODE)")
		args = append(args, *q.TagID)
	}
	if q.Published != nil {
		where = append(where, "`published` = ?")
		args = append(args, q.Published)
	}
	if q.PublishedAtStart != nil {
		where = append(where, "`published_at` >= FROM_UNIXTIME(?)")
		args = append(args, q.PublishedAtStart.Unix())
	}
	if q.PublishedAtEnd != nil {
		where = append(where, "`published_at` < FROM_UNIXTIME(?)")
		args = append(args, q.PublishedAtEnd.Unix())
	}
	if q.CreatedAtStart != nil {
		where = append(where, "`created_at` >= FROM_UNIXTIME(?)")
		args = append(args, q.CreatedAtStart.Unix())
	}
	if q.CreatedAtEnd != nil {
		where = append(where, "`created_at` < FROM_UNIXTIME(?)")
		args = append(args, q.CreatedAtEnd.Unix())
	}
	queryString := "SELECT `article_id` FROM `articles_search_index`"
	if len(where) > 0 {
		queryString += " WHERE " + strings.Join(where, " AND ")
	}
	queryString += " ORDER BY `created_at` DESC "
	queryString += " LIMIT ? "
	args = append(args, limit)
	queryString += " OFFSET ? "
	args = append(args, offset)
	rows, err := t.Pool.QueryContext(ctx, queryString, args...)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	defer rows.Close()
	articleIDs := []entity.ArticleID{}
	for rows.Next() {
		articleID := entity.ArticleID("")
		if err := rows.Scan(&articleID); err != nil {
			return nil, nil, nil, terrors.Wrap(err)
		}
		articleIDs = append(articleIDs, articleID)
	}
	articles, err := getArticles(ctx, t.Pool, articleIDs...)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}

	var nextOffset *int
	var prevOffset *int
	if len(articleIDs) >= limit {
		nextOffsetValue := offset + len(articleIDs)
		nextOffset = &nextOffsetValue
	}
	if offset > 0 {
		prevOffsetValue := offset - len(articleIDs)
		prevOffset = &prevOffsetValue
	}
	return articles, prevOffset, nextOffset, nil
}
