package infra

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

func (t *RepositoryArticleIndex) Search(ctx context.Context, query *entity.ArticleSearchQuery) ([]*entity.ArticleSearchIndex, error) {
	// TODO impl
	return []*entity.ArticleSearchIndex{}, nil
}
