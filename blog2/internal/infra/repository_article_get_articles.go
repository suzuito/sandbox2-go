package infra

import (
	"context"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

func (t *RepositoryArticle) GetArticles(ctx context.Context, ids ...entity.ArticleID) ([]*entity.Article, error) {
	now := time.Now()
	// TODO impl
	return []*entity.Article{
		{
			ID:          "id01",
			Title:       "title01",
			Summary:     "summary01",
			Published:   true,
			PublishedAt: &now,
			Tags: []entity.Tag{
				{ID: "tag01", Name: "tag01name"},
			},
		},
		{
			ID:          "id02",
			Title:       "title02",
			Summary:     "summary02",
			Published:   false,
			PublishedAt: &now,
			Tags: []entity.Tag{
				{ID: "tag01", Name: "tag01name"},
			},
		},
		{
			ID:          "id03",
			Title:       "title03",
			Summary:     "summary03",
			Published:   true,
			PublishedAt: &now,
			Tags: []entity.Tag{
				{ID: "tag01", Name: "tag01name"},
			},
		},
	}, nil
}
