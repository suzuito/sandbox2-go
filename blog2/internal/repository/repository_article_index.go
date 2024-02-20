package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryArticleIndex interface {
	Search(
		ctx context.Context,
		query *entity.ArticleSearchQuery,
	) ([]*entity.ArticleSearchIndex, error)
}
